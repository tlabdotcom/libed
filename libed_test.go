package libed

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExampleNewGCMEncrypter(t *testing.T) {
	type DataModel struct {
		Message string `json:"message"`
	}
	req := DataModel{
		Message: "Hi",
	}
	data, err := json.Marshal(req)
	require.NoError(t, err)
	enc, err := GCMEncrypter(data, "key_secret")
	require.NoError(t, err)
	t.Log(enc) // q8rBym7n5C8DRenLoRbuM9GXCbOlwvfoIwnrTFJmbmUYbX+8RzLA2uNqH8k=

	var res DataModel
	err = GCMDecrypter(enc, "key_secret", &res)
	require.NoError(t, err)
	t.Log(res.Message) // Hi
}
