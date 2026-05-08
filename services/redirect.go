package services

import (
	"HtmxBlog/config"
	"HtmxBlog/model"
	"strings"
)

// FindRedirectBySource looks up a redirect rule by the request path.
// It normalizes the path by ensuring a leading "/" before querying.
// Returns nil and an error if no match is found.
func FindRedirectBySource(path string) (*model.Redirect, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	var redirect model.Redirect
	err := config.DB.Where("source_path = ?", path).First(&redirect).Error
	if err != nil {
		return nil, err
	}
	return &redirect, nil
}

// ReadAllRedirects returns all redirect rules.
func ReadAllRedirects() ([]model.Redirect, error) {
	var redirects []model.Redirect
	err := config.DB.Order("id desc").Find(&redirects).Error
	return redirects, err
}

// CreateRedirect creates a new redirect rule.
func CreateRedirect(sourcePath, targetPath string, statusCode int) error {
	redirect := model.Redirect{
		SourcePath: sourcePath,
		TargetPath: targetPath,
		StatusCode: statusCode,
	}
	return config.DB.Create(&redirect).Error
}

// UpdateRedirect updates an existing redirect rule.
func UpdateRedirect(id uint, sourcePath, targetPath string, statusCode int) error {
	return config.DB.Model(&model.Redirect{}).Where("id = ?", id).Updates(map[string]any{
		"source_path": sourcePath,
		"target_path": targetPath,
		"status_code": statusCode,
	}).Error
}

// DeleteRedirect deletes a redirect rule by ID.
func DeleteRedirect(id uint) error {
	return config.DB.Delete(&model.Redirect{}, id).Error
}
