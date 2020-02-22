package gedcom

// File automatically generated with ./genStructures.sh

import "errors"

// Header is a GEDCOM structure type
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

func (s *Header) parse(l *Line, o options) error {
	var SourceSet, SubmitterSet, VersionSet, CharacterSetSet, ReceivingSystemNameSet, TransmissionLDateSet, SubmissionSet, FileNameSet, CopyrightSet, LanguageSet, PlaceSet, ContentDescriptionSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case cSOUR:
			if SourceSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Header", cSOUR, ErrSingleMultiple}
			}
			SourceSet = true
			if err := s.Source.parse(&sl, o); err != nil {
				return ErrContext{"Header", cSOUR, err}
			}
		case cDEST:
			if ReceivingSystemNameSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Header", cDEST, ErrSingleMultiple}
			}
			ReceivingSystemNameSet = true
			if err := s.ReceivingSystemName.parse(&sl, o); err != nil {
				return ErrContext{"Header", cDEST, err}
			}
		case cDATE:
			if TransmissionLDateSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Header", cDATE, ErrSingleMultiple}
			}
			TransmissionLDateSet = true
			if err := s.TransmissionLDate.parse(&sl, o); err != nil {
				return ErrContext{"Header", cDATE, err}
			}
		case cSUBM:
			if SubmitterSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Header", cSUBM, ErrSingleMultiple}
			}
			SubmitterSet = true
			if err := s.Submitter.parse(&sl, o); err != nil {
				return ErrContext{"Header", cSUBM, err}
			}
		case cSUMB:
			if SubmissionSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Header", cSUMB, ErrSingleMultiple}
			}
			SubmissionSet = true
			if err := s.Submission.parse(&sl, o); err != nil {
				return ErrContext{"Header", cSUMB, err}
			}
		case cFILE:
			if FileNameSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Header", cFILE, ErrSingleMultiple}
			}
			FileNameSet = true
			if err := s.FileName.parse(&sl, o); err != nil {
				return ErrContext{"Header", cFILE, err}
			}
		case cCOPR:
			if CopyrightSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Header", cCOPR, ErrSingleMultiple}
			}
			CopyrightSet = true
			if err := s.Copyright.parse(&sl, o); err != nil {
				return ErrContext{"Header", cCOPR, err}
			}
		case cGEDC:
			if VersionSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Header", cGEDC, ErrSingleMultiple}
			}
			VersionSet = true
			if err := s.Version.parse(&sl, o); err != nil {
				return ErrContext{"Header", cGEDC, err}
			}
		case cCHAR:
			if CharacterSetSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Header", cCHAR, ErrSingleMultiple}
			}
			CharacterSetSet = true
			if err := s.CharacterSet.parse(&sl, o); err != nil {
				return ErrContext{"Header", cCHAR, err}
			}
		case cLANG:
			if LanguageSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Header", cLANG, ErrSingleMultiple}
			}
			LanguageSet = true
			if err := s.Language.parse(&sl, o); err != nil {
				return ErrContext{"Header", cLANG, err}
			}
		case cPLAC:
			if PlaceSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Header", cPLAC, ErrSingleMultiple}
			}
			PlaceSet = true
			if err := s.Place.parse(&sl, o); err != nil {
				return ErrContext{"Header", cPLAC, err}
			}
		case cNOTE:
			if ContentDescriptionSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Header", cNOTE, ErrSingleMultiple}
			}
			ContentDescriptionSet = true
			if err := s.ContentDescription.parse(&sl, o); err != nil {
				return ErrContext{"Header", cNOTE, err}
			}
		default:
			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
				return ErrContext{"Header", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	if !o.allowMissingRequired {
		if !SourceSet {
			return ErrContext{"Header", "Source", ErrRequiredMissing}
		}
		if !SubmitterSet {
			return ErrContext{"Header", "Submitter", ErrRequiredMissing}
		}
		if !VersionSet {
			return ErrContext{"Header", "Version", ErrRequiredMissing}
		}
		if !CharacterSetSet {
			return ErrContext{"Header", "CharacterSet", ErrRequiredMissing}
		}
	}
	return nil
}

// HeaderSource is a GEDCOM structure type
type HeaderSource struct {
	SystemID      ApprovedSystemID
	VersionNumber VersionNumber
	Name          NameOfProduct
	Business      HeaderBusiness
	Data          HeaderDataSource
}

func (s *HeaderSource) parse(l *Line, o options) error {
	if err := s.SystemID.parse(l, o); err != nil {
		return ErrContext{"HeaderSource", "line_value", err}
	}
	var VersionNumberSet, NameSet, BusinessSet, DataSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case cVERS:
			if VersionNumberSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"HeaderSource", cVERS, ErrSingleMultiple}
			}
			VersionNumberSet = true
			if err := s.VersionNumber.parse(&sl, o); err != nil {
				return ErrContext{"HeaderSource", cVERS, err}
			}
		case cNAME:
			if NameSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"HeaderSource", cNAME, ErrSingleMultiple}
			}
			NameSet = true
			if err := s.Name.parse(&sl, o); err != nil {
				return ErrContext{"HeaderSource", cNAME, err}
			}
		case cCORP:
			if BusinessSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"HeaderSource", cCORP, ErrSingleMultiple}
			}
			BusinessSet = true
			if err := s.Business.parse(&sl, o); err != nil {
				return ErrContext{"HeaderSource", cCORP, err}
			}
		case cDATA:
			if DataSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"HeaderSource", cDATA, ErrSingleMultiple}
			}
			DataSet = true
			if err := s.Data.parse(&sl, o); err != nil {
				return ErrContext{"HeaderSource", cDATA, err}
			}
		default:
			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
				return ErrContext{"HeaderSource", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// TransmissionDateTime is a GEDCOM structure type
type TransmissionDateTime struct {
	TransmissionDate TransmissionDate
	Time             TimeValue
}

func (s *TransmissionDateTime) parse(l *Line, o options) error {
	if err := s.TransmissionDate.parse(l, o); err != nil {
		return ErrContext{"TransmissionDateTime", "line_value", err}
	}
	var TimeSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case cTIME:
			if TimeSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"TransmissionDateTime", cTIME, ErrSingleMultiple}
			}
			TimeSet = true
			if err := s.Time.parse(&sl, o); err != nil {
				return ErrContext{"TransmissionDateTime", cTIME, err}
			}
		default:
			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
				return ErrContext{"TransmissionDateTime", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// HeaderBusiness is a GEDCOM structure type
type HeaderBusiness struct {
	NameOfBusiness NameOfBusiness
	Address        AddressStructure
	PhoneNumber    []PhoneNumber
}

func (s *HeaderBusiness) parse(l *Line, o options) error {
	if err := s.NameOfBusiness.parse(l, o); err != nil {
		return ErrContext{"HeaderBusiness", "line_value", err}
	}
	var AddressSet bool
	s.PhoneNumber = make([]PhoneNumber, 0, 3)
	for _, sl := range l.Sub {
		switch sl.tag {
		case cADDR:
			if AddressSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"HeaderBusiness", cADDR, ErrSingleMultiple}
			}
			AddressSet = true
			if err := s.Address.parse(&sl, o); err != nil {
				return ErrContext{"HeaderBusiness", cADDR, err}
			}
		case cPHON:
			if len(s.PhoneNumber) == 3 {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"HeaderBusiness", cPHON, ErrTooMany(3)}
			}
			var t PhoneNumber
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"HeaderBusiness", cPHON, err}
			}
			s.PhoneNumber = append(s.PhoneNumber, t)
		default:
			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
				return ErrContext{"HeaderBusiness", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// HeaderDataSource is a GEDCOM structure type
type HeaderDataSource struct {
	SourceName          NameOfSourceData
	PublicationDate     PublicationDate
	CopyrightSourceData CopyrightSourceData
}

func (s *HeaderDataSource) parse(l *Line, o options) error {
	if err := s.SourceName.parse(l, o); err != nil {
		return ErrContext{"HeaderDataSource", "line_value", err}
	}
	var PublicationDateSet, CopyrightSourceDataSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case cDATE:
			if PublicationDateSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"HeaderDataSource", cDATE, ErrSingleMultiple}
			}
			PublicationDateSet = true
			if err := s.PublicationDate.parse(&sl, o); err != nil {
				return ErrContext{"HeaderDataSource", cDATE, err}
			}
		case cCOPR:
			if CopyrightSourceDataSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"HeaderDataSource", cCOPR, ErrSingleMultiple}
			}
			CopyrightSourceDataSet = true
			if err := s.CopyrightSourceData.parse(&sl, o); err != nil {
				return ErrContext{"HeaderDataSource", cCOPR, err}
			}
		default:
			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
				return ErrContext{"HeaderDataSource", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// Version is a GEDCOM structure type
type Version struct {
	VersionNumber VersionNumber
	Form          Form
}

func (s *Version) parse(l *Line, o options) error {
	var VersionNumberSet, FormSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case cVERS:
			if VersionNumberSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Version", cVERS, ErrSingleMultiple}
			}
			VersionNumberSet = true
			if err := s.VersionNumber.parse(&sl, o); err != nil {
				return ErrContext{"Version", cVERS, err}
			}
		case cFORM:
			if FormSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Version", cFORM, ErrSingleMultiple}
			}
			FormSet = true
			if err := s.Form.parse(&sl, o); err != nil {
				return ErrContext{"Version", cFORM, err}
			}
		default:
			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
				return ErrContext{"Version", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	if !o.allowMissingRequired {
		if !VersionNumberSet {
			return ErrContext{"Version", "VersionNumber", ErrRequiredMissing}
		}
		if !FormSet {
			return ErrContext{"Version", "Form", ErrRequiredMissing}
		}
	}
	return nil
}

// CharacterSetStructure is a GEDCOM structure type
type CharacterSetStructure struct {
	CharacterSet  CharacterSet
	VersionNumber VersionNumber
}

