package hello

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ONSdigital/dp-frontend-router/config"
	"github.com/ONSdigital/dp-frontend-router/lang"
	"github.com/ONSdigital/dp-frontend-router/resolver"
	"github.com/ONSdigital/go-ns/log"
)

const xRequestIDHeaderParam = "X-Request-Id"

func Handler(babbageProxy http.Handler) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		b, err := resolver.Get("/", req.Header.Get(xRequestIDHeaderParam))
		if err == resolver.ErrUnauthorised {
			log.ErrorR(req, err, log.Data{"unauthorised user": err.Error()})
			babbageProxy.ServeHTTP(w, req)
			return
		} else if err != nil {
			log.ErrorR(req, err, log.Data{"failed to resolve request": err.Error()})
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		rdr := bytes.NewReader(b)

		rendererReq, err := http.NewRequest("POST", config.RendererURL+"/hello", rdr)
		if err != nil {
			log.ErrorR(req, err, nil)
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		// FIXME there's other headers we want
		rendererReq.Header.Set("Accept-Language", string(lang.Get(req)))
		rendererReq.Header.Set("X-Request-Id", req.Header.Get("X-Request-Id"))

		res, err := http.DefaultClient.Do(rendererReq)
		if err != nil {
			log.ErrorR(req, err, nil)
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		defer res.Body.Close()

		if res.StatusCode != 200 {
			err = fmt.Errorf("Handler.handler: unexpected status code: %d", res.StatusCode)
			log.ErrorR(req, err, nil)
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		// FIXME should stream this using a io.Reader etc
		b, err = ioutil.ReadAll(res.Body)
		if err != nil {
			log.ErrorR(req, err, nil)
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		for hdr, v := range res.Header {
			for _, v2 := range v {
				w.Header().Add(hdr, v2)
			}
		}
		w.WriteHeader(res.StatusCode)
		w.Write(b)
	}
}
