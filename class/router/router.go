package router

import (
	"20241112/config"
	"20241112/handler"
	"20241112/lib"
	"20241112/middleware"
	"20241112/repository"
	"20241112/service"
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"html/template"
	"log"
	"net/http"
)

func initTemplate() (*repository.WebPageData, *template.Template) {
	tmpl, err := template.ParseGlob("view/*.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
		return nil, nil
	}

	return &repository.WebPageData{}, tmpl
}

func NewRouter(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()

	logger := lib.InitLog()
	mwLogger := middleware.NewMiddleware(logger)

	repositoryOrder := repository.InitOrderRepo(db, logger)
	serviceOrder := service.InitOrderService(*repositoryOrder)
	handleOrder := handler.InitOrderHandler(*serviceOrder)

	//OrderHandler := InitializeOrderHandler()
	//log.Println(OrderHandler)

	handlePaymentMethod := handler.InitPaymentMethodHandler(*service.InitPaymentMethodService(*repository.InitPaymentMethodRepo(db, logger)), logger)
	handleWebTemplate := handler.InitWebPageHandler(*service.InitWebPageService(*repository.InitWebPageRepo(initTemplate())))

	//r.Use(middleware.BasicAuth)

	//r.Get("/login", handleWebTemplate.Login)
	//r.Post("/login", handleWebTemplate.Authenticate)
	//
	//r.Get("/logout", handleWebTemplate.Logout)
	//
	//r.Get("/dashboard", handleWebTemplate.Dashboard)
	//
	//r.Route("/books", func(r chi.Router) {
	//	r.Get("/", handleWebTemplate.BookIndex)
	//	r.Get("/{id}", handleWebTemplate.BookShow)
	//	r.Get("/create", handleWebTemplate.BookCreate)
	//	r.Get("/{id}/edit", handleWebTemplate.BookEdit)
	//
	//	r.Get("/{id}/discount/create", handleWebTemplate.BookDiscountCreate)
	//})
	//
	r.Route("/api", func(r chi.Router) {
		r.Use(middleware.JsonResponse())
		//r.With(mwLogger.MiddlewareLogger).Post("/store_book", handlerBookHandler.CreateBook)
		r.Post("/orders", handleOrder.Create)
		r.With(mwLogger.MiddlewareLogger).Route("/payment-methods", func(r chi.Router) {
			r.Post("/", handlePaymentMethod.Create)
			r.Get("/", handlePaymentMethod.All)
			r.Get("/{id}", handlePaymentMethod.Get)
			r.Put("/{id}", handlePaymentMethod.Update)
			r.Delete("/{id}", handlePaymentMethod.Delete)
		})
	})

	r.Get("/style.css", handleWebTemplate.Static)

	fileServer := http.FileServer(http.Dir("./uploads/"))
	r.Handle(fmt.Sprintf("/%s/*", config.UploadDir), http.StripPrefix("/uploads", fileServer))

	return r
}
