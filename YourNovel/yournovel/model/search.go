package model

import "html/template"

type SearchResult struct {
	Href    string `json:"href"`
	Title   string `json:"title"`
	IsParse int    `json:"is_parse"`
	Host    string `json:"host"`
}
type ChapterInfo struct {
	ChapterName string `json:"chapter_name"`
	ChapterUrl  string `json:"chapter_url"`
}

type NovelChapterApi struct {
	Name        string         `json:"name"`
	OriginUrl   string         `json:"origin_url"`
	ChapterInfo []*ChapterInfo `json:"chapter_info"`
	LinkPrefix  string         `json:"link_prefix"`
	Domain      string         `json:"domain"`
}
type NovelChapter struct {
	Name       string        `json:"name"`
	OriginUrl  string        `json:"origin_url"`
	Content    template.HTML `json:"content"`
	LinkPrefix string        `json:"link_prefix"`
	Domain     string        `json:"domain"`
}

type NovelContent struct {
	NovelName   string        `json:"novel_name"`
	Title       string        `json:"title"`
	ContentUrl  string        `json:"content_url"`
	Content     template.HTML `json:"content"`
	PreChapter  string        `json:"pre_chapter"`
	NextChapter string        `json:"next_chapter"`
}
