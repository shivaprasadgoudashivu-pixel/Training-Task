<<<<<<< HEAD
package main

import (
	"demo/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", handler.Ping) // register handler with standard net/http
	http.ListenAndServe(":8080", nil)      // start server on port 8080
}
=======
package main

import (
	"demo/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", handler.Ping) // register handler with standard net/http
	http.ListenAndServe(":8080", nil)      // start server on port 8080
}
>>>>>>> 6b7cfda (changes on MFtask)
