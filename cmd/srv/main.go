package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"

	"myoidc/internal/server"
)

func main() {
	var listen = flag.String("listen", "", "listen address")
  var secret = flag.String("secret", "", "required secret for encrypting token")
  var issuer = flag.String("issuer", "", "optional issuer")
  var insecure = flag.Bool("insecure", false, "allow insecure OP")
  flag.Parse()


	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)

	srv, err := server.New(&server.Config{
    Listen: *listen,
    Secret: *secret,
    Issuer: *issuer,
    Insecure: *insecure,
  })

	if err != nil {
		log.Printf("server.New err: %v", err)
	}

	err = srv.Run(ctx)
	stop()
	if err != nil {
		log.Printf("srv.Run err: %v", err)
	}
}
