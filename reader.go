package gedcom

import (
	"errors"
	"io"
	"strconv"
)

type line struct {
	level  uint64
	xrefID string
	tag    string
	value  string
}

type Reader struct {
	t       *tokeniser
	options options

	peeked bool
	line   line
	err    error

	hadHeader, hadRecord bool
}

func NewReader(r io.Reader, opts ...Option) *Reader {
	var o options
	for _, opt := range opts {
		opt(&o)
	}
	return &Reader{
		t:       newTokeniser(r, o),
		options: o,
	}
}

func (r *Reader) readLine() {
	if r.err != nil {
		return
	}
	var t token
	t, r.err = r.t.GetToken()
	if r.err != nil {
		return
	}
	if t.typ != tokenLevel {
		r.err = ErrNotLevel
		return
	}
	r.line.level, r.err = strconv.ParseUint(t.data, 10, 64)
	if r.err != nil {
		return
	}
	t, r.err = r.t.GetToken()
	if r.err != nil {
		return
	}
	if t.typ == tokenXref {
		r.line.xrefID = t.data
		t, r.err = r.t.GetToken()
		if r.err != nil {
			return
		}
	} else {
		r.line.xrefID = ""
	}
	if t.typ != tokenTag {
		r.err = ErrNotTag
		return
	}
	r.line.tag = t.data
	t, r.err = r.t.GetToken()
	if r.err != nil {
		return
	}
	if t.typ == tokenEndLine {
		r.line.value = ""
		return
	}
	if t.typ != tokenLine && t.typ != tokenPointer {
		r.err = ErrNotLine
		return
	}
	r.line.value = t.data
}

func (r *Reader) Record() (Record, error) {
	if !r.peeked {
		r.readLine()
		r.peeked = true
	}
	if r.err != nil {
		return nil, r.err
	}
	if !r.hadHeader {
		if r.line.tag != "HEAD" {
			r.peeked = false
			return nil, ErrNoHeader
		}
		r.hadHeader = true
	} else if !r.hadRecord {
		switch r.line.tag {
		case "FAM", "INDI", "OBJE", "NOTE", "REPO", "SOUR", "SUBN":
			r.hadRecord = true
		case "TRLR":
			r.peeked = false
			return nil, ErrNoRecords
		}
	} else if r.line.tag == "TRLR" {
		r.peeked = false
		return &Trailer{}, nil
	}

	lines := make([]line, 1, 32)
	lines[0] = r.line
	var lastlevel uint64
	for {
		if r.err != nil {
			return nil, r.err
		}
		if r.line.level > lastlevel+1 {
			return nil, ErrInvalidLevel
		}
		lastlevel = r.line.level
		lines = append(lines, r.line)
		r.readLine()
		if r.line.level == 0 {
			plines := parseLines(lines)
			var record Record
			switch lines[0].tag {
			case "HEAD":
				record = &Header{}
			case "SUBM":
				record = &SubmissionRecord{}
			case "FAM":
				record = &Family{}
			case "INDI":
				record = &Individual{}
			case "OBJE":
				record = &MultimediaRecord{}
			case "NOTE":
				record = &NoteRecord{}
			case "REPO":
				record = &RepositoryRecord{}
			case "SOUR":
				record = &SourceRecord{}
			case "SUBN":
				record = &SubmitterRecord{}
			default:
				if lines[0].tag[0] == '_' {
					return plines, nil
				}
				return plines, ErrContext{"root", lines[0].tag, ErrUnknownTag}
			}
			err := record.parse(&plines, r.options)
			if err != nil {
				return nil, ErrContext{"root", lines[0].tag, err}
			}
			return record, nil
		}
	}

}

// Errors
var (
	ErrNoHeader  = errors.New("no header")
	ErrNoRecords = errors.New("no records")
	ErrNotLevel  = errors.New("not level token")
	ErrNotTag    = errors.New("not tag token")
	ErrNotLine   = errors.New("not line_value token")
)
