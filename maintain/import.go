package maintain

import (
	"HtmxBlog/config"
	"HtmxBlog/model"
	"encoding/json"
	"fmt"
	"os"

	"gorm.io/gorm"
)

// ImportAll reads a JSON snapshot file and restores all database records.
func ImportAll(filePath string) error {
	config.Init()
	config.InitDB()

	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	var snapshot Snapshot
	if err := json.Unmarshal(data, &snapshot); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	if snapshot.Version == "" {
		return fmt.Errorf("invalid snapshot file: missing version")
	}

	fmt.Printf("Importing snapshot v%s (exported at %s)\n\n", snapshot.Version, snapshot.ExportedAt)

	// Run import in a transaction
	err = config.DB.Transaction(func(tx *gorm.DB) error {
		// Clear tables in reverse dependency order
		fmt.Println("Clearing existing data...")
		if err := clearTables(tx); err != nil {
			return fmt.Errorf("failed to clear tables: %w", err)
		}

		// Insert categories
		fmt.Printf("  importing categories: %d records\n", len(snapshot.Tables.Categories))
		if err := insertAll(tx, snapshot.Tables.Categories); err != nil {
			return fmt.Errorf("failed to import categories: %w", err)
		}

		// Insert tags
		fmt.Printf("  importing tags: %d records\n", len(snapshot.Tables.Tags))
		if err := insertAll(tx, snapshot.Tables.Tags); err != nil {
			return fmt.Errorf("failed to import tags: %w", err)
		}

		// Insert attaches (without Refers to avoid circular dependency)
		fmt.Printf("  importing attaches: %d records\n", len(snapshot.Tables.Attaches))
		for _, a := range snapshot.Tables.Attaches {
			a.Refers = nil // clear Refers; will be restored via post association
			if err := tx.Create(&a).Error; err != nil {
				return fmt.Errorf("failed to import attach (id=%d): %w", a.ID, err)
			}
		}

		// Insert posts with associations
		fmt.Printf("  importing posts: %d records\n", len(snapshot.Tables.Posts))
		for _, p := range snapshot.Tables.Posts {
			post := p
			// GORM will handle post_tags and post_attaches join tables
			if err := tx.Create(&post).Error; err != nil {
				return fmt.Errorf("failed to import post (id=%d): %w", p.ID, err)
			}
		}

		// Insert pages
		fmt.Printf("  importing pages: %d records\n", len(snapshot.Tables.Pages))
		if err := insertAll(tx, snapshot.Tables.Pages); err != nil {
			return fmt.Errorf("failed to import pages: %w", err)
		}

		// Insert comments
		fmt.Printf("  importing comments: %d records\n", len(snapshot.Tables.Comments))
		if err := insertAll(tx, snapshot.Tables.Comments); err != nil {
			return fmt.Errorf("failed to import comments: %w", err)
		}

		// Insert settings
		fmt.Printf("  importing settings: %d records\n", len(snapshot.Tables.Settings))
		if err := insertAll(tx, snapshot.Tables.Settings); err != nil {
			return fmt.Errorf("failed to import settings: %w", err)
		}

		// Insert redirects
		fmt.Printf("  importing redirects: %d records\n", len(snapshot.Tables.Redirects))
		if err := insertAll(tx, snapshot.Tables.Redirects); err != nil {
			return fmt.Errorf("failed to import redirects: %w", err)
		}

		return nil
	})

	if err != nil {
		return err
	}

	fmt.Println("\nImport completed successfully.")
	return nil
}

// clearTables deletes all records from all tables in dependency order.
func clearTables(tx *gorm.DB) error {
	if err := tx.Exec("DELETE FROM post_tags").Error; err != nil {
		return err
	}
	if err := tx.Exec("DELETE FROM post_attaches").Error; err != nil {
		return err
	}

	tables := []interface{}{
		&model.Comment{},
		&model.Post{},
		&model.Attach{},
		&model.Page{},
		&model.Category{},
		&model.Tag{},
		&model.Setting{},
		&model.Redirect{},
	}

	for _, table := range tables {
		if err := tx.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(table).Error; err != nil {
			return err
		}
	}

	return nil
}

// insertAll inserts a slice of records in batches.
func insertAll[T any](tx *gorm.DB, records []T) error {
	for _, record := range records {
		if err := tx.Create(&record).Error; err != nil {
			return err
		}
	}
	return nil
}
