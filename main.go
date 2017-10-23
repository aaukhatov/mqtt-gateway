package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
	"io"
	"regexp"
	"time"
	"github.com/aukhatov/MqttService/api"
	"github.com/gorilla/mux"
)

const defaultHttpPort = ":80"

func main()  {
	loggerInitialize()
	args := readCommandLineArguments()
	httpPort := parseHttpPort(args)
	var rest RestService
	rest.Run(httpPort)
}

type RestService struct {
	Router *mux.Router
}

func (rest *RestService) Run(httpPort string) {
	// обработчики
	rest.Router = mux.NewRouter().StrictSlash(true)
	// home controller
	rest.Router.HandleFunc("/", defaultHandler).Methods("GET")
	// отправка SMS
	rest.Router.HandleFunc("/sendSms", api.SendSms).Methods("POST")
	// запуск веб-сервиса
	err := http.ListenAndServe(httpPort, rest.Router)
	log.Println("Web service has been started on port", httpPort[1:])
	if err != nil {
		log.Fatalf("Couldn't start web service. %v", err)
	}
}


func loggerInitialize() {
	timeSuffix := time.Now().Format("2006-01-02")
	logFileName := "access-" + timeSuffix + ".log"
	accessLogFile, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	multiWriter := io.MultiWriter(os.Stdout, accessLogFile)
	log.SetOutput(multiWriter)
}

// Ничего не делает. Приветственная страница.
func defaultHandler(writer http.ResponseWriter, request *http.Request) {
	log.Printf("%v %v %v", request.Proto, request.Method, request.RequestURI)
	fmt.Fprint(writer, "The web-service is working by Go!")
}

func readCommandLineArguments() []string {
	args := os.Args[1:]
	return args
}

func parseHttpPort(args []string) string {
	var httpPort = defaultHttpPort
	var argumentPattern = "port="
	var valueStartIndex = 5

	for _, entry := range args {
		found, _ := regexp.MatchString(argumentPattern, entry)
		if found {
			httpPort = ":" + entry[valueStartIndex:]
		}
	}
	return httpPort
}

func Travis(first int, second int) int {
	return first + second
}
