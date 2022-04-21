//go:generate go run generate.go

package ead

// Based on: "Data model for parsing EAD <archdesc> elements": https://jira.nyu.edu/jira/browse/FADESIGN-29.

const (
	Version = "0.14.0"
)

type EAD struct {
	RunInfo   RunInfo   `json:"runinfo"`
	PubInfo   PubInfo   `json:"pubinfo"`
	Donors    Donors    `json:"donors,omitempty"`
	ArchDesc  *ArchDesc `xml:"archdesc" json:"archdesc,omitempty"`
	EADHeader EADHeader `xml:"eadheader" json:"eadheader,omitempty"`
}

type Abstract struct {
	ID FilteredString `xml:"id,attr" json:"id,omitempty"`

	Title []*Title `xml:"title" json:"title,omitempty"`
	Value string   `xml:",innerxml" json:"value,chardata,omitempty"`
}

type AccessTermWithRole struct {
	Role string `xml:"role,attr" json:"role,omitempty"`

	Value string `xml:",innerxml" json:"value,chardata,omitempty"`
}

type Address struct {
	AddressLine []*AddressLine `xml:"addressline" json:"addressline,omitempty"`
}

type AddressLine struct {
	ExtPtr []*ExtPtr `xml:"extptr" json:"extptr,omitempty"`

	Value string `xml:",innerxml" json:"value,chardata,omitempty"`
}

type ArchDesc struct {
	Level    FilteredString `xml:"level,attr" json:"level,attr,omitempty"`
	Children []*EADChild    `xml:",any"`
	/*
		AccessRestrict    []*FormattedNoteWithHead `xml:"accessrestrict" json:"accessrestrict,omitempty"`
		Accruals          []*FormattedNoteWithHead `xml:"accruals" json:"accruals,omitempty"`
		AcqInfo           []*FormattedNoteWithHead `xml:"acqinfo" json:"acqinfo,omitempty"`
		AltFormAvailable  []*FormattedNoteWithHead `xml:"altformavailable" json:"altformavailable,omitempty"`
		Appraisal         []*FormattedNoteWithHead `xml:"appraisal" json:"appraisal,omitempty"`
		Arrangement       []*FormattedNoteWithHead `xml:"arrangement" json:"arrangement,omitempty"`
		Bibliography      []*Bibliography          `xml:"bibliography" json:"bibliography,omitempty"`
		BiogHist          []*FormattedNoteWithHead `xml:"bioghist" json:"bioghist,omitempty"`
		ControlAccess     []*ControlAccess         `xml:"controlaccess" json:"controlaccess,omitempty"`
		CustodHist        []*FormattedNoteWithHead `xml:"custodhist" json:"custodhist,omitempty"`
		DID               DID                      `xml:"did" json:"did,omitempty"`
		DSC               *DSC                     `xml:"dsc" json:"dsc,omitempty"`
		Odd               []*FormattedNoteWithHead `xml:"odd" json:"odd,omitempty"`
		OtherFindAid      []*FormattedNoteWithHead `xml:"otherfindaid" json:"otherfindaid,omitempty"`
		OriginalsLoc      []*FormattedNoteWithHead `xml:"originalsloc" json:"originalsloc,omitempty"`
		PhysTech          []*FormattedNoteWithHead `xml:"phystech" json:"phystech,omitempty"`
		PreferCite        []*FormattedNoteWithHead `xml:"prefercite" json:"prefercite,omitempty"`
		ProcessInfo       []*FormattedNoteWithHead `xml:"processinfo" json:"processinfo,omitempty"`
		RelatedMaterial   []*FormattedNoteWithHead `xml:"relatedmaterial" json:"relatedmaterial,omitempty"`
		ScopeContent      []*FormattedNoteWithHead `xml:"scopecontent" json:"scopecontent,omitempty"`
		SeparatedMaterial []*FormattedNoteWithHead `xml:"separatedmaterial" json:"separatedmaterial,omitempty"`
		UserRestrict      []*FormattedNoteWithHead `xml:"userestrict" json:"userestrict,omitempty"`
	*/
}

type ArchRef struct {
	PhysLoc []*PhysLoc `xml:"physloc" json:"physloc,omitempty"`

	Value string `xml:",innerxml" json:"value,chardata,omitempty"`
}

type Bibliography struct {
	ID FilteredString `xml:"id,attr" json:"id,omitempty"`

	Head   *Head     `xml:"head,omitemtpy" json:"head,omitempty"`
	BibRef []*BibRef `xml:"bibref,omitempty" json:"bibref,omitempty"`
	P      []*P      `xml:"p,omitempty" json:"p,omitempty"`
}

type BibRef struct {
	Title []*Title `xml:"title" json:"title,omitempty"`

	Value string `xml:",innerxml" json:"value,chardata,omitempty"`
}

type C struct {
	ID         FilteredString `xml:"id,attr" json:"id,omitempty"`
	Level      FilteredString `xml:"level,attr" json:"level,omitempty"`
	OtherLevel FilteredString `xml:"otherlevel,attr" json:"otherlevel,omitempty"`

	AccessRestrict    []*FormattedNoteWithHead `xml:"accessrestrict,omitempty" json:"accessrestrict,omitempty"`
	Accruals          []*FormattedNoteWithHead `xml:"accruals,omitempty" json:"accruals,omitempty"`
	AcqInfo           []*FormattedNoteWithHead `xml:"acqinfo,omitempty" json:"acqinfo,omitempty"`
	AltFormAvailable  []*FormattedNoteWithHead `xml:"altformavailable" json:"altformavailable,omitempty"`
	Appraisal         []*FormattedNoteWithHead `xml:"appraisal,omitempty" json:"appraisal,omitempty"`
	Arrangement       []*FormattedNoteWithHead `xml:"arrangement,omitempty" json:"arrangement,omitempty"`
	BiogHist          []*FormattedNoteWithHead `xml:"bioghist,omitempty" json:"bioghist,omitempty"`
	C                 []*C                     `xml:"c,omitempty" json:"c,omitempty"`
	ControlAccess     []*ControlAccess         `xml:"controlaccess" json:"controlaccess,omitempty"`
	CustodHist        []*FormattedNoteWithHead `xml:"custodhist" json:"custodhist,omitempty"`
	DID               DID                      `xml:"did,omitempty" json:"did,omitempty"`
	FilePlan          []*FormattedNoteWithHead `xml:"fileplan,omitempty" json:"fileplan,omitempty"`
	Index             []*Index                 `xml:"index,omitempty" json:"index,omitempty"`
	Odd               []*FormattedNoteWithHead `xml:"odd" json:"odd,omitempty"`
	OtherFindAid      []*FormattedNoteWithHead `xml:"otherfindaid" json:"otherfindaid,omitempty"`
	OriginalsLoc      []*FormattedNoteWithHead `xml:"originalsloc" json:"originalsloc,omitempty"`
	PhysTech          []*FormattedNoteWithHead `xml:"phystech,omitempty" json:"phystech,omitempty"`
	PreferCite        []*FormattedNoteWithHead `xml:"prefercite,omitempty" json:"prefercite,omitempty"`
	ProcessInfo       []*FormattedNoteWithHead `xml:"processinfo,omitempty" json:"processinfo,omitempty"`
	RelatedMaterial   []*FormattedNoteWithHead `xml:"relatedmaterial,omitempty" json:"relatedmaterial,omitempty"`
	ScopeContent      []*FormattedNoteWithHead `xml:"scopecontent,omitempty" json:"scopecontent,omitempty"`
	SeparatedMaterial []*FormattedNoteWithHead `xml:"separatedmaterial" json:"separatedmaterial,omitempty"`
	UserRestrict      []*FormattedNoteWithHead `xml:"userrestrict,omitempty" json:"userrestrict,omitempty"`
}

type Change struct {
	Date []*Date `xml:"date" json:"date,omitempty"`
	Item []*Item `xml:"item" json:"item,omitempty"`
}

type ChronItem struct {
	Date     []*Date     `xml:"date" json:"date,omitempty"`
	EventGrp []*EventGrp `xml:"eventgrp,omitempty" json:"eventgrp,omitempty"`

	Value string `xml:",innerxml" json:"value,chardata,omitempty"`
}

