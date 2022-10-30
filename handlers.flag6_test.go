package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestFlag6Handler testes getFlag6()
func TestFlag6Handler(t *testing.T) {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/flag6", nil)
	router.ServeHTTP(w, req)

	// test the status code
	testStatusCode(t, w.Code, http.StatusOK)

	// test the header, footer, and nav renders
	testCommonTemplatesRender(t, w.Body.String())

	// test the body of the index
	testBodyContains(t, `only IPs from '11.0.0.0/8' are allowed</p>`, w.Body.String())

	req.Header.Set("x-forwarded-for", "11.0.0.1")
	router.ServeHTTP(w, req)
	testFlag6Success(t, w.Body.String())
}

// testFlag6Success does what's necessary for flag0 to succeed
// and confirms that the received flag is as expected.
func testFlag6Success(t *testing.T, body string) {
	id := 6
	f, err := getFlag(id)
	if err != nil {
		t.Error("The flag at the provided ID was not found: ", id)
	}

	testBodyContains(t, f.Flag, body)
}
