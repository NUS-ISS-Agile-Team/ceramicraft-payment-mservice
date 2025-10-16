package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/NUS-ISS-Agile-Team/ceramicraft-payment-mservice/server/config"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-payment-mservice/server/grpc"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-payment-mservice/server/http"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-payment-mservice/server/log"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-payment-mservice/server/metrics"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-payment-mservice/server/mq"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-payment-mservice/server/repository"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-user-mservice/common/utils"
)

var (
	sigCh = make(chan os.Signal, 1)
)

func main() {
	config.Init()
	log.InitLogger()
	repository.Init()
	utils.InitJwtSecret()
	mq.Init()
	metrics.RegisterMetrics()
	go grpc.Init(sigCh)
	go http.Init(sigCh)
	// listen terminage signal
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigCh // Block until signal is received
	log.Logger.Infof("Received signal: %v, shutting down...", sig)
}
