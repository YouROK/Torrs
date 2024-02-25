package gincert

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/sync/errgroup"
)

type tlsContextKey string

var (
	ctxKey  = tlsContextKey("autls")
	todoCtx = context.WithValue(context.Background(), ctxKey, "done")
)

func run(ctx context.Context, r http.Handler, domain ...string) error {
	var g errgroup.Group

	s2 := &http.Server{
		Handler: r,
	}

	g.Go(func() error {
		return s2.Serve(autocert.NewListener(domain...))
	})

	g.Go(func() error {
		if v := ctx.Value(ctxKey); v != nil {
			return nil
		}

		<-ctx.Done()

		var gShutdown errgroup.Group
		gShutdown.Go(func() error {
			return s2.Shutdown(context.Background())
		})

		return gShutdown.Wait()
	})
	return g.Wait()
}

// Run support 1-line LetsEncrypt HTTPS servers with graceful shutdown
func RunWithContext(ctx context.Context, r http.Handler, domain ...string) error {
	return run(ctx, r, domain...)
}

// Run support 1-line LetsEncrypt HTTPS servers
func Run(r http.Handler, domain ...string) error {
	return run(todoCtx, r, domain...)
}

// RunWithManager support custom autocert manager
func RunWithManager(r http.Handler, m *autocert.Manager) error {
	return RunWithManagerAndTLSConfig(r, m, m.TLSConfig())
}

// RunWithManagerAndTLSConfig support custom autocert manager and tls.Config
func RunWithManagerAndTLSConfig(r http.Handler, m *autocert.Manager, tlsc *tls.Config) error {
	var g errgroup.Group
	if m.Cache == nil {
		var e error
		m.Cache, e = getCacheDir()
		if e != nil {
			log.Println(e)
		}
	}
	defaultTLSConfig := m.TLSConfig()
	tlsc.GetCertificate = defaultTLSConfig.GetCertificate
	tlsc.NextProtos = defaultTLSConfig.NextProtos
	s := &http.Server{
		Addr:      ":https",
		TLSConfig: tlsc,
		Handler:   r,
	}
	g.Go(func() error {
		return http.ListenAndServe(":http", r)
	})
	g.Go(func() error {
		return s.ListenAndServeTLS("", "")
	})
	return g.Wait()
}
