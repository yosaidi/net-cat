package server

import "fmt"

func HandleError(err error) {
	if err != nil {
		fmt.Println("Error encountered: ", err)
		return
	}
}
