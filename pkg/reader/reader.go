package reader

import (
	"io/ioutil"

	"github.com/dotpy3/compose-preview/pkg/types"
	yaml "gopkg.in/yaml.v2"
)

func readBytes(v []byte) (types.Config, error) {
	res := types.Config{}
	err := yaml.Unmarshal(v, &res)
	return res, err
}

// ReadString returns the docker-compose config read from a string.
func ReadString(s string) (types.Config, error) {
	return readBytes([]byte(s))
}

// ReadFile returns the docker-compose config read from a file.
func ReadFile(path string) (types.Config, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return types.Config{}, err
	}

	return readBytes(content)
}
