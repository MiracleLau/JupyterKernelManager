package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

const PATH_SEPARATOR = string(os.PathSeparator)

type KernelJson struct {
	Argv []string	`json:"argv"`
	DisplayName string	`json:"display_name"`
	Language string `json:"language"`
}

// 获取所有的conda路径
func GetCondaPath() string {
	var condaPath = ""
	if path := os.Getenv("PATH"); path != "" {
		// 分割变量
		envs := strings.Split(path, ";")
		for _, env := range envs {
			if IsFileExist(strings.Join([]string{env, "Scripts", "conda.exe"}, PATH_SEPARATOR)) {
				condaPath = env
			}
		}
	}
	return condaPath
}

// 判断是否存在指定文件
func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// 获取所有的conda环境
func GetCondaEnvs() (map[string]string,error) {
	path := GetCondaPath()
	envPath := strings.Join([]string{path, "envs"}, PATH_SEPARATOR)
	dir, err := ioutil.ReadDir(envPath)
	if err != nil {
		return nil, err
	}
	var envs map[string]string = make(map[string]string)
	for _, p := range dir {
		if p.IsDir() {
			pythonPath := strings.Join([]string{envPath,p.Name(),"python.exe"},PATH_SEPARATOR)
			envs[p.Name()] = pythonPath
		}
	}
	return envs, nil
}

// 获取用户家目录
func GetUserHome() (string, error) {
	path, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return path, nil
}

// 在用户目录生成配置文件
func GenerateConfigFile(userHome string,envPath string,displayName string) error  {
	path := strings.Join([]string{userHome,"AppData","Roaming","jupyter","kernels",displayName},PATH_SEPARATOR)
	if !IsFileExist(path){
		// 不存在就创建目录
		err := os.MkdirAll(path,os.ModePerm)
		if err != nil{
			return err
		}
	}
	kernelConfig := KernelJson{
		Argv: []string{
			envPath,
			"-m",
			"ipykernel_launcher",
			"-f",
			"{connection_file}",
		},
		DisplayName: displayName,
		Language: "python",
	}
	jsonBytes,err := json.Marshal(kernelConfig)
	if err!=nil{
		return err
	}
	err = ioutil.WriteFile(strings.Join([]string{path,"kernel.json"},PATH_SEPARATOR),jsonBytes,0644)
	if err!=nil{
		return err
	}
	err = CopyFile("./Images/logo-32x32.png",strings.Join([]string{path,"logo-32x32.png"},PATH_SEPARATOR))
	if err != nil {
		return err
	}
	err = CopyFile("./Images/logo-64x64.png",strings.Join([]string{path,"logo-64x64.png"},PATH_SEPARATOR))
	if err != nil {
		return err
	}
	return nil
}

// 复制文件
func CopyFile(src string,dst string) error {
	input, err := ioutil.ReadFile(src)
	if err!=nil{
		return err
	}
	err = ioutil.WriteFile(dst, input, 0644)
	if err!=nil{
		return err
	}
	return nil
}