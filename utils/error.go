package utils

import "fmt"

func Catch(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
