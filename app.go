package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"golang.design/x/clipboard"
)

func main(){
	var password_size int
	var err_convert error

	if len(os.Args) < 2 || os.Args[1] == ""{
		fmt.Println("Enter password size")
		fmt.Scan(&password_size)
	} else {
		password_size,_ = strconv.Atoi(os.Args[1])

		if err_convert != nil{
			panic("Not able to convert OS Arg to int")
		}

	}
	
	url := fmt.Sprintf("https://www.random.org/strings/?num=1&len=%d&digits=on&upperalpha=on&loweralpha=on&unique=on&format=plain&rnd=new", password_size) 

	resp, err_api := http.Get(url)

	if err_api != nil{
		fmt.Println("Unable to parse random.org API")
	}
	password_byte, _ := io.ReadAll(resp.Body)

	password := string(password_byte)

	fmt.Println((string(password)))

	clipboard.Write(clipboard.FmtText,[]byte(password))
	}
