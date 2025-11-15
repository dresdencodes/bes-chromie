package files

import (
	"os"
	"strings"
	"io/ioutil"
	"encoding/json"
)

type JSONContainer struct{}

func JSON() *JSONContainer { return &JSONContainer{}}

func (j *JSONContainer) OpenOrMake(file string, i interface{}) error {
	
	// open
	err := j.Open(file, i)

	// missing
	if err!=nil && FileMissingErr(err) { 

		// make dir 
		err = MakeDirectory(file)
		if err!= nil {return err}

		// write 
		err = j.Write(file, i)
		if err != nil {return err}

	}

	return err

}

func (j *JSONContainer) Open(file string, i interface{}) error {
	
	// read json file
	jsonFile, err := os.Open(file)
	if err!=nil {return err}

	// io util 
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err!=nil {return err}

	// json unmarshal
	err = json.Unmarshal(byteValue, i)
	if err!=nil {return err}
	
	return nil

}

func (j *JSONContainer) Write(file string, i interface{}) error {
	
	if !strings.HasSuffix(file, ".json") {
		file = file + ".json"
	}
    jsonBytes, err := json.Marshal(i)
	if err!=nil {
		return err
	}
    err = ioutil.WriteFile(file, jsonBytes, 0644)
	if err!=nil {
		return err
	}
	
	return nil
}
