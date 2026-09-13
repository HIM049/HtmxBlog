package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/log"
)

type ThemeVar struct {
	Key         string            `json:"key"`
	Name        map[string]string `json:"name"`
	Description map[string]string `json:"description"`
	Type        string            `json:"type"`
}

type ThemeVarTranslated struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

type ThemeMeta struct {
	Path     string `json:"-"`
	ID       string `json:"id"`
	Metadata int    `json:"metadata"`
	Version  string `json:"version"`
	Author   string `json:"author"`
}

type ThemeMetaFile struct {
	ThemeMeta
	Name        map[string]string `json:"name"`
	Description map[string]string `json:"description"`
	GlobalVars  []ThemeVar        `json:"global_vars"`
	PostVars    []ThemeVar        `json:"post_vars"`
	Pages       []ThemeVar        `json:"pages"`
}

type ThemeMetaTranslated struct {
	ThemeMeta
	Name        string                `json:"name"`
	Description string                `json:"description"`
	GlobalVars  []ThemeVarTranslated  `json:"global_vars"`
	PostVars    []ThemeVarTranslated  `json:"post_vars"`
	Pages       []ThemeVarTranslated  `json:"pages"`
}

func InitTheme() {
	themes, err := findTheme()
	if err != nil {
		panic("failed to find themes: " + err.Error())
	}
	if len(themes) <= 0 {
		panic("failed to find themes (len=0)")
	}

	targetID := Cfg.Settings["theme"]
	if targetID == "" {
		log.Warnf("theme not set, use default")
		targetID = "default"
	}

	var targetTheme *ThemeMetaFile
	for _, theme := range themes {
		if theme.ID == targetID {
			targetTheme = &theme
			break
		}
	}

	if targetTheme == nil {
		log.Errorf("Cannot found theme: %s, use first one", targetID)
		targetTheme = &themes[0]
	}

	lang := Cfg.Settings["language"]
	var theme ThemeMetaTranslated
	theme.Path = targetTheme.Path
	theme.ThemeMeta = targetTheme.ThemeMeta

	theme.Name = langFallback(lang, targetTheme.Name)
	theme.Description = langFallback(lang, targetTheme.Description)
	for _, v := range targetTheme.GlobalVars {
		theme.GlobalVars = append(theme.GlobalVars, v.TranslatedFallback(lang))
	}
	for _, v := range targetTheme.PostVars {
		theme.PostVars = append(theme.PostVars, v.TranslatedFallback(lang))
	}
	for _, v := range targetTheme.Pages {
		theme.Pages = append(theme.Pages, v.TranslatedFallback(lang))
	}

	Cfg.Theme = theme

}

func (t ThemeVar) TranslatedFallback(lang string) ThemeVarTranslated {
	return ThemeVarTranslated{
		Key:         t.Key,
		Name:        langFallback(lang, t.Name),
		Description: langFallback(lang, t.Description),
		Type:        t.Type,
	}
}

func langFallback(lang string, translate map[string]string) string {
	val := translate[lang]
	if val != "" {
		return val
	}

	val = translate["en_us"]
	if val != "" {
		return val
	}

	for _, val := range translate {
		if val != "" {
			return val
		}
	}

	return "Name Mismatch"
}

func findTheme() ([]ThemeMetaFile, error) {
	var themes []ThemeMetaFile
	err := filepath.WalkDir(Cfg.Service.ThemeDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && strings.EqualFold(d.Name(), "metadata.json") {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			var theme ThemeMetaFile
			if err := json.Unmarshal(content, &theme); err != nil {
				return err
			}
			theme.Path = filepath.Dir(path)
			themes = append(themes, theme)
		}
		return nil
	})

	return themes, err
}
