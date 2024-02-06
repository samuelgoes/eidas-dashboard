package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/samuelgoes/eidas-dashboard/models"
)

const (
	eidasV1Host              = "https://eidas.ec.europa.eu/efda/tl-browser/api/v1/"
	endpointListByCountry    = "browser/tl/"
	endpointSearchContryList = "search/countries_list"
	endpointSearchTspList    = "search/tsp_list"
)

func GetTrustedListByCountry(countryCode string) (*models.TrustedListDTO, error) {
	url := eidasV1Host + endpointListByCountry + countryCode
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)

	var trustedList models.TrustedListDTO
	err = json.NewDecoder(response.Body).Decode(&trustedList)
	if err != nil {
		return nil, err
	}

	return &trustedList, nil
}

func GetTrustedList() ([]models.TspDTO, error) {
	url := eidasV1Host + endpointSearchTspList
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)

	var trustedList []models.TspDTO
	err = json.NewDecoder(response.Body).Decode(&trustedList)
	if err != nil {
		return nil, err
	}

	return trustedList, nil
}

func SearchForTSPByTrustmark(trustedList *models.TrustedListDTO, trustmark string) (string, error) {
	// Iterate through the trusted list to find the TSP with the given trustmark
	/*for _, pointer := range trustedList. {
		// Here you can add your own logic to search for a specific trustmark
		// For simplicity, let's assume we are searching for a TSP with a specific trustmark
		// You can replace this logic with your actual search criteria
		if pointer.TSLPointer.SchemeOperatorName == trustmark {
			return pointer.TSLPointer.SchemeOperatorName, nil
		}
	}*/
	return "", fmt.Errorf("TSP with trustmark '%s' not found", trustmark)
}
