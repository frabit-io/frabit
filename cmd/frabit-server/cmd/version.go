/* (c) 2022 Frabit Project maintained and limited by Blylei < blylei.info@gmail.com >
GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)

This file is part of Frabit

*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/frabits/frabit/common/version"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display frabit version information",
	Run:   runVersion,
}

func runVersion(cmd *cobra.Command, args []string) {
	fmt.Printf("%s\n", version.InfoStr.String())
	fmt.Printf("%s", version.InfoStr.BuildInfo())
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
