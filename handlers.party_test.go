// handlers.party_test.go

package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

func getPartyPOSTPayload() string {
	params := url.Values{}
	params.Add("title", "Test Party Title")
	params.Add("content", "Test Party Content")

	return params.Encode()
}

// Test that a GET request to the home page returns the home page with
// the HTTP code 200 for an unauthenticated user
func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := getRouter(true)

	r.GET("/", showIndexPage)

	// Create a request to send  the above route
	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that http status code is 200
		statusOk := w.Code == http.StatusOK

		// Test that the page title is "Home Page"
		p, err := ioutil.ReadAll(w.Body)
		pageOk := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOk && pageOk
	},
	)
}

// Test that a GET request to display an party returns the party with
// the HTTP code 200 for an unauthenticated user
func TestPartyUnauthenticated(t *testing.T) {
	r := getRouter(true)

	r.GET("/party/view/:party_id", getParty)

	// Create a request to send  the above route
	req, _ := http.NewRequest("GET", "/party/view/1", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that http status code is 200
		statusOk := w.Code == http.StatusOK

		// Test that the page title is "Party 1"
		p, err := ioutil.ReadAll(w.Body)
		pageOk := err == nil && strings.Index(string(p), "<title>Party 1</title>") > 0

		return statusOk && pageOk
	},
	)
}

// Test for party creation will be similar to the test that tested the registration functionality
func TestPartyCreationAuthenticated(t *testing.T) {
	saveLists()
	w := httptest.NewRecorder()

	r := getRouter(true)

	http.SetCookie(w, &http.Cookie{Name: "token", Value: "123"})

	r.POST("/party/create", createParty)

	partyPayload := getPartyPOSTPayload()
	req, _ := http.NewRequest("POST", "/party/create", strings.NewReader(partyPayload))
	req.Header = http.Header{"Cookie": w.HeaderMap["Set-Cookie"]}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(partyPayload)))

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fail()
	}

	p, err := ioutil.ReadAll(w.Body)
	s := string(p)
	if err != nil || strings.Index(s, "<title>Submission") < 0 {
		t.Fail()
	}
	restoreLists()
}
