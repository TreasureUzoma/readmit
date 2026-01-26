package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/treasureuzoma/readmit/readmit/controllers"
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push [branch]",
	Short: "Automatically stage, commit with AI, and push changes",
	Long: `readmit push streamlines your workflow by:
1. Running git add .
2. Generating a commit message using Readmit AI
3. Committing the changes
4. Pushing to the remote repository`,
	Example: `  readmit push
  readmit push main`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(Ascii)
		time.Sleep(200 * time.Millisecond)

		// 1. Git add .
		fmt.Println("✓ Staging changes (git add .)...")
		if err := exec.Command("git", "add", ".").Run(); err != nil {
			log.Fatalf("[ERROR] git add failed: %v", err)
		}

		// 2. Generate Commit Message
		fmt.Println("✓ Generating AI commit message...")
		commitMsg, err := controllers.GenerateAIContent("commit", false)
		if err != nil {
			log.Fatalf("[ERROR] Failed to generate commit message: %v", err)
		}
		fmt.Printf("✓ AI Suggested: %s\n", commitMsg)

		// 3. Git commit
		fmt.Println("✓ Committing changes...")
		commitCmd := exec.Command("git", "commit", "-m", commitMsg)
		if out, err := commitCmd.CombinedOutput(); err != nil {
			if strings.Contains(string(out), "nothing to commit") {
				fmt.Println("! Nothing new to commit.")
			} else {
				log.Fatalf("[ERROR] git commit failed: %v. Output: %s", err, string(out))
			}
		}

		// 4. Git Push
		var branch string
		if len(args) > 0 {
			branch = args[0]
		} else {
			out, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").CombinedOutput()
			if err != nil {
				log.Fatalf("[ERROR] Failed to get current branch: %v", err)
			}
			branch = strings.TrimSpace(string(out))
		}

		fmt.Printf("✓ Pushing to origin %s...\n", branch)
		gitPush := exec.Command("git", "push", "origin", branch)
		gitPush.Stdout = os.Stdout
		gitPush.Stderr = os.Stderr
		
		if err := gitPush.Run(); err != nil {
			log.Fatalf("[ERROR] git push failed: %v", err)
		}
		
		fmt.Printf("✓ Successfully pushed to origin %s\n", branch)
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
}
