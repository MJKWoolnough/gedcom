package gedcom

type options struct {
	allowUnknownTags        bool
	allowWrongLength        bool
	allowMissingRequired    bool
	allowMoreThanAllowed    bool
	ignoreInvalidValue      bool
	allowUnknownCharset     bool
	allowTerminatorsInValue bool
	allowInvalidEscape      bool
	allowInvalidChars       bool
}

// Option type for turning off generated errors.
type Option func(o *options)

// AllowUnknownTags turns off error reporting when an unknown tag is read.
func AllowUnknownTags(o *options) {
	o.allowUnknownTags = true
}

// AllowWrongLength turns off error reporting when a base type has an incorrect
// length.
func AllowWrongLength(o *options) {
	o.allowWrongLength = true
}

// AllowMissingRequired turns off error reporting for a missing required field.
func AllowMissingRequired(o *options) {
	o.allowMissingRequired = true
}

// AllowMoreThanAllowed turns off error reporting when more than the maxumum
// number of a certain item are read.
func AllowMoreThanAllowed(o *options) {
	o.allowMoreThanAllowed = true
}

// IgnoreInvalidValue turns off error reporting when a non-valid value if read.
func IgnoreInvalidValue(o *options) {
	o.ignoreInvalidValue = true
}

// AllowUnknownCharset allows the parser to read any character, not just those
// within ANSEL.
func AllowUnknownCharset(o *options) {
	o.allowUnknownCharset = true
}

// AllowTerminatorsInValue allows a line_value to contain newlines.
func AllowTerminatorsInValue(o *options) {
	o.allowTerminatorsInValue = true
}

// AllowInvalidEscape turns off error reporting for invalid escape sequences.
func AllowInvalidEscape(o *options) {
	o.allowInvalidEscape = true
}

// AllowInvalidChars allows control characters (<32) in line_values.
func AllowInvalidChars(o *options) {
	o.allowInvalidChars = true
}
