package main

import (
	"fmt"
	"log"
	"net/http/httputil"
	"net/url"
)

func buildTargets(cfg Config) map[string]url.URL {
	targets := make(map[string]url.URL)
	for host, route := range cfg.Routes {
		targets[host] = url.URL{
			Scheme: route.Target.Scheme,
			Host:   route.Target.Host,
			Path:   route.Target.Path,
		}
	}
	return targets
}

func buildProxy(debug bool, targets map[string]url.URL) httputil.ReverseProxy {
	proxy := httputil.ReverseProxy{
		Rewrite: func(r *httputil.ProxyRequest) {
			if debug {
				dump, err := httputil.DumpRequest(r.In, true)
				if err != nil {
					log.Println("Error while trying to dump the request : ", err.Error())
				}
				fmt.Printf("Dump IN : %s\n", dump)
			}
			host := r.In.Host
			target := targets[host]
			r.SetURL(&target)
			r.SetXForwarded()
			if debug {
				dump, err := httputil.DumpRequestOut(r.Out, true)
				if err != nil {
					log.Println("Error while trying to dump the request : ", err.Error())
				}
				fmt.Printf("Dump OUT : %s\n", dump)
			}
		},
	}
	return proxy
}
