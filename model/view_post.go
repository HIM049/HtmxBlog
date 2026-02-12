package model

import (
	"HtmxBlog/utils"
	"fmt"
	"html/template"
	"os"
	"sort"
	"strings"
)

type ViewPost struct {
	Post
	Content string
}

// LoadContent loads the content of the post
func (vp *ViewPost) LoadContent() error {
	content, err := os.ReadFile(vp.ContentPath())
	if err != nil {
		return err
	}

	vp.Content = string(content)
	return nil
}

// ParseContent parses the md content of the post
func (vp *ViewPost) ParseContent() template.HTML {
	if vp.Content == "" {
		return template.HTML("")
	}
	md, err := utils.ParseMarkdown([]byte(vp.Content))
	if err != nil {
		return template.HTML("")
	}
	return template.HTML(md)
}

func (p *ViewPost) TagsToString() string {
	return strings.Join(p.Tags, ", ")
}

func (p *ViewPost) CustomVarsToString() string {
	var lines []string
	for k, v := range p.CustomVars {
		lines = append(lines, fmt.Sprintf("%s: %v", k, v))
	}
	sort.Strings(lines)
	return strings.Join(lines, "\n")
}
