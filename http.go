// Package httpcache introduces an in-memory-cached http client into the KrakenD stack
package httpcache

import (
	"context"
	"net/http"

	"github.com/davron112/lura/config"
	"github.com/davron112/lura/proxy"
	"github.com/davron112/lura/transport/http/client"
	"github.com/gregjones/httpcache"
)

// Namespace is the key to use to store and access the custom config data
const Namespace = "github.com/davron112/krakend-httpcache"

var (
	memTransport = httpcache.NewMemoryCacheTransport()
	memClient    = http.Client{Transport: memTransport}
)

// NewHTTPClient creates a HTTPClientFactory using an in-memory-cached http client
func NewHTTPClient(cfg *config.Backend) client.HTTPClientFactory {
	_, ok := cfg.ExtraConfig[Namespace]
	if !ok {
		return client.NewHTTPClient
	}
	return func(_ context.Context) *http.Client {
		return &memClient
	}
}

// BackendFactory returns a proxy.BackendFactory that creates backend proxies using
// an in-memory-cached http client
func BackendFactory(cfg *config.Backend) proxy.BackendFactory {
	return proxy.CustomHTTPProxyFactory(NewHTTPClient(cfg))
}