func (s *CharacterSetStructure) parse(l *Line, o options) error {
	if err := s.CharacterSet.parse(l, o); err != nil {
		return ErrContext{"CharacterSetStructure", "line_value", err}
	}
	var VersionNumberSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case cVERS:
			if VersionNumberSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"CharacterSetStructure", cVERS, ErrSingleMultiple}
			}
			VersionNumberSet = true
			if err := s.VersionNumber.parse(&sl, o); err != nil {
				return ErrContext{"CharacterSetStructure", cVERS, err}
			}
		default:
			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
				return ErrContext{"CharacterSetStructure", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// HeaderPlace is a GEDCOM structure type
type HeaderPlace struct {
	PlaceHierarchy PlaceHierarchy
}

func (s *HeaderPlace) parse(l *Line, o options) error {
	var PlaceHierarchySet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case cFORM:
			if PlaceHierarchySet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"HeaderPlace", cFORM, ErrSingleMultiple}
			}
			PlaceHierarchySet = true
			if err := s.PlaceHierarchy.parse(&sl, o); err != nil {
				return ErrContext{"HeaderPlace", cFORM, err}
			}
		default:
			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
				return ErrContext{"HeaderPlace", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	if !o.allowMissingRequired {
		if !PlaceHierarchySet {
			return ErrContext{"HeaderPlace", "PlaceHierarchy", ErrRequiredMissing}
		}
	}
	return nil
}

// Family is a GEDCOM structure type
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

func (s *Family) parse(l *Line, o options) error {
	if err := s.ID.parse(&Line{line: line{value: l.xrefID}}, o); err != nil {
		return ErrContext{"Family", "xrefID", err}
	}
	var AnnulmentSet, CensusSet, DivorceSet, DivorceFiledSet, EngagementSet, MarriageSet, MarriageBannSet, MarriageContractSet, MarriageLicenseSet, MarriageSettlementSet, HusbandSet, WifeSet, NumChildrenSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case cANUL:
			if AnnulmentSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Family", cANUL, ErrSingleMultiple}
			}
			AnnulmentSet = true
			if err := s.Annulment.parse(&sl, o); err != nil {
				return ErrContext{"Family", cANUL, err}
			}
		case cCENS:
			if CensusSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Family", cCENS, ErrSingleMultiple}
			}
			CensusSet = true
			if err := s.Census.parse(&sl, o); err != nil {
				return ErrContext{"Family", cCENS, err}
			}
		case cDIV:
			if DivorceSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Family", cDIV, ErrSingleMultiple}
			}
			DivorceSet = true
			if err := s.Divorce.parse(&sl, o); err != nil {
				return ErrContext{"Family", cDIV, err}
			}
		case cDIVF:
			if DivorceFiledSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Family", cDIVF, ErrSingleMultiple}
			}
			DivorceFiledSet = true
			if err := s.DivorceFiled.parse(&sl, o); err != nil {
				return ErrContext{"Family", cDIVF, err}
			}
		case cENGA:
			if EngagementSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Family", cENGA, ErrSingleMultiple}
			}
			EngagementSet = true
			if err := s.Engagement.parse(&sl, o); err != nil {
				return ErrContext{"Family", cENGA, err}
			}
		case cMARR:
			if MarriageSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Family", cMARR, ErrSingleMultiple}
			}
			MarriageSet = true
			if err := s.Marriage.parse(&sl, o); err != nil {
				return ErrContext{"Family", cMARR, err}
			}
		case cMARB:
			if MarriageBannSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Family", cMARB, ErrSingleMultiple}
			}
			MarriageBannSet = true
			if err := s.MarriageBann.parse(&sl, o); err != nil {
				return ErrContext{"Family", cMARB, err}
			}
		case cMARC:
			if MarriageContractSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Family", cMARC, ErrSingleMultiple}
			}
			MarriageContractSet = true
			if err := s.MarriageContract.parse(&sl, o); err != nil {
				return ErrContext{"Family", cMARC, err}
			}
		case cMARL:
			if MarriageLicenseSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Family", cMARL, ErrSingleMultiple}
			}
			MarriageLicenseSet = true
			if err := s.MarriageLicense.parse(&sl, o); err != nil {
				return ErrContext{"Family", cMARL, err}
			}
		case cMARS:
			if MarriageSettlementSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Family", cMARS, ErrSingleMultiple}
			}
			MarriageSettlementSet = true
			if err := s.MarriageSettlement.parse(&sl, o); err != nil {
				return ErrContext{"Family", cMARS, err}
			}
		case cEVEN:
			var t FamilyEventDetail
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"Family", cEVEN, err}
			}
			s.Events = append(s.Events, t)
		case cHUSB:
			if HusbandSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Family", cHUSB, ErrSingleMultiple}
			}
			HusbandSet = true
			if err := s.Husband.parse(&sl, o); err != nil {
				return ErrContext{"Family", cHUSB, err}
			}
		case cWIFE:
			if WifeSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Family", cWIFE, ErrSingleMultiple}
			}
			WifeSet = true
			if err := s.Wife.parse(&sl, o); err != nil {
				return ErrContext{"Family", cWIFE, err}
			}
		case cCHIL:
			var t Xref
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"Family", cCHIL, err}
			}
			s.Children = append(s.Children, t)
		case cNCHI:
			if NumChildrenSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Family", cNCHI, ErrSingleMultiple}
			}
			NumChildrenSet = true
			if err := s.NumChildren.parse(&sl, o); err != nil {
				return ErrContext{"Family", cNCHI, err}
			}
		case cSUBM:
			var t Xref
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"Family", cSUBM, err}
			}
			s.Submitters = append(s.Submitters, t)
		case cSLGS:
			var t LDSSpouseSealing
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"Family", cSLGS, err}
			}
			s.LDSSpouseSealing = append(s.LDSSpouseSealing, t)
		case cSOUR:
			var t SourceCitation
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"Family", cSOUR, err}
			}
			s.Sources = append(s.Sources, t)
		case cOBJE:
			var t MultimediaLink
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"Family", cOBJE, err}
			}
			s.Multimedia = append(s.Multimedia, t)
		case cNOTE:
			var t NoteStructure
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"Family", cNOTE, err}
			}
			s.Notes = append(s.Notes, t)
		default:
			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
				return ErrContext{"Family", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// VerifiedFamilyEventDetail is a GEDCOM structure type
type VerifiedFamilyEventDetail struct {
	Verified Verified
	FamilyEventDetail
}

func (s *VerifiedFamilyEventDetail) parse(l *Line, o options) error {
	if err := s.Verified.parse(l, o); err != nil {
		return ErrContext{"VerifiedFamilyEventDetail", "line_value", err}
	}
	return s.FamilyEventDetail.parse(l, o)
}

// FamilyEventDetail is a GEDCOM structure type
type FamilyEventDetail struct {
	HusbandAge AgeStructure
	WifeAge    AgeStructure
	EventDetail
}

func (s *FamilyEventDetail) parse(l *Line, o options) error {
	var HusbandAgeSet, WifeAgeSet bool
	for i := 0; i < len(l.Sub); i++ {
		sl := l.Sub[i]
		switch sl.tag {
		case cHUSB:
			if HusbandAgeSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"FamilyEventDetail", cHUSB, ErrSingleMultiple}
			}
			HusbandAgeSet = true
			if err := s.HusbandAge.parse(&sl, o); err != nil {
				return ErrContext{"FamilyEventDetail", cHUSB, err}
			}
			l.Sub = append(l.Sub[:i], l.Sub[i+1:]...)
			i--
		case cWIFE:
			if WifeAgeSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"FamilyEventDetail", cWIFE, ErrSingleMultiple}
			}
			WifeAgeSet = true
			if err := s.WifeAge.parse(&sl, o); err != nil {
				return ErrContext{"FamilyEventDetail", cWIFE, err}
			}
			l.Sub = append(l.Sub[:i], l.Sub[i+1:]...)
			i--
		}
	}
	return s.EventDetail.parse(l, o)
}

// AgeStructure is a GEDCOM structure type
type AgeStructure struct {
	Age AgeAtEvent
}

func (s *AgeStructure) parse(l *Line, o options) error {
	var AgeSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case cAGE:
			if AgeSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"AgeStructure", cAGE, ErrSingleMultiple}
			}
			AgeSet = true
			if err := s.Age.parse(&sl, o); err != nil {
				return ErrContext{"AgeStructure", cAGE, err}
			}
		default:
			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
				return ErrContext{"AgeStructure", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	if !o.allowMissingRequired {
		if !AgeSet {
			return ErrContext{"AgeStructure", "Age", ErrRequiredMissing}
		}
	}
	return nil
}

// Individual is a GEDCOM structure type
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

func (s *Individual) parse(l *Line, o options) error {
	if err := s.ID.parse(&Line{line: line{value: l.xrefID}}, o); err != nil {
		return ErrContext{"Individual", "xrefID", err}
	}
	var RestrictionNoticeSet, GenderSet, BirthSet, ChristeningSet, DeathSet, BuriedSet, CremationSet, AdoptionSet, MaptismSet, BarMitzvahSet, BasMitzvahSet, BlessingSet, AdultChristeningSet, ConfirmationSet, FirstCommunionSet, OrdinationSet, NaturalizationSet, EmigratedSet, ImmigratedSet, CensusSet, ProbateSet, WillSet, GraduatedSet, RetiredSet, PermanentRecordSet, AncestralFileNumberSet, AutomatedRecordIDSet, ChangeDateSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case cRESN:
			if RestrictionNoticeSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Individual", cRESN, ErrSingleMultiple}
			}
			RestrictionNoticeSet = true
			if err := s.RestrictionNotice.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cRESN, err}
			}
		case cNAME:
			var t PersonalNameStructure
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cNAME, err}
			}
			s.PersonalNameStructure = append(s.PersonalNameStructure, t)
		case cSEX:
			if GenderSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Individual", cSEX, ErrSingleMultiple}
			}
			GenderSet = true
			if err := s.Gender.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cSEX, err}
			}
		case cBIRT:
			if BirthSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Individual", cBIRT, ErrSingleMultiple}
			}
			BirthSet = true
			if err := s.Birth.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cBIRT, err}
			}
		case cCHR:
			if ChristeningSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Individual", cCHR, ErrSingleMultiple}
			}
			ChristeningSet = true
			if err := s.Christening.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cCHR, err}
			}
		case cDEAT:
			if DeathSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Individual", cDEAT, ErrSingleMultiple}
			}
			DeathSet = true
			if err := s.Death.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cDEAT, err}
			}
		case cBURI:
			if BuriedSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Individual", cBURI, ErrSingleMultiple}
			}
			BuriedSet = true
			if err := s.Buried.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cBURI, err}
			}
		case cCREM:
			if CremationSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Individual", cCREM, ErrSingleMultiple}
			}
			CremationSet = true
			if err := s.Cremation.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cCREM, err}
			}
		case cADOP:
			if AdoptionSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Individual", cADOP, ErrSingleMultiple}
			}
			AdoptionSet = true
			if err := s.Adoption.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cADOP, err}
			}
		case cBAPM:
			if MaptismSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Individual", cBAPM, ErrSingleMultiple}
			}
			MaptismSet = true
			if err := s.Maptism.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cBAPM, err}
			}
		case cBARM:
			if BarMitzvahSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Individual", cBARM, ErrSingleMultiple}
			}
			BarMitzvahSet = true
			if err := s.BarMitzvah.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cBARM, err}
			}
		case cBASM:
			if BasMitzvahSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Individual", cBASM, ErrSingleMultiple}
			}
			BasMitzvahSet = true
			if err := s.BasMitzvah.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cBASM, err}
			}
		case cBLES:
			if BlessingSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Individual", cBLES, ErrSingleMultiple}
			}
			BlessingSet = true
			if err := s.Blessing.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cBLES, err}
			}
		case cCHRA:
			if AdultChristeningSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Individual", cCHRA, ErrSingleMultiple}
			}
			AdultChristeningSet = true
			if err := s.AdultChristening.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cCHRA, err}
			}
		case cCONF:
			if ConfirmationSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Individual", cCONF, ErrSingleMultiple}
			}
			ConfirmationSet = true
			if err := s.Confirmation.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cCONF, err}
			}
		case cFCOM:
			if FirstCommunionSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Individual", cFCOM, ErrSingleMultiple}
			}
			FirstCommunionSet = true
			if err := s.FirstCommunion.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cFCOM, err}
			}
		case cORDN:
			if OrdinationSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Individual", cORDN, ErrSingleMultiple}
			}
			OrdinationSet = true
			if err := s.Ordination.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cORDN, err}
			}
		case cNATU:
			if NaturalizationSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Individual", cNATU, ErrSingleMultiple}
			}
			NaturalizationSet = true
			if err := s.Naturalization.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cNATU, err}
			}
		case cEMIG:
			if EmigratedSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Individual", cEMIG, ErrSingleMultiple}
			}
			EmigratedSet = true
			if err := s.Emigrated.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cEMIG, err}
			}
		case cIMMI:
			if ImmigratedSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Individual", cIMMI, ErrSingleMultiple}
			}
			ImmigratedSet = true
			if err := s.Immigrated.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cIMMI, err}
			}
		case cCENS:
			if CensusSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Individual", cCENS, ErrSingleMultiple}
			}
			CensusSet = true
			if err := s.Census.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cCENS, err}
			}
		case cPROB:
			if ProbateSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Individual", cPROB, ErrSingleMultiple}
			}
			ProbateSet = true
			if err := s.Probate.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cPROB, err}
			}
		case cWILL:
			if WillSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Individual", cWILL, ErrSingleMultiple}
			}
			WillSet = true
			if err := s.Will.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cWILL, err}
			}
		case cGRAD:
			if GraduatedSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Individual", cGRAD, ErrSingleMultiple}
			}
			GraduatedSet = true
			if err := s.Graduated.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cGRAD, err}
			}
		case cRETI:
			if RetiredSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Individual", cRETI, ErrSingleMultiple}
			}
			RetiredSet = true
			if err := s.Retired.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cRETI, err}
			}
		case cEVEN:
			var t IndividualEventDetail
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cEVEN, err}
			}
			s.Events = append(s.Events, t)
		case cCAST:
			var t CasteEvent
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cCAST, err}
			}
			s.Caste = append(s.Caste, t)
		case cDSCR:
			var t DescriptionEvent
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cDSCR, err}
			}
			s.Description = append(s.Description, t)
		case cEDUC:
			var t ScholasticEvent
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cEDUC, err}
			}
			s.ScholasticAchievement = append(s.ScholasticAchievement, t)
		case cIDNO:
			var t NationalIDEvent
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cIDNO, err}
			}
			s.NationalID = append(s.NationalID, t)
		case cNATI:
			var t NationalOriginEvent
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cNATI, err}
			}
			s.NationalTribalOrigin = append(s.NationalTribalOrigin, t)
		case cNCHI:
			var t ChildrenEvent
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cNCHI, err}
			}
			s.CountOfChildren = append(s.CountOfChildren, t)
		case cNMR:
			var t MarriagesEvent
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cNMR, err}
			}
			s.CountOfMarriages = append(s.CountOfMarriages, t)
		case cOCCU:
			var t OccupationEvent
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cOCCU, err}
			}
			s.Occupation = append(s.Occupation, t)
		case cPROP:
			var t PossessionEvent
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cPROP, err}
			}
			s.Possessions = append(s.Possessions, t)
		case cRELI:
			var t ReligiousEvent
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cRELI, err}
			}
			s.ReligiousAffiliation = append(s.ReligiousAffiliation, t)
		case cRESI:
			var t ResidenceEvent
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cRESI, err}
			}
			s.Residences = append(s.Residences, t)
		case cSSN:
			var t SSNEvent
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cSSN, err}
			}
			s.SocialSecurity = append(s.SocialSecurity, t)
		case cTITL:
			var t NobilityEvent
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cTITL, err}
			}
			s.NobilityTypeTitle = append(s.NobilityTypeTitle, t)
		case cFAMC:
			var t ChildToFamilyLink
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cFAMC, err}
			}
			s.ChildOf = append(s.ChildOf, t)
		case cFAMS:
			var t SpouseToFamilyLink
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cFAMS, err}
			}
			s.SpouseOf = append(s.SpouseOf, t)
		case cSUBM:
			var t Xref
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cSUBM, err}
			}
			s.Submitters = append(s.Submitters, t)
		case cASSO:
			var t AssociationStructure
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cASSO, err}
			}
			s.Associations = append(s.Associations, t)
		case cALIA:
			var t Xref
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cALIA, err}
			}
			s.Aliases = append(s.Aliases, t)
		case cANCI:
			var t Xref
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cANCI, err}
			}
			s.AncestorInterest = append(s.AncestorInterest, t)
		case cDESI:
			var t Xref
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cDESI, err}
			}
			s.DescendentInterest = append(s.DescendentInterest, t)
		case cSOUR:
			var t SourceCitation
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cSOUR, err}
			}
			s.Sources = append(s.Sources, t)
		case cOBJE:
			var t MultimediaLink
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cOBJE, err}
			}
			s.Multimedia = append(s.Multimedia, t)
		case cNOTE:
			var t NoteStructure
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cNOTE, err}
			}
			s.Notes = append(s.Notes, t)
		case cRFN:
			if PermanentRecordSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Individual", cRFN, ErrSingleMultiple}
			}
			PermanentRecordSet = true
			if err := s.PermanentRecord.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cRFN, err}
			}
		case cAFN:
			if AncestralFileNumberSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Individual", cAFN, ErrSingleMultiple}
			}
			AncestralFileNumberSet = true
			if err := s.AncestralFileNumber.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cAFN, err}
			}
		case cREFN:
			var t UserReferenceStructure
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cREFN, err}
			}
			s.UserReferences = append(s.UserReferences, t)
		case cRIN:
			if AutomatedRecordIDSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Individual", cRIN, ErrSingleMultiple}
			}
			AutomatedRecordIDSet = true
			if err := s.AutomatedRecordID.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cRIN, err}
			}
		case cCHAN:
			if ChangeDateSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"Individual", cCHAN, ErrSingleMultiple}
			}
			ChangeDateSet = true
			if err := s.ChangeDate.parse(&sl, o); err != nil {
				return ErrContext{"Individual", cCHAN, err}
			}
		default:
			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
				return ErrContext{"Individual", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// VerifiedIndividualFamEventDetail is a GEDCOM structure type
type VerifiedIndividualFamEventDetail struct {
	Famc Xref
	VerifiedEventDetail
}

func (s *VerifiedIndividualFamEventDetail) parse(l *Line, o options) error {
	var FamcSet bool
	for i := 0; i < len(l.Sub); i++ {
		sl := l.Sub[i]
		switch sl.tag {
		case cFAMC:
			if FamcSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"VerifiedIndividualFamEventDetail", cFAMC, ErrSingleMultiple}
			}
			FamcSet = true
			if err := s.Famc.parse(&sl, o); err != nil {
				return ErrContext{"VerifiedIndividualFamEventDetail", cFAMC, err}
			}
			l.Sub = append(l.Sub[:i], l.Sub[i+1:]...)
			i--
		}
	}
	return s.VerifiedEventDetail.parse(l, o)
}

// VerifiedEventDetail is a GEDCOM structure type
type VerifiedEventDetail struct {
	Verified Verified
	EventDetail
}

func (s *VerifiedEventDetail) parse(l *Line, o options) error {
	if err := s.Verified.parse(l, o); err != nil {
		return ErrContext{"VerifiedEventDetail", "line_value", err}
	}
	return s.EventDetail.parse(l, o)
}

// IndividualEventDetail is a GEDCOM structure type
type IndividualEventDetail struct {
	EventDetail
}

func (s *IndividualEventDetail) parse(l *Line, o options) error {
	return s.EventDetail.parse(l, o)
}

// AdoptionEvent is a GEDCOM structure type
type AdoptionEvent struct {
	Family AdoptionReference
	EventDetail
}

func (s *AdoptionEvent) parse(l *Line, o options) error {
	var FamilySet bool
	for i := 0; i < len(l.Sub); i++ {
		sl := l.Sub[i]
		switch sl.tag {
		case cFAMC:
			if FamilySet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"AdoptionEvent", cFAMC, ErrSingleMultiple}
			}
			FamilySet = true
			if err := s.Family.parse(&sl, o); err != nil {
				return ErrContext{"AdoptionEvent", cFAMC, err}
			}
			l.Sub = append(l.Sub[:i], l.Sub[i+1:]...)
			i--
		}
	}
	return s.EventDetail.parse(l, o)
}

// AdoptionReference is a GEDCOM structure type
type AdoptionReference struct {
	ID        Xref
	AdoptedBy AdoptedBy
}

func (s *AdoptionReference) parse(l *Line, o options) error {
	if err := s.ID.parse(&Line{line: line{value: l.xrefID}}, o); err != nil {
		return ErrContext{"AdoptionReference", "xrefID", err}
	}
	var AdoptedBySet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case cADOP:
			if AdoptedBySet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"AdoptionReference", cADOP, ErrSingleMultiple}
			}
			AdoptedBySet = true
			if err := s.AdoptedBy.parse(&sl, o); err != nil {
				return ErrContext{"AdoptionReference", cADOP, err}
			}
		default:
			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
				return ErrContext{"AdoptionReference", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// CasteEvent is a GEDCOM structure type
type CasteEvent struct {
	CasteName CasteName
	EventDetail
}

func (s *CasteEvent) parse(l *Line, o options) error {
	if err := s.CasteName.parse(l, o); err != nil {
		return ErrContext{"CasteEvent", "line_value", err}
	}
	return s.EventDetail.parse(l, o)
}

// DescriptionEvent is a GEDCOM structure type
type DescriptionEvent struct {
	PhysicalDescription PhysicalDescription
	EventDetail
}

func (s *DescriptionEvent) parse(l *Line, o options) error {
	if err := s.PhysicalDescription.parse(l, o); err != nil {
		return ErrContext{"DescriptionEvent", "line_value", err}
	}
	return s.EventDetail.parse(l, o)
}

// ScholasticEvent is a GEDCOM structure type
type ScholasticEvent struct {
	ScholasticAchievement ScholasticAchievement
	EventDetail
}

func (s *ScholasticEvent) parse(l *Line, o options) error {
	if err := s.ScholasticAchievement.parse(l, o); err != nil {
		return ErrContext{"ScholasticEvent", "line_value", err}
	}
	return s.EventDetail.parse(l, o)
}

// NationalIDEvent is a GEDCOM structure type
type NationalIDEvent struct {
	NationalIDNumber NationalIDNumber
	EventDetail
}

func (s *NationalIDEvent) parse(l *Line, o options) error {
	if err := s.NationalIDNumber.parse(l, o); err != nil {
		return ErrContext{"NationalIDEvent", "line_value", err}
	}
	return s.EventDetail.parse(l, o)
}

// NationalOriginEvent is a GEDCOM structure type
type NationalOriginEvent struct {
	NationalOrTribalOrigin NationalOrTribalOrigin
	EventDetail
}

func (s *NationalOriginEvent) parse(l *Line, o options) error {
	if err := s.NationalOrTribalOrigin.parse(l, o); err != nil {
		return ErrContext{"NationalOriginEvent", "line_value", err}
	}
	return s.EventDetail.parse(l, o)
}

// ChildrenEvent is a GEDCOM structure type
type ChildrenEvent struct {
	CountOfChildren CountOfChildren
	EventDetail
}

func (s *ChildrenEvent) parse(l *Line, o options) error {
	if err := s.CountOfChildren.parse(l, o); err != nil {
		return ErrContext{"ChildrenEvent", "line_value", err}
	}
	return s.EventDetail.parse(l, o)
}

// MarriagesEvent is a GEDCOM structure type
type MarriagesEvent struct {
	CountOfMarriages CountOfMarriages
	EventDetail
}

func (s *MarriagesEvent) parse(l *Line, o options) error {
	if err := s.CountOfMarriages.parse(l, o); err != nil {
		return ErrContext{"MarriagesEvent", "line_value", err}
	}
	return s.EventDetail.parse(l, o)
}

// OccupationEvent is a GEDCOM structure type
type OccupationEvent struct {
	Occupation Occupation
	EventDetail
}

func (s *OccupationEvent) parse(l *Line, o options) error {
	if err := s.Occupation.parse(l, o); err != nil {
		return ErrContext{"OccupationEvent", "line_value", err}
	}
	return s.EventDetail.parse(l, o)
}

// PossessionEvent is a GEDCOM structure type
type PossessionEvent struct {
	Possessions Possessions
	EventDetail
}

func (s *PossessionEvent) parse(l *Line, o options) error {
	if err := s.Possessions.parse(l, o); err != nil {
		return ErrContext{"PossessionEvent", "line_value", err}
	}
	return s.EventDetail.parse(l, o)
}

// ReligiousEvent is a GEDCOM structure type
type ReligiousEvent struct {
	ReligiousAffiliation ReligiousAffiliation
	EventDetail
}

func (s *ReligiousEvent) parse(l *Line, o options) error {
	if err := s.ReligiousAffiliation.parse(l, o); err != nil {
		return ErrContext{"ReligiousEvent", "line_value", err}
	}
	return s.EventDetail.parse(l, o)
}

// ResidenceEvent is a GEDCOM structure type
type ResidenceEvent struct {
	EventDetail
}

func (s *ResidenceEvent) parse(l *Line, o options) error {
	return s.EventDetail.parse(l, o)
}

// SSNEvent is a GEDCOM structure type
type SSNEvent struct {
	SocialSecurityNumber SocialSecurityNumber
	EventDetail
}

func (s *SSNEvent) parse(l *Line, o options) error {
	if err := s.SocialSecurityNumber.parse(l, o); err != nil {
		return ErrContext{"SSNEvent", "line_value", err}
	}
	return s.EventDetail.parse(l, o)
}

// NobilityEvent is a GEDCOM structure type
type NobilityEvent struct {
	NobilityTypeTitle NobilityTypeTitle
	EventDetail
}

func (s *NobilityEvent) parse(l *Line, o options) error {
	if err := s.NobilityTypeTitle.parse(l, o); err != nil {
		return ErrContext{"NobilityEvent", "line_value", err}
	}
	return s.EventDetail.parse(l, o)
}

// UserReferenceStructure is a GEDCOM structure type
type UserReferenceStructure struct {
	UserReferenceNumber UserReferenceNumber
	Type                UserReferenceType
}

func (s *UserReferenceStructure) parse(l *Line, o options) error {
	if err := s.UserReferenceNumber.parse(l, o); err != nil {
		return ErrContext{"UserReferenceStructure", "line_value", err}
	}
	var TypeSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case cTYPE:
			if TypeSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"UserReferenceStructure", cTYPE, ErrSingleMultiple}
			}
			TypeSet = true
			if err := s.Type.parse(&sl, o); err != nil {
				return ErrContext{"UserReferenceStructure", cTYPE, err}
			}
		default:
			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
				return ErrContext{"UserReferenceStructure", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// MultimediaRecord is a GEDCOM structure type
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

func (s *MultimediaRecord) parse(l *Line, o options) error {
	if err := s.ID.parse(&Line{line: line{value: l.xrefID}}, o); err != nil {
		return ErrContext{"MultimediaRecord", "xrefID", err}
	}
	var FormatSet, BlobSet, TitleSet, ContinuedObjectSet, AutomatedRecordIDSet, ChangeDateSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case cFORM:
			if FormatSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"MultimediaRecord", cFORM, ErrSingleMultiple}
			}
			FormatSet = true
			if err := s.Format.parse(&sl, o); err != nil {
				return ErrContext{"MultimediaRecord", cFORM, err}
			}
		case cTITLE:
			if TitleSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"MultimediaRecord", cTITLE, ErrSingleMultiple}
			}
			TitleSet = true
			if err := s.Title.parse(&sl, o); err != nil {
				return ErrContext{"MultimediaRecord", cTITLE, err}
			}
		case cNOTE:
			var t NoteStructure
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"MultimediaRecord", cNOTE, err}
			}
			s.Notes = append(s.Notes, t)
		case cBLOB:
			if BlobSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"MultimediaRecord", cBLOB, ErrSingleMultiple}
			}
			BlobSet = true
			if err := s.Blob.parse(&sl, o); err != nil {
				return ErrContext{"MultimediaRecord", cBLOB, err}
			}
		case cOBJE:
			if ContinuedObjectSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"MultimediaRecord", cOBJE, ErrSingleMultiple}
			}
			ContinuedObjectSet = true
			if err := s.ContinuedObject.parse(&sl, o); err != nil {
				return ErrContext{"MultimediaRecord", cOBJE, err}
			}
		case cREFN:
			var t UserReferenceStructure
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"MultimediaRecord", cREFN, err}
			}
			s.UserReferences = append(s.UserReferences, t)
		case cRIN:
			if AutomatedRecordIDSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"MultimediaRecord", cRIN, ErrSingleMultiple}
			}
			AutomatedRecordIDSet = true
			if err := s.AutomatedRecordID.parse(&sl, o); err != nil {
				return ErrContext{"MultimediaRecord", cRIN, err}
			}
		case cCHAN:
			if ChangeDateSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"MultimediaRecord", cCHAN, ErrSingleMultiple}
			}
			ChangeDateSet = true
			if err := s.ChangeDate.parse(&sl, o); err != nil {
				return ErrContext{"MultimediaRecord", cCHAN, err}
			}
		default:
			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
				return ErrContext{"MultimediaRecord", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	if !o.allowMissingRequired {
		if !FormatSet {
			return ErrContext{"MultimediaRecord", "Format", ErrRequiredMissing}
		}
		if !BlobSet {
			return ErrContext{"MultimediaRecord", "Blob", ErrRequiredMissing}
		}
	}
	return nil
}

// NoteRecord is a GEDCOM structure type
type NoteRecord struct {
	ID                Xref
	SubmitterText     SubmitterText
	Sources           []SourceCitation
	UserReferences    []UserReferenceStructure
	AutomatedRecordID AutomatedRecordID
	ChangeDate        ChangeDateStructure
}

