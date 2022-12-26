package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/sirupsen/logrus"

	"github.com/choirulanwar/kitsune-cli/assets"
)

func createNewNestJs(values appValues) {
	rootFs := &assets.Assets

	rootFsMapping := map[string]string{
		"domain.go.tmpl":       "internal/core/domain/domain.ts",
		"ports.go.tmpl":        "internal/core/ports/ports.ts",
		"service.go.tmpl":      "internal/core/service/service.ts",
		"input.go.tmpl":        "internal/dto/",
		"handlers.gql.go.tmpl": "internal/handlers/gql.ts",
		"module.go.tmpl":       fmt.Sprintf("%s.module.ts", values.AppName),
	}

	subdirs = []string{
		"internal/core/domain",
		"internal/core/ports",
		"internal/core/service",
		"internal/dto",
		"internal/handlers",
	}

	for _, dirname := range subdirs {
		if err = os.MkdirAll(dirname, 0755); err != nil {
			logrus.WithError(err).Fatalf("[] Unable to create subdirectory %s", dirname)
		}
	}

	var schemaData SchemaStruct
	data, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", values.Cwd, values.Schema))
	if err != nil {
		fmt.Println(err)
		return
	}
	schemaData.readSchema(data)

	if templates, err = template.New("").Funcs(template.FuncMap{
		"ToUpper":    strings.ToUpper,
		"ToTitle":    strings.Title,
		"ToLower":    strings.ToLower,
		"Contains":   strings.Contains,
		"HasPrefix":  strings.HasPrefix,
		"HasSuffix":  strings.HasSuffix,
		"Replace":    strings.Replace,
		"RemoveDash": RemoveDash,
	}).ParseFS(rootFs, "templates/nestjs/*.go.tmpl"); err != nil {
		logrus.WithError(err).Fatal("[] Error parsing root templates files")
	}

	for templateName, outputPath := range rootFsMapping {
		if templateName == "input.go.tmpl" {
			for _, val := range schemaData.Input {
				if fp, err = os.Create(fmt.Sprintf("%s%s.input.ts", outputPath, val.Operation)); err != nil {
					logrus.WithError(err).Fatalf("[] Unable to create file %s for writing", outputPath)
				}

				defer fp.Close()

				data := map[string]interface{}{
					"ModuleName": schemaData.ModuleName,
					"Input":      val,
				}

				if err = templates.ExecuteTemplate(fp, templateName, data); err != nil {
					logrus.WithError(err).Fatalf("[] Unable to execute template %s", templateName)
				}
			}
		} else {
			if fp, err = os.Create(outputPath); err != nil {
				logrus.WithError(err).Fatalf("[] Unable to create file %s for writing", outputPath)
			}

			defer fp.Close()

			if err = templates.ExecuteTemplate(fp, templateName, schemaData); err != nil {
				logrus.WithError(err).Fatalf("[] Unable to execute template %s", templateName)
			}
		}
	}
}

func prettifyTS(values appValues) {
	err = exec.Command("npx", "prettier", "--write", fmt.Sprintf("%s/%s", values.Cwd, values.AppName)).Run()
	if err != nil {
		logrus.WithError(err).Errorf("[] Error prettify output '%s'", fmt.Sprintf("%s/%s", values.Cwd, values.AppName))
	}
}
