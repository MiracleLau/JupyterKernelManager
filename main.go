package main

import "fmt"

func main() {
	path := GetCondaPath()
	envs,err := GetCondaEnvs(path)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(envs)
}