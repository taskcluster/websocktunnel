package main

import (
	"crypto/tls"
	"encoding/base64"
	"log/syslog"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	lSyslog "github.com/sirupsen/logrus/hooks/syslog"

	"github.com/gorilla/websocket"
	mozlog "github.com/mozilla-services/go-mozlogrus"
	"github.com/taskcluster/webhooktunnel/whproxy"
)

// starts proxy on a random port on the system
func main() {
	// Load required env variables
	// Load Hostname
	hostname := os.Getenv("HOSTNAME")
	if hostname == "" {
		panic("hostname required")
	}

	logger := log.New()

	if env := os.Getenv("ENV"); env == "production" {
		// add mozlog formatter
		logger.Formatter = &mozlog.MozLogFormatter{
			LoggerName: "webhookproxy",
		}

		// add syslog hook if addr is provided
		syslogAddr := os.Getenv("SYSLOG_ADDR")
		if syslogAddr != "" {
			hook, err := lSyslog.NewSyslogHook("udp", syslogAddr, syslog.LOG_DEBUG, "proxy")
			if err != nil {
				panic(err)
			}
			logger.Hooks.Add(hook)
		}
	}

	// Load secrets
	signingSecretA := os.Getenv("SECRET_A")
	signingSecretB := os.Getenv("SECRET_B")

	// Load config
	useTLS := os.Getenv("USE_TLS") != ""
	domainHosted := os.Getenv("DOMAIN_HOSTED") != ""

	//load port
	port := os.Getenv("PORT")
	if port == "" {
		if useTLS {
			port = "443"
		} else {
			port = "80"
		}
	}

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	// will panic if secrets are not loaded
	proxy, _ := whproxy.New(whproxy.Config{
		Logger:       logger,
		Upgrader:     upgrader,
		JWTSecretA:   []byte(signingSecretA),
		JWTSecretB:   []byte(signingSecretB),
		Domain:       hostname,
		TLS:          useTLS,
		DomainHosted: domainHosted,
	})

	// TODO: Read TLS config
	server := &http.Server{Addr: ":" + port, Handler: proxy}
	defer func() {
		_ = server.Close()
	}()
	logger.WithFields(log.Fields{
		"server-addr": server.Addr,
		"hostname":    hostname,
	}).Infof("starting server")

	// create tls config and serve
	if useTLS {
		// Load TLS certificates
		tlsKeyEnc := os.Getenv("TLS_KEY")
		tlsCertEnc := os.Getenv("TLS_CERTIFICATE")

		tlsKey, _ := base64.StdEncoding.DecodeString(tlsKeyEnc)
		tlsCert, _ := base64.StdEncoding.DecodeString(tlsCertEnc)
		cert, err := tls.X509KeyPair([]byte(tlsCert), []byte(tlsKey))
		if err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}

		config := &tls.Config{
			Certificates: []tls.Certificate{cert},
		}
		config.BuildNameToCertificate()
		listener, err := tls.Listen("tcp", ":"+port, config)
		if err != nil {
			panic(err)
		}
		_ = server.Serve(listener)
	} else {
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}
}
