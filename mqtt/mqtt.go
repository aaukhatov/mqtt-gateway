// Здесь интферфейс взаимодействия с датчиками
package mqtt

import (
	"github.com/eclipse/paho.mqtt.golang"
	"net/url"
	"time"
	"log"
	"fmt"
)

func connect(clientId string, uri *url.URL) mqtt.Client {
	opts := createClientOptions(clientId, uri)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}

func createClientOptions(clientId string, uri *url.URL) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	opts.SetUsername(uri.User.Username())
	password, _ := uri.User.Password()
	opts.SetPassword(password)
	opts.SetClientID(clientId)
	return opts
}

func listen(uri *url.URL, topic string) {
	log.Println("Try to connect")
	client := connect("sub", uri)
	log.Println("Connected")
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		log.Println("Subscribe...")
		fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
	})
}

func Test(mqttUrl string) {
	uri, err := url.Parse(mqttUrl)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("URL parsed", uri)
	topic := uri.Path[1:len(uri.Path)]
	if topic == "" {
		topic = "linoxide/topic"
	}

	go listen(uri, topic)

	client := connect("pub", uri)
	timer := time.NewTicker(1 * time.Second)
	for t := range timer.C {
		client.Publish(topic, 0, false, t.String())
	}
}