package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/jenkins-x/jx/v2/cmd/codegen/util"

	"github.com/ghodss/yaml"

	"k8s.io/kube-openapi/pkg/builder"

	"k8s.io/kube-openapi/pkg/common"

	"github.com/go-openapi/spec"
	"github.com/pkg/errors"
)

const (
	openapiTemplateSrc = `// +build !ignore_autogenerated

// Code generated by jx create client. DO NOT EDIT.
package openapi

import (
	openapicore "{{ $.Path }}"
	{{ range $i, $path := $.Dependents }}
	openapi{{ $i }} "{{ $path }}"
	{{ end }}
	"k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	result := make(map[string]common.OpenAPIDefinition)
    // This is our core openapi definitions (the ones for this module)
	for k, v := range openapicore.GetOpenAPIDefinitions(ref) {
		result[k] = v
	}
	// These are the ones we depend on
	{{ range $i, $path := $.Dependents }}
	for k, v := range openapi{{ $i}}.GetOpenAPIDefinitions(ref) {
		result[k] = v
	}
	{{ end }}
	return result
}

func GetNames(ref common.ReferenceCallback) []string {
	result := make([]string, 0)
	for k, _ := range openapicore.GetOpenAPIDefinitions(ref) {
		result = append(result, k)
	}
	return result
}
`

	schemaWriterTemplateSrc = `package main

import (
	"flag"
	"os"
	"strings"

	openapi "{{ $.AllImportPath }}"

	"github.com/go-openapi/spec"

	"github.com/pkg/errors"

	"github.com/jenkins-x/jx/v2/cmd/codegen/generator"
)

func main() {
	var outputDir, namesStr, title, version string
	flag.StringVar(&outputDir, "output-directory", "", "directory to write generated files to")
	flag.StringVar(&namesStr, "names", "", "comma separated list of resources to generate schema for, "+
		"if empty all resources will be generated")
	flag.StringVar(&title, "title", "", "title for OpenAPI and HTML generated docs")
	flag.StringVar(&version, "version", "", "version for OpenAPI and HTML generated docs")
	flag.Parse()
	if outputDir == "" {
		panic(errors.New("--output-directory cannot be empty"))
	}
	var names []string
	if namesStr != "" {
		names = strings.Split(namesStr, ",")
	} else {
		refCallback := func(path string) spec.Ref {
			return spec.Ref{}
		}
		names = openapi.GetNames(refCallback)
	}
	err := generator.WriteSchemaToDisk(outputDir, title, version, openapi.GetOpenAPIDefinitions, names)
	if err != nil {
		panic(errors.Wrapf(err, "writing schema to %s", outputDir))
	}
	os.Exit(0)
}
`
	OpenApiDir              = "openapi"
	SchemaWriterSrcFileName = "schema_writer_generated.go"
	OpenApiV2JSON           = "openapiv2.json"
	OpenApiV2YAML           = "openapiv2.yaml"
	openApiGenerator        = "openapi-gen"

	bootstrapJsUrl      = "https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js"
	bootstrapJsFileName = "bootstrap-3.3.7.min.js"
	jqueryUrl           = "https://code.jquery.com/jquery-3.2.1.min.js"
	jqueryFileName      = "jquery-3.2.1.min.js"

	openApiGen = "k8s.io/kube-openapi/cmd/openapi-gen"
)

var (
	fonts = []string{
		"FontAwesome.otf",
		"fontawesome-webfont.eot",
		"fontawesome-webfont.svg",
		"fontawesome-webfont.ttf",
		"fontawesome-webfont.woff",
		"fontawesome-webfont.woff2",
	}

	css = []string{
		"stylesheet.css",
		"bootstrap.min.css",
		"font-awesome.min.css",
	}

	js = []string{
		"jquery-3.2.1.min.js",
		"bootstrap-3.3.7.min.js",
	}

	jsroot = []string{
		"scroll.js",
		"jquery.scrollTo.min.js",
	}

	build = []string{
		"index.html",
		"navData.js",
	}
)

type openapiTemplateData struct {
	Dependents []string
	Path       string
}

type schemaWriterTemplateData struct {
	AllImportPath string
}

