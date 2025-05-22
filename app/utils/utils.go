package utils

import (
	"bytes"
	"compress/gzip"
	"errors"
	"io"
	"log/slog"
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
			slog.Error("error reading request content: %v", "error", err)
			return nil, err
		}
	}
	content := strings.Split(string(buffer), "\r\n")
	return content, nil
}

func ReadFile(sourcePath string) ([]byte, error) {
	byt, err := os.ReadFile(sourcePath)
	if err != nil {
		slog.Error("error reading file located in %s. %s", sourcePath, err)
		return nil, err
	}

	return byt, nil
}

func WriteFile(sourcePath string, data string) error {
	dataBytes := bytes.Trim([]byte(data), "\x00")
	err := os.WriteFile(sourcePath, dataBytes, 0644)
	if err != nil {
		slog.Error("error writing data into the file %s. %v", sourcePath, err)
		return err
	}
	return nil
}

func CompressContent(content string) ([]byte, error) {
	var buf bytes.Buffer
	writer := gzip.NewWriter(&buf)

	_, err := writer.Write([]byte(content))
	if err != nil {
		slog.Error("error compressing content: %v", "error", err)
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		slog.Error("error closing writer compression: %v", "error", err)
		return nil, err
	}

	return buf.Bytes(), nil
}
