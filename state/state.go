package state

import (
	"html/template"
	"time"
)

var CurrentToken string
var CreateTime time.Time

var Tmpl *template.Template
var AdminTmpl *template.Template
