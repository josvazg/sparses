package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"testing"
)

const Testdata = "testdata"

func TestMain(m *testing.M) {
	flag.Parse()
	removeFiles()
	createFiles()
	os.Exit(m.Run())
	removeFiles()
}

func TestSparses(t *testing.T) {
	expected := Testdata + "/sparse-file"
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
	// sparse-file
	sf, err := os.Create(Testdata + "/sparse-file")
	dieOnError(err)
	defer sf.Close()
	_, err = sf.Seek(24576, 0)
	dieOnError(err)
	_, err = sf.Write(zero)
	dieOnError(err)
	// zeroes-file
	zf, err := os.Create(Testdata + "/zeroes-file")
	dieOnError(err)
	defer zf.Close()
	for i := 0; i <= 4098; i++ {
		_, err := zf.Write(zero)
		dieOnError(err)
	}
	// ones-file
	of, err := os.Create(Testdata + "/ones-file")
	dieOnError(err)
	defer of.Close()
	for i := 0; i <= 1090; i++ {
		_, err := of.Write(one)
		dieOnError(err)
	}
	// ls to check sizes on command line
	run("ls", "-ls", Testdata)
}

// creates testdata/sparse-file & testdata/zeroes-file
func removeFiles() {
	err := os.RemoveAll(Testdata)
	dieOnError(err)
}

func run(cmd string, args ...string) {
	out, err := exec.Command(cmd, args...).CombinedOutput()
	dieOnError(err)
	fmt.Println(string(out))
}

// dieOnError reports error and exists if error e is not nil
func dieOnError(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(-1)
	}
}