func (s *NoteRecord) parse(l *Line, o options) error {
	if err := s.ID.parse(&Line{line: line{value: l.xrefID}}, o); err != nil {
		return ErrContext{"NoteRecord", "xrefID", err}
	}
	if err := s.SubmitterText.parse(l, o); err != nil {
		return ErrContext{"NoteRecord", "line_value", err}
	}
	var AutomatedRecordIDSet, ChangeDateSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case cSOUR:
			var t SourceCitation
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"NoteRecord", cSOUR, err}
			}
			s.Sources = append(s.Sources, t)
		case cREFN:
			var t UserReferenceStructure
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"NoteRecord", cREFN, err}
			}
			s.UserReferences = append(s.UserReferences, t)
		case cRIN:
			if AutomatedRecordIDSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"NoteRecord", cRIN, ErrSingleMultiple}
			}
			AutomatedRecordIDSet = true
			if err := s.AutomatedRecordID.parse(&sl, o); err != nil {
				return ErrContext{"NoteRecord", cRIN, err}
			}
		case cCHAN:
			if ChangeDateSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"NoteRecord", cCHAN, ErrSingleMultiple}
			}
			ChangeDateSet = true
			if err := s.ChangeDate.parse(&sl, o); err != nil {
				return ErrContext{"NoteRecord", cCHAN, err}
			}
		default:
			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
				return ErrContext{"NoteRecord", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// RepositoryRecord is a GEDCOM structure type
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

func (s *RepositoryRecord) parse(l *Line, o options) error {
	if err := s.ID.parse(&Line{line: line{value: l.xrefID}}, o); err != nil {
		return ErrContext{"RepositoryRecord", "xrefID", err}
	}
	var NameOfRepositorySet, AddressSet, AutomatedRecordIDSet, ChangeDateSet bool
	s.PhoneNumber = make([]PhoneNumber, 0, 3)
	for _, sl := range l.Sub {
		switch sl.tag {
		case cNAME:
			if NameOfRepositorySet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"RepositoryRecord", cNAME, ErrSingleMultiple}
			}
			NameOfRepositorySet = true
			if err := s.NameOfRepository.parse(&sl, o); err != nil {
				return ErrContext{"RepositoryRecord", cNAME, err}
			}
		case cADDR:
			if AddressSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"RepositoryRecord", cADDR, ErrSingleMultiple}
			}
			AddressSet = true
			if err := s.Address.parse(&sl, o); err != nil {
				return ErrContext{"RepositoryRecord", cADDR, err}
			}
		case cPHON:
			if len(s.PhoneNumber) == 3 {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"RepositoryRecord", cPHON, ErrTooMany(3)}
			}
			var t PhoneNumber
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"RepositoryRecord", cPHON, err}
			}
			s.PhoneNumber = append(s.PhoneNumber, t)
		case cNOTE:
			var t NoteStructure
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"RepositoryRecord", cNOTE, err}
			}
			s.Notes = append(s.Notes, t)
		case cREFN:
			var t UserReferenceStructure
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"RepositoryRecord", cREFN, err}
			}
			s.UserReferences = append(s.UserReferences, t)
		case cRIN:
			if AutomatedRecordIDSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"RepositoryRecord", cRIN, ErrSingleMultiple}
			}
			AutomatedRecordIDSet = true
			if err := s.AutomatedRecordID.parse(&sl, o); err != nil {
				return ErrContext{"RepositoryRecord", cRIN, err}
			}
		case cCHAN:
			if ChangeDateSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"RepositoryRecord", cCHAN, ErrSingleMultiple}
			}
			ChangeDateSet = true
			if err := s.ChangeDate.parse(&sl, o); err != nil {
				return ErrContext{"RepositoryRecord", cCHAN, err}
			}
		default:
			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
				return ErrContext{"RepositoryRecord", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// SourceRecord is a GEDCOM structure type
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

func (s *SourceRecord) parse(l *Line, o options) error {
	if err := s.ID.parse(&Line{line: line{value: l.xrefID}}, o); err != nil {
		return ErrContext{"SourceRecord", "xrefID", err}
	}
	var SourceRepositoryCitationSet, DataSet, OriginatorSet, TitleSet, FiledBySet, PublicationFactsSet, TextFromSourceSet, ContinuedObjectSet, AutomatedRecordIDSet, ChangeDateSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case cDATA:
			if DataSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"SourceRecord", cDATA, ErrSingleMultiple}
			}
			DataSet = true
			if err := s.Data.parse(&sl, o); err != nil {
				return ErrContext{"SourceRecord", cDATA, err}
			}
		case cAUTH:
			if OriginatorSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"SourceRecord", cAUTH, ErrSingleMultiple}
			}
			OriginatorSet = true
			if err := s.Originator.parse(&sl, o); err != nil {
				return ErrContext{"SourceRecord", cAUTH, err}
			}
		case cTITL:
			if TitleSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"SourceRecord", cTITL, ErrSingleMultiple}
			}
			TitleSet = true
			if err := s.Title.parse(&sl, o); err != nil {
				return ErrContext{"SourceRecord", cTITL, err}
			}
		case cABBR:
			if FiledBySet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"SourceRecord", cABBR, ErrSingleMultiple}
			}
			FiledBySet = true
			if err := s.FiledBy.parse(&sl, o); err != nil {
				return ErrContext{"SourceRecord", cABBR, err}
			}
		case cPUBL:
			if PublicationFactsSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"SourceRecord", cPUBL, ErrSingleMultiple}
			}
			PublicationFactsSet = true
			if err := s.PublicationFacts.parse(&sl, o); err != nil {
				return ErrContext{"SourceRecord", cPUBL, err}
			}
		case cTEXT:
			if TextFromSourceSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"SourceRecord", cTEXT, ErrSingleMultiple}
			}
			TextFromSourceSet = true
			if err := s.TextFromSource.parse(&sl, o); err != nil {
				return ErrContext{"SourceRecord", cTEXT, err}
			}
		case cREPO:
			if SourceRepositoryCitationSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"SourceRecord", cREPO, ErrSingleMultiple}
			}
			SourceRepositoryCitationSet = true
			if err := s.SourceRepositoryCitation.parse(&sl, o); err != nil {
				return ErrContext{"SourceRecord", cREPO, err}
			}
		case cOBJE:
			if ContinuedObjectSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"SourceRecord", cOBJE, ErrSingleMultiple}
			}
			ContinuedObjectSet = true
			if err := s.ContinuedObject.parse(&sl, o); err != nil {
				return ErrContext{"SourceRecord", cOBJE, err}
			}
		case cNOTE:
			var t NoteStructure
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"SourceRecord", cNOTE, err}
			}
			s.Notes = append(s.Notes, t)
		case cREFN:
			var t UserReferenceStructure
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"SourceRecord", cREFN, err}
			}
			s.UserReferences = append(s.UserReferences, t)
		case cRIN:
			if AutomatedRecordIDSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"SourceRecord", cRIN, ErrSingleMultiple}
			}
			AutomatedRecordIDSet = true
			if err := s.AutomatedRecordID.parse(&sl, o); err != nil {
				return ErrContext{"SourceRecord", cRIN, err}
			}
		case cCHAN:
			if ChangeDateSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"SourceRecord", cCHAN, ErrSingleMultiple}
			}
			ChangeDateSet = true
			if err := s.ChangeDate.parse(&sl, o); err != nil {
				return ErrContext{"SourceRecord", cCHAN, err}
			}
		default:
			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
				return ErrContext{"SourceRecord", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	if !o.allowMissingRequired {
		if !SourceRepositoryCitationSet {
			return ErrContext{"SourceRecord", "SourceRepositoryCitation", ErrRequiredMissing}
		}
	}
	return nil
}

// SourceRecordDataStructure is a GEDCOM structure type
type SourceRecordDataStructure struct {
	EventsRecorded    []EventsRecordedStructure
	ResponsibleAgency ResponsibleAgency
	Notes             []NoteStructure
}

