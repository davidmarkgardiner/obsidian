```go

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Kustomization struct {
	CommonLabels struct {
		OpEnvironment string `yaml:"opEnvironment"`
	} `yaml:"commonLabels"`
	CommonAnnotations struct {
		AksClusterResourceId string `yaml:"aksClusterResourceId"`
		FullNamespaceName    string `yaml:"fullNamespaceName"`
	} `yaml:"commonAnnotations"`
}

type Payload struct {
	AksResourceId   string `json:"aksresourceid"`
	OpenEnvironment string `json:"openvironment"`
	NamespaceName   string `json:"namespacename"`
}

func main() {
	// Get region and clustername from environment variables
	region := os.Getenv("REGION")
	clusterName := os.Getenv("CLUSTER_NAME")

	if region == "" || clusterName == "" {
		log.Fatal("REGION and CLUSTER_NAME environment variables must be set")
	}

	// Base directory
	baseDir := filepath.Join("environment", region, clusterName)

	var payloads []Payload

	// Walk through the directory structure
	err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the file is kustomization.yaml
		if info.IsDir() || filepath.Base(path) != "kustomization.yaml" {
			return nil
		}

		// Process each kustomization.yaml file
		payload, err := processKustomizationFile(path)
		if err != nil {
			log.Printf("Error processing file %s: %v", path, err)
			return nil
		}

		payloads = append(payloads, payload)
		return nil
	})

	if err != nil {
		log.Fatalf("Error walking through directory: %v", err)
	}

	// Write payloads to JSON file
	outputFile := "payloads.json"
	err = writePayloadsToFile(payloads, outputFile)
	if err != nil {
		log.Fatalf("Error writing payloads to file: %v", err)
	}

	fmt.Printf("Processed %d kustomization files. Output written to %s\n", len(payloads), outputFile)
}

func processKustomizationFile(path string) (Payload, error) {
	// Read and parse the kustomization.yaml file
	data, err := os.ReadFile(path)
	if err != nil {
		return Payload{}, fmt.Errorf("error reading file: %v", err)
	}

	var kustomization Kustomization
	err = yaml.Unmarshal(data, &kustomization)
	if err != nil {
		return Payload{}, fmt.Errorf("error unmarshaling YAML: %v", err)
	}

	// Extract required information
	payload := Payload{
		AksResourceId:   kustomization.CommonAnnotations.AksClusterResourceId,
		OpenEnvironment: kustomization.CommonLabels.OpEnvironment,
		NamespaceName:   kustomization.CommonAnnotations.FullNamespaceName,
	}

	log.Printf("Processed file %s - aksresourceid: %s, openvironment: %s, namespacename: %s", 
		path, payload.AksResourceId, payload.OpenEnvironment, payload.NamespaceName)

	return payload, nil
}

func writePayloadsToFile(payloads []Payload, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pretty print the JSON
	if err := encoder.Encode(payloads); err != nil {
		return fmt.Errorf("error encoding JSON: %v", err)
	}

	return nil
}