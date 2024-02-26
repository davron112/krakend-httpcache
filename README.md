Krakend HTTP Cache
====

A cached http client for the [KrakenD](github.com/davron112/krakend) framework

## Using it

This package exposes two simple factories capable to create a instances of the `proxy.HTTPClientFactory` and the `proxy.BackendFactory` interfaces, respectively, embedding an in-memory-cached http client using the package [github.com/gregjones/httpcache](https://github.com/gregjones/httpcache). The client will cache the responses honoring the defined Cache HTTP header.

	import 	(
		"context"
		"net/http"
		"github.com/davron112/lura/config"
		"github.com/davron112/lura/proxy"
		"github.com/davron112/krakend-httpcache"
	)

	requestExecutorFactory := func(cfg *config.Backend) proxy.HTTPRequestExecutor {
		clientFactory := httpcache.NewHTTPClient(cfg)
		return func(ctx context.Context, req *http.Request) (*http.Response, error) {
			return clientFactory(ctx).Do(req.WithContext(ctx))
		}
	}

You can create your own proxy.HTTPRequestExecutor and inject it into your BackendFactory