// InstallOpenApiGen installs the openapi-gen tool from the github.com/kubernetes/kube-openapi repository.
func InstallOpenApiGen(version string, gopath string) error {
	util.AppLogger().Infof("installing %s with version %s via 'go get' to %s", openApiGen, version, gopath)
	err := util.GoGet(openApiGen, version, gopath, true, false, true)
	if err != nil {
		return err
	}

	return nil
}

// GenerateOpenApi generates the OpenAPI structs and schema files.
// It looks at the specified groupsWithVersions in inputPackage and generates to outputPackage (
// relative to the module outputBase). Any openApiDependencies also have OpenAPI structs generated.
// A boilerplateFile is written to the top of any generated files.
// The gitter client is used to ensure the correct versions of dependencies are loaded.
func GenerateOpenApi(groupsWithVersions []string, inputPackage string, outputPackage string, relativePackage string,
	outputBase string, openApiDependencies []string, moduleDir string, moduleName string, boilerplateFile string, gopath string, semVer string) error {
	basePkg := fmt.Sprintf("%s/openapi", outputPackage)
	corePkg := fmt.Sprintf("%s/core", basePkg)
	allPkg := fmt.Sprintf("%s/all", basePkg)

	// Generate the dependent openapi structs as these are missing from the k8s client
	dependentPackages, err := generateOpenApiDependenciesStruct(outputPackage, relativePackage, outputBase,
		openApiDependencies, moduleDir, moduleName, boilerplateFile, gopath)
	if err != nil {
		return err
	}
	// Generate the main openapi struct
	err = defaultGenerate(openApiGenerator, "openapi", groupsWithVersions, inputPackage,
		corePkg, outputBase, boilerplateFile, gopath, "--output-package", corePkg)
	if err != nil {
		return err
	}
	_, err = writeOpenApiAll(outputBase, allPkg, corePkg, dependentPackages, semVer)
	if err != nil {
		return err
	}
	_, err = writeSchemaWriterToDisk(outputBase, basePkg, allPkg, semVer)
	if err != nil {
		return err
	}
	return nil
}

// writeOpenApiAll code generates a file in openapi/all that reads in all the generated openapi structs and puts them
// in a single map, allowing them to be used by the schema writer and the CRD registration.
// baseDir is the root of the module, outputPackage is the base path of the output package,
// path is the path to the core openapi package (those that are generated for module the generator is run against),
// and dependents is the paths to the dependent openapi packages
func writeOpenApiAll(baseDir string, outputPackage string, path string, dependents []string, semVer string) (string,
	error) {
	tmpl, err := template.New("openapi").Parse(openapiTemplateSrc)
	if err != nil {
		return "", errors.Wrapf(err, "parsing template for openapi_generated.go")
	}
	outputDir := filepath.Join(baseDir, outputPackage)
	err = os.MkdirAll(outputDir, 0700)
	if err != nil {
		return "", errors.Wrapf(err, "creating directory %s", outputDir)
	}
	outFilePath := filepath.Join(outputDir, "openapi_generated.go")
	outFile, err := os.Create(outFilePath)
	if err != nil {
		return "", errors.Wrapf(err, "creating file %s", outFilePath)
	}
	data := &openapiTemplateData{
		Path:       path,
		Dependents: dependents,
	}
	if semVer != "" {
		data.Path = strings.ReplaceAll(path, "/pkg/", fmt.Sprintf("/%s/pkg/", semVer))
		data.Dependents = []string{}
		for _, d := range dependents {
			data.Dependents = append(data.Dependents, strings.ReplaceAll(d, "/pkg/", fmt.Sprintf("/%s/pkg/", semVer)))
		}
	}
	err = tmpl.Execute(outFile, data)
	defer func() {
		err := outFile.Close()
		if err != nil {
			util.AppLogger().Errorf("error closing %s %v\n", outFilePath, err)
		}
	}()
	if err != nil {
		return "", errors.Wrapf(err, "templating %s", outFilePath)
	}
	return outputPackage, nil
}

