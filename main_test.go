package main

import (
	"testing"
	"net/http"
	"bytes"
	"fmt"
	"net/http/httptest"
	"github.com/aukhatov/MqttService/api"
)

func TestSendSms(t *testing.T) {
	requestPayload := []byte(`{"number": "8-800", "text": "super code 666", "receipt": false}`)
	req, _ := http.NewRequest("POST", "/sendSms", bytes.NewBuffer(requestPayload))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.SendSms)
	handler.ServeHTTP(rr, req)
	fmt.Println(rr.Body)
}
