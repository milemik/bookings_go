package forms

import (
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestValid(t *testing.T) {
	f := Form{url.Values{}, errors{}}

	if !f.Valid() {
		t.Error("Form has no errors and should be valid")
	}

	f.Errors.Add("err", "errawv")

	if f.Valid() {
		t.Error("Form has errors and should not be valid")
	}
}

func TestNew(t *testing.T) {
	testData := url.Values{}

	_ = New(testData)

}

func TestNewWithValues(t *testing.T) {
	const key = "key"
	const val = "val"
	testData := url.Values{}

	testData.Add(key, val)

	f := New(testData)

	result := f.Values.Get(key)
	if result != val {
		t.Errorf("Exptexted %s and got %s", val, result)
	}
}

func TestRequired(t *testing.T) {
	f := Form{url.Values{}, errors{}}

	fields := []string{"a", "b", "c"}

	f.Required(fields...)

	if len(f.Errors) < 3 {
		t.Error("3 erros did not add")
	}
}

func TestRequredWithFields(t *testing.T) {
	f := Form{url.Values{}, errors{}}

	f.Values.Add("a", "aa")
	f.Values.Add("b", "bb")
	f.Values.Add("c", "cc")

	fields := []string{"a", "b", "c"}

	f.Required(fields...)

	if len(f.Errors) > 0 {
		t.Error("Have errors here and expected none")
	}
}

func TestRequredWithFieldsMissingOne(t *testing.T) {
	f := Form{url.Values{}, errors{}}

	f.Values.Add("a", "aa")
	f.Values.Add("c", "cc")

	fields := []string{"a", "b", "c"}

	f.Required(fields...)

	if len(f.Errors) != 1 {
		t.Error("Expected one error and got different num of errors")
	}
}

func TestHasEmpty(t *testing.T) {
	testRequest := httptest.NewRequest("GET", "/something", &strings.Reader{})
	testForm := New(testRequest.PostForm)

	result := testForm.Has("a", testRequest)

	if result != false {
		t.Error("Expected false got true")
	}
}

func TestHasValueIsHere(t *testing.T) {
	tr := httptest.NewRequest("POST", "/something", nil)

	postedData := url.Values{}
	postedData.Add("a", "aaa")

	tr.PostForm = postedData
	from := New(tr.PostForm)
	tr.ParseForm()

	result := from.Has("a", tr)

	if result != true {
		t.Errorf("Expected true got %v", result)
	}
}

func TestMinLengthBad(t *testing.T) {
	tr := httptest.NewRequest("POST", "/something", nil)

	postedData := url.Values{}
	postedData.Add("a", "aaa")

	tr.PostForm = postedData
	from := New(tr.PostForm)
	tr.ParseForm()

	result := from.MinLength("a", 5, tr)

	if result != false {
		t.Errorf("Expected true got %v", result)
	}
}

func TestMinLengthOk(t *testing.T) {
	tr := httptest.NewRequest("POST", "/something", nil)

	postedData := url.Values{}
	postedData.Add("a", "aaaaa")

	tr.PostForm = postedData
	from := New(tr.PostForm)
	tr.ParseForm()

	result := from.MinLength("a", 5, tr)

	if result != true {
		t.Errorf("Expected true got %v", result)
	}
}

func TestIsEmailBad(t *testing.T) {
	tr := httptest.NewRequest("POST", "/something", nil)

	postedData := url.Values{}
	postedData.Add("a", "aaaaa")

	tr.PostForm = postedData
	from := New(tr.PostForm)
	tr.ParseForm()

	from.IsEmail("a")

	result := from.Errors.Get("a")

	if result != "Invalid email address" {
		t.Error("Expect to see arror not a valid email adderss")
	}
}


func TestIsEmailOk(t *testing.T) {
	tr := httptest.NewRequest("POST", "/something", nil)

	postedData := url.Values{}
	postedData.Add("a", "aaaaa@aaa.com")

	tr.PostForm = postedData
	from := New(tr.PostForm)
	tr.ParseForm()

	from.IsEmail("a")

	result := from.Errors

	if len(result) != 0 {
		t.Errorf("Expect to see 0 errors and got %d", len(result))
	}
}

