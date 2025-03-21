package ahhhhh

import "encoding/json"

// parsing json response
func parseJSON(jsonStr string) (StudentData, error) {
	var data StudentData
	err := json.Unmarshal([]byte(jsonStr), &data)
	return data, err
}

func parseNames(jsonStr string) ([]Student, error) {
	var data []Student
	err := json.Unmarshal([]byte(jsonStr), &data)
	return data, err
}
