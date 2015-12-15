package gedcom

// File automatically generated with ./genMap.sh

import (
	"strconv"
	"strings"
)

type AddressCity string

func (e *AddressCity) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 60 {
		return ErrInvalidLength{"AddressCity", l.value, 1, 60}
	}
	*e = AddressCity(l.value)
	return nil
}

type AddressCountry string

func (e *AddressCountry) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 60 {
		return ErrInvalidLength{"AddressCountry", l.value, 1, 60}
	}
	*e = AddressCountry(l.value)
	return nil
}

type AddressLine string

func (e *AddressLine) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 60 {
		return ErrInvalidLength{"AddressLine", l.value, 1, 60}
	}
	*e = AddressLine(l.value)
	return nil
}

type AddressLine1 string

func (e *AddressLine1) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 60 {
		return ErrInvalidLength{"AddressLine1", l.value, 1, 60}
	}
	*e = AddressLine1(l.value)
	return nil
}

type AddressLine2 string

func (e *AddressLine2) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 60 {
		return ErrInvalidLength{"AddressLine2", l.value, 1, 60}
	}
	*e = AddressLine2(l.value)
	return nil
}

type AddressPostalCose string

func (e *AddressPostalCose) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 10 {
		return ErrInvalidLength{"AddressPostalCose", l.value, 1, 10}
	}
	*e = AddressPostalCose(l.value)
	return nil
}

type AddressState string

func (e *AddressState) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 60 {
		return ErrInvalidLength{"AddressState", l.value, 1, 60}
	}
	*e = AddressState(l.value)
	return nil
}

type AdoptedBy string

func (e *AdoptedBy) parse(l Line) error {
	switch strings.ToUpper(l.value) {
	case "HUSB":
		*e = "HUSB"
	case "WIFE":
		*e = "WIFE"
	case "BOTH":
		*e = "BOTH"
	default:
		return ErrInvalidValue{"AdoptedBy", l.value}
	}
	return nil
}

type AgeAtEvent string

func (e *AgeAtEvent) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 12 {
		return ErrInvalidLength{"AgeAtEvent", l.value, 1, 12}
	}
	*e = AgeAtEvent(l.value)
	return nil
}

type AncestralFileNumber string

func (e *AncestralFileNumber) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 12 {
		return ErrInvalidLength{"AncestralFileNumber", l.value, 1, 12}
	}
	*e = AncestralFileNumber(l.value)
	return nil
}

type ApprovedSystemID string

func (e *ApprovedSystemID) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 20 {
		return ErrInvalidLength{"ApprovedSystemID", l.value, 1, 20}
	}
	*e = ApprovedSystemID(l.value)
	return nil
}

type AttributeType string

func (e *AttributeType) parse(l Line) error {
	switch strings.ToUpper(l.value) {
	case "CAST":
		*e = "CAST"
	case "EDUC":
		*e = "EDUC"
	case "NATI":
		*e = "NATI"
	case "OCCU":
		*e = "OCCU"
	case "PROP":
		*e = "PROP"
	case "RELI":
		*e = "RELI"
	case "RESI":
		*e = "RESI"
	case "TITL":
		*e = "TITL"
	default:
		return ErrInvalidValue{"AttributeType", l.value}
	}
	return nil
}

type AutomatedRecordID string

func (e *AutomatedRecordID) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 12 {
		return ErrInvalidLength{"AutomatedRecordID", l.value, 1, 12}
	}
	*e = AutomatedRecordID(l.value)
	return nil
}

type CasteName string

func (e *CasteName) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 90 {
		return ErrInvalidLength{"CasteName", l.value, 1, 90}
	}
	*e = CasteName(l.value)
	return nil
}

type CauseOfEvent string

func (e *CauseOfEvent) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 90 {
		return ErrInvalidLength{"CauseOfEvent", l.value, 1, 90}
	}
	*e = CauseOfEvent(l.value)
	return nil
}

type CertaintyAssessment uint

func (e *CertaintyAssessment) parse(l Line) error {
	switch strings.ToUpper(l.value) {
	case "0":
		*e = 0
	case "1":
		*e = 1
	case "2":
		*e = 2
	case "3":
		*e = 3
	default:
		return ErrInvalidValue{"CertaintyAssessment", l.value}
	}
	return nil
}

type ChangeDate string

func (e *ChangeDate) parse(l Line) error {
	if len(l.value) < 10 || len(l.value) > 11 {
		return ErrInvalidLength{"ChangeDate", l.value, 10, 11}
	}
	*e = ChangeDate(l.value)
	return nil
}

type CharacterSet string

func (e *CharacterSet) parse(l Line) error {
	switch strings.ToUpper(l.value) {
	case "ANSWL":
		*e = "ANSWL"
	case "UNICODE":
		*e = "UNICODE"
	case "ASCII":
		*e = "ASCII"
	default:
		return ErrInvalidValue{"CharacterSet", l.value}
	}
	return nil
}

