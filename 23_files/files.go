package main

import (
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("file name:", fileInfo.Name())
	// fmt.Println("file or dir:", fileInfo.IsDir())
	// fmt.Println("file size:", fileInfo.Size())
	// fmt.Println("file permissions:", fileInfo.Mode())
	// fmt.Println("file modified at:", fileInfo.ModTime())

	// Read file 

	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	
	// defer f.Close()

	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	panic(err)
	// }
	// size := fileInfo.Size()
	
	// buf := make([]byte, size)

	// d, err := f.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("read", d, "bytes")
	// fmt.Println("data", string(buf))

	f, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(f))

}
