// Здесь интерфейс доступа к веб-сервису
package api
import (
	"net/http"
	"log"
	"fmt"
	"encoding/json"
	"strconv"
	"time"

	"github.com/aukhatov/mqtt-gateway/mqtt"
	"github.com/gorilla/mux"
)
const MQTT_URL = "tcp://mqtt-spy:123@192.168.1.133:1883/"
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

func LedOnOff(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	var payload Payload
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		log.Print("Couldn't parse request", err)
	}
	vars := mux.Vars(request)
	chipId, err := strconv.Atoi(vars["chipId"])
	if err != nil {
		log.Fatal(err)
		respondWithError(writer, http.StatusBadRequest, "Invalid chipId")
		return
	}
	topic := fmt.Sprintf("/ESP/%d/CONTROL/LED", chipId)
	log.Printf("Send '%s' to %s", payload.Data, topic)
	go publish(topic, payload)

	//buildHttpHeader201(writer)
	//response, err := json.Marshal(payload)
	//fmt.Fprintf(writer, string(response))
}
func publish(topic string, payload Payload) {
	client := mqtt.Connect("pub", MQTT_URL)
	token := client.Publish(topic, 0, false, payload.Data)
	for !token.WaitTimeout(2 * time.Second) && token.Error() != nil {
		log.Println("Wait...")
	}
	if token.Error() != nil {
		log.Fatal("Error", token.Error())
	}
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

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
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

type Payload struct {
	Data string	`json:"data"`
}