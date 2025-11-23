package launcher

import (
	"errors"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func waitForChrome(port string) error {
	start := time.Now()
	for {
		if isChromeRunning(port) {
			return nil
		}
		if time.Since(start) > time.Duration(10) * time.Second {return errors.New("starting chrome timed out")}
	}
	return nil
}

func isChromeRunning(port string) bool {
	client := http.Client{Timeout: 100 * time.Millisecond}
	_, err := client.Get("http://localhost:"+port+"/json/version")
	return err == nil
}

type DevToolsMetadata struct {
	Browser              string `json:"Browser"`
	ProtocolVersion      string `json:"Protocol-Version"`
	UserAgent            string `json:"User-Agent"`
	V8Version            string `json:"V8-Version"`
	WebKitVersion        string `json:"WebKit-Version"`
	WebSocketDebuggerURL string `json:"webSocketDebuggerUrl"`
}

func getChromeOutput(port string) (*DevToolsMetadata, error) {
	client := http.Client{Timeout: 100 * time.Millisecond}

	resp, err := client.Get("http://localhost:"+port+"/json/version")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Optional: validate it's valid JSON
	var js DevToolsMetadata
	if err := json.Unmarshal(body, &js); err != nil {
		return nil, fmt.Errorf("invalid JSON: %w", err)
	}

	return &js, nil
}