package main

import (
  "context"
  "log"
  "os"
  "net/http"

  "github.com/DedovR/events_test/gateway"
  "github.com/DedovR/events_test/server"
  "github.com/joho/godotenv"
  "go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
  err := godotenv.Load()
  if err != nil {
      log.Fatalf("err loading: %v", err)
  }

  clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
  clientOptions.SetAuth(options.Credential{
    Username: os.Getenv("MONGODB_USER"),
    Password: os.Getenv("MONGODB_PASSWORD"),
  })
  client, err := mongo.Connect(clientOptions)
  if err != nil {
    log.Fatal(err)
  }
  log.Println("mongo connected...")

  defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

  srv := gateway.NewServer()
  r := http.NewServeMux()
  h := api.HandlerFromMux(srv, r)

  log.Println(os.Getenv("HTTP_ADDR"))
  s := &http.Server{
    Handler: h,
    Addr:    os.Getenv("HTTP_ADDR"),
  }
  log.Println("server init...")

  log.Fatal(s.ListenAndServe())
}
