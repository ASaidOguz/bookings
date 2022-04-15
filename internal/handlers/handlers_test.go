package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var TheTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{

	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"gq", "/generals-quarters", "GET", []postData{}, http.StatusOK},
	{"ms", "/majors-suite", "GET", []postData{}, http.StatusOK},
	{"sa", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"con", "/contact", "GET", []postData{}, http.StatusOK},
	{"mr", "/make-reservation", "GET", []postData{}, http.StatusOK},
	{"post-search-avail", "/search-availability", "POST", []postData{
		{key: "2022-10-10"},
		{value: "2022-10-15"},
	}, http.StatusOK},
	{"post-search-avail-json", "/search-availability-json", "POST", []postData{
		{key: "2022-10-10"},
		{value: "2022-10-15"},
	}, http.StatusOK},
	{"make-reserve", "/make-reservation", "POST", []postData{
		{key: "first_name"},
		{value: "John"},
		{key: "last_name"},
		{value: "Smith"},
		{key: "email"},
		{value: "aso33@gmail.com"},
		{key: "phone"},
		{value: "212-580-46-45"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {

	routes := getRoots()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range TheTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s ,expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}

		} else {
			values := url.Values{}

			for _, x := range e.params {
				values.Add(x.key, x.value)
			}
			resp, err := ts.Client().PostForm(ts.URL+e.url, values)

			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s ,expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}

		}
	}

}
