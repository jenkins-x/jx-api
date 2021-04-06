package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jenkins-x/jx-api/v4/pkg/apis/core/v4beta1"
	"github.com/jenkins-x/jx-api/v4/pkg/util"
	"github.com/jenkins-x/jx-logging/v3/pkg/log"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
)

const ()

func main() {
	o := &Options{}
	if len(os.Args) > 1 {
		o.Out = os.Args[1]
	}
	err := o.Run()
	if err != nil {
		log.Logger().Errorf("failed: %v", err)
		os.Exit(1)
	}
	log.Logger().Infof("completed the plugin generator")
	os.Exit(0)
}

type Options struct {
	Out string
}

// Run implements this command
func (o *Options) Run() error {
	if o.Out == "" {
		o.Out = "schema/jx-requirements.json"
	}
	return o.Generate("jx-requirements.yml", &v4beta1.Requirements{})
}

// Generate generates the schema document
func (o *Options) Generate(schemaName string, schemaTarget interface{}) error {
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

	if o.Out != "" {
		err := ioutil.WriteFile(o.Out, []byte(output), util.DefaultWritePermissions)
		if err != nil {
			return errors.Wrapf(err, "failed to save file %s", o.Out)
		}
		log.Logger().Infof("wrote file %s", o.Out)
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
