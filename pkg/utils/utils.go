package utils

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strings"
)

func LoadJson(filePath string) []byte {
	resource := LoadCommonFile(filePath)
	minifyJsonB, err := minify(resource)
	if err != nil {
		log.Printf("[Testutils] Cannot minify json resource %s", err)
	}
	return minifyJsonB
}

func LoadCommonFile(filePath string) []byte {
	resourceFolderPath := getResourceFolderPath()
	fullPath := fmt.Sprintf("%s%s", resourceFolderPath, filePath)
	resource, err := ioutil.ReadFile(fullPath)
	if err != nil {
		log.Infof("[Testutils] Cannot load the resource %s", err)
	}
	return decode(resource)
}

func getResourceFolderPath() string {
	workingDir, _ := os.Getwd()
	var dirSplit []string
	dirSplit = strings.Split(workingDir, "car-smile-mngr-go")
	if len(dirSplit) == 2 {
		return dirSplit[0] + "car-smile-mngr-go/pkg/test/resources/"
	} else {
		dirSplit = strings.Split(workingDir, "/app")
		return "/app/pkg/test/resources/"
	}
}

func minify(jsonB []byte) ([]byte, error) {

	var buff *bytes.Buffer = new(bytes.Buffer)
	errCompact := json.Compact(buff, jsonB)
	if errCompact != nil {
		newErr := fmt.Errorf("failure encountered compacting json := %v", errCompact)
		return []byte{}, newErr
	}

	b, err := ioutil.ReadAll(buff)
	if err != nil {
		readErr := fmt.Errorf("read buffer error encountered := %v", err)
		return []byte{}, readErr
	}

	return b, nil
}

func decode(dataCode []byte) []byte {
	value, err := b64.StdEncoding.DecodeString(string(dataCode))

	if err != nil {
		log.Error(err)
	}
	return value
}

func DecodeString(dataCode string) string {
	value, err := b64.StdEncoding.DecodeString(dataCode)

	if err != nil {
		log.Error(err)
	}
	return string(value)
}
