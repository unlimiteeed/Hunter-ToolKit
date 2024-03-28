package yamlReader

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Configuration struct {
	Tools     map[string]bool `yaml:"Tools"`
	WordLists map[string]bool `yaml:"WordLists"`
}

func ReadFunction() Configuration {
	file, err := os.Open("config/installation.yaml")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	// Read the content of the file into a byte slice
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading file content: %v", err)
	}

	config := Configuration{}

	// Unmarshal the byte slice containing YAML data
	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatalf("Error unmarshalling YAML: %v", err)
	}

	return config

}
