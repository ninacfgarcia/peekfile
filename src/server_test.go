package main

import (
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

type MockStubFileMode struct {
	isDir     bool
	isRegular bool
}

func (s MockStubFileMode) IsDir() bool {
	return s.isDir
}
func (s MockStubFileMode) IsRegular() bool {
	return s.isRegular
}

func TestPathDataHandler(t *testing.T) {
	fileMock := MockFunction{calls: make([]string, 0), rval: ResponseData{123, "fake-payload"}}
	dirMock := MockFunction{}
	fo := MockStubFileMode{isDir: false, isRegular: true} // file
	response := pathDataHandler(dirMock.dataFrom, fileMock.dataFrom)("test-string", fo)

	assert.Equal(t, 1, len(fileMock.calls))
	assert.Equal(t, fileMock.calls[0], "test-string")
	assert.EqualValues(t, response, fileMock.rval)
}
