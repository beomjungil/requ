package parser_test

import (
	"testing"

	"github.com/go-requ/requ/parser"
	"github.com/stretchr/testify/assert"
)

func TestSimpleGet(t *testing.T) {
	result, _ := parser.Parse("./__mock__/simple_get_test.http", "")

	assert.Equal(t, len(result), 1, "Length must be '1'")
	assert.Equal(t, result[0].Method, "GET", "Method must be 'GET'")
	assert.Equal(t, result[0].Url, "https://example.com", "Url must be 'https://example.com'")
}

func TestGetWithoutMethod(t *testing.T) {
	result, _ := parser.Parse("./__mock__/get_without_method_test.http", "")

	assert.Equal(t, len(result), 1, "Length must be '1'")
	assert.Equal(t, result[0].Method, "GET", "Method must be 'GET'")
	assert.Equal(t, result[0].Url, "https://example.com", "Url must be 'https://example.com'")
}

func TestMultiGet(t *testing.T) {
	result, _ := parser.Parse("./__mock__/multi_get_test.http", "")

	assert.Equal(t, len(result), 3, "Length must be '3'")
}

func TestGetWithHeader(t *testing.T) {
	test_header := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer test",
	}
	result, _ := parser.Parse("./__mock__/get_with_header_test.http", "")

	assert.Equal(t, len(result), 1, "Length must be '1'")
	assert.Equal(t, result[0].Method, "GET", "Method must be 'GET'")
	assert.Equal(t, result[0].Headers, test_header, "Header must be Content-Type: application/json")
}
