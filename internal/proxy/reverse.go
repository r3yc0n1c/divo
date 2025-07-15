package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/charmbracelet/log"
)

func StartReverseProxy(targetURL string, listenAddr string) error {
	target, err := url.Parse(targetURL)
	if err != nil {
		log.Error("Failed to parse target URL", "error", err)
		return err
	}

	log.SetLevel(log.InfoLevel)

	proxy := httputil.NewSingleHostReverseProxy(target)

	proxy.Director = func(req *http.Request) {
		originalURL := req.URL.String()
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		log.Info("Request", "method", req.Method, "path", originalURL)
	}

	proxy.ModifyResponse = func(resp *http.Response) error {
		log.Info("Response", "status", resp.StatusCode, "url", resp.Request.URL.Path)
		return nil
	}

	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		log.Error("Proxy error", "method", r.Method, "url", r.URL.Path, "error", err)
		http.Error(w, "Proxy Error", http.StatusBadGateway)
	}

	return http.ListenAndServe(listenAddr, proxy)
}
