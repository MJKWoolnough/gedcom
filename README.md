# gedcom
--
    import "github.com/MJKWoolnough/gedcom"

Package gedcom implements a parser to read genealogical data in a standard format

## Usage

```go
var (
	ErrNoHeader  errors.Error = "no header"
	ErrNoRecords errors.Error = "no records"
	ErrNotLevel  errors.Error = "not level token"
	ErrNotTag    errors.Error = "not tag token"
	ErrNotLine   errors.Error = "not line_value token"
)
```
Errors

```go
var (
	ErrRequiredMissing errors.Error = "required tag missing"
	ErrSingleMultiple  errors.Error = "tag was specified more than the one time allowed"
	ErrUnknownTag      errors.Error = "unknown tag"
)
```
Errors

```go
var (
	ErrInvalidLevel   errors.Error = "invalid level num"
	ErrMissingDelim   errors.Error = "missing delminitator"
	ErrInvalidPointer errors.Error = "invalid pointer string"
	ErrInvalidTag     errors.Error = "invalid tag"
	ErrBadEscape      errors.Error = "bad escape sequence"
	ErrBadChar        errors.Error = "bad character"
)
```
Errors

#### func  AllowInvalidChars

```go
func AllowInvalidChars(o *options)
```
AllowInvalidChars allows control characters (<32) in line_values

#### func  AllowInvalidEscape

```go
func AllowInvalidEscape(o *options)
```
AllowInvalidEscape turns off error reporting for invalid escape sequences

#### func  AllowMissingRequired

```go
func AllowMissingRequired(o *options)
```
AllowMissingRequired turns off error reporting for a missing required field

#### func  AllowMoreThanAllowed

```go
func AllowMoreThanAllowed(o *options)
```
AllowMoreThanAllowed turns off error reporting when more than the maxumum number
of a certain item are read

#### func  AllowTerminatorsInValue

```go
func AllowTerminatorsInValue(o *options)
```
AllowTerminatorsInValue allows a line_value to contain newlines

#### func  AllowUnknownCharset

```go
func AllowUnknownCharset(o *options)
```
AllowUnknownCharset allows the parser to read any character, not just those
within ANSEL

#### func  AllowUnknownTags

```go
func AllowUnknownTags(o *options)
```
AllowUnknownTags turns off error reporting when an unknown tag is read

#### func  AllowWrongLength

```go
func AllowWrongLength(o *options)
```
AllowWrongLength turns off error reporting when a base type has an incorrect
length

#### func  IgnoreInvalidValue

```go
func IgnoreInvalidValue(o *options)
```
IgnoreInvalidValue turns off error reporting when a non-valid value if read

#### type AddressCity

```go
type AddressCity string
```

AddressCity is a GEDCOM base type

#### type AddressCountry

```go
type AddressCountry string
```

AddressCountry is a GEDCOM base type

#### type AddressLine

```go
type AddressLine string
```

AddressLine is a GEDCOM base type

#### type AddressLine1

```go
type AddressLine1 string
```

AddressLine1 is a GEDCOM base type

#### type AddressLine2

```go
type AddressLine2 string
```

AddressLine2 is a GEDCOM base type

#### type AddressPostalCode

```go
type AddressPostalCode string
```

AddressPostalCode is a GEDCOM base type

#### type AddressState

```go
type AddressState string
```

AddressState is a GEDCOM base type

#### type AddressStructure

```go
type AddressStructure struct {
	AddressLine  AddressLine
	AddressLine1 AddressLine1
	AddressLine2 AddressLine2
	City         AddressCity
	State        AddressState
	PostalCode   AddressPostalCode
	Country      AddressCountry
}
```

AddressStructure is a GEDCOM structure type

#### type AdoptedBy

```go
type AdoptedBy string
```

AdoptedBy is a GEDCOM base type

#### type AdoptionEvent

```go
type AdoptionEvent struct {
	Family AdoptionReference
	EventDetail
}
```

AdoptionEvent is a GEDCOM structure type

#### type AdoptionReference

```go
type AdoptionReference struct {
	ID        Xref
	AdoptedBy AdoptedBy
}
```

AdoptionReference is a GEDCOM structure type

#### type AgeAtEvent

```go
type AgeAtEvent string
```

AgeAtEvent is a GEDCOM base type

#### type AgeStructure

```go
type AgeStructure struct {
	Age AgeAtEvent
}
```

AgeStructure is a GEDCOM structure type

#### type AncestralFileNumber

```go
type AncestralFileNumber string
```

AncestralFileNumber is a GEDCOM base type

#### type ApprovedSystemID

```go
type ApprovedSystemID string
```

ApprovedSystemID is a GEDCOM base type

#### type AssociationStructure

```go
type AssociationStructure struct {
	ID         Xref
	RecordType RecordType
	Relation   RelationIsDescriptor
	Notes      []NoteStructure
	Sources    []SourceCitation
}
```

AssociationStructure is a GEDCOM structure type

#### type AttributeType

```go
type AttributeType string
```

AttributeType is a GEDCOM base type

#### type AutomatedRecordID

```go
type AutomatedRecordID string
```

