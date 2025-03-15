package utils

import (
	"fmt"
	"os"
)

func DeleteFile(path string) {
	err := os.Remove(path)
	if err != nil {
		fmt.Println(err.Error())
	}
}
