package state

import (
	"HtmxBlog/model"
	"html/template"
	"time"
)

// admin auth
var CurrentToken string
var CreateTime time.Time

// App state
var CurrentState App

// templates
var Tmpl *template.Template
var AdminTmpl *template.Template

// i18n
var I18n *model.I18n