AutomatedRecordID is a GEDCOM base type

#### type CasteEvent

```go
type CasteEvent struct {
	CasteName CasteName
	EventDetail
}
```

CasteEvent is a GEDCOM structure type

#### type CasteName

```go
type CasteName string
```

CasteName is a GEDCOM base type

#### type CauseOfEvent

```go
type CauseOfEvent string
```

CauseOfEvent is a GEDCOM base type

#### type CertaintyAssessment

```go
type CertaintyAssessment uint
```

CertaintyAssessment is a GEDCOM base type

#### type ChangeDate

```go
type ChangeDate string
```

ChangeDate is a GEDCOM base type

#### type ChangeDateStructure

```go
type ChangeDateStructure struct {
	Date  ChangeDateTime
	Notes []NoteStructure
}
```

ChangeDateStructure is a GEDCOM structure type

#### type ChangeDateTime

```go
type ChangeDateTime struct {
	ChangeDate ChangeDate
	Time       TimeValue
}
```

ChangeDateTime is a GEDCOM structure type

#### type CharacterSet

```go
type CharacterSet string
```

CharacterSet is a GEDCOM base type

#### type CharacterSetStructure

```go
type CharacterSetStructure struct {
	CharacterSet  CharacterSet
	VersionNumber VersionNumber
}
```

CharacterSetStructure is a GEDCOM structure type

#### type ChildToFamilyLink

```go
type ChildToFamilyLink struct {
	ID                  Xref
	PedigreeLinkageType []PedigreeLinkageType
	Notes               []NoteStructure
}
```

ChildToFamilyLink is a GEDCOM structure type

#### type ChildrenEvent

```go
type ChildrenEvent struct {
	CountOfChildren CountOfChildren
	EventDetail
}
```

ChildrenEvent is a GEDCOM structure type

#### type ContentDescription

```go
type ContentDescription string
```

ContentDescription is a GEDCOM base type

#### type CopyrightGedcomFile

```go
type CopyrightGedcomFile string
```

CopyrightGedcomFile is a GEDCOM base type

#### type CopyrightSourceData

```go
type CopyrightSourceData string
```

CopyrightSourceData is a GEDCOM base type

#### type CountOfChildren

```go
type CountOfChildren uint8
```

CountOfChildren is a GEDCOM base type

#### type CountOfMarriages

```go
type CountOfMarriages uint8
```

CountOfMarriages is a GEDCOM base type

#### type Date

```go
type Date string
```

Date is a GEDCOM base type

#### type DateApproximated

```go
type DateApproximated string
```

DateApproximated is a GEDCOM base type

#### type DateCalendar

```go
type DateCalendar string
```

DateCalendar is a GEDCOM base type

#### type DateCalendarEscape

```go
type DateCalendarEscape string
```

DateCalendarEscape is a GEDCOM base type

#### type DateExact

```go
type DateExact string
```

DateExact is a GEDCOM base type

#### type DateFren

```go
type DateFren string
```

DateFren is a GEDCOM base type

#### type DateGreg

```go
type DateGreg string
```

DateGreg is a GEDCOM base type

#### type DateHebr

```go
type DateHebr string
```

DateHebr is a GEDCOM base type

#### type DateJuln

```go
type DateJuln string
```

DateJuln is a GEDCOM base type

#### type DateLDSOrd

```go
type DateLDSOrd string
```

DateLDSOrd is a GEDCOM base type

#### type DatePeriod

```go
type DatePeriod string
```

DatePeriod is a GEDCOM base type

#### type DatePhrase

```go
type DatePhrase string
```

DatePhrase is a GEDCOM base type

#### type DateRange

```go
type DateRange string
```

DateRange is a GEDCOM base type

#### type DateValue

```go
type DateValue string
```

DateValue is a GEDCOM base type

#### type Day

```go
type Day uint8
```

Day is a GEDCOM base type

#### type DescriptionEvent

```go
type DescriptionEvent struct {
	PhysicalDescription PhysicalDescription
	EventDetail
}
```

DescriptionEvent is a GEDCOM structure type

#### type DescriptiveTitle

```go
type DescriptiveTitle string
```

DescriptiveTitle is a GEDCOM base type

#### type Digit

```go
type Digit uint8
```

Digit is a GEDCOM base type

#### type EncodedMultimediaLine

```go
type EncodedMultimediaLine string
```

EncodedMultimediaLine is a GEDCOM base type

#### type EntryRecordingDate

```go
type EntryRecordingDate string
```

EntryRecordingDate is a GEDCOM base type

#### type ErrContext

```go
type ErrContext struct {
	Structure, Tag string
	Err            error
}
```

ErrContext adds context to a returned error

#### func (ErrContext) Error

```go
func (e ErrContext) Error() string
```
Error implements the error interface

#### func (ErrContext) Underlying

```go
func (e ErrContext) Underlying() error
```
Underlying goes through the error list to retrieve the underlying
(non-ErrContext) error

#### type ErrInvalidLength

```go
type ErrInvalidLength struct {
	Type, Value string
	Min, Max    uint
}
```

