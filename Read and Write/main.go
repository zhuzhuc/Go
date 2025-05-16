package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("/Users/Apple/Desktop/Go/Read and Write/text.txt")
	defer file.Close() // 延迟关闭文件句柄,必须关闭文件
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}
	fmt.Println(file)
	// 读取文件内容
	// 1. 首先创建一个字节切片，用于存储读取到的内容
	// 2. 然后调用 file.Read() 方法，将文件内容读取到字节切片中
	// 3. 如果读取成功，返回读取到的字节数和 nil 错误
	// 4. 如果读取失败，返回读取到的字节数和错误信息
	// 5. 最后打印读取到的字节数和内容
	slice := []byte{}
	tempSlice := make([]byte, 1024)
	// num, err := file.Read(tempSlice)
	// if err != nil {
	// 	fmt.Println("读取文件失败:", err)
	// 	return
	// }
	// fmt.Println("读取的字节数:", num)
	// fmt.Println("读取的内容:", string(tempSlice[:num]))

	for {
		num, err := file.Read(tempSlice)
		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}
		if err != nil {
			fmt.Println("读取文件失败:", err)
			return
		}
		slice = append(slice, tempSlice[:num]...)
	}
	fmt.Println("读取的内容:", string(slice))

	fmt.Println("-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------")

	ff, err := os.Open("/Users/Apple/Desktop/Go/Read and Write/text2.txt")
	defer ff.Close() // 延迟关闭文件句柄,必须关闭文件
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}
	reader := bufio.NewReader(ff)
	// strr, err := reader.ReadString('\n') // 一次读取一行
	// if err == io.EOF {
	// 	fmt.Println("读取完毕")
	// 	return
	// }
	// if err != nil {
	// 	fmt.Println("读取文件失败:", err)
	// 	return
	// }
	var fff string
	for {
		strr, err := reader.ReadString('\n') // 一次读取一行
		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}
		if err != nil {
			fmt.Println("读取文件失败:", err)
			return
		}
		fff += strr
	}
	fmt.Println(fff)

	fmt.Println("-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------")

	bytestr, err := ioutil.ReadFile("/Users/Apple/Desktop/Go/Read and Write/text3.txt")
	if err != nil {
		fmt.Println("读取文件失败:", err)
		return
	}
	fmt.Println(string(bytestr))

	fmt.Println("-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------")

	// 写入文件
	// 1. 首先打开文件，获取文件句柄
	// 2. 然后调用 file.Write() 方法，将内容写入文件
	// 3. 如果写入成功，返回写入的字节数和 nil 错误
	// 4. 如果写入失败，返回写入的字节数和错误信息
	// 5. 最后关闭文件句柄
	fmt.Println("写入文件")

	www, err := os.OpenFile("/Users/Apple/Desktop/Go/Read and Write/1.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	defer www.Close() // 延迟关闭文件句柄,必须关闭文件
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}

	_, err = www.Write([]byte(`
                          '--'   '--'                           '--'   '--'                           '--'   '--'                 
`))
	if err != nil {
		fmt.Println("写入文件失败:", err)
		return
	}
	fmt.Println("写入完毕")
	fmt.Println("-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------")
	// 追加文件
	// 1. 首先打开文件，获取文件句柄
	// 2. 然后调用 file.Write() 方法，将内容写入文件
	// 3. 如果写入成功，返回写入的字节数和 nil 错误
	// 4. 如果写入失败，返回写入的字节数和错误信息
	// 5. 最后关闭文件句柄
	fmt.Println("追加文件")
	www, err = os.OpenFile("/Users/Apple/Desktop/Go/Read and Write/1.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer www.Close() // 延迟关闭文件句柄,必须关闭文件
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}
	_, err = www.Write([]byte(`
                          '--'   '--'                           '--'   '--'                           '--'   '--'
`))
	if err != nil {
		fmt.Println("写入文件失败:", err)
		return
	}
	fmt.Println("写入完毕")
	fmt.Println("-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------")

	// 追加文件内容，使用带缓冲的写入器
	fsf, err := os.OpenFile("/Users/Apple/Desktop/Go/Read and Write/1.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// 延迟关闭文件句柄，确保文件在函数结束时关闭
	defer fsf.Close()
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}

	// 创建带缓冲的写入器，使用新的变量名 writer
	writer := bufio.NewWriter(fsf)

	// 向缓冲区写入字符串
	n, err := writer.WriteString(`
	                          '--'   '--'                           '--'   '--'                           '--'   '--'`)
	if err != nil {
		fmt.Println("写入缓冲区失败:", err)
		return
	}
	fmt.Printf("成功写入 %d 个字节到缓冲区\n", n)

	// 刷新缓冲区，将内容写入文件
	err = writer.Flush()
	if err != nil {
		fmt.Println("刷新缓冲区失败:", err)
		return
	}
	fmt.Println("缓冲区内容已成功写入文件")

	str := "hello world"
	errs := ioutil.WriteFile("/Users/Apple/Desktop/Go/Read and Write/2.txt", []byte(str), 0666)
	if errs != nil {
		fmt.Println("写入文件失败:", err)
		return
	}
	fmt.Println("ioutil()写入完毕")
}
