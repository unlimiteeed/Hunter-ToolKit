package yamlReader

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Configuration struct {
	Tools     map[string]bool `yaml:"Tools"`
	WordLists map[string]bool `yaml:"WordLists"`
}

func ReadFunction() {
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

	fmt.Println("Tools:")
	for tool, available := range config.Tools {
		fmt.Printf("%s - %v\n", tool, available)
	}

	fmt.Println("Wordlists:")
	for wordlist, available := range config.WordLists {
		fmt.Printf("%s - %v\n", wordlist, available)
	}
}