ErrInvalidLength is an error that is generated when a type is given more or less
data than is required

#### func (ErrInvalidLength) Error

```go
func (e ErrInvalidLength) Error() string
```
Error is an implementation of the error interface

#### type ErrInvalidValue

```go
type ErrInvalidValue struct {
	Type, Value string
}
```

ErrInvalidValue is an error that is generated when a type is not one of the
specified values

#### func (ErrInvalidValue) Error

```go
func (e ErrInvalidValue) Error() string
```
Error is an implementation of the error interface

#### type ErrTooMany

```go
type ErrTooMany int
```

ErrTooMany is an error returned when too many of a particular tag exist

#### func (ErrTooMany) Error

```go
func (ErrTooMany) Error() string
```
Error implements the error interface

#### type EventAttributeType

```go
type EventAttributeType string
```

EventAttributeType is a GEDCOM base type

#### type EventDescriptor

```go
type EventDescriptor string
```

EventDescriptor is a GEDCOM base type

#### type EventDetail

```go
type EventDetail struct {
	Type              EventDescriptor
	Date              DateValue
	Place             PlaceStructure
	Address           AddressStructure
	PhoneNumber       []PhoneNumber
	Age               AgeAtEvent
	ResponsibleAgency ResponsibleAgency
	CauseOfEvent      CauseOfEvent
	Sources           []SourceCitation
	Multimedia        []MultimediaLink
	Notes             []NoteStructure
}
```

EventDetail is a GEDCOM structure type

#### type EventTypeCitedFrom

```go
type EventTypeCitedFrom string
```

EventTypeCitedFrom is a GEDCOM base type

#### type EventTypeFamile

```go
type EventTypeFamile string
```

EventTypeFamile is a GEDCOM base type

#### type EventTypeIndividual

```go
type EventTypeIndividual string
```

EventTypeIndividual is a GEDCOM base type

#### type EventsRecorded

```go
type EventsRecorded string
```

EventsRecorded is a GEDCOM base type

#### type EventsRecordedStructure

```go
type EventsRecordedStructure struct {
	DatePeriod              DatePeriod
	SourceJurisdictionPlace SourceJurisdictionPlace
}
```

EventsRecordedStructure is a GEDCOM structure type

#### type Family

```go
type Family struct {
	ID                 Xref
	Annulment          VerifiedFamilyEventDetail
	Census             VerifiedFamilyEventDetail
	Divorce            VerifiedFamilyEventDetail
	DivorceFiled       VerifiedFamilyEventDetail
	Engagement         VerifiedFamilyEventDetail
	Marriage           VerifiedFamilyEventDetail
	MarriageBann       VerifiedFamilyEventDetail
	MarriageContract   VerifiedFamilyEventDetail
	MarriageLicense    VerifiedFamilyEventDetail
	MarriageSettlement VerifiedFamilyEventDetail
	Events             []FamilyEventDetail
	Husband            Xref
	Wife               Xref
	Children           []Xref
	NumChildren        CountOfChildren
	Submitters         []Xref
	LDSSpouseSealing   []LDSSpouseSealing
	Sources            []SourceCitation
	Multimedia         []MultimediaLink
	Notes              []NoteStructure
}
```

Family is a GEDCOM structure type

#### func (Family) Type

```go
func (Family) Type() RecordID
```
Type implements the Record interface

#### type FamilyEventDetail

```go
type FamilyEventDetail struct {
	HusbandAge AgeStructure
	WifeAge    AgeStructure
	EventDetail
}
```

FamilyEventDetail is a GEDCOM structure type

#### type FileName

```go
type FileName string
```

FileName is a GEDCOM base type

#### type Form

```go
type Form string
```

Form is a GEDCOM base type

#### type GenerationsOfAncestors

```go
type GenerationsOfAncestors uint16
```

GenerationsOfAncestors is a GEDCOM base type

#### type GenerationsOfDescendants

```go
type GenerationsOfDescendants uint16
```

GenerationsOfDescendants is a GEDCOM base type

#### type Header

```go
type Header struct {
	Source              HeaderSource
	ReceivingSystemName ReceivingSystemName
	TransmissionLDate   TransmissionDateTime
	Submitter           Xref
	Submission          Xref
	FileName            FileName
	Copyright           CopyrightGedcomFile
	Version             Version
	CharacterSet        CharacterSetStructure
	Language            LanguageOfText
	Place               HeaderPlace
	ContentDescription  ContentDescription
}
```

Header is a GEDCOM structure type

#### func (Header) Type

```go
func (Header) Type() RecordID
```
Type implements the Record interface

#### type HeaderBusiness

```go
type HeaderBusiness struct {
	NameOfBusiness NameOfBusiness
	Address        AddressStructure
	PhoneNumber    []PhoneNumber
}
```

HeaderBusiness is a GEDCOM structure type

#### type HeaderDataSource

```go
type HeaderDataSource struct {
	SourceName          NameOfSourceData
	PublicationDate     PublicationDate
	CopyrightSourceData CopyrightSourceData
}
```

HeaderDataSource is a GEDCOM structure type

#### type HeaderPlace

```go
type HeaderPlace struct {
	PlaceHierarchy PlaceHierarchy
}
```

HeaderPlace is a GEDCOM structure type

#### type HeaderSource

```go
type HeaderSource struct {
	SystemID      ApprovedSystemID
	VersionNumber VersionNumber
	Name          NameOfProduct
	Business      HeaderBusiness
	Data          HeaderDataSource
}
```

HeaderSource is a GEDCOM structure type

#### type Individual

```go
type Individual struct {
	ID                    Xref
	RestrictionNotice     RestrictionNotice
	PersonalNameStructure []PersonalNameStructure
	Gender                SexValue
	Birth                 VerifiedIndividualFamEventDetail
	Christening           VerifiedIndividualFamEventDetail
	Death                 VerifiedEventDetail
	Buried                VerifiedEventDetail
	Cremation             VerifiedEventDetail
	Adoption              AdoptionEvent
	Maptism               VerifiedEventDetail
	BarMitzvah            VerifiedEventDetail
	BasMitzvah            VerifiedEventDetail
	Blessing              VerifiedEventDetail
	AdultChristening      VerifiedEventDetail
	Confirmation          VerifiedEventDetail
	FirstCommunion        VerifiedEventDetail
	Ordination            VerifiedEventDetail
	Naturalization        VerifiedEventDetail
	Emigrated             VerifiedEventDetail
	Immigrated            VerifiedEventDetail
	Census                VerifiedEventDetail
	Probate               VerifiedEventDetail
	Will                  VerifiedEventDetail
	Graduated             VerifiedEventDetail
	Retired               VerifiedEventDetail
	Events                []IndividualEventDetail
	Caste                 []CasteEvent
	Description           []DescriptionEvent
	ScholasticAchievement []ScholasticEvent
	NationalID            []NationalIDEvent
	NationalTribalOrigin  []NationalOriginEvent
	CountOfChildren       []ChildrenEvent
	CountOfMarriages      []MarriagesEvent
	Occupation            []OccupationEvent
	Possessions           []PossessionEvent
	ReligiousAffiliation  []ReligiousEvent
	Residences            []ResidenceEvent
	SocialSecurity        []SSNEvent
	NobilityTypeTitle     []NobilityEvent
	ChildOf               []ChildToFamilyLink
	SpouseOf              []SpouseToFamilyLink
	Submitters            []Xref
	Associations          []AssociationStructure
	Aliases               []Xref
	AncestorInterest      []Xref
	DescendentInterest    []Xref
	Sources               []SourceCitation
	Multimedia            []MultimediaLink
	Notes                 []NoteStructure
	PermanentRecord       PermanentRecordFileNumber
	AncestralFileNumber   AncestralFileNumber
	UserReferences        []UserReferenceStructure
	AutomatedRecordID     AutomatedRecordID
	ChangeDate            ChangeDateStructure
}
```

Individual is a GEDCOM structure type

#### func (Individual) Type

```go
func (Individual) Type() RecordID
```
Type implements the Record interface

#### type IndividualEventDetail

```go
type IndividualEventDetail struct {
	EventDetail
}
```

IndividualEventDetail is a GEDCOM structure type

#### type LDSBaptismDateStatus

```go
type LDSBaptismDateStatus string
```

LDSBaptismDateStatus is a GEDCOM base type

#### type LDSChildSealingDateStatus

```go
type LDSChildSealingDateStatus string
```

LDSChildSealingDateStatus is a GEDCOM base type

#### type LDSEndowmentDateStatus

```go
type LDSEndowmentDateStatus string
```

LDSEndowmentDateStatus is a GEDCOM base type

#### type LDSSpouseSealing

```go
type LDSSpouseSealing struct {
	LDSSpouseSealingDateStatus LDSSpouseSealingDateStatus
	Date                       DateLDSOrd
	TempleCode                 TempleCode
	Place                      PlaceLivingOrdinance
	Sources                    []SourceCitation
	Notes                      []NoteStructure
}
```

LDSSpouseSealing is a GEDCOM structure type

#### type LDSSpouseSealingDateStatus

```go
type LDSSpouseSealingDateStatus string
```

LDSSpouseSealingDateStatus is a GEDCOM base type

#### type LanguageID

```go
type LanguageID string
```

LanguageID is a GEDCOM base type

#### type LanguageOfText

```go
type LanguageOfText string
```

LanguageOfText is a GEDCOM base type

#### type LanguagePreference

```go
type LanguagePreference string
```

LanguagePreference is a GEDCOM base type

#### type Line

```go
type Line struct {
	Sub []Line
}
```

Line represents an unknown GEDCOM record

#### func (Line) Type

```go
func (Line) Type() RecordID
```
Type implements the Record interface

#### type MarriagesEvent

```go
type MarriagesEvent struct {
	CountOfMarriages CountOfMarriages
	EventDetail
}
```

MarriagesEvent is a GEDCOM structure type

#### type Month

```go
type Month string
```

Month is a GEDCOM base type

#### type MonthFren

```go
type MonthFren string
```

MonthFren is a GEDCOM base type

#### type MonthHebr

