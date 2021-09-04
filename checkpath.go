package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"sort"
	"strings"
)

const COPYRIGHT = "Copyright (C) 2021 Tomasz Stochmal <stochmal@gmail.com>"
const LICENSE = "https://github.com/stochmal/checkpath-go/blob/main/LICENSE"

func main() {

	var separator = ":"
	if runtime.GOOS == "windows" {
		separator = ";"
	}

	fmt.Println(COPYRIGHT)
	fmt.Println(LICENSE)

	fmt.Printf("\nos name is %s-%s use separator %s\n", runtime.GOOS, runtime.GOARCH, separator)

	var path = os.Getenv("PATH")
	fmt.Printf("\nPATH len is %d (4095 max allowed)\n", len(path))

	fmt.Println("\n*** folders that can be removed from PATH")
	fmt.Println()

	folders := strings.Split(path, separator)
	sort.Strings(folders)

	for _, folder := range folders {

		exist, _ := exists(folder)

		if !exist {
			fmt.Printf("%-25s %s\n", "*** FOLDER NOT FOUND", folder)
		} else {
			files, _ := ioutil.ReadDir(folder)

			if len(files) == 0 {
				fmt.Printf("%-25s %s\n", "*** FOLDER IS EMPTY", folder)
			} else {

				fmt.Printf("%9d files found %3s %s\n", len(files), "", folder)
			}
		}
	}
	//	fmt.Println("press [Enter] to exit")
	//	fmt.Scanln()
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