type ChronList struct {
	Head      *Head        `xml:"head,omitemtpy" json:"head,omitempty"`
	ChronItem []*ChronItem `xml:"chronitem,omitempty" json:"chronitem,omitempty"`
}

type Container struct {
	AltRender FilteredString      `xml:"altrender,attr" json:"altrender,omitempty"`
	ID        FilteredString      `xml:"id,attr" json:"id,omitempty"`
	Label     FilteredLabelString `xml:"label,attr" json:"label,omitempty"`
	Parent    FilteredString      `xml:"parent,attr" json:"parent,omitempty"`
	Type      FilteredString      `xml:"type,attr" json:"type,omitempty"`

	Value string `xml:",innerxml" json:"value,chardata,omitempty"`
}

type ControlAccess struct {
	CorpName   []*AccessTermWithRole `xml:"corpname" json:"corpname,omitempty"`
	FamName    []*AccessTermWithRole `xml:"famname" json:"famname,omitempty"`
	Function   []*FilteredString     `xml:"function" json:"function,omitempty"`
	GenreForm  []*FilteredString     `xml:"genreform" json:"genreform,omitempty"`
	GeogName   []*FilteredString     `xml:"geogname" json:"geogname,omitempty"`
	Occupation []*FilteredString     `xml:"occupation" json:"occupation,omitempty"`
	PersName   []*AccessTermWithRole `xml:"persname" json:"persname,omitempty"`
	Subject    []*FilteredString     `xml:"subject" json:"subject,omitempty"`
	Title      []*Title              `xml:"title" json:"title,omitempty"`
}

type Creation struct {
	Value string `xml:",innerxml" json:"value,chardata,omitempty"`
}

type DAO struct {
	Actuate FilteredString `xml:"actuate,attr" json:"actuate,omitempty"`
	Href    FilteredString `xml:"href,attr" json:"href,omitempty"`
	Role    FilteredString `xml:"role,attr" json:"role,omitempty"`
	Show    FilteredString `xml:"show,attr" json:"show,omitempty"`
	DOType  FilteredString `json:"do_type,omitempty"`
	Count   uint64         `json:"count,omitempty"`
	Title   FilteredString `xml:"title,attr" json:"title,omitempty"`
	Type    FilteredString `xml:"type,attr" json:"type,omitempty"`

	DAODesc DAODesc `xml:"daodesc" json:"daodesc,omitempty"`
}

type DAODesc struct {
	P []*P `xml:"p,omitempty" json:"p,omitempty"`
}

type DAOGrp struct {
	Title FilteredString `xml:"title,attr" json:"title,omitempty"`
	Type  FilteredString `xml:"type,attr"  json:"type,omitempty"`

	DAODesc DAODesc   `xml:"daodesc" json:"daodesc,omitempty"`
	DAOLoc  []*DAOLoc `xml:"daoloc" json:"daoloc,omitempty"`
}

type DAOLoc struct {
	Href  FilteredString `xml:"href,attr" json:"href,omitempty"`
	Role  FilteredString `xml:"role,attr" json:"role,omitempty"`
	Title FilteredString `xml:"title,attr" json:"title,omitempty"`
	Type  FilteredString `xml:"type,attr" json:"type,omitempty"`
}

type Date struct {
	Type FilteredString `xml:"type,attr" json:"type,omitempty"`

	Value string `xml:",innerxml" json:"value,chardata,omitempty"`
}

type DefItem struct {
	Item  []*Item             `xml:"item" json:"item,omitempty"`
	Label FilteredLabelString `xml:"label" json:"label,omitempty"`
}

type DID struct {
	Abstract     []*Abstract              `xml:"abstract" json:"abstract,omitempty"`
	Container    []*Container             `xml:"container" json:"container,omitempty"`
	DAO          []*DAO                   `xml:"dao" json:"dao,omitempty"`
	DAOGrp       []*DAOGrp                `xml:"daogrp" json:"daogrp,omitempty"`
	LangMaterial []*LangMaterial          `xml:"langmaterial" json:"langmaterial,omitempty"`
	MaterialSpec []*FormattedNoteWithHead `xml:"materialspec" json:"materialspec,omitempty"`
	Origination  []*Origination           `xml:"origination" json:"origination,omitempty"`
	PhysDesc     []*PhysDesc              `xml:"physdesc" json:"physdesc,omitempty"`
	PhysLoc      []*PhysLoc               `xml:"physloc" json:"physloc,omitempty"`
	Repository   *Repository              `xml:"repository" json:"repository,omitempty"`
	UnitDate     []*UnitDate              `xml:"unitdate" json:"unitdate,omitempty"`
	UnitID       FilteredString           `xml:"unitid" json:"unitid,omitempty"`
	UnitTitle    *UnitTitle               `xml:"unittitle" json:"unittitle,omitempty"`
}

