package reader_test

import (
	"strings"
	"testing"

	"github.com/dotpy3/compose-preview/pkg/reader"
	"github.com/dotpy3/compose-preview/pkg/types"
	"github.com/go-test/deep"
)

var addressDb27017 = "db:27017"

var testCases = []struct {
	compose string
	config  types.Config
}{
	{
		compose: ``,
		config:  types.Config{},
	},
	{
		compose: `version: '3'
services:
  db:
    image: mongo:3.4
  web:
    extends:
      file: docker-compose.versions.yml
      service: web
    depends_on:
    - db
    environment:
      APP_MONGO_ADDRESS: db:27017`,
		config: types.Config{
			Version: "3",
			Services: types.Services{
				{
					Name:  "db",
					Image: "mongo:3.4",
				},
				{
					Name: "web",
					Extends: &types.ExtensionConfig{
						Filename: "docker-compose.versions.yml",
						Service:  "web",
					},
					DependsOn: []string{"db"},
					Environment: types.MappingWithEquals{
						"APP_MONGO_ADDRESS": &addressDb27017,
					},
				},
			},
		},
	},
}

func TestReader(t *testing.T) {
	for i, testCase := range testCases {
		res, err := reader.ReadString(testCase.compose)
		if err != nil {
			t.Fatalf("Could not read test case %d: %s", i, err)
		}

		diffs := deep.Equal(res, testCase.config)
		if len(diffs) > 0 {
			t.Fatalf(
				"Differences found between expected and actual result: \n%s",
				strings.Join(diffs, "\n"),
			)
		}
	}
}

func TestWriter(t *testing.T) {

}
