# 🔍 dupscanner — Concurrent File Deduplicator (Go)

A fast CLI tool to find duplicate files by content using Go concurrency
---
## Features

- Recursive directory scanning
- Concurrent file hashing (SHA256)
- Detects duplicate files by content
- Configurable worker pool
- Early filtering (file size, extension)
- Deterministic output

---

## Installation

```bash
git clone <your-repo-url>
cd dupfind
go build -o dupfind ./cmd

---

## Usage 

```
./dupfind --dir ./downloads --workers 8


## Flags

| Flag         | Description                            | Default   |
| ------------ | -------------------------------------- | --------- |
| `--dir`      | Directory to scan                      | `.`       |
| `--workers`  | Number of concurrent workers           | NumCPU()  |
| `--min-size` | Ignore files smaller than N bytes      | `0`       |
| `--ext`      | Filter by extension (e.g. `.jpg`)      | all files |


## Sample Output 

Scanned 120 files

Duplicate Files Found:

Hash: a1b2c3
  ./downloads/file1.txt
  ./downloads/copy/file1.txt

Hash: x9y8z7
  ./downloads/img.png
  ./downloads/backup/img.png

Total duplicate groups: 2