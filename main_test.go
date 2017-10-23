package main

import (
	"testing"
	"net/http"
	"bytes"
	"fmt"
	"net/http/httptest"
	"github.com/aukhatov/MqttService/api"
)

func TestTravis(t *testing.T) {
	if Travis(2, 2) != 4 {
		t.Error("Expected 4")
	}
}

func TestSendSms(t *testing.T) {
	requestPayload := []byte(`{"number": "8-800", "text": "super code 666", "receipt": false}`)
	req, _ := http.NewRequest("POST", "/sendSms", bytes.NewBuffer(requestPayload))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.SendSms)
	handler.ServeHTTP(rr, req)
	fmt.Println(rr.Body.String())
}
