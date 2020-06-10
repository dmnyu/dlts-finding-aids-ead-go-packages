package ead

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
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

func (bibRef *BibRef) MarshalJSON() ([]byte, error) {
	type BibRefWithTags BibRef

	result, err := getConvertedTextWithTags(bibRef.Value)
	if err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(&struct {
		Value string `json:"value,chardata,omitempty"`
		*BibRefWithTags
	}{
		Value:          string(result),
		BibRefWithTags: (*BibRefWithTags)(bibRef),
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

func (unitTitle *UnitTitle) MarshalJSON() ([]byte, error) {
	type UnitTitleWithTags UnitTitle

	result, err := getConvertedTextWithTags(unitTitle.Value)
	if err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(&struct {
		Value string `json:"value,chardata,omitempty"`
		*UnitTitleWithTags
	}{
		Value:             string(result),
		UnitTitleWithTags: (*UnitTitleWithTags)(unitTitle),
	})
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func getConvertedTextWithTags(text string) ([]byte, error) {
	decoder := xml.NewDecoder(strings.NewReader(text))

	var result string
	needClosingTag := true
	for {
		token, err := decoder.Token()

		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		switch token := token.(type) {
		case xml.StartElement:
			switch token.Name.Local {
			default:
				result += fmt.Sprintf("<span class=\"ead-%s\">", token.Name.Local)
			case "emph":
				{
					var render string
					for i := range token.Attr {
						if token.Attr[i].Name.Local == "render" {
							render = token.Attr[i].Value
							break
						}
					}
					result += fmt.Sprintf("<span class=\"%s\">", "ead-emph ead-emph-" + render)
				}
			case "lb":
				{
					result += "<br>"
					needClosingTag = false
				}
			}

		case xml.EndElement:
			if needClosingTag {
				result += "</span>"
			} else {
				// Reset
				needClosingTag = true
			}
		case xml.CharData:
			result += string(token)
		}
	}

	return []byte(result), nil
}
