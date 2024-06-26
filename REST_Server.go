package main

import (
	"fmt"
	"log"
	"net/http"
)

const userApiResp = `
<html> 
<body>
<p>Hi, thanks for calling my /users API with HTTP method '%v'
<p>This is the %v call
</body>
</html>
`

var userCounter int

func main() {
	http.HandleFunc("/users", usersHandlefunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func usersHandlefunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("We got a request on /users")
	userCounter++
	s := fmt.Sprintf(userApiResp, r.Method, userCounter)
	fmt.Fprintf(w, s)
}
