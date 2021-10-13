package ead

import (
	"encoding/json"
	"regexp"
)

// Note that this custom marshalling for DID will prevent PhysDesc from having a Value field
// that is all whitespace if Extent is nil, but won't prevent PhysDesc from having
// a Value field that is all whitespace if Extent is not nil.
// We need to convert Value field values like "\n    \n    \n" to empty string
// so they can be removed by omitempty struct tag.  This is done in the PhysDesc.MarshalJSON
// in marshal-generated.go.
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

func (accessTermWithRole *AccessTermWithRole) MarshalJSON() ([]byte, error) {
	type accessTermWithRoleWithTranslatedRelatorCode AccessTermWithRole

	var (
		role string
		err error
	)
	if accessTermWithRole.Role != "" {
		role, err = getRelatorAuthoritativeLabel(accessTermWithRole.Role)
		if err != nil {
			return nil, err
		}
	}

	result, err := getConvertedTextWithTags(accessTermWithRole.Value)
	if err != nil {
		return nil, err
	}
	accessTermWithRole.Value = string(result)

	jsonData, err := json.Marshal(&struct {
		Role string `xml:"role,attr" json:"role,omitempty"`
		*accessTermWithRoleWithTranslatedRelatorCode
	}{
		Role: role,
		accessTermWithRoleWithTranslatedRelatorCode: (*accessTermWithRoleWithTranslatedRelatorCode)(accessTermWithRole),
	})
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func (titleproper *TitleProper) MarshalJSON() ([]byte, error) {
	type TitleProperWithTags TitleProper

	result, err := getConvertedTextWithTagsNoLBConversion(titleproper.Value)
	if err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(&struct {
		Value string `json:"value,chardata,omitempty"`
		*TitleProperWithTags
	}{
		Value:               string(result),
		TitleProperWithTags: (*TitleProperWithTags)(titleproper),
	})
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}


/* TitleStmt.TitleProper flattening strategy:
   func (titleStmt *TitleStmt) MarshalJSON() ([]byte, error) {
      var titleStmtOutput  TitleStmtOutput
      // convert TitleProper array into a single string
      // marshall JSON on the output variable and return the result
*/

func (titleStmt *TitleStmt) MarshalJSON() ([]byte, error) {
	return titleStmt.MarshalJSON()
}


