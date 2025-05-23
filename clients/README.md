# Clients

Clients service provides an HTTP API for managing platform resources: clients and channels.
Through this API clients are able to do the following actions:

- provision new clients
- create new channels
- "connect" clients into the channels

For an in-depth explanation of the aforementioned scenarios, as well as thorough
understanding of SuperMQ, please check out the [official documentation][doc].

## Configuration

The service is configured using the environment variables presented in the
following table. Note that any unset variables will be replaced with their
default values.

| Variable                       | Description                                                             | Default                        |
| ------------------------------ | ----------------------------------------------------------------------- | ------------------------------ |
| SMQ_CLIENTS_LOG_LEVEL          | Log level for Clients (debug, info, warn, error)                        | info                           |
| SMQ_CLIENTS_HTTP_HOST          | Clients service HTTP host                                               | localhost                      |
| SMQ_CLIENTS_HTTP_PORT          | Clients service HTTP port                                               | 9000                           |
| SMQ_CLIENTS_SERVER_CERT        | Path to the PEM encoded server certificate file                         | ""                             |
| SMQ_CLIENTS_SERVER_KEY         | Path to the PEM encoded server key file                                 | ""                             |
| SMQ_CLIENTS_GRPC_HOST          | Clients service gRPC host                                               | localhost                      |
| SMQ_CLIENTS_GRPC_PORT          | Clients service gRPC port                                               | 7000                           |
| SMQ_CLIENTS_GRPC_SERVER_CERT   | Path to the PEM encoded server certificate file                         | ""                             |
| SMQ_CLIENTS_GRPC_SERVER_KEY    | Path to the PEM encoded server key file                                 | ""                             |
| SMQ_CLIENTS_DB_HOST            | Database host address                                                   | localhost                      |
| SMQ_CLIENTS_DB_PORT            | Database host port                                                      | 5432                           |
| SMQ_CLIENTS_DB_USER            | Database user                                                           | supermq                        |
| SMQ_CLIENTS_DB_PASS            | Database password                                                       | supermq                        |
| SMQ_CLIENTS_DB_NAME            | Name of the database used by the service                                | clients                        |
| SMQ_CLIENTS_DB_SSL_MODE        | Database connection SSL mode (disable, require, verify-ca, verify-full) | disable                        |
| SMQ_CLIENTS_DB_SSL_CERT        | Path to the PEM encoded certificate file                                | ""                             |
| SMQ_CLIENTS_DB_SSL_KEY         | Path to the PEM encoded key file                                        | ""                             |
| SMQ_CLIENTS_DB_SSL_ROOT_CERT   | Path to the PEM encoded root certificate file                           | ""                             |
| SMQ_CLIENTS_CACHE_URL          | Cache database URL                                                      | <redis://localhost:6379/0>     |
| SMQ_CLIENTS_CACHE_KEY_DURATION | Cache key duration in seconds                                           | 3600                           |
| SMQ_CLIENTS_ES_URL             | Event store URL                                                         | <localhost:6379>               |
| SMQ_CLIENTS_ES_PASS            | Event store password                                                    | ""                             |
| SMQ_CLIENTS_ES_DB              | Event store instance name                                               | 0                              |
| SMQ_CLIENTS_STANDALONE_ID      | User ID for standalone mode (no gRPC communication with Auth)           | ""                             |
| SMQ_CLIENTS_STANDALONE_TOKEN   | User token for standalone mode that should be passed in auth header     | ""                             |
| SMQ_JAEGER_URL                 | Jaeger server URL                                                       | <http://jaeger:4318/v1/traces> |
| SMQ_AUTH_GRPC_URL              | Auth service gRPC URL                                                   | localhost:7001                 |
| SMQ_AUTH_GRPC_TIMEOUT          | Auth service gRPC request timeout in seconds                            | 1s                             |
| SMQ_AUTH_GRPC_CLIENT_TLS       | Enable TLS for gRPC client                                              | false                          |
| SMQ_AUTH_GRPC_CA_CERT          | Path to the CA certificate file                                         | ""                             |
| SMQ_SEND_TELEMETRY             | Send telemetry to supermq call home server.                             | true                           |
| Clients_INSTANCE_ID            | Clients instance ID                                                     | ""                             |

