package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	cm "framework/common"
	"framework/middleware"
	"framework/services"
	"framework/transport"

	log "github.com/Sirupsen/logrus"
	httptransport "github.com/go-kit/kit/transport/http"
)

func initHandlers() {
	var svc services.PaymentServices

	svc = services.PaymentService{}
	svc = middleware.BasicMiddleware()(svc)

	root := cm.Config.RootURL
	fmt.Println(cm.Config.RootURL)

	// orders endpoint
	http.Handle(fmt.Sprintf("%s/order", root), httptransport.NewServer(
		transport.OrderEndpoint(svc), transport.DecodeOrderRequest, transport.EncodeResponse,
	))

	// customers endpoint
	http.Handle(fmt.Sprintf("%s/customer", root), httptransport.NewServer(
		transport.CustomerEndpoint(svc), transport.DecodeCustomerRequest, transport.EncodeResponse,
	))

	// customers endpoint
	http.Handle(fmt.Sprintf("%s/product", root), httptransport.NewServer(
		transport.ProductEndpoint(svc), transport.DecodeProductRequest, transport.EncodeResponse,
	))

	// faspay endpoint
	http.Handle(fmt.Sprintf("%s/faspay", root), httptransport.NewServer(
		transport.FaspayEndpoint(svc), transport.DecodeFaspayRequest, transport.EncodeResponse,
	))
}

var logger *log.Entry

func initLogger() {
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.999",
	})

	//log.SetReportCaller(true)
}

func main() {
	configFile := flag.String("conf", "conf-dev.yml", "main configuration file")
	flag.Parse()
	initLogger()
	log.WithField("file", *configFile).Info("Loading configuration file")
	cm.LoadConfigFromFile(configFile)
	initHandlers()

	var err error
	if cm.Config.RootURL != "" || cm.Config.ListenPort != "" {
		err = http.ListenAndServe(cm.Config.ListenPort, nil)
	}

	if err != nil {
		log.WithField("error", err).Error("Unable to start the server")
		os.Exit(1)
	}
}
