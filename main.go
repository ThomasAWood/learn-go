package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func resetAlgorithms() {
	fmt.Println("⚠️  This will reset ALL algorithm implementations.")
	fmt.Print("Are you sure? (y/N): ")
	var answer string
	fmt.Scan(&answer)

	if answer == "n" || answer == "N" {
		return
	}

	fmt.Println("🔄 Resetting algorithms...")

	// Walk through templates directory
	err := filepath.Walk("templates", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasSuffix(path, ".template") {
			// Convert template path to actual file path
			targetPath := strings.Replace(path, "templates/", "algorithms/", 1)
			targetPath = strings.Replace(targetPath, ".template", "", 1)

			return copyFile(path, targetPath)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error resetting: %v\n", err)
	} else {
		fmt.Println("✅ All algorithms reset!")
	}
}

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// Create destination directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}

func main() {
	resetAlgorithms()
}
