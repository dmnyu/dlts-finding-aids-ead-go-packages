package ead

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

// Based on: "Data model for parsing EAD <archdesc> elements": https://jira.nyu.edu/jira/browse/FADESIGN-29
// ...using code generated by zek against https://github.com/nyudlts/samplesFromWeatherly/tree/216c4aab0eae53c1c06a1433b357b1d7df933bd6/FADESIGN-3
// as a starting point for non-<archdesc> elements.

const (
	Version = "0.1.0"
)

type EAD struct {
	ArchDesc  *ArchDesc    `xml:"archdesc" json:"archdesc,omitempty"`
	EADHeader []*EADHeader `xml:"eadheader" json:"eadheader,omitempty"`
}

// https://jira.nyu.edu/jira/browse/FADESIGN-29 additions
type Abstract struct {
	ID    string `xml:"id,attr" json:"id,attr,omitempty"`
	Label string `xml:"label,attr" json:"label,attr,omitempty"`

	Value string `xml:",innerxml" json:"value,chardata,omitempty"`
}

type AddressLine struct {
	ExtPtr []*ExtPtr `xml:"extptr" json:"extptr,omitempty"`
	Value  string    `xml:",chardata" json:"value,chardata,omitempty"`
}

type ArchDesc struct {
	Level string `xml:"level,attr" json:"level,attr,omitempty"`

	AccessRestrict     []*FormattedNoteWithHead `xml:"accessrestrict" json:"accessrestrict,omitempty"`
	Accruals           []*FormattedNoteWithHead `xml:"accruals" json:"accruals,omitempty"`
	AcqInfo            []*FormattedNoteWithHead `xml:"acqinfo" json:"acqinfo,omitempty"`
	AltFormatAvailable []*FormattedNoteWithHead `xml:"altformatavailable" json:"altformatavailable,omitempty"`
	Appraisal          []*FormattedNoteWithHead `xml:"appraisal" json:"appraisal,omitempty"`
	Arrangement        []*FormattedNoteWithHead `xml:"arrangement" json:"arrangement,omitempty"`
	Bibliography       []*Bibliography          `xml:"bibliography" json:"bibliography,omitempty"`
	BiogHist           []*FormattedNoteWithHead `xml:"bioghist" json:"bioghist,omitempty"`
	ControlAccess      []*ControlAccess         `xml:"controlaccess" json:"controlaccess,omitempty"`
	CustodHist         []*FormattedNoteWithHead `xml:"custodhist" json:"custodhist,omitempty"`
	DID                []*DID                   `xml:"did" json:"did,omitempty"`
	DSC                []*DSC                   `xml:"dsc" json:"dsc,omitempty"`
	Odd                []*FormattedNoteWithHead `xml:"odd" json:"odd,omitempty"`
	PhysTech           []*FormattedNoteWithHead `xml:"phystech" json:"phystech,omitempty"`
	PreferCite         []*FormattedNoteWithHead `xml:"prefercite" json:"prefercite,omitempty"`
	ProcessInfo        []*FormattedNoteWithHead `xml:"processinfo" json:"processinfo,omitempty"`
	RelatedMaterial    []*FormattedNoteWithHead `xml:"relatedmaterial" json:"relatedmaterial,omitempty"`
	ScopeContent       []*FormattedNoteWithHead `xml:"scopecontent" json:"scopecontent,omitempty"`
	SeparatedMaterial  []*FormattedNoteWithHead `xml:"separatedmaterial" json:"separatedmaterial,omitempty"`
	UserRestrict       []*FormattedNoteWithHead `xml:"userestrict" json:"userestrict,omitempty"`
}

type Bibliography struct {
	ID string `xml:"id,attr" json:"id,attr,omitempty"`

	Head   []Head    `xml:"head,omitemtpy" json:"head,omitempty"`
	BibRef []*BibRef `xml:"bibref,omitempty" json:"bibref,omitempty"`
}

type BibRef struct {
	Value string `xml:",innerxml" json:"value,chardata,omitempty"`
}

