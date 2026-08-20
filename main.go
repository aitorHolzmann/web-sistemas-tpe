package main

import (
	"fmt"
	"net/http"
)

func main(){
  dirPath := "./static"
  server := http.FileServer(http.Dir(dirPath))

  http.Handle("/", server)
  
  port := ":8080"
  fmt.Printf("Servidor escuchand en https://localhost%s\n", port)

  err := http.ListenAndServe(port, nil)
  if (err != nil){
    fmt.Printf("Error al iniciar el servidor: %s\n", err)
  }

}
