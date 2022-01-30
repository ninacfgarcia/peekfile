## Run

from top-level repo directory, pass an absolute path:
`./peekfile.sh ${ABSOLUTE_PATH_TO_BROWSE}`
`./peekfile.sh ${PWD}/${RELATIVE_PATH_TO_BROWSE}`

With an http tool like curl, call the API.
API Handles valid paths relative to the $PATH_TO_BROWSE, depending on what's on your machine's filesystem.

## Stop

`docker stop peek-app`

docker stop will stop and also remove the container

### List directory

`curl -s http://localhost:8080/<PATH>`

### View file contents

`curl -s http://localhost:8080/<PATH.ext>`



## Security

You can scan the built container for vulnerabilities if you are logged in to Docker.

```
docker build . -t simple_server --rm
docker login
## enter credentials or use CLI token
docker scan simple_server

# Package manager:   apk
# Project name:      docker-image|simple_server
# Docker image:      simple_server
# Platform:          linux/arm64
```
  

## Prerequisites

Install dependencies: (docker)[https://hub.docker.com/editions/community/docker-ce-desktop-mac] 

---

## Reference

- https://vsupalov.com/docker-arg-env-variable-guide/
- https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/mocking
- https://www.practical-go-lessons.com/chap-26-basic-http-server#anatomy-of-an-http-response
- https://github.com/stretchr/testify#mock-package
