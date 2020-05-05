# LIBED
> libed is a golang lib to encrypt and decrypt data.

# How to use

## Data example
```go
  type DataModel struct {
		Message string `json:"message"`
	}
	req := DataModel{
		Message: "Hi",
	}
  
	data, err := json.Marshal(req)
	if err != nil {
    return err
  }
```

## Encrypt
```go
	enc, err := GCMEncrypter(data, "key_secret")
	if err != nil {
    return err
  }
	fmt.Println(enc) // q8rBym7n5C8DRenLoRbuM9GXCbOlwvfoIwnrTFJmbmUYbX+8RzLA2uNqH8k=
```

## Decrypt
```go
  var res DataModel
  err = GCMDecrypter(enc, "key_secret", &res)
  if err != nil {
    return err
  }
  fmt.Println(res.Message) // Hi
```