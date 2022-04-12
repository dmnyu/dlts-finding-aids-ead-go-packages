package ead

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"path/filepath"
	"testing"
)

var testFixtureDataPath = filepath.Join("testdata", "xmlorder")

func TestNoteChildOrder(t *testing.T) {

	t.Run("Test Marshal JSON", func(t *testing.T) {
		ead := getOrderXMLOmega(t)
		jsonData, err := json.MarshalIndent(ead, "", "    ")
		failOnError(t, err, "Unexpected error marshaling JSON")

		// reference file includes newline at end of file so
		// add newline to jsonData
		jsonData = append(jsonData, '\n')
		if err := ioutil.WriteFile("omega-ead-test.xml", jsonData, 0755); err != nil {
			failOnError(t, err, "Unexpected error writing json to file")
		}
	})
}

func getOrderXMLOmega(t *testing.T) EAD {
	EADXML, err := ioutil.ReadFile(filepath.Join(testFixtureDataPath, "Omega-EAD.xml"))
	failOnError(t, err, "Unexpected error")

	var ead EAD
	err = xml.Unmarshal([]byte(EADXML), &ead)
	failOnError(t, err, "Unexpected error")
	return ead
}
