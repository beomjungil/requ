package parser

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/go-requ/requ/model"
)

func Parse(filePath string) ([]model.HttpRequestConfig, error) {

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	requestList := make([]model.HttpRequestConfig, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		parseHttpFile(&requestList, text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return requestList, nil
}

func parseHttpFile(list *[]model.HttpRequestConfig, str string) {
	if strings.HasPrefix(str, "GET") {
		*list = append(*list, model.HttpRequestConfig{
			Method: "GET",
			Url:    strings.Split(str, " ")[1],
		})
	}

	if strings.HasPrefix(str, "http") {
		*list = append(*list, model.HttpRequestConfig{
			Method: "GET",
			Url:    str,
		})
	}

	if strings.HasPrefix(str, "POST") {
		*list = append(*list, model.HttpRequestConfig{
			Method: "POST",
			Url:    strings.Split(str, " ")[1],
		})
	}

	if strings.HasPrefix(str, "PUT") {
		*list = append(*list, model.HttpRequestConfig{
			Method: "PUT",
			Url:    strings.Split(str, " ")[1],
		})
	}

	if strings.HasPrefix(str, "DELETE") {
		*list = append(*list, model.HttpRequestConfig{
			Method: "DELETE",
			Url:    strings.Split(str, " ")[1],
		})
	}
}
