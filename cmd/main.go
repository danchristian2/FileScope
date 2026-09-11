package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type ProjectStats struct {
	Files		int
	Directories int
	TotalSize   int64
}
 func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: filescope<directory>")
		return
	}


	root := os.Args[1]

	stats := ProjectStats{}

	err := filepath.WalkDir(root, func(path string, entry os.DirEntry, err error) error{
		

		if err != nil {
			return err
		}
		if entry.IsDir() {
			stats.Directories++


			return nil
		}
		stats.Files++


		info, err := entry.Info()
		if err != nil {
			return err
		}


		stats.TotalSize += info.Size()


		return nil
	
	})
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("==== FileScope ====")
	fmt.Println()

	fmt.Println("Directory: ", root)
	fmt.Println("Files: ", stats.Files)
	fmt.Println("Directories: ",stats.Directories)
	fmt.Println("Total size: ",stats.TotalSize)
 }

 // a function to format the file sizes

 func formatSize(bytes int64) string {
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
	)
	if bytes >= GB {
		return fmt.Sprintf("%.2f GB", float64(bytes)/float64(GB))
	}
	if bytes >= MB {
		return fmt.Sprint("%.2f MB", float64(bytes)/float64(MB))
	}
	if bytes >= KB {
		return fmt.Sprintf("%.2f KB",float64(bytes)/float64(KB))
	}
	return fmt.Sprintf("%d B", bytes)
 }