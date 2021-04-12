package ead

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"testing"
	"time"
)

const testFixturePath string = "./testdata/v0.1.1"

func failOnError(t *testing.T, err error, label string) {
	if err != nil {
		t.Errorf("%s: %s", label, err)
		t.FailNow()
	}
}

func assertEqual(t *testing.T, want string, got string, label string) {
	if want != got {
		t.Errorf("%s Mismatch: want: %s, got: %s", label, want, got)
	}
}

func getOmegaEAD(t *testing.T) EAD {
	EADXML, err := ioutil.ReadFile(testFixturePath + "/" + "Omega-EAD.xml")
	failOnError(t, err, "Unexpected error")

	var ead EAD
	err = xml.Unmarshal([]byte(EADXML), &ead)
	failOnError(t, err, "Unexpected error")

	return ead
}

func TestXMLParsing(t *testing.T) {
	t.Run("XML Parsing", func(t *testing.T) {
		ead := getOmegaEAD(t)

		want := "collection"
		got := ead.ArchDesc.Level
		assertEqual(t, want, got, "ArchDesc.Level")
	})
}

func TestJSONMarshaling(t *testing.T) {
	t.Run("JSON Marshaling", func(t *testing.T) {
		ead := getOmegaEAD(t)

		jsonData, err := json.MarshalIndent(ead, "", "    ")
		failOnError(t, err, "Unexpected error marshaling JSON")

		// reference file includes newline at end of file so
		// add newline to jsonData
		jsonData = append(jsonData, '\n')

		referenceFile := testFixturePath + "/" + "mos_2021.json"
		referenceFileContents, err := ioutil.ReadFile(referenceFile)
		failOnError(t, err, "Unexpected error reading reference file")

		if !bytes.Equal(referenceFileContents, jsonData) {
			jsonFile := "./testdata/tmp/failing-marshal.json"
			err = ioutil.WriteFile(jsonFile, []byte(jsonData), 0644)
			failOnError(t, err, fmt.Sprintf("Unexpected error writing %s", jsonFile))

			errMsg := fmt.Sprintf("JSON Data does not match reference file %s. Wrote marshaled JSON to: %s", referenceFile, jsonFile)
			t.Errorf(errMsg)
		}
	})
}

func TestUpdateRunInfo(t *testing.T) {
	t.Run("JSON Marshaling", func(t *testing.T) {
		var sut EAD

		want := ""
		got := sut.RunInfo.PkgVersion
		assertEqual(t, want, got, "Initial ead.RunInfo.PkgVersion")

		want = "0001-01-01T00:00:00Z"
		got = sut.RunInfo.TimeStamp.Format(time.RFC3339)
		assertEqual(t, want, got, "Initial ead.RunInfo.TimeStamp")

		now := time.Now()
		version := Version // from ead package constant

		sut.RunInfo.SetRunInfo(version, now)

		want = version
		got = sut.RunInfo.PkgVersion
		assertEqual(t, want, got, "Post-assignment ead.RunInfo.PkgVersion")

		want = now.Format(time.RFC3339)
		got = sut.RunInfo.TimeStamp.Format(time.RFC3339)
		assertEqual(t, want, got, "Initial ead.RunInfo.TimeStamp")

	})
}
