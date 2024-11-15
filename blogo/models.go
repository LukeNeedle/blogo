package main

import (
	"html/template"
	"time"
)

type ArticleData struct {
	Title    string
	Author   string
	Summary  string
	Tags     []string
	Image    string
	Date     time.Time
	Slug     string
	Draft    bool
	Layout   string
	Md       string
	Html     template.HTML
}

type Config struct {
	Title       string
	Description string
	Url         string
	Keywords    string
	Analytics   string
	Timezone    string
}
