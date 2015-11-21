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
	"fmt"
	"os"

	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
)

var cfgFile string
var accessToken string
var Verbose bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "gogetgithubstats",
	Short: "Get some useful statistics from github",
	Long: `gogetgithubstats is utility to get some useful statistics
on github repositories and users.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var rootCmdV *cobra.Command

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gogetgithubstats.json)")
	RootCmd.PersistentFlags().StringVar(&accessToken, "access-token", "ACCESSTOKEN", "Access Token to be used to authentication to github's API")

	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

	RootCmd.SuggestionsMinimumDistance = 1

	rootCmdV = RootCmd

	viper.BindPFlag("access-token", RootCmd.PersistentFlags().Lookup("access-token"))
	viper.BindPFlag("verbose", RootCmd.PersistentFlags().Lookup("verbose"))

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".gogetgithubstats") // name of config file (without extension)
	viper.AddConfigPath("$HOME")             // adding home directory as first search path
	viper.AutomaticEnv()                     // read in environment variables that match

	// This is the defaults
	viper.SetDefault("Verbose", true)

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigParseError); ok {
			jww.ERROR.Println(err)
		} else {
			jww.ERROR.Println("Unable to locate Config file.", err)
		}
	}

	if rootCmdV.PersistentFlags().Lookup("verbose").Changed {
		viper.Set("Verbose", Verbose)
	}

	if rootCmdV.PersistentFlags().Lookup("access-token").Changed {
		viper.Set("access-token", accessToken)
	}

	if viper.GetBool("verbose") {
		jww.SetStdoutThreshold(jww.LevelDebug)
	}
}
