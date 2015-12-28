package gedcom

// RecordID stores the type of record
type RecordID uint

// RecordIDs
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

// Record is the interface to contain all of the record types
type Record interface {
	Type() RecordID
	parse(*Line, options) error
}

// Type implements the Record interface
func (Header) Type() RecordID {
	return RecordHeader
}

// Type implements the Record interface
func (SubmissionRecord) Type() RecordID {
	return RecordSubmission
}

// Type implements the Record interface
func (Family) Type() RecordID {
	return RecordFamily
}

// Type implements the Record interface
func (Individual) Type() RecordID {
	return RecordIndividual
}

// Type implements the Record interface
func (MultimediaRecord) Type() RecordID {
	return RecordMultimedia
}

// Type implements the Record interface
func (NoteRecord) Type() RecordID {
	return RecordNote
}

// Type implements the Record interface
func (RepositoryRecord) Type() RecordID {
	return RecordRepository
}

// Type implements the Record interface
func (SourceRecord) Type() RecordID {
	return RecordSource
}

// Type implements the Record interface
func (SubmitterRecord) Type() RecordID {
	return RecordSubmitter
}

// Type implements the Record interface
func (Trailer) Type() RecordID {
	return RecordTrailer
}

// Type implements the Record interface
func (MultimediaLinkID) Type() RecordID {
	return RecordMultimediaLinkID
}

// Type implements the Record interface
func (MultimediaLinkFile) Type() RecordID {
	return RecordMultimediaLinkFile
}

// Type implements the Record interface
func (NoteID) Type() RecordID {
	return RecordNoteID
}

// Type implements the Record interface
func (NoteText) Type() RecordID {
	return RecordNoteText
}

// Type implements the Record interface
func (SourceID) Type() RecordID {
	return RecordSourceID
}

// Type implements the Record interface
func (SourceText) Type() RecordID {
	return RecordSourceText
}

// Type implements the Record interface
func (Line) Type() RecordID {
	return RecordUnknown
}

func (Line) parse(*Line, options) error {
	return nil
}
