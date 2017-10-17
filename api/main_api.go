// Здесь интерфейс доступа к веб-сервису
package api
import (
	"net/http"
	"log"
	"fmt"
)

func EnableDevice(writer http.ResponseWriter, request *http.Request) {
	log.Printf("Called EnableDevice method %v", request.RequestURI)
	fmt.Fprint(writer, "Called EnableDevice method")
	if request.Method == "GET" {
		log.Println("Some device enabled")
	} else {
		log.Println("Unsupported method!")
	}
}
