package geoip2

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const ApiVersion = "v1.23.0"
const ApiUrl = "https://minfraud.maxmind.com/minfraud/v2.0"

type MaxmindClient struct {
	//validateInput bool

	accountID  string
	licenseKey string
	httpClient *http.Client
}

func NewMaxmindClient(accountID string, licenseKey string) *MaxmindClient {
	return &MaxmindClient{
		accountID:  accountID,
		licenseKey: licenseKey,
		httpClient: &http.Client{
			Timeout: 20 * time.Second,
		},
	}
}

func (c *MaxmindClient) Post(path string, data interface{}) ([]byte, int, error) {
	jsonValue, err := json.Marshal(data)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to marshal MaxMind JSON: %+v", err)
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", ApiUrl, path), bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, 0, fmt.Errorf("failed to create MaxMind req: %+v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Content-Length", fmt.Sprintf("%d", len(jsonValue)))
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(c.accountID+":"+c.licenseKey)))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "minFraud-API/"+ApiVersion)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to do MaxMind req: %+v", err)
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to do MaxMind response body: %+v", err)
	}

	return bodyBytes, resp.StatusCode, nil
}
