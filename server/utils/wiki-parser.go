package utils

import (
	"github.com/gocolly/colly"
)

type FandomPageParser struct {
	PageContents *colly.Collector
	Url          string
}

func NewPageParser(wikiPage *colly.Collector) *FandomPageParser {
	return &FandomPageParser{}
}

type InfoboxParser struct {
}

func (p *FandomPageParser) Infobox() {
}

type ParagraphParser struct {
	FilterBlocks []string
}

func (p *FandomPageParser) Paragraph() {
}

type CategoryFilters []string

func (p *FandomPageParser) Categories(filters CategoryFilters) {
}
