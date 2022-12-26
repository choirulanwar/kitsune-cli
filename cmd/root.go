/*
Copyright © 2022 Choirul Anwar <labs.choirulanwar@gmail.com>
*/
package cmd

import (
	"os"
	"os/signal"

	cc "github.com/ivanpirog/coloredcobra"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kitsune-cli",
	Short: "A command line tool to automatically generate some or all feature set files for NestJS & Go-Micro.",
	Long:  `Using the command line tool, developers can quickly and easily generate some or all of the feature set files needed to create a NestJS & Go-Micro application, saving them time and effort when starting a new project.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt)

	go func() {
		for range done {
			os.Exit(0)
		}
	}()

	cc.Init(&cc.Config{
		RootCmd:  rootCmd,
		Headings: cc.HiCyan + cc.Bold + cc.Underline,
		Commands: cc.HiYellow + cc.Bold,
		Example:  cc.Italic,
		ExecName: cc.Bold,
		Flags:    cc.Bold,
	})

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kitsune-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
