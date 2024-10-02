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

	// Read immediate subdirectories
	entries, err := os.ReadDir(baseDir)
	if err != nil {
		log.Fatalf("Error reading base directory: %v", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			subDir := filepath.Join(baseDir, entry.Name())
			kustomizationPath := filepath.Join(subDir, "kustomization.yaml")

			// Check if kustomization.yaml exists in the subdirectory
			if _, err := os.Stat(kustomizationPath); err == nil {
				// Process the kustomization.yaml file
				payload, err := processKustomizationFile(kustomizationPath)
				if err != nil {
					log.Printf("Error processing file %s: %v", kustomizationPath, err)
					continue
				}
				payloads = append(payloads, payload)
			}
		}
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