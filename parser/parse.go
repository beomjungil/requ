package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/go-requ/requ/model"
	"github.com/thoas/go-funk"
)

func Parse(filePath string, variableFilePath string) ([]model.HttpRequestConfig, error) {
	requestList := make([]model.HttpRequestConfig, 0)

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	dataWithVariable, err := ReplaceVariable(string(data), variableFilePath)
	if err != nil {
		return nil, err
	}

	rawRequests := strings.Split(dataWithVariable, "###\n")

	for _, rawReq := range rawRequests {
		parseHttpFile(&requestList, strings.TrimSpace(rawReq))
	}

	return requestList, nil
}

func ReplaceVariable(data string, variableFilePath string) (string, error) {
	if variableFilePath == "" {
		return data, nil
	}

	variable, err := ioutil.ReadFile(variableFilePath)
	if err != nil {
		return "", err
	}

	variableMap := make(map[string]string)
	err = json.Unmarshal(variable, &variableMap)
	if err != nil {
		return "", err
	}

	for key, value := range variableMap {
		data = strings.Replace(data, fmt.Sprintf("{{%s}}", key), value, -1)
	}

	return data, nil
}

func parseHttpFile(list *[]model.HttpRequestConfig, str string) {
	lines := funk.FilterString(strings.Split(str, "\n"), func(s string) bool {
		return !strings.HasPrefix(s, "#") && !strings.HasPrefix(s, "@")
	})

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
