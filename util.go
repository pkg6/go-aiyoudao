package aiyoudao

import (
	"encoding/base64"
	"io"
	"os"
)

func SaveFile(path string, data []byte, needDecode bool) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if needDecode {
		data, _ = base64.StdEncoding.DecodeString(string(data))
	}
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func ReadFileAsBase64(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	fd, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(fd), nil
}
