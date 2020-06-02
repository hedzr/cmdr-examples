/*
 */

// Copyright © 2020 Hedzr Yeh.

package unsorted

import (
	"log"
	"os"
	"path/filepath"
)

// GetExecutableDir returns this executable file pathname
func GetExecutableDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(dir)
	return dir
}

// GetCurrentDir returns the current os working directory pathname.
func GetCurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(dir)
	return dir
}

// FileExists returns true if the file `name` was existed.
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
