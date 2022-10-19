package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestFlag2Handler testes getFlag2()
func TestFlag2Handler(t *testing.T) {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/flag2", nil)
	router.ServeHTTP(w, req)

	// test the status code
	testStatusCode(t, w.Code, http.StatusOK)

	// test the header, footer, and nav renders
	testCommonTemplatesRender(t, w.Body.String())

	// test the body of the index
	testBodyContains(t, `<h1>Om nom 🍪</h1>`, w.Body.String())

	testFlag2Success(t, w.Result().Cookies())
}

// testFlag2Success does what's necessary for flag2 to succeed
// and confirms that the received flag is as expected.
func testFlag2Success(t *testing.T, c []*http.Cookie) {
	id := 2
	cookieName := "flag2"
	f, err := getFlag(id)
	if err != nil {
		t.Error("The flag at the provided ID was not found: ", id)
	}

	// the cookie when accessed via recorder is flag2=val
	// so we need to split to test the value
	fkv := strings.Split(c[0].String(), ";")[0]
	cookieVal := strings.Split(fkv, "=")[1]

	if f.Flag.String() != cookieVal {
		t.Errorf("Unable to find flag in expected cookie %s.\n Got: %s\nWant: %s",
			cookieName, cookieVal, f.Flag)
	}
}
