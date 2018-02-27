// Copyright © 2018 James Qualls https://github.com/sourcec0de
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// secretsVersionCmd represents the version command
var secretsVersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints a uint64 hash of all the secrets in the vault",
	Run: func(cmd *cobra.Command, args []string) {
		version, err := gvault.HashSecrets()
		if err != nil {
			logger.Fatal(err)
		}

		fmt.Print(version)
	},
}

func init() {
	secretsCmd.AddCommand(secretsVersionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// secretsVersionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// secretsVersionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
