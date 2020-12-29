package main

import "fmt"

func main() {
	path := GetCondaPath()
	envs,err := GetCondaEnvs(path)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(envs)

	userPath,err := GetUserHome()
	if err !=nil{
		fmt.Println(err)
	}
	fmt.Println(userPath)
}