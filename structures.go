package gedcom

// File automatically generated with ./genStructures.sh

import "errors"

// Header is a GEDCOM structure type
type Header struct {
	Source                   HeaderSource
	ReceivingSystemName      ReceivingSystemName
	TransmissionLDate        TransmissionDateTime
	Submitter                Xref
	Submission               Xref
	FileName                 FileName
	Copyright                CopyrightGedcomFile
	Version                  Version
	CharacterSet             CharacterSetStructure
	Language                 LanguageOfText
	Place                    HeaderPlace
	GedcomContentDescription GedcomContentDescription
}

func (s *Header) parse(l Line) error {
	var SourceSet, SubmitterSet, VersionSet, CharacterSetSet, ReceivingSystemNameSet, TransmissionLDateSet, SubmissionSet, FileNameSet, CopyrightSet, LanguageSet, PlaceSet, GedcomContentDescriptionSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "SOUR":
			if SourceSet {
				return ErrContext{"Header", "SOUR", ErrSingleMultiple}
			}
			SourceSet = true
			if err := s.Source.parse(sl); err != nil {
				return ErrContext{"Header", "SOUR", err}
			}
		case "DEST":
			if ReceivingSystemNameSet {
				return ErrContext{"Header", "DEST", ErrSingleMultiple}
			}
			ReceivingSystemNameSet = true
			if err := s.ReceivingSystemName.parse(sl); err != nil {
				return ErrContext{"Header", "DEST", err}
			}
		case "DATE":
			if TransmissionLDateSet {
				return ErrContext{"Header", "DATE", ErrSingleMultiple}
			}
			TransmissionLDateSet = true
			if err := s.TransmissionLDate.parse(sl); err != nil {
				return ErrContext{"Header", "DATE", err}
			}
		case "SUBM":
			if SubmitterSet {
				return ErrContext{"Header", "SUBM", ErrSingleMultiple}
			}
			SubmitterSet = true
			if err := s.Submitter.parse(sl); err != nil {
				return ErrContext{"Header", "SUBM", err}
			}
		case "SUMB":
			if SubmissionSet {
				return ErrContext{"Header", "SUMB", ErrSingleMultiple}
			}
			SubmissionSet = true
			if err := s.Submission.parse(sl); err != nil {
				return ErrContext{"Header", "SUMB", err}
			}
		case "FILE":
			if FileNameSet {
				return ErrContext{"Header", "FILE", ErrSingleMultiple}
			}
			FileNameSet = true
			if err := s.FileName.parse(sl); err != nil {
				return ErrContext{"Header", "FILE", err}
			}
		case "COPR":
			if CopyrightSet {
				return ErrContext{"Header", "COPR", ErrSingleMultiple}
			}
			CopyrightSet = true
			if err := s.Copyright.parse(sl); err != nil {
				return ErrContext{"Header", "COPR", err}
			}
		case "GEDC":
			if VersionSet {
				return ErrContext{"Header", "GEDC", ErrSingleMultiple}
			}
			VersionSet = true
			if err := s.Version.parse(sl); err != nil {
				return ErrContext{"Header", "GEDC", err}
			}
		case "CHAR":
			if CharacterSetSet {
				return ErrContext{"Header", "CHAR", ErrSingleMultiple}
			}
			CharacterSetSet = true
			if err := s.CharacterSet.parse(sl); err != nil {
				return ErrContext{"Header", "CHAR", err}
			}
		case "LANG":
			if LanguageSet {
				return ErrContext{"Header", "LANG", ErrSingleMultiple}
			}
			LanguageSet = true
			if err := s.Language.parse(sl); err != nil {
				return ErrContext{"Header", "LANG", err}
			}
		case "PLAC":
			if PlaceSet {
				return ErrContext{"Header", "PLAC", ErrSingleMultiple}
			}
			PlaceSet = true
			if err := s.Place.parse(sl); err != nil {
				return ErrContext{"Header", "PLAC", err}
			}
		case "NOTE":
			if GedcomContentDescriptionSet {
				return ErrContext{"Header", "NOTE", ErrSingleMultiple}
			}
			GedcomContentDescriptionSet = true
			if err := s.GedcomContentDescription.parse(sl); err != nil {
				return ErrContext{"Header", "NOTE", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
				return ErrContext{"Header", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
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

func (s *HeaderSource) parse(l Line) error {
	if err := s.SystemID.parse(l); err != nil {
		return ErrContext{"HeaderSource", "line_value", err}
	}
	var VersionNumberSet, NameSet, BusinessSet, DataSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "VERS":
			if VersionNumberSet {
				return ErrContext{"HeaderSource", "VERS", ErrSingleMultiple}
			}
			VersionNumberSet = true
			if err := s.VersionNumber.parse(sl); err != nil {
				return ErrContext{"HeaderSource", "VERS", err}
			}
		case "NAME":
			if NameSet {
				return ErrContext{"HeaderSource", "NAME", ErrSingleMultiple}
			}
			NameSet = true
			if err := s.Name.parse(sl); err != nil {
				return ErrContext{"HeaderSource", "NAME", err}
			}
		case "CORP":
			if BusinessSet {
				return ErrContext{"HeaderSource", "CORP", ErrSingleMultiple}
			}
			BusinessSet = true
			if err := s.Business.parse(sl); err != nil {
				return ErrContext{"HeaderSource", "CORP", err}
			}
		case "DATA":
			if DataSet {
				return ErrContext{"HeaderSource", "DATA", ErrSingleMultiple}
			}
			DataSet = true
			if err := s.Data.parse(sl); err != nil {
				return ErrContext{"HeaderSource", "DATA", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
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

func (s *TransmissionDateTime) parse(l Line) error {
	if err := s.TransmissionDate.parse(l); err != nil {
		return ErrContext{"TransmissionDateTime", "line_value", err}
	}
	var TimeSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "TIME":
			if TimeSet {
				return ErrContext{"TransmissionDateTime", "TIME", ErrSingleMultiple}
			}
			TimeSet = true
			if err := s.Time.parse(sl); err != nil {
				return ErrContext{"TransmissionDateTime", "TIME", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
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

func (s *HeaderBusiness) parse(l Line) error {
	if err := s.NameOfBusiness.parse(l); err != nil {
		return ErrContext{"HeaderBusiness", "line_value", err}
	}
	var AddressSet bool
	var PhoneNumberCount int
	for _, sl := range l.Sub {
		switch sl.tag {
		case "ADDR":
			if AddressSet {
				return ErrContext{"HeaderBusiness", "ADDR", ErrSingleMultiple}
			}
			AddressSet = true
			if err := s.Address.parse(sl); err != nil {
				return ErrContext{"HeaderBusiness", "ADDR", err}
			}
		case "PHON":
			if PhoneNumberCount == 3 {
				return ErrContext{"HeaderBusiness", "PHON", ErrTooMany(3)}
			}
			PhoneNumberCount++
			var t PhoneNumber
			if err != t.parse(sl); err != nil {
				return ErrContext{"HeaderBusiness", "PHON", err}
			}
			s.PhoneNumber = append(s.PhoneNumber, t)
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
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

func (s *HeaderDataSource) parse(l Line) error {
	if err := s.SourceName.parse(l); err != nil {
		return ErrContext{"HeaderDataSource", "line_value", err}
	}
	var PublicationDateSet, CopyrightSourceDataSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "DATE":
			if PublicationDateSet {
				return ErrContext{"HeaderDataSource", "DATE", ErrSingleMultiple}
			}
			PublicationDateSet = true
			if err := s.PublicationDate.parse(sl); err != nil {
				return ErrContext{"HeaderDataSource", "DATE", err}
			}
		case "COPR":
			if CopyrightSourceDataSet {
				return ErrContext{"HeaderDataSource", "COPR", ErrSingleMultiple}
			}
			CopyrightSourceDataSet = true
			if err := s.CopyrightSourceData.parse(sl); err != nil {
				return ErrContext{"HeaderDataSource", "COPR", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
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
	GedcomForm    GedcomForm
}

func (s *Version) parse(l Line) error {
	var VersionNumberSet, GedcomFormSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "VERS":
			if VersionNumberSet {
				return ErrContext{"Version", "VERS", ErrSingleMultiple}
			}
			VersionNumberSet = true
			if err := s.VersionNumber.parse(sl); err != nil {
				return ErrContext{"Version", "VERS", err}
			}
		case "FORM":
			if GedcomFormSet {
				return ErrContext{"Version", "FORM", ErrSingleMultiple}
			}
			GedcomFormSet = true
			if err := s.GedcomForm.parse(sl); err != nil {
				return ErrContext{"Version", "FORM", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
				return ErrContext{"Version", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	if !VersionNumberSet {
		return ErrContext{"Version", "VersionNumber", ErrRequiredMissing}
	}
	if !GedcomFormSet {
		return ErrContext{"Version", "GedcomForm", ErrRequiredMissing}
	}
	return nil
}

// CharacterSetStructure is a GEDCOM structure type
type CharacterSetStructure struct {
	CharacterSet  CharacterSet
	VersionNumber VersionNumber
}

func (s *CharacterSetStructure) parse(l Line) error {
	if err := s.CharacterSet.parse(l); err != nil {
		return ErrContext{"CharacterSetStructure", "line_value", err}
	}
	var VersionNumberSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "VERS":
			if VersionNumberSet {
				return ErrContext{"CharacterSetStructure", "VERS", ErrSingleMultiple}
			}
			VersionNumberSet = true
			if err := s.VersionNumber.parse(sl); err != nil {
				return ErrContext{"CharacterSetStructure", "VERS", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
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

func (s *HeaderPlace) parse(l Line) error {
	var PlaceHierarchySet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "FORM":
			if PlaceHierarchySet {
				return ErrContext{"HeaderPlace", "FORM", ErrSingleMultiple}
			}
			PlaceHierarchySet = true
			if err := s.PlaceHierarchy.parse(sl); err != nil {
				return ErrContext{"HeaderPlace", "FORM", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
				return ErrContext{"HeaderPlace", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	if !PlaceHierarchySet {
		return ErrContext{"HeaderPlace", "PlaceHierarchy", ErrRequiredMissing}
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

func (s *Family) parse(l Line) error {
	if err := s.ID.parse(Line{value: l.xrefID}); err != nil {
		return ErrContext{"Family", "xrefID", err}
	}
	var AnnulmentSet, CensusSet, DivorceSet, DivorceFiledSet, EngagementSet, MarriageSet, MarriageBannSet, MarriageContractSet, MarriageLicenseSet, MarriageSettlementSet, HusbandSet, WifeSet, NumChildrenSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "ANUL":
			if AnnulmentSet {
				return ErrContext{"Family", "ANUL", ErrSingleMultiple}
			}
			AnnulmentSet = true
			if err := s.Annulment.parse(sl); err != nil {
				return ErrContext{"Family", "ANUL", err}
			}
		case "CENS":
			if CensusSet {
				return ErrContext{"Family", "CENS", ErrSingleMultiple}
			}
			CensusSet = true
			if err := s.Census.parse(sl); err != nil {
				return ErrContext{"Family", "CENS", err}
			}
		case "DIV":
			if DivorceSet {
				return ErrContext{"Family", "DIV", ErrSingleMultiple}
			}
			DivorceSet = true
			if err := s.Divorce.parse(sl); err != nil {
				return ErrContext{"Family", "DIV", err}
			}
		case "DIVF":
			if DivorceFiledSet {
				return ErrContext{"Family", "DIVF", ErrSingleMultiple}
			}
			DivorceFiledSet = true
			if err := s.DivorceFiled.parse(sl); err != nil {
				return ErrContext{"Family", "DIVF", err}
			}
		case "ENGA":
			if EngagementSet {
				return ErrContext{"Family", "ENGA", ErrSingleMultiple}
			}
			EngagementSet = true
			if err := s.Engagement.parse(sl); err != nil {
				return ErrContext{"Family", "ENGA", err}
			}
		case "MARR":
			if MarriageSet {
				return ErrContext{"Family", "MARR", ErrSingleMultiple}
			}
			MarriageSet = true
			if err := s.Marriage.parse(sl); err != nil {
				return ErrContext{"Family", "MARR", err}
			}
		case "MARB":
			if MarriageBannSet {
				return ErrContext{"Family", "MARB", ErrSingleMultiple}
			}
			MarriageBannSet = true
			if err := s.MarriageBann.parse(sl); err != nil {
				return ErrContext{"Family", "MARB", err}
			}
		case "MARC":
			if MarriageContractSet {
				return ErrContext{"Family", "MARC", ErrSingleMultiple}
			}
			MarriageContractSet = true
			if err := s.MarriageContract.parse(sl); err != nil {
				return ErrContext{"Family", "MARC", err}
			}
		case "MARL":
			if MarriageLicenseSet {
				return ErrContext{"Family", "MARL", ErrSingleMultiple}
			}
			MarriageLicenseSet = true
			if err := s.MarriageLicense.parse(sl); err != nil {
				return ErrContext{"Family", "MARL", err}
			}
		case "MARS":
			if MarriageSettlementSet {
				return ErrContext{"Family", "MARS", ErrSingleMultiple}
			}
			MarriageSettlementSet = true
			if err := s.MarriageSettlement.parse(sl); err != nil {
				return ErrContext{"Family", "MARS", err}
			}
		case "EVEN":
			var t FamilyEventDetail
			if err != t.parse(sl); err != nil {
				return ErrContext{"Family", "EVEN", err}
			}
			s.Events = append(s.Events, t)
		case "HUSB":
			if HusbandSet {
				return ErrContext{"Family", "HUSB", ErrSingleMultiple}
			}
			HusbandSet = true
			if err := s.Husband.parse(sl); err != nil {
				return ErrContext{"Family", "HUSB", err}
			}
		case "WIFE":
			if WifeSet {
				return ErrContext{"Family", "WIFE", ErrSingleMultiple}
			}
			WifeSet = true
			if err := s.Wife.parse(sl); err != nil {
				return ErrContext{"Family", "WIFE", err}
			}
		case "CHIL":
			var t Xref
			if err != t.parse(sl); err != nil {
				return ErrContext{"Family", "CHIL", err}
			}
			s.Children = append(s.Children, t)
		case "NCHI":
			if NumChildrenSet {
				return ErrContext{"Family", "NCHI", ErrSingleMultiple}
			}
			NumChildrenSet = true
			if err := s.NumChildren.parse(sl); err != nil {
				return ErrContext{"Family", "NCHI", err}
			}
		case "SUBM":
			var t Xref
			if err != t.parse(sl); err != nil {
				return ErrContext{"Family", "SUBM", err}
			}
			s.Submitters = append(s.Submitters, t)
		case "SLGS":
			var t LDSSpouseSealing
			if err != t.parse(sl); err != nil {
				return ErrContext{"Family", "SLGS", err}
			}
			s.LDSSpouseSealing = append(s.LDSSpouseSealing, t)
		case "SOUR":
			var t SourceCitation
			if err != t.parse(sl); err != nil {
				return ErrContext{"Family", "SOUR", err}
			}
			s.Sources = append(s.Sources, t)
		case "OBJE":
			var t MultimediaLink
			if err != t.parse(sl); err != nil {
				return ErrContext{"Family", "OBJE", err}
			}
			s.Multimedia = append(s.Multimedia, t)
		case "NOTE":
			var t NoteStructure
			if err != t.parse(sl); err != nil {
				return ErrContext{"Family", "NOTE", err}
			}
			s.Notes = append(s.Notes, t)
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
				return ErrContext{"Family", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// VerifiedFamilyEventDetail is a GEDCOM structure type
type VerifiedFamilyEventDetail struct {
	Verified          Verified
	FamilyEventDetail FamilyEventDetail
}

func (s *VerifiedFamilyEventDetail) parse(l Line) error {
	if err := s.Verified.parse(l); err != nil {
		return ErrContext{"VerifiedFamilyEventDetail", "line_value", err}
	}
	var FamilyEventDetailSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "":
			if FamilyEventDetailSet {
				return ErrContext{"VerifiedFamilyEventDetail", "", ErrSingleMultiple}
			}
			FamilyEventDetailSet = true
			if err := s.FamilyEventDetail.parse(sl); err != nil {
				return ErrContext{"VerifiedFamilyEventDetail", "", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
				return ErrContext{"VerifiedFamilyEventDetail", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	if !FamilyEventDetailSet {
		return ErrContext{"VerifiedFamilyEventDetail", "FamilyEventDetail", ErrRequiredMissing}
	}
	return nil
}

// FamilyEventDetail is a GEDCOM structure type
type FamilyEventDetail struct {
	HusbandAge AgeStructure
	WifeAge    AgeStructure
	Event      EventDetail
}

func (s *FamilyEventDetail) parse(l Line) error {
	var HusbandAgeSet, WifeAgeSet, EventSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "HUSB":
			if HusbandAgeSet {
				return ErrContext{"FamilyEventDetail", "HUSB", ErrSingleMultiple}
			}
			HusbandAgeSet = true
			if err := s.HusbandAge.parse(sl); err != nil {
				return ErrContext{"FamilyEventDetail", "HUSB", err}
			}
		case "WIFE":
			if WifeAgeSet {
				return ErrContext{"FamilyEventDetail", "WIFE", ErrSingleMultiple}
			}
			WifeAgeSet = true
			if err := s.WifeAge.parse(sl); err != nil {
				return ErrContext{"FamilyEventDetail", "WIFE", err}
			}
		case "EVEN":
			if EventSet {
				return ErrContext{"FamilyEventDetail", "EVEN", ErrSingleMultiple}
			}
			EventSet = true
			if err := s.Event.parse(sl); err != nil {
				return ErrContext{"FamilyEventDetail", "EVEN", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
				return ErrContext{"FamilyEventDetail", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// AgeStructure is a GEDCOM structure type
type AgeStructure struct {
	Age AgeAtEvent
}

func (s *AgeStructure) parse(l Line) error {
	var AgeSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "AGE":
			if AgeSet {
				return ErrContext{"AgeStructure", "AGE", ErrSingleMultiple}
			}
			AgeSet = true
			if err := s.Age.parse(sl); err != nil {
				return ErrContext{"AgeStructure", "AGE", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
				return ErrContext{"AgeStructure", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	if !AgeSet {
		return ErrContext{"AgeStructure", "Age", ErrRequiredMissing}
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
	Emmigrated            VerifiedEventDetail
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

func (s *Individual) parse(l Line) error {
	if err := s.ID.parse(Line{value: l.xrefID}); err != nil {
		return ErrContext{"Individual", "xrefID", err}
	}
	var RestrictionNoticeSet, GenderSet, BirthSet, ChristeningSet, DeathSet, BuriedSet, CremationSet, AdoptionSet, MaptismSet, BarMitzvahSet, BasMitzvahSet, BlessingSet, AdultChristeningSet, ConfirmationSet, FirstCommunionSet, OrdinationSet, NaturalizationSet, EmmigratedSet, ImmigratedSet, CensusSet, ProbateSet, WillSet, GraduatedSet, RetiredSet, PermanentRecordSet, AncestralFileNumberSet, AutomatedRecordIDSet, ChangeDateSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "RESN":
			if RestrictionNoticeSet {
				return ErrContext{"Individual", "RESN", ErrSingleMultiple}
			}
			RestrictionNoticeSet = true
			if err := s.RestrictionNotice.parse(sl); err != nil {
				return ErrContext{"Individual", "RESN", err}
			}
		case "NAME":
			var t PersonalNameStructure
			if err != t.parse(sl); err != nil {
				return ErrContext{"Individual", "NAME", err}
			}
			s.PersonalNameStructure = append(s.PersonalNameStructure, t)
		case "SEX":
			if GenderSet {
				return ErrContext{"Individual", "SEX", ErrSingleMultiple}
			}
			GenderSet = true
			if err := s.Gender.parse(sl); err != nil {
				return ErrContext{"Individual", "SEX", err}
			}
		case "BIRT":
			if BirthSet {
				return ErrContext{"Individual", "BIRT", ErrSingleMultiple}
			}
			BirthSet = true
			if err := s.Birth.parse(sl); err != nil {
				return ErrContext{"Individual", "BIRT", err}
			}
		case "CHR":
			if ChristeningSet {
				return ErrContext{"Individual", "CHR", ErrSingleMultiple}
			}
			ChristeningSet = true
			if err := s.Christening.parse(sl); err != nil {
				return ErrContext{"Individual", "CHR", err}
			}
		case "DEAT":
			if DeathSet {
				return ErrContext{"Individual", "DEAT", ErrSingleMultiple}
			}
			DeathSet = true
			if err := s.Death.parse(sl); err != nil {
				return ErrContext{"Individual", "DEAT", err}
			}
		case "BURI":
			if BuriedSet {
				return ErrContext{"Individual", "BURI", ErrSingleMultiple}
			}
			BuriedSet = true
			if err := s.Buried.parse(sl); err != nil {
				return ErrContext{"Individual", "BURI", err}
			}
		case "CREM":
			if CremationSet {
				return ErrContext{"Individual", "CREM", ErrSingleMultiple}
			}
			CremationSet = true
			if err := s.Cremation.parse(sl); err != nil {
				return ErrContext{"Individual", "CREM", err}
			}
		case "ADOP":
			if AdoptionSet {
				return ErrContext{"Individual", "ADOP", ErrSingleMultiple}
			}
			AdoptionSet = true
			if err := s.Adoption.parse(sl); err != nil {
				return ErrContext{"Individual", "ADOP", err}
			}
		case "BAPM":
			if MaptismSet {
				return ErrContext{"Individual", "BAPM", ErrSingleMultiple}
			}
			MaptismSet = true
			if err := s.Maptism.parse(sl); err != nil {
				return ErrContext{"Individual", "BAPM", err}
			}
		case "BARM":
			if BarMitzvahSet {
				return ErrContext{"Individual", "BARM", ErrSingleMultiple}
			}
			BarMitzvahSet = true
			if err := s.BarMitzvah.parse(sl); err != nil {
				return ErrContext{"Individual", "BARM", err}
			}
		case "BASM":
			if BasMitzvahSet {
				return ErrContext{"Individual", "BASM", ErrSingleMultiple}
			}
			BasMitzvahSet = true
			if err := s.BasMitzvah.parse(sl); err != nil {
				return ErrContext{"Individual", "BASM", err}
			}
		case "BLES":
			if BlessingSet {
				return ErrContext{"Individual", "BLES", ErrSingleMultiple}
			}
			BlessingSet = true
			if err := s.Blessing.parse(sl); err != nil {
				return ErrContext{"Individual", "BLES", err}
			}
		case "CHRA":
			if AdultChristeningSet {
				return ErrContext{"Individual", "CHRA", ErrSingleMultiple}
			}
			AdultChristeningSet = true
			if err := s.AdultChristening.parse(sl); err != nil {
				return ErrContext{"Individual", "CHRA", err}
			}
		case "CONF":
			if ConfirmationSet {
				return ErrContext{"Individual", "CONF", ErrSingleMultiple}
			}
			ConfirmationSet = true
			if err := s.Confirmation.parse(sl); err != nil {
				return ErrContext{"Individual", "CONF", err}
			}
		case "FCOM":
			if FirstCommunionSet {
				return ErrContext{"Individual", "FCOM", ErrSingleMultiple}
			}
			FirstCommunionSet = true
			if err := s.FirstCommunion.parse(sl); err != nil {
				return ErrContext{"Individual", "FCOM", err}
			}
		case "ORDN":
			if OrdinationSet {
				return ErrContext{"Individual", "ORDN", ErrSingleMultiple}
			}
			OrdinationSet = true
			if err := s.Ordination.parse(sl); err != nil {
				return ErrContext{"Individual", "ORDN", err}
			}
		case "NATU":
			if NaturalizationSet {
				return ErrContext{"Individual", "NATU", ErrSingleMultiple}
			}
			NaturalizationSet = true
			if err := s.Naturalization.parse(sl); err != nil {
				return ErrContext{"Individual", "NATU", err}
			}
		case "EMIG":
			if EmmigratedSet {
				return ErrContext{"Individual", "EMIG", ErrSingleMultiple}
			}
			EmmigratedSet = true
			if err := s.Emmigrated.parse(sl); err != nil {
				return ErrContext{"Individual", "EMIG", err}
			}
		case "IMMI":
			if ImmigratedSet {
				return ErrContext{"Individual", "IMMI", ErrSingleMultiple}
			}
			ImmigratedSet = true
			if err := s.Immigrated.parse(sl); err != nil {
				return ErrContext{"Individual", "IMMI", err}
			}
		case "CENS":
			if CensusSet {
				return ErrContext{"Individual", "CENS", ErrSingleMultiple}
			}
			CensusSet = true
			if err := s.Census.parse(sl); err != nil {
				return ErrContext{"Individual", "CENS", err}
			}
		case "PROB":
			if ProbateSet {
				return ErrContext{"Individual", "PROB", ErrSingleMultiple}
			}
			ProbateSet = true
			if err := s.Probate.parse(sl); err != nil {
				return ErrContext{"Individual", "PROB", err}
			}
		case "WILL":
			if WillSet {
				return ErrContext{"Individual", "WILL", ErrSingleMultiple}
			}
			WillSet = true
			if err := s.Will.parse(sl); err != nil {
				return ErrContext{"Individual", "WILL", err}
			}
		case "GRAD":
			if GraduatedSet {
				return ErrContext{"Individual", "GRAD", ErrSingleMultiple}
			}
			GraduatedSet = true
			if err := s.Graduated.parse(sl); err != nil {
				return ErrContext{"Individual", "GRAD", err}
			}
		case "RETI":
			if RetiredSet {
				return ErrContext{"Individual", "RETI", ErrSingleMultiple}
			}
			RetiredSet = true
			if err := s.Retired.parse(sl); err != nil {
				return ErrContext{"Individual", "RETI", err}
			}
		case "EVEN":
			var t IndividualEventDetail
			if err != t.parse(sl); err != nil {
				return ErrContext{"Individual", "EVEN", err}
			}
			s.Events = append(s.Events, t)
		case "CAST":
			var t CasteEvent
			if err != t.parse(sl); err != nil {
				return ErrContext{"Individual", "CAST", err}
			}
			s.Caste = append(s.Caste, t)
		case "DSCR":
			var t DescriptionEvent
			if err != t.parse(sl); err != nil {
				return ErrContext{"Individual", "DSCR", err}
			}
			s.Description = append(s.Description, t)
		case "EDUC":
			var t ScholasticEvent
			if err != t.parse(sl); err != nil {
				return ErrContext{"Individual", "EDUC", err}
			}
			s.ScholasticAchievement = append(s.ScholasticAchievement, t)
		case "IDNO":
			var t NationalIDEvent
			if err != t.parse(sl); err != nil {
				return ErrContext{"Individual", "IDNO", err}
			}
			s.NationalID = append(s.NationalID, t)
		case "NATI":
			var t NationalOriginEvent
			if err != t.parse(sl); err != nil {
				return ErrContext{"Individual", "NATI", err}
			}
			s.NationalTribalOrigin = append(s.NationalTribalOrigin, t)
		case "NCHI":
			var t ChildrenEvent
			if err != t.parse(sl); err != nil {
				return ErrContext{"Individual", "NCHI", err}
			}
			s.CountOfChildren = append(s.CountOfChildren, t)
		case "NMR":
			var t MarriagesEvent
			if err != t.parse(sl); err != nil {
				return ErrContext{"Individual", "NMR", err}
			}
			s.CountOfMarriages = append(s.CountOfMarriages, t)
		case "OCCU":
			var t OccupationEvent
			if err != t.parse(sl); err != nil {
				return ErrContext{"Individual", "OCCU", err}
			}
			s.Occupation = append(s.Occupation, t)
		case "PROP":
			var t PossessionEvent
			if err != t.parse(sl); err != nil {
				return ErrContext{"Individual", "PROP", err}
			}
			s.Possessions = append(s.Possessions, t)
		case "RELI":
			var t ReligiousEvent
			if err != t.parse(sl); err != nil {
				return ErrContext{"Individual", "RELI", err}
			}
			s.ReligiousAffiliation = append(s.ReligiousAffiliation, t)
		case "RESI":
			var t ResidenceEvent
			if err != t.parse(sl); err != nil {
				return ErrContext{"Individual", "RESI", err}
			}
			s.Residences = append(s.Residences, t)
		case "SSN":
			var t SSNEvent
			if err != t.parse(sl); err != nil {
				return ErrContext{"Individual", "SSN", err}
			}
			s.SocialSecurity = append(s.SocialSecurity, t)
		case "TITL":
			var t NobilityEvent
			if err != t.parse(sl); err != nil {
				return ErrContext{"Individual", "TITL", err}
			}
			s.NobilityTypeTitle = append(s.NobilityTypeTitle, t)
		case "FAMC":
			var t ChildToFamilyLink
			if err != t.parse(sl); err != nil {
				return ErrContext{"Individual", "FAMC", err}
			}
			s.ChildOf = append(s.ChildOf, t)
		case "FAMS":
			var t SpouseToFamilyLink
			if err != t.parse(sl); err != nil {
				return ErrContext{"Individual", "FAMS", err}
			}
			s.SpouseOf = append(s.SpouseOf, t)
		case "SUBM":
			var t Xref
			if err != t.parse(sl); err != nil {
				return ErrContext{"Individual", "SUBM", err}
			}
			s.Submitters = append(s.Submitters, t)
		case "ASSO":
			var t AssociationStructure
			if err != t.parse(sl); err != nil {
				return ErrContext{"Individual", "ASSO", err}
			}
			s.Associations = append(s.Associations, t)
		case "ALIA":
			var t Xref
			if err != t.parse(sl); err != nil {
				return ErrContext{"Individual", "ALIA", err}
			}
			s.Aliases = append(s.Aliases, t)
		case "ANCI":
			var t Xref
			if err != t.parse(sl); err != nil {
				return ErrContext{"Individual", "ANCI", err}
			}
			s.AncestorInterest = append(s.AncestorInterest, t)
		case "DESI":
			var t Xref
			if err != t.parse(sl); err != nil {
				return ErrContext{"Individual", "DESI", err}
			}
			s.DescendentInterest = append(s.DescendentInterest, t)
		case "SOUR":
			var t SourceCitation
			if err != t.parse(sl); err != nil {
				return ErrContext{"Individual", "SOUR", err}
			}
			s.Sources = append(s.Sources, t)
		case "OBJE":
			var t MultimediaLink
			if err != t.parse(sl); err != nil {
				return ErrContext{"Individual", "OBJE", err}
			}
			s.Multimedia = append(s.Multimedia, t)
		case "NOTE":
			var t NoteStructure
			if err != t.parse(sl); err != nil {
				return ErrContext{"Individual", "NOTE", err}
			}
			s.Notes = append(s.Notes, t)
		case "RFN":
			if PermanentRecordSet {
				return ErrContext{"Individual", "RFN", ErrSingleMultiple}
			}
			PermanentRecordSet = true
			if err := s.PermanentRecord.parse(sl); err != nil {
				return ErrContext{"Individual", "RFN", err}
			}
		case "AFN":
			if AncestralFileNumberSet {
				return ErrContext{"Individual", "AFN", ErrSingleMultiple}
			}
			AncestralFileNumberSet = true
			if err := s.AncestralFileNumber.parse(sl); err != nil {
				return ErrContext{"Individual", "AFN", err}
			}
		case "REFN":
			var t UserReferenceStructure
			if err != t.parse(sl); err != nil {
				return ErrContext{"Individual", "REFN", err}
			}
			s.UserReferences = append(s.UserReferences, t)
		case "RIN":
			if AutomatedRecordIDSet {
				return ErrContext{"Individual", "RIN", ErrSingleMultiple}
			}
			AutomatedRecordIDSet = true
			if err := s.AutomatedRecordID.parse(sl); err != nil {
				return ErrContext{"Individual", "RIN", err}
			}
		case "CHAN":
			if ChangeDateSet {
				return ErrContext{"Individual", "CHAN", ErrSingleMultiple}
			}
			ChangeDateSet = true
			if err := s.ChangeDate.parse(sl); err != nil {
				return ErrContext{"Individual", "CHAN", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
				return ErrContext{"Individual", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// VerifiedIndividualFamEventDetail is a GEDCOM structure type
type VerifiedIndividualFamEventDetail struct {
	Famc                Xref
	VerifiedEventDetail VerifiedEventDetail
}

func (s *VerifiedIndividualFamEventDetail) parse(l Line) error {
	var VerifiedEventDetailSet, FamcSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "FAMC":
			if FamcSet {
				return ErrContext{"VerifiedIndividualFamEventDetail", "FAMC", ErrSingleMultiple}
			}
			FamcSet = true
			if err := s.Famc.parse(sl); err != nil {
				return ErrContext{"VerifiedIndividualFamEventDetail", "FAMC", err}
			}
		case "":
			if VerifiedEventDetailSet {
				return ErrContext{"VerifiedIndividualFamEventDetail", "", ErrSingleMultiple}
			}
			VerifiedEventDetailSet = true
			if err := s.VerifiedEventDetail.parse(sl); err != nil {
				return ErrContext{"VerifiedIndividualFamEventDetail", "", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
				return ErrContext{"VerifiedIndividualFamEventDetail", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	if !VerifiedEventDetailSet {
		return ErrContext{"VerifiedIndividualFamEventDetail", "VerifiedEventDetail", ErrRequiredMissing}
	}
	return nil
}

// VerifiedEventDetail is a GEDCOM structure type
type VerifiedEventDetail struct {
	Verified Verified
	Event    EventDetail
}

func (s *VerifiedEventDetail) parse(l Line) error {
	if err := s.Verified.parse(l); err != nil {
		return ErrContext{"VerifiedEventDetail", "line_value", err}
	}
	var EventSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "EVEN":
			if EventSet {
				return ErrContext{"VerifiedEventDetail", "EVEN", ErrSingleMultiple}
			}
			EventSet = true
			if err := s.Event.parse(sl); err != nil {
				return ErrContext{"VerifiedEventDetail", "EVEN", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
				return ErrContext{"VerifiedEventDetail", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// IndividualEventDetail is a GEDCOM structure type
type IndividualEventDetail struct {
	Event EventDetail
}

func (s *IndividualEventDetail) parse(l Line) error {
	var EventSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "EVEN":
			if EventSet {
				return ErrContext{"IndividualEventDetail", "EVEN", ErrSingleMultiple}
			}
			EventSet = true
			if err := s.Event.parse(sl); err != nil {
				return ErrContext{"IndividualEventDetail", "EVEN", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
				return ErrContext{"IndividualEventDetail", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// AdoptionEvent is a GEDCOM structure type
type AdoptionEvent struct {
	Family AdoptionReference
	Event  EventDetail
}

func (s *AdoptionEvent) parse(l Line) error {
	var FamilySet, EventSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "FAMC":
			if FamilySet {
				return ErrContext{"AdoptionEvent", "FAMC", ErrSingleMultiple}
			}
			FamilySet = true
			if err := s.Family.parse(sl); err != nil {
				return ErrContext{"AdoptionEvent", "FAMC", err}
			}
		case "EVEN":
			if EventSet {
				return ErrContext{"AdoptionEvent", "EVEN", ErrSingleMultiple}
			}
			EventSet = true
			if err := s.Event.parse(sl); err != nil {
				return ErrContext{"AdoptionEvent", "EVEN", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
				return ErrContext{"AdoptionEvent", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// AdoptionReference is a GEDCOM structure type
type AdoptionReference struct {
	ID        Xref
	AdoptedBy AdoptedBy
}

func (s *AdoptionReference) parse(l Line) error {
	if err := s.ID.parse(Line{value: l.xrefID}); err != nil {
		return ErrContext{"AdoptionReference", "xrefID", err}
	}
	var AdoptedBySet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "ADOP":
			if AdoptedBySet {
				return ErrContext{"AdoptionReference", "ADOP", ErrSingleMultiple}
			}
			AdoptedBySet = true
			if err := s.AdoptedBy.parse(sl); err != nil {
				return ErrContext{"AdoptionReference", "ADOP", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
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
	Event     EventDetail
}

func (s *CasteEvent) parse(l Line) error {
	if err := s.CasteName.parse(l); err != nil {
		return ErrContext{"CasteEvent", "line_value", err}
	}
	var EventSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "EVEN":
			if EventSet {
				return ErrContext{"CasteEvent", "EVEN", ErrSingleMultiple}
			}
			EventSet = true
			if err := s.Event.parse(sl); err != nil {
				return ErrContext{"CasteEvent", "EVEN", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
				return ErrContext{"CasteEvent", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// DescriptionEvent is a GEDCOM structure type
type DescriptionEvent struct {
	PhysicalDescription PhysicalDescription
	Event               EventDetail
}

func (s *DescriptionEvent) parse(l Line) error {
	if err := s.PhysicalDescription.parse(l); err != nil {
		return ErrContext{"DescriptionEvent", "line_value", err}
	}
	var EventSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "EVEN":
			if EventSet {
				return ErrContext{"DescriptionEvent", "EVEN", ErrSingleMultiple}
			}
			EventSet = true
			if err := s.Event.parse(sl); err != nil {
				return ErrContext{"DescriptionEvent", "EVEN", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
				return ErrContext{"DescriptionEvent", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// ScholasticEvent is a GEDCOM structure type
type ScholasticEvent struct {
	ScholasticAchievement ScholasticAchievement
	Event                 EventDetail
}

func (s *ScholasticEvent) parse(l Line) error {
	if err := s.ScholasticAchievement.parse(l); err != nil {
		return ErrContext{"ScholasticEvent", "line_value", err}
	}
	var EventSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "EVEN":
			if EventSet {
				return ErrContext{"ScholasticEvent", "EVEN", ErrSingleMultiple}
			}
			EventSet = true
			if err := s.Event.parse(sl); err != nil {
				return ErrContext{"ScholasticEvent", "EVEN", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
				return ErrContext{"ScholasticEvent", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// NationalIDEvent is a GEDCOM structure type
type NationalIDEvent struct {
	NationalIDNumber NationalIDNumber
	Event            EventDetail
}

func (s *NationalIDEvent) parse(l Line) error {
	if err := s.NationalIDNumber.parse(l); err != nil {
		return ErrContext{"NationalIDEvent", "line_value", err}
	}
	var EventSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "EVEN":
			if EventSet {
				return ErrContext{"NationalIDEvent", "EVEN", ErrSingleMultiple}
			}
			EventSet = true
			if err := s.Event.parse(sl); err != nil {
				return ErrContext{"NationalIDEvent", "EVEN", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
				return ErrContext{"NationalIDEvent", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// NationalOriginEvent is a GEDCOM structure type
type NationalOriginEvent struct {
	NationalOrTribalOrigin NationalOrTribalOrigin
	Event                  EventDetail
}

func (s *NationalOriginEvent) parse(l Line) error {
	if err := s.NationalOrTribalOrigin.parse(l); err != nil {
		return ErrContext{"NationalOriginEvent", "line_value", err}
	}
	var EventSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "EVEN":
			if EventSet {
				return ErrContext{"NationalOriginEvent", "EVEN", ErrSingleMultiple}
			}
			EventSet = true
			if err := s.Event.parse(sl); err != nil {
				return ErrContext{"NationalOriginEvent", "EVEN", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
				return ErrContext{"NationalOriginEvent", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// ChildrenEvent is a GEDCOM structure type
type ChildrenEvent struct {
	CountOfChildren CountOfChildren
	Event           EventDetail
}

func (s *ChildrenEvent) parse(l Line) error {
	if err := s.CountOfChildren.parse(l); err != nil {
		return ErrContext{"ChildrenEvent", "line_value", err}
	}
	var EventSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "EVEN":
			if EventSet {
				return ErrContext{"ChildrenEvent", "EVEN", ErrSingleMultiple}
			}
			EventSet = true
			if err := s.Event.parse(sl); err != nil {
				return ErrContext{"ChildrenEvent", "EVEN", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
				return ErrContext{"ChildrenEvent", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// MarriagesEvent is a GEDCOM structure type
type MarriagesEvent struct {
	CountOfMarriages CountOfMarriages
	Event            EventDetail
}

func (s *MarriagesEvent) parse(l Line) error {
	if err := s.CountOfMarriages.parse(l); err != nil {
		return ErrContext{"MarriagesEvent", "line_value", err}
	}
	var EventSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "EVEN":
			if EventSet {
				return ErrContext{"MarriagesEvent", "EVEN", ErrSingleMultiple}
			}
			EventSet = true
			if err := s.Event.parse(sl); err != nil {
				return ErrContext{"MarriagesEvent", "EVEN", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
				return ErrContext{"MarriagesEvent", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// OccupationEvent is a GEDCOM structure type
type OccupationEvent struct {
	Occupation Occupation
	Event      EventDetail
}

func (s *OccupationEvent) parse(l Line) error {
	if err := s.Occupation.parse(l); err != nil {
		return ErrContext{"OccupationEvent", "line_value", err}
	}
	var EventSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "EVEN":
			if EventSet {
				return ErrContext{"OccupationEvent", "EVEN", ErrSingleMultiple}
			}
			EventSet = true
			if err := s.Event.parse(sl); err != nil {
				return ErrContext{"OccupationEvent", "EVEN", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
				return ErrContext{"OccupationEvent", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// PossessionEvent is a GEDCOM structure type
type PossessionEvent struct {
	Possessions Possessions
	Event       EventDetail
}

func (s *PossessionEvent) parse(l Line) error {
	if err := s.Possessions.parse(l); err != nil {
		return ErrContext{"PossessionEvent", "line_value", err}
	}
	var EventSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "EVEN":
			if EventSet {
				return ErrContext{"PossessionEvent", "EVEN", ErrSingleMultiple}
			}
			EventSet = true
			if err := s.Event.parse(sl); err != nil {
				return ErrContext{"PossessionEvent", "EVEN", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
				return ErrContext{"PossessionEvent", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// ReligiousEvent is a GEDCOM structure type
type ReligiousEvent struct {
	ReligiousAffiliation ReligiousAffiliation
	Event                EventDetail
}

func (s *ReligiousEvent) parse(l Line) error {
	if err := s.ReligiousAffiliation.parse(l); err != nil {
		return ErrContext{"ReligiousEvent", "line_value", err}
	}
	var EventSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "EVEN":
			if EventSet {
				return ErrContext{"ReligiousEvent", "EVEN", ErrSingleMultiple}
			}
			EventSet = true
			if err := s.Event.parse(sl); err != nil {
				return ErrContext{"ReligiousEvent", "EVEN", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
				return ErrContext{"ReligiousEvent", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// ResidenceEvent is a GEDCOM structure type
type ResidenceEvent struct {
	Event EventDetail
}

func (s *ResidenceEvent) parse(l Line) error {
	var EventSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "EVEN":
			if EventSet {
				return ErrContext{"ResidenceEvent", "EVEN", ErrSingleMultiple}
			}
			EventSet = true
			if err := s.Event.parse(sl); err != nil {
				return ErrContext{"ResidenceEvent", "EVEN", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
				return ErrContext{"ResidenceEvent", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// SSNEvent is a GEDCOM structure type
type SSNEvent struct {
	SocialSecurityNumber SocialSecurityNumber
	Event                EventDetail
}

func (s *SSNEvent) parse(l Line) error {
	if err := s.SocialSecurityNumber.parse(l); err != nil {
		return ErrContext{"SSNEvent", "line_value", err}
	}
	var EventSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "EVEN":
			if EventSet {
				return ErrContext{"SSNEvent", "EVEN", ErrSingleMultiple}
			}
			EventSet = true
			if err := s.Event.parse(sl); err != nil {
				return ErrContext{"SSNEvent", "EVEN", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
				return ErrContext{"SSNEvent", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// NobilityEvent is a GEDCOM structure type
type NobilityEvent struct {
	NobilityTypeTitle NobilityTypeTitle
	Event             EventDetail
}

func (s *NobilityEvent) parse(l Line) error {
	if err := s.NobilityTypeTitle.parse(l); err != nil {
		return ErrContext{"NobilityEvent", "line_value", err}
	}
	var EventSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "EVEN":
			if EventSet {
				return ErrContext{"NobilityEvent", "EVEN", ErrSingleMultiple}
			}
			EventSet = true
			if err := s.Event.parse(sl); err != nil {
				return ErrContext{"NobilityEvent", "EVEN", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
				return ErrContext{"NobilityEvent", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	return nil
}

// UserReferenceStructure is a GEDCOM structure type
type UserReferenceStructure struct {
	UserReferenceNumber UserReferenceNumber
	Type                UserReferenceType
}

func (s *UserReferenceStructure) parse(l Line) error {
	if err := s.UserReferenceNumber.parse(l); err != nil {
		return ErrContext{"UserReferenceStructure", "line_value", err}
	}
	var TypeSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "TYPE":
			if TypeSet {
				return ErrContext{"UserReferenceStructure", "TYPE", ErrSingleMultiple}
			}
			TypeSet = true
			if err := s.Type.parse(sl); err != nil {
				return ErrContext{"UserReferenceStructure", "TYPE", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
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

func (s *MultimediaRecord) parse(l Line) error {
	if err := s.ID.parse(Line{value: l.xrefID}); err != nil {
		return ErrContext{"MultimediaRecord", "xrefID", err}
	}
	var FormatSet, BlobSet, TitleSet, ContinuedObjectSet, AutomatedRecordIDSet, ChangeDateSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "FORM":
			if FormatSet {
				return ErrContext{"MultimediaRecord", "FORM", ErrSingleMultiple}
			}
			FormatSet = true
			if err := s.Format.parse(sl); err != nil {
				return ErrContext{"MultimediaRecord", "FORM", err}
			}
		case "TITLE":
			if TitleSet {
				return ErrContext{"MultimediaRecord", "TITLE", ErrSingleMultiple}
			}
			TitleSet = true
			if err := s.Title.parse(sl); err != nil {
				return ErrContext{"MultimediaRecord", "TITLE", err}
			}
		case "NOTE":
			var t NoteStructure
			if err != t.parse(sl); err != nil {
				return ErrContext{"MultimediaRecord", "NOTE", err}
			}
			s.Notes = append(s.Notes, t)
		case "BLOB":
			if BlobSet {
				return ErrContext{"MultimediaRecord", "BLOB", ErrSingleMultiple}
			}
			BlobSet = true
			if err := s.Blob.parse(sl); err != nil {
				return ErrContext{"MultimediaRecord", "BLOB", err}
			}
		case "OBJE":
			if ContinuedObjectSet {
				return ErrContext{"MultimediaRecord", "OBJE", ErrSingleMultiple}
			}
			ContinuedObjectSet = true
			if err := s.ContinuedObject.parse(sl); err != nil {
				return ErrContext{"MultimediaRecord", "OBJE", err}
			}
		case "REFN":
			var t UserReferenceStructure
			if err != t.parse(sl); err != nil {
				return ErrContext{"MultimediaRecord", "REFN", err}
			}
			s.UserReferences = append(s.UserReferences, t)
		case "RIN":
			if AutomatedRecordIDSet {
				return ErrContext{"MultimediaRecord", "RIN", ErrSingleMultiple}
			}
			AutomatedRecordIDSet = true
			if err := s.AutomatedRecordID.parse(sl); err != nil {
				return ErrContext{"MultimediaRecord", "RIN", err}
			}
		case "CHAN":
			if ChangeDateSet {
				return ErrContext{"MultimediaRecord", "CHAN", ErrSingleMultiple}
			}
			ChangeDateSet = true
			if err := s.ChangeDate.parse(sl); err != nil {
				return ErrContext{"MultimediaRecord", "CHAN", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
				return ErrContext{"MultimediaRecord", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	if !FormatSet {
		return ErrContext{"MultimediaRecord", "Format", ErrRequiredMissing}
	}
	if !BlobSet {
		return ErrContext{"MultimediaRecord", "Blob", ErrRequiredMissing}
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

func (s *NoteRecord) parse(l Line) error {
	if err := s.ID.parse(Line{value: l.xrefID}); err != nil {
		return ErrContext{"NoteRecord", "xrefID", err}
	}
	if err := s.SubmitterText.parse(l); err != nil {
		return ErrContext{"NoteRecord", "line_value", err}
	}
	var AutomatedRecordIDSet, ChangeDateSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "SOUR":
			var t SourceCitation
			if err != t.parse(sl); err != nil {
				return ErrContext{"NoteRecord", "SOUR", err}
			}
			s.Sources = append(s.Sources, t)
		case "REFN":
			var t UserReferenceStructure
			if err != t.parse(sl); err != nil {
				return ErrContext{"NoteRecord", "REFN", err}
			}
			s.UserReferences = append(s.UserReferences, t)
		case "RIN":
			if AutomatedRecordIDSet {
				return ErrContext{"NoteRecord", "RIN", ErrSingleMultiple}
			}
			AutomatedRecordIDSet = true
			if err := s.AutomatedRecordID.parse(sl); err != nil {
				return ErrContext{"NoteRecord", "RIN", err}
			}
		case "CHAN":
			if ChangeDateSet {
				return ErrContext{"NoteRecord", "CHAN", ErrSingleMultiple}
			}
			ChangeDateSet = true
			if err := s.ChangeDate.parse(sl); err != nil {
				return ErrContext{"NoteRecord", "CHAN", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
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

func (s *RepositoryRecord) parse(l Line) error {
	if err := s.ID.parse(Line{value: l.xrefID}); err != nil {
		return ErrContext{"RepositoryRecord", "xrefID", err}
	}
	var NameOfRepositorySet, AddressSet, AutomatedRecordIDSet, ChangeDateSet bool
	var PhoneNumberCount int
	for _, sl := range l.Sub {
		switch sl.tag {
		case "Name":
			if NameOfRepositorySet {
				return ErrContext{"RepositoryRecord", "Name", ErrSingleMultiple}
			}
			NameOfRepositorySet = true
			if err := s.NameOfRepository.parse(sl); err != nil {
				return ErrContext{"RepositoryRecord", "Name", err}
			}
		case "ADDR":
			if AddressSet {
				return ErrContext{"RepositoryRecord", "ADDR", ErrSingleMultiple}
			}
			AddressSet = true
			if err := s.Address.parse(sl); err != nil {
				return ErrContext{"RepositoryRecord", "ADDR", err}
			}
		case "PHON":
			if PhoneNumberCount == 3 {
				return ErrContext{"RepositoryRecord", "PHON", ErrTooMany(3)}
			}
			PhoneNumberCount++
			var t PhoneNumber
			if err != t.parse(sl); err != nil {
				return ErrContext{"RepositoryRecord", "PHON", err}
			}
			s.PhoneNumber = append(s.PhoneNumber, t)
		case "NOTE":
			var t NoteStructure
			if err != t.parse(sl); err != nil {
				return ErrContext{"RepositoryRecord", "NOTE", err}
			}
			s.Notes = append(s.Notes, t)
		case "REFN":
			var t UserReferenceStructure
			if err != t.parse(sl); err != nil {
				return ErrContext{"RepositoryRecord", "REFN", err}
			}
			s.UserReferences = append(s.UserReferences, t)
		case "RIN":
			if AutomatedRecordIDSet {
				return ErrContext{"RepositoryRecord", "RIN", ErrSingleMultiple}
			}
			AutomatedRecordIDSet = true
			if err := s.AutomatedRecordID.parse(sl); err != nil {
				return ErrContext{"RepositoryRecord", "RIN", err}
			}
		case "CHAN":
			if ChangeDateSet {
				return ErrContext{"RepositoryRecord", "CHAN", ErrSingleMultiple}
			}
			ChangeDateSet = true
			if err := s.ChangeDate.parse(sl); err != nil {
				return ErrContext{"RepositoryRecord", "CHAN", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
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

func (s *SourceRecord) parse(l Line) error {
	if err := s.ID.parse(Line{value: l.xrefID}); err != nil {
		return ErrContext{"SourceRecord", "xrefID", err}
	}
	var SourceRepositoryCitationSet, DataSet, OriginatorSet, TitleSet, FiledBySet, PublicationFactsSet, TextFromSourceSet, ContinuedObjectSet, AutomatedRecordIDSet, ChangeDateSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "DATA":
			if DataSet {
				return ErrContext{"SourceRecord", "DATA", ErrSingleMultiple}
			}
			DataSet = true
			if err := s.Data.parse(sl); err != nil {
				return ErrContext{"SourceRecord", "DATA", err}
			}
		case "AUTH":
			if OriginatorSet {
				return ErrContext{"SourceRecord", "AUTH", ErrSingleMultiple}
			}
			OriginatorSet = true
			if err := s.Originator.parse(sl); err != nil {
				return ErrContext{"SourceRecord", "AUTH", err}
			}
		case "TITL":
			if TitleSet {
				return ErrContext{"SourceRecord", "TITL", ErrSingleMultiple}
			}
			TitleSet = true
			if err := s.Title.parse(sl); err != nil {
				return ErrContext{"SourceRecord", "TITL", err}
			}
		case "ABBR":
			if FiledBySet {
				return ErrContext{"SourceRecord", "ABBR", ErrSingleMultiple}
			}
			FiledBySet = true
			if err := s.FiledBy.parse(sl); err != nil {
				return ErrContext{"SourceRecord", "ABBR", err}
			}
		case "PUBL":
			if PublicationFactsSet {
				return ErrContext{"SourceRecord", "PUBL", ErrSingleMultiple}
			}
			PublicationFactsSet = true
			if err := s.PublicationFacts.parse(sl); err != nil {
				return ErrContext{"SourceRecord", "PUBL", err}
			}
		case "TEXT":
			if TextFromSourceSet {
				return ErrContext{"SourceRecord", "TEXT", ErrSingleMultiple}
			}
			TextFromSourceSet = true
			if err := s.TextFromSource.parse(sl); err != nil {
				return ErrContext{"SourceRecord", "TEXT", err}
			}
		case "REPO":
			if SourceRepositoryCitationSet {
				return ErrContext{"SourceRecord", "REPO", ErrSingleMultiple}
			}
			SourceRepositoryCitationSet = true
			if err := s.SourceRepositoryCitation.parse(sl); err != nil {
				return ErrContext{"SourceRecord", "REPO", err}
			}
		case "OBJE":
			if ContinuedObjectSet {
				return ErrContext{"SourceRecord", "OBJE", ErrSingleMultiple}
			}
			ContinuedObjectSet = true
			if err := s.ContinuedObject.parse(sl); err != nil {
				return ErrContext{"SourceRecord", "OBJE", err}
			}
		case "NOTE":
			var t NoteStructure
			if err != t.parse(sl); err != nil {
				return ErrContext{"SourceRecord", "NOTE", err}
			}
			s.Notes = append(s.Notes, t)
		case "REFN":
			var t UserReferenceStructure
			if err != t.parse(sl); err != nil {
				return ErrContext{"SourceRecord", "REFN", err}
			}
			s.UserReferences = append(s.UserReferences, t)
		case "RIN":
			if AutomatedRecordIDSet {
				return ErrContext{"SourceRecord", "RIN", ErrSingleMultiple}
			}
			AutomatedRecordIDSet = true
			if err := s.AutomatedRecordID.parse(sl); err != nil {
				return ErrContext{"SourceRecord", "RIN", err}
			}
		case "CHAN":
			if ChangeDateSet {
				return ErrContext{"SourceRecord", "CHAN", ErrSingleMultiple}
			}
			ChangeDateSet = true
			if err := s.ChangeDate.parse(sl); err != nil {
				return ErrContext{"SourceRecord", "CHAN", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
				return ErrContext{"SourceRecord", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	if !SourceRepositoryCitationSet {
		return ErrContext{"SourceRecord", "SourceRepositoryCitation", ErrRequiredMissing}
	}
	return nil
}

// SourceRecordDataStructure is a GEDCOM structure type
type SourceRecordDataStructure struct {
	EventsRecorded    []EventsRecordedStructure
	ResponsibleAgency ResponsibleAgency
	Notes             []NoteStructure
}

func (s *SourceRecordDataStructure) parse(l Line) error {
	var ResponsibleAgencySet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "EVEN":
			var t EventsRecordedStructure
			if err != t.parse(sl); err != nil {
				return ErrContext{"SourceRecordDataStructure", "EVEN", err}
			}
			s.EventsRecorded = append(s.EventsRecorded, t)
		case "AGNC":
			if ResponsibleAgencySet {
				return ErrContext{"SourceRecordDataStructure", "AGNC", ErrSingleMultiple}
			}
			ResponsibleAgencySet = true
			if err := s.ResponsibleAgency.parse(sl); err != nil {
				return ErrContext{"SourceRecordDataStructure", "AGNC", err}
			}
		case "NOTE":
			var t NoteStructure
			if err != t.parse(sl); err != nil {
				return ErrContext{"SourceRecordDataStructure", "NOTE", err}
			}
			s.Notes = append(s.Notes, t)
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
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

func (s *EventsRecordedStructure) parse(l Line) error {
	var DatePeriodSet, SourceJurisdictionPlaceSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "DATE":
			if DatePeriodSet {
				return ErrContext{"EventsRecordedStructure", "DATE", ErrSingleMultiple}
			}
			DatePeriodSet = true
			if err := s.DatePeriod.parse(sl); err != nil {
				return ErrContext{"EventsRecordedStructure", "DATE", err}
			}
		case "PLACE":
			if SourceJurisdictionPlaceSet {
				return ErrContext{"EventsRecordedStructure", "PLACE", ErrSingleMultiple}
			}
			SourceJurisdictionPlaceSet = true
			if err := s.SourceJurisdictionPlace.parse(sl); err != nil {
				return ErrContext{"EventsRecordedStructure", "PLACE", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
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

func (s *SubmissionRecord) parse(l Line) error {
	if err := s.ID.parse(Line{value: l.xrefID}); err != nil {
		return ErrContext{"SubmissionRecord", "xrefID", err}
	}
	var NameOfFamilyFileSet, TempleCodeSet, GenerationsOfAncestorsSet, GenerationsOfDescendantsSet, OrdinanceProcessFlagSet, AutomatedRecordIDSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "FAMF":
			if NameOfFamilyFileSet {
				return ErrContext{"SubmissionRecord", "FAMF", ErrSingleMultiple}
			}
			NameOfFamilyFileSet = true
			if err := s.NameOfFamilyFile.parse(sl); err != nil {
				return ErrContext{"SubmissionRecord", "FAMF", err}
			}
		case "TEMP":
			if TempleCodeSet {
				return ErrContext{"SubmissionRecord", "TEMP", ErrSingleMultiple}
			}
			TempleCodeSet = true
			if err := s.TempleCode.parse(sl); err != nil {
				return ErrContext{"SubmissionRecord", "TEMP", err}
			}
		case "ANCE":
			if GenerationsOfAncestorsSet {
				return ErrContext{"SubmissionRecord", "ANCE", ErrSingleMultiple}
			}
			GenerationsOfAncestorsSet = true
			if err := s.GenerationsOfAncestors.parse(sl); err != nil {
				return ErrContext{"SubmissionRecord", "ANCE", err}
			}
		case "DESC":
			if GenerationsOfDescendantsSet {
				return ErrContext{"SubmissionRecord", "DESC", ErrSingleMultiple}
			}
			GenerationsOfDescendantsSet = true
			if err := s.GenerationsOfDescendants.parse(sl); err != nil {
				return ErrContext{"SubmissionRecord", "DESC", err}
			}
		case "ORDI":
			if OrdinanceProcessFlagSet {
				return ErrContext{"SubmissionRecord", "ORDI", ErrSingleMultiple}
			}
			OrdinanceProcessFlagSet = true
			if err := s.OrdinanceProcessFlag.parse(sl); err != nil {
				return ErrContext{"SubmissionRecord", "ORDI", err}
			}
		case "RIN":
			if AutomatedRecordIDSet {
				return ErrContext{"SubmissionRecord", "RIN", ErrSingleMultiple}
			}
			AutomatedRecordIDSet = true
			if err := s.AutomatedRecordID.parse(sl); err != nil {
				return ErrContext{"SubmissionRecord", "RIN", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
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

func (s *SubmitterRecord) parse(l Line) error {
	if err := s.ID.parse(Line{value: l.xrefID}); err != nil {
		return ErrContext{"SubmitterRecord", "xrefID", err}
	}
	var SubmitterNameSet, AddressSet, SubmitterRegisteredRFNSet, AutomatedRecordIDSet, ChangeDateSet bool
	var PhoneNumberCount, LanguagePreferenceCount int
	for _, sl := range l.Sub {
		switch sl.tag {
		case "NAME":
			if SubmitterNameSet {
				return ErrContext{"SubmitterRecord", "NAME", ErrSingleMultiple}
			}
			SubmitterNameSet = true
			if err := s.SubmitterName.parse(sl); err != nil {
				return ErrContext{"SubmitterRecord", "NAME", err}
			}
		case "ADDR":
			if AddressSet {
				return ErrContext{"SubmitterRecord", "ADDR", ErrSingleMultiple}
			}
			AddressSet = true
			if err := s.Address.parse(sl); err != nil {
				return ErrContext{"SubmitterRecord", "ADDR", err}
			}
		case "PHON":
			if PhoneNumberCount == 3 {
				return ErrContext{"SubmitterRecord", "PHON", ErrTooMany(3)}
			}
			PhoneNumberCount++
			var t PhoneNumber
			if err != t.parse(sl); err != nil {
				return ErrContext{"SubmitterRecord", "PHON", err}
			}
			s.PhoneNumber = append(s.PhoneNumber, t)
		case "OBJE":
			var t MultimediaLink
			if err != t.parse(sl); err != nil {
				return ErrContext{"SubmitterRecord", "OBJE", err}
			}
			s.Multimedia = append(s.Multimedia, t)
		case "LANG":
			if LanguagePreferenceCount == 3 {
				return ErrContext{"SubmitterRecord", "LANG", ErrTooMany(3)}
			}
			LanguagePreferenceCount++
			var t LanguagePreference
			if err != t.parse(sl); err != nil {
				return ErrContext{"SubmitterRecord", "LANG", err}
			}
			s.LanguagePreference = append(s.LanguagePreference, t)
		case "RFN":
			if SubmitterRegisteredRFNSet {
				return ErrContext{"SubmitterRecord", "RFN", ErrSingleMultiple}
			}
			SubmitterRegisteredRFNSet = true
			if err := s.SubmitterRegisteredRFN.parse(sl); err != nil {
				return ErrContext{"SubmitterRecord", "RFN", err}
			}
		case "RIN":
			if AutomatedRecordIDSet {
				return ErrContext{"SubmitterRecord", "RIN", ErrSingleMultiple}
			}
			AutomatedRecordIDSet = true
			if err := s.AutomatedRecordID.parse(sl); err != nil {
				return ErrContext{"SubmitterRecord", "RIN", err}
			}
		case "CHAN":
			if ChangeDateSet {
				return ErrContext{"SubmitterRecord", "CHAN", ErrSingleMultiple}
			}
			ChangeDateSet = true
			if err := s.ChangeDate.parse(sl); err != nil {
				return ErrContext{"SubmitterRecord", "CHAN", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
				return ErrContext{"SubmitterRecord", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	if !SubmitterNameSet {
		return ErrContext{"SubmitterRecord", "SubmitterName", ErrRequiredMissing}
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

func (s *AddressStructure) parse(l Line) error {
	if err := s.AddressLine.parse(l); err != nil {
		return ErrContext{"AddressStructure", "line_value", err}
	}
	var AddressLine1Set, AddressLine2Set, CitySet, StateSet, PostalCodeSet, CountrySet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "ADR1":
			if AddressLine1Set {
				return ErrContext{"AddressStructure", "ADR1", ErrSingleMultiple}
			}
			AddressLine1Set = true
			if err := s.AddressLine1.parse(sl); err != nil {
				return ErrContext{"AddressStructure", "ADR1", err}
			}
		case "ADR2":
			if AddressLine2Set {
				return ErrContext{"AddressStructure", "ADR2", ErrSingleMultiple}
			}
			AddressLine2Set = true
			if err := s.AddressLine2.parse(sl); err != nil {
				return ErrContext{"AddressStructure", "ADR2", err}
			}
		case "CITY":
			if CitySet {
				return ErrContext{"AddressStructure", "CITY", ErrSingleMultiple}
			}
			CitySet = true
			if err := s.City.parse(sl); err != nil {
				return ErrContext{"AddressStructure", "CITY", err}
			}
		case "STAE":
			if StateSet {
				return ErrContext{"AddressStructure", "STAE", ErrSingleMultiple}
			}
			StateSet = true
			if err := s.State.parse(sl); err != nil {
				return ErrContext{"AddressStructure", "STAE", err}
			}
		case "POST":
			if PostalCodeSet {
				return ErrContext{"AddressStructure", "POST", ErrSingleMultiple}
			}
			PostalCodeSet = true
			if err := s.PostalCode.parse(sl); err != nil {
				return ErrContext{"AddressStructure", "POST", err}
			}
		case "CTRY":
			if CountrySet {
				return ErrContext{"AddressStructure", "CTRY", ErrSingleMultiple}
			}
			CountrySet = true
			if err := s.Country.parse(sl); err != nil {
				return ErrContext{"AddressStructure", "CTRY", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
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

func (s *AssociationStructure) parse(l Line) error {
	if err := s.ID.parse(Line{value: l.xrefID}); err != nil {
		return ErrContext{"AssociationStructure", "xrefID", err}
	}
	var RecordTypeSet, RelationSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "TYPE":
			if RecordTypeSet {
				return ErrContext{"AssociationStructure", "TYPE", ErrSingleMultiple}
			}
			RecordTypeSet = true
			if err := s.RecordType.parse(sl); err != nil {
				return ErrContext{"AssociationStructure", "TYPE", err}
			}
		case "RELA":
			if RelationSet {
				return ErrContext{"AssociationStructure", "RELA", ErrSingleMultiple}
			}
			RelationSet = true
			if err := s.Relation.parse(sl); err != nil {
				return ErrContext{"AssociationStructure", "RELA", err}
			}
		case "NOTE":
			var t NoteStructure
			if err != t.parse(sl); err != nil {
				return ErrContext{"AssociationStructure", "NOTE", err}
			}
			s.Notes = append(s.Notes, t)
		case "SOUR":
			var t SourceCitation
			if err != t.parse(sl); err != nil {
				return ErrContext{"AssociationStructure", "SOUR", err}
			}
			s.Sources = append(s.Sources, t)
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
				return ErrContext{"AssociationStructure", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	if !RecordTypeSet {
		return ErrContext{"AssociationStructure", "RecordType", ErrRequiredMissing}
	}
	if !RelationSet {
		return ErrContext{"AssociationStructure", "Relation", ErrRequiredMissing}
	}
	return nil
}

// ChangeDateStructure is a GEDCOM structure type
type ChangeDateStructure struct {
	Date  ChangeDateTime
	Notes []NoteStructure
}

func (s *ChangeDateStructure) parse(l Line) error {
	var DateSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "DATE":
			if DateSet {
				return ErrContext{"ChangeDateStructure", "DATE", ErrSingleMultiple}
			}
			DateSet = true
			if err := s.Date.parse(sl); err != nil {
				return ErrContext{"ChangeDateStructure", "DATE", err}
			}
		case "NOTE":
			var t NoteStructure
			if err != t.parse(sl); err != nil {
				return ErrContext{"ChangeDateStructure", "NOTE", err}
			}
			s.Notes = append(s.Notes, t)
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
				return ErrContext{"ChangeDateStructure", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	if !DateSet {
		return ErrContext{"ChangeDateStructure", "Date", ErrRequiredMissing}
	}
	return nil
}

// ChangeDateTime is a GEDCOM structure type
type ChangeDateTime struct {
	ChangeDate ChangeDate
	Time       TimeValue
}

func (s *ChangeDateTime) parse(l Line) error {
	if err := s.ChangeDate.parse(l); err != nil {
		return ErrContext{"ChangeDateTime", "line_value", err}
	}
	var TimeSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "TIME":
			if TimeSet {
				return ErrContext{"ChangeDateTime", "TIME", ErrSingleMultiple}
			}
			TimeSet = true
			if err := s.Time.parse(sl); err != nil {
				return ErrContext{"ChangeDateTime", "TIME", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
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

func (s *ChildToFamilyLink) parse(l Line) error {
	if err := s.ID.parse(Line{value: l.xrefID}); err != nil {
		return ErrContext{"ChildToFamilyLink", "xrefID", err}
	}
	for _, sl := range l.Sub {
		switch sl.tag {
		case "PEDI":
			var t PedigreeLinkageType
			if err != t.parse(sl); err != nil {
				return ErrContext{"ChildToFamilyLink", "PEDI", err}
			}
			s.PedigreeLinkageType = append(s.PedigreeLinkageType, t)
		case "NOTE":
			var t NoteStructure
			if err != t.parse(sl); err != nil {
				return ErrContext{"ChildToFamilyLink", "NOTE", err}
			}
			s.Notes = append(s.Notes, t)
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
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

func (s *EventDetail) parse(l Line) error {
	var TypeSet, DateSet, PlaceSet, AddressSet, AgeSet, ResponsibleAgencySet, CauseOfEventSet bool
	var PhoneNumberCount int
	for _, sl := range l.Sub {
		switch sl.tag {
		case "TYPE":
			if TypeSet {
				return ErrContext{"EventDetail", "TYPE", ErrSingleMultiple}
			}
			TypeSet = true
			if err := s.Type.parse(sl); err != nil {
				return ErrContext{"EventDetail", "TYPE", err}
			}
		case "DATE":
			if DateSet {
				return ErrContext{"EventDetail", "DATE", ErrSingleMultiple}
			}
			DateSet = true
			if err := s.Date.parse(sl); err != nil {
				return ErrContext{"EventDetail", "DATE", err}
			}
		case "PLAC":
			if PlaceSet {
				return ErrContext{"EventDetail", "PLAC", ErrSingleMultiple}
			}
			PlaceSet = true
			if err := s.Place.parse(sl); err != nil {
				return ErrContext{"EventDetail", "PLAC", err}
			}
		case "ADDR":
			if AddressSet {
				return ErrContext{"EventDetail", "ADDR", ErrSingleMultiple}
			}
			AddressSet = true
			if err := s.Address.parse(sl); err != nil {
				return ErrContext{"EventDetail", "ADDR", err}
			}
		case "PHON":
			if PhoneNumberCount == 3 {
				return ErrContext{"EventDetail", "PHON", ErrTooMany(3)}
			}
			PhoneNumberCount++
			var t PhoneNumber
			if err != t.parse(sl); err != nil {
				return ErrContext{"EventDetail", "PHON", err}
			}
			s.PhoneNumber = append(s.PhoneNumber, t)
		case "AGE":
			if AgeSet {
				return ErrContext{"EventDetail", "AGE", ErrSingleMultiple}
			}
			AgeSet = true
			if err := s.Age.parse(sl); err != nil {
				return ErrContext{"EventDetail", "AGE", err}
			}
		case "AGNC":
			if ResponsibleAgencySet {
				return ErrContext{"EventDetail", "AGNC", ErrSingleMultiple}
			}
			ResponsibleAgencySet = true
			if err := s.ResponsibleAgency.parse(sl); err != nil {
				return ErrContext{"EventDetail", "AGNC", err}
			}
		case "CAUS":
			if CauseOfEventSet {
				return ErrContext{"EventDetail", "CAUS", ErrSingleMultiple}
			}
			CauseOfEventSet = true
			if err := s.CauseOfEvent.parse(sl); err != nil {
				return ErrContext{"EventDetail", "CAUS", err}
			}
		case "SOUR":
			var t SourceCitation
			if err != t.parse(sl); err != nil {
				return ErrContext{"EventDetail", "SOUR", err}
			}
			s.Sources = append(s.Sources, t)
		case "OBJE":
			var t MultimediaLink
			if err != t.parse(sl); err != nil {
				return ErrContext{"EventDetail", "OBJE", err}
			}
			s.Multimedia = append(s.Multimedia, t)
		case "NOTE":
			var t NoteStructure
			if err != t.parse(sl); err != nil {
				return ErrContext{"EventDetail", "NOTE", err}
			}
			s.Notes = append(s.Notes, t)
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
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

func (s *LDSSpouseSealing) parse(l Line) error {
	var LDSSpouseSealingDateStatusSet, DateSet, TempleCodeSet, PlaceSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "STAT":
			if LDSSpouseSealingDateStatusSet {
				return ErrContext{"LDSSpouseSealing", "STAT", ErrSingleMultiple}
			}
			LDSSpouseSealingDateStatusSet = true
			if err := s.LDSSpouseSealingDateStatus.parse(sl); err != nil {
				return ErrContext{"LDSSpouseSealing", "STAT", err}
			}
		case "DATE":
			if DateSet {
				return ErrContext{"LDSSpouseSealing", "DATE", ErrSingleMultiple}
			}
			DateSet = true
			if err := s.Date.parse(sl); err != nil {
				return ErrContext{"LDSSpouseSealing", "DATE", err}
			}
		case "TEMP":
			if TempleCodeSet {
				return ErrContext{"LDSSpouseSealing", "TEMP", ErrSingleMultiple}
			}
			TempleCodeSet = true
			if err := s.TempleCode.parse(sl); err != nil {
				return ErrContext{"LDSSpouseSealing", "TEMP", err}
			}
		case "PLAC":
			if PlaceSet {
				return ErrContext{"LDSSpouseSealing", "PLAC", ErrSingleMultiple}
			}
			PlaceSet = true
			if err := s.Place.parse(sl); err != nil {
				return ErrContext{"LDSSpouseSealing", "PLAC", err}
			}
		case "SOUR":
			var t SourceCitation
			if err != t.parse(sl); err != nil {
				return ErrContext{"LDSSpouseSealing", "SOUR", err}
			}
			s.Sources = append(s.Sources, t)
		case "NOTE":
			var t NoteStructure
			if err != t.parse(sl); err != nil {
				return ErrContext{"LDSSpouseSealing", "NOTE", err}
			}
			s.Notes = append(s.Notes, t)
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
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

func (s *MultimediaLinkID) parse(l Line) error {
	if err := s.ID.parse(Line{value: l.xrefID}); err != nil {
		return ErrContext{"MultimediaLinkID", "xrefID", err}
	}
	for _, sl := range l.Sub {
		if len(sl.tag) < 1 || sl.tag[0] != '_' {
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

func (s *MultimediaLinkFile) parse(l Line) error {
	var FormatSet, TitleSet, FileSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "FORM":
			if FormatSet {
				return ErrContext{"MultimediaLinkFile", "FORM", ErrSingleMultiple}
			}
			FormatSet = true
			if err := s.Format.parse(sl); err != nil {
				return ErrContext{"MultimediaLinkFile", "FORM", err}
			}
		case "TITLE":
			if TitleSet {
				return ErrContext{"MultimediaLinkFile", "TITLE", ErrSingleMultiple}
			}
			TitleSet = true
			if err := s.Title.parse(sl); err != nil {
				return ErrContext{"MultimediaLinkFile", "TITLE", err}
			}
		case "FILE":
			if FileSet {
				return ErrContext{"MultimediaLinkFile", "FILE", ErrSingleMultiple}
			}
			FileSet = true
			if err := s.File.parse(sl); err != nil {
				return ErrContext{"MultimediaLinkFile", "FILE", err}
			}
		case "NOTE":
			var t NoteStructure
			if err != t.parse(sl); err != nil {
				return ErrContext{"MultimediaLinkFile", "NOTE", err}
			}
			s.Notes = append(s.Notes, t)
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
				return ErrContext{"MultimediaLinkFile", sl.tag, ErrUnknownTag}
			}
			// possibly store in a Other field
		}
	}
	if !FormatSet {
		return ErrContext{"MultimediaLinkFile", "Format", ErrRequiredMissing}
	}
	return nil
}

// NoteID is a GEDCOM structure type
type NoteID struct {
	ID Xref
}

func (s *NoteID) parse(l Line) error {
	if err := s.ID.parse(Line{value: l.xrefID}); err != nil {
		return ErrContext{"NoteID", "xrefID", err}
	}
	for _, sl := range l.Sub {
		if len(sl.tag) < 1 || sl.tag[0] != '_' {
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

func (s *NoteText) parse(l Line) error {
	if err := s.SubmitterText.parse(l); err != nil {
		return ErrContext{"NoteText", "line_value", err}
	}
	for _, sl := range l.Sub {
		switch sl.tag {
		case "SOUR":
			var t SourceCitation
			if err != t.parse(sl); err != nil {
				return ErrContext{"NoteText", "SOUR", err}
			}
			s.Sources = append(s.Sources, t)
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
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

func (s *PersonalNameStructure) parse(l Line) error {
	if err := s.NamePersonal.parse(l); err != nil {
		return ErrContext{"PersonalNameStructure", "line_value", err}
	}
	var PrefixSet, GivenSet, NicknameSet, SurnamePrefixSet, SurnameSet, SuffixSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "NPFX":
			if PrefixSet {
				return ErrContext{"PersonalNameStructure", "NPFX", ErrSingleMultiple}
			}
			PrefixSet = true
			if err := s.Prefix.parse(sl); err != nil {
				return ErrContext{"PersonalNameStructure", "NPFX", err}
			}
		case "GIVN":
			if GivenSet {
				return ErrContext{"PersonalNameStructure", "GIVN", ErrSingleMultiple}
			}
			GivenSet = true
			if err := s.Given.parse(sl); err != nil {
				return ErrContext{"PersonalNameStructure", "GIVN", err}
			}
		case "NICK":
			if NicknameSet {
				return ErrContext{"PersonalNameStructure", "NICK", ErrSingleMultiple}
			}
			NicknameSet = true
			if err := s.Nickname.parse(sl); err != nil {
				return ErrContext{"PersonalNameStructure", "NICK", err}
			}
		case "SPFX":
			if SurnamePrefixSet {
				return ErrContext{"PersonalNameStructure", "SPFX", ErrSingleMultiple}
			}
			SurnamePrefixSet = true
			if err := s.SurnamePrefix.parse(sl); err != nil {
				return ErrContext{"PersonalNameStructure", "SPFX", err}
			}
		case "SURN":
			if SurnameSet {
				return ErrContext{"PersonalNameStructure", "SURN", ErrSingleMultiple}
			}
			SurnameSet = true
			if err := s.Surname.parse(sl); err != nil {
				return ErrContext{"PersonalNameStructure", "SURN", err}
			}
		case "NSFX":
			if SuffixSet {
				return ErrContext{"PersonalNameStructure", "NSFX", ErrSingleMultiple}
			}
			SuffixSet = true
			if err := s.Suffix.parse(sl); err != nil {
				return ErrContext{"PersonalNameStructure", "NSFX", err}
			}
		case "SOUR":
			var t SourceCitation
			if err != t.parse(sl); err != nil {
				return ErrContext{"PersonalNameStructure", "SOUR", err}
			}
			s.Sources = append(s.Sources, t)
		case "NOTE":
			var t NoteStructure
			if err != t.parse(sl); err != nil {
				return ErrContext{"PersonalNameStructure", "NOTE", err}
			}
			s.Notes = append(s.Notes, t)
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
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

func (s *PlaceStructure) parse(l Line) error {
	if err := s.Place.parse(l); err != nil {
		return ErrContext{"PlaceStructure", "line_value", err}
	}
	var PlaceHierarchySet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "FORM":
			if PlaceHierarchySet {
				return ErrContext{"PlaceStructure", "FORM", ErrSingleMultiple}
			}
			PlaceHierarchySet = true
			if err := s.PlaceHierarchy.parse(sl); err != nil {
				return ErrContext{"PlaceStructure", "FORM", err}
			}
		case "SOUR":
			var t SourceCitation
			if err != t.parse(sl); err != nil {
				return ErrContext{"PlaceStructure", "SOUR", err}
			}
			s.Sources = append(s.Sources, t)
		case "NOTE":
			var t NoteStructure
			if err != t.parse(sl); err != nil {
				return ErrContext{"PlaceStructure", "NOTE", err}
			}
			s.Notes = append(s.Notes, t)
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
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

func (s *SourceID) parse(l Line) error {
	if err := s.ID.parse(Line{value: l.xrefID}); err != nil {
		return ErrContext{"SourceID", "xrefID", err}
	}
	var WhereWithinSourceSet, SourceCitationEventSet, DataSet, CertaintySet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "PAGE":
			if WhereWithinSourceSet {
				return ErrContext{"SourceID", "PAGE", ErrSingleMultiple}
			}
			WhereWithinSourceSet = true
			if err := s.WhereWithinSource.parse(sl); err != nil {
				return ErrContext{"SourceID", "PAGE", err}
			}
		case "EVEN":
			if SourceCitationEventSet {
				return ErrContext{"SourceID", "EVEN", ErrSingleMultiple}
			}
			SourceCitationEventSet = true
			if err := s.SourceCitationEvent.parse(sl); err != nil {
				return ErrContext{"SourceID", "EVEN", err}
			}
		case "DATA":
			if DataSet {
				return ErrContext{"SourceID", "DATA", ErrSingleMultiple}
			}
			DataSet = true
			if err := s.Data.parse(sl); err != nil {
				return ErrContext{"SourceID", "DATA", err}
			}
		case "QUAY":
			if CertaintySet {
				return ErrContext{"SourceID", "QUAY", ErrSingleMultiple}
			}
			CertaintySet = true
			if err := s.Certainty.parse(sl); err != nil {
				return ErrContext{"SourceID", "QUAY", err}
			}
		case "OBJE":
			var t MultimediaLink
			if err != t.parse(sl); err != nil {
				return ErrContext{"SourceID", "OBJE", err}
			}
			s.Multimedia = append(s.Multimedia, t)
		case "NOTE":
			var t NoteStructure
			if err != t.parse(sl); err != nil {
				return ErrContext{"SourceID", "NOTE", err}
			}
			s.Notes = append(s.Notes, t)
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
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

func (s *SourceCitationEvent) parse(l Line) error {
	if err := s.EventTypeCitedFrom.parse(l); err != nil {
		return ErrContext{"SourceCitationEvent", "line_value", err}
	}
	var RoleSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "ROLE":
			if RoleSet {
				return ErrContext{"SourceCitationEvent", "ROLE", ErrSingleMultiple}
			}
			RoleSet = true
			if err := s.Role.parse(sl); err != nil {
				return ErrContext{"SourceCitationEvent", "ROLE", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
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

func (s *SourceData) parse(l Line) error {
	var DateSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "DATE":
			if DateSet {
				return ErrContext{"SourceData", "DATE", ErrSingleMultiple}
			}
			DateSet = true
			if err := s.Date.parse(sl); err != nil {
				return ErrContext{"SourceData", "DATE", err}
			}
		case "TEXT":
			var t TextFromSource
			if err != t.parse(sl); err != nil {
				return ErrContext{"SourceData", "TEXT", err}
			}
			s.Text = append(s.Text, t)
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
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

func (s *SourceText) parse(l Line) error {
	if err := s.Description.parse(l); err != nil {
		return ErrContext{"SourceText", "line_value", err}
	}
	for _, sl := range l.Sub {
		switch sl.tag {
		case "TEXT":
			var t TextFromSource
			if err != t.parse(sl); err != nil {
				return ErrContext{"SourceText", "TEXT", err}
			}
			s.Texts = append(s.Texts, t)
		case "NOTE":
			var t NoteStructure
			if err != t.parse(sl); err != nil {
				return ErrContext{"SourceText", "NOTE", err}
			}
			s.Notes = append(s.Notes, t)
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
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

func (s *SourceRepositoryCitation) parse(l Line) error {
	if err := s.ID.parse(Line{value: l.xrefID}); err != nil {
		return ErrContext{"SourceRepositoryCitation", "xrefID", err}
	}
	for _, sl := range l.Sub {
		switch sl.tag {
		case "CALN":
			var t SourceCallStructure
			if err != t.parse(sl); err != nil {
				return ErrContext{"SourceRepositoryCitation", "CALN", err}
			}
			s.Numbers = append(s.Numbers, t)
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
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

func (s *SourceCallStructure) parse(l Line) error {
	if err := s.SourceCallNumber.parse(l); err != nil {
		return ErrContext{"SourceCallStructure", "line_value", err}
	}
	var TypeSet bool
	for _, sl := range l.Sub {
		switch sl.tag {
		case "MEDI":
			if TypeSet {
				return ErrContext{"SourceCallStructure", "MEDI", ErrSingleMultiple}
			}
			TypeSet = true
			if err := s.Type.parse(sl); err != nil {
				return ErrContext{"SourceCallStructure", "MEDI", err}
			}
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
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

func (s *SpouseToFamilyLink) parse(l Line) error {
	if err := s.ID.parse(Line{value: l.xrefID}); err != nil {
		return ErrContext{"SpouseToFamilyLink", "xrefID", err}
	}
	for _, sl := range l.Sub {
		switch sl.tag {
		case "NOTE":
			var t NoteStructure
			if err != t.parse(sl); err != nil {
				return ErrContext{"SpouseToFamilyLink", "NOTE", err}
			}
			s.Notes = append(s.Notes, t)
		default:
			if len(sl.tag) < 1 || sl.tag[0] != '_' {
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

func (s *MultimediaLink) parse(l Line) error {
	if l.xrefID != "" {
		s.Data = MultimediaLinkID{}
	} else {
		s.Data = MultimediaLinkFile{}
	}
	return s.Data.parse(l)
}

// NoteStructure splits between NoteID and NoteText
type NoteStructure struct {
	Data Record
}

func (s *NoteStructure) parse(l Line) error {
	if l.xrefID != "" {
		s.Data = NoteID{}
	} else {
		s.Data = NoteText{}
	}
	return s.Data.parse(l)
}

// SourceCitation splits between SourceID and SourceText
type SourceCitation struct {
	Data Record
}

func (s *SourceCitation) parse(l Line) error {
	if l.xrefID != "" {
		s.Data = SourceID{}
	} else {
		s.Data = SourceText{}
	}
	return s.Data.parse(l)
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
	return e.Tag + ":" + t.Err.Error()
}

// ErrTooMany is an error returned when too many of a particular tag exist
type ErrTooMany int

// Error implements the error interface
func (ErrTooMany) Error() string {
	return "too many tags"
}
