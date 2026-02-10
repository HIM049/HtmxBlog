package services

import (
	"HtmxBlog/config"
	"HtmxBlog/model"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const ATTACHES_DIR = "./app_data/attaches"

// CreateAttach handle uploads an attach file to the server.
func CreateAttach(file *multipart.File, name, mime string, postId uint) (*model.Attach, error) {
	isSuccess := false

	// get post info
	post, err := ReadPost(postId)
	if err != nil {
		return nil, err
	}

	// Generate unique ID (UID)
	uuid := uuid.New().String()
	dstPath := filepath.Join(ATTACHES_DIR, uuid)

	// create local file
	dst, err := os.Create(dstPath)
	if err != nil {
		return nil, err
	}
	defer func() {
		dst.Close()
		if !isSuccess {
			os.Remove(dstPath)
		}
	}()

	// create hasher
	hasher := sha256.New()
	writer := io.MultiWriter(dst, hasher)
	if _, err = io.Copy(writer, *file); err != nil {
		return nil, err
	}
	hash := hex.EncodeToString(hasher.Sum(nil))

	// check if attach already exists
	attach, err := ReadAttachByHash(hash)
	if err == nil {
		// duplicated

		// append reference
		attach.Refers = append(attach.Refers, *post)
		err = UpdateAttach(attach)
		if err != nil {
			return nil, err
		}
		return attach, nil
	} else {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			// if error happend and is db err
			dst.Close()
			os.Remove(dstPath)
			return nil, err
		}
	}

	// create attach record
	attach = &model.Attach{
		Hash:       hash,
		Uid:        uuid,
		Name:       name,
		Mime:       mime,
		Permission: model.VisibilityPublic,
		Refers:     []model.Post{*post},
	}
	// record to db
	err = config.DB.Create(attach).Error
	if err != nil {
		return nil, err
	}

	isSuccess = true
	return attach, nil
}

// ReadAttachById reads an attach by its ID.
func ReadAttachByUid(uid string) (*model.Attach, error) {
	var attach model.Attach
	err := config.DB.Where("uid = ?", uid).First(&attach).Error
	return &attach, err
}

// ReadAttachByHash reads an attach by its hash.
func ReadAttachByHash(hash string) (*model.Attach, error) {
	var attach model.Attach
	err := config.DB.Where("hash = ?", hash).First(&attach).Error
	return &attach, err
}

// ReadAllAttaches reads all attaches.
func ReadAllAttaches(limit, offset int) ([]model.Attach, error) {
	var attaches []model.Attach
	err := config.DB.Limit(limit).Offset(offset).Find(&attaches).Error
	return attaches, err
}

// UpdateAttach updates an attach.
func UpdateAttach(attach *model.Attach) error {
	return config.DB.Save(attach).Error
}

// DeleteAttach deletes an attach.
func DeleteAttach(id uint) error {
	return config.DB.Delete(&model.Attach{}, id).Error
}