type CopyrightGedcomFile string

func (e *CopyrightGedcomFile) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 90 {
		return ErrInvalidLength{"CopyrightGedcomFile", l.value, 1, 90}
	}
	*e = CopyrightGedcomFile(l.value)
	return nil
}

type CopyrightSourceData string

func (e *CopyrightSourceData) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 90 {
		return ErrInvalidLength{"CopyrightSourceData", l.value, 1, 90}
	}
	*e = CopyrightSourceData(l.value)
	return nil
}

type CountOfChildren uint8

func (e *CountOfChildren) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 3 {
		return ErrInvalidLength{"CountOfChildren", l.value, 1, 3}
	}
	n, err := strconv.ParseUint(l.value, 10, 8)
	if err != nil {
		return err
	}
	*e = CountOfChildren(n)
	return nil
}

type CountOfMarriages uint8

func (e *CountOfMarriages) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 3 {
		return ErrInvalidLength{"CountOfMarriages", l.value, 1, 3}
	}
	n, err := strconv.ParseUint(l.value, 10, 8)
	if err != nil {
		return err
	}
	*e = CountOfMarriages(n)
	return nil
}

type Date string

func (e *Date) parse(l Line) error {
	if len(l.value) < 3 || len(l.value) > 35 {
		return ErrInvalidLength{"Date", l.value, 3, 35}
	}
	*e = Date(l.value)
	return nil
}

type DateApproximated string

func (e *DateApproximated) parse(l Line) error {
	if len(l.value) < 4 || len(l.value) > 35 {
		return ErrInvalidLength{"DateApproximated", l.value, 4, 35}
	}
	*e = DateApproximated(l.value)
	return nil
}

type DateCalendar string

func (e *DateCalendar) parse(l Line) error {
	if len(l.value) < 4 || len(l.value) > 35 {
		return ErrInvalidLength{"DateCalendar", l.value, 4, 35}
	}
	*e = DateCalendar(l.value)
	return nil
}

type DateCalendarEscape string

func (e *DateCalendarEscape) parse(l Line) error {
	if len(l.value) < 4 || len(l.value) > 15 {
		return ErrInvalidLength{"DateCalendarEscape", l.value, 4, 15}
	}
	*e = DateCalendarEscape(l.value)
	return nil
}

type DateExact string

func (e *DateExact) parse(l Line) error {
	if len(l.value) < 10 || len(l.value) > 11 {
		return ErrInvalidLength{"DateExact", l.value, 10, 11}
	}
	*e = DateExact(l.value)
	return nil
}

type DateFren string

func (e *DateFren) parse(l Line) error {
	if len(l.value) < 4 || len(l.value) > 35 {
		return ErrInvalidLength{"DateFren", l.value, 4, 35}
	}
	*e = DateFren(l.value)
	return nil
}

type DateGreg string

func (e *DateGreg) parse(l Line) error {
	if len(l.value) < 4 || len(l.value) > 35 {
		return ErrInvalidLength{"DateGreg", l.value, 4, 35}
	}
	*e = DateGreg(l.value)
	return nil
}

type DateHebr string

func (e *DateHebr) parse(l Line) error {
	if len(l.value) < 4 || len(l.value) > 35 {
		return ErrInvalidLength{"DateHebr", l.value, 4, 35}
	}
	*e = DateHebr(l.value)
	return nil
}

type DateJuln string

func (e *DateJuln) parse(l Line) error {
	if len(l.value) < 4 || len(l.value) > 35 {
		return ErrInvalidLength{"DateJuln", l.value, 4, 35}
	}
	*e = DateJuln(l.value)
	return nil
}

type DateLSDOrd string

func (e *DateLSDOrd) parse(l Line) error {
	if len(l.value) < 4 || len(l.value) > 35 {
		return ErrInvalidLength{"DateLSDOrd", l.value, 4, 35}
	}
	*e = DateLSDOrd(l.value)
	return nil
}

type DatePeriod string

func (e *DatePeriod) parse(l Line) error {
	if len(l.value) < 7 || len(l.value) > 35 {
		return ErrInvalidLength{"DatePeriod", l.value, 7, 35}
	}
	*e = DatePeriod(l.value)
	return nil
}

type DatePhrase string

func (e *DatePhrase) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 35 {
		return ErrInvalidLength{"DatePhrase", l.value, 1, 35}
	}
	*e = DatePhrase(l.value)
	return nil
}

type DateRange string

func (e *DateRange) parse(l Line) error {
	if len(l.value) < 8 || len(l.value) > 35 {
		return ErrInvalidLength{"DateRange", l.value, 8, 35}
	}
	*e = DateRange(l.value)
	return nil
}

type DateValue string

func (e *DateValue) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 35 {
		return ErrInvalidLength{"DateValue", l.value, 1, 35}
	}
	*e = DateValue(l.value)
	return nil
}

type Day uint8

