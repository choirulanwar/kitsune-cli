package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/sirupsen/logrus"

	"github.com/choirulanwar/kitsune-cli/assets"
)

func transformYamlToProto(values appValues) {
	rootFs := &assets.Assets

	if _, err := os.Stat(values.AppName); os.IsNotExist(err) {
		if err = os.Mkdir(values.AppName, 0755); err != nil {
			logrus.WithError(err).Errorf("[] Error attempting to create application directory '%s'", values.AppName)
		}
	}

	if err = os.Chdir(values.AppName); err != nil {
		logrus.WithError(err).Errorf("[] Error changing to application directory '%s'", values.AppName)
	}

	var schemaData SchemaStruct
	data, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", values.Cwd, values.Schema))
	if err != nil {
		fmt.Println(err)
		return
	}
	schemaData.readSchema(data)

	rootFsMapping := map[string]string{
		"proto.go.tmpl": fmt.Sprintf("%s.proto", schemaData.ModuleName),
	}

	if templates, err = template.New("").Funcs(template.FuncMap{
		"ToUpper":     strings.ToUpper,
		"ToTitle":     strings.Title,
		"ToLower":     strings.ToLower,
		"Contains":    strings.Contains,
		"HasPrefix":   strings.HasPrefix,
		"HasSuffix":   strings.HasSuffix,
		"Replace":     strings.Replace,
		"RemoveDash":  RemoveDash,
		"RemoveSpace": RemoveSpace,
		"Add":         Add,
	}).ParseFS(rootFs, "templates/domain/*.go.tmpl"); err != nil {
		logrus.WithError(err).Fatal("[] Error parsing root templates files")
	}

	for templateName, outputPath := range rootFsMapping {
		if fp, err = os.Create(outputPath); err != nil {
			logrus.WithError(err).Fatalf("[] Unable to create file %s for writing", outputPath)
		}

		defer fp.Close()

		if err = templates.ExecuteTemplate(fp, templateName, schemaData); err != nil {
			logrus.WithError(err).Fatalf("[] Unable to execute template %s", templateName)
		}

	}
}
