package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
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
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does.")
	}

}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	has := form.Has("whatever")
	if has {
		t.Error("form shows has field when it does not.")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)

	has = form.Has("a")
	if !has {
		t.Error("shows form does not have field when it should")
	}
}

func TestForm_MinLength(t *testing.T) {

	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	postedData := url.Values{}
	postedData.Add("a", "steven")
	form = New(postedData)

	result := form.MinLength("a", 3)

	if !result {
		t.Error("Min length not met")
	}

}

func TestForm_IsEmail(t *testing.T) {
	postedValues := url.Values{}
	form := New(postedValues)

	form.IsEmail("x")
	if form.Valid() {
		t.Error("forms show valid email for non-existent field")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "me@here.com")
	form = New(postedValues)

	form.IsEmail("email")
	if !form.Valid() {
		t.Error("got an invalid email when whe should not have")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "x")
	form = New(postedValues)

	form.IsEmail("email")
	if form.Valid() {
		t.Error("got valid for invalid email address.")
	}
}
