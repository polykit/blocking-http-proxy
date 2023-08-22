# Blocking HTTP Proxy

This repository contains a simple http-proxy which allows blocking of certain CIDR ranges. It's meant for running inside of a Kubernetes cluster and to prevent applications from accessing internal services.

## Usage

Run `blocking-http-proxy` locally, eg.

```
$ ./blocking-http-proxy -v --block=127.0.0.0/8
2023/08/22 14:05:49 Listening on :8080
```

Try it out:
```
$ curl --proxy http://localhost:8080 -L -v http://localhost
*   Trying 127.0.0.1:8080...
* TCP_NODELAY set
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET http://localhost/ HTTP/1.1
> Host: localhost
> User-Agent: curl/7.68.0
> Accept: */*
> Proxy-Connection: Keep-Alive
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 403 Forbidden
< Content-Type: text/plain
< Date: Tue, 22 Aug 2023 12:14:09 GMT
< Content-Length: 31
<
Blocked by blocking-http-proxy
```

See [yaml](yaml/) directory for an actual deployment in a Kubernetes environment.

## References

https://github.com/elazarl/goproxy
