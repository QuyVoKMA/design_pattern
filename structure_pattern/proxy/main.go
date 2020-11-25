package main

import (
	"fmt"
	p "proxy/a"
)

func main() {
	nginxServer := p.NewNginxServer()
	appStatusURL :="app/status"
	createUser :="create/user"

	httpCode, message :=nginxServer.HandleRequest(appStatusURL, "GET")
	fmt.Printf("\nURL: %s\nHttpCode: %d;\nMessage: %s\n", appStatusURL, httpCode, message)

	httpCode, message =nginxServer.HandleRequest(appStatusURL, "GET")
	fmt.Printf("\nURL: %s\nHttpCode: %d;\nMessage: %s\n", appStatusURL, httpCode, message)

	httpCode, message =nginxServer.HandleRequest(appStatusURL, "GET")
	fmt.Printf("\nURL: %s\nHttpCode: %d;\nMessage: %s\n", appStatusURL, httpCode, message)



	httpCode, message =nginxServer.HandleRequest(createUser, "POST")
	fmt.Printf("\nURL: %s\nHttpCode: %d;\nMessage: %s\n", appStatusURL, httpCode, message)

	httpCode, message =nginxServer.HandleRequest(createUser, "GET")
	fmt.Printf("\nURL: %s\nHttpCode: %d;\nMessage: %s\n", appStatusURL, httpCode, message)
}
