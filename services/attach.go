package services

import (
	"HtmxBlog/database"
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

func UploadAttach(file *multipart.File, name, mime string) (*model.Attach, error) {
	isSuccess := false

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
	attach, err := database.ReadAttachByHash(hash)
	if err == nil {
		return attach, nil
	} else {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			// if error happend and is db err
			dst.Close()
			os.Remove(dstPath)
			return nil, err
		}
	}

	attach = &model.Attach{
		Hash:       hash,
		Uid:        uuid,
		Name:       name,
		Mime:       mime,
		Permission: model.PermissionPublic,
	}

	// record to db
	err = database.CreateAttach(attach)
	if err != nil {
		return nil, err
	}

	isSuccess = true

	return attach, nil

}
