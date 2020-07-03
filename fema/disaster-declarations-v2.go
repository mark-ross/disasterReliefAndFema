package fema

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type DisasterDeclarationV2 struct {
	MetaData                     MetaData                       `json:"metadata"`
	DisasterDeclarationSummaries []*DisasterDeclarationsElement `json:"DisasterDeclarationsSummaries"`
}

// DisasterDeclarationV2
//
// This struct contains the definitions as put forth in
// the FEMA Data Fields element of their Disaster Declaration
// Summaries - V2 API.
//
// For more information, please visit:
// https://www.fema.gov/openfema-dataset-disaster-declarations-summaries-v2
type DisasterDeclarationsElement struct {
	// FEMADeclarationString
	//
	// Agency standard method for uniquely identifying Stafford Act
	// declarations - Concatenation of declaration type, disaster
	// number and state code. Ex: DR-4393-NC
	FEMADeclarationString string `json:"femaDeclarationString";storm:"index"`

	// DisasterNumber
	//
	// Sequentially assigned number used to designate an event or
	// incident declared as a disaster. For more information on the
	// disaster process, please visit https://www.fema.gov
	DisasterNumber int `json:"disasterNumber";storm:"unique"`

	// State
	//
	// The name or phrase describing the U.S. state, districtor, or territory
	State string `json:"state";storm:"index"`

	// DeclarationType
	//
	// Two character code that defines if this is a major disaster,
	// fire management, or emergency declaration. For more information
	// on the disaster process, please visit https://www.fema.gov
	DeclarationType string `json:"declarationType";storm:"index"`

	// DeclarationDate
	//
	// Date the disaster was declared
	DeclarationDate string `json:"declarationDate";storm:"index"`

	// FiscalYearDeclared
	//
	// Fiscal year in which the disaster was declared
	FiscalYearDeclared int `json:"fyDeclared"`

	// IncidentType
	//
	// Type of incident such as fire or flood. The incident type will
	// affect the types of assistance available. For more information
	// on incident types, please visit https://www.fema.gov
	IncidentType string `json:"incidentType"`

	// DeclarationTitle
	//
	// Title for the disaster
	DeclarationTitle string `json:"declarationTitle"`

	// IndividualAndHouseholdProgramDeclared
	//
	// Denotes whether the Individuals and Households program was
	// declared for this disaster. For more information on the program,
	// please visit https://www.fema.gov
	IndividualAndHouseholdProgramDeclared bool `json:"ihProgramDeclared"`

	// IndividualAssistanceProgramDeclared
	//
	// Denotes whether the Individual Assistance program was declared
	// for this disaster. For more information on the program, please
	// visit https://www.fema.gov
	IndividualAssistanceProgramDeclared bool `json:"iaProgramDeclared"`

	// TODO: Add paProgramDeclared, hmProgramDeclared

	// The following fields were pulled from the request but were not listed
	// in the API documentation
	IncidentBeginDate        string `json:"incidentBeingDate"`
	IncidentEndDate          string `json:"incidentEndDate"`
	DisasterCloseOutDate     string `json:"disasterCloseOutDate"`
	FIPSStateCode            string `json:"fipsStateCode"`
	FIPESCountyCode          string `json:"fipsCountyCode"`
	PlaceCode                string `json:"placeCode"`
	DesignatedArea           string `json:"designatedArea"`
	DeclarationRequestNumber string `json:"declarationRequestNumber"`
	Hash                     string `json:"hash"`
	LastRefresh              string `json:"lastRefresh"`
	ID                       string `json:"id"`
}

func GetDisasterDeclarationsV2Data(options ...QueryFunc) (*DisasterDeclarationV2, error) {
	// Query the data with the given options and generate the full URL
	queryOpts := resolveQueryOptions(options)
	addr := "https://www.fema.gov/api/open/v2/DisasterDeclarationsSummaries"
	if queryOpts != "" {
		addr += "?" + queryOpts
	}

	// Perform the request and parse the body of the request (should be the JSON struct)
	response, err := http.Get(addr)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// parse the returned results into our defined structure and return
	// the object containing all the info
	var results DisasterDeclarationV2
	err = json.Unmarshal(body, &results)
	if err != nil {
		return nil, err
	}

	return &results, nil
}
