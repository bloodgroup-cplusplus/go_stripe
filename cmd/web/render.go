package main

import (
	"embed"
	"text/template"
)

// we know we will pass information to our templates

type templateData struct {
	StringMap       map[string]string
	IntMap          map[string]int
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string
	Flash           string
	Warning         string
	Error           string
	IsAuthenticated int
	API             string
	CSSVersion      string
}

// we want to pass function to our templteas

var functions = template.FuncMap{}

// go:embed embed the directory templates

var templateFS embed.FS
