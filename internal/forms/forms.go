package forms

import (
	"net/http"
	"net/url"
)

// Creates a custom form struct
type Form struct {
	url.Values
	Errors errors
}

// Initialize a fomr struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Checks if Field is not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		f.Errors.Add(field, "This field cannot be blank")
		return false
	}
	return true
}

// valid returns true if there are no errors
func (f *Form) Valid() bool {

	return len(f.Errors) == 0

}
