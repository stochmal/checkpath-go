// package main is entry point for console application
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"sort"
)

const copyright = "Copyright (C) 2021 Tomasz Stochmal <stochmal@gmail.com>"
const license = "https://github.com/stochmal/checkpath-go/blob/main/LICENSE"

// main contains application logic
func main() {

	fmt.Println(copyright)
	fmt.Println(license)

	fmt.Printf("\nos name is %s-%s\n", runtime.GOOS, runtime.GOARCH)

	var path = os.Getenv("PATH")
	fmt.Printf("\nPATH len is %d (4095 max allowed)\n", len(path))

	fmt.Println("\n*** folders that can be removed from PATH")
	fmt.Println()

	folders := filepath.SplitList(path)

	sort.Strings(folders)

	var last_folder = ""
	for _, folder := range folders {

		exist, _ := exists(folder)

		if !exist {
			fmt.Printf("%-25s %s\n", "*** FOLDER NOT FOUND", folder)
		} else if last_folder == folder {
			fmt.Printf("%-25s %s\n", "*** DUPLICATE", folder)
		} else {
			files, _ := ioutil.ReadDir(folder)

			if len(files) == 0 {
				fmt.Printf("%-25s %s\n", "*** FOLDER IS EMPTY", folder)
			} else {
				fmt.Printf("%9d files found %3s %s\n", len(files), "", folder)
			}
		}

		last_folder = folder
	}
}

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
