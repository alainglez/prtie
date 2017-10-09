// models.party.go

package main

import (
	"errors"

	_ "github.com/lib/pq"
)

type party struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var partyList = []party{
	party{ID: 1, Title: "Party 1", Content: "Party 1 body"},
	party{ID: 2, Title: "Party 2", Content: "Party 2 body"},
}

// Return a list of all the parties
func getAllparties() []party {
	/*
		// Connect to the "party" database.
		db, err := sql.Open("postgres", "postgresql://alain@localhost:26257/party?sslmode=disable")
		if err != nil {
			log.Fatal("error connecting to the database: ", err)
		}
		// Populate the partyList slice.
		rows, err := db.Query("SELECT id, title, content FROM parties")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			var id int
			var title, content string
			if err := rows.Scan(&id, &title, &content); err != nil {
				log.Fatal(err)
			}
			partyList
		}
	*/
	return partyList
}

// Return an party by ID
func getPartyByID(id int) (*party, error) {
	for _, a := range partyList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Party not found")
}

func createNewParty(title, content string) (*party, error) {
	a := party{ID: len(partyList) + 1, Title: title, Content: content}

	partyList = append(partyList, a)

	return &a, nil
}