type C struct {
	ID         string `xml:"id,attr" json:"id,attr,omitempty"`
	Level      string `xml:"level,attr" json:"level,attr,omitempty"`
	OtherLevel string `xml:"otherlevel,attr" json:"otherlevel,attr,omitempty"`

	AccessRestrict  []*FormattedNoteWithHead `xml:"accessrestrict,omitempty" json:"accessrestrict,omitempty"`
	Accruals        []*FormattedNoteWithHead `xml:"accruals,omitempty" json:"accruals,omitempty"`
	Appraisal       []*FormattedNoteWithHead `xml:"appraisal,omitempty" json:"appraisal,omitempty"`
	Arrangement     []*FormattedNoteWithHead `xml:"arrangement,omitempty" json:"arrangement,omitempty"`
	BiogHist        []*FormattedNoteWithHead `xml:"bioghist,omitempty" json:"bioghist,omitempty"`
	C               []*C                     `xml:"c,omitempty" json:"c,omitempty"`
	ControlAccess   []*ControlAccess         `xml:"controlaccess" json:"controlaccess,omitempty"`
	CustodHist      []*FormattedNoteWithHead `xml:"custodhist" json:"custodhist,omitempty"`
	DID             []*DID                   `xml:"did,omitempty" json:"did,omitempty"`
	PhysTech        []*FormattedNoteWithHead `xml:"phystech,omitempty" json:"phystech,omitempty"`
	PreferCite      []*FormattedNoteWithHead `xml:"prefercite,omitempty" json:"prefercite,omitempty"`
	ProcessInfo     []*FormattedNoteWithHead `xml:"processinfo,omitempty" json:"processinfo,omitempty"`
	RelatedMaterial []*FormattedNoteWithHead `xml:"relatedmaterial,omitempty" json:"relatedmaterial,omitempty"`
	ScopeContent    []*FormattedNoteWithHead `xml:"scopecontent,omitempty" json:"scopecontent,omitempty"`
	UserRestrict    []*FormattedNoteWithHead `xml:"userrestrict,omitempty" json:"userrestrict,omitempty"`
}

type Change struct {
	Date  []string `xml:"date" json:"date,omitempty"`
	Item  []*Item  `xml:"item" json:"item,omitempty"`
	Value string   `xml:",chardata" json:"value,chardata,omitempty"`
}

type ControlAccess struct {
	CorpName  []string `xml:"corpname" json:"corpname,omitempty"`
	FamName   []string `xml:"famname" json:"famname,omitempty"`
	GenreForm []string `xml:"genreform" json:"genreform,omitempty"`
	PersName  []string `xml:"persname" json:"persname,omitempty"`
	Subject   []string `xml:"subject" json:"subject,omitempty"`
}

type Creation struct {
	Date  []string `xml:"date" json:"date,omitempty"`
	Value string   `xml:",chardata" json:"value,chardata,omitempty"`
}

type DID struct {
	Abstract     []*Abstract     `xml:"abstract" json:"abstract,omitempty"`
	LangMaterial []*LangMaterial `xml:"langmaterial" json:"langmaterial,omitempty"`
	Origination  []*Origination  `xml:"origination" json:"origination,omitempty"`
	PhysDesc     []*PhysDesc     `xml:"physdesc" json:"physdesc,omitempty"`
	PhysLoc      []*PhysLoc      `xml:"physloc" json:"physloc,omitempty"`
	Repository   []*Repository   `xml:"repository" json:"repository,omitempty"`
	UnitDate     []*UnitDate     `xml:"unitdate" json:"unitdate,omitempty"`
	UnitID       []*UnitID       `xml:"unitid" json:"unitid,omitempty"`
	UnitTitle    []*UnitTitle    `xml:"unittitle" json:"unittitle,omitempty"`
}

type DSC struct {
	C    []*C   `xml:"c,omitempty" json:"c,omitempty"`
	Head []Head `xml:"head,omitemtpy" json:"head,omitempty"`
	P    []*P   `xml:"p,omitempty" json:"p,omitempty"`
}

type EADHeader struct {
	EADID        []*EADID        `xml:"eadid" json:"eadid,omitempty"`
	FileDesc     []*FileDesc     `xml:"filedesc" json:"filedesc,omitempty"`
	ProfileDesc  []*ProfileDesc  `xml:"profiledesc" json:"profiledesc,omitempty"`
	RevisionDesc []*RevisionDesc `xml:"revisiondesc" json:"revisiondesc,omitempty"`
}

type EADID struct {
	CountryCode    string `xml:"countrycode,attr" json:"countrycode,attr,omitempty"`
	MainAgencyCode string `xml:"mainagencycode,attr" json:"mainagencycode,attr,omitempty"`
	URL            string `xml:"url,attr" json:"url,attr,omitempty"`

	Value string `xml:",chardata" json:"value,chardata,omitempty"`
}

type Extent struct {
	Value string `xml:",chardata" json:"value,chardata,omitempty"`
}

type ExtPtr struct {
	Href  string   `xml:"href,attr" json:"href,attr,omitempty"`
	Show  string   `xml:"show,attr" json:"show,attr,omitempty"`
	Title string   `xml:"title,attr" json:"title,attr,omitempty"`
	Type  string   `xml:"type,attr" json:"type,attr,omitempty"`
}

