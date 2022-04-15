package forms

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

type Form struct {
	url.Values
	Errors errors
}

// Valid returns true if there r no errors ,otherwise false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

//New initialize a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}

}

//Required check for required fields.
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")

		}
	}
}

//MinLength checks strings minimum length
func (f *Form) MinLength(field string, length int) bool {
	x := f.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("This field has to at least %d character long", length))
		return false
	}
	return true
}

//Isemail checks the field if its valid email form
func (f *Form) Isemail(field string) bool {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
		return false
	}
	return true
}

//Has check if form field is in post and not empty useful function for check box methods!!!!
func (f *Form) Has(field string) bool {

	x := f.Get(field)
	if x == "" {

		return false
	}
	return true
}
