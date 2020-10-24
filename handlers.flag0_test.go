package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestFlag0Handler testes getFlag0()
func TestFlag0Handler(t *testing.T) {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/flag0", nil)
	router.ServeHTTP(w, req)

	// test the status code
	testStatusCode(t, w.Code, http.StatusOK)

	// test the header, footer, and nav renders
	testCommonTemplatesRender(t, w.Body.String())

	// test the body of the index
	testBodyContains(t, `<h1>The flag is in the source...</h1>`, w.Body.String())

	testFlag0Success(t, w.Body.String())
}

// testFlag0Success does what's necessary for flag0 to succeed
// and confirms that the received flag is as expected.
func testFlag0Success(t *testing.T, body string) {
	id := 0
	f, err := getFlag(id)
	if err != nil {
		t.Error("The flag at the provided ID was not found: ", id)
	}

	testBodyContains(t, f.Flag, body)
}
