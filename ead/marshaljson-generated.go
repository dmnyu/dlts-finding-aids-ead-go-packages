// Code generated by generate.go; DO NOT EDIT.

package ead

import (
	"encoding/json"
	"regexp"
)

func (abstract *Abstract) MarshalJSON() ([]byte, error) {
	type AbstractWithTags Abstract

	result, err := getConvertedTextWithTags(abstract.Value)
	if err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(&struct {
		Value string `json:"value,chardata,omitempty"`
		*AbstractWithTags
	}{
		Value:            string(result),
		AbstractWithTags: (*AbstractWithTags)(abstract),
	})
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func (bibref *BibRef) MarshalJSON() ([]byte, error) {
	type BibRefWithTags BibRef

	result, err := getConvertedTextWithTags(bibref.Value)
	if err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(&struct {
		Value string `json:"value,chardata,omitempty"`
		*BibRefWithTags
	}{
		Value:          string(result),
		BibRefWithTags: (*BibRefWithTags)(bibref),
	})
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func (head *Head) MarshalJSON() ([]byte, error) {
	type HeadWithTags Head

	result, err := getConvertedTextWithTags(head.Value)
	if err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(&struct {
		Value string `json:"value,chardata,omitempty"`
		*HeadWithTags
	}{
		Value:        string(result),
		HeadWithTags: (*HeadWithTags)(head),
	})
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func (p *P) MarshalJSON() ([]byte, error) {
	type PWithTags P

	result, err := getConvertedTextWithTags(p.Value)
	if err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(&struct {
		Value string `json:"value,chardata,omitempty"`
		*PWithTags
	}{
		Value:     string(result),
		PWithTags: (*PWithTags)(p),
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

func (unittitle *UnitTitle) MarshalJSON() ([]byte, error) {
	type UnitTitleWithTags UnitTitle

	result, err := getConvertedTextWithTags(unittitle.Value)
	if err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(&struct {
		Value string `json:"value,chardata,omitempty"`
		*UnitTitleWithTags
	}{
		Value:             string(result),
		UnitTitleWithTags: (*UnitTitleWithTags)(unittitle),
	})
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

func (physdesc *PhysDesc) MarshalJSON() ([]byte, error) {
	type PhysDescWithNoWhitespaceOnlyValues PhysDesc

	containsNonWhitespace, err := regexp.MatchString(`\S`, physdesc.Value)
	if err != nil {
		return nil, err
	}

	var value string
	if containsNonWhitespace {
		value = physdesc.Value
	} else {
		value = ""
	}

	jsonData, err := json.Marshal(&struct {
		Value string `json:"value,chardata,omitempty"`
		*PhysDescWithNoWhitespaceOnlyValues
	}{
		Value:                              value,
		PhysDescWithNoWhitespaceOnlyValues: (*PhysDescWithNoWhitespaceOnlyValues)(physdesc),
	})
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}
