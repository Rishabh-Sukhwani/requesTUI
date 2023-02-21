package services

import (
	"net/http"
	"github.com/rivo/tview"
	"fmt"
	"bytes"
	"encoding/json"
)

func GetRequest(url string, headers string, textView *tview.TextView) {
	
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		textView.SetText(fmt.Sprintf("%v", err))
		return
	} 

	if headers != "" {
		var headerMap map[string]string
		err = json.Unmarshal([]byte(headers), &headerMap)
		if err != nil {
			textView.SetText(fmt.Sprintf("%v", err))
			return
		}

		for key, value := range headerMap {
			req.Header.Set(key, value)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		textView.SetText(fmt.Sprintf("%v", err))
		return
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		textView.SetText(fmt.Sprintf("%v", err))
		return
	}

	var jsonData interface{}
	err = json.Unmarshal(buf.Bytes(), &jsonData)
	if err != nil {
		textView.SetText(fmt.Sprintf("%v", err))
		return
	}

	prettyJson, err := json.MarshalIndent(jsonData, "", "	")
	if err != nil {
		textView.SetText(fmt.Sprintf("%v", err))
		return
	}



	textView.SetText(string(prettyJson))
	

}