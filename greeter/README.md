# Greeter

An example Greeter application

## Contents

- **[Dependencies](#dependencies)**
- **[Run Greeter](#run-greeter)**
- **[API-Gateway](#api-gateway)** - examples of RESTful API

## Dependencies

- Consul
- NATS
- Health Check Script

### Consul

```
$ brew install consul
$ consul agent  -bootstrap-expect 1 -server -data-dir /tmp/datadir -enable-script-checks=true -ui
```

### Nats
```
$ go get github.com/nats
$ gnatsd
```

### Health Check Script
One of the primary roles of the Consul agent is management of system-level and application-level health checks. There are several different kinds of checks, see Checks Definition. The kind of check used in micro is [Script + Interval](https://www.consul.io/docs/agent/checks.html).
[Check Scripts](https://www.consul.io/docs/agent/checks.html#check-scripts)
To get started quickly, we define a script that always return with passing:
```shell
$ vim /etc/gogo/healthcheck.sh
```
```shell
#!/bin/sh
# A health check script that always return 0 (Passing).
exit 0
```
```shell
$ chmod +x /etc/consul/healthcheck.sh
```
```shell

## Run Greeter
### Service Configurations
An example of servce configuration file:
```json
{
"nats_addr": "",
"consul_addr":"",
"hc_load_critical_threshold":"3",
"hc_load_warning_threshold":"2",
"hc_memory_critical_threshold":"5",
"hc_memory_warning_threshold":"10",
"hc_script":"/etc/gogo/healthcheck.sh",
"hc_deregister_critical_service_after":"2m",
"hc_interval":"1m"
}
```
- "nats_addr". if left empty, default address (nats//:127.0.0.1:4222) is used.
- "consul_addr". if left empty, default address (127.0.0.1:8500) is used.
- "hc_script" field specifies the location of [health check scrpit](https://gitlab.com/gogoexpress/gogoexpress-go-consul-healthcheck-v1).
- "hc_interval", Consul health check interval.

### Start greeter
```shell
go run srv/server.go
```

## API-Gateway

An API gateway uses the [Microservice Architecture](https://microservices.io/patterns/microservices.html) pattern to provide a single entry point for all services.
The implementation accepts HTTP requests and dynamically routes to the appropriate service using service discovery.

### Run api-gateway
```shell
$ go run api-gateway/main.go
```

## HTTP request
```shell
$ curl http://localhost:8080/gogox/v1/greeter/hello
```
or
Browse http://localhost:8080/v1/core/greeter/hello