type FileDesc struct {
	NoteStmt        []*FormattedNoteWithHead `xml:"notestmt" json:"notestmt,omitempty"`
	PublicationStmt []*PublicationStmt       `xml:"publicationstmt" json:"publicationstmt,omitempty"`
	TitleStmt       []*TitleStmt             `xml:"titlestmt" json:"titlestmt,omitempty"`
}

// "eadnote" in current draft of the data model
type FormattedNoteWithHead struct {
	ID string `xml:"id,attr" json:"id,attr,omitempty"`

	Head []Head `xml:"head,omitemtpy" json:"head,omitempty"`
	P    []*P   `xml:"p,omitempty" json:"p,omitempty"`
}

type Head struct {
	ExtPtr []*ExtPtr `xml:"extptr" json:"extptr,omitempty"`
	Value  string    `xml:",innerxml" json:"value,chardata,omitempty"`
}

type Item struct {
	Value string `xml:",chardata" json:"value,chardata,omitempty"`
}

type LangMaterial struct {
	ID string `xml:"id,attr" json:"id,attr,omitempty"`

	Language []*Language `xml:"language" json:"language,omitempty"`
	Value    string      `xml:",chardata" json:"value,chardata,omitempty"`
}

type Language struct {
	LangCode   string `xml:"langcode,attr" json:"langcode,attr,omitempty"`
	ScriptCode string `xml:"scriptcode,attr" json:"scriptcode,attr,omitempty"`

	Value string `xml:",chardata" json:"value,chardata,omitempty"`
}

type LangUsage struct {
	LangCode string `xml:"langcode,attr" json:"langcode,attr,omitempty"`
}

type Origination struct {
	Label string `xml:"label,attr" json:"label,attr,omitempty"`

	PersName []string `xml:"persname"   json:"persname,omitempty"`
}

type P struct {
	ID string `xml:"id,attr" json:"id,attr,omitempty"`

	Value string `xml:",innerxml" json:"value,chardata,omitempty"`
}

type PhysDesc struct {
	Extent []*Extent `xml:"extent" json:"extent,omitempty"`
}

type PhysLoc struct {
	ID string `xml:"id,attr" json:"id,attr,omitempty"`

	Value string `xml:",chardata" json:"value,chardata,omitempty"`
}

type ProfileDesc struct {
	Creation  []*Creation  `xml:"creation" json:"creation,omitempty"`
	DescRules []*string `xml:"descrules" json:"descrules,omitempty"`
	LangUsage []*LangUsage `xml:"langusage" json:"langusage,omitempty"`
}

type PublicationStmt struct {
	Address   []*AddressLine `xml:"address" json:"address,omitempty"`
	Publisher string         `xml:"publication" json:"publication,omitempty"`
	P         []*P           `xml:"p" json:"p,omitempty"`
}

type Repository struct {
	CorpName []string `xml:"corpname" json:"corpname,omitempty"`
}

type RevisionDesc struct {
	Change []*Change `xml:"change" json:"change,omitempty"`
}

type Title struct {
	Render string `xml:"render,attr" json:"render,attr,omitempty"`
	Value  string `xml:",chardata" json:"value,chardata,omitempty"`
}

type TitleProper struct {
	Value string `xml:",chardata" json:"value,chardata,omitempty"`
}

type TitleStmt struct {
	Author      []string       `xml:"author" json:"author,omitempty"`
	Sponsor     []string       `xml:"sponsor" json:"sponsor,omitempty"`
	TitleProper []*TitleProper `xml:"titleproper" json:"titleproper,omitempty"`
}

type UnitDate struct {
	Normal string `xml:"normal,attr" json:"normal,attr,omitempty"`
	Type   string `xml:"type,attr" json:"type,attr,omitempty"`

	Value string `xml:",chardata" json:"value,chardata,omitempty"`
}

type UnitID struct {
	Value string `xml:",chardata" json:"value,chardata,omitempty"`
}

type UnitTitle struct {
	Title []*Title `xml:"title" json:"title,omitempty"`
	Value string   `xml:",innerxml" json:"value,chardata,omitempty"`
}

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
				result += fmt.Sprintf("<span class=\"%s\">", token.Name.Local)
			case "emph":
				{
					var render string
					for i := range token.Attr {
						if token.Attr[i].Name.Local == "render" {
							render = token.Attr[i].Value
							break
						}
					}
					result += fmt.Sprintf("<span class=\"%s\">", "emph emph-" + render)
				}
			case "lb":
				{
					result += "<br>"
					needClosingTag = false
				}
			}

		case xml.EndElement:
			if (needClosingTag) {
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
