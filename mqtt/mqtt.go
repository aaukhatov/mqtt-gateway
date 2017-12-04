// Здесь интферфейс взаимодействия с датчиками
package mqtt

import (
	"github.com/eclipse/paho.mqtt.golang"
	"net/url"
	"time"
	"log"
	"fmt"
)
// clientId - sub is subscribe
// clientId - pub is publish
func Connect(clientId string, mqttUrl string) mqtt.Client {
	uri, err := url.Parse(mqttUrl)
	if err != nil {
		log.Fatal("Couldn't parse URI", err)
	}
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

func Listen(mqttUrl string, topic string) {
	client := Connect("sub", mqttUrl)
	for {
		client.Subscribe(topic, 0, printOut())
	}
}
func printOut() func(client mqtt.Client, msg mqtt.Message) {
	return func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
	}
}