/*
Copyright © 2025 Readmit
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd is the base command for the Readmit CLI.
var rootCmd = &cobra.Command{
	Use:   "readmit",
	Short: "AI-powered file and commit generator for your projects",
	Long: `Readmit helps you quickly generate common project files and commit messages 
using AI. It analyzes your codebase and Git history to produce:

- README.md
- CONTRIBUTION.md
- Commit messages (based on git diffs)
- Custom text files (design-doc, changelog, etc.)

Examples:
  readmit generate readme
  readmit generate contribution
  readmit generate commit
  readmit generate design-doc
`,
	SilenceUsage:  true,
	SilenceErrors: true,
}

// Execute runs the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Remove the placeholder toggle flag
	// Add global flags here if needed later
	// Example: verbose logging, custom config file, etc.
}
