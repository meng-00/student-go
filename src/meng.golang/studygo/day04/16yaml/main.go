package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func readJson(filePath string) (result string) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	buf := bufio.NewReader(file)
	for {
		s, err := buf.ReadString('\n')
		result += s
		if err != nil {
			if err == io.EOF {
				fmt.Println("Read is ok")
				break
			} else {
				fmt.Println("ERROR:", err)
				return
			}
		}
	}
	return result
}

func main() {
	var data struct {
		Id         int
		Customerld int
		Raised     string
		Due        string
		Paid       bool
		Note       string
		Items      []struct {
			Id       string
			Price    float64
			Quantity int
			Note     string
		}
	}
	result := readJson("C:\\GolangProject\\io\\file_option\\json\\json.json")
	err := json.Unmarshal([]byte(result), &data)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	fmt.Println(data)
}
