package cstools

import (
	"io/ioutil"
)

func ReadFile(fileName string) (string, error) {
	b, err := ioutil.ReadFile(fileName) // just pass the file name
	if err != nil {
		return "", err
	}

	return string(b), nil
}
