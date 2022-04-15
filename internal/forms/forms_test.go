package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when we should have been valid")
	}
}

func TestNew(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)
	if form == nil {
		t.Error("function deosnt return form")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "b")
	postedData.Add("b", "k")
	postedData.Add("c", "l")

	r = httptest.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields")
	}

}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.MinLength("x", 10)
	if form.Valid() {
		t.Error("Form shows min length for non-existent field")
	}

	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("should have an error ,but did not get one ")
	}

	postedValues := url.Values{}
	postedValues.Add("some_field", "some_values")
	form = New(postedValues)

	form.MinLength("some_field", 100)
	if form.Valid() {
		t.Error("shows minlength when field is shorter than 100 !")
	}

	postedValues = url.Values{}
	postedValues.Add("a", "abc123")
	form = New(postedValues)

	form.MinLength("a", 1)
	if !form.Valid() {
		t.Error("Form is long to be expected but it fails")
	}
	isError = form.Errors.Get("a")
	if isError != "" {
		t.Error("should not have an error ,but  get one ")
	}
}

func TestForm_Isemail(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)
	form.Isemail("x")
	if form.Valid() {
		t.Error("form shows valid email for non-existent field")
	}

	postedValues := url.Values{}
	postedValues.Add("email", "xyexy@gmail.com")
	form = New(postedValues)

	form.Isemail("email")

	if !form.Valid() {
		t.Error(" we got invalid email when we should not have ")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "xyexygmail.com")
	form = New(postedValues)

	form.Isemail("email")

	if form.Valid() {
		t.Error("got valid email for invalid ")
	}

}

func TestForm_Has(t *testing.T) {

	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)
	has := form.Has("whatever")
	if has {
		t.Error("form shows has field when it doesnt")
	}

	postedData := url.Values{}
	postedData.Add("a", "b")
	form = New(postedData)
	has = form.Has("a")
	if !has {
		t.Error("shows form has field when it doesnt ")
	}
}
