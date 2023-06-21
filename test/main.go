package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	root := "C:\\Program Files\\Java\\jdk1.8.0_241\\jre\\lib\\" // 指定目录路径
	files, err := filepath.Glob(filepath.Join(root, "*.jar"))
	if err != nil {
		fmt.Printf("Error globbing the path %q: %v", root, err)
		return
	}

	for _, file := range files {
		fileInfo, err := os.Stat(file)
		if err != nil {
			fmt.Printf("Error getting file info for %q: %v", file, err)
			continue
		}

		if !fileInfo.IsDir() {
			fmt.Println(file)
		}
	}
}
