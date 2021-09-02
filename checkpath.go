package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
)

func main() {

	var separator = ":"
	if runtime.GOOS == "windows" {
		separator = ";"
	}

	fmt.Printf("os name is %s-%s use separator %s\n", runtime.GOOS, runtime.GOARCH, separator)

	fmt.Println("\n*** folders that can be removed from PATH")
	fmt.Println()

	folders := strings.Split(os.Getenv("PATH"), separator)

	for _, folder := range folders {

		exist, _ := exists(folder)

		if !exist {
			fmt.Println("*** FOLDER NOT FOUND", "\t", folder)
		} else {
			files, _ := ioutil.ReadDir(folder)

			if len(files) == 0 {
				fmt.Println("*** FOLDER IS EMPTY", "\t", folder)
			} else {

				fmt.Printf("%6d files found\t%s\n", len(files), folder)
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