func (e *Day) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 2 {
		return ErrInvalidLength{"Day", l.value, 1, 2}
	}
	n, err := strconv.ParseUint(l.value, 10, 8)
	if err != nil {
		return err
	}
	*e = Day(n)
	return nil
}

type DescriptiveTitle string

func (e *DescriptiveTitle) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 248 {
		return ErrInvalidLength{"DescriptiveTitle", l.value, 1, 248}
	}
	*e = DescriptiveTitle(l.value)
	return nil
}

type Digit uint8

func (e *Digit) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 1 {
		return ErrInvalidLength{"Digit", l.value, 1, 1}
	}
	n, err := strconv.ParseUint(l.value, 10, 8)
	if err != nil {
		return err
	}
	*e = Digit(n)
	return nil
}

type EncodedMultimediaLine string

func (e *EncodedMultimediaLine) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 87 {
		return ErrInvalidLength{"EncodedMultimediaLine", l.value, 1, 87}
	}
	*e = EncodedMultimediaLine(l.value)
	return nil
}

type EntryRecordingDate string

func (e *EntryRecordingDate) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 90 {
		return ErrInvalidLength{"EntryRecordingDate", l.value, 1, 90}
	}
	*e = EntryRecordingDate(l.value)
	return nil
}

type EventAttributeType string

func (e *EventAttributeType) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 15 {
		return ErrInvalidLength{"EventAttributeType", l.value, 1, 15}
	}
	*e = EventAttributeType(l.value)
	return nil
}

type EventDescriptor string

func (e *EventDescriptor) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 90 {
		return ErrInvalidLength{"EventDescriptor", l.value, 1, 90}
	}
	*e = EventDescriptor(l.value)
	return nil
}

type EventTypeCitedFrom string

func (e *EventTypeCitedFrom) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 15 {
		return ErrInvalidLength{"EventTypeCitedFrom", l.value, 1, 15}
	}
	*e = EventTypeCitedFrom(l.value)
	return nil
}

type EventTypeFamile string

func (e *EventTypeFamile) parse(l Line) error {
	switch strings.ToUpper(l.value) {
	case "ANUL":
		*e = "ANUL"
	case "CENS":
		*e = "CENS"
	case "DIV":
		*e = "DIV"
	case "DIVF":
		*e = "DIVF"
	case "ENGA":
		*e = "ENGA"
	case "MARR":
		*e = "MARR"
	case "MARB":
		*e = "MARB"
	case "MARC":
		*e = "MARC"
	case "MARL":
		*e = "MARL"
	case "MARS":
		*e = "MARS"
	case "EVEN":
		*e = "EVEN"
	default:
		return ErrInvalidValue{"EventTypeFamile", l.value}
	}
	return nil
}

type EventTypeIndividual string

func (e *EventTypeIndividual) parse(l Line) error {
	switch strings.ToUpper(l.value) {
	case "ADOP":
		*e = "ADOP"
	case "BIRT":
		*e = "BIRT"
	case "BAPM":
		*e = "BAPM"
	case "BARM":
		*e = "BARM"
	case "BASM":
		*e = "BASM"
	case "BLES":
		*e = "BLES"
	case "BURI":
		*e = "BURI"
	case "CENS":
		*e = "CENS"
	case "CHR":
		*e = "CHR"
	case "CHRA":
		*e = "CHRA"
	case "CONF":
		*e = "CONF"
	case "CREM":
		*e = "CREM"
	case "DEAT":
		*e = "DEAT"
	case "EMIG":
		*e = "EMIG"
	case "FCOM":
		*e = "FCOM"
	case "GRAD":
		*e = "GRAD"
	case "IMMI":
		*e = "IMMI"
	case "NATU":
		*e = "NATU"
	case "ORDN":
		*e = "ORDN"
	case "RETI":
		*e = "RETI"
	case "PROB":
		*e = "PROB"
	case "WILL":
		*e = "WILL"
	case "EVEN":
		*e = "EVEN"
	default:
		return ErrInvalidValue{"EventTypeIndividual", l.value}
	}
	return nil
}

type EventsRecorded string

func (e *EventsRecorded) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 90 {
		return ErrInvalidLength{"EventsRecorded", l.value, 1, 90}
	}
	*e = EventsRecorded(l.value)
	return nil
}

type FileName string

func (e *FileName) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 90 {
		return ErrInvalidLength{"FileName", l.value, 1, 90}
	}
	*e = FileName(l.value)
	return nil
}

type GedcomContentDescription string

func (e *GedcomContentDescription) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 248 {
		return ErrInvalidLength{"GedcomContentDescription", l.value, 1, 248}
	}
	*e = GedcomContentDescription(l.value)
	return nil
}

type GedcomForm string

func (e *GedcomForm) parse(l Line) error {
	if len(l.value) < 14 || len(l.value) > 20 {
		return ErrInvalidLength{"GedcomForm", l.value, 14, 20}
	}
	*e = GedcomForm(l.value)
	return nil
}

type GenerationsOfAncestors uint16

func (e *GenerationsOfAncestors) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 4 {
		return ErrInvalidLength{"GenerationsOfAncestors", l.value, 1, 4}
	}
	n, err := strconv.ParseUint(l.value, 10, 16)
	if err != nil {
		return err
	}
	*e = GenerationsOfAncestors(n)
	return nil
}

type GenerationsOfDescendants uint16

func (e *GenerationsOfDescendants) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 4 {
		return ErrInvalidLength{"GenerationsOfDescendants", l.value, 1, 4}
	}
	n, err := strconv.ParseUint(l.value, 10, 16)
	if err != nil {
		return err
	}
	*e = GenerationsOfDescendants(n)
	return nil
}

type LanguageID string

func (e *LanguageID) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 15 {
		return ErrInvalidLength{"LanguageID", l.value, 1, 15}
	}
	*e = LanguageID(l.value)
	return nil
}

type LanguageOfText string

func (e *LanguageOfText) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 15 {
		return ErrInvalidLength{"LanguageOfText", l.value, 1, 15}
	}
	*e = LanguageOfText(l.value)
	return nil
}

type LanguagePreference string

func (e *LanguagePreference) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 90 {
		return ErrInvalidLength{"LanguagePreference", l.value, 1, 90}
	}
	*e = LanguagePreference(l.value)
	return nil
}

type LDSBaptismDateStatus string

func (e *LDSBaptismDateStatus) parse(l Line) error {
	switch strings.ToUpper(l.value) {
	case "CHILD":
		*e = "CHILD"
	case "CLEARED":
		*e = "CLEARED"
	case "COMPLETED":
		*e = "COMPLETED"
	case "INFANT":
		*e = "INFANT"
	case "PRE-1970":
		*e = "PRE-1970"
	case "QUALIFIED":
		*e = "QUALIFIED"
	case "STILLBORN":
		*e = "STILLBORN"
	case "SUBMITTED":
		*e = "SUBMITTED"
	case "UNCLEARED":
		*e = "UNCLEARED"
	default:
		return ErrInvalidValue{"LDSBaptismDateStatus", l.value}
	}
	return nil
}

type LDSChildSealingDateStatus string

func (e *LDSChildSealingDateStatus) parse(l Line) error {
	switch strings.ToUpper(l.value) {
	case "BIC":
		*e = "BIC"
	case "CLEARED":
		*e = "CLEARED"
	case "COMPLETED":
		*e = "COMPLETED"
	case "DNS":
		*e = "DNS"
	case "PRE-1970":
		*e = "PRE-1970"
	case "QUALIFIED":
		*e = "QUALIFIED"
	case "STILLBORN":
		*e = "STILLBORN"
	case "SUBMITTES":
		*e = "SUBMITTES"
	case "UNCLEARED":
		*e = "UNCLEARED"
	default:
		return ErrInvalidValue{"LDSChildSealingDateStatus", l.value}
	}
	return nil
}

type LDSEndowmentDateStatus string

func (e *LDSEndowmentDateStatus) parse(l Line) error {
	switch strings.ToUpper(l.value) {
	case "CHILD":
		*e = "CHILD"
	case "CLEARED":
		*e = "CLEARED"
	case "COMPLETED":
		*e = "COMPLETED"
	case "INFANT":
		*e = "INFANT"
	case "PRE-1970":
		*e = "PRE-1970"
	case "QUALIFIED":
		*e = "QUALIFIED"
	case "STILLBORN":
		*e = "STILLBORN"
	case "SUBMITTED":
		*e = "SUBMITTED"
	case "UNCLEARED":
		*e = "UNCLEARED"
	default:
		return ErrInvalidValue{"LDSEndowmentDateStatus", l.value}
	}
	return nil
}

type LDSSpouseSealingDateStatus string

func (e *LDSSpouseSealingDateStatus) parse(l Line) error {
	switch strings.ToUpper(l.value) {
	case "CANCELED":
		*e = "CANCELED"
	case "CLEARED":
		*e = "CLEARED"
	case "COMPLETED":
		*e = "COMPLETED"
	case "DNS":
		*e = "DNS"
	case "DNS/CAN":
		*e = "DNS/CAN"
	case "PRE-1970":
		*e = "PRE-1970"
	case "QUALIFIED":
		*e = "QUALIFIED"
	case "SUBMITTED":
		*e = "SUBMITTED"
	case "UNCLEARED":
		*e = "UNCLEARED"
	default:
		return ErrInvalidValue{"LDSSpouseSealingDateStatus", l.value}
	}
	return nil
}

type Month string

