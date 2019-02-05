package logger

import (
	"fmt"
	"os"
)

func Log(msg string) bool {
	// file descriptor
	// os.O_APPEND|os.O_CREATE|os.O_WRONLY are flags to indicate what to do with that file
	// 0644 is filemode when its created
	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// close the file descriptor at the end of the function
	defer f.Close()

	if err != nil {
		fmt.Println(err)
		return false
	}

	// write string to file
	_, err = f.WriteString(fmt.Sprintf("%s\n", msg))
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
