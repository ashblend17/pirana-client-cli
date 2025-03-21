package ahhhhh

import (
	"encoding/json"
	"fmt"
	"strings"
)

// request handler function + formatting
// need to fix password thing
func GetData(rolls []string, flag string, password string) {

	jsonData, err := json.Marshal(rolls)
	if err != nil {
		fmt.Print("Error converting roll to json: ", err)
	}

	payload := []byte(fmt.Sprintf(`{"Roll": %s, "AlfredPassword": "%s", "Flags" : "%s"}`, jsonData, password, flag))
	data := reqData(payload)
	parsedJson, err := parseJSON(string(data))
	if err != nil {
		fmt.Println("Error parsing json: ", err)
		fmt.Println()
	}
	// fmt.Println(parsedJson.Password)
	if strings.Contains(flag, "s") && strings.Contains(flag, "c") {
		PrintGroupedTableByRoll(parsedJson, rolls)
	} else if strings.Contains(flag, "c") {
		PrintGroupedCourseTable(parsedJson, rolls)
	} else if strings.Contains(flag, "s") {
		PrintGroupedSGPATable(parsedJson, rolls)
	}

	if strings.Contains(flag, "p") {
		printPasswordTable(parsedJson, rolls)
	}

}

// get names when running find command
func GetNames(password string, name string) {
	payload := []byte(fmt.Sprintf(`{"alfredPassword": "%s", "name": "%s"}`, password, name))
	data := reqNames(payload)
	parsedJson, err := parseNames(string(data))
	if err != nil {
		fmt.Println("Error parsing json: ", err)
		fmt.Println()
	}
	PrintStudentTable(parsedJson)
}
