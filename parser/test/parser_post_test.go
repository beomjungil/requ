package parser_test

import (
	"testing"

	"github.com/go-requ/requ/parser"
	"github.com/stretchr/testify/assert"
)

func TestSimplePost(t *testing.T) {
	result, _ := parser.Parse("./__mock__/simple_post_test.http")

	assert.Equal(t, len(result), 1, "Length must be '1'")
	assert.Equal(t, result[0].Method, "POST", "Method must be 'POST'")
	assert.Equal(t, result[0].Url, "https://example.com", "Url must be 'https://example.com'")
}
