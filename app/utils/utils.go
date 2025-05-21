package utils

import (
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

type Handler struct{}

func ReadRequestContent(conn net.Conn) ([]string, error) {
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		if !errors.Is(err, io.EOF) {
			fmt.Printf("error reading request content: %v\n", err)
			return nil, err
		}
	}
	content := strings.Split(string(buffer), "\r\n")
	return content, nil
}

func ReadFile(sourcePath string) ([]byte, error) {
	byt, err := os.ReadFile(sourcePath)
	if err != nil {
		fmt.Printf("error reading file located in %s. %v\n", sourcePath, err)
		return nil, err
	}

	return byt, nil
}

func WriteFile(sourcePath string, data string) error {
	dataBytes := bytes.Trim([]byte(data), "\x00")
	err := os.WriteFile(sourcePath, dataBytes, 0666)
	if err != nil {
		fmt.Printf("error writing data into the file %s. %v\n", sourcePath, err)
		return err
	}
	return nil
}

func CompressContent(content string) ([]byte, error) {
	var buf bytes.Buffer
	writer := gzip.NewWriter(&buf)

	_, err := writer.Write([]byte(content))
	if err != nil {
		fmt.Printf("error compressing content: %v\n", err)
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		fmt.Printf("error closing writer compression: %v\n", err)
		return nil, err
	}

	return buf.Bytes(), nil
}
