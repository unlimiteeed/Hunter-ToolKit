package yamlReader

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func ReadFunction() {
	obj := make(map[string]interface{})

	ConfigFile, err := ioutil.ReadFile("/config/installation.yaml")
	if err != nil {
		fmt.Printf("We Have Error %v", err)

	}
	err = yaml.Unmarshal(ConfigFile, obj)
	if err != nil {
		fmt.Println("Unmarsall Faild")
	}
	fmt.Println(obj)
}
