package a

type application struct {

}

func (a *application) handleRequest(url, method string) (int, string){
	if url=="app/status" && method == "GET" {
		return 200,  "OK"

	}
	if url=="create/user" && method=="POST"{
		return 201, "Created user"
	}

	return 404, "Not found"
}
