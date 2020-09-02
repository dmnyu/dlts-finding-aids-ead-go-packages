package ead

import (
	"encoding/json"
	"regexp"
)

func (container *Container) MarshalJSON() ([]byte, error) {
	type ContainerWithBarcode Container

	var barcode string
	var label string
	labelContainsBarcodeRegexp := regexp.MustCompile(`^(.+)\s+\[(\d{14})\]$`)
	submatches := labelContainsBarcodeRegexp.FindStringSubmatch(container.Label)
	if (len(submatches) > 0) {
		barcode = submatches[2]
		label = submatches[1]
	} else {
		label = container.Label
	}

	jsonData, err := json.Marshal(&struct {
		Barcode string `json:"barcode,omitempty"`
		Label   string `json:"label,omitempty"`
		*ContainerWithBarcode
	}{
		Barcode:              barcode,
		Label:                label,
		ContainerWithBarcode: (*ContainerWithBarcode)(container),
	})
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func (did *DID) MarshalJSON() ([]byte, error) {
	type DIDWithNoEmptyPhysDesc DID

	containsNonWhitespaceRegexp := regexp.MustCompile(`\S`)
	var physDescNoEmpties []*PhysDesc
	for _, el := range did.PhysDesc {
		if el.Extent != nil || containsNonWhitespaceRegexp.MatchString(el.Value) {
			physDescNoEmpties = append(physDescNoEmpties, el)
		}
	}

	var jsonData []byte
	var err error
	if physDescNoEmpties != nil {
		jsonData, err = json.Marshal(&struct {
			PhysDesc     []*PhysDesc     `xml:"physdesc" json:"physdesc,omitempty"`
			*DIDWithNoEmptyPhysDesc
		}{
			PhysDesc: physDescNoEmpties,
			DIDWithNoEmptyPhysDesc: (*DIDWithNoEmptyPhysDesc)(did),
		})
	} else {
		jsonData, err = json.Marshal(&struct {
			PhysDesc     []*PhysDesc     `xml:"physdesc" json:"physdesc,omitempty"`
			*DIDWithNoEmptyPhysDesc
		}{
			PhysDesc: nil,
			DIDWithNoEmptyPhysDesc: (*DIDWithNoEmptyPhysDesc)(did),
		})
	}

	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func (dao *DAO) MarshalJSON() ([]byte, error) {
	type DAOWithNoWhitespaceOnlyValues DAO

	containsNonWhitespace, err := regexp.MatchString(`\S`, dao.Value)
	if err != nil {
		return nil, err
	}

	var value string
	if containsNonWhitespace {
		value = dao.Value
	} else {
		value = ""
	}

	jsonData, err := json.Marshal(&struct {
		Value string `json:"value,chardata,omitempty"`
		*DAOWithNoWhitespaceOnlyValues
	}{
		Value:                         value,
		DAOWithNoWhitespaceOnlyValues: (*DAOWithNoWhitespaceOnlyValues)(dao),
	})
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

// The custom marshalling for DID will prevent PhysDesc from having a Value field
// that is all whitespace if Extent is nil, but won't prevent PhysDesc from having
// a Value field that is all whitespace if Extent is not nil.
// We need to convert Value field values like "\n    \n    \n" to empty string
// so they can be removed by omitempty struct tag.
func (physDesc *PhysDesc) MarshalJSON() ([]byte, error) {
	type PhysDescWithNoWhitespaceOnlyValues PhysDesc

	containsNonWhitespace, err := regexp.MatchString(`\S`, physDesc.Value)
	if err != nil {
		return nil, err
	}

	var value string
	if containsNonWhitespace {
		value = physDesc.Value
	} else {
		value = ""
	}

	jsonData, err := json.Marshal(&struct {
		Value string `json:"value,chardata,omitempty"`
		*PhysDescWithNoWhitespaceOnlyValues
	}{
		Value:                              value,
		PhysDescWithNoWhitespaceOnlyValues: (*PhysDescWithNoWhitespaceOnlyValues)(physDesc),
	})
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

