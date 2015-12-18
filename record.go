package gedcom

type RecordID uint

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

type Record interface {
	Type() RecordID
	parse(Line) error
}

func (Header) Type() RecordID {
	return RecordHeader
}

func (SubmissionRecord) Type() RecordID {
	return RecordSubmission
}

func (Family) Type() RecordID {
	return RecordFamily
}

func (Individual) Type() RecordID {
	return RecordIndividual
}

func (MultimediaRecord) Type() RecordID {
	return RecordMultimedia
}

func (NoteRecord) Type() RecordID {
	return RecordNote
}

func (RepositoryRecord) Type() RecordID {
	return RecordRepository
}

func (SourceRecord) Type() RecordID {
	return RecordSource
}

func (SubmitterRecord) Type() RecordID {
	return RecordSubmitter
}

func (Trailer) Type() RecordID {
	return RecordTrailer
}

func (MultimediaLinkID) Type() RecordID {
	return RecordMultimediaLinkID
}

func (MultimediaLinkFile) Type() RecordID {
	return RecordMultimediaLinkFile
}

func (NoteID) Type() RecordID {
	return RecordNoteID
}

func (NoteText) Type() RecordID {
	return RecordNoteText
}

func (SourceID) Type() RecordID {
	return RecordSourceID
}

func (SourceText) Type() RecordID {
	return RecordSourceText
}

func (Line) Type() RecordID {
	return RecordUnknown
}

func (Line) parse(Line) error {
	return nil
}
