package util

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	indexLibrary "github.com/devfile/registry-support/index/generator/library"
	indexSchema "github.com/devfile/registry-support/index/generator/schema"
)

// IsHtmlRequested checks the accept header if html has been requested
func IsHtmlRequested(acceptHeader []string) bool {
	for _, header := range acceptHeader {
		if strings.Contains(header, "text/html") {
			return true
		}
	}
	return false
}

// EncodeIndexIconToBase64 encodes all index icons to base64 format given the index file path
func EncodeIndexIconToBase64(indexPath string, base64IndexPath string) ([]byte, error) {
	// load index
	bytes, err := ioutil.ReadFile(indexPath)
	if err != nil {
		return nil, err
	}
	var index []indexSchema.Schema
	err = json.Unmarshal(bytes, &index)
	if err != nil {
		return nil, err
	}

	// encode all index icons to base64 format
	for i, indexEntry := range index {
		if indexEntry.Icon != "" {
			base64Icon, err := encodeToBase64(indexEntry.Icon)
			index[i].Icon = base64Icon
			if err != nil {
				return nil, err
			}
		}
	}
	err = indexLibrary.CreateIndexFile(index, base64IndexPath)
	if err != nil {
		return nil, err
	}
	bytes, err = json.MarshalIndent(&index, "", "  ")
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// encodeToBase64 encodes the content from the given uri to base64 format
func encodeToBase64(uri string) (string, error) {
	url, err := url.Parse(uri)
	if err != nil {
		return "", err
	}

	// load the content from the given uri
	var bytes []byte
	if url.Scheme == "http" || url.Scheme == "https" {
		resp, err := http.Get(uri)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()

		bytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
	} else {
		bytes, err = ioutil.ReadFile(uri)
		if err != nil {
			return "", err
		}
	}

	// encode the content to base64 format
	var base64Encoding string
	mimeType := http.DetectContentType(bytes)
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	default:
		base64Encoding += "data:image/svg+xml;base64,"
	}
	base64Encoding += base64.StdEncoding.EncodeToString(bytes)
	return base64Encoding, nil
}

// ReadIndexPath reads the index from the path and unmarshalls it into the index
func ReadIndexPath(indexPath string) ([]indexSchema.Schema, error) {
	// load index
	bytes, err := ioutil.ReadFile(indexPath)
	if err != nil {
		return nil, err
	}
	var index []indexSchema.Schema
	err = json.Unmarshal(bytes, &index)
	if err != nil {
		return nil, err
	}

	return index, nil
}
