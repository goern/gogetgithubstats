/* Copyright © 2015 Christoph Görn

This file is part of gogetgithubstats.

gogetgithubstats is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

gogetgithubstats is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with gogetgithubstats. If not, see <http://www.gnu.org/licenses/>.
*/

package cmd

import (
	"strings"

	"github.com/goern/gogetgithubstats/pkg/githubstats"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
)

// getCmd respresents the get command
var getCmd = &cobra.Command{
	Use:   "get USERNAME/REPO",
	Short: "get statistics of a repository USERNAME/REPO",
	Long:  `This will get some statistics on github repository 'USERNAME/REPO'.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			jww.ERROR.Println("get needs a USERNAME/REPO")
			return
		}

		// lets get user's name and repository name
		splitString := strings.Split(args[0], "/")

		if len(splitString) != 2 {
			jww.ERROR.Println("get needs a USERNAME/REPO")
			return
		}

		if viper.GetString("access-token") == "ACCESSTOKEN" {
			jww.ERROR.Println("get needs a ACCESSTOKEN")
			return
		}

		jww.DEBUG.Printf("access-token=%s", accessToken)
		jww.DEBUG.Printf("u=%s, r=%s\n", splitString[0], splitString[1])

		githubstats.GetStatsByRepo(splitString[0], splitString[1], accessToken)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)

}
