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

	log "github.com/sirupsen/logrus"

	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// kubeCmd represents the kube command
var kubeCmd = &cobra.Command{
	Use:   "kube",
	Short: "Sync your gvault secrets with kubernetes",
	Run: func(cmd *cobra.Command, args []string) {

		// Authenticate against the cluster
		client, err := getClient()
		if err != nil {
			log.Fatal(err)
		}

		secret := &v1.Secret{
			Type: v1.SecretTypeOpaque,
		}

		secret.SetName("gvault-" + fmt.Sprintf("%v", secretsCmd.vault.Version))

		if err := secretsCmd.vault.DecryptAll(); err != nil {
			log.Fatal(err)
		}

		secret.StringData = secretsCmd.vault.Base64Encode()

		if _, err := client.CoreV1().Secrets("default").Create(secret); err != nil {
			log.Fatal(errors.Wrap(err, "failed to create kubernetes secret"))
		}
	},
}

// Create a client so we're allowed to perform requests
// Because of the use of `os.Getenv("HOME")`, this only works on unix environments
func getClient() (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(os.Getenv("HOME"), ".kube", "config"))
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}

func init() {
	rootCmd.AddCommand(kubeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kubeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kubeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
