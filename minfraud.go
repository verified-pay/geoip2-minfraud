package geoip2

import (
	"encoding/json"
	"errors"
	"github.com/verified-pay/geoip2-minfraud/model"
	"net/http"
)

var ErrReceived = errors.New("received error from Minfraud API")

type MinFraud struct {
	AccountID  string
	LicenseKey string
	Req        *model.MinfraudReq

	client *MaxmindClient
}

func NewMinFraud(accountID string, licenseKey string) (*MinFraud, error) {
	// official PHP API has options: hashEmail, locales and in client proxy etc..

	return &MinFraud{
		AccountID:  accountID,
		LicenseKey: licenseKey,
		Req:        &model.MinfraudReq{},
		client:     NewMaxmindClient(accountID, licenseKey),
	}, nil
}

// Insights returns the MaxMind Risk Score along with device and user data.
// The 3rd return value is if error == ErrReceived.
func (mf *MinFraud) Insights() (*model.Insights, error, *model.Error) {
	bodyBytes, statusCode, err := mf.client.Post("/insights", mf.Req)
	if err != nil {
		return nil, err, nil
	} else if statusCode != http.StatusOK {
		res := &model.Error{}
		err = json.Unmarshal(bodyBytes, res)
		if err != nil {
			return nil, err, nil
		}
		return nil, ErrReceived, res
	}

	res := &model.Insights{}
	err = json.Unmarshal(bodyBytes, res)
	if err != nil {
		return nil, err, nil
	}
	return res, nil, nil
}

// Score returns the MaxMind Risk Score.
// The 3rd return value is if error == ErrReceived.
func (mf *MinFraud) Score() (*model.Score, error, *model.Error) {
	bodyBytes, statusCode, err := mf.client.Post("/score", mf.Req)
	if err != nil {
		return nil, err, nil
	} else if statusCode != http.StatusOK {
		res := &model.Error{}
		err = json.Unmarshal(bodyBytes, res)
		if err != nil {
			return nil, err, nil
		}
		return nil, ErrReceived, res
	}

	res := &model.Score{}
	err = json.Unmarshal(bodyBytes, res)
	if err != nil {
		return nil, err, nil
	}
	return res, nil, nil
}

func (mf *MinFraud) Factors() (*model.Factors, error) {
	panic("Factors is not yet implemented") // TODO
}
