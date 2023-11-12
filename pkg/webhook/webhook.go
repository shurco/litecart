package webhook

import (
	"bytes"
	"net/http"
	"fmt"
	
)

func SendHook(url string, payload []byte) (*http.Response, error) {
	
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {	
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}

	return resp, nil
}
