// Frabit - The next-generation database automatic operation platform
// Copyright © 2022-2023 Frabit Labs
//
// Licensed under the GNU General Public License, Version 3.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://www.gnu.org/licenses/gpl-3.0.txt
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package auth

import (
	"fmt"
	"github.com/frabits/frabit/pkg/common/version"

	"github.com/spf13/cobra"
)

type loginOpt struct {
	username string
	password string
	token    string
}

// CmdLogin authenticate frabit-admin to frabit-server via token
var CmdLogin = &cobra.Command{
	Use:   "login [flag]",
	Short: "Login frabit-admin to frabit-server via token",
	Long: `
frabit-admin auth login
`,
	Run: runLogin,
}

func runLogin(cmd *cobra.Command, args []string) {
	fmt.Printf("%s\n", version.InfoStr.String())
}

func init() {
	// cmd.rootCmd.AddCommand(newVersionCmd)
}
