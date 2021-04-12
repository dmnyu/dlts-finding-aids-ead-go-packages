package ead

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"bytes"
	"fmt"
	"testing"
)


func failOnError(t *testing.T, err error, label string) {
	if err != nil {
		t.Errorf("%s: %s", label, err)
		t.FailNow()
	}
}

func assert(t *testing.T, want string, got string, label string) {
	if want != got {
		t.Errorf("%s Mismatch: want: %s, got: %s", label, want, got)
	}
}

func TestXMLParsing(t *testing.T) {
	t.Run("XML Parsing", func(t *testing.T) {
		EADXML, err := ioutil.ReadFile("./testdata/v0.0.0/Omega-EAD.xml")
		var ead EAD
		err = xml.Unmarshal([]byte(EADXML), &ead)
		failOnError(t, err, "Unexpected error")

		want := "collection"
		got := ead.ArchDesc.Level
		assert(t, want, got, "ArchDesc.Level")
	})
}

func TestJSONMarshaling(t *testing.T) {
	t.Run("JSON Marshaling", func(t *testing.T) {
		EADXML, err := ioutil.ReadFile("./testdata/v0.0.0/Omega-EAD.xml")
		var ead EAD

		err = xml.Unmarshal([]byte(EADXML), &ead)
		failOnError(t, err, "Unexpected error unmarshaling XML")

		jsonData, err := json.MarshalIndent(ead, "", "    ")
		failOnError(t, err, "Unexpected error marshaling JSON")

		// reference file includes newline at end of file so
		// add newline to jsonData
		jsonData = append(jsonData, '\n')

		referenceFileContents, err := ioutil.ReadFile("./testdata/v0.0.0/mos_2021.json")
		failOnError(t, err, "Unexpected error reading reference file")

		if !bytes.Equal(referenceFileContents, jsonData) {
			jsonFile := "./testdata/tmp/failing-marshal.json"
			errMsg := fmt.Sprintf("JSON Data does not match reference file. Writing marshaled JSON to: %s", jsonFile)
			t.Errorf(errMsg)
			_ = ioutil.WriteFile(jsonFile, []byte(jsonData), 0644)
		}
	})
}
