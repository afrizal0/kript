/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package encryptor

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kript",
	Short: "A tools for encrypt a text",
	Long:  `Kript is a CLI tools for encrypt any text with many supported encryption algorithm. Kript support AES ...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the Kript. Please use kript --help if you need help")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
