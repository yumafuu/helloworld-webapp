# helloworld-webapp

This is a HelloWorld webapp that runs on `PORT` environment variables.

## Endpoints

All return status 200.

- `/`
- `/health`

## Environment variables

- `PORT`: The port the server will listen to. Default is 8080.

## Running

```bash
$ docker run -p 3000:3000 -e PORT=3000 yumafuu/helloworld-webapp
```
