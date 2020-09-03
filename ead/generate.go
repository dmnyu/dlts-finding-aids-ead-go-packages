// +build ignore

package main

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"sort"
	"strings"
	"text/template"
)

// NOTE: Currently there is no way to escape backticks in raw strings:
// https://github.com/golang/go/issues/18221
const convertTextWithTagsMarshalJSONCodeTemplate = `func ({{.VarName}} *{{.TypeName}}) MarshalJSON() ([]byte, error) {
	type {{.TypeName}}WithTags {{.TypeName}}

	result, err := {{.ConversionFunction}}({{.VarName}}.Value)
	if err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(&struct {
		Value string ` + "`" + `json:"value,chardata,omitempty"` + "`\n" +
`		*{{.TypeName}}WithTags
	}{
		Value:             string(result),
		{{.TypeName}}WithTags: (*{{.TypeName}}WithTags)({{.VarName}}),
	})
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}`

const omitWhitespaceOnlyValueFieldsMarshalJSONCodeTemplate = `func ({{.VarName}} *{{.TypeName}}) MarshalJSON() ([]byte, error) {
	type {{.TypeName}}WithNoWhitespaceOnlyValues {{.TypeName}}

	containsNonWhitespace, err := regexp.MatchString(` + "`\\S`" + `, {{.VarName}}.Value)
	if err != nil {
		return nil, err
	}

	var value string
	if containsNonWhitespace {
		value = {{.VarName}}.Value
	} else {
		value = ""
	}

	jsonData, err := json.Marshal(&struct {
		Value string ` + "`" + `json:"value,chardata,omitempty"` + "`" + `
		*{{.TypeName}}WithNoWhitespaceOnlyValues
	}{
		Value: value,
		{{.TypeName}}WithNoWhitespaceOnlyValues: (*{{.TypeName}}WithNoWhitespaceOnlyValues)({{.VarName}}),
	})
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}`

func main() {
	w := new(bytes.Buffer)

	w.WriteString(`// Code generated by generate.go; DO NOT EDIT.

package ead

import (
	"encoding/json"
	"regexp"
)`)

	writeConvertTextWithTagsCodeToBuffer(w)
	writeOmitWhitespaceOnlyValueFieldsCodeToBuffer(w)

	// Format with gofmt
	out, err := format.Source(w.Bytes())
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("marshaljson-generated.go", out, 0600)
	if err != nil {
		panic(err)
	}
}

func writeConvertTextWithTagsCodeToBuffer(w *bytes.Buffer) {
	type templateData struct{
		ConversionFunction string
		TypeName string
		VarName string
	}

	t := template.Must(template.New("").Parse(convertTextWithTagsMarshalJSONCodeTemplate))

	conversionFunctionsForTypes := map[string]string{
		"Abstract" : "getConvertedTextWithTags",
		"BibRef" : "getConvertedTextWithTags",
		"Head" : "getConvertedTextWithTags",
		"P" : "getConvertedTextWithTags",
		"TitleProper" : "getConvertedTextWithTagsNoLBConversion",
		"UnitTitle" : "getConvertedTextWithTags",
	}

	sortedTypes := make([]string, len(conversionFunctionsForTypes))
	i := 0
	for k := range conversionFunctionsForTypes {
		sortedTypes[i] = k
		i++
	}
	sort.Strings(sortedTypes)

	for _, typeName := range sortedTypes {
		conversionFunction := conversionFunctionsForTypes[typeName]
		w.WriteString("\n\n")

		err := t.Execute(w, templateData{
			ConversionFunction : conversionFunction,
			TypeName: typeName,
			VarName:  strings.ToLower(typeName),
		})
		if err != nil {
			panic(err)
		}
	}
}

func writeOmitWhitespaceOnlyValueFieldsCodeToBuffer(w *bytes.Buffer) {
	type templateData struct{
		TypeName string
		VarName string
	}

	t := template.Must(template.New("").Parse(omitWhitespaceOnlyValueFieldsMarshalJSONCodeTemplate))

	sortedTypes := []string{
		"DAO",
		"PhysDesc",
	}

	for _, typeName := range sortedTypes {
		w.WriteString("\n\n")

		err := t.Execute(w, templateData{
			TypeName: typeName,
			VarName:  strings.ToLower(typeName),
		})
		if err != nil {
			panic(err)
		}
	}
}
