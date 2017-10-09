// models.party_test.go

package main

import "testing"

// Test the function that fetches all parties
func TestGetAllparties(t *testing.T) {
	alist := getAllparties()

	// Check that the length of the list of parties returned is the
	// same as the length of the global variable holding the list
	if len(alist) != len(partyList) {
		t.Fail()
	}

	// Check that each member is identical
	for i, v := range alist {
		if v.Content != partyList[i].Content ||
			v.ID != partyList[i].ID ||
			v.Title != partyList[i].Title {

			t.Fail()
			break
		}
	}
}

// Test the function that fetches an Party by its ID
func TestGetPartyByID(t *testing.T) {
	a, err := getPartyByID(1)

	if err != nil || a.ID != 1 || a.Title != "Party 1" || a.Content != "Party 1 body" {
		t.Fail()
	}
}

// Test the createNewParty by creating a new party and checking that
// party list contains the new party
func TestCreateNewParty(t *testing.T) {
	originalLength := len(getAllparties())

	a, err := createNewParty("New test party", "New test content")

	allParty := getAllparties()
	newLength := len(allParty)

	if err != nil || newLength != originalLength+1 ||
		a.Title != "New test party" || a.Content != "New test content" {

		t.Fail()
	}
}