type Dimensions struct {
	ID    FilteredString      `xml:"id,attr" json:"id,omitempty"`
	Label FilteredLabelString `xml:"label,attr" json:"label,omitempty"`

	Value string `xml:",innerxml" json:"value,chardata,omitempty"`
}

type DSC struct {
	C []*C `xml:"c,omitempty" json:"c,omitempty"`
	P []*P `xml:"p,omitempty" json:"p,omitempty"`
}

type EADHeader struct {
	EADID        EADID         `xml:"eadid" json:"eadid,omitempty"`
	FileDesc     FileDesc      `xml:"filedesc" json:"filedesc,omitempty"`
	ProfileDesc  ProfileDesc   `xml:"profiledesc" json:"profiledesc,omitempty"`
	RevisionDesc *RevisionDesc `xml:"revisiondesc" json:"revisiondesc,omitempty"`
}

// NOTE: Event though we are process Value as innerxml, we do not create a
// MarshalJSON for it that processes it as mixed content because we have strict
// validation rules for <eadid> that automatically reject any values that contain
// mixed content.
type EADID struct {
	URL FilteredString `xml:"url,attr" json:"url,omitempty"`

	Value string `xml:",innerxml" json:"value,chardata,omitempty"`
}

type EditionStmt struct {
	P []*P `xml:"p,omitempty" json:"p,omitempty"`
}

type Event struct {
	Title []*Title `xml:"title" json:"title,omitempty"`

	Value string `xml:",innerxml" json:"value,chardata,omitempty"`
}

type EventGrp struct {
	Event []*Event `xml:"event" json:"event,omitempty"`
}

type Extent struct {
	AltRender FilteredString `xml:"altrender,attr" json:"altrender,omitempty"`

	Unit FilteredString `xml:"unit,attr" json:"unit,omitempty"`

	Value string `xml:",innerxml" json:"value,chardata,omitempty"`
}

type ExtPtr struct {
	Href  FilteredString `xml:"href,attr" json:"href,omitempty"`
	Show  FilteredString `xml:"show,attr" json:"show,omitempty"`
	Title FilteredString `xml:"title,attr" json:"title,omitempty"`
	Type  FilteredString `xml:"type,attr" json:"type,omitempty"`
}

type ExtRef struct {
	Actuate FilteredString `xml:"actuate,attr" json:"actuate,omitempty"`
	Href    FilteredString `xml:"href,attr" json:"href,omitempty"`
	Show    FilteredString `xml:"show,attr" json:"show,omitempty"`
	Title   FilteredString `xml:"title,attr" json:"title,omitempty"`
	Type    FilteredString `xml:"type,attr" json:"type,omitempty"`
}

type FileDesc struct {
	EditionStmt     *EditionStmt    `xml:"editionstmt" json:"editionstmt,omitempty"`
	NoteStmt        *NoteStmt       `xml:"notestmt" json:"notestmt,omitempty"`
	PublicationStmt PublicationStmt `xml:"publicationstmt" json:"publicationstmt,omitempty"`
	TitleStmt       *TitleStmt      `xml:"titlestmt" json:"titlestmt,omitempty"`
}

type FormattedNoteWithHead struct {
	ID       FilteredString `xml:"id,attr" json:"id,omitempty"`
	Head     *Head          `xml:"head" json:"head,omitempty"`
	Children []*EADChild    `xml:",any"`

	/*	ChronList   []*ChronList `xml:"chronlist" json:"chronlist,omitempty"`
		DefItem     []*DefItem   `xml:"defitem,omitemtpy" json:"defitem,omitempty"`
		Head        *Head        `xml:"head,omitemtpy" json:"head,omitempty"`
		LegalStatus *LegalStatus `xml:"legalstatus,omitemtpy" json:"legalstatus,omitempty"`
		List        []*List      `xml:"list,omitemtpy" json:"list,omitempty"`
		P           []*P         `xml:"p,omitempty" json:"p,omitempty"`
	*/
}

