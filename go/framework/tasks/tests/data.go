package tests

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type DataType string

const (
	Mock DataType = "mock"
	Real          = "real"
)

type DataDescriptor struct {
	dataType DataType
	path     string
}

func (d DataDescriptor) Path() (string, error) {
	switch d.dataType {
	case Mock:
		return filepath.Abs(filepath.Join("testdata", d.path))
	case Real:
		return filepath.Abs(filepath.Join("../../data", d.path))
	default:
		return "", errors.New(fmt.Sprintf("unknown data type: %v", d.dataType))
	}
}

func MockData(name string) DataDescriptor {
	return DataDescriptor{Mock, name}
}

func RealData(name string) DataDescriptor {
	return DataDescriptor{Real, name}
}

func ReadData(data DataDescriptor) (string, error) {
	var bytes []byte
	if path, err := data.Path(); err != nil {
		return "", err
	} else if bytes, err = os.ReadFile(path); err != nil {
		return "", err
	}

	return strings.TrimRight(string(bytes), " \n"), nil
}
