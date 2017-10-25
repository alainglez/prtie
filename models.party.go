// models.party.go

package main

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/lib/pq"
)

type party struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type amenity struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var partyList = []party{
	party{ID: 1, Title: "Party 1", Content: "Party 1 body"},
	party{ID: 2, Title: "Party 2", Content: "Party 2 body"},
}

var amenityList = make([]amenity, 1)

func createTables() bool {
	// Connect to the "party" database.
	db, err := sql.Open("postgres", "postgresql://alain@localhost:26257/party?sslmode=disable")
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	//create the "amenities" table
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS amenities (id INT PRIMARY KEY, name STRING NOT NULL);"); err != nil {
		log.Fatal(err)
	}

	//create the "parties" table
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS parties (id SERIAL PRIMARY KEY, name STRING NULL, inviteMessage STRING NULL, organizer INT NOT NULL, partyType INT NULL, place INT NOT NULL, partyDate DATE NOT NULL, startTime TIMESTAMPTZ NOT NULL, endTime TIMESTAMPTZ NULL, isPrivate BOOL DEFAULT TRUE, isFree BOOL DEFAULT TRUE, attendanceFee DECIMAL NULL, guestLimit INT NULL, giftShopURL STRING NULL, foodShopURL STRING NULL, drinksShopURL STRING NULL, amenities INT ARRAY NULL, vipfee DECIMAL NULL, iscancelled BOOL NULL DEFAULT false, lastUpdated TIMESTAMPTZ DEFAULT current_timestamp())"); err != nil {
		log.Fatal(err)
	}

	//create the "users" table
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, userName STRING NOT NULL, phone STRING NOT NULL, email STRING NULL, password STRING NOT NULL, firstName STRING NOT NULL, lastName STRING NOT NULL, dateOfBirth DATE NOT NULL, paymentDetailsID INT NULL, isprivate BOOL NOT NULL DEFAULT false, lastUpdated TIMESTAMPTZ DEFAULT current_timestamp())"); err != nil {
		log.Fatal(err)
	}

	//create the "partiers" table
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS partiers (partyId INT, userId INT, isOrganizer BOOL DEFAULT FALSE, isConfirmed BOOL DEFAULT FALSE, isNotified BOOL DEFAULT FALSE, hasPaid BOOL DEFAULT FALSE, guests INT DEFAULT 0, passImage BYTES NULL, passRedeemed BOOL DEFAULT FALSE, isvip BOOL NULL DEFAULT false, isbanned BOOL NULL DEFAULT false, attended BOOL NULL, lastUpdated TIMESTAMPTZ DEFAULT current_timestamp(), PRIMARY KEY(partyId, userId))"); err != nil {
		log.Fatal(err)
	}

	//create the "places" table
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS places (id SERIAL PRIMARY KEY, name STRING NOT NULL, street STRING NOT NULL, city STRING NOT NULL, state INT NOT NULL, zip STRING NOT NULL, country INT NOT NULL, isbanned BOOL NOT NULL DEFAULT false, lastUpdated TIMESTAMPTZ DEFAULT current_timestamp())"); err != nil {
		log.Fatal(err)
	}

	return true
}

// Return a list of all active parties - from yesterday and into the future
func getAllParties() []party {
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

// Return a list of all active parties - from yesterday and into the future
func getAllAmenities() []amenity {
	// Connect to the "party" database.
	db, err := sql.Open("postgres", "postgresql://root@localhost:26257/party?sslmode=disable")
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	// Populate the amenitiesList slice.
	rows, err := db.Query("SELECT id, name FROM amenities")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var i, id int
	var name string
	for rows.Next() {
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		if i == 0 {
			amenityList[i] = amenity{id, name}
		} else {
			amenityList = append(amenityList, amenity{id, name})
		}
		//amenityList[i].Name = name
		i++
	}

	return amenityList
}
