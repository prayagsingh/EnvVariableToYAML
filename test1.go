package main

import (
	//"github.com/go-yaml/yaml"
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"

	"github.com/ghodss/yaml"
	"github.com/joho/godotenv"
)

// Params :- to store the env variables
type Params struct {
	ORDERER_PROFILE string
	ORD_CHANNEL_ID  string
	CHANNEL_PROFILE string
	CHANNEL_NAME    string
	ORG1_NAME       string
	ORG2_NAME       string
	DOMAIN          string
	PROJECT_NAME    string
}

// Taking User input if a particular value not found in env, kind of fallback func
func readInput(message string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Value of ", message, " not found. Please enter the value\n")
	scanner.Scan()
	text := scanner.Text()
	fmt.Print("\n Value of text is: ", text)
	return text

}

// Simple helper function to read an environment or return a default value
func getEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = readInput(key)
	}
	return value
}

func main() {
	fmt.Println("\n Reading env file ")
	err := godotenv.Load(os.ExpandEnv("../.env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	confContent, err := ioutil.ReadFile("docker-compose-e2e-template.tmpl.yaml")
	if err != nil {
		panic(err)
	}
	j, err := yaml.YAMLToJSON(confContent)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Print(string(j), "\n\n")

	// checking if the output.json file exists. if True then remove it
	_, err = os.Stat("output.json")
	//fmt.Println("\n\nValue of fout is; ", fout)
	if os.IsNotExist(err) {
		fmt.Println("\n ### File Not Found. Creating output.json file ###")
		// writing output as json file
		err = ioutil.WriteFile("output.json", j, 0777)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("\n### File found. Removing the file then recreating it ###")
		// remove the output.json file then recreate it
		err = os.Remove("output.json")
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile("output.json", j, 0777)
		if err != nil {
			panic(err)
		}
	}

	// parse the template
	tpl, err := template.ParseFiles("output.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n Value of tpl is: ", tpl)

	param := Params{
		ORDERER_PROFILE: getEnv("ORDERER_PROFILE"),
		ORD_CHANNEL_ID:  getEnv("ORD_CHANNEL_ID"),
		CHANNEL_PROFILE: getEnv("CHANNEL_PROFILE"),
		CHANNEL_NAME:    getEnv("CHANNEL_NAME"),
		ORG1_NAME:       getEnv("ORG1_NAME"),
		ORG2_NAME:       getEnv("ORG2_NAME"),
		DOMAIN:          getEnv("DOMAIN"),
		PROJECT_NAME:    getEnv("PROJECT_NAME"),
	}
	// checking readInput func
	//test := getEnv("Hello")

	fmt.Println("\nValue of param struct is: ", param)
	//fmt.Println("\n value of test is: ", test)

	// execute the template with the given data
	var ts bytes.Buffer
	err = tpl.Execute(&ts, param) // Execute will fill the buffer so pass as reference
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nJSON:\n%v\n", ts.String())

	// convert to yaml
	y, err := yaml.JSONToYAML([]byte(ts.String()))
	if err != nil {
		panic(err)
	}

	fmt.Println("\n JSON TO YAML \n\n", string(y))

	err = ioutil.WriteFile("docker-compose-e2e-temp.yaml", y, 0777)
	if err != nil {
		panic(err)
	}
}
