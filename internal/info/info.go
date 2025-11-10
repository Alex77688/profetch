package info

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func FetchInfo(root string) ([]string, error) {
	res := make([]string, 0)
	files, err := addFilesAndLines(root)
	if err != nil {
		return nil, err
	}
	res = append(res, files...)
	return res, nil
}

func addFilesAndLines(root string) ([]string, error) {
	fileCount, lineCount, err := countFilesAndLines(root)
	if err != nil {
		return make([]string, 0), err
	}
	res := make([]string, 2)
	res[0] = utils.coloText(fmt.Sprintf("Number of lines %d", lineCount))
	res[1] = fmt.Sprintf("Number of files %d", fileCount)
	return res, nil
}

func countFilesAndLines(root string) (int, int, error) {
	var fileCount, lineCount int

	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		fileCount++
		lines, err := countLines(path)
		if err != nil {
			return nil
		}
		lineCount += lines
		return nil
	})

	return fileCount, lineCount, err
}
func countLines(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lines := 0
	for scanner.Scan() {
		lines++
	}
	return lines, scanner.Err()
}