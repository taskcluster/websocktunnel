# webhooktunnel
--
[![Task Status](https://github.taskcluster.net/v1/repository/taskcluster/webhooktunnel/master/badge.svg)](https://github.taskcluster.net/v1/repository/taskcluster/webhooktunnel/master/latest)

Webhooktunnel allows processes that lives behind a firewall to expose web-hooks,
without exposing local ports. Instead they connect to webhooktunnel, requesting
a specific JWT covered `tunnelId`, requests to the `tunnelId` is then proxied to
the process over a websocket.

This is useful for exposing livelogs and interactive sessions from workers in
taskcluster. In many ways this is similar to [ngrok](https://ngrok.com/) or
[localtunnel.me](https://localtunnel.github.io/www/), except server connections
are protected by a JWT which covers the `tunnelId`.

## Configuration

The `webhooktunnel` binary takes the following configuration options via.
environment variables:

 * `HOSTNAME`, hostname this server lives under.
 * `ENV`, set to `'production'` to enable logging.
 * `SYSLOG_ADDR`, syslog name if desired.
 * `SECRET_A`, secret that JWTs can be signed with,
 * `SECRET_B`, secret that JWTs can be signed with,
 * `USE_TLS`, whether to use HTTPS or HTTP (requires `TLS_KEY` and `TLS_CERTIFICATE`),
 * `TLS_KEY`, private TLS key,
 * `TLS_CERTIFICATE`, TLS certificate with intermediate certificates,
 * `DOMAIN_HOSTED`, set to non-empty to use domain-hosted mode, see below.
 * `PORT`, port to listen on (defaults to `80` or `443`).

## Hosting Mode (`DOMAIN_HOSTED`)

If the environment variable `DOMAIN_HOSTED` is non-empty, `webhooktunnel` will
run in _domain hosted mode_. This requires that DNS for `*.HOSTNAME` is pointing
to the server running `webhooktunnel`, usually using a CNAME.

In _domain hosted mode_ a `tunnelId` is exposed as `http(s)://<tunnelId>.<HOSTNAME>/<path>`.
If not running in _domain hosted mode_ a `tunnelId` is exposed as
`http(s)://<HOSTNAME>/<tunnelId>/<path>`, notice that in this mode all tunnels
share the _same-origin_ in browsers. For security reason this may not be desirable.


## Hosting Behind Reverse Proxy

Hosting behind a TLS terminating reverse proxy is possible by setting the
environment variable `USE_TLS` to empty string (or undefined), and ensuring that
the reverse proxy sets the header `X-Forwarded-Proto: https` when forwarding
requests. Failure to do this, will result in clients being told to use HTTP.

**Warning**, notice that it is not possible to load balance over between multiple
instances of webhooktunnel. Such feature might be added in the future.

## Usage in taskcluster
In taskcluster we have a service that issues JWT tokens for use with webhooktunnel.
The service requires callers to have a scope and generates random `tunnelId`s,
so that no two callers gets the same `tunnelId`.

Workers exploit this to be able to expose web-hooks that can be referenced from
_reference artifacts_ that redirect to the web-hook. This is pretty solid as
the web hook is never intended to outlive the runtime of the task.
