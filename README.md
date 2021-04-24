# Micromock (umock)

A simple mocking service for application development.

## Getting started

You can run Micromock by building Micromock into executable file and execute it in command-line in your machine or use Dockerized Micromock.

### Command line Micromock

```
umock start --conf=<yaml-configuration-file>
```

### Dockerized Micromock

You can start using Micromock by using [Micromock docker image](https://hub.docker.com/repository/docker/chonla/umock).

```
docker run -p <published-port> -v <volume-to-be-mount> chonla/umock:latest ./umock start --conf=<yaml-configuration-file>
```

#### Example

```
docker run -p 8000:8000 -v /my-project/conf.yml:/app/conf.yml chonla/umock:latest ./umock start --conf=/app/conf.yml
```

## Configuration

Configuration is in YAML format. Here is the template and example values.

```
server:
  host: <listening-host-address> # e.g., 0.0.0.0
  port: <listening-port> # 8000
routes:
  - name: <name-of-route> # can be omitted
    method: <request-method> # get, post, delete, put, patch, ...
    path: <request-path> # e.g., /. request path can contain parameters with ":parameter-name" format
    when:
      content_type: <request-content-type> # can be omitted, default = application/x-www-form-urlencoded
      param: # use param to match path parameters
        - id=1721
      query: # use query when method is get
        - id=7771
      form: # use form when content type is application/x-www-form-urlencoded
        - id=3331
      json_body: # use json_body when content type is application/json
        - employee.id=123
    then:
      status: <http-response-code> # e.g., 200
      headers: # headers to be sent back with response
        - Content-Type=application/json
      body: | # content if any
        {}
```

## Project

Micromock on [GitHub](https://github.com/chonla/umock)

## License

[MIT](LICENSE)