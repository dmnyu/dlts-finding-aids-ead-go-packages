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
	switch start.Name.Local {
	case "accessrestrict", "accruals", "acqinfo", "altformavail", "appraisal", "arrangement", "bioghist",
		"custodhist", "odd", "otherfindaid", "originalsloc", "phystech", "prefercite",
		"processinfo", "relatedmaterial", "scopecontent", "separatedmaterial", "userestrict":
		return decodeFormattedNote(eadChild, d, start)
	case "bibliography":
		return decodeBibliography(eadChild, d, start)
	case "controlaccess":
		return decodeControlAccess(eadChild, d, start)
	case "chronlist":
		return decodeChronList(eadChild, d, start)
	case "defitem":
		return decodeDefItem(eadChild, d, start)
	case "did":
		return decodeDID(eadChild, d, start)
	case "dsc":
		return decodeDSC(eadChild, d, start)
	case "extref":
		return decodeExtref(eadChild, d, start)
	case "legalstatus":
		return decodeLegalStatus(eadChild, d, start)
	case "list":
		return decodeList(eadChild, d, start)
	case "p":
		return decodeP(eadChild, d, start)
	default:
		return fmt.Errorf("unsupported element error %s", start.Name.Local)
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

func decodeBibliography(eadChild *EADChild, d *xml.Decoder, start xml.StartElement) error {
	var bib *Bibliography
	if err := d.DecodeElement(&bib, &start); err != nil {
		return err
	}
	eadChild.Name = start.Name.Local
	eadChild.Value = bib
	return nil
}

func decodeControlAccess(eadChild *EADChild, d *xml.Decoder, start xml.StartElement) error {
	var ca *ControlAccess
	if err := d.DecodeElement(&ca, &start); err != nil {
		return err
	}
	eadChild.Name = start.Name.Local
	eadChild.Value = ca
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

func decodeDID(eadChild *EADChild, d *xml.Decoder, start xml.StartElement) error {
	var did *DID
	if err := d.DecodeElement(&did, &start); err != nil {
		return err
	}
	eadChild.Name = start.Name.Local
	eadChild.Value = did
	return nil
}

func decodeDSC(eadChild *EADChild, d *xml.Decoder, start xml.StartElement) error {
	var dsc *DSC
	if err := d.DecodeElement(&dsc, &start); err != nil {
		return err
	}
	eadChild.Name = start.Name.Local
	eadChild.Value = dsc
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

func decodeFormattedNote(eadChild *EADChild, d *xml.Decoder, start xml.StartElement) error {
	var fnh *FormattedNoteWithHead
	if err := d.DecodeElement(&fnh, &start); err != nil {
		return err
	}
	eadChild.Name = start.Name.Local
	eadChild.Value = fnh
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
