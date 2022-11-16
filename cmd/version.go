/*
Copyright © 2022 Blylei <blylei.info@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display frabit-ctl version number",
	Run:   runVersion,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func runVersion(cmd *cobra.Command, args []string) {
	v := newVersion()
	versionStr := v.String()
	fmt.Println(versionStr)
}

type Version struct {
	Major int
	Minor int
	Patch int
}

func newVersion() *Version {
	return &Version{
		Major: 2,
		Minor: 0,
		Patch: 0,
	}
}

func (v *Version) String() string {
	return fmt.Sprintf("frabit-ctl %v.%v.%v", v.Major, v.Minor, v.Patch)
}
