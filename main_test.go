package main

import (
	"testing"
	"net/http"
	"bytes"
	"net/http/httptest"
	"github.com/aukhatov/mqtt-gateway/api"
)

func TestSendMessage(t *testing.T) {
	requestPayload := []byte(`{"number": "8-800", "text": "super code 666", "receipt": false}`)
	req, _ := http.NewRequest("POST", "/esp", bytes.NewBuffer(requestPayload))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.SendMessage)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Failed()
	}
}
