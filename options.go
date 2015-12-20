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

type Option func(o *options)

func AllowUnknownTags(o *options) {
	o.allowUnknownTags = true
}

func AllowWrongLength(o *options) {
	o.allowWrongLength = true
}

func AllowMissingRequired(o *options) {
	o.allowMissingRequired = true
}

func AllowMoreThanAllowed(o *options) {
	o.allowMoreThanAllowed = true
}

func IgnoreInvalidValue(o *options) {
	o.ignoreInvalidValue = true
}

func AllowUnknownCharset(o *options) {
	o.allowUnknownCharset = true
}

func AllowTerminatorsInValue(o *options) {
	o.allowTerminatorsInValue = true
}

func AllowInvalidEscape(o *options) {
	o.allowInvalidEscape = true
}

func AllowInvalidChars(o *options) {
	o.allowInvalidChars = true
}