```go
type MonthHebr string
```

MonthHebr is a GEDCOM base type

#### type MultimediaFileReference

```go
type MultimediaFileReference string
```

MultimediaFileReference is a GEDCOM base type

#### type MultimediaFormat

```go
type MultimediaFormat string
```

MultimediaFormat is a GEDCOM base type

#### type MultimediaLink

```go
type MultimediaLink struct {
	Data Record
}
```

MultimediaLink splite between MultimediaLinkID and MultimediaLinkFile

#### type MultimediaLinkFile

```go
type MultimediaLinkFile struct {
	Format MultimediaFormat
	Title  DescriptiveTitle
	File   MultimediaFileReference
	Notes  []NoteStructure
}
```

MultimediaLinkFile is a GEDCOM structure type

#### func (MultimediaLinkFile) Type

```go
func (MultimediaLinkFile) Type() RecordID
```
Type implements the Record interface

#### type MultimediaLinkID

```go
type MultimediaLinkID struct {
	ID Xref
}
```

MultimediaLinkID is a GEDCOM structure type

#### func (MultimediaLinkID) Type

```go
func (MultimediaLinkID) Type() RecordID
```
Type implements the Record interface

#### type MultimediaRecord

```go
type MultimediaRecord struct {
	ID                Xref
	Format            MultimediaFormat
	Title             DescriptiveTitle
	Notes             []NoteStructure
	Blob              EncodedMultimediaLine
	ContinuedObject   Xref
	UserReferences    []UserReferenceStructure
	AutomatedRecordID AutomatedRecordID
	ChangeDate        ChangeDateStructure
}
```

MultimediaRecord is a GEDCOM structure type

#### func (MultimediaRecord) Type

```go
func (MultimediaRecord) Type() RecordID
```
Type implements the Record interface

#### type NameOfBusiness

```go
type NameOfBusiness string
```

NameOfBusiness is a GEDCOM base type

#### type NameOfFamilyFile

```go
type NameOfFamilyFile string
```

NameOfFamilyFile is a GEDCOM base type

#### type NameOfProduct

```go
type NameOfProduct string
```

NameOfProduct is a GEDCOM base type

#### type NameOfRepository

```go
type NameOfRepository string
```

NameOfRepository is a GEDCOM base type

#### type NameOfSourceData

```go
type NameOfSourceData string
```

NameOfSourceData is a GEDCOM base type

#### type NamePersonal

```go
type NamePersonal string
```

NamePersonal is a GEDCOM base type

#### type NamePiece

```go
type NamePiece string
```

NamePiece is a GEDCOM base type

#### type NamePieceGiven

```go
type NamePieceGiven string
```

NamePieceGiven is a GEDCOM base type

#### type NamePieceNickname

```go
type NamePieceNickname string
```

NamePieceNickname is a GEDCOM base type

#### type NamePiecePrefix

```go
type NamePiecePrefix string
```

NamePiecePrefix is a GEDCOM base type

#### type NamePieceSuffix

```go
type NamePieceSuffix string
```

NamePieceSuffix is a GEDCOM base type

#### type NamePieceSurname

```go
type NamePieceSurname string
```

NamePieceSurname is a GEDCOM base type

#### type NamePieceSurnamePrefix

```go
type NamePieceSurnamePrefix string
```

NamePieceSurnamePrefix is a GEDCOM base type

#### type NationalIDEvent

```go
type NationalIDEvent struct {
	NationalIDNumber NationalIDNumber
	EventDetail
}
```

NationalIDEvent is a GEDCOM structure type

#### type NationalIDNumber

```go
type NationalIDNumber string
```

NationalIDNumber is a GEDCOM base type

#### type NationalOrTribalOrigin

```go
type NationalOrTribalOrigin string
```

NationalOrTribalOrigin is a GEDCOM base type

#### type NationalOriginEvent

```go
type NationalOriginEvent struct {
	NationalOrTribalOrigin NationalOrTribalOrigin
	EventDetail
}
```

NationalOriginEvent is a GEDCOM structure type

#### type NewTag

```go
type NewTag string
```

NewTag is a GEDCOM base type

#### type NobilityEvent

```go
type NobilityEvent struct {
	NobilityTypeTitle NobilityTypeTitle
	EventDetail
}
```

NobilityEvent is a GEDCOM structure type

#### type NobilityTypeTitle

```go
type NobilityTypeTitle string
```

NobilityTypeTitle is a GEDCOM base type

#### type NoteID

```go
type NoteID struct {
	ID Xref
}
```

NoteID is a GEDCOM structure type

#### func (NoteID) Type

```go
func (NoteID) Type() RecordID
```
Type implements the Record interface

#### type NoteRecord

```go
type NoteRecord struct {
	ID                Xref
	SubmitterText     SubmitterText
	Sources           []SourceCitation
	UserReferences    []UserReferenceStructure
	AutomatedRecordID AutomatedRecordID
	ChangeDate        ChangeDateStructure
}
```

NoteRecord is a GEDCOM structure type

#### func (NoteRecord) Type

```go
func (NoteRecord) Type() RecordID
```
Type implements the Record interface

