package main

import (
	"context"
	"math/rand"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/ONSdigital/dp-frontend-router/router"

	dphttp "github.com/ONSdigital/dp-net/http"

	"github.com/ONSdigital/dp-api-clients-go/zebedee"
	"github.com/ONSdigital/dp-frontend-router/assets"
	"github.com/ONSdigital/dp-frontend-router/config"
	"github.com/ONSdigital/dp-frontend-router/handlers/analytics"
	"github.com/ONSdigital/dp-frontend-router/middleware/redirects"
	"github.com/ONSdigital/dp-healthcheck/healthcheck"
	"github.com/ONSdigital/log.go/v2/log"
)

var (
	// BuildTime represents the time in which the service was built
	BuildTime string
	// GitCommit represents the commit (SHA-1) hash of the service that is running
	GitCommit string
	// Version represents the version of the service that is running
	Version string
)

func main() {
	log.Namespace = "dp-frontend-router"

	ctx := context.Background()

	cfg, err := config.Get()
	if err != nil {
		log.Fatal(ctx, "unable to retrieve service configuration", err)
		os.Exit(1)
	}

	log.Info(ctx, "got service configuration", log.Data{"config": cfg})

	cookiesControllerURL, err := url.Parse(cfg.CookiesControllerURL)
	if err != nil {
		log.Fatal(ctx, "configuration value is invalid", err, log.Data{"config_name": "CookiesControllerURL", "value": cfg.CookiesControllerURL})
		os.Exit(1)
	}

	datasetControllerURL, err := url.Parse(cfg.DatasetControllerURL)
	if err != nil {
		log.Fatal(ctx, "configuration value is invalid", err, log.Data{"config_name": "DatasetControllerURL", "value": cfg.DatasetControllerURL})
		os.Exit(1)
	}

	filterDatasetControllerURL, err := url.Parse(cfg.FilterDatasetControllerURL)
	if err != nil {
		log.Fatal(ctx, "configuration value is invalid", err, log.Data{"config_name": "FilterDatasetControllerURL", "value": cfg.FilterDatasetControllerURL})
		os.Exit(1)
	}

	geographyControllerURL, err := url.Parse(cfg.GeographyControllerURL)
	if err != nil {
		log.Fatal(ctx, "configuration value is invalid", err, log.Data{"config_name": "GeographyControllerURL", "value": cfg.GeographyControllerURL})
		os.Exit(1)
	}

	homepageControllerURL, err := url.Parse(cfg.HomepageControllerURL)
	if err != nil {
		log.Fatal(ctx, "configuration value is invalid", err, log.Data{"config_name": "HomepageControllerURL", "value": cfg.HomepageControllerURL})
		os.Exit(1)
	}

	searchControllerURL, err := url.Parse(cfg.SearchControllerURL)
	if err != nil {
		log.Fatal(ctx, "configuration value is invalid", err, log.Data{"config_name": "SearchControllerURL", "value": cfg.SearchControllerURL})
		os.Exit(1)
	}

	babbageURL, err := url.Parse(cfg.BabbageURL)
	if err != nil {
		log.Fatal(ctx, "configuration value is invalid", err, log.Data{"config_name": "BabbageURL", "value": cfg.BabbageURL})
		os.Exit(1)
	}

	downloaderURL, err := url.Parse(cfg.DownloaderURL)
	if err != nil {
		log.Fatal(ctx, "configuration value is invalid", err, log.Data{"config_name": "DownloaderURL", "value": cfg.DownloaderURL})
		os.Exit(1)
	}

	feedbackControllerURL, err := url.Parse(cfg.FeedbackControllerURL)
	if err != nil {
		log.Fatal(ctx, "configuration value is invalid", err, log.Data{"config_name": "FeedbackControllerURL", "value": cfg.FeedbackControllerURL})
		os.Exit(1)
	}

	areaProfileControllerURL, err := url.Parse(cfg.AreaProfilesControllerURL)
	if err != nil {
		log.Fatal(ctx, "configuration value is invalid", err, log.Data{"config_name": "AreaProfileControllerURL", "value": cfg.AreaProfilesControllerURL})
		os.Exit(1)
	}

	enableSearchABTest := config.IsEnableSearchABTest(*cfg)

	redirects.Init(assets.Asset)

	// create ZebedeeClient proxying calls through the API Router
	hcClienter := dphttp.NewClient()
	hcClienter.SetMaxRetries(cfg.ZebedeeRequestMaximumRetries)
	hcClienter.SetTimeout(cfg.ZebedeeRequestMaximumTimeoutSeconds)

	zebedeeClient := zebedee.NewClientWithClienter(cfg.APIRouterURL, hcClienter)

	// Healthcheck API
	versionInfo, err := healthcheck.NewVersionInfo(BuildTime, GitCommit, Version)
	if err != nil {
		log.Fatal(ctx, "Failed to obtain VersionInfo for healthcheck", err)
		os.Exit(1)
	}
	hc := healthcheck.New(versionInfo, cfg.HealthckeckCriticalTimeout, cfg.HealthckeckInterval)
	if err = hc.AddCheck("API router", zebedeeClient.Checker); err != nil {
		log.Fatal(ctx, "Failed to add api router checker to healthcheck", err)
		os.Exit(1)
	}

	analyticsHandler, err := analytics.NewSearchHandler(ctx, cfg.SQSAnalyticsURL, cfg.RedirectSecret)
	if err != nil {
		log.Fatal(ctx, "error creating search analytics handler", err)
		os.Exit(1)
	}

	downloadHandler := createReverseProxy("download", downloaderURL)
	cookieHandler := createReverseProxy("cookies", cookiesControllerURL)
	datasetHandler := createReverseProxy("datasets", datasetControllerURL)
	filterHandler := createReverseProxy("filters", filterDatasetControllerURL)
	feedbackHandler := createReverseProxy("feedback", feedbackControllerURL)
	searchHandler := createReverseProxy("search", searchControllerURL)
	homepageHandler := createReverseProxy("homepage", homepageControllerURL)
	babbageHandler := createReverseProxy("babbage", babbageURL)
	areaProfileHandler := createReverseProxy("areas", areaProfileControllerURL)
	var geographyHandler http.Handler
	if cfg.AreaProfilesRoutesEnabled {
		geographyHandler = redirects.DynamicRedirectHandler("/geography", "/areas")
	} else {
		geographyHandler = createReverseProxy("geography", geographyControllerURL)
	}

	routerConfig := router.Config{
		AnalyticsHandler:       analyticsHandler,
		AreaProfileEnabled:     cfg.AreaProfilesRoutesEnabled,
		AreaProfileHandler:     areaProfileHandler,
		DownloadHandler:        downloadHandler,
		CookieHandler:          cookieHandler,
		DatasetHandler:         datasetHandler,
		HealthCheckHandler:     hc.Handler,
		FilterHandler:          filterHandler,
		FeedbackHandler:        feedbackHandler,
		GeographyEnabled:       cfg.GeographyEnabled,
		GeographyHandler:       geographyHandler,
		SearchRoutesEnabled:    cfg.SearchRoutesEnabled,
		SearchHandler:          searchHandler,
		EnableSearchABTest:     enableSearchABTest,
		SearchABTestPercentage: cfg.SearchABTestPercentage,
		SiteDomain:             cfg.SiteDomain,
		HomepageHandler:        homepageHandler,
		BabbageHandler:         babbageHandler,
		ZebedeeClient:          zebedeeClient,
		ContentTypeByteLimit:   cfg.ContentTypeByteLimit,
	}

	httpHandler := router.New(routerConfig)

	log.Info(ctx, "Starting server", log.Data{"config": cfg})

	s := &http.Server{
		Addr:         cfg.BindAddr,
		Handler:      httpHandler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Start health check
	hc.Start(ctx)

	// Start server
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(ctx, "error starting server", err)
		hc.Stop()
		os.Exit(2)
	}
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

// abHandler ... percentA is the percentage of request that handler 'a' is used
//
// FIXME this isn't used anymore, it could be removed, but seems like it might be useful?
func abHandler(a, b http.Handler, percentA int) http.Handler {
	if percentA == 0 {
		log.Info(context.Background(), "abHandler decision", log.Data{"percentA": percentA, "destination": "b"})
		return b
	} else if percentA == 100 {
		log.Info(context.Background(), "abHandler decision", log.Data{"percentA": percentA, "destination": "a"})
		return a
	}

	if percentA < 0 || percentA > 100 {
		panic("Percent 'a' must be between 0 and 100")
	}
	rand.Seed(time.Now().UnixNano())

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Detect cookie
		cookie, _ := req.Cookie("homepage-version")

	RETRY:
		if cookie == nil {
			var cookieValue string
			sel := rand.Intn(100)
			if sel < percentA {
				cookieValue = "A"
			} else {
				cookieValue = "B"
			}

			log.Info(req.Context(), "abHandler decision", log.Data{"sel": sel, "handler": cookieValue})

			expiration := time.Now().Add(365 * 24 * time.Hour)
			cookie = &http.Cookie{Name: "homepage-version", Value: cookieValue, Expires: expiration}
			http.SetCookie(w, cookie)
		}

		// Use cookie value to direct to a or b handler
		switch cookie.Value {
		case "A":
			log.Info(req.Context(), "abHandler decision", log.Data{"cookie": "A", "destination": "a"})
			a.ServeHTTP(w, req)
		case "B":
			log.Info(req.Context(), "abHandler decision", log.Data{"cookie": "B", "destination": "b"})
			b.ServeHTTP(w, req)
		default:
			log.Info(req.Context(), "abHandler invalid cookie value, reselecting")
			cookie = nil
			goto RETRY
		}
	})
}

func createReverseProxy(proxyName string, proxyURL *url.URL) http.Handler {
	proxy := httputil.NewSingleHostReverseProxy(proxyURL)
	director := proxy.Director
	proxy.Transport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   5 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	proxy.Director = func(req *http.Request) {
		log.Info(req.Context(), "proxying request", log.HTTP(req, 0, 0, nil, nil), log.Data{
			"destination": proxyURL,
			"proxy_name":  proxyName,
		})
		director(req)
	}
	return proxy
}
