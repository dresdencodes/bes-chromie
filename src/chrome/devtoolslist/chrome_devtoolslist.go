package chromeattach

import (
	"errors"
    "net/http"
    "encoding/json"
)

type DevToolsTarget struct {
    ID      string 					`json:"id"`
    Type    string 					`json:"type"`
    URL     string 					`json:"url"`
    WebSocketDebuggerURL string 	`json:"webSocketDebuggerUrl"`
}

var chromeListURL = "http://localhost:9222/json"

func getTab() (string, error) {

    // Step 1: Get existing targets
    resp, err := http.Get(chromeListURL)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    var targets []DevToolsTarget
    err = json.NewDecoder(resp.Body).Decode(&targets); 
	if err != nil {
        return "", err
    }

    // Step 2: Pick the first "page" type tab
    var tabURL string
    for _, target := range targets {
        if target.Type == "page" {
            tabURL = target.WebSocketDebuggerURL
            break
        }
    }

	// tab url
    if tabURL == "" { return "", errors.New("") }

	// return tab url
	return tabURL, nil

}