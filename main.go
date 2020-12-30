package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var (
	list bool
	help bool
	add string
	display string
)

func main() {
	flag.Parse()
	if help || len(os.Args) == 1{
		flag.Usage()
	}
	if list{
		env,err := GetCondaEnvs()
		if err!=nil{
			fmt.Println(err)
			return
		}
		fmt.Println("环境名称\t环境所在位置")
		for i,e := range env{
			fmt.Printf("%s\t%s\r\n",i,e)
		}
	}

	if add != ""{
		// 获取用户家目录
		path,err := GetUserHome()
		if err != nil {
			fmt.Println(err)
			return
		}
		displayName := add
		if display!=""{
			displayName = display
		}
		env,err := GetCondaEnvs()
		if err!=nil{
			fmt.Println(err)
			return
		}

		envPath,ok := env[add]
		if ok{
			fmt.Println("安装必须的模块....")
			command := exec.Command(envPath,"-m","pip","install","ipykernel")
			err = command.Run()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println("添加内核配置文件")
			err = GenerateConfigFile(path,envPath,displayName)
			if err != nil {
				fmt.Println(err)
				return
			}
		}else{
			fmt.Printf("环境变量\"%s\"不存在\r\n",add)
		}
		fmt.Println("添加成功")
	}
}

func init()  {
	flag.BoolVar(&help,"h",false,"显示本帮助信息")
	flag.BoolVar(&list,"l",false,"列出所有环境")
	flag.StringVar(&add,"add","","添加一个内核，后跟环境名称")
	flag.StringVar(&display,"dn","","Jupyter中显示的名称，默认为环境名称")
}
