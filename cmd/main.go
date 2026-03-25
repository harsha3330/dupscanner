package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"
)

type Config struct {
	Workers int
	Dir     string
	MinSize int64
	Ext     string
}

type AppError struct {
	Path string
	Err  error
}

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: dupscanner --dir <path> [--workers N] [--min-size bytes] [--ext .ext]")
		flag.PrintDefaults()
	}

	var config Config

	flag.StringVar(&config.Dir, "dir", ".", "directory to scan")
	flag.IntVar(&config.Workers, "workers", runtime.NumCPU(), "number of concurrent workers")
	flag.Int64Var(&config.MinSize, "min-size", 0, "minimum file size in bytes")
	flag.StringVar(&config.Ext, "ext", "", "filter by file extension (e.g. .jpg)")

	flag.Parse()

	if config.Workers <= 0 {
		fmt.Println("Error: workers must be > 0")
		os.Exit(1)
	}

	if config.MinSize < 0 {
		fmt.Println("Error: min-size must be >= 0")
		os.Exit(1)
	}
	config.Ext = strings.ToLower(config.Ext)

	results, appErrors := ExtractResults(&config)

	printResults(results, appErrors)
}

func printResults(results []Result, errorsList []AppError) {
	fmt.Println()
	hashMap := make(map[string][]string)

	for _, r := range results {
		hashMap[r.Hash] = append(hashMap[r.Hash], r.File)
	}

	fmt.Println("Duplicate Files Found:")

	totalGroups := 0

	for hash, files := range hashMap {
		if len(files) < 2 {
			continue
		}

		totalGroups++

		fmt.Printf("Hash: %s\n", hash)
		for _, f := range files {
			fmt.Printf("  %s\n", f)
		}
		fmt.Println()
	}

	fmt.Printf("Total duplicate groups: %d\n", totalGroups)

	fmt.Printf("\nErrors: %d\n", len(errorsList))

	for _, e := range errorsList {
		fmt.Printf("  %s: %v\n", e.Path, e.Err)
	}
}