// writeSchemaWriterToDisk code generates a simple main function that can be called to write the contents of all the
// OpenAPI structs out to JSON and YAML. It's implemented like this to allow us to automatically call the schema
// writer without requiring the user to write a command themselves. baseDir is the path to the module,
// outputPackage is the path to the outputPacakge for the code generator,
// and allImportPath is the path to the package where the generated map of all the structs is
func writeSchemaWriterToDisk(baseDir string, outputPackage string, allImportPath string, semVer string) (string, error) {
	tmpl, err := template.New("schema_writer").Parse(schemaWriterTemplateSrc)
	if err != nil {
		return "", errors.Wrapf(err, "parsing template for %s", SchemaWriterSrcFileName)
	}
	outputDir := filepath.Join(baseDir, outputPackage)
	err = os.MkdirAll(outputDir, 0700)
	if err != nil {
		return "", errors.Wrapf(err, "creating directory %s", outputDir)
	}
	outFilePath := filepath.Join(outputDir, SchemaWriterSrcFileName)
	outFile, err := os.Create(outFilePath)
	if err != nil {
		return "", errors.Wrapf(err, "creating file %s", outFilePath)
	}
	data := &schemaWriterTemplateData{
		AllImportPath: allImportPath,
	}
	if semVer != "" {
		data.AllImportPath = strings.ReplaceAll(allImportPath, "/pkg/", fmt.Sprintf("/%s/pkg/", semVer))
	}
	err = tmpl.Execute(outFile, data)
	defer func() {
		err := outFile.Close()
		if err != nil {
			util.AppLogger().Errorf("error closing %s %v\n", outFilePath, err)
		}
	}()
	if err != nil {
		return "", errors.Wrapf(err, "templating %s", outFilePath)
	}
	return outputPackage, nil
}

// WriteSchemaToDisk is called by the code generated main function to marshal the contents of the OpenAPI structs and
// write them to disk. outputDir is the dir to write the json and yaml files to,
// you can also provide the title and version for the OpenAPI spec.
// definitions is the function that returns all the openapi definitions.
// WriteSchemaToDisk will rewrite the definitions to a dot-separated notation, reversing the initial domain name
func WriteSchemaToDisk(outputDir string, title string, version string, definitions common.GetOpenAPIDefinitions,
	names []string) error {
	err := os.MkdirAll(outputDir, 0700)
	if err != nil {
		return errors.Wrapf(err, "creating --output-directory %s", outputDir)
	}
	config := common.Config{
		Info: &spec.Info{
			InfoProps: spec.InfoProps{
				Version: version,
				Title:   title,
			},
		},
		GetDefinitions: definitions,
		GetDefinitionName: func(name string) (string, spec.Extensions) {
			// For example "github.com/jenkins-x/jx/v2/pkg/apis/jenkins.io/v1.AppSpec"
			parts := strings.Split(name, "/")
			if len(parts) < 3 {
				// Can't do anything with it, return raw
				return name, nil
			}
			var result []string
			for i, part := range parts {
				// handle the domain at the start of the package
				if i == 0 {
					subparts := strings.Split(part, ".")
					for j := len(subparts) - 1; j >= 0; j-- {
						result = append(result, subparts[j])
					}
				} else if i < len(parts)-1 {
					// The docs generator can't handle a dot in the group name, so we remove it
					result = append(result, strings.Replace(part, ".", "_", -1))
				} else {
					result = append(result, part)
				}
			}
			return strings.Join(result, "."), nil
		},
	}

	spec, err := builder.BuildOpenAPIDefinitionsForResources(&config, names...)
	if err != nil {
		return errors.Wrapf(err, "building openapi definitions for %s", names)
	}
	bytes, err := json.Marshal(spec)
	if err != nil {
		return errors.Wrapf(err, "marshaling openapi definitions to json for %s", names)
	}
	outFile := filepath.Join(outputDir, OpenApiV2JSON)
	err = ioutil.WriteFile(outFile, bytes, 0644)
	if err != nil {
		return errors.Wrapf(err, "writing openapi definitions for %s to %s", names, outFile)
	}
	return nil
}

func packageToDirName(pkg string) string {
	str := strings.Join(strings.Split(pkg, "/"), "_")
	str = strings.Join(strings.Split(str, "."), "_")
	return str
}