#### type NoteStructure

```go
type NoteStructure struct {
	Data Record
}
```

NoteStructure splits between NoteID and NoteText

#### type NoteText

```go
type NoteText struct {
	SubmitterText SubmitterText
	Sources       []SourceCitation
}
```

NoteText is a GEDCOM structure type

#### func (NoteText) Type

```go
func (NoteText) Type() RecordID
```
Type implements the Record interface

#### type Number

```go
type Number uint
```

Number is a GEDCOM base type

#### type Occupation

```go
type Occupation string
```

Occupation is a GEDCOM base type

#### type OccupationEvent

```go
type OccupationEvent struct {
	Occupation Occupation
	EventDetail
}
```

OccupationEvent is a GEDCOM structure type

#### type Option

```go
type Option func(o *options)
```

Option type for turning off generated errors

#### type OrdinanceProcessFlag

```go
type OrdinanceProcessFlag string
```

OrdinanceProcessFlag is a GEDCOM base type

#### type PedigreeLinkageType

```go
type PedigreeLinkageType string
```

PedigreeLinkageType is a GEDCOM base type

#### type PermanentRecordFileNumber

```go
type PermanentRecordFileNumber string
```

PermanentRecordFileNumber is a GEDCOM base type

#### type PersonalNameStructure

```go
type PersonalNameStructure struct {
	NamePersonal  NamePersonal
	Prefix        NamePiecePrefix
	Given         NamePieceGiven
	Nickname      NamePieceNickname
	SurnamePrefix NamePieceSurnamePrefix
	Surname       NamePieceSurname
	Suffix        NamePieceSuffix
	Sources       []SourceCitation
	Notes         []NoteStructure
}
```

PersonalNameStructure is a GEDCOM structure type

#### type PhoneNumber

```go
type PhoneNumber string
```

PhoneNumber is a GEDCOM base type

#### type PhysicalDescription

```go
type PhysicalDescription string
```

PhysicalDescription is a GEDCOM base type

#### type PlaceHierarchy

```go
type PlaceHierarchy string
```

PlaceHierarchy is a GEDCOM base type

#### type PlaceLivingOrdinance

```go
type PlaceLivingOrdinance string
```

PlaceLivingOrdinance is a GEDCOM base type

#### type PlaceStructure

```go
type PlaceStructure struct {
	Place          PlaceValue
	PlaceHierarchy PlaceHierarchy
	Sources        []SourceCitation
	Notes          []NoteStructure
}
```

PlaceStructure is a GEDCOM structure type

#### type PlaceValue

```go
type PlaceValue string
```

PlaceValue is a GEDCOM base type

#### type PossessionEvent

```go
type PossessionEvent struct {
	Possessions Possessions
	EventDetail
}
```

PossessionEvent is a GEDCOM structure type

#### type Possessions

```go
type Possessions string
```

Possessions is a GEDCOM base type

#### type PublicationDate

```go
type PublicationDate string
```

PublicationDate is a GEDCOM base type

#### type Reader

```go
type Reader struct {
}
```

Reader reads Records from the underlying GEDCOM file

#### func  NewReader

```go
func NewReader(r io.Reader, opts ...Option) *Reader
```
NewReader creates a new Reader, setting the given options

#### func (*Reader) Record

```go
func (r *Reader) Record() (Record, error)
```
Record returns a GEDCOM Record. Record types are: -

    *Header
    *Submission
    *Family
    *Invididual
    *MultimediaNote
    *Repository
    *Source
    *Submitter
    *Trailer

#### type ReceivingSystemName

```go
type ReceivingSystemName string
```

ReceivingSystemName is a GEDCOM base type

#### type Record

```go
type Record interface {
	Type() RecordID
	// contains filtered or unexported methods
}
```

Record is the interface to contain all of the record types

#### type RecordID

```go
type RecordID uint
```

RecordID stores the type of record

```go
const (
	RecordHeader RecordID = iota
	RecordSubmission
	RecordFamily
	RecordIndividual
	RecordMultimedia
	RecordNote
	RecordRepository
	RecordSource
	RecordSubmitter
	RecordTrailer
	RecordUnknown
	RecordMultimediaLinkID
	RecordMultimediaLinkFile
	RecordNoteID
	RecordNoteText
	RecordSourceID
	RecordSourceText
)
```
RecordIDs

#### type RecordIdentifier

```go
type RecordIdentifier string
```

RecordIdentifier is a GEDCOM base type

#### type RecordType

```go
type RecordType string
```

RecordType is a GEDCOM base type

#### type RegisteredResourceIdentifier

```go
type RegisteredResourceIdentifier string
```

RegisteredResourceIdentifier is a GEDCOM base type

#### type RelationIsDescriptor

```go
type RelationIsDescriptor string
```

RelationIsDescriptor is a GEDCOM base type

#### type ReligiousAffiliation

```go
type ReligiousAffiliation string
```

ReligiousAffiliation is a GEDCOM base type

#### type ReligiousEvent

```go
type ReligiousEvent struct {
	ReligiousAffiliation ReligiousAffiliation
	EventDetail
}
```

