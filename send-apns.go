package main

import (
	"flag"
	"fmt"
	"github.com/hoverruan/Go-Apns"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const (
	// Sandbox env
	sandboxGateway  = "gateway.sandbox.push.apple.com:2195"
	sandboxFeedback = "feedback.sandbox.push.apple.com:2196"
	sandboxCert     = "dev-cert.pem"
	sandboxKey      = "dev-key.pem"

	// Production env
	productGateway  = "gateway.push.apple.com:2195"
	productFeedback = "feedback.push.apple.com:2196"
	productCert     = "prod-cert.pem"
	productKey      = "prod-key.pem"
)

var (
	prod  = flag.Bool("p", false, "Using production destination")
	sound = flag.String("s", "", "Sound")

	customFieldsInput = flag.String("C", "", "Setting custom fields, separated with comma, eg: key1=value1,key2=value2")
)

func isBlank(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func main() {
	flag.Parse()

	token, body := "", ""
	var badge uint64 = 0

	if len(flag.Args()) > 1 {
		token = flag.Args()[0]
		body = flag.Args()[1]
	}

	if len(flag.Args()) > 2 {
		badge, _ = strconv.ParseUint(flag.Args()[2], 10, 32)
	}

	if isBlank(token) || isBlank(body) {
		fmt.Fprintf(os.Stderr, "Usage: %s [options...] <token> <body> [badge]\n"+
			"Options:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr)

		os.Exit(-1)
	}

	customFields := make(map[string]interface{})
	if !isBlank(*customFieldsInput) {
		parts := strings.Split(*customFieldsInput, ",")
		for _, part := range parts {
			keyValuePair := strings.Split(part, "=")
			key := keyValuePair[0]
			value := keyValuePair[1]
			customFields[key] = value
		}
	}

	var cert, key, gateway string
	if *prod {
		cert, key, gateway = productCert, productKey, productGateway
	} else {
		cert, key, gateway = sandboxCert, sandboxKey, sandboxGateway
	}

	apn, err := goapns.Connect(cert, key, gateway)
	if err != nil {
		log.Fatalf("Connect error: %s\n", err.Error())
	}

	notify := goapns.Notification{
		DeviceToken: token,
		Aps: goapns.SimpleAps{
			Alert: body,
			Badge: uint(badge),
			Sound: *sound,
		},
		Identifier:  uint32(rand.Intn(10000)),
		CustomFiels: customFields,
	}

	model, _ := notify.MarshalJSON()
	fmt.Println("Sending", string(model))
	err = apn.SendNotification(&notify)
	if err != nil {
		log.Fatalf("Send error: %s\n", err.Error())
	}
}