func (s *SourceRecordDataStructure) parse(l *Line, o options) error {
	var ResponsibleAgencySet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case cEVEN:
			var t EventsRecordedStructure
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"SourceRecordDataStructure", cEVEN, err}
			}
			s.EventsRecorded = append(s.EventsRecorded, t)
		case cAGNC:
			if ResponsibleAgencySet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"SourceRecordDataStructure", cAGNC, ErrSingleMultiple}
			}
			ResponsibleAgencySet = true
			if err := s.ResponsibleAgency.parse(&sl, o); err != nil {
				return ErrContext{"SourceRecordDataStructure", cAGNC, err}
			}
		case cNOTE:
			var t NoteStructure
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"SourceRecordDataStructure", cNOTE, err}
			}
			s.Notes = append(s.Notes, t)
		default:
			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
				return ErrContext{"SourceRecordDataStructure", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// EventsRecordedStructure is a GEDCOM structure type
type EventsRecordedStructure struct {
	DatePeriod              DatePeriod
	SourceJurisdictionPlace SourceJurisdictionPlace
}

func (s *EventsRecordedStructure) parse(l *Line, o options) error {
	var DatePeriodSet, SourceJurisdictionPlaceSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case cDATE:
			if DatePeriodSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"EventsRecordedStructure", cDATE, ErrSingleMultiple}
			}
			DatePeriodSet = true
			if err := s.DatePeriod.parse(&sl, o); err != nil {
				return ErrContext{"EventsRecordedStructure", cDATE, err}
			}
		case cPLACE:
			if SourceJurisdictionPlaceSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"EventsRecordedStructure", cPLACE, ErrSingleMultiple}
			}
			SourceJurisdictionPlaceSet = true
			if err := s.SourceJurisdictionPlace.parse(&sl, o); err != nil {
				return ErrContext{"EventsRecordedStructure", cPLACE, err}
			}
		default:
			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
				return ErrContext{"EventsRecordedStructure", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// SubmissionRecord is a GEDCOM structure type
type SubmissionRecord struct {
	ID                       Xref
	NameOfFamilyFile         NameOfFamilyFile
	TempleCode               TempleCode
	GenerationsOfAncestors   GenerationsOfAncestors
	GenerationsOfDescendants GenerationsOfDescendants
	OrdinanceProcessFlag     OrdinanceProcessFlag
	AutomatedRecordID        AutomatedRecordID
}

func (s *SubmissionRecord) parse(l *Line, o options) error {
	if err := s.ID.parse(&Line{line: line{value: l.xrefID}}, o); err != nil {
		return ErrContext{"SubmissionRecord", "xrefID", err}
	}
	var NameOfFamilyFileSet, TempleCodeSet, GenerationsOfAncestorsSet, GenerationsOfDescendantsSet, OrdinanceProcessFlagSet, AutomatedRecordIDSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case cFAMF:
			if NameOfFamilyFileSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"SubmissionRecord", cFAMF, ErrSingleMultiple}
			}
			NameOfFamilyFileSet = true
			if err := s.NameOfFamilyFile.parse(&sl, o); err != nil {
				return ErrContext{"SubmissionRecord", cFAMF, err}
			}
		case cTEMP:
			if TempleCodeSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"SubmissionRecord", cTEMP, ErrSingleMultiple}
			}
			TempleCodeSet = true
			if err := s.TempleCode.parse(&sl, o); err != nil {
				return ErrContext{"SubmissionRecord", cTEMP, err}
			}
		case cANCE:
			if GenerationsOfAncestorsSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"SubmissionRecord", cANCE, ErrSingleMultiple}
			}
			GenerationsOfAncestorsSet = true
			if err := s.GenerationsOfAncestors.parse(&sl, o); err != nil {
				return ErrContext{"SubmissionRecord", cANCE, err}
			}
		case cDESC:
			if GenerationsOfDescendantsSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"SubmissionRecord", cDESC, ErrSingleMultiple}
			}
			GenerationsOfDescendantsSet = true
			if err := s.GenerationsOfDescendants.parse(&sl, o); err != nil {
				return ErrContext{"SubmissionRecord", cDESC, err}
			}
		case cORDI:
			if OrdinanceProcessFlagSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"SubmissionRecord", cORDI, ErrSingleMultiple}
			}
			OrdinanceProcessFlagSet = true
			if err := s.OrdinanceProcessFlag.parse(&sl, o); err != nil {
				return ErrContext{"SubmissionRecord", cORDI, err}
			}
		case cRIN:
			if AutomatedRecordIDSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"SubmissionRecord", cRIN, ErrSingleMultiple}
			}
			AutomatedRecordIDSet = true
			if err := s.AutomatedRecordID.parse(&sl, o); err != nil {
				return ErrContext{"SubmissionRecord", cRIN, err}
			}
		default:
			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
				return ErrContext{"SubmissionRecord", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// SubmitterRecord is a GEDCOM structure type
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

func (s *SubmitterRecord) parse(l *Line, o options) error {
	if err := s.ID.parse(&Line{line: line{value: l.xrefID}}, o); err != nil {
		return ErrContext{"SubmitterRecord", "xrefID", err}
	}
	var SubmitterNameSet, AddressSet, SubmitterRegisteredRFNSet, AutomatedRecordIDSet, ChangeDateSet bool
	s.PhoneNumber = make([]PhoneNumber, 0, 3)
	s.LanguagePreference = make([]LanguagePreference, 0, 3)
	for _, sl := range l.Sub {
		switch sl.tag {
		case cNAME:
			if SubmitterNameSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"SubmitterRecord", cNAME, ErrSingleMultiple}
			}
			SubmitterNameSet = true
			if err := s.SubmitterName.parse(&sl, o); err != nil {
				return ErrContext{"SubmitterRecord", cNAME, err}
			}
		case cADDR:
			if AddressSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"SubmitterRecord", cADDR, ErrSingleMultiple}
			}
			AddressSet = true
			if err := s.Address.parse(&sl, o); err != nil {
				return ErrContext{"SubmitterRecord", cADDR, err}
			}
		case cPHON:
			if len(s.PhoneNumber) == 3 {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"SubmitterRecord", cPHON, ErrTooMany(3)}
			}
			var t PhoneNumber
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"SubmitterRecord", cPHON, err}
			}
			s.PhoneNumber = append(s.PhoneNumber, t)
		case cOBJE:
			var t MultimediaLink
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"SubmitterRecord", cOBJE, err}
			}
			s.Multimedia = append(s.Multimedia, t)
		case cLANG:
			if len(s.LanguagePreference) == 3 {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"SubmitterRecord", cLANG, ErrTooMany(3)}
			}
			var t LanguagePreference
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"SubmitterRecord", cLANG, err}
			}
			s.LanguagePreference = append(s.LanguagePreference, t)
		case cRFN:
			if SubmitterRegisteredRFNSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"SubmitterRecord", cRFN, ErrSingleMultiple}
			}
			SubmitterRegisteredRFNSet = true
			if err := s.SubmitterRegisteredRFN.parse(&sl, o); err != nil {
				return ErrContext{"SubmitterRecord", cRFN, err}
			}
		case cRIN:
			if AutomatedRecordIDSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"SubmitterRecord", cRIN, ErrSingleMultiple}
			}
			AutomatedRecordIDSet = true
			if err := s.AutomatedRecordID.parse(&sl, o); err != nil {
				return ErrContext{"SubmitterRecord", cRIN, err}
			}
		case cCHAN:
			if ChangeDateSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"SubmitterRecord", cCHAN, ErrSingleMultiple}
			}
			ChangeDateSet = true
			if err := s.ChangeDate.parse(&sl, o); err != nil {
				return ErrContext{"SubmitterRecord", cCHAN, err}
			}
		default:
			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
				return ErrContext{"SubmitterRecord", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	if !o.allowMissingRequired {
		if !SubmitterNameSet {
			return ErrContext{"SubmitterRecord", "SubmitterName", ErrRequiredMissing}
		}
	}
	return nil
}

// AddressStructure is a GEDCOM structure type
type AddressStructure struct {
	AddressLine  AddressLine
	AddressLine1 AddressLine1
	AddressLine2 AddressLine2
	City         AddressCity
	State        AddressState
	PostalCode   AddressPostalCode
	Country      AddressCountry
}

func (s *AddressStructure) parse(l *Line, o options) error {
	if err := s.AddressLine.parse(l, o); err != nil {
		return ErrContext{"AddressStructure", "line_value", err}
	}
	var AddressLine1Set, AddressLine2Set, CitySet, StateSet, PostalCodeSet, CountrySet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case cADR1:
			if AddressLine1Set {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"AddressStructure", cADR1, ErrSingleMultiple}
			}
			AddressLine1Set = true
			if err := s.AddressLine1.parse(&sl, o); err != nil {
				return ErrContext{"AddressStructure", cADR1, err}
			}
		case cADR2:
			if AddressLine2Set {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"AddressStructure", cADR2, ErrSingleMultiple}
			}
			AddressLine2Set = true
			if err := s.AddressLine2.parse(&sl, o); err != nil {
				return ErrContext{"AddressStructure", cADR2, err}
			}
		case cCITY:
			if CitySet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"AddressStructure", cCITY, ErrSingleMultiple}
			}
			CitySet = true
			if err := s.City.parse(&sl, o); err != nil {
				return ErrContext{"AddressStructure", cCITY, err}
			}
		case cSTAE:
			if StateSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"AddressStructure", cSTAE, ErrSingleMultiple}
			}
			StateSet = true
			if err := s.State.parse(&sl, o); err != nil {
				return ErrContext{"AddressStructure", cSTAE, err}
			}
		case cPOST:
			if PostalCodeSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"AddressStructure", cPOST, ErrSingleMultiple}
			}
			PostalCodeSet = true
			if err := s.PostalCode.parse(&sl, o); err != nil {
				return ErrContext{"AddressStructure", cPOST, err}
			}
		case cCTRY:
			if CountrySet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"AddressStructure", cCTRY, ErrSingleMultiple}
			}
			CountrySet = true
			if err := s.Country.parse(&sl, o); err != nil {
				return ErrContext{"AddressStructure", cCTRY, err}
			}
		default:
			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
				return ErrContext{"AddressStructure", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// AssociationStructure is a GEDCOM structure type
type AssociationStructure struct {
	ID         Xref
	RecordType RecordType
	Relation   RelationIsDescriptor
	Notes      []NoteStructure
	Sources    []SourceCitation
}

func (s *AssociationStructure) parse(l *Line, o options) error {
	if err := s.ID.parse(&Line{line: line{value: l.xrefID}}, o); err != nil {
		return ErrContext{"AssociationStructure", "xrefID", err}
	}
	var RecordTypeSet, RelationSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case cTYPE:
			if RecordTypeSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"AssociationStructure", cTYPE, ErrSingleMultiple}
			}
			RecordTypeSet = true
			if err := s.RecordType.parse(&sl, o); err != nil {
				return ErrContext{"AssociationStructure", cTYPE, err}
			}
		case cRELA:
			if RelationSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"AssociationStructure", cRELA, ErrSingleMultiple}
			}
			RelationSet = true
			if err := s.Relation.parse(&sl, o); err != nil {
				return ErrContext{"AssociationStructure", cRELA, err}
			}
		case cNOTE:
			var t NoteStructure
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"AssociationStructure", cNOTE, err}
			}
			s.Notes = append(s.Notes, t)
		case cSOUR:
			var t SourceCitation
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"AssociationStructure", cSOUR, err}
			}
			s.Sources = append(s.Sources, t)
		default:
			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
				return ErrContext{"AssociationStructure", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	if !o.allowMissingRequired {
		if !RecordTypeSet {
			return ErrContext{"AssociationStructure", "RecordType", ErrRequiredMissing}
		}
		if !RelationSet {
			return ErrContext{"AssociationStructure", "Relation", ErrRequiredMissing}
		}
	}
	return nil
}

// ChangeDateStructure is a GEDCOM structure type
type ChangeDateStructure struct {
	Date  ChangeDateTime
	Notes []NoteStructure
}

func (s *ChangeDateStructure) parse(l *Line, o options) error {
	var DateSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case cDATE:
			if DateSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"ChangeDateStructure", cDATE, ErrSingleMultiple}
			}
			DateSet = true
			if err := s.Date.parse(&sl, o); err != nil {
				return ErrContext{"ChangeDateStructure", cDATE, err}
			}
		case cNOTE:
			var t NoteStructure
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"ChangeDateStructure", cNOTE, err}
			}
			s.Notes = append(s.Notes, t)
		default:
			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
				return ErrContext{"ChangeDateStructure", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	if !o.allowMissingRequired {
		if !DateSet {
			return ErrContext{"ChangeDateStructure", "Date", ErrRequiredMissing}
		}
	}
	return nil
}

// ChangeDateTime is a GEDCOM structure type
type ChangeDateTime struct {
	ChangeDate ChangeDate
	Time       TimeValue
}

