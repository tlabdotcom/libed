package libed

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExampleNewGCMEncrypter(t *testing.T) {
	type DataModel struct {
		Message string
	}
	req := DataModel{
		Message: "Hi",
	}
	data, err := json.Marshal(req)
	require.NoError(t, err)
	enc, err := GCMEncrypter(data, "key_secret")
	require.NoError(t, err)
	t.Log(enc) // q8rBym7n5C8DRenLoRbuM9GXCbOlwvfoIwnrTFJmbmUYbX+8RzLA2uNqH8k=

	denc, err := GCMDecrypter(enc, "key_secret")
	require.NoError(t, err)
	t.Log(denc) // {"Message":"Hi"}
}
