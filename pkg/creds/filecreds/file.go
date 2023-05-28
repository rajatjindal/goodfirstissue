package filecreds

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type file struct {
	filename string
}

func Provider(filename string) *file {
	return &file{
		filename: filename,
	}
}

func (f *file) GetCredentials(svc string) (map[string]string, error) {
	r, err := os.ReadFile(f.filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read credentials file with err: %s", err.Error())
	}

	t := map[string]map[string]string{}
	err = yaml.Unmarshal(r, &t)
	if err != nil {
		return nil, err
	}

	return t[svc], nil
}
