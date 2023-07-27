/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package encryptor

import (
	"fmt"
	"kript/pkg/encryptor"

	"github.com/spf13/cobra"
)

// aesCmd represents the aes command
var aesCmd = &cobra.Command{
	Use:   "aes",
	Short: "Encrypt text with AES Algorithm",
	Long: `The algorithm described by AES is a symmetric-key algorithm,
	meaning the same key is used for both encrypting and decrypting the data`,
	Run: func(cmd *cobra.Command, args []string) {
		key, _ := cmd.Flags().GetString("key")
		filename, _ := cmd.Flags().GetString("filename")
		if filename != "" {
			data := encryptor.ReadFiles("./data/" + filename)
			result := encryptor.AES(data, []byte(key))
			fmt.Println(result)
		} else {
			plaintext, _ := cmd.Flags().GetString("plaintext")
			result := encryptor.AES(plaintext, []byte(key))
			fmt.Println(result)
		}
	},
}

func init() {
	rootCmd.AddCommand(aesCmd)

	aesCmd.PersistentFlags().StringP("plaintext", "p", "", "A plaintext to be encrypted")

	aesCmd.PersistentFlags().StringP("filename", "f", "", "Filename of your data you want to encrypt")

	aesCmd.Flags().StringP("key", "k", "", "A key")
	aesCmd.MarkFlagRequired("key")
}
