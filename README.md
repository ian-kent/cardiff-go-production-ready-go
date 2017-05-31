Production Ready Go
===================

Main focus will be on HTTP servers, primarily APIs, but is also
applicable to lots of client or command line apps.

## The 'goes without saying' stuff

- Race detector!
- Vendoring

## Other stuff

- Healthchecks
  - run healthchecks on a timer
  - expose result only
- Logging/telemetry
  - pull - e.g. expvar
  - push - e.g. statsd
  - structured logging to stdout/stderr 
    - same as push, but more 12-factor
    - means capturing the app output
- Timeouts
  - mitigated by proxy, e.g. nginx
  - good DDOS vector
  - https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
  - https://blog.cloudflare.com/exposing-go-on-the-internet/
- OS signals
  - graceful shutdown
  - rolling upgrades
- Cancellation
- Tracing
- Never start a goroutine without knowing how it will stop - Dave Cheney
