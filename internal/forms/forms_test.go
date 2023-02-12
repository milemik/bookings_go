package forms

import (
	"net/http/httptest"
	"net/url"
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

var hasValueTestData = []struct {
	fieldName      string
	fieldValue     string
	expectedResult bool
}{
	{"a", "", false},
	{"a", "aaa", true},
}

func TestHasValueIsHere(t *testing.T) {
	for _, data := range hasValueTestData {
		tr := httptest.NewRequest("POST", "/something", nil)

		postedData := url.Values{}
		postedData.Add(data.fieldName, data.fieldValue)

		tr.PostForm = postedData
		from := New(tr.PostForm)
		tr.ParseForm()

		result := from.Has(data.fieldName, tr)

		if result != data.expectedResult {
			t.Errorf("Expected true got %v", result)
		}
	}
}

var minLengthTestData = []struct {
	fieldName      string
	fieldValue     string
	expectedResult bool
}{
	{"a", "", false},
	{"a", "aaa", false},
	{"a", "aaaaa", true},
	{"a", "aaaaaaaa", true},
}

func TestMinLengthBad(t *testing.T) {
	for _, data := range minLengthTestData {
		tr := httptest.NewRequest("POST", "/something", nil)

		postedData := url.Values{}
		postedData.Add(data.fieldName, data.fieldValue)

		tr.PostForm = postedData
		from := New(tr.PostForm)
		tr.ParseForm()

		result := from.MinLength(data.fieldName, 5, tr)

		if result != data.expectedResult {
			t.Errorf("Expected true got %v", result)
		}
	}
}

var badEmailTestData = []struct {
	field string
	value string
}{
	{"a", "aaa"},
	{"email", "test@test"},
}

func TestIsEmailBad(t *testing.T) {
	for _, data := range badEmailTestData {
		tr := httptest.NewRequest("POST", "/something", nil)

		postedData := url.Values{}
		postedData.Add(data.field, data.value)

		tr.PostForm = postedData
		from := New(tr.PostForm)
		tr.ParseForm()

		from.IsEmail(data.field)

		result := from.Errors.Get(data.field)

		if result != "Invalid email address" {
			t.Error("Expect to see arror not a valid email adderss")
		}
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
