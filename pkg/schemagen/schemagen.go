package schemagen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jenkins-x/jx-api/v4/pkg/util"
	"github.com/jenkins-x/jx-logging/v3/pkg/log"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	// DefaultDirWritePermissions default permissions when creating a directory
	DefaultDirWritePermissions = 0766
)

// ResourceKind defines a resource kind to have its schema generated
type ResourceKind struct {
	APIVersion string
	Name       string
	Resource   interface{}
}

// GenerateSchemas generates the schemas for the given kinds
func GenerateSchemas(resourceKinds []ResourceKind, out string) error {
	for _, k := range resourceKinds {
		name := k.Name
		out := filepath.Join(out, k.APIVersion, name+".json")
		dir := filepath.Dir(out)
		err := os.MkdirAll(dir, DefaultDirWritePermissions)
		if err != nil {
			return errors.Wrapf(err, "failed to create dir %s", dir)
		}

		err = generate(name, out, k.Resource)
		if err != nil {
			return errors.Wrapf(err, "failed to generate file %s", out)
		}
	}
	return nil
}

// Generate generates the schema document
func generate(schemaName string, out string, schemaTarget interface{}) error {
	schema := util.GenerateSchema(schemaTarget)
	if schema == nil {
		return fmt.Errorf("could not generate schema for %s", schemaName)
	}

	output := prettyPrintJSON(schema)

	if output == "" {
		tempOutput, err := json.Marshal(schema)
		if err != nil {
			return errors.Wrapf(err, "error outputting schema for %s", schemaName)
		}
		output = string(tempOutput)
	}
	log.Logger().Infof("JSON schema for %s:", schemaName)

	if out != "" {
		err := ioutil.WriteFile(out, []byte(output), util.DefaultWritePermissions)
		if err != nil {
			return errors.Wrapf(err, "failed to save file %s", out)
		}
		log.Logger().Infof("wrote file %s", out)
		return nil
	}
	log.Logger().Infof("%s", output)
	return nil
}

func prettyPrintJSON(input interface{}) string {
	output := &bytes.Buffer{}
	if err := json.NewEncoder(output).Encode(input); err != nil {
		return ""
	}
	formatted := &bytes.Buffer{}
	if err := json.Indent(formatted, output.Bytes(), "", "  "); err != nil {
		return ""
	}
	return string(formatted.Bytes())
}
