package main

import (
	"os"
	"path/filepath"
	"syscall"
)

const SectorSize = 512

// FindSparses returns all found sparse files in the given path
func FindSparses(path string) ([]string, error) {
	sparses := []string{}
	dir, err := os.Open(path) // For read access.
	if err != nil {
		return nil, err
	}
	files, err := dir.Readdir(0)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		filename := filepath.Join(path, f.Name())
		if f.IsDir() {
			dirSparses, err := FindSparses(filename)
			if err != nil {
				return nil, err
			}
			sparses = append(sparses, dirSparses...)
		} else {
			isSparse, err := IsSparse(filename)
			if err != nil {
				return nil, err
			}
			if isSparse {
				sparses = append(sparses, filename)
			}
		}
	}
	return sparses, nil
}

// IsSparse returns true if the given file is a Sparse file
func IsSparse(filename string) (bool, error) {
	stat := &syscall.Stat_t{}
	if err := syscall.Stat(filename, stat); err != nil {
		return false, err
	}
	correction := stat.Blksize / SectorSize
	actualBlocks := stat.Blocks / int64(correction)
	/*
		fmt.Printf("%s Stat_t {Size %d, BlkSize %d Blocks %d Actual Blocks %d Sparse %v}\n", filename,
			stat.Size,
			stat.Blksize,
			stat.Blocks, actualBlocks,
			(stat.Size > (actualBlocks * int64(stat.Blksize))))
	*/
	return stat.Size > (actualBlocks * int64(stat.Blksize)), nil
}