func (e *Month) parse(l Line) error {
	switch strings.ToUpper(l.value) {
	case "JAN":
		*e = "JAN"
	case "FEB":
		*e = "FEB"
	case "MAR":
		*e = "MAR"
	case "APR":
		*e = "APR"
	case "MAY":
		*e = "MAY"
	case "JUN":
		*e = "JUN"
	case "JUL":
		*e = "JUL"
	case "AUG":
		*e = "AUG"
	case "SEP":
		*e = "SEP"
	case "OCT":
		*e = "OCT"
	case "NOV":
		*e = "NOV"
	case "DEC":
		*e = "DEC"
	default:
		return ErrInvalidValue{"Month", l.value}
	}
	return nil
}

type MonthFren string

func (e *MonthFren) parse(l Line) error {
	switch strings.ToUpper(l.value) {
	case "VEND":
		*e = "VEND"
	case "BRUM":
		*e = "BRUM"
	case "FRIM":
		*e = "FRIM"
	case "NIVO":
		*e = "NIVO"
	case "PLUV":
		*e = "PLUV"
	case "VENT":
		*e = "VENT"
	case "GERM":
		*e = "GERM"
	case "FLOR":
		*e = "FLOR"
	case "PRAI":
		*e = "PRAI"
	case "MESS":
		*e = "MESS"
	case "THER":
		*e = "THER"
	case "FRUC":
		*e = "FRUC"
	case "COMP":
		*e = "COMP"
	default:
		return ErrInvalidValue{"MonthFren", l.value}
	}
	return nil
}

type MonthHebr string

func (e *MonthHebr) parse(l Line) error {
	switch strings.ToUpper(l.value) {
	case "TSH":
		*e = "TSH"
	case "CSH":
		*e = "CSH"
	case "KSL":
		*e = "KSL"
	case "TVT":
		*e = "TVT"
	case "SHV":
		*e = "SHV"
	case "ADR":
		*e = "ADR"
	case "ADS":
		*e = "ADS"
	case "NSN":
		*e = "NSN"
	case "IYR":
		*e = "IYR"
	case "SVN":
		*e = "SVN"
	case "TMZ":
		*e = "TMZ"
	case "AAV":
		*e = "AAV"
	case "ELL":
		*e = "ELL"
	default:
		return ErrInvalidValue{"MonthHebr", l.value}
	}
	return nil
}

type MultimediaFileReference string

func (e *MultimediaFileReference) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 30 {
		return ErrInvalidLength{"MultimediaFileReference", l.value, 1, 30}
	}
	*e = MultimediaFileReference(l.value)
	return nil
}

type MultimediaFormat string

func (e *MultimediaFormat) parse(l Line) error {
	switch strings.ToLower(l.value) {
	case "bmp":
		*e = "bmp"
	case "gif":
		*e = "gif"
	case "jpeg":
		*e = "jpeg"
	case "ole":
		*e = "ole"
	case "pcx":
		*e = "pcx"
	case "tiff":
		*e = "tiff"
	case "wav":
		*e = "wav"
	default:
		return ErrInvalidValue{"MultimediaFormat", l.value}
	}
	return nil
}

type NameOfBusiness string

func (e *NameOfBusiness) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 90 {
		return ErrInvalidLength{"NameOfBusiness", l.value, 1, 90}
	}
	*e = NameOfBusiness(l.value)
	return nil
}

type NameOfFamilyFile string

func (e *NameOfFamilyFile) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 20 {
		return ErrInvalidLength{"NameOfFamilyFile", l.value, 1, 20}
	}
	*e = NameOfFamilyFile(l.value)
	return nil
}

type NameOfProduct string

func (e *NameOfProduct) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 90 {
		return ErrInvalidLength{"NameOfProduct", l.value, 1, 90}
	}
	*e = NameOfProduct(l.value)
	return nil
}

type NameOfRepository string

func (e *NameOfRepository) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 90 {
		return ErrInvalidLength{"NameOfRepository", l.value, 1, 90}
	}
	*e = NameOfRepository(l.value)
	return nil
}

type NameOfSourceData string

func (e *NameOfSourceData) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 90 {
		return ErrInvalidLength{"NameOfSourceData", l.value, 1, 90}
	}
	*e = NameOfSourceData(l.value)
	return nil
}

type NameOfPersonal string

func (e *NameOfPersonal) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 20 {
		return ErrInvalidLength{"NameOfPersonal", l.value, 1, 20}
	}
	*e = NameOfPersonal(l.value)
	return nil
}

type NamePiece string

func (e *NamePiece) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 90 {
		return ErrInvalidLength{"NamePiece", l.value, 1, 90}
	}
	*e = NamePiece(l.value)
	return nil
}

type NamePieceGiven string

func (e *NamePieceGiven) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 120 {
		return ErrInvalidLength{"NamePieceGiven", l.value, 1, 120}
	}
	*e = NamePieceGiven(l.value)
	return nil
}

type NamePieceNickname string

func (e *NamePieceNickname) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 30 {
		return ErrInvalidLength{"NamePieceNickname", l.value, 1, 30}
	}
	*e = NamePieceNickname(l.value)
	return nil
}

type NamePiecePrefix string

