package netboxapi

import (
	"log"
	"net/http"
	"strings"

	"github.com/thunderstorm99/apihandler"
)

// NetboxConnection is a type that holds all necessary info
type NetboxConnection struct {
	Token    string
	BaseURL  string
	Insecure bool
}

// APIAnswer is the answer from the netbox API
type APIAnswer struct {
	Count   int    `json:"count"`
	Next    string `json:"next"`
	Results []any  `json:"results"`
}

// NewNetboxConnection creates a new connection to a netbox via url and token
func NewNetboxConnection(url string, token string) NetboxConnection {
	return NetboxConnection{Token: token, BaseURL: url}
}

// getAPI executes an API call with an array as an answer
func (n *NetboxConnection) getAPI(url string, output any) error {
	a, err := n.getAPIRaw(url)
	if err != nil {
		return err
	}

	err = resultToOtherStructure(a, &output)
	if err != nil {
		return err
	}

	return nil
}

// getAPISingle executes an API call with a single answer
func (n *NetboxConnection) getAPISingle(url string, output any) error {
	api := apihandler.APICall{URL: n.BaseURL + url, Method: http.MethodGet, Header: map[string]string{"Authorization": "Token " + n.Token}, Insecure: n.Insecure}

	// calling API
	log.Println("calling URL", api.URL)
	_, _, err := api.Exec(&output)
	return err
}

// getAPIRaw performs a GET request onto any netbox API endpoint specified with url
func (n *NetboxConnection) getAPIRaw(url string) (APIAnswer, error) {
	api := apihandler.APICall{URL: n.BaseURL + url, Method: http.MethodGet, Header: map[string]string{"Authorization": "Token " + n.Token}, Insecure: n.Insecure}

	// create new variable that holds the answer of the API
	var answer APIAnswer

	// calling API
	// log.Println("calling URL", api.URL)
	_, _, err := api.Exec(&answer)
	if err != nil {
		return APIAnswer{}, err
	}

	// check for pagination
	if answer.Next == "" {
		return answer, nil
	}

	// check if BaseURL is https
	// check if Next URL is http (which it should not be), but if so upgrade to https
	if strings.HasPrefix(n.BaseURL, "https://") && strings.HasPrefix(answer.Next, "http://") {
		// upgrade http to https
		answer.Next = strings.Replace(answer.Next, "http://", "https://", 1)
	}

	// trim the base off of the next url
	nextURL := strings.TrimPrefix(answer.Next, n.BaseURL)
	next, err := n.getAPIRaw(nextURL)
	if err != nil {
		return APIAnswer{}, err
	}

	// append all results of next to these results
	answer.Results = append(answer.Results, next.Results...)

	return answer, nil
}
