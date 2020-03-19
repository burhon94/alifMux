package main

import (
	"github.com/burhon94/alifMux/pkg/mux"
	"log"
	"net/http"
)

func main() {
	mu := mux.NewExactMux()
	mu.GET("/posts/hot", func(writer http.ResponseWriter, request *http.Request) {
		log.Print("hot")
	})
	mu.GET("/posts/{id}", func(writer http.ResponseWriter, request *http.Request) {
		value, _ := mux.FromContext(request.Context(), "id")
		log.Printf("id: %s", value)
	})
	mu.GET("/posts/{postId}/comments/{commentId}", func(writer http.ResponseWriter, request *http.Request) {
		postId, _ := mux.FromContext(request.Context(), "postId")
		commentId, _ := mux.FromContext(request.Context(), "commentId")
		log.Printf("postId: %s, commentId: %s", postId, commentId)
	})

	panic(http.ListenAndServe("0.0.0.0:9999", mu))
}
