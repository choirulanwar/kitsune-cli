/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

type appValues struct {
	Cwd        string
	AppName    string
	ProjectApp string
	Schema     string
}

var (
	err       error
	fp        *os.File
	templates *template.Template
	subdirs   []string
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		values := appValues{}

		values.Cwd = GetWorkingDirectory()

		appNameLabel := "app-name"
		appNamePromptContent := promptContent{
			"Please input app name.",
			appNameLabel,
		}
		values.AppName = promptGetInput("app-name", appNamePromptContent)

		projectAppTypeLabel := "project-type"
		projectAppType := promptContent{
			"Please choose project app.", projectAppTypeLabel,
		}
		values.ProjectApp = promptGetSelect(projectAppTypeLabel, projectAppType)

		schemaLabel := "schema"
		schemaPromptContent := promptContent{
			"Please input yaml schema.",
			schemaLabel,
		}
		values.Schema = promptGetInput(schemaLabel, schemaPromptContent)

		transformYamlToProto(values)

		if values.ProjectApp == "go-micro" {
			createNewGoMicro(values)
		} else if values.ProjectApp == "nestjs" {
			createNewNestJs(values)
			prettifyTS(values)
		}

		fmt.Printf("\n🎉 Congratulations! Your new application is ready.\n")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