func (s *ChangeDateTime) parse(l *Line, o options) error {
	if err := s.ChangeDate.parse(l, o); err != nil {
		return ErrContext{"ChangeDateTime", "line_value", err}
	}
	var TimeSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case cTIME:
			if TimeSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"ChangeDateTime", cTIME, ErrSingleMultiple}
			}
			TimeSet = true
			if err := s.Time.parse(&sl, o); err != nil {
				return ErrContext{"ChangeDateTime", cTIME, err}
			}
		default:
			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
				return ErrContext{"ChangeDateTime", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// ChildToFamilyLink is a GEDCOM structure type
type ChildToFamilyLink struct {
	ID                  Xref
	PedigreeLinkageType []PedigreeLinkageType
	Notes               []NoteStructure
}

func (s *ChildToFamilyLink) parse(l *Line, o options) error {
	if err := s.ID.parse(l, o); err != nil {
		return ErrContext{"ChildToFamilyLink", "line_value", err}
	}
	for _, sl := range l.Sub {
		switch sl.tag {
		case cPEDI:
			var t PedigreeLinkageType
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"ChildToFamilyLink", cPEDI, err}
			}
			s.PedigreeLinkageType = append(s.PedigreeLinkageType, t)
		case cNOTE:
			var t NoteStructure
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"ChildToFamilyLink", cNOTE, err}
			}
			s.Notes = append(s.Notes, t)
		default:
			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
				return ErrContext{"ChildToFamilyLink", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// EventDetail is a GEDCOM structure type
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

func (s *EventDetail) parse(l *Line, o options) error {
	var TypeSet, DateSet, PlaceSet, AddressSet, AgeSet, ResponsibleAgencySet, CauseOfEventSet bool
	s.PhoneNumber = make([]PhoneNumber, 0, 3)
	for _, sl := range l.Sub {
		switch sl.tag {
		case cTYPE:
			if TypeSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"EventDetail", cTYPE, ErrSingleMultiple}
			}
			TypeSet = true
			if err := s.Type.parse(&sl, o); err != nil {
				return ErrContext{"EventDetail", cTYPE, err}
			}
		case cDATE:
			if DateSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"EventDetail", cDATE, ErrSingleMultiple}
			}
			DateSet = true
			if err := s.Date.parse(&sl, o); err != nil {
				return ErrContext{"EventDetail", cDATE, err}
			}
		case cPLAC:
			if PlaceSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"EventDetail", cPLAC, ErrSingleMultiple}
			}
			PlaceSet = true
			if err := s.Place.parse(&sl, o); err != nil {
				return ErrContext{"EventDetail", cPLAC, err}
			}
		case cADDR:
			if AddressSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"EventDetail", cADDR, ErrSingleMultiple}
			}
			AddressSet = true
			if err := s.Address.parse(&sl, o); err != nil {
				return ErrContext{"EventDetail", cADDR, err}
			}
		case cPHON:
			if len(s.PhoneNumber) == 3 {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"EventDetail", cPHON, ErrTooMany(3)}
			}
			var t PhoneNumber
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"EventDetail", cPHON, err}
			}
			s.PhoneNumber = append(s.PhoneNumber, t)
		case cAGE:
			if AgeSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"EventDetail", cAGE, ErrSingleMultiple}
			}
			AgeSet = true
			if err := s.Age.parse(&sl, o); err != nil {
				return ErrContext{"EventDetail", cAGE, err}
			}
		case cAGNC:
			if ResponsibleAgencySet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"EventDetail", cAGNC, ErrSingleMultiple}
			}
			ResponsibleAgencySet = true
			if err := s.ResponsibleAgency.parse(&sl, o); err != nil {
				return ErrContext{"EventDetail", cAGNC, err}
			}
		case cCAUS:
			if CauseOfEventSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"EventDetail", cCAUS, ErrSingleMultiple}
			}
			CauseOfEventSet = true
			if err := s.CauseOfEvent.parse(&sl, o); err != nil {
				return ErrContext{"EventDetail", cCAUS, err}
			}
		case cSOUR:
			var t SourceCitation
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"EventDetail", cSOUR, err}
			}
			s.Sources = append(s.Sources, t)
		case cOBJE:
			var t MultimediaLink
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"EventDetail", cOBJE, err}
			}
			s.Multimedia = append(s.Multimedia, t)
		case cNOTE:
			var t NoteStructure
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"EventDetail", cNOTE, err}
			}
			s.Notes = append(s.Notes, t)
		default:
			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
				return ErrContext{"EventDetail", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// LDSSpouseSealing is a GEDCOM structure type
type LDSSpouseSealing struct {
	LDSSpouseSealingDateStatus LDSSpouseSealingDateStatus
	Date                       DateLDSOrd
	TempleCode                 TempleCode
	Place                      PlaceLivingOrdinance
	Sources                    []SourceCitation
	Notes                      []NoteStructure
}

func (s *LDSSpouseSealing) parse(l *Line, o options) error {
	var LDSSpouseSealingDateStatusSet, DateSet, TempleCodeSet, PlaceSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case cSTAT:
			if LDSSpouseSealingDateStatusSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"LDSSpouseSealing", cSTAT, ErrSingleMultiple}
			}
			LDSSpouseSealingDateStatusSet = true
			if err := s.LDSSpouseSealingDateStatus.parse(&sl, o); err != nil {
				return ErrContext{"LDSSpouseSealing", cSTAT, err}
			}
		case cDATE:
			if DateSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"LDSSpouseSealing", cDATE, ErrSingleMultiple}
			}
			DateSet = true
			if err := s.Date.parse(&sl, o); err != nil {
				return ErrContext{"LDSSpouseSealing", cDATE, err}
			}
		case cTEMP:
			if TempleCodeSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"LDSSpouseSealing", cTEMP, ErrSingleMultiple}
			}
			TempleCodeSet = true
			if err := s.TempleCode.parse(&sl, o); err != nil {
				return ErrContext{"LDSSpouseSealing", cTEMP, err}
			}
		case cPLAC:
			if PlaceSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"LDSSpouseSealing", cPLAC, ErrSingleMultiple}
			}
			PlaceSet = true
			if err := s.Place.parse(&sl, o); err != nil {
				return ErrContext{"LDSSpouseSealing", cPLAC, err}
			}
		case cSOUR:
			var t SourceCitation
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"LDSSpouseSealing", cSOUR, err}
			}
			s.Sources = append(s.Sources, t)
		case cNOTE:
			var t NoteStructure
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"LDSSpouseSealing", cNOTE, err}
			}
			s.Notes = append(s.Notes, t)
		default:
			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
				return ErrContext{"LDSSpouseSealing", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// MultimediaLinkID is a GEDCOM structure type
type MultimediaLinkID struct {
	ID Xref
}

func (s *MultimediaLinkID) parse(l *Line, o options) error {
	if err := s.ID.parse(&Line{line: line{value: l.xrefID}}, o); err != nil {
		return ErrContext{"MultimediaLinkID", "xrefID", err}
	}
	for _, sl := range l.Sub {
		if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
			return ErrContext{"MultimediaLinkID", sl.tag, ErrUnknownTag}
		}
		// possibly store in a Other field
	}
	return nil
}

// MultimediaLinkFile is a GEDCOM structure type
type MultimediaLinkFile struct {
	Format MultimediaFormat
	Title  DescriptiveTitle
	File   MultimediaFileReference
	Notes  []NoteStructure
}

func (s *MultimediaLinkFile) parse(l *Line, o options) error {
	var FormatSet, TitleSet, FileSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case cFORM:
			if FormatSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"MultimediaLinkFile", cFORM, ErrSingleMultiple}
			}
			FormatSet = true
			if err := s.Format.parse(&sl, o); err != nil {
				return ErrContext{"MultimediaLinkFile", cFORM, err}
			}
		case cTITLE:
			if TitleSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"MultimediaLinkFile", cTITLE, ErrSingleMultiple}
			}
			TitleSet = true
			if err := s.Title.parse(&sl, o); err != nil {
				return ErrContext{"MultimediaLinkFile", cTITLE, err}
			}
		case cFILE:
			if FileSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"MultimediaLinkFile", cFILE, ErrSingleMultiple}
			}
			FileSet = true
			if err := s.File.parse(&sl, o); err != nil {
				return ErrContext{"MultimediaLinkFile", cFILE, err}
			}
		case cNOTE:
			var t NoteStructure
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"MultimediaLinkFile", cNOTE, err}
			}
			s.Notes = append(s.Notes, t)
		default:
			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
				return ErrContext{"MultimediaLinkFile", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	if !o.allowMissingRequired {
		if !FormatSet {
			return ErrContext{"MultimediaLinkFile", "Format", ErrRequiredMissing}
		}
	}
	return nil
}

// NoteID is a GEDCOM structure type
type NoteID struct {
	ID Xref
}

func (s *NoteID) parse(l *Line, o options) error {
	if err := s.ID.parse(&Line{line: line{value: l.xrefID}}, o); err != nil {
		return ErrContext{"NoteID", "xrefID", err}
	}
	for _, sl := range l.Sub {
		if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
			return ErrContext{"NoteID", sl.tag, ErrUnknownTag}
		}
		// possibly store in a Other field
	}
	return nil
}

// NoteText is a GEDCOM structure type
type NoteText struct {
	SubmitterText SubmitterText
	Sources       []SourceCitation
}

