package transport

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httputil"

	cm "framework/common"

	ex "framework/error"

	log "github.com/Sirupsen/logrus"
)

func DecodeOrderRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var body []byte

	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.WithField("error", err).Error("Exception caught")
	}
	log.Debug(string(requestDump))

	//request.IPAddress = r.RemoteAddr

	//decode request body
	body, err = ioutil.ReadAll(r.Body)
	log.WithField("info", string(body[:])).Info("Decode Request Simobi Payment API")
	if err != nil {
		return ex.Error(err, 100).Rem("Unable to read request body"), nil
	}

	var request cm.Order

	if err = json.Unmarshal(body, &request); err != nil {
		return ex.Error(err, 100).Rem("Failed decoding json message"), nil
	}

	return request, nil

	//return nil, nil
}

func DecodeCustomerRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var body []byte

	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.WithField("error", err).Error("Exception caught")
	}
	log.Debug(string(requestDump))

	body, err = ioutil.ReadAll(r.Body)
	log.WithField("info", string(body[:])).Info("Decode Request Customers")
	if err != nil {
		return ex.Error(err, 100).Rem("Unable to read request body"), nil
	}

	var request cm.Customer
	if err = json.Unmarshal(body, &request); err != nil {
		return ex.Error(err, 100).Rem("Failed decoding json message"), nil
	}

	return request, nil
}

func DecodeProductRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var body []byte

	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.WithField("error", err).Error("Exception caught")
	}
	log.Debug(string(requestDump))

	body, err = ioutil.ReadAll(r.Body)
	log.WithField("info", string(body[:])).Info("Decode Request Product")
	if err != nil {
		return ex.Error(err, 100).Rem("Unable to read request body"), nil
	}

	var request cm.Product
	if err = json.Unmarshal(body, &request); err != nil {
		return ex.Error(err, 100).Rem("Failed decoding json message"), nil
	}

	return request, nil
}

func DecodeFaspayRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var body []byte

	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.WithField("error", err).Error("Exception caught")
	}
	log.Debug(string(requestDump))

	body, err = ioutil.ReadAll(r.Body)
	log.WithField("info", string(body[:])).Info("Decode Request Faspay")
	if err != nil {
		return ex.Error(err, 100).Rem("Unable to read request body"), nil
	}

	var request cm.RequestFaspay
	if err = json.Unmarshal(body, &request); err != nil {
		return ex.Error(err, 100).Rem("Failed decoding json message"), nil
	}

	return request, nil
}

func DecodeTripRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var body []byte

	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.WithField("error", err).Error("Exception caught")
	}
	log.Debug(string(requestDump))

	body, err = ioutil.ReadAll(r.Body)
	log.WithField("info", string(body[:])).Info("Decode Request Trip")
	if err != nil {
		return ex.Error(err, 100).Rem("Unable to read request body"), nil
	}

	var request cm.RequestTrip
	if err = json.Unmarshal(body, &request); err != nil {
		return ex.Error(err, 100).Rem("Failed decoding json message"), nil
	}

	return request, nil
}

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	var body []byte
	body, err := json.Marshal(&response)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")

	if _, ok := response.(int); ok {
		//respond back to backend
		var e = response.(int)
		if e < 2 {
			w.WriteHeader(http.StatusOK)
		} else if e < 90 {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else if _, ok := response.(int); ok {
		w.WriteHeader(http.StatusOK)
	}

	_, err = w.Write(body)

	return err
}
