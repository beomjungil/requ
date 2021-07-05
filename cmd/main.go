package cmd

import (
	"fmt"

	"github.com/go-requ/requ/parser"
	"github.com/spf13/cobra"
)

func main(cmd *cobra.Command, args []string) {
	var filePath string = args[0]
	var config, err = parser.Parse(filePath)

	if err == nil {
		fmt.Println(config)
	}
}
