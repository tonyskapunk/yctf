package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestFlag4Handler testes getFlag4()
func TestFlag4Handler(t *testing.T) {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/flag4", nil)
	router.ServeHTTP(w, req)

	// test the status code
	testStatusCode(t, w.Code, http.StatusOK)

	// test the header, footer, and nav renders
	testCommonTemplatesRender(t, w.Body.String())

	// test the body of the index
	testBodyContains(t, `<p>You don't look like a bot`, w.Body.String())

	req.Header.Set("User-Agent", "Wall-E Scraperbot")
	router.ServeHTTP(w, req)

	testFlag4Success(t, w.Body.String())
}

// testFlag4Success does what's necessary for flag4 to succeed
// and confirms that the received flag is as expected.
func testFlag4Success(t *testing.T, body string) {
	id := 4
	f, err := getFlag(id)
	if err != nil {
		t.Error("The flag at the provided ID was not found: ", id)
	}

	testBodyContains(t, f.Flag, body)
}
