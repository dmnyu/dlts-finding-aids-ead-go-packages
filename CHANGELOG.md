# Changelog

#### v0.10.0
  - add Donors member to EAD struct

#### v0.9.0
  - flatten TitleProper from an array into a single FilteredString

#### v0.8.0
  - add PubInfo type
  - add PubInfo member to EAD struct

#### v0.7.0
  - remove bracketed text from Label values
  - remove leading/trailing spaces from FilteredString values

#### v0.6.0
  - change Address in PublicationStmt to an array
  - change ChronList in FormattedNoteWithHead to an array
  - change ControlAccess in ArchDesc to an array
  - change CorpName in IndexEntry to an array
  - change CorpName in Origination to an array
  - change CorpName in Repository to an array
  - change Date in Change to an array
  - change Date in ChronItem to an array
  - change FamName in Origination to an array
  - change Item in Change to an array
  - change Item in DefItem to an array
  - change Name in IndexEntry to an array
  - change PersName in Origination to an array
  - change PhysLoc in ArchRef to an array
  - change Subject in IndexEntry to an array
  - change Title in Event to an array

#### v0.5.0
  - replace all instances of \r, \t, \n, and consecutive spaces in
    EAD element values with a single space

#### v0.4.0
  - add RunInfo.SourceFile to record the source EAD file path

#### v0.3.0
  - add FilteredString type to strip out newlines from string fields
  - add RunInfo type to capture JSON-creation timestamp and EAD package version
  - add P.PersName
  - add P.GeogName
  - add P.ChronList
  - remove P.ID field
  - remove Head.ExtPtr
  - add Extref.Actuate
  - correct parsing tag for Extent.Unit (it is an attr).
  - remove Head from DSC per data model v8.0.1
  - remove ID   from DID per data model v8.0.1
  - add AltFormAvailable to type C
  - rename AltFormatAvailable to AltFormAvailable
  - remove Abtract.Label to reflect updated data model
  - rename `AltFormatAvailable` to `AltFormAvailable`, correct XML tag, JSON tag
  - add `Date` field to `Creation` struct (matches data model v8.0.1)

#### v0.2.0
  - replace instances of `\n` with spaces in `value` fields processed by `_getConvertedTextWithTags`
  
