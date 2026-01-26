/*
Copyright © 2026 Readmit
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
	Long: `Readmit helps you quickly generate project documentation and commit messages 
using AI. It analyzes your codebase to produce:

- README.md & CONTRIBUTION.md
- Comprehensive App & API Documentation
- Commit messages (based on git diffs)
- Security vulnerability reports (watchtower)

Examples:
  readmit generate readme
  readmit generate docs
  readmit generate commit
  readmit push
`,
	SilenceUsage:               true,
	SilenceErrors:              false,
	SuggestionsMinimumDistance: 1,
	Version:                    "0.6.1",
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