func (e *NamePiecePrefix) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 30 {
		return ErrInvalidLength{"NamePiecePrefix", l.value, 1, 30}
	}
	*e = NamePiecePrefix(l.value)
	return nil
}

type NamePieceSuffix string

func (e *NamePieceSuffix) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 30 {
		return ErrInvalidLength{"NamePieceSuffix", l.value, 1, 30}
	}
	*e = NamePieceSuffix(l.value)
	return nil
}

type NamePieceSurname string

func (e *NamePieceSurname) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 120 {
		return ErrInvalidLength{"NamePieceSurname", l.value, 1, 120}
	}
	*e = NamePieceSurname(l.value)
	return nil
}

type NamePieceSurnamePrefix string

func (e *NamePieceSurnamePrefix) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 30 {
		return ErrInvalidLength{"NamePieceSurnamePrefix", l.value, 1, 30}
	}
	*e = NamePieceSurnamePrefix(l.value)
	return nil
}

type NationalIDNumber string

func (e *NationalIDNumber) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 30 {
		return ErrInvalidLength{"NationalIDNumber", l.value, 1, 30}
	}
	*e = NationalIDNumber(l.value)
	return nil
}

type NationalOrTribalOrigin string

func (e *NationalOrTribalOrigin) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 120 {
		return ErrInvalidLength{"NationalOrTribalOrigin", l.value, 1, 120}
	}
	*e = NationalOrTribalOrigin(l.value)
	return nil
}

type NewTag string

func (e *NewTag) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 15 {
		return ErrInvalidLength{"NewTag", l.value, 1, 15}
	}
	*e = NewTag(l.value)
	return nil
}

type NobilityTypeTitle string

func (e *NobilityTypeTitle) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 120 {
		return ErrInvalidLength{"NobilityTypeTitle", l.value, 1, 120}
	}
	*e = NobilityTypeTitle(l.value)
	return nil
}

type Number uint

func (e *Number) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 9 {
		return ErrInvalidLength{"Number", l.value, 1, 9}
	}
	n, err := strconv.ParseUint(l.value, 10, 0)
	if err != nil {
		return err
	}
	*e = Number(n)
	return nil
}

type Occupation string

func (e *Occupation) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 90 {
		return ErrInvalidLength{"Occupation", l.value, 1, 90}
	}
	*e = Occupation(l.value)
	return nil
}

type OrginanceProcessFlag string

func (e *OrginanceProcessFlag) parse(l Line) error {
	switch strings.ToLower(l.value) {
	case "yes":
		*e = "yes"
	case "no":
		*e = "no"
	default:
		return ErrInvalidValue{"OrginanceProcessFlag", l.value}
	}
	return nil
}

type PedigreeLinkageType string

func (e *PedigreeLinkageType) parse(l Line) error {
	switch strings.ToLower(l.value) {
	case "adopted":
		*e = "adopted"
	case "birth":
		*e = "birth"
	case "foster":
		*e = "foster"
	case "sealing":
		*e = "sealing"
	default:
		return ErrInvalidValue{"PedigreeLinkageType", l.value}
	}
	return nil
}

type PermanentRecordFileNumber string

func (e *PermanentRecordFileNumber) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 90 {
		return ErrInvalidLength{"PermanentRecordFileNumber", l.value, 1, 90}
	}
	*e = PermanentRecordFileNumber(l.value)
	return nil
}

type PhoneNumber string

func (e *PhoneNumber) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 25 {
		return ErrInvalidLength{"PhoneNumber", l.value, 1, 25}
	}
	*e = PhoneNumber(l.value)
	return nil
}

type PhysicalDescription string

func (e *PhysicalDescription) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 248 {
		return ErrInvalidLength{"PhysicalDescription", l.value, 1, 248}
	}
	*e = PhysicalDescription(l.value)
	return nil
}

type PlaceHierachy string

func (e *PlaceHierachy) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 120 {
		return ErrInvalidLength{"PlaceHierachy", l.value, 1, 120}
	}
	*e = PlaceHierachy(l.value)
	return nil
}

type PlaceLivingOrdinance string

func (e *PlaceLivingOrdinance) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 120 {
		return ErrInvalidLength{"PlaceLivingOrdinance", l.value, 1, 120}
	}
	*e = PlaceLivingOrdinance(l.value)
	return nil
}

type PlaceValue string

func (e *PlaceValue) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 120 {
		return ErrInvalidLength{"PlaceValue", l.value, 1, 120}
	}
	*e = PlaceValue(l.value)
	return nil
}

type Possessions string

func (e *Possessions) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 248 {
		return ErrInvalidLength{"Possessions", l.value, 1, 248}
	}
	*e = Possessions(l.value)
	return nil
}

type PublicationDate string

func (e *PublicationDate) parse(l Line) error {
	if len(l.value) < 10 || len(l.value) > 11 {
		return ErrInvalidLength{"PublicationDate", l.value, 10, 11}
	}
	*e = PublicationDate(l.value)
	return nil
}

type ReceivingSystemName string