ReligiousEvent is a GEDCOM structure type

#### type RepositoryRecord

```go
type RepositoryRecord struct {
	ID                Xref
	NameOfRepository  NameOfRepository
	Address           AddressStructure
	PhoneNumber       []PhoneNumber
	Notes             []NoteStructure
	UserReferences    []UserReferenceStructure
	AutomatedRecordID AutomatedRecordID
	ChangeDate        ChangeDateStructure
}
```

RepositoryRecord is a GEDCOM structure type

#### func (RepositoryRecord) Type

```go
func (RepositoryRecord) Type() RecordID
```
Type implements the Record interface

#### type ResidenceEvent

```go
type ResidenceEvent struct {
	EventDetail
}
```

ResidenceEvent is a GEDCOM structure type

#### type ResponsibleAgency

```go
type ResponsibleAgency string
```

ResponsibleAgency is a GEDCOM base type

#### type RestrictionNotice

```go
type RestrictionNotice string
```

RestrictionNotice is a GEDCOM base type

#### type RoleDescriptor

```go
type RoleDescriptor string
```

RoleDescriptor is a GEDCOM base type

#### type RoleInEvent

```go
type RoleInEvent string
```

RoleInEvent is a GEDCOM base type

#### type SSNEvent

```go
type SSNEvent struct {
	SocialSecurityNumber SocialSecurityNumber
	EventDetail
}
```

SSNEvent is a GEDCOM structure type

#### type ScholasticAchievement

```go
type ScholasticAchievement string
```

ScholasticAchievement is a GEDCOM base type

#### type ScholasticEvent

```go
type ScholasticEvent struct {
	ScholasticAchievement ScholasticAchievement
	EventDetail
}
```

ScholasticEvent is a GEDCOM structure type

#### type SexValue

```go
type SexValue string
```

SexValue is a GEDCOM base type

#### type SocialSecurityNumber

```go
type SocialSecurityNumber string
```

SocialSecurityNumber is a GEDCOM base type

#### type SourceCallNumber

```go
type SourceCallNumber string
```

SourceCallNumber is a GEDCOM base type

#### type SourceCallStructure

```go
type SourceCallStructure struct {
	SourceCallNumber SourceCallNumber
	Type             SourceMediaType
}
```

SourceCallStructure is a GEDCOM structure type

#### type SourceCitation

```go
type SourceCitation struct {
	Data Record
}
```

SourceCitation splits between SourceID and SourceText

#### type SourceCitationEvent

```go
type SourceCitationEvent struct {
	EventTypeCitedFrom EventTypeCitedFrom
	Role               RoleInEvent
}
```

SourceCitationEvent is a GEDCOM structure type

#### type SourceData

```go
type SourceData struct {
	Date EntryRecordingDate
	Text []TextFromSource
}
```

SourceData is a GEDCOM structure type

#### type SourceDescription

```go
type SourceDescription string
```

SourceDescription is a GEDCOM base type

#### type SourceDescriptiveTitle

```go
type SourceDescriptiveTitle string
```

SourceDescriptiveTitle is a GEDCOM base type

#### type SourceFiledByEntry

```go
type SourceFiledByEntry string
```

SourceFiledByEntry is a GEDCOM base type

#### type SourceID

```go
type SourceID struct {
	ID                  Xref
	WhereWithinSource   WhereWithinSource
	SourceCitationEvent SourceCitationEvent
	Data                SourceData
	Certainty           CertaintyAssessment
	Multimedia          []MultimediaLink
	Notes               []NoteStructure
}
```

SourceID is a GEDCOM structure type

#### func (SourceID) Type

```go
func (SourceID) Type() RecordID
```
Type implements the Record interface

#### type SourceJurisdictionPlace

```go
type SourceJurisdictionPlace string
```

SourceJurisdictionPlace is a GEDCOM base type

#### type SourceMediaType

```go
type SourceMediaType string
```

SourceMediaType is a GEDCOM base type

#### type SourceOriginator

```go
type SourceOriginator string
```

SourceOriginator is a GEDCOM base type

#### type SourcePublicationFacts

```go
type SourcePublicationFacts string
```

SourcePublicationFacts is a GEDCOM base type

#### type SourceRecord

```go
type SourceRecord struct {
	ID                       Xref
	Data                     SourceRecordDataStructure
	Originator               SourceOriginator
	Title                    SourceDescriptiveTitle
	FiledBy                  SourceFiledByEntry
	PublicationFacts         SourcePublicationFacts
	TextFromSource           TextFromSource
	SourceRepositoryCitation SourceRepositoryCitation
	ContinuedObject          Xref
	Notes                    []NoteStructure
	UserReferences           []UserReferenceStructure
	AutomatedRecordID        AutomatedRecordID
	ChangeDate               ChangeDateStructure
}
```

SourceRecord is a GEDCOM structure type

#### func (SourceRecord) Type

```go
func (SourceRecord) Type() RecordID
```
Type implements the Record interface

#### type SourceRecordDataStructure

```go
type SourceRecordDataStructure struct {
	EventsRecorded    []EventsRecordedStructure
	ResponsibleAgency ResponsibleAgency
	Notes             []NoteStructure
}
```

