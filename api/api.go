// Здесь интерфейс доступа к веб-сервису
package api
import (
	"net/http"
	"log"
	"fmt"
	"encoding/json"
)

// принимает структуру MessageRequest
func SendMessage(writer http.ResponseWriter, request *http.Request) {
	var messageRequest MessageRequest
	decoder  := json.NewDecoder(request.Body)
	err := decoder.Decode(&messageRequest)
	if err != nil {
		log.Println("Couldn't parse MessageRequest", err)
	}
	defer request.Body.Close()
	buildHttpHeader201(writer)
	//для демонстрации ответ упаковываем в json
	response, err := json.Marshal(messageRequest)
	fmt.Fprintf(writer, string(response))
}

func GetEspList(writer http.ResponseWriter, request *http.Request) {

}

func buildHttpHeader201(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json;charset=UTF-8")
	writer.WriteHeader(http.StatusCreated)
}

func buildHttpHeader200(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json;charset=UTF-8")
	writer.WriteHeader(http.StatusOK)
}
// структура ответа
type Response struct {
	Code string		`json:"code"`
}
// структура запроса
type MessageRequest struct {
	// a phone number
	Number string	`json:"number"`
	// sms text
	Text string		`json:"text"`
	// need receipt?
	Receipt bool	`json:"receipt"`
}

// массив запросов
type Requests []MessageRequest