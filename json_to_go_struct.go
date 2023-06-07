package main

import (
	"encoding/json"
	"fmt"
	"github.com/atotto/clipboard"
	"sort"
	"strings"
)

func JsonToStruct(data string) {
	trimData := strings.Trim(data, "")
	err, result := jsonToGolan(trimData)
	if err != nil {
		return
	}

	s := createTypeStruct(result)

	err = clipboard.WriteAll(s)
	if err != nil {
		fmt.Println("Error copying to clipboard:", err)
		return
	}

	fmt.Println("###################################################################")
	fmt.Println("###################################################################")
	fmt.Println("###################################################################")
	fmt.Println("###################################################################")
	fmt.Println()
	fmt.Println()
	fmt.Printf(s)
	fmt.Println()
	fmt.Println()
	fmt.Println("Struct copied to clipboard")
	fmt.Println("###################################################################")
	fmt.Println("###################################################################")
	fmt.Println("###################################################################")
	fmt.Println("###################################################################")
}

func jsonToGolan(d string) (error, interface{}) {
	var result interface{}
	var err error
	if string(d[0]) == "{" {
		err, result = jsonToMap(d)
	} else {
		err, result = jsonToArrayMap(d)
	}
	if err != nil {
		err, result = jsonToAnything(d)
		if err != nil {
			return err, nil
		}
	}

	return nil, result
}

func toCapital(s string) string {
	return strings.ToUpper(string(s[0])) + s[1:]
}

func createTypeStruct(r interface{}) string {
	var toReturn strings.Builder
	toReturn.WriteString("type AutogeneratedStruct ")

	switch value := r.(type) {
	case []map[string]interface{}:
		if len(value) > 0 {
			toReturn.WriteString("[]" + createStruct(value[0]))
		} else {
			toReturn.WriteString("[]any")
		}
	case map[string]interface{}:
		toReturn.WriteString(createStruct(value))
	case []interface{}:
		if len(value) > 0 {
			toReturn.WriteString("[]" + typeToString(value[0]))
		} else {
			toReturn.WriteString("[]any")
		}
	}
	return toReturn.String()
}

func createStruct(r map[string]interface{}) string {
	var toReturn strings.Builder
	toReturn.WriteString("struct {\n")
	// Extract the keys into a slice
	keys := make([]string, 0, len(r))
	for key := range r {
		keys = append(keys, key)
	}
	// Sort the keys
	sort.Strings(keys)
	// Iterate over the sorted keys and access values from the map
	for _, key := range keys {
		toReturn.WriteString("  " + toCapital(key) + " " + typeConversion(key, r[key]) + "\n")
	}
	toReturn.WriteString("}")
	return toReturn.String()
}

func typeConversion(k string, v interface{}) string {
	var toReturn string
	toReturn += typeToString(v) + " " + fmt.Sprintf("`json:\"%s\"`", strings.ToLower(k))
	return toReturn
}

func typeToString(v interface{}) string {
	var toReturn string
	switch value := v.(type) {
	case int:
		toReturn += "int"
	case string:
		toReturn += "string"
	case float64:
		toReturn += "float64"
	case bool:
		toReturn += "bool"
	case []interface{}:
		if len(value) > 0 {
			toReturn += "[]" + typeToString(value[0])
		} else {
			toReturn += "[]any"
		}
	case map[string]interface{}:
		toReturn += createStruct(v.(map[string]interface{}))
	default:
		toReturn += fmt.Sprintf("%v", value)
	}
	return toReturn
}

func jsonToMap(data string) (error, map[string]interface{}) {
	var response map[string]interface{}
	err := json.Unmarshal([]byte(data), &response)
	if err != nil {
		fmt.Println("Error:", err)
		return err, nil
	}
	return nil, response
}

func jsonToAnything(data string) (error, interface{}) {
	var response []interface{}
	err := json.Unmarshal([]byte(data), &response)
	if err != nil {
		fmt.Println("Error:", err)
		return err, nil
	}
	return nil, response
}

func jsonToArrayMap(data string) (error, []map[string]interface{}) {
	var response []map[string]interface{}
	err := json.Unmarshal([]byte(data), &response)
	if err != nil {
		fmt.Println("Error:", err)
		return err, nil
	}
	return nil, response
}
