package utils

import (
	"fmt"
	"os"
)

// HandleErr is a helper function to print errors
func HandleErr(err error) {
	//TODO: print this is red color
	fmt.Println(err)
	os.Exit(0)
}
