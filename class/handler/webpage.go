package handler

import (
	"20241112/service"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type WebPageHandler struct {
	WebPageService service.WebPageService
}

func InitWebPageHandler(webPageService service.WebPageService) WebPageHandler {
	return WebPageHandler{WebPageService: webPageService}
}

//func (handle *WebPageHandler) Login(w http.ResponseWriter, r *http.Request) {
//	handle.WebPageService.Render(w, "login.html", "Login")
//}
//
//func (handle *WebPageHandler) Authenticate(w http.ResponseWriter, r *http.Request) {
//
//	//handle.WebPageService.Render(w, "login.html", "Login")
//}
//
//func (handle *WebPageHandler) Logout(w http.ResponseWriter, r *http.Request) {
//	handle.WebPageService.Render(w, "logout.html", "Logout")
//}
//
//func (handle *WebPageHandler) Dashboard(w http.ResponseWriter, r *http.Request) {
//	handle.WebPageService.Render(w, "dashboard.html", "Dashboard")
//}
//
//func (handle *WebPageHandler) BookCreate(w http.ResponseWriter, r *http.Request) {
//	handle.WebPageService.Render(w, "add-book.html", "Add a Book")
//}
//
//func (handle *WebPageHandler) BookIndex(w http.ResponseWriter, r *http.Request) {
//	handle.WebPageService.Render(w, "book-list.html", "Books")
//}
//
//func (handle *WebPageHandler) BookEdit(w http.ResponseWriter, r *http.Request) {
//	handle.WebPageService.Render(w, "edit-book.html", "Edit Book")
//}
//
//func (handle *WebPageHandler) BookShow(w http.ResponseWriter, r *http.Request) {
//	handle.WebPageService.Render(w, "view-book.html", "Book")
//}
//
//func (handle *WebPageHandler) BookDiscountCreate(w http.ResponseWriter, r *http.Request) {
//	handle.WebPageService.Render(w, "discount-book.html", "Set Discount")
//}
//
//func (handle *WebPageHandler) OrderIndex(w http.ResponseWriter, r *http.Request) {
//	handle.WebPageService.Render(w, "order-list.html", "Orders")
//}

func (handle *WebPageHandler) Static(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	var data []byte
	log.Println(path)
	w.Header().Set("Content-Type", "text/css")
	if strings.HasSuffix(path, "js") {
		w.Header().Set("Content-Type", "application/javascript")
	}
	data, _ = os.ReadFile(fmt.Sprintf("view/%s", path))
	w.Write(data)
}
