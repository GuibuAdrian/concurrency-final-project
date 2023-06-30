package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *Config) routes() http.Handler {
	// create router
	mux := chi.NewRouter()

	// set up middleware
	mux.Use(middleware.Recoverer)
	mux.Use(app.SessionLoad)

	// define application routes
	mux.Get("/", app.homePage)

	mux.Get("/login", app.LoginPage)
	mux.Post("/login", app.PostLoginPage)
	mux.Get("/logout", app.Logout)
	mux.Get("/register", app.RegisterPage)
	mux.Post("/register", app.PostRegisterPage)
	mux.Get("/activate", app.ActivateAccount)
	mux.Get("/plans", app.ChooseSubscription)
	mux.Get("/subscribe", app.SubscribeToPlan)

	mux.Get("/test-mail", app.ActivateAccount)

	mux.Mount("members", app.authRouter())
	return mux
	/*
	   mux.Get(`/test-mail2`, {
	       m := Mail {
	           Domain: "localhost",
	           Host: "localhost",
	           Port: 1025,
	           Encryptation: "none",
	           FromAddress: "info@mycompany.com",
	           FromName: "info",
	           ErrorChan: make(chan error),
	       }

	       msg := Message {
	           To: "me@here.com",
	           Subject: "Test mail",
	           Data: "Hello, world!",
	       }

	       m.sendMail(msg, make(chan error))
	   })*/
}

func (app *Config) authRouter() http.Handler {
	mux := chi.NewRouter()
	mux.Use(app.Auth)

	mux.Get("/plans", app.ChooseSubscription)
	mux.Get("/subscribe", app.SubscribeToPlan)

	return mux
}
