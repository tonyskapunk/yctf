package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func init() {
	router = setupRouter()
}

// TestIndexHandler tests showIndex()
func TestIndexHandler(t *testing.T) {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	// test the status code
	testStatusCode(t, w.Code, http.StatusOK)

	// test the header, footer, and nav renders
	testCommonTemplatesRender(t, w.Body.String())

	// test the body of the index
	testBodyContains(t, `<h1>Welcome to yCTF</h1>`, w.Body.String())

	// test the session cookie named yctf-session is set
	testCookieExists(t, defaultSessionName, w.Result().Cookies())

	// test the score from the session is rendered
	testBodyContains(t, `Score: [0/`, w.Body.String())
}

// TestValidateFlag tests validateFlag(), which is unimplemented at this point
// this is is a scaffolded test for when it is implemented.
func TestValidateFlag(t *testing.T) {
	t.Log("Unimplemented!")
}

// testCookieContains verifies a string is found in a list of cookies
func testCookieContains(t *testing.T, want string, got []string) {
	var found bool
	for _, cookie := range got {
		if strings.Contains(cookie, want) {
			found = true
		}
	}
	if !found {
		t.Errorf("No cookie matched.\n Got: %v\nWant: %v\n", got, want)
	}
}

// testCookieExists verifies a named cookie is found in a list of cookies
func testCookieExists(t *testing.T, want string, got []*http.Cookie) {
	var found bool
	for _, cookie := range got {
		if want == cookie.Name {
			found = true
		}
	}
	if !found {
		t.Errorf("No cookie matched.\n Got: %v\nWant: %v\n", got, want)
	}
}

// testStatusCode confirms that the wanted status code matches that received.
func testStatusCode(t *testing.T, want, got int) {
	if want != got {
		t.Errorf("The status code is incorrect.\n Got: %v\nWant %v\n", got, want)
	}
}

// testBodyContains checks the body for a substring to confirm a template rendered.
func testBodyContains(t *testing.T, want, got string) {
	if !strings.Contains(got, want) {
		t.Errorf(
			"The page did not contain a substring expected from the rendered template.\n Got: %v\nWant: %v\n", got, want)
	}
}

// testCommonTemplatesRender is an aggregate function checking that the
// header, footer, and nav rendered.
func testCommonTemplatesRender(t *testing.T, in string) {
	testHeaderRenders(t, in)
	testFooterRenders(t, in)
	testNavRenders(t, in)
}

// testHeaderRenders checks if the header rendered by matching a string
// expected in the header only.
func testHeaderRenders(t *testing.T, in string) {
	// test the header renders
	testBodyContains(t, "<!doctype html>", in)
}

// testFooterRenders checks if the footer renedered by matching a string
// expected in the footer only.
func testFooterRenders(t *testing.T, in string) {
	// test the footer renders
	testBodyContains(t, "</html>", in)
}

// testNavRenders checks if the nav rendered by matching a string expected
// in the body only.
func testNavRenders(t *testing.T, in string) {
	// test the nav renders
	testBodyContains(t, `<nav class="navbar navbar-default">`, in)
}
