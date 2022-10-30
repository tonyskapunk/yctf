package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestFlag5Handler testes getFlag5()
func TestFlag5Handler(t *testing.T) {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/flag5", nil)
	router.ServeHTTP(w, req)

	// test the status code
	testStatusCode(t, w.Code, http.StatusUnauthorized)

	// test the header, footer, and nav renders
	testCommonTemplatesRender(t, w.Body.String())

	// test the body of the index
	testBodyContains(t, `<a href="./">Only authorized users are allowed`, w.Body.String())

	req.AddCookie(&http.Cookie{
		Name:  "auth",
		Value: "true",
	})
	router.ServeHTTP(w, req)

	testFlag5Success(t, w.Body.String())
}

// testFlag5Success does what's necessary for flag5 to succeed
// and confirms that the received flag is as expected.
func testFlag5Success(t *testing.T, body string) {
	id := 5
	f, err := getFlag(id)
	if err != nil {
		t.Error("The flag at the provided ID was not found: ", id)
	}

	testBodyContains(t, f.Flag, body)
}
