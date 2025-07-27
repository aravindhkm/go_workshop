package zexecute

import (
	"fmt"
	"os"
	"path/filepath"
)

// 	dir = dir[:len(dir)-1]

func AddCmd(path string) {
	// err := os.MkdirAll(path, os.ModePerm)
	// if err != nil {
	// 	fmt.Printf("Error creating directory: %v\n", err)
	// 	return
	// }

	zrunFilePath := filepath.Join(path, "zrun.md")

	// Create or overwrite zrun.md
	zrunFile, err := os.OpenFile(zrunFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("Error creating zrun.md: %v\n", err)
		return
	}
	defer zrunFile.Close()

	// Start block comment
	_, err = zrunFile.WriteString("/*\n\n")
	if err != nil {
		fmt.Printf("Error writing start block: %v\n", err)
		return
	}

	// Read directory entries
	dir, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for index, file := range dir {
		if file.IsDir() || filepath.Ext(file.Name()) != ".go" {
			continue
		}

		// fmt.Println("file name", file)

		// Relative file path
		// relPath := filepath.Join(filepath.Base(path), file.Name())

		block := fmt.Sprintf(" %s\n ./run.sh %s/ %d\n\n", file.Name(), filepath.Base(path), (index + 1))
		_, err := zrunFile.WriteString(block)
		if err != nil {
			fmt.Printf("Error writing block: %v\n", err)
			return
		}
	}

	// End block comment
	_, err = zrunFile.WriteString("*/\n")
	if err != nil {
		fmt.Printf("Error writing end block: %v\n", err)
		return
	}

	fmt.Printf("✅ zrun.md created at: %s\n", zrunFilePath)
}
