package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/treasureuzoma/readmit/readmit/controllers"
	"path/filepath"

	"github.com/spf13/cobra"
)


var (
	withCommit bool
	singleFile bool
)

var generateCmd = &cobra.Command{
	Use:   "generate [type]",
	Short: "Generate files like README, CONTRIBUTION guide, or commit messages",
	Long: `Supported types:
	- readme 		 Generates README.md
	- contribution 	 Generates CONTRIBUTION.md
	- commit 		 Suggests commit message (printed to console)
	- watchtower Pick up all vulnerabilities found and creates a report.md file
	- docs 		 Generates comprehensive documentation`,
	Example: `	readmit generate readme
				readmit generate contribution
				readmit generate commit
				readmit generate docs
				readmit generate docs --single-file
				readmit watchtower`,
	Args: cobra.ExactArgs(1),
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
                                                                                                        
                                                                                                        
                                                                                                        
`)
		time.Sleep(500 * time.Millisecond)

		fileType := strings.ToLower(args[0])
		validTypes := map[string]bool{"readme": true, "contribution": true, "commit": true, "watchtower": true, "docs": true}
		if !validTypes[fileType] {
			log.Printf("[ERROR] Unsupported type: %s (valid: readme, contribution, commit, docs)", fileType)
			return
		}

		fmt.Printf("✓ Analyzing codebase...\n")
		time.Sleep(500 * time.Millisecond)

		// Optimization: Only include full codebase for non-commit types
		includeFullCodebase := fileType != "commit"
		
		fmt.Printf("✓ Feeding codebase to Readmit AI...\n")
		generatedContent, err := controllers.GenerateAIContent(fileType, includeFullCodebase)
		if err != nil {
			log.Fatalf("[ERROR] %v", err)
		}

		fmt.Printf("✓ Generating %s content...\n", fileType)
		time.Sleep(500 * time.Millisecond)

		switch fileType {
		case "readme":
			if err := os.WriteFile("README.md", []byte(generatedContent), 0644); err != nil {
				log.Printf("[ERROR] Failed to write README.md: %v", err)
			}
			fmt.Printf("✓ Successfully created README.md\n")

		case "contribution":
			if err := os.WriteFile("CONTRIBUTION.md", []byte(generatedContent), 0644); err != nil {
				log.Printf("[ERROR] Failed to write CONTRIBUTION.md: %v", err)
			}
			fmt.Printf("✓ Successfully created CONTRIBUTION.md\n")

		case "commit":
			fmt.Println("\n-- Suggested Commit Message --")
			fmt.Println(generatedContent)
			fmt.Println("------------------------------")
			
			if withCommit {
				fmt.Println("Running git add . and git commit...")
				// Git add .
				addCmd := exec.Command("git", "add", ".")
				if err := addCmd.Run(); err != nil {
					log.Printf("[ERROR] Failed to run git add .: %v", err)
					return
				}

				// Git commit
				commitCmd := exec.Command("git", "commit", "-m", generatedContent)
				commitCmd.Stdout = os.Stdout
				commitCmd.Stderr = os.Stderr
				if err := commitCmd.Run(); err != nil {
					log.Printf("[ERROR] Failed to run git commit: %v", err)
					return
				}
				fmt.Println("✓ Changes committed successfully.")
			} else {
				fmt.Println("✓ Commit message generated and printed to console.")
			}

		case "docs":
			if singleFile {
				if err := os.WriteFile("docs.md", []byte(generatedContent), 0644); err != nil {
					log.Printf("[ERROR] Failed to write docs.md: %v", err)
				}
				fmt.Printf("✓ Successfully created docs.md\n")
			} else {
				fmt.Println("✓ Creating docs folder structure...")
				if err := os.MkdirAll("docs", 0755); err != nil {
					log.Printf("[ERROR] Failed to create docs folder: %v", err)
					return
				}

				// Basic parsing of AI response for multiple files
				// Expecting format like --- filename.md ---
				parts := strings.Split(generatedContent, "---")
				if len(parts) <= 1 {
					// Fallback if AI didn't follow format: write everything to index.md
					if err := os.WriteFile("docs/index.md", []byte(generatedContent), 0644); err != nil {
						log.Printf("[ERROR] Failed to write docs/index.md: %v", err)
					}
					fmt.Printf("✓ Successfully created docs/index.md\n")
				} else {
					for i := 1; i < len(parts); i += 2 {
						if i+1 >= len(parts) {
							break
						}
						header := strings.TrimSpace(parts[i])
						content := strings.TrimSpace(parts[i+1])
						// Remove potential extra dashes or whitespace from header
						filename := strings.TrimSpace(strings.ReplaceAll(header, "---", ""))
						// Basic security check for filename
						filename = filepath.Base(filename)
						if filename == "." || filename == "/" {
							continue
						}
						
						filePath := filepath.Join("docs", filename)
						if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
							log.Printf("[ERROR] Failed to write %s: %v", filePath, err)
						} else {
							fmt.Printf("✓ Created %s\n", filePath)
						}
					}
					fmt.Printf("✓ Successfully populated docs folder\n")
				}
			}

		default:
			// For generic types, we still use the filename generated inside (but wait, GenerateAIContent doesn't return the filename)
			// Actually we can just write it to a default name or handle it.
			// The original code used fileType-uuid.txt.
			log.Printf("✓ Successfully generated %s content", fileType)
			fmt.Println(generatedContent)
		}
	},
}

func init() {
	generateCmd.Flags().BoolVar(&withCommit, "with-commit", false, "Automatically commit changes with the generated message")
	generateCmd.Flags().BoolVar(&singleFile, "single-file", false, "Generate documentation in a single docs.md file instead of a folder")
	rootCmd.AddCommand(generateCmd)
}
