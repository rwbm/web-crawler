package parser

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractLinksEmptyReader(t *testing.T) {
	reader := strings.NewReader("")
	links := extractLinks(reader)

	assert.Equal(t, 0, len(links))
}

func TestExtractLinksInvalidHTML(t *testing.T) {
	reader := strings.NewReader("<html><body")
	links := extractLinks(reader)

	assert.Equal(t, 0, len(links))
}

func TestExtractLinksWithNoHref(t *testing.T) {
	reader := strings.NewReader("<html><a/></html>")
	links := extractLinks(reader)

	assert.Equal(t, 0, len(links))
}

func TestExtractLinksWithHref(t *testing.T) {
	reader := strings.NewReader(`<html><a href="https://mywebsite.com"></a></html>`)
	links := extractLinks(reader)

	assert.Equal(t, 1, len(links))
	assert.Equal(t, "https://mywebsite.com", links[0])
}
