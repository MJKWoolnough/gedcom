package gedcom

import "errors"

type Generator struct {
	SystemID    string
	Version     string
	ProductName string
	Corp
	SourceData
}

type Corp struct {
	BusinessName    string
	BusinessAddress Address
}

type SourceData struct {
	SourceName string
	Date       string
	Copyright  string
}

type TransmissionDate struct {
	Date string
	Time string
}

type Gedcom struct {
	Version string
	Form    string
}

type Character struct {
	Set     string
	Version string
}

type Header struct {
	Generator
	Dest             string
	TransmissionDate TransmissionDate
	Submitter        string
	Submission       string
	File             string
	Copyright        string
	Gedcom           Gedcom
	Character        Character
	Language         string
	Place            string
	Note             string
	Custom           []Line
}

func parseHeader(l Line) (h Header, err error) {
	defer func() {
		if e := recover(); err != nil {
			if etm, ok := e.(*ErrTooMany); ok {
				err = etm
			}
		}
		if err != nil {
			h = Header{}
		}
	}()
	var sour, dest, date, subm, subn, file, copr, gedc, char, lang, plac, note bool
	for _, s := range l.Sub {
		switch s.tag {
		case "SOUR":
			setCheck(&sour, "SOUR")
		case "DEST":
			setCheck(&dest, "DEST")
		case "DATE":
			setCheck(&date, "DATE")
		case "SUBM":
			setCheck(&subm, "SUBM")
		case "SUBN":
			setCheck(&subn, "SUBN")
		case "FILE":
			setCheck(&file, "FILE")
		case "COPR":
			setCheck(&copr, "COPR")
		case "GEDC":
			setCheck(&gedc, "GEDC")
		case "CHAR":
			setCheck(&char, "CHAR")
		case "LANG":
			setCheck(&lang, "LANG")
		case "PLAC":
			setCheck(&plac, "PLAC")
		case "NOTE":
			setCheck(&note, "NOTE")
		default:
			if len(s.tag) < 1 || s.tag[0] != '_' {
				return Header{}, ErrUnknownTag
			}
			h.Custom = append(h.Misc, s)
		}
	}
	if !sour {
		return Header{}, ErrNoApprovedSystemID
	}
	if !subm {
		return Header{}, ErrNoSubmitter
	}
	if !char {
		return Header{}, ErrNoCharacterSet
	}
	return h, nil
}

func (Header) Type() RecordType {
	return RecordHeader
}

// Errors
var (
	ErrNoApprovedSystemID = errors.New("no required approved system id")
	ErrNoSubmitter        = errors.New("no required submitter")
	ErrNoCharacterSet     = errors.New("no required character set")
)
