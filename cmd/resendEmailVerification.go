/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// ResendEmailVerificationCmd represents the resend email verification command
var ResendEmailVerificationCmd = &cobra.Command{
	Use:   "resendEmailVerification",
	Short: "Resend email verification",
	Long:  "Resend email verification for a specific user",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
	Args:       cobra.OnlyValidArgs,
	SuggestFor: []string{"resendEmailVerification"},
	ValidArgs:  []string{"user"},
}

func init() {
	rootCmd.AddCommand(ResendEmailVerificationCmd)
}
