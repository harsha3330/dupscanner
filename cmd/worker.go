package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func hashFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	hasher := sha256.New()

	_, err = io.Copy(hasher, f)
	if err != nil {
		return "", err
	}

	hash := hasher.Sum(nil)
	return fmt.Sprintf("%x", hash), nil
}

func HashWorker(
	config *Config,
	files <-chan string,
	results chan<- Result,
	errs chan<- AppError,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for path := range files {
		info, err := os.Stat(path)
		if err != nil {
			errs <- AppError{
				Path: path,
				Err:  fmt.Errorf("worker: %w", err),
			}
			continue
		}

		if config.MinSize > 0 && info.Size() < config.MinSize {
			continue
		}

		if config.Ext != "" && strings.ToLower(filepath.Ext(path)) != config.Ext {
			continue
		}

		hash, err := hashFile(path)
		if err != nil {
			errs <- AppError{
				Path: path,
				Err:  fmt.Errorf("worker: %w", err),
			}
			continue
		}

		results <- Result{
			Hash: hash,
			File: path,
		}
	}
}
