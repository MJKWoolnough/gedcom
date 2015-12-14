package gedcom

type RecordType uint8

const (
	RecordHeader RecordType = iota
	RecordSubmitter
	RecordFamily
	RecordIndividual
	RecordMultimedia
	RecordNote
	RecordRepository
	RecordSource
	RecordSubmission
	RecordTrailer
	RecordUnknown
)

type Record interface {
	Type() RecordType
}
