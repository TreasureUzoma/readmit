package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push [branch]",
	Short: "Push changes to the remote repository",
	Long: `readmit push pushes the current or specified branch to origin.
It ensures your changes are synchronized with the remote repository.`,
	Example: `  readmit push
  readmit push main
  readmit push develop`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(`
  o__ __o         o__ __o__/_          o           o__ __o        o          o   __o__  ____o__ __o____ 
 <|     v\       <|    v              <|>         <|     v\      <|\        /|>    |     /   \   /   \  
 / \     <\      < >                  / \         / \     <\     / \\o    o// \   / \         \o/       
 \o/     o/       |                 o/   \o       \o/       \o   \o/ v\  /v \o/   \o/          |        
  |__  _<|        o__/_            <|__ __|>       |         |>   |   <\/>   |     |          < >       
  |       \       |                /       \      / \       //   / \        / \   < >          |        
 <o>       \o    <o>             o/         \o    \o/      /     \o/        \o/    |           o        
  |         v\    |             /v           v\    |      o       |          |     o          <|        
 / \         <\  / \  _\o__/_  />             <\  / \  __/>      / \        / \  __|>_        / \       
                                                                                                        
         🚀 Pushing...
`)
		time.Sleep(500 * time.Millisecond)

		var branch string
		if len(args) > 0 {
			branch = args[0]
		} else {
			// Get current branch
			out, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").CombinedOutput()
			if err != nil {
				log.Fatalf("[ERROR] Failed to get current branch: %v. Output: %s", err, string(out))
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
