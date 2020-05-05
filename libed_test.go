package libed

import (
	"encoding/json"
	"testing"
)

func TestExampleNewGCMEncrypter(t *testing.T) {
	type DataModel struct {
		Message string `json:"message"`
	}
	req := DataModel{
		Message: "Hi",
	}
	data, err := json.Marshal(req)
	if err != nil {
		t.Log(err)
	}
	enc, err := GCMEncrypter(data, "key_secret")
	if err != nil {
		t.Log(err)
	}
	t.Log(enc) // q8rBym7n5C8DRenLoRbuM9GXCbOlwvfoIwnrTFJmbmUYbX+8RzLA2uNqH8k=

	var res DataModel
	err = GCMDecrypter(enc, "key_secret", &res)
	if err != nil {
		t.Log(err)
	}
	t.Log(res.Message) // Hi
}