func (e *ReceivingSystemName) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 20 {
		return ErrInvalidLength{"ReceivingSystemName", l.value, 1, 20}
	}
	*e = ReceivingSystemName(l.value)
	return nil
}

type RecordIdentifier string

func (e *RecordIdentifier) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 18 {
		return ErrInvalidLength{"RecordIdentifier", l.value, 1, 18}
	}
	*e = RecordIdentifier(l.value)
	return nil
}

type RecordType string

func (e *RecordType) parse(l Line) error {
	switch strings.ToUpper(l.value) {
	case "FAM":
		*e = "FAM"
	case "INDI":
		*e = "INDI"
	case "NOTE":
		*e = "NOTE"
	case "OBJE":
		*e = "OBJE"
	case "REPO":
		*e = "REPO"
	case "SOUR":
		*e = "SOUR"
	case "SUBM":
		*e = "SUBM"
	case "SUBN":
		*e = "SUBN"
	default:
		return ErrInvalidValue{"RecordType", l.value}
	}
	return nil
}

type RegisteredResourceIdentifier string

func (e *RegisteredResourceIdentifier) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 25 {
		return ErrInvalidLength{"RegisteredResourceIdentifier", l.value, 1, 25}
	}
	*e = RegisteredResourceIdentifier(l.value)
	return nil
}

type RelationIsDescriptor string

func (e *RelationIsDescriptor) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 25 {
		return ErrInvalidLength{"RelationIsDescriptor", l.value, 1, 25}
	}
	*e = RelationIsDescriptor(l.value)
	return nil
}

type ReligiousAffiliation string

func (e *ReligiousAffiliation) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 90 {
		return ErrInvalidLength{"ReligiousAffiliation", l.value, 1, 90}
	}
	*e = ReligiousAffiliation(l.value)
	return nil
}

type ResponsibleAgency string

func (e *ResponsibleAgency) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 120 {
		return ErrInvalidLength{"ResponsibleAgency", l.value, 1, 120}
	}
	*e = ResponsibleAgency(l.value)
	return nil
}

type RestrictionNotice string

func (e *RestrictionNotice) parse(l Line) error {
	switch strings.ToLower(l.value) {
	case "locked":
		*e = "locked"
	case "privacy":
		*e = "privacy"
	default:
		return ErrInvalidValue{"RestrictionNotice", l.value}
	}
	return nil
}

type RoleDescriptor string

func (e *RoleDescriptor) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 25 {
		return ErrInvalidLength{"RoleDescriptor", l.value, 1, 25}
	}
	*e = RoleDescriptor(l.value)
	return nil
}

type RoleInEvent string

func (e *RoleInEvent) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 15 {
		return ErrInvalidLength{"RoleInEvent", l.value, 1, 15}
	}
	*e = RoleInEvent(l.value)
	return nil
}

type ScholasticAchievement string

func (e *ScholasticAchievement) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 248 {
		return ErrInvalidLength{"ScholasticAchievement", l.value, 1, 248}
	}
	*e = ScholasticAchievement(l.value)
	return nil
}

type SexValue string

func (e *SexValue) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 7 {
		return ErrInvalidLength{"SexValue", l.value, 1, 7}
	}
	*e = SexValue(l.value)
	return nil
}

type SocialSecurityNumber string

func (e *SocialSecurityNumber) parse(l Line) error {
	if len(l.value) < 9 || len(l.value) > 11 {
		return ErrInvalidLength{"SocialSecurityNumber", l.value, 9, 11}
	}
	*e = SocialSecurityNumber(l.value)
	return nil
}

type SourceCallNumber string

func (e *SourceCallNumber) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 120 {
		return ErrInvalidLength{"SourceCallNumber", l.value, 1, 120}
	}
	*e = SourceCallNumber(l.value)
	return nil
}

type SourceDescription string

func (e *SourceDescription) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 248 {
		return ErrInvalidLength{"SourceDescription", l.value, 1, 248}
	}
	*e = SourceDescription(l.value)
	return nil
}

type SourceDescriptiveTitle string

func (e *SourceDescriptiveTitle) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 248 {
		return ErrInvalidLength{"SourceDescriptiveTitle", l.value, 1, 248}
	}
	*e = SourceDescriptiveTitle(l.value)
	return nil
}

type SourceFiledByEntry string

func (e *SourceFiledByEntry) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 60 {
		return ErrInvalidLength{"SourceFiledByEntry", l.value, 1, 60}
	}
	*e = SourceFiledByEntry(l.value)
	return nil
}

type SourceJurisdictionPlace string

func (e *SourceJurisdictionPlace) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 120 {
		return ErrInvalidLength{"SourceJurisdictionPlace", l.value, 1, 120}
	}
	*e = SourceJurisdictionPlace(l.value)
	return nil
}

type SourceMediaType string

