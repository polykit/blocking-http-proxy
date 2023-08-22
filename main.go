package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/elazarl/goproxy"
)

var cidrs []*net.IPNet

func main() {
	verbose := flag.Bool("v", false, "should every proxy request be logged to stdout")
	ranges := flag.String("block", "10.0.0.0/8,127.0.0.0/8,172.16.0.0/12,192.168.0.0/16", "comma separated list of CIDR ranges to block")
	addr := flag.String("listen", ":8080", "proxy listen address")
	flag.Parse()

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = *verbose

	cs := strings.Split(*ranges, ",")
	cidrs = make([]*net.IPNet, len(cs))
	for i, c := range cs {
		_, cidr, err := net.ParseCIDR(strings.TrimSpace(c))
		if err != nil {
			log.Fatalf("Invalid CIDR: %s", c)
		}
		cidrs[i] = cidr
	}

	proxy.OnRequest().DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			ips, err := net.LookupIP(r.URL.Hostname())
			if err != nil {
				return r, nil
			}

			for _, ip := range ips {
				for _, cidr := range cidrs {
					if cidr.Contains(ip) {
						log.Printf("Blocked %s (%s)", ip, r.RequestURI)
						return r, goproxy.NewResponse(r,
							goproxy.ContentTypeText, http.StatusForbidden,
							"Blocked by blocking-http-proxy\n")
					}
				}
			}

			return r, nil
		})

	log.Printf("Listening on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, proxy))
}
