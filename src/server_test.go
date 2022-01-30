package main

import (
	"errors"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockFunction struct {
	calls []string
	rval  ResponseData
}

func (mock *MockFunction) dataFrom(path string) ResponseData {
	mock.calls = append(mock.calls, path)
	return mock.rval
}

type MockFileMode struct {
	isDir     bool
	isRegular bool
}

func (s MockFileMode) IsDir() bool {
	return s.isDir
}
func (s MockFileMode) IsRegular() bool {
	return s.isRegular
}

func TestPathDataHandler(t *testing.T) {
	fileMock := MockFunction{rval: ResponseData{123, "fake-payload-file"}}
	dirMock := MockFunction{rval: ResponseData{123, "fake-payload-dir"}}
	handler := pathDataHandler(dirMock.dataFrom, fileMock.dataFrom)

	fileFileMode := MockFileMode{isDir: false, isRegular: true}
	response := handler("file-mock-path/", fileFileMode)

	assert.Equal(t, 1, len(fileMock.calls))
	assert.Equal(t, fileMock.calls[0], "file-mock-path/")
	assert.EqualValues(t, response, fileMock.rval)

	dirFileMode := MockFileMode{isDir: true, isRegular: false}
	response = handler("dir-mock-path/", dirFileMode)
	assert.Equal(t, 1, len(fileMock.calls))
	assert.Equal(t, dirMock.calls[0], "dir-mock-path/")
	assert.EqualValues(t, response, dirMock.rval)
}

func TestGetResponseForPathReturnsError(t *testing.T) {
	f := GetResponseForPath(func(string) (StubFileMode, error) {
		return nil, errors.New("")
	})
	assert.Equal(t, http.StatusNotFound, f("bad-path").Status,
		"GetResponseForPath handler should respond with StatusNotFound",
	)
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