func (s *NoteText) parse(l *Line, o options) error {
	if err := s.SubmitterText.parse(l, o); err != nil {
		return ErrContext{"NoteText", "line_value", err}
	}
	for _, sl := range l.Sub {
		switch sl.tag {
		case cSOUR:
			var t SourceCitation
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"NoteText", cSOUR, err}
			}
			s.Sources = append(s.Sources, t)
		default:
			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
				return ErrContext{"NoteText", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// PersonalNameStructure is a GEDCOM structure type
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

func (s *PersonalNameStructure) parse(l *Line, o options) error {
	if err := s.NamePersonal.parse(l, o); err != nil {
		return ErrContext{"PersonalNameStructure", "line_value", err}
	}
	var PrefixSet, GivenSet, NicknameSet, SurnamePrefixSet, SurnameSet, SuffixSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case cNPFX:
			if PrefixSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"PersonalNameStructure", cNPFX, ErrSingleMultiple}
			}
			PrefixSet = true
			if err := s.Prefix.parse(&sl, o); err != nil {
				return ErrContext{"PersonalNameStructure", cNPFX, err}
			}
		case cGIVN:
			if GivenSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"PersonalNameStructure", cGIVN, ErrSingleMultiple}
			}
			GivenSet = true
			if err := s.Given.parse(&sl, o); err != nil {
				return ErrContext{"PersonalNameStructure", cGIVN, err}
			}
		case cNICK:
			if NicknameSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"PersonalNameStructure", cNICK, ErrSingleMultiple}
			}
			NicknameSet = true
			if err := s.Nickname.parse(&sl, o); err != nil {
				return ErrContext{"PersonalNameStructure", cNICK, err}
			}
		case cSPFX:
			if SurnamePrefixSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"PersonalNameStructure", cSPFX, ErrSingleMultiple}
			}
			SurnamePrefixSet = true
			if err := s.SurnamePrefix.parse(&sl, o); err != nil {
				return ErrContext{"PersonalNameStructure", cSPFX, err}
			}
		case cSURN:
			if SurnameSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"PersonalNameStructure", cSURN, ErrSingleMultiple}
			}
			SurnameSet = true
			if err := s.Surname.parse(&sl, o); err != nil {
				return ErrContext{"PersonalNameStructure", cSURN, err}
			}
		case cNSFX:
			if SuffixSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"PersonalNameStructure", cNSFX, ErrSingleMultiple}
			}
			SuffixSet = true
			if err := s.Suffix.parse(&sl, o); err != nil {
				return ErrContext{"PersonalNameStructure", cNSFX, err}
			}
		case cSOUR:
			var t SourceCitation
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"PersonalNameStructure", cSOUR, err}
			}
			s.Sources = append(s.Sources, t)
		case cNOTE:
			var t NoteStructure
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"PersonalNameStructure", cNOTE, err}
			}
			s.Notes = append(s.Notes, t)
		default:
			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
				return ErrContext{"PersonalNameStructure", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// PlaceStructure is a GEDCOM structure type
type PlaceStructure struct {
	Place          PlaceValue
	PlaceHierarchy PlaceHierarchy
	Sources        []SourceCitation
	Notes          []NoteStructure
}

func (s *PlaceStructure) parse(l *Line, o options) error {
	if err := s.Place.parse(l, o); err != nil {
		return ErrContext{"PlaceStructure", "line_value", err}
	}
	var PlaceHierarchySet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case cFORM:
			if PlaceHierarchySet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"PlaceStructure", cFORM, ErrSingleMultiple}
			}
			PlaceHierarchySet = true
			if err := s.PlaceHierarchy.parse(&sl, o); err != nil {
				return ErrContext{"PlaceStructure", cFORM, err}
			}
		case cSOUR:
			var t SourceCitation
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"PlaceStructure", cSOUR, err}
			}
			s.Sources = append(s.Sources, t)
		case cNOTE:
			var t NoteStructure
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"PlaceStructure", cNOTE, err}
			}
			s.Notes = append(s.Notes, t)
		default:
			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
				return ErrContext{"PlaceStructure", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// SourceID is a GEDCOM structure type
type SourceID struct {
	ID                  Xref
	WhereWithinSource   WhereWithinSource
	SourceCitationEvent SourceCitationEvent
	Data                SourceData
	Certainty           CertaintyAssessment
	Multimedia          []MultimediaLink
	Notes               []NoteStructure
}

func (s *SourceID) parse(l *Line, o options) error {
	if err := s.ID.parse(&Line{line: line{value: l.xrefID}}, o); err != nil {
		return ErrContext{"SourceID", "xrefID", err}
	}
	var WhereWithinSourceSet, SourceCitationEventSet, DataSet, CertaintySet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case cPAGE:
			if WhereWithinSourceSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"SourceID", cPAGE, ErrSingleMultiple}
			}
			WhereWithinSourceSet = true
			if err := s.WhereWithinSource.parse(&sl, o); err != nil {
				return ErrContext{"SourceID", cPAGE, err}
			}
		case cEVEN:
			if SourceCitationEventSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"SourceID", cEVEN, ErrSingleMultiple}
			}
			SourceCitationEventSet = true
			if err := s.SourceCitationEvent.parse(&sl, o); err != nil {
				return ErrContext{"SourceID", cEVEN, err}
			}
		case cDATA:
			if DataSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"SourceID", cDATA, ErrSingleMultiple}
			}
			DataSet = true
			if err := s.Data.parse(&sl, o); err != nil {
				return ErrContext{"SourceID", cDATA, err}
			}
		case cQUAY:
			if CertaintySet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"SourceID", cQUAY, ErrSingleMultiple}
			}
			CertaintySet = true
			if err := s.Certainty.parse(&sl, o); err != nil {
				return ErrContext{"SourceID", cQUAY, err}
			}
		case cOBJE:
			var t MultimediaLink
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"SourceID", cOBJE, err}
			}
			s.Multimedia = append(s.Multimedia, t)
		case cNOTE:
			var t NoteStructure
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"SourceID", cNOTE, err}
			}
			s.Notes = append(s.Notes, t)
		default:
			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
				return ErrContext{"SourceID", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// SourceCitationEvent is a GEDCOM structure type
type SourceCitationEvent struct {
	EventTypeCitedFrom EventTypeCitedFrom
	Role               RoleInEvent
}

func (s *SourceCitationEvent) parse(l *Line, o options) error {
	if err := s.EventTypeCitedFrom.parse(l, o); err != nil {
		return ErrContext{"SourceCitationEvent", "line_value", err}
	}
	var RoleSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case cROLE:
			if RoleSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"SourceCitationEvent", cROLE, ErrSingleMultiple}
			}
			RoleSet = true
			if err := s.Role.parse(&sl, o); err != nil {
				return ErrContext{"SourceCitationEvent", cROLE, err}
			}
		default:
			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
				return ErrContext{"SourceCitationEvent", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// SourceData is a GEDCOM structure type
type SourceData struct {
	Date EntryRecordingDate
	Text []TextFromSource
}

func (s *SourceData) parse(l *Line, o options) error {
	var DateSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case cDATE:
			if DateSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"SourceData", cDATE, ErrSingleMultiple}
			}
			DateSet = true
			if err := s.Date.parse(&sl, o); err != nil {
				return ErrContext{"SourceData", cDATE, err}
			}
		case cTEXT:
			var t TextFromSource
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"SourceData", cTEXT, err}
			}
			s.Text = append(s.Text, t)
		default:
			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
				return ErrContext{"SourceData", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// SourceText is a GEDCOM structure type
type SourceText struct {
	Description SourceDescription
	Texts       []TextFromSource
	Notes       []NoteStructure
}

func (s *SourceText) parse(l *Line, o options) error {
	if err := s.Description.parse(l, o); err != nil {
		return ErrContext{"SourceText", "line_value", err}
	}
	for _, sl := range l.Sub {
		switch sl.tag {
		case cTEXT:
			var t TextFromSource
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"SourceText", cTEXT, err}
			}
			s.Texts = append(s.Texts, t)
		case cNOTE:
			var t NoteStructure
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"SourceText", cNOTE, err}
			}
			s.Notes = append(s.Notes, t)
		default:
			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
				return ErrContext{"SourceText", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// SourceRepositoryCitation is a GEDCOM structure type
type SourceRepositoryCitation struct {
	ID      Xref
	Numbers []SourceCallStructure
}

func (s *SourceRepositoryCitation) parse(l *Line, o options) error {
	if err := s.ID.parse(&Line{line: line{value: l.xrefID}}, o); err != nil {
		return ErrContext{"SourceRepositoryCitation", "xrefID", err}
	}
	for _, sl := range l.Sub {
		switch sl.tag {
		case cCALN:
			var t SourceCallStructure
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"SourceRepositoryCitation", cCALN, err}
			}
			s.Numbers = append(s.Numbers, t)
		default:
			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
				return ErrContext{"SourceRepositoryCitation", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// SourceCallStructure is a GEDCOM structure type
type SourceCallStructure struct {
	SourceCallNumber SourceCallNumber
	Type             SourceMediaType
}

func (s *SourceCallStructure) parse(l *Line, o options) error {
	if err := s.SourceCallNumber.parse(l, o); err != nil {
		return ErrContext{"SourceCallStructure", "line_value", err}
	}
	var TypeSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case cMEDI:
			if TypeSet {
				if !o.allowMoreThanAllowed {
					continue
				}
				return ErrContext{"SourceCallStructure", cMEDI, ErrSingleMultiple}
			}
			TypeSet = true
			if err := s.Type.parse(&sl, o); err != nil {
				return ErrContext{"SourceCallStructure", cMEDI, err}
			}
		default:
			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
				return ErrContext{"SourceCallStructure", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// SpouseToFamilyLink is a GEDCOM structure type
type SpouseToFamilyLink struct {
	ID    Xref
	Notes []NoteStructure
}

func (s *SpouseToFamilyLink) parse(l *Line, o options) error {
	if err := s.ID.parse(l, o); err != nil {
		return ErrContext{"SpouseToFamilyLink", "line_value", err}
	}
	for _, sl := range l.Sub {
		switch sl.tag {
		case cNOTE:
			var t NoteStructure
			if err := t.parse(&sl, o); err != nil {
				return ErrContext{"SpouseToFamilyLink", cNOTE, err}
			}
			s.Notes = append(s.Notes, t)
		default:
			if !o.allowMissingRequired && (len(sl.tag) < 1 || sl.tag[0] != '_') {
				return ErrContext{"SpouseToFamilyLink", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// MultimediaLink splite between MultimediaLinkID and MultimediaLinkFile
type MultimediaLink struct {
	Data Record
}

func (s *MultimediaLink) parse(l *Line, o options) error {
	var err error
	if l.xrefID != "" {
		t := &MultimediaLinkID{}
		err = t.parse(l, o)
		s.Data = t
	} else {
		t := &MultimediaLinkFile{}
		err = t.parse(l, o)
		s.Data = t
	}
	return err
}

// NoteStructure splits between NoteID and NoteText
type NoteStructure struct {
	Data Record
}

func (s *NoteStructure) parse(l *Line, o options) error {
	var err error
	if l.xrefID != "" {
		t := &NoteID{}
		err = t.parse(l, o)
		s.Data = t
	} else {
		t := &NoteText{}
		err = t.parse(l, o)
		s.Data = t
	}
	return err
}

// SourceCitation splits between SourceID and SourceText
type SourceCitation struct {
	Data Record
}

func (s *SourceCitation) parse(l *Line, o options) error {
	var err error
	if l.xrefID != "" {
		t := &SourceID{}
		err = t.parse(l, o)
		s.Data = t
	} else {
		t := &SourceText{}
		err = t.parse(l, o)
		s.Data = t
	}
	return err
}

// Trailer type
type Trailer struct{}

func (s *Trailer) parse(*Line, options) error {
	return nil
}

// Errors
var (
	ErrRequiredMissing = errors.New("required tag missing")
	ErrSingleMultiple  = errors.New("tag was specified more than the one time allowed")
	ErrUnknownTag      = errors.New("unknown tag")
)

// ErrContext adds context to a returned error
type ErrContext struct {
	Structure, Tag string
	Err            error
}

// Error implements the error interface
func (e ErrContext) Error() string {
	return e.Tag + ":" + e.Err.Error()
}

// Unwrap goes through the error list to retrieve the underlying
// (non-ErrContext) error
func (e ErrContext) Unwrap() error {
	err := e.Err
	for {
		if er, ok := err.(ErrContext); ok {
			err = er.Err
			continue
		}
		return err
	}
}

// ErrTooMany is an error returned when too many of a particular tag exist
type ErrTooMany int

// Error implements the error interface
func (ErrTooMany) Error() string {
	return "too many tags"
}
