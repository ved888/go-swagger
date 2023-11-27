// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"fmt"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"go-swagger/dbhelper"
	"go-swagger/jwt"
	"go-swagger/restapi/operations/user"

	"go-swagger/database"

	"go-swagger/models"
	"go-swagger/restapi/operations"
	"go-swagger/restapi/operations/movie"

	"net/http"
)

//go:generate swagger generate server --target ..\..\go-swagger --name Movie --spec ..\swagger.yml --principal interface{}

func configureFlags(api *operations.MovieAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.MovieAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// call database connection function
	db, err := database.DbConnection()
	if err != nil {
		fmt.Println("db error ", err)
	}
	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.MovieGetMovieGetHandler == nil {
		api.MovieGetMovieGetHandler = movie.GetMovieGetHandlerFunc(func(params movie.GetMovieGetParams) middleware.Responder {
			getMovie, err := dbhelper.GetMovieById(db, params.ID)
			if err != nil {
				return movie.NewGetMovieGetDefault(500).WithPayload(&models.Error{Code: 200, Message: swag.String(err.Error())})
			}
			return movie.NewGetMovieGetOK().WithPayload(getMovie)
		})
	}
	//if api.UserPostLoginHandler == nil {
	api.UserPostLoginHandler = user.PostLoginHandlerFunc(func(params user.PostLoginParams) middleware.Responder {
		login, err := dbhelper.GetUserByEmailAndPhone(db, params.Login.UserPhone, params.Login.UserEmail)
		if err != nil {
			return user.NewPostLoginDefault(500).WithPayload(&models.Error{Code: 200, Message: swag.String(err.Error())})
		}
		token, err := jwt.GenerateJWTToken(*login.Phone, *login.Email)
		if err != nil {
			return user.NewPostLoginDefault(500).WithPayload(&models.Error{Code: 200, Message: swag.String(err.Error())})
		}
		return user.NewPostLoginOK().WithPayload(&models.LoginSuccess{Success: true, Token: token})
	})
	//}
	if api.MoviePostMovieCreateHandler == nil {
		api.MoviePostMovieCreateHandler = movie.PostMovieCreateHandlerFunc(func(params movie.PostMovieCreateParams) middleware.Responder {
			createMovie, err := dbhelper.MovieCreate(db, params.Movie)
			if err != nil {
				return movie.NewPostMovieCreateDefault(500).WithPayload(&models.Error{Code: 200, Message: swag.String(err.Error())})
			}
			return movie.NewPostMovieCreateOK().WithPayload(createMovie)
		})
	}
	//if api.UserPostUserCreateHandler == nil {
	api.UserPostUserCreateHandler = user.PostUserCreateHandlerFunc(func(params user.PostUserCreateParams) middleware.Responder {
		createUser, err := dbhelper.CreateUser(db, params.Register)
		if err != nil {
			return user.NewPostUserCreateDefault(500).WithPayload(&models.Error{Code: 200, Message: swag.String(err.Error())})
		}
		return user.NewPostUserCreateOK().WithPayload(createUser)
	})
	//}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
