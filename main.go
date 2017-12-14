package main

import (
	"html/template"
	"net/http"
	"log"
	"os"
	"io"
	"regexp"
	"time"
	"github.com/aukhatov/mqtt-gateway/api"
	"github.com/gorilla/mux"
)

const defaultHttpPort = ":80"
const MQTT_URL = "tcp://mqtt-spy:123@192.168.1.133:1883/"
const ESP_API = "/ESP/"

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
	defineHandlers(rest)
	// запуск веб-сервиса
	log.Println("Web service has been started on port", httpPort[1:])
	err := http.ListenAndServe(httpPort, rest.Router)
	if err != nil {
		log.Fatalf("Couldn't start web service. %v", err)
	}
}

func defineHandlers(rest *RestService) {
	// home controller
	rest.Router.HandleFunc("/", defaultHandler).Methods("GET")

	rest.Router.HandleFunc("/esp", api.GetEspList).Methods("GET")
	rest.Router.HandleFunc("/esp/{chipId:[0-9]+}", api.LedOnOff).Methods("POST")
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
}

func viewTemplate(writer http.ResponseWriter) {
	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(writer, nil)
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
