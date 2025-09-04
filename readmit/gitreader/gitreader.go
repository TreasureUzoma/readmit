package gitreader

import (
	"bytes"
	"fmt"
	"os/exec"
)

func GetStagedDiff() (string, error) {
	cmd := exec.Command("git", "diff", "--cached")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to run git diff --cached: %w", err)
	}

	return out.String(), nil
}

func GetUnstagedDiff() (string, error) {
	cmd := exec.Command("git", "diff")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to run git diff: %w", err)
	}

	return out.String(), nil
}

func GetLastCommitDiff() (string, error) {
	cmd := exec.Command("git", "show", "HEAD")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to run git show HEAD: %w", err)
	}

	return out.String(), nil
}

// GetBestDiff tries staged, then unstaged, then last commit.
// If all are empty, returns "" with nil error.
func GetBestDiff() (string, error) {
	diff, err := GetStagedDiff()
	if err != nil {
		return "", fmt.Errorf("staged diff failed: %w", err)
	}
	if diff != "" {
		return diff, nil
	}

	diff, err = GetUnstagedDiff()
	if err != nil {
		return "", fmt.Errorf("unstaged diff failed: %w", err)
	}
	if diff != "" {
		return diff, nil
	}

	diff, err = GetLastCommitDiff()
	if err != nil {
		return "", fmt.Errorf("last commit diff failed: %w", err)
	}
	if diff != "" {
		return diff, nil
	}

	return "", nil // no diffs at all
}

