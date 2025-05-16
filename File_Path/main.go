package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func copy(src string, dist string) (err error) {
	byte, err := ioutil.ReadFile(src)
	if err != nil {
		fmt.Println("读取文件失败:", err)
		return err
	}
	err2 := ioutil.WriteFile(dist, byte, 0666)
	if err2 != nil {
		fmt.Println("写入文件失败:", err2)
		return err2
	}
	return nil
}

func Copyfile(src string, dist string) (err error) {
	fmt.Println("---------------------------------")
	file, err := os.Open(src)
	defer file.Close() // 延迟关闭文件句柄,必须关闭文件
	file2, err2 := os.OpenFile(dist, os.O_CREATE|os.O_WRONLY, 0666)
	defer file2.Close() // 延迟关闭文件句柄,必须关闭文件
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return err
	}
	if err2 != nil {
		fmt.Println("打开文件失败:", err2)
		return err2
	}
	tempSlice := make([]byte, 1024)
	for {
		n, err4 := file.Read(tempSlice)
		if err4 == io.EOF {
			fmt.Println("读取完毕")
			break
		}
		if err4 != nil {
			fmt.Println("读取文件失败:", err4)
			return err4
		}
		if _, err5 := file2.Write(tempSlice[:n]); err5 != nil {
			fmt.Println("写入文件失败:", err5)
			return err5
		}

	}
	fmt.Println("CopyFile复制文件成功")
	fmt.Println("---------------------------------")
	return nil
}

func main() {
	src := "/Users/Apple/Desktop/Go/Read and Write/3.txt"
	dist := "/Users/Apple/Desktop/Go/Read and Write/4.txt"
	err := copy(src, dist)
	if err != nil {
		fmt.Println("复制文件失败:", err)
		return
	}
	fmt.Println("复制文件成功")
	err2 := Copyfile(src, dist)
	if err2 != nil {
		fmt.Println("复制文件失败:", err2)
		return
	}
	fmt.Println("复制文件成功")
	fmt.Println("---------------------------------")

	// errss := os.Mkdir("./abc", 0666)
	// if errss != nil {
	// 	fmt.Println("创建目录失败:", errss)
	// 	return
	// }
	os.MkdirAll("./abc/1/2/3/4", 0666)
	fmt.Println("创建目录成功")
	os.MkdirAll("./a", 0666)

	os.Remove("./a")
	// os.RemoveAll("./abc")

	os.Rename("./abc", "./abc2")
}
