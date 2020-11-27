# Ports Project
Demo project, example of microservice architecture

## How to use it after `git clone ...`
* Just go to `$GOPATH`/src/.../microports
* `docker-compose up`

## Try, only two endpoins:
### Upload:
~~~bash
curl -X POST \
  http://localhost:8080/api/v1/ports/upload \
  -F file=@ports.json
~~~

### Get all uploaded ports:
~~~bash
curl -X GET \
  http://localhost:8080/api/v1/ports/
~~~