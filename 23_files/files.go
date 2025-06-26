package main

import (
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

	// f, err := os.ReadFile("example.txt") // not good for optimization
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(f))

	// Read Folders

	// dir, err := os.Open("../")

	// if err != nil {
	// 	panic(err)
	// }

	// defer dir.Close()

	// fileInfo, err := dir.ReadDir(-1)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, fi := range fileInfo {
	// 	fmt.Println(fi.Name(), fi.IsDir())
	// }

	f, err := os.Create("example2.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// f.WriteString("hello golang, from files.go ")
	// f.WriteString("go is nice language")

	bytes := []byte("Hello Golang")
	f.Write(bytes)

	

}
