package main

import (
	"io/ioutil"
	"os"
	"strings"
)

const PATH_SEPARATOR = string(os.PathSeparator)
// 获取所有的conda路径
func GetCondaPath() string  {
	var condaPath = ""
	if path := os.Getenv("PATH"); path != "" {
		// 分割变量
		envs := strings.Split(path,";")
		for _,env := range envs{
			if IsFileExist(strings.Join([]string{env,"Scripts","conda.exe"},PATH_SEPARATOR)){
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
func GetCondaEnvs(path string) ([]string,error)  {
	envPath := strings.Join([]string{path,"envs"},PATH_SEPARATOR)
	dir,err := ioutil.ReadDir(envPath)
	if err != nil{
		return nil, err
	}
	var envs []string
	for _,p := range dir{
		if p.IsDir(){
			envs = append(envs, p.Name())
		}
	}
	return envs,nil
}