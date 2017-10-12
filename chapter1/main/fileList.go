package main

import (
	"fmt"
	"readFile"
)

func main() {
	r, args := readFile.New(), []int{}
	file := "1kints.txt"
	path, err := r.GetDataPath()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = r.ReadToInt(path+file, &args)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(args[0], args[9])
}
