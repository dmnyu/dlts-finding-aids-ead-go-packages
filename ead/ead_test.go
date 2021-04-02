package ead

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"bytes"
	"testing"
)

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
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

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
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		jsonData, err := json.MarshalIndent(ead, "", "    ")
		if err != nil {
			t.Errorf("Unexpected error marshaling JSON: %s", err)
		}

		referenceFileContents, err := ioutil.ReadFile("./testdata/v0.0.0/mos_2021.json")
		if err != nil {
			t.Errorf("Unexpected error reading reference file: %s", err)
		}

		if !bytes.Equal(referenceFileContents, jsonData) {
			t.Errorf("JSON Data does not match reference file")
		}
	})
}