type Head struct {
	Value string `xml:",innerxml" json:"value,chardata,omitempty"`
}

type Index struct {
	ID FilteredString `xml:"id,attr" json:"id,omitempty"`

	Head       *Head         `xml:"head,omitemtpy" json:"head,omitempty"`
	IndexEntry []*IndexEntry `xml:"indexentry" json:"indexentry,omitempty"`
}

type IndexEntry struct {
	CorpName []*AccessTermWithRole `xml:"corpname" json:"corpname,omitempty"`
	Name     []*AccessTermWithRole `xml:"name" json:"name,omitempty"`
	Subject  []*FilteredString     `xml:"subject" json:"subject,omitempty"`
}

type Item struct {
	BibRef   []*BibRef             `xml:"bibref" json:"bibref,omitempty"`
	CorpName []*AccessTermWithRole `xml:"corpname" json:"corpname,omitempty"`
	Name     []*AccessTermWithRole `xml:"name" json:"name,omitempty"`
	PersName []*AccessTermWithRole `xml:"persname" json:"persname,omitempty"`
	Title    []*Title              `xml:"title" json:"title,omitempty"`

	Value string `xml:",innerxml" json:"value,chardata,omitempty"`
}

type LangMaterial struct {
	ID FilteredString `xml:"id,attr" json:"id,omitempty"`

	Language *FilteredString `xml:"language" json:"language,omitempty"`

	Value string `xml:",innerxml" json:"value,chardata,omitempty"`
}

type LangUsage struct {
	Language FilteredString `xml:"language" json:"language,omitempty"`

	Value string `xml:",innerxml" json:"value,chardata,omitempty"`
}

type LegalStatus struct {
	ID FilteredString `xml:"id,attr" json:"id,omitempty"`

	Value string `xml:",innerxml" json:"value,chardata,omitempty"`
}

type List struct {
	Numeration FilteredString `xml:"numeration,attr" json:"numeration,omitempty"`
	Type       FilteredString `xml:"type,attr"  json:"type,omitempty"`

	Head    *Head      `xml:"head" json:"head,omitempty"`
	Item    []*Item    `xml:"item" json:"item,omitempty"`
	DefItem []*DefItem `xml:"defitem" json:"defitem,omitempty"`
}

type Note struct {
	P []*P `xml:"p" json:"p,omitempty"`
}

type NoteStmt struct {
	Note []*Note `xml:"note" json:"note,omitempty"`
}

type Num struct {
	Type FilteredString `xml:"type,attr" json:"type,omitempty"`

	Value string `xml:",innerxml" json:"value,chardata,omitempty"`
}

type Origination struct {
	Label FilteredLabelString `xml:"label,attr" json:"label,omitempty"`

	CorpName []*AccessTermWithRole `xml:"corpname" json:"corpname,omitempty"`
	FamName  []*AccessTermWithRole `xml:"famname" json:"famname,omitempty"`
	PersName []*AccessTermWithRole `xml:"persname" json:"persname,omitempty"`
}

type P struct {
	Abbr       []*FilteredString     `xml:"abbr" json:"abbr,omitempty"`
	Address    []*Address            `xml:"address" json:"address,omitempty"`
	ArchRef    []*ArchRef            `xml:"archref" json:"archref,omitempty"`
	BibRef     []*BibRef             `xml:"bibref" json:"bibref,omitempty"`
	ChronList  []*ChronList          `xml:"chronlist" json:"chronlist,omitempty"`
	CorpName   []*AccessTermWithRole `xml:"corpname" json:"corpname,omitempty"`
	Date       []*Date               `xml:"date" json:"date,omitempty"`
	ExtRef     []*ExtRef             `xml:"extref" json:"extref,omitempty"`
	GenreForm  []*FilteredString     `xml:"genreform" json:"genreform,omitempty"`
	GeogName   []*FilteredString     `xml:"geogname" json:"geogname,omitempty"`
	List       []*List               `xml:"list" json:"list,omitempty"`
	Name       []*AccessTermWithRole `xml:"name" json:"name,omitempty"`
	Num        []*Num                `xml:"num" json:"num,omitempty"`
	Occupation []*FilteredString     `xml:"occupation" json:"occupation,omitempty"`
	PersName   []*AccessTermWithRole `xml:"persname" json:"persname,omitempty"`
	Subject    []*FilteredString     `xml:"subject" json:"subject,omitempty"`
	Title      []*Title              `xml:"title" json:"title,omitempty"`

	Value string `xml:",innerxml" json:"value,chardata,omitempty"`
}

