package zexecute

import (
	"fmt"
	"os"
	"path/filepath"
)

func AddCmd(path string) {
	fmt.Println("args", path)

	zrunFilePath := filepath.Join(path, "zrun.md")

	// Create or truncate zrun.md
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

	// Read directory
	dir, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	dir = dir[:len(dir)-1]
	for index, file := range dir {
		if file.IsDir() || filepath.Ext(file.Name()) != ".go" {
			continue
		}

		// Write line with a blank line after
		line := fmt.Sprintf(" ./run.sh %s/ %d\n\n", filepath.Base(path), (index + 1))
		_, err := zrunFile.WriteString(line)
		if err != nil {
			fmt.Printf("Error writing line: %v\n", err)
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

// func AddCmd(path string) {
// 	fmt.Println("args", path)

// 	dir, err := os.ReadDir(path)
// 	if err != nil {
// 		fmt.Println("Error reading directory:", err)
// 		return
// 	}

// 	dir = dir[:len(dir)-1]

// 	for index, file := range dir {
// 		fmt.Println(index, file)

// 		if file.IsDir() || filepath.Ext(file.Name()) != ".go" {
// 			continue
// 		}

// 		filePath := filepath.Join(path, file.Name())

// 		// Read the entire content
// 		contentBytes, err := os.ReadFile(filePath)
// 		if err != nil {
// 			fmt.Printf("Error reading file %s: %v\n", filePath, err)
// 			continue
// 		}
// 		content := string(contentBytes)

// 		// Remove any existing run.sh block comment
// 		content = removeRunShComment(content)

// 		// Add new block comment
// 		newComment := fmt.Sprintf("\n/*\n \n   ./run.sh %s %d   \n \n*/\n", path, index+1)
// 		content += newComment

// 		// Write back to file
// 		err = os.WriteFile(filePath, []byte(content), 0644)
// 		if err != nil {
// 			fmt.Printf("Error writing to file %s: %v\n", filePath, err)
// 		} else {
// 			fmt.Printf("Updated file: %s\n", filePath)
// 		}
// 	}
// }

// // Helper to remove any existing `/* ./run.sh ... */` block
// func removeRunShComment(content string) string {
// 	lines := strings.Split(content, "\n")
// 	var newLines []string
// 	inBlock := false

// 	for _, line := range lines {
// 		trim := strings.TrimSpace(line)
// 		if trim == "/*" {
// 			inBlock = true
// 			continue
// 		}
// 		if inBlock {
// 			if strings.HasPrefix(trim, "./run.sh") {
// 				continue
// 			}
// 			if trim == "*/" {
// 				inBlock = false
// 				continue
// 			}
// 			continue
// 		}
// 		newLines = append(newLines, line)
// 	}
// 	return strings.Join(newLines, "\n")
// }
