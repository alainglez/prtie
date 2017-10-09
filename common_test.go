// common_test.go

package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var tmpPartyList []party
var tmpUserList []user

// This function is used for setup before executing the test functions
func TestMain(m *testing.M) {
	//Set Gin to Test mode
	gin.SetMode(gin.TestMode)

	// Run the other tests
	os.Exit(m.Run())
}

//Helper function to create a router during testing
func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("templates/*")
		r.Use(setUserStatus())
	}
	return r
}

// Helper function to process a request a test its response
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the above request
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

// This function is used to store the main lists into a temporary one
func saveLists() {
	tmpUserList = userList
	tmpPartyList = partyList
}

// This function is used to restore the main lists from the temporary one
func restoreLists() {
	userList = tmpUserList
	partyList = tmpPartyList
}

func testMiddlewareRequest(t *testing.T, r *gin.Engine, expectedHTTPCode int) {
	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		return w.Code == expectedHTTPCode
	})
}
