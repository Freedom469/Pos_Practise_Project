package pos

import "fmt"

func ErrHandler(err error, message string) {
	if err != nil {
		fmt.Println(message)
	}
}