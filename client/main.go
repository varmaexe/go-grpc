package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/varmaexe/go-grpc/proto"
	"google.golang.org/grpc"
)

type Name struct {
	Name string
}

func main() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	client := pb.NewGreetUserServiceClient(conn)

	g := gin.Default()

	g.LoadHTMLGlob("./templates/*")

	g.GET("/", func(ctx *gin.Context) {

		// name, err := strconv.ParseInt(ctx.Param("name"), 10, 64)
		name := ctx.Request.FormValue("name")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		req := &pb.Request{Name: name}

		r, err := client.GreetUser(ctx, req)
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greetings: %s", r.GetGreetings())
		ctx.HTML(200, "index.html", name)
	})
	g.Run()
}
