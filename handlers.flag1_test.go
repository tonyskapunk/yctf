package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestFlag1Handler testes getFlag1()
func TestFlag1Handler(t *testing.T) {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/flag1", nil)
	router.ServeHTTP(w, req)

	// test the status code
	testStatusCode(t, w.Code, http.StatusOK)

	// test the header, footer, and nav renders
	testCommonTemplatesRender(t, w.Body.String())

	// test the body of the index
	testBodyContains(t, `<h1>The headers</h1`, w.Body.String())

	testFlag1Success(t, w.Header())
}

// testFlag1Success does what's necessary for flag1 to succeed
// and confirms that the received flag is as expected.
func testFlag1Success(t *testing.T, headers http.Header) {
	id := 1
	header := "X-Yctf"
	f, err := getFlag(id)
	if err != nil {
		t.Error("The flag at the provided ID was not found: ", id)
	}

	if f.Flag != headers[header][0] {
		t.Errorf("Unable to find flag in expected response header %s.\n Got: %s\nWant %s",
			header, headers[header][0], f.Flag)
	}

}
