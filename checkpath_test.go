// package main contains tests
package main

import "testing"

func Test_existsFound(t *testing.T) {
	exist, err := exists("checkpath_test.go")

	if err != nil {
		t.Error("got error", err)
	}

	if exist != true {
		t.Error("exist is", exist)
	}
}

func Test_existsNotFound(t *testing.T) {
	exist, err := exists("checkpath_test_NOTTHERE.go")

	if err != nil {
		t.Error("got error", err)
	}

	if exist != false {
		t.Error("exist is", exist)
	}
}