type PhysDesc struct {
	AltRender FilteredString      `xml:"altrender,attr" json:"altrender,omitempty"`
	ID        FilteredString      `xml:"id,attr" json:"id,omitempty"`
	Label     FilteredLabelString `xml:"label,attr" json:"label,omitempty"`

	Extent     []*Extent   `xml:"extent" json:"extent,omitempty"`
	Dimensions *Dimensions `xml:"dimensions" json:"dimensions,omitempty"`
	PhysFacet  *PhysFacet  `xml:"physfacet" json:"physfacet,omitempty"`

	Value string `xml:",innerxml" json:"value,chardata,omitempty"`
}

type PhysFacet struct {
	ID    FilteredString      `xml:"id,attr" json:"id,omitempty"`
	Label FilteredLabelString `xml:"label,attr" json:"label,omitempty"`

	Value string `xml:",innerxml" json:"value,chardata,omitempty"`
}

type PhysLoc struct {
	ID FilteredString `xml:"id,attr" json:"id,omitempty"`

	Value string `xml:",innerxml" json:"value,chardata,omitempty"`
}

type ProfileDesc struct {
	Creation  *Creation      `xml:"creation" json:"creation,omitempty"`
	DescRules FilteredString `xml:"descrules" json:"descrules,omitempty"`
	LangUsage *LangUsage     `xml:"langusage" json:"langusage,omitempty"`
}

type PublicationStmt struct {
	Address   []*Address     `xml:"address" json:"address,omitempty"`
	P         []*P           `xml:"p" json:"p,omitempty"`
	Publisher FilteredString `xml:"publisher" json:"publisher,omitempty"`
}

type Repository struct {
	CorpName []*AccessTermWithRole `xml:"corpname" json:"corpname,omitempty"`

	Value string `xml:",innerxml" json:"value,chardata,omitempty"`
}

type RevisionDesc struct {
	Change []*Change `xml:"change" json:"change,omitempty"`
}

type Title struct {
	Render FilteredString `xml:"render,attr" json:"render,omitempty"`
	Source FilteredString `xml:"source,attr" json:"source,omitempty"`
	Type   FilteredString `xml:"type,attr" json:"type,omitempty"`

	Value string `xml:",innerxml" json:"value,chardata,omitempty"`
}

type TitleProper struct {
	Type FilteredString `xml:"type,attr" json:"type,omitempty"`

	Num []*Num `xml:"num" json:"num,omitempty"`

	Value string `xml:",innerxml" json:"value,chardata,omitempty"`
}

type TitleStmt struct {
	Author               FilteredString `xml:"author" json:"author,omitempty"`
	Sponsor              FilteredString `xml:"sponsor" json:"sponsor,omitempty"`
	SubTitle             FilteredString `xml:"subtitle" json:"subtitle,omitempty"`
	TitleProper          []*TitleProper `xml:"titleproper" json:"-"`
	FlattenedTitleProper FilteredString `xml:"-" json:"titleproper,omitempty"`
}

type UnitDate struct {
	Type FilteredString `xml:"type,attr" json:"type,omitempty"`

	Value string `xml:",innerxml" json:"value,chardata,omitempty"`
}

type UnitTitle struct {
	CorpName []*AccessTermWithRole `xml:"corpname" json:"corpname,omitempty"`
	Name     []*AccessTermWithRole `xml:"name" json:"name,omitempty"`
	PersName []*AccessTermWithRole `xml:"persname" json:"persname,omitempty"`
	Title    []*Title              `xml:"title" json:"title,omitempty"`

	Value string `xml:",innerxml" json:"value,chardata,omitempty"`
}
