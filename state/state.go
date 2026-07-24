package state

import (
	"HtmxBlog/model"
	"html/template"
	"time"
)

var CurrentToken string
var CreateTime time.Time

var CurrentState App

var Tmpl *template.Template
var AdminTmpl *template.Template

var I18n *model.I18n