// GenerateSchema calls the generated schema writer and then loads the output and also writes out a yaml version. The
// outputDir is the base directory for writing the schemas to (they get put in the openapi-spec subdir),
// inputPackage is the package in which generated code lives, inputBase is the path to the module,
// title and version are used in the OpenAPI spec files.
func GenerateSchema(outputDir string, inputPackage string, inputBase string, title string, version string, gopath string) error {
	schemaWriterSrc := filepath.Join(inputPackage, OpenApiDir, SchemaWriterSrcFileName)
	schemaWriterBinary, err := ioutil.TempFile("", "")
	outputDir = filepath.Join(outputDir, "openapi-spec")
	defer func() {
		err := util.DeleteFile(schemaWriterBinary.Name())
		if err != nil {
			util.AppLogger().Warnf("error cleaning up tempfile %s created to compile %s to %v",
				schemaWriterBinary.Name(), SchemaWriterSrcFileName, err)
		}
	}()
	if err != nil {
		return errors.Wrapf(err, "creating tempfile to compile %s to %v", SchemaWriterSrcFileName, err)
	}
	cmd := util.Command{
		Dir:  inputBase,
		Name: "go",
		Args: []string{
			"build",
			"-o",
			schemaWriterBinary.Name(),
			schemaWriterSrc,
		},
		Env: map[string]string{
			"GO111MODULE": "on",
			"GOPATH":      gopath,
		},
	}
	out, err := cmd.RunWithoutRetry()
	if err != nil {
		return errors.Wrapf(err, "running %s, output %s", cmd.String(), out)
	}
	fileJSON := filepath.Join(outputDir, OpenApiV2JSON)
	fileYAML := filepath.Join(outputDir, OpenApiV2YAML)
	cmd = util.Command{
		Name: schemaWriterBinary.Name(),
		Args: []string{
			"--output-directory",
			outputDir,
			"--title",
			title,
			"--version",
			version,
		},
	}
	out, err = cmd.RunWithoutRetry()
	if err != nil {
		return errors.Wrapf(err, "running %s, output %s", cmd.String(), out)
	}
	// Convert to YAML as well
	bytes, err := ioutil.ReadFile(fileJSON)
	if err != nil {
		return errors.Wrapf(err, "reading %s", fileJSON)
	}
	yamlBytes, err := yaml.JSONToYAML(bytes)
	if err != nil {
		return errors.Wrapf(err, "converting %s to yaml", fileJSON)
	}
	err = ioutil.WriteFile(fileYAML, yamlBytes, 0644)
	if err != nil {
		return errors.Wrapf(err, "writing %s", fileYAML)
	}
	return nil
}

func getOutputPackageForOpenApi(pkg string, groupWithVersion []string, outputPackage string) (string, error) {
	if len(groupWithVersion) != 2 {
		return "", errors.Errorf("groupWithVersion must be of length 2 but is %s", groupWithVersion)
	}
	version := groupWithVersion[1]
	if version == "" {
		version = "unversioned"
	}
	return filepath.Join(outputPackage, "openapi", fmt.Sprintf("%s_%s_%s", toValidPackageName(packageToDirName(pkg)),
		toValidPackageName(groupWithVersion[0]),
		toValidPackageName(version))), nil
}

func toValidPackageName(pkg string) string {
	return strings.Replace(strings.Replace(pkg, "-", "_", -1), ".", "_", -1)
}

func generateOpenApiDependenciesStruct(outputPackage string, relativePackage string, outputBase string,
	openApiDependencies []string, moduleDir string, moduleName string, boilerplateFile string, gopath string) ([]string, error) {
	paths := make([]string, 0)
	modulesRequirements, err := util.GetModuleRequirements(moduleDir, gopath)
	if err != nil {
		return nil, errors.Wrapf(err, "getting module requirements for %s", moduleDir)
	}
	for _, d := range openApiDependencies {
		outputPackage, err := generate(d, outputPackage, relativePackage, outputBase, moduleName, boilerplateFile, gopath, modulesRequirements)
		if err != nil {
			return nil, errors.Wrapf(err, "generating open api dependency %s", d)
		}
		paths = append(paths, outputPackage)
	}
	return paths, nil
}