func (e *SourceMediaType) parse(l Line) error {
	switch strings.ToLower(l.value) {
	case "audio":
		*e = "audio"
	case "book":
		*e = "book"
	case "card":
		*e = "card"
	case "electronic":
		*e = "electronic"
	case "fiche":
		*e = "fiche"
	case "film":
		*e = "film"
	case "magazine":
		*e = "magazine"
	case "manuscript":
		*e = "manuscript"
	case "map":
		*e = "map"
	case "newspaper":
		*e = "newspaper"
	case "photo":
		*e = "photo"
	case "tombstone":
		*e = "tombstone"
	case "video":
		*e = "video"
	default:
		return ErrInvalidValue{"SourceMediaType", l.value}
	}
	return nil
}

type SourceOriginator string

func (e *SourceOriginator) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 248 {
		return ErrInvalidLength{"SourceOriginator", l.value, 1, 248}
	}
	*e = SourceOriginator(l.value)
	return nil
}

type SourcePublicationFacts string

func (e *SourcePublicationFacts) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 248 {
		return ErrInvalidLength{"SourcePublicationFacts", l.value, 1, 248}
	}
	*e = SourcePublicationFacts(l.value)
	return nil
}

type SubmitterName string

func (e *SubmitterName) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 60 {
		return ErrInvalidLength{"SubmitterName", l.value, 1, 60}
	}
	*e = SubmitterName(l.value)
	return nil
}

type SubmitterRegisteredRFN string

func (e *SubmitterRegisteredRFN) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 30 {
		return ErrInvalidLength{"SubmitterRegisteredRFN", l.value, 1, 30}
	}
	*e = SubmitterRegisteredRFN(l.value)
	return nil
}

type SubmitterText string

func (e *SubmitterText) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 248 {
		return ErrInvalidLength{"SubmitterText", l.value, 1, 248}
	}
	*e = SubmitterText(l.value)
	return nil
}

type TempleCode string

func (e *TempleCode) parse(l Line) error {
	if len(l.value) < 4 || len(l.value) > 5 {
		return ErrInvalidLength{"TempleCode", l.value, 4, 5}
	}
	*e = TempleCode(l.value)
	return nil
}

type Text string

func (e *Text) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 248 {
		return ErrInvalidLength{"Text", l.value, 1, 248}
	}
	*e = Text(l.value)
	return nil
}

type TextFromSource string

func (e *TextFromSource) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 248 {
		return ErrInvalidLength{"TextFromSource", l.value, 1, 248}
	}
	*e = TextFromSource(l.value)
	return nil
}

type TimeValue string

func (e *TimeValue) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 12 {
		return ErrInvalidLength{"TimeValue", l.value, 1, 12}
	}
	*e = TimeValue(l.value)
	return nil
}

type TransmissionDate string

func (e *TransmissionDate) parse(l Line) error {
	if len(l.value) < 10 || len(l.value) > 11 {
		return ErrInvalidLength{"TransmissionDate", l.value, 10, 11}
	}
	*e = TransmissionDate(l.value)
	return nil
}

type UserReferenceNumber string

func (e *UserReferenceNumber) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 20 {
		return ErrInvalidLength{"UserReferenceNumber", l.value, 1, 20}
	}
	*e = UserReferenceNumber(l.value)
	return nil
}

type UserReferenceType string

func (e *UserReferenceType) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 40 {
		return ErrInvalidLength{"UserReferenceType", l.value, 1, 40}
	}
	*e = UserReferenceType(l.value)
	return nil
}

type VersionNumber string

func (e *VersionNumber) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 15 {
		return ErrInvalidLength{"VersionNumber", l.value, 1, 15}
	}
	*e = VersionNumber(l.value)
	return nil
}

type WhereWithinSource string

func (e *WhereWithinSource) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 248 {
		return ErrInvalidLength{"WhereWithinSource", l.value, 1, 248}
	}
	*e = WhereWithinSource(l.value)
	return nil
}

type Xref string

func (e *Xref) parse(l Line) error {
	if len(l.value) < 1 || len(l.value) > 22 {
		return ErrInvalidLength{"Xref", l.value, 1, 22}
	}
	*e = Xref(l.value)
	return nil
}

type Year string

func (e *Year) parse(l Line) error {
	if len(l.value) < 3 || len(l.value) > 4 {
		return ErrInvalidLength{"Year", l.value, 3, 4}
	}
	*e = Year(l.value)
	return nil
}

type YearGreg string

func (e *YearGreg) parse(l Line) error {
	if len(l.value) < 3 || len(l.value) > 7 {
		return ErrInvalidLength{"YearGreg", l.value, 3, 7}
	}
	*e = YearGreg(l.value)
	return nil
}

type ErrInvalidValue struct {
	Type, Value string
}

func (e ErrInvalidValue) Error() string {
	return "Value for " + e.Type + " is invalid"
}

type ErrInvalidLength struct {
	Type, Value string
	Min, Max    uint
}

func (e ErrInvalidLength) Error() string {
	return "Value for " + e.Type + " has an invalid length"
}
