// Здесь интерфейс доступа к веб-сервису
package api
import (
	"net/http"
	"log"
	"fmt"
	"encoding/json"
)

// принимает структуру SmsRequest
func SendSms(writer http.ResponseWriter, request *http.Request) {
	var smsRequest SmsRequest
	decoder  := json.NewDecoder(request.Body)
	err := decoder.Decode(&smsRequest)
	if err != nil {
		log.Println("Couldn't parse SmsRequest", err)
	}
	defer request.Body.Close()
	buildHttpHeader201(writer)
	//для демонстрации ответ упаковываем в json
	response, err := json.Marshal(smsRequest)
	fmt.Fprintf(writer, "Received request: %v", string(response))
}

func buildHttpHeader201(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json;charset=UTF-8")
	writer.WriteHeader(http.StatusCreated)
}
// структура ответа
type Response struct {
	Code string		`json:"code"`
}
// структура запроса
type SmsRequest struct {
	// a phone number
	Number string	`json:"number"`
	// sms text
	Text string		`json:"text"`
	// need receipt?
	Receipt bool	`json:"receipt"`
}

// массив запросов
type Requests []SmsRequest