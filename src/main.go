package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func fileDataFrom(path string) ResponseData {
	data, readErr := os.ReadFile(path)
	if readErr != nil {
		return ResponseData{
			http.StatusInternalServerError,
			BadPayload{readErr.Error()},
		}
	}
	return ResponseData{
		http.StatusOK,
		FilePayload{string(data)},
	}
}

func dirDataFrom(path string) ResponseData {
	data, readErr := os.ReadDir(path)
	if readErr != nil {
		return ResponseData{
			http.StatusForbidden,
			BadPayload{readErr.Error()},
		}
	}
	payload, formatErr := FormatEntries(data)
	if formatErr != nil {
		return ResponseData{
			http.StatusInternalServerError,
			BadPayload{formatErr.Error()},
		}
	}
	return ResponseData{http.StatusOK, DirPayload{payload}}
}

type StubFileMode interface {
	IsRegular() bool
	IsDir() bool
}

func pathDataHandler(
	dirDataFrom func(string) ResponseData,
	fileDataFrom func(string) ResponseData,
) func(string, StubFileMode) ResponseData {
	return func(path string, info StubFileMode) ResponseData {
		if info != nil {
			if info.IsDir() {
				return dirDataFrom(path)
			} else if info.IsRegular() {
				return fileDataFrom(path)
			}
		}
		return ResponseData{
			http.StatusNotFound,
			BadPayload{fmt.Sprintf("Not a valid path %s", path)},
		}
	}
}

func GetLStatMode(path string) (StubFileMode, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return nil, err
	}
	return info.Mode(), nil
}

func GetResponseForPath(getMode func(string) (StubFileMode, error)) func(string) ResponseData {
	return func(path string) ResponseData {
		mode, _ := getMode(path)

		handler := pathDataHandler(dirDataFrom, fileDataFrom)
		return handler(path, mode)
	}
}

func marshalResponseData(data ResponseData) ([]byte, int) {
	body, err := json.MarshalIndent(data.Payload, "", "  ")
	if err != nil {
		return []byte("{\"error\":\"Could not marshal response\"}"),
			http.StatusInternalServerError
	}
	return body, data.Status
}

func pathHandler(rw http.ResponseWriter, req *http.Request) {
	path := filepath.Join(os.Args[1], req.URL.Path)
	defer req.Body.Close()

	dataHandler := GetResponseForPath(GetLStatMode)
	data := dataHandler(path)
	rw.Header().Set("Content-Type", "application/json")
	responseBody, status := marshalResponseData(data)
	rw.WriteHeader(status)
	rw.Write(responseBody)
}

func main() {
	http.HandleFunc("/", pathHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
