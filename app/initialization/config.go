package initialization

import (
	"io/ioutil"
	"log"
	"secKill/app/common"

	"gopkg.in/yaml.v3"
)

func initConfig() {
	fileContent, err := ioutil.ReadFile("./conf/conf.yaml")
	if err != nil {
		log.Printf("ioutil.ReadFile failed. err:[%+v]\n", err)
		panic(err)
	}

	err = yaml.Unmarshal(fileContent, &common.Config)
	if err != nil {
		log.Printf("yaml.Unmarshal failed. err:[%+v]\n", err)
		panic(err)
	}
}
