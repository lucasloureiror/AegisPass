package main

import (
	"fmt"
	"golang.design/x/clipboard"
)

func main(){

	fmt.Println("Enter password size")
	var password_size int
	fmt.Scan(&password_size)
	
	var password string

	for i:= 1; i <= password_size; i++{
		var digit string
		fmt.Scan(&digit)
		password = password + digit
	} 

	fmt.Println("Password:",password)
	fmt.Println(string(password[0]))
	err := clipboard.Init()
	if err != nil {
		fmt.Println("Não será possível copiar para o clipboard")
	}

	clipboard.Write(clipboard.FmtText,[]byte(password))

}