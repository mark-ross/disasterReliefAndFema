package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/mark-ross/disasterReliefAndFema/fema"
	"github.com/sanity-io/litter"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type connInfo struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

func createFemaTable(db *sql.DB) error {

	femaQuery := `
CREATE TABLE IF NOT EXISTS fema_disasters(
	id                      CHAR(255) PRIMARY KEY,
	disaster_number         INT,
	FEMA_Declaration_String CHAR(255),
	State                   CHAR(255),
	Declaration_Type        CHAR(255),
	Declaration_Date        DATE,
	Fiscal_Year_Declared    CHAR(4),
	Incident_Type           CHAR(255),
	Declaration_Title       CHAR(255),
	Designated_Area         Char(255),
	Place_Code              Char(255)
);`

	_, err := db.Exec(femaQuery)
	return err
}

func insertDataIntoTable(db *sql.DB, elements []*fema.DisasterDeclarationsElement) error {
	insertQuery := `
INSERT INTO fema_disasters(
	id,
	disaster_number,
	fema_declaration_string,
	state,
	declaration_type,
	declaration_date,
	fiscal_year_declared,
	incident_type, 
	declaration_title,
	designated_area,
	place_code
) values ( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);`

	layout := "2006-01-02T15:04:05.000Z"
	for _, ele := range elements {
		declarationDate, err := time.Parse(layout, ele.DeclarationDate)
		if err != nil {
			log.Println("Unable to parse the date properly: ", err, "\n", litter.Sdump(ele))
			return err
		}

		_, err = db.Exec(insertQuery,
			ele.ID,
			ele.DisasterNumber,
			ele.FEMADeclarationString,
			ele.State,
			ele.DeclarationType,
			declarationDate,
			ele.FiscalYearDeclared,
			ele.IncidentType,
			ele.DeclarationTitle,
			ele.DesignatedArea,
			ele.PlaceCode)
		if err != nil {
			log.Println("Unable to insert the previous data: ", litter.Sdump(ele), "\n:::\n", err, "\n")
			return err
		}
	}
	return nil
}

func main() {
	settingsFile := flag.String("filepath", "./settings.json", "filepath to setting file")

	f, err := os.Open(*settingsFile)
	if err != nil {
		log.Println("Unable to open the settings file: ", err)
		panic(err)
	}

	settingsBytes, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println("Unable to read the settings file: ", err)
		panic(err)
	}

	var settings connInfo
	err = json.Unmarshal(settingsBytes, &settings)
	if err != nil {
		log.Println("Settings file contents malformed: ", err)
		panic(err)
	}

	connectionStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		settings.Host, settings.Port, settings.Username, settings.Password,
		settings.DBName)

	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		log.Println("Unable to open the database connection: ", err)
		return
	}

	err = db.Ping()
	if err != nil {
		log.Println("Unable to ping the database...: ", err)
		return
	}

	err = createFemaTable(db)
	if err != nil {
		log.Println("Unable to perform the FEMA table creation: ", err)
	}

	data, err := fema.GetDisasterDeclarationsV2Data(fema.WithMaxCount(10000))
	if err != nil {
		log.Println("Unable to get the data from FEMA: ", err)
		return
	}

	err = insertDataIntoTable(db, data.DisasterDeclarationSummaries)
	if err != nil {
		log.Println("Unable to insert the data into the database: ", err)
		return
	}
}