**Note** that if you want `clients` service to have only one user locally, you should use `CLIENTS_STANDALONE` env vars. By specifying these, you don't need `auth` service in your deployment for users' authorization.

## Deployment

The service itself is distributed as Docker container. Check the [`clients `](https://github.com/absmach/supermq/blob/main/docker/docker-compose.yaml#L167-L194) service section in
docker-compose file to see how service is deployed.

To start the service outside of the container, execute the following shell script:

```bash
# download the latest version of the service
git clone https://github.com/absmach/supermq

cd supermq

# compile the clients
make clients

# copy binary to bin
make install

# set the environment variables and run the service
Clients_LOG_LEVEL=[Clients log level] \
Clients_STANDALONE_ID=[User ID for standalone mode (no gRPC communication with auth)] \
Clients_STANDALONE_TOKEN=[User token for standalone mode that should be passed in auth header] \
Clients_CACHE_KEY_DURATION=[Cache key duration in seconds] \
Clients_HTTP_HOST=[Clients service HTTP host] \
Clients_HTTP_PORT=[Clients service HTTP port] \
Clients_HTTP_SERVER_CERT=[Path to server certificate in pem format] \
Clients_HTTP_SERVER_KEY=[Path to server key in pem format] \
Clients_AUTH_GRPC_HOST=[Clients service gRPC host] \
Clients_AUTH_GRPC_PORT=[Clients service gRPC port] \
Clients_AUTH_GRPC_SERVER_CERT=[Path to server certificate in pem format] \
Clients_AUTH_GRPC_SERVER_KEY=[Path to server key in pem format] \
Clients_DB_HOST=[Database host address] \
Clients_DB_PORT=[Database host port] \
Clients_DB_USER=[Database user] \
Clients_DB_PASS=[Database password] \
Clients_DB_NAME=[Name of the database used by the service] \
Clients_DB_SSL_MODE=[SSL mode to connect to the database with] \
Clients_DB_SSL_CERT=[Path to the PEM encoded certificate file] \
Clients_DB_SSL_KEY=[Path to the PEM encoded key file] \
Clients_DB_SSL_ROOT_CERT=[Path to the PEM encoded root certificate file] \
Clients_CACHE_URL=[Cache database URL] \
Clients_ES_URL=[Event store URL] \
Clients_ES_PASS=[Event store password] \
Clients_ES_DB=[Event store instance name] \
SMQ_AUTH_GRPC_URL=[Auth service gRPC URL] \
SMQ_AUTH_GRPC_TIMEOUT=[Auth service gRPC request timeout in seconds] \
SMQ_AUTH_GRPC_CLIENT_TLS=[Enable TLS for gRPC client] \
SMQ_AUTH_GRPC_CA_CERT=[Path to trusted CA certificate file] \
SMQ_JAEGER_URL=[Jaeger server URL] \
SMQ_SEND_TELEMETRY=[Send telemetry to supermq call home server] \
Clients_INSTANCE_ID=[Clients instance ID] \
$GOBIN/supermq-clients
```

Setting `Clients_CA_CERTS` expects a file in PEM format of trusted CAs. This will enable TLS against the Auth gRPC endpoint trusting only those CAs that are provided.

In constrained environments, sometimes it makes sense to run Clients service as a standalone to reduce network traffic and simplify deployment. This means that Clients service
operates only using a single user and is able to authorize it without gRPC communication with Auth service.
To run service in a standalone mode, set `Clients_STANDALONE_EMAIL` and `Clients_STANDALONE_TOKEN`.

## Usage

For more information about service capabilities and its usage, please check out
the [API documentation](https://docs.api.supermq.abstractmachines.fr/?urls.primaryName=clients-openapi.yaml).

[doc]: https://docs.supermq.abstractmachines.fr
