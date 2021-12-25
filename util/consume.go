package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ConsumeAPIPost(targetConsume string, bodyType string, bodyTemplate map[string]string) (body []byte, err error) {
	bodyJson, err := json.Marshal(bodyTemplate)
	if err != nil {
		return
	}
	resp, err := http.Post(targetConsume, bodyType, bytes.NewBuffer(bodyJson))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, _ = ioutil.ReadAll(resp.Body)
	return
}
