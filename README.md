# helloworld-webapp

This is a HelloWorld webapp that runs on `PORT` environment variables.

## Endpoints

All path return status 200.

## Environment variables

- `PORT` (optional)

The port the server will listen to.

Default is `8080`.

- `TCP_CONNECTION_ADDRS` (optional)

List of TCP connection addresses to check.

Default is ` `(empty) .

- `TCP_CONNECTION_TIMEOUT` (optional)

Timeout in seconds for TCP connection check.

Default is `5`.


## Running

```bash
$ docker run -p 3000:3000 -e PORT=3000 yumafuu/helloworld-webapp
```


# Development

```bash
$ docker build -t yumafuu/helloworld-webapp:tagname .
$ docker push yumafuu/helloworld-webapp:tagname
```
