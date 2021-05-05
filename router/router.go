package router

import (
	"github.com/ONSdigital/dp-frontend-router/middleware/allRoutes"
	"github.com/ONSdigital/dp-frontend-router/middleware/redirects"
	"github.com/ONSdigital/dp-frontend-router/middleware/serverError"
	dprequest "github.com/ONSdigital/dp-net/request"
	"github.com/ONSdigital/log.go/log"
	"github.com/gorilla/mux"
	"github.com/gorilla/pat"
	"github.com/justinas/alice"
	"net/http"
	"path/filepath"
	"strings"
)

//go:generate moq -out routertest/handler.go -pkg routertest . Handler
type Handler http.Handler

type Config struct {
	HealthCheckHandler   func(w http.ResponseWriter, req *http.Request)
	AnalyticsHandler     http.Handler
	DownloadHandler      http.Handler
	DatasetHandler       http.Handler
	CookieHandler        http.Handler
	FilterHandler        http.Handler
	FeedbackHandler      http.Handler
	ContentTypeByteLimit int
	ZebedeeClient        allRoutes.ZebedeeClient
	GeographyEnabled     bool
	GeographyHandler     http.Handler
	SearchRoutesEnabled  bool
	SearchHandler        http.Handler
	HomepageHandler      http.Handler
	BabbageHandler       http.Handler
}

func New(cfg Config) http.Handler {

	router := pat.New()

	middleware := []alice.Constructor{
		dprequest.HandlerRequestID(16),
		log.Middleware,
		securityHandler,
		healthcheckHandler(cfg.HealthCheckHandler),
		serverError.Handler,
		redirects.Handler,
	}

	alice := alice.New(middleware...).Then(router)

	router.Handle("/", cfg.HomepageHandler)

	router.Handle("/redir/{data:.*}", cfg.AnalyticsHandler)
	router.Handle("/download/{uri:.*}", cfg.DownloadHandler)
	router.Handle("/cookies{uri:.*}", cfg.CookieHandler)
	router.Handle("/datasets/{uri:.*}", cfg.DatasetHandler)
	router.Handle("/filters/{uri:.*}", cfg.FilterHandler)
	router.Handle("/filter-outputs/{uri:.*}", cfg.FilterHandler)
	router.Handle("/feedback{uri:.*}", cfg.FeedbackHandler)

	if cfg.GeographyEnabled {
		router.Handle("/geography{uri:.*}", cfg.GeographyHandler)
	}

	if cfg.SearchRoutesEnabled {
		router.Handle("/search", cfg.SearchHandler)
	}

	// if the request is for a file go directly to babbage instead of using the allRoutesMiddleware
	router.MatcherFunc(hasFileExtMatcher).Handler(cfg.BabbageHandler)

	// If it is a known babbage endpoint go directly to babbage instead of using the allRoutesMiddleware
	router.MatcherFunc(isKnownBabbageEndpointMatcher).Handler(cfg.BabbageHandler)

	// all other requests go through the allRoutesMiddleware to check the page type first
	allRoutesMiddleware := allRoutes.Handler(map[string]http.Handler{
		"dataset_landing_page": cfg.DatasetHandler,
	}, cfg.ZebedeeClient, cfg.ContentTypeByteLimit)

	babbageRouter := router.PathPrefix("/").Subrouter()
	babbageRouter.Use(allRoutesMiddleware)
	babbageRouter.PathPrefix("/").Handler(cfg.BabbageHandler)

	return alice
}

// securityHandler ...
func securityHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/embed" && !strings.HasPrefix(req.URL.Path, "/visualisations/") {
			w.Header().Set("X-Frame-Options", "SAMEORIGIN")
		}
		h.ServeHTTP(w, req)
	})
}

// healthcheckHandler uses the provided handler for /health endpoint, and serves any other traffic to the next handler in chain
func healthcheckHandler(hc func(w http.ResponseWriter, req *http.Request)) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			if req.URL.Path == "/health" {
				hc(w, req)
				return
			}
			h.ServeHTTP(w, req)
		})
	}
}

var knownBabbageEndpoints = []string{
	"/chartconfig",
	"/chartimage",
	"/generator",
	"/visualisations/",
	"/timeseriestool",
	"/search",
	"/file",
	"/resource",
	"/ons/",
	"/file",
	"/calendar",
	"/chart",
	"/embed",
	"/export",
	"/hash",
}

// IsKnownBabbageEndpoint returns true if the given path matches a known babbage endpoint
func IsKnownBabbageEndpoint(path string) bool {
	for _, endpoint := range knownBabbageEndpoints {
		if strings.HasPrefix(path, endpoint) {
			return true
		}
	}
	return false
}

// hasFileExtMatcher is a mux MatcherFunc implementation, allowing routes to be matched on having a file extension
func isKnownBabbageEndpointMatcher(request *http.Request, match *mux.RouteMatch) bool {
	return IsKnownBabbageEndpoint(request.URL.Path)
}

// HasFileExt returns true if the given path has a file extension
func HasFileExt(path string) bool {
	return len(filepath.Ext(path)) > 0
}

// hasFileExtMatcher is a mux MatcherFunc implementation, allowing routes to be matched on having a file extension
func hasFileExtMatcher(request *http.Request, match *mux.RouteMatch) bool {
	return HasFileExt(request.URL.Path)
}
