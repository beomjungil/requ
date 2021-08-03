package parser

import (
	"io/ioutil"
	"strings"

	"github.com/go-requ/requ/model"
)

func Parse(filePath string) ([]model.HttpRequestConfig, error) {
	requestList := make([]model.HttpRequestConfig, 0)

	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		return nil, err
	}

	rawRequests := strings.Split(string(data), "###\n")

	for _, rawReq := range rawRequests {
		parseHttpFile(&requestList, strings.TrimSpace(rawReq))
	}

	return requestList, nil
}

func parseHttpFile(list *[]model.HttpRequestConfig, str string) {
	lines := strings.Split(str, "\n")

	method, url := func() (string, string) {
		switch firstLine := lines[0]; {
		case strings.HasPrefix(firstLine, "POST"):
			return "POST", strings.Split(firstLine, " ")[1]
		case strings.HasPrefix(firstLine, "PUT"):
			return "PUT", strings.Split(firstLine, " ")[1]
		case strings.HasPrefix(firstLine, "DELETE"):
			return "DELETE", strings.Split(firstLine, " ")[1]
		case strings.HasPrefix(firstLine, "GET"):
			return "GET", strings.Split(firstLine, " ")[1]
		default:
			return "GET", firstLine
		}
	}()

	if len(lines) == 0 {
		*list = append(*list, model.HttpRequestConfig{
			Method: method,
			Url:    url,
		})
		return
	}

	headerAndBody := lines[1:]

	headers, body := func() (map[string]string, string) {
		m := map[string]string{}

		for i := 0; len(headerAndBody) != 0 && headerAndBody[0] != ""; i++ {
			splitted := strings.SplitN(headerAndBody[0], ":", 2)
			m[splitted[0]] = strings.TrimSpace(splitted[1])

			headerAndBody = headerAndBody[1:]
		}

		return m, strings.TrimSpace(strings.Join(headerAndBody, "\n"))
	}()

	*list = append(*list, model.HttpRequestConfig{
		Method:  method,
		Url:     url,
		Headers: headers,
		Body:    body,
	})
}
