package boot

import (
	"io/ioutil"
	"log"
	"os"
)

// LocalConfigFile 读取配置文件
func LocalConfigFile() []byte {
	dir, _ := os.Getwd()
	file := dir + "/application.yaml"
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println(err)
		return nil
	}
	return b
}
