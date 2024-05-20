package cmd

import (
	"github.com/dyammarcano/version/internal/generate"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		path, err := cmd.Flags().GetString("project")
		if err != nil {
			return err
		}

		ver, err := generate.NewVersion(path)
		if err != nil {
			return err
		}

		yes, err := cmd.Flags().GetBool("yes")
		if err != nil {
			return err
		}

		if !yes {
			prompt := promptui.Prompt{
				Label:     "Do you want to generate a new version?",
				IsConfirm: true,
			}

			if _, err := prompt.Run(); err != nil {
				return generate.OperationCanceled()
			}
		}

		if err = ver.Generate(); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringP("project", "p", "", "Project path")
	generateCmd.Flags().BoolP("yes", "y", false, "Skip confirmation")
}
