package cmd

import "gopkg.in/yaml.v3"

type SchemaStruct struct {
	ModuleName string `yaml:"module_name"`
	Schema     struct {
		Properties []struct {
			Name     string `yaml:"name"`
			Type     string `yaml:"type"`
			Required bool   `yaml:"required"`
		} `yaml:"properties"`
	} `yaml:"schema"`
	Input []struct {
		Operation string `yaml:"operation"`
		Label     string `yaml:"label"`
		Fields    []struct {
			Name     string `yaml:"name"`
			Type     string `yaml:"type"`
			Required bool   `yaml:"required"`
		} `yaml:"fields"`
	} `yaml:"input"`
}

func (d *SchemaStruct) readSchema(data []byte) *SchemaStruct {
	yaml.Unmarshal(data, d)
	return d
}
