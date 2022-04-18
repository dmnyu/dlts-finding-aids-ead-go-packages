package ead

import (
	"encoding/xml"
	"fmt"
)

type EADChild struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value,omitempty"`
}

func (eadChild *EADChild) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tagName := start.Name.Local
	switch tagName {
	case "chronlist":
		return decodeChronList(eadChild, d, start)
	case "defitem":
		return decodeDefItem(eadChild, d, start)
	case "extref":
		return decodeExtref(eadChild, d, start)
	case "legalstatus":
		return decodeLegalStatus(eadChild, d, start)
	case "list":
		return decodeList(eadChild, d, start)
	case "p":
		return decodeP(eadChild, d, start)
	default:
		return fmt.Errorf("unsupported element error %s", tagName)
	}
}

func decodeLegalStatus(eadChild *EADChild, d *xml.Decoder, start xml.StartElement) error {
	var ls LegalStatus
	if err := d.DecodeElement(&ls, &start); err != nil {
		return err
	}
	eadChild.Name = start.Name.Local
	eadChild.Value = ls
	return nil
}

func decodeChronList(eadChild *EADChild, d *xml.Decoder, start xml.StartElement) error {
	var cl *ChronList
	if err := d.DecodeElement(&cl, &start); err != nil {
		return err
	}
	eadChild.Name = start.Name.Local
	eadChild.Value = cl
	return nil
}

func decodeDefItem(eadChild *EADChild, d *xml.Decoder, start xml.StartElement) error {
	var di *DefItem
	if err := d.DecodeElement(&di, &start); err != nil {
		return err
	}
	eadChild.Name = start.Name.Local
	eadChild.Value = di
	return nil
}

func decodeExtref(eadChild *EADChild, d *xml.Decoder, start xml.StartElement) error {
	var er *ExtRef
	if err := d.DecodeElement(&er, &start); err != nil {
		return err
	}
	eadChild.Name = start.Name.Local
	eadChild.Value = er
	return nil
}

func decodeList(eadChild *EADChild, d *xml.Decoder, start xml.StartElement) error {
	var l *List
	if err := d.DecodeElement(&l, &start); err != nil {
		return err
	}
	eadChild.Name = start.Name.Local
	eadChild.Value = l
	return nil
}

func decodeP(eadChild *EADChild, d *xml.Decoder, start xml.StartElement) error {
	var p *P
	if err := d.DecodeElement(&p, &start); err != nil {
		return err
	}
	eadChild.Name = start.Name.Local
	eadChild.Value = p
	return nil
}
