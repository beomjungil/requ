package cmd

import (
	"fmt"

	"github.com/go-requ/requ/model"
	"github.com/go-requ/requ/network"
	"github.com/go-requ/requ/parser"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func main(cmd *cobra.Command, args []string) {
	var filePath string = args[0]
	var configList, err = parser.Parse(filePath, variableFilePath)
	if err != nil {
		fmt.Println(configList)
	}

	config, err := selectFromConfigList(configList)

	if err != nil {
		network.Request(config)
	} else {
		panic(err)
	}
}

func selectFromConfigList(configList []model.HttpRequestConfig) (model.HttpRequestConfig, error) {
	templates := &promptui.SelectTemplates{
		Label:    "{{ .Method }}  {{ .Url }}",
		Active:   "\U00002705 {{ .Method | cyan }} ({{ .Url | red }})",
		Inactive: "  {{ .Method | cyan }} ({{ .Url | red }})",
		Selected: "\U00002705 {{ .Method }} {{ .Url }}",
		Details: `
--------- Config ----------
{{ "Method:" | faint }}	{{ .Method }}
{{ "Url:" | faint }}	{{ .Url }}
{{ "Headers:" | faint }} {{ .Headers }}
{{ "Body:" | faint }} {{ .Body }}`,
	}

	prompt := promptui.Select{
		Label:     "Select Request",
		Items:     configList,
		Templates: templates,
		Size:      len(configList),
	}

	i, _, err := prompt.Run()
	if err != nil {
		return configList[0], err
	}
	return configList[i], nil
}