func generate(d string, outputPackage string, relativePackage string, outputBase string, moduleName string, boilerplateFile string, gopath string, modulesRequirements map[string]map[string]string) (string, error) {
	// First, make sure we have the right files in our .jx GOPATH
	// Use go mod to find out the dependencyVersion for our main tree
	ds := strings.Split(d, ":")
	if len(ds) != 4 {
		return "", errors.Errorf("--open-api-dependency %s must be of the format path:package:group"+
			":apiVersion", d)
	}
	path := ds[0]
	pkg := ds[1]
	group := ds[2]
	version := ds[3]
	groupsWithVersions := []string{
		fmt.Sprintf("%s:%s", group, version),
	}
	modules := false
	if strings.Contains(path, "?modules") {
		path = strings.TrimSuffix(path, "?modules")
		modules = true
	}

	dependencyVersion := "master"
	if moduleRequirement, ok := modulesRequirements[moduleName]; !ok {
		util.AppLogger().Warnf("unable to find module requirement for %s, please add it to your go.mod, "+
			"for now using HEAD of the master branch", moduleName)
	} else {
		if requirementVersion, ok := moduleRequirement[path]; !ok {
			util.AppLogger().Warnf("unable to find module requirement version for %s (module %s), "+
				"please add it to your go.mod, "+
				"for now using HEAD of the master branch", path, moduleName)
		} else {
			dependencyVersion = requirementVersion
		}
	}

	if strings.HasPrefix(dependencyVersion, "v0.0.0-") {
		parts := strings.Split(dependencyVersion, "-")
		if len(parts) != 3 {
			return "", errors.Errorf("unable to parse dependencyVersion %s", dependencyVersion)
		}
		// this is the sha
		dependencyVersion = parts[2]
	}
	err := util.GoGet(path, dependencyVersion, gopath, modules, true, true)
	if err != nil {
		return "", errors.WithStack(err)
	}
	// Now we can run the generator against it
	generator := openApiGenerator
	modifiedOutputPackage, err := getOutputPackageForOpenApi(path, []string{group, version}, outputPackage)
	if err != nil {
		return "", errors.Wrapf(err, "getting filename for openapi structs for %s", d)
	}
	err = defaultGenerate(generator,
		"openapi",
		groupsWithVersions,
		filepath.Join(path, pkg),
		modifiedOutputPackage,
		outputBase,
		boilerplateFile,
		gopath,
		"--output-package",
		modifiedOutputPackage)
	if err != nil {
		return "", errors.WithStack(err)
	}
	relativeOutputPackage, err := getOutputPackageForOpenApi(path, []string{group, version}, relativePackage)
	if err != nil {
		return "", errors.Wrapf(err, "getting filename for openapi structs for %s", d)
	}
	// the generator forgets to add the spec import in some cases
	generatedFile := filepath.Join(relativeOutputPackage, "openapi_generated.go")
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, generatedFile, nil, parser.ParseComments)
	if err != nil {
		return "", errors.Wrapf(err, "parsing %s", generatedFile)
	}
	found := false
	for _, imp := range f.Imports {
		if strings.Trim(imp.Path.Value, "\"") == "github.com/go-openapi/spec" {
			found = true
			break
		}
	}
	if !found {
		// Add the imports
		for i := 0; i < len(f.Decls); i++ {
			d := f.Decls[i]

			switch d.(type) {
			case *ast.FuncDecl:
				// No action
			case *ast.GenDecl:
				dd := d.(*ast.GenDecl)

				// IMPORT Declarations
				if dd.Tok == token.IMPORT {
					// Add the new import
					iSpec := &ast.ImportSpec{Path: &ast.BasicLit{Value: strconv.Quote("github.com/go-openapi/spec")}}
					dd.Specs = append(dd.Specs, iSpec)
				}
			}
		}

		// Sort the imports
		ast.SortImports(fs, f)
		var buf bytes.Buffer
		err = format.Node(&buf, fs, f)
		if err != nil {
			return "", errors.Wrapf(err, "convert AST to []byte for %s", generatedFile)
		}
		// Manually add new lines after build tags
		lines := strings.Split(string(buf.Bytes()), "\n")
		buf.Reset()
		for _, line := range lines {
			buf.WriteString(line)
			buf.WriteString("\r\n")
			if strings.HasPrefix(line, "// +") {
				buf.WriteString("\r\n")
			}
		}

		err = ioutil.WriteFile(generatedFile, buf.Bytes(), 0644)
		if err != nil {
			return "", errors.Wrapf(err, "writing %s", generatedFile)
		}
	}
	return modifiedOutputPackage, nil
}
