package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/bitforth/reqio/comms"
	"github.com/bitforth/reqio/parser"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "reqio [filename]",
	Short: "CLI tool that issues HTTP requests based on a file input",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		request, err := parser.ParseFile(filename)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		response, err := comms.MakeRequest(request)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Response Status:", response.Status)
		bodyBytes, err :=  io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			os.Exit(1)
		}
		response.Body.Close()
		fmt.Println("Response Body:", string(bodyBytes))
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}