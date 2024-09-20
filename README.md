# helloworld-webapp

This is a HelloWorld webapp that runs on `PORT` environment variables.

If you set `TCP_CONNECTION_ADDRS` environment variables, You can check TCP connection to a list of addresses.

This is intended for use during first time setup of ECS fargate, etc.

Preparing images for ECR is the biggest pain in my life!

## Endpoints

Returns status 200 for all paths.

## Environment variables

### `PORT` (optional)

The port the server will listen to.

Default is `8080`.

### `TCP_CONNECTION_ADDRS` (optional)

List of TCP connection addresses to check. (comma separated)

Default is ` `(empty) .

For example, `TCP_CONNECTION_ADDRS=mysql.example.com:3306,redis.example.com:6379`


### `TCP_CONNECTION_TIMEOUT` (optional)

Timeout in seconds for TCP connection check.

Default is `5`.


## Running

[compose.yaml](https://github.com/yumafuu/helloworld-webapp/blob/main/compose.yaml) show full example.

```bash
# This will run the webapp on port 3000 without checking TCP connection.
$ docker run -p 3000:3000 -e PORT=3000 yumafuu/helloworld-webapp

# This will run the webapp on port 9000 with checking TCP connection.
$ docker run -p 9000:9000 -e PORT=9000 -e TCP_CONNECTION_ADDRS=mysql.example.com:3306,redis.example.com:6379 yumafuu/helloworld-webapp
```


# Development

```bash
$ docker build -t yumafuu/helloworld-webapp:tagname .
$ docker push yumafuu/helloworld-webapp:tagname
```
