package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestFlag7Handler testes getFlag7()
func TestFlag7Handler(t *testing.T) {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/flag7", nil)
	router.ServeHTTP(w, req)

	// test the status code
	testStatusCode(t, w.Code, http.StatusBadRequest)

	// test the header, footer, and nav renders
	testCommonTemplatesRender(t, w.Body.String())

	// test the body of the index
	testBodyContains(t, `flag</code> and ask it nicely`, w.Body.String())

	q := req.URL.Query()
	q.Add("flag", "flag7")
	q.Add("please", "true")
	req.URL.RawQuery = q.Encode()
	router.ServeHTTP(w, req)
	testFlag7Success(t, w.Body.String())
}

// testFlag7Success does what's necessary for flag7 to succeed
// and confirms that the received flag is as expected.
func testFlag7Success(t *testing.T, body string) {
	id := 7
	f, err := getFlag(id)
	if err != nil {
		t.Error("The flag at the provided ID was not found: ", id)
	}

	testBodyContains(t, f.Flag.String(), body)
}
