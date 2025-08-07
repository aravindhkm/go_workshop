package interview

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			// If the entry is a file, send its size to the channel.
			f, err := entry.Info()
			if err == nil {
				fileSizes <- f.Size()
			} else {
				fmt.Fprintf(os.Stderr, "du1: %v\n", err)
			}
		}
	}
}

// dirents returns the entries of directory dir.
func dirents(dir string) []os.DirEntry {
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func PrintDiskUsageOne() {
	// Determine the initial directories.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Write code to traverse the file tree.

	// remaining code
	// Create a channel to collect file sizes.
	fileSizes := make(chan int64)

	// Traverse each root directory using walkDir (in a goroutine).
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes) // close the channel after traversal
	}()

	// Receive sizes from the channel and compute totals.
	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}

	// Print total disk usage.
	printDiskUsage(nfiles, nbytes)
}