SourceRecordDataStructure is a GEDCOM structure type

#### type SourceRepositoryCitation

```go
type SourceRepositoryCitation struct {
	ID      Xref
	Numbers []SourceCallStructure
}
```

SourceRepositoryCitation is a GEDCOM structure type

#### type SourceText

```go
type SourceText struct {
	Description SourceDescription
	Texts       []TextFromSource
	Notes       []NoteStructure
}
```

SourceText is a GEDCOM structure type

#### func (SourceText) Type

```go
func (SourceText) Type() RecordID
```
Type implements the Record interface

#### type SpouseToFamilyLink

```go
type SpouseToFamilyLink struct {
	ID    Xref
	Notes []NoteStructure
}
```

SpouseToFamilyLink is a GEDCOM structure type

#### type SubmissionRecord

```go
type SubmissionRecord struct {
	ID                       Xref
	NameOfFamilyFile         NameOfFamilyFile
	TempleCode               TempleCode
	GenerationsOfAncestors   GenerationsOfAncestors
	GenerationsOfDescendants GenerationsOfDescendants
	OrdinanceProcessFlag     OrdinanceProcessFlag
	AutomatedRecordID        AutomatedRecordID
}
```

SubmissionRecord is a GEDCOM structure type

#### func (SubmissionRecord) Type

```go
func (SubmissionRecord) Type() RecordID
```
Type implements the Record interface

#### type SubmitterName

```go
type SubmitterName string
```

SubmitterName is a GEDCOM base type

#### type SubmitterRecord

```go
type SubmitterRecord struct {
	ID                     Xref
	SubmitterName          SubmitterName
	Address                AddressStructure
	PhoneNumber            []PhoneNumber
	Multimedia             []MultimediaLink
	LanguagePreference     []LanguagePreference
	SubmitterRegisteredRFN SubmitterRegisteredRFN
	AutomatedRecordID      AutomatedRecordID
	ChangeDate             ChangeDateStructure
}
```

SubmitterRecord is a GEDCOM structure type

#### func (SubmitterRecord) Type

```go
func (SubmitterRecord) Type() RecordID
```
Type implements the Record interface

#### type SubmitterRegisteredRFN

```go
type SubmitterRegisteredRFN string
```

SubmitterRegisteredRFN is a GEDCOM base type

#### type SubmitterText

```go
type SubmitterText string
```

SubmitterText is a GEDCOM base type

#### type TempleCode

```go
type TempleCode string
```

TempleCode is a GEDCOM base type

#### type Text

```go
type Text string
```

Text is a GEDCOM base type

#### type TextFromSource

```go
type TextFromSource string
```

TextFromSource is a GEDCOM base type

#### type TimeValue

```go
type TimeValue string
```

TimeValue is a GEDCOM base type

#### type Trailer

```go
type Trailer struct{}
```

Trailer type

#### func (Trailer) Type

```go
func (Trailer) Type() RecordID
```
Type implements the Record interface

#### type TransmissionDate

```go
type TransmissionDate string
```

TransmissionDate is a GEDCOM base type

#### type TransmissionDateTime

```go
type TransmissionDateTime struct {
	TransmissionDate TransmissionDate
	Time             TimeValue
}
```

TransmissionDateTime is a GEDCOM structure type

#### type UserReferenceNumber

```go
type UserReferenceNumber string
```

UserReferenceNumber is a GEDCOM base type

#### type UserReferenceStructure

```go
type UserReferenceStructure struct {
	UserReferenceNumber UserReferenceNumber
	Type                UserReferenceType
}
```

UserReferenceStructure is a GEDCOM structure type

#### type UserReferenceType

```go
type UserReferenceType string
```

UserReferenceType is a GEDCOM base type

#### type Verified

```go
type Verified string
```

Verified is a GEDCOM base type

#### type VerifiedEventDetail

```go
type VerifiedEventDetail struct {
	Verified Verified
	EventDetail
}
```

VerifiedEventDetail is a GEDCOM structure type

#### type VerifiedFamilyEventDetail

```go
type VerifiedFamilyEventDetail struct {
	Verified Verified
	FamilyEventDetail
}
```

VerifiedFamilyEventDetail is a GEDCOM structure type

#### type VerifiedIndividualFamEventDetail

```go
type VerifiedIndividualFamEventDetail struct {
	Famc Xref
	VerifiedEventDetail
}
```

VerifiedIndividualFamEventDetail is a GEDCOM structure type

#### type Version

```go
type Version struct {
	VersionNumber VersionNumber
	Form          Form
}
```

Version is a GEDCOM structure type

#### type VersionNumber

```go
type VersionNumber string
```

VersionNumber is a GEDCOM base type

#### type WhereWithinSource

```go
type WhereWithinSource string
```

WhereWithinSource is a GEDCOM base type

#### type Xref

```go
type Xref string
```

Xref is a GEDCOM base type

#### type Year

```go
type Year string
```

Year is a GEDCOM base type

#### type YearGreg

```go
type YearGreg string
```

YearGreg is a GEDCOM base type
