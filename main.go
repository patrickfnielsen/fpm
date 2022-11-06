package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
	"github.com/patrickfnielsen/fpm/internal/types"
)

func main() {
	file, err := os.ReadFile("settings.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal our input YAML file into empty Settings struct
	var settings Settings
	if err := yaml.Unmarshal(file, &settings); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", settings)
}
