package maintain

import (
	"HtmxBlog/config"
	"HtmxBlog/model"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/log"
)

// Snapshot holds all database tables for export/import.
type Snapshot struct {
	Version    string         `json:"version"`
	ExportedAt string         `json:"exported_at"`
	Tables     SnapshotTables `json:"tables"`
}

type SnapshotTables struct {
	Categories []model.Category `json:"categories"`
	Tags       []model.Tag      `json:"tags"`
	Posts      []model.Post     `json:"posts"`
	Attaches   []model.Attach   `json:"attaches"`
	Pages      []model.Page     `json:"pages"`
	Comments   []model.Comment  `json:"comments"`
	Settings   []model.Setting  `json:"settings"`
	Redirects  []model.Redirect `json:"redirects"`
}

// ExportAll reads all database records and writes them to a JSON file.
// The output file defaults to "backup.json" unless outPath is specified.
func ExportAll(outPath string) error {
	config.Init()
	config.InitDB()

	if outPath == "" {
		outPath = fmt.Sprintf("backup_%s.json", time.Now().Format("20060102_150405"))
	}

	snapshot := Snapshot{
		Version:    "1.0",
		ExportedAt: time.Now().Format(time.RFC3339),
	}

	// Read categories
	if err := config.DB.Find(&snapshot.Tables.Categories).Error; err != nil {
		return fmt.Errorf("failed to read categories: %w", err)
	}
	log.Infof("  categories: %d records\n", len(snapshot.Tables.Categories))

	// Read tags
	if err := config.DB.Find(&snapshot.Tables.Tags).Error; err != nil {
		return fmt.Errorf("failed to read tags: %w", err)
	}
	log.Infof("  tags: %d records\n", len(snapshot.Tables.Tags))

	// Read posts with associations
	if err := config.DB.Preload("Category").Preload("Tags").Preload("Attachs").Find(&snapshot.Tables.Posts).Error; err != nil {
		return fmt.Errorf("failed to read posts: %w", err)
	}
	log.Infof("  posts: %d records\n", len(snapshot.Tables.Posts))

	// Read attaches
	if err := config.DB.Find(&snapshot.Tables.Attaches).Error; err != nil {
		return fmt.Errorf("failed to read attaches: %w", err)
	}
	log.Infof("  attaches: %d records\n", len(snapshot.Tables.Attaches))

	// Read pages
	if err := config.DB.Find(&snapshot.Tables.Pages).Error; err != nil {
		return fmt.Errorf("failed to read pages: %w", err)
	}
	log.Infof("  pages: %d records\n", len(snapshot.Tables.Pages))

	// Read comments
	if err := config.DB.Find(&snapshot.Tables.Comments).Error; err != nil {
		return fmt.Errorf("failed to read comments: %w", err)
	}
	log.Infof("  comments: %d records\n", len(snapshot.Tables.Comments))

	// Read settings
	if err := config.DB.Find(&snapshot.Tables.Settings).Error; err != nil {
		return fmt.Errorf("failed to read settings: %w", err)
	}
	log.Infof("  settings: %d records\n", len(snapshot.Tables.Settings))

	// Read redirects
	if err := config.DB.Find(&snapshot.Tables.Redirects).Error; err != nil {
		return fmt.Errorf("failed to read redirects: %w", err)
	}
	log.Infof("  redirects: %d records\n", len(snapshot.Tables.Redirects))

	// Serialize to JSON
	data, err := json.MarshalIndent(snapshot, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	if err := os.WriteFile(outPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	log.Infof("\nExported successfully to: %s", outPath)
	return nil
}
