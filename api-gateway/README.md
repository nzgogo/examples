# API Gateway

An API gateway uses the [Microservice Architecture](https://microservices.io/patterns/microservices.html) pattern to provide a single entry point for all services.
The implementation accepts HTTP requests and dynamically routes to the appropriate service using service discovery.

## Dependencies

- Consul
- NATS

## Run api-gateway
```shell
$ go run api-gateway/main.go
```
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
"hc_script":"gghc",
"hc_deregister_critical_service_after":"2m",
"hc_interval":"1m"
}
```
- "nats_addr". if left empty, default address (nats//:127.0.0.1:4222) is used.
- "consul_addr". if left empty, default address (127.0.0.1:8500) is used.
- "hc_script" field specifies the location of [health check scrpit](https://gitlab.com/gogoexpress/gogoexpress-go-consul-healthcheck-v1).
- "hc_interval", Consul health check interval.




