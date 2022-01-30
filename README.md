Peekfile app reads valid paths on your machine's filesystem .

## Run

from top-level repo directory, pass an absolute path as the root of peekfile's file browser tree:

`./peekfile.sh ${ABSOLUTE_PATH_TO_BROWSE}`

`./peekfile.sh ${PWD}/${RELATIVE_PATH_TO_BROWSE}`

With an http tool like curl, call the API at port :8080.

-------------

## Use
-------------

"data" field wraps valid response data, and "error" wraps the error response message

### List directory

GET `/<PATH>`

[200 OK] data: Array of file entries under the directory at PATH.

Each file entry has this shape:

field        | type   | description
------       | ------ | -----
filename     | String | basename
owner_id     | String | uid
size         | String | size in bytes
permissions  | String | permissions as octal string

### View file contents

GET `/<PATH.ext>`

[200 OK] data: file content of the file at PATH.

field   | type   | description
------- | ------ | ----------
content | String | the text contained in the file

### Errors

error: the error trying to read the PATH.

[404 Not Found] Path does not exist  
[500 Internal Server Error] a formatting or OS operation has failed

-------------

## Example
-------------

You might start the app CLI 
with `./peekfile.sh /home/me`  
and this file structure:
```
/home/me/
+----- file0
\----- foo/
    +----- bar/
    │   \----- file2
    \----- .file1
```

GET the following paths and the JSON will correspond to the below output:

Path       | Does what  | Summary of output
-----------|------------|-----------------
/          | ls         | file0, foo/    
foo        | ls  foo    | bar/, .file1    
foo/bar    | ls  bar    | file2          
foo/.file1 | cat .file1 | "text-content" 

## Stop

`docker stop peek-app`

docker stop will stop and also remove the container

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
  
------

## Prerequisites

Install dependencies: (docker)[https://hub.docker.com/editions/community/docker-ce-desktop-mac] 
