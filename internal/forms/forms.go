package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Form creates a form struct, embades an url.Values object
type Form struct {
	url.Values
	Errors errors
}

// Valid returns true if there are no errors - form valid
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// New initialize a form strunct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Required check for required fields
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field can't be blank")
		}
	}
}

// Has checks if form field is in post or empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		// f.Errors.Add(field, "This field can't be blank") - don't need this any more
		return false
	}
	return true
}

// MinLength check for sting minimum length
func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	x := r.Form.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d chars long", length))
		return false
	}
	return true
}
