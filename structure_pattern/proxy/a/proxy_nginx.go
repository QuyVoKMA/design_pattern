package a

type Nginx struct {
	application *application
	maxAllowedRequest int //request tối đa của application
	rateLimiter map[string]int //Lưu lại số lần request đến webserver
}

/*Hàm khởi tạo*/

func NewNginxServer() *Nginx{
	return &Nginx{
		application: &application{},
		maxAllowedRequest :2,
		rateLimiter: make(map[string]int),
	}
}

func (n *Nginx) HandleRequest(url, method string) (int, string){
	allowed :=n.checkRateLimiting(url)
	if !allowed {
		return 403, "Not Allowed"
	}
	return n.application.handleRequest(url, method)
}

func (n *Nginx) checkRateLimiting(url string ) bool{
	if n.rateLimiter[url] ==0{
		n.rateLimiter[url] =1
	} else if n.rateLimiter[url] > n.maxAllowedRequest{
		return false
	}
		n.rateLimiter[url]++
		return true


}