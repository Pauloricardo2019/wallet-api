package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"wallet-api/adapter/rest"
)

var chanSignal chan os.Signal

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @termsOfService http://swagger.io/terms/
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	rest.SetupHttpEngine()
	handleExit()
	defer handlePanic()

	/*
		Deixa o sistema em 'loop'
		Caso essa linha seja apagada, o sistema vai iniciar e finalizar em milésimos de segundo,
	*/
	for {
		time.Sleep(2 * time.Second)
	}
}

func handlePanic() {
	if err := recover(); err != nil {
		chanSignal <- syscall.SIGTERM
		for {
			time.Sleep(2 * time.Second)
		}
	}
}

func handleExit() {
	chanSignal = make(chan os.Signal, 1)
	signal.Notify(chanSignal, os.Interrupt)

	go func() {
		for exitSignal := range chanSignal {
			fmt.Printf("Capturing signal... %v\n", exitSignal)
			ctx := context.Background()
			rest.StopRest(ctx)

			fmt.Println("Exiting system")
			os.Exit(0)
		}
	}()
}
