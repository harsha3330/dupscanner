package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func FileProducer(dir string) (<-chan string, <-chan AppError) {
	files := make(chan string)
	errs := make(chan AppError)
	go func() {
		defer close(files)
		defer close(errs)

		var walk func(string)

		walk = func(current string) {
			entries, err := os.ReadDir(current)
			if err != nil {
				errs <- AppError{
					Err:  fmt.Errorf("producer: %w", err),
					Path: current,
				}
				return
			}

			for _, entry := range entries {
				fullPath := filepath.Join(current, entry.Name())

				if entry.IsDir() {
					walk(fullPath)
				} else {
					files <- fullPath
				}
			}
		}

		walk(dir)
	}()

	return files, errs
}
