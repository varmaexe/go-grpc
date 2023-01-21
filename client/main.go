package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/varmaexe/go-grpc/proto"
	"google.golang.org/grpc"
)

func main() {
	//communicating with 2nd microservice
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	client := pb.NewGreetUserServiceClient(conn)

	g := gin.Default()

	//parses the html files
	g.LoadHTMLGlob("./templates/*")

	g.GET("/", func(ctx *gin.Context) {

		name := ctx.Request.FormValue("name")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//req receives the user input from the client
		req := &pb.Request{Name: name}

		//passing user input to the server
		r, err := client.GreetUser(ctx, req)
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}

		//Gets the greeting message from the server side and
		// prints in the 1st microservice terminal
		log.Printf("Greetings: %s", r.GetGreetings())

		// prints the greeting msg to the client side i.e frontend
		ctx.HTML(200, "index.html", name)

	})
	g.Run()
}
