package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goget",
	Short: "goget helps you download files over HTTP.",
	Run: func(cmd *cobra.Command, args []string) {
		url, err := cmd.Flags().GetString("url")
		if err != nil {
			fmt.Printf("Error : %s", err)
			return
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error : %s", err)
			return
		}
		defer resp.Body.Close()
		filepath := path.Base(resp.Request.URL.String())

		file, err := os.Create(filepath)
		if err != nil {
			fmt.Printf("Error : %s", err)
			return
		}
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			fmt.Printf("Error : %s", err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().String("url", "", "Please provide a url")
}
