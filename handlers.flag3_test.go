package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestFlag1Handler testes getFlag1()
func TestFlag3Handler(t *testing.T) {
	htmlw := httptest.NewRecorder()
	jsonw := httptest.NewRecorder()

	htmlreq, _ := http.NewRequest("GET", "/flag3", nil)
	router.ServeHTTP(htmlw, htmlreq)

	jsonreq, _ := http.NewRequest("GET", "/flag3", nil)
	jsonreq.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(jsonw, jsonreq)

	// test the status code
	testStatusCode(t, htmlw.Code, http.StatusOK)

	// test the header, footer, and nav renders
	testCommonTemplatesRender(t, htmlw.Body.String())

	// test the body of the index
	testBodyContains(t, `<a href="https://json.org">This is not JSON</a>`, htmlw.Body.String())

	testFlag3Success(t, jsonw.Body.String())
}

// testFlag3Success does what's necessary for flag3 to succeed
// and confirms that the received flag is as expected.
func testFlag3Success(t *testing.T, jsonbody string) {
	id := 3
	f, err := getFlag(id)
	if err != nil {
		t.Error("The flag at the provided ID was not found: ", id)
	}

	// the json response from the serve is quoted, we need to remove that to test.
	jsonflag := strings.Trim(jsonbody, `"`)

	if f.Flag.String() != jsonflag {
		t.Errorf("Unable to find flag in expected JSON payload.\n Got: %s\nWant %s",
			jsonflag, f.Flag)
	}
}
