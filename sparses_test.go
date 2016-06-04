package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

const (
	ExpectedSparse = "sparse-file"
	Testdata       = "testdata"
)

// TestMain creates and removes the testdata files
func TestMain(m *testing.M) {
	flag.Parse()
	removeFiles()
	createFiles()
	os.Exit(m.Run())
	removeFiles()
}

// TestSparses searchs for sparse files.
// Will Only work on FS with sparse file support like in Linux (Not in OSX)
func TestSparses(t *testing.T) {
	expected := filepath.Join(Testdata, ExpectedSparse)
	actual, err := FindSparses(Testdata)
	if err != nil {
		t.Fatalf("Error in Sparses call!: %s\n", err)
	}
	if len(actual) != 1 || actual[0] != expected {
		t.Fatalf("Expected [ %s] but got %v\n", expected, actual)
	}
}

// creates testdata/sparse-file & testdata/zeroes-file
func createFiles() {
	zero := []byte{0}
	one := []byte{1}
	// directory
	err := os.MkdirAll(Testdata, 0750)
	dieOnError(err)
	// files
	createTestFile(ExpectedSparse, zero, 1, 24576)
	createTestFile("zeroes-file", zero, 4098, 0)
	createTestFile("ones-file", one, 1090, 0)
}

// createTestFile creates a file named filename & writes buf n times from pos
func createTestFile(filename string, buf []byte, n int, pos int64) {
	f, err := os.Create(filepath.Join(Testdata, filename))
	dieOnError(err)
	defer f.Close()
	_, err = f.Seek(pos, 0)
	dieOnError(err)
	_, err = f.Write(buf)
	dieOnError(err)
}

// creates testdata/sparse-file & testdata/zeroes-file
func removeFiles() {
	err := os.RemoveAll(Testdata)
	dieOnError(err)
}

// dieOnError reports error and exists if error e is not nil
func dieOnError(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(-1)
	}
}
