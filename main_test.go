package main

import (
	"testing"
	"net/http"
	"bytes"
	"net/http/httptest"
	"github.com/aukhatov/mqtt-gateway/api"
)

func TestSendSms(t *testing.T) {
	requestPayload := []byte(`{"number": "8-800", "text": "super code 666", "receipt": false}`)
	req, _ := http.NewRequest("POST", "/sendSms", bytes.NewBuffer(requestPayload))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.SendSms)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Failed()
	}
}
