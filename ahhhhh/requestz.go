package ahhhhh

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

const DATA_URL string = "http://localhost:8080/getData"
const NAME_URL string = "http://localhost:8080/getNames"

// actual post request for data
func reqData(payload []byte) []byte {
	full_url := DATA_URL
	fmt.Println("Fetching Data...")

	req, err := http.NewRequest("POST", full_url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil
	}
	res.Body.Close()

	return body
}

// req for getNames
func reqNames(payload []byte) []byte {
	req, err := http.NewRequest("POST", NAME_URL, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		fmt.Println("Error with sending req", err)
		return nil
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil
	}

	res.Body.Close()
	return body
}
