/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package decryptor

import (
	"fmt"
	"kript/pkg/decryptor"

	"github.com/spf13/cobra"
)

// aesCmd represents the aes command
var aesCmd = &cobra.Command{
	Use:   "aes",
	Short: "Encrypt text with AES Algorithm",
	Long: `The algorithm described by AES is a symmetric-key algorithm,
	meaning the same key is used for both encrypting and decrypting the data`,
	Run: func(cmd *cobra.Command, args []string) {
		plaintext, _ := cmd.Flags().GetString("plaintext")
		key, _ := cmd.Flags().GetString("key")
		result := decryptor.AES(plaintext, []byte(key))
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(aesCmd)

	aesCmd.PersistentFlags().StringP("plaintext", "pt", "", "A plaintext to be encrypted")
	aesCmd.PersistentFlags().StringP("key", "k", "", "A key")
}
