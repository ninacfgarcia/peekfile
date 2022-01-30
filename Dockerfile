FROM golang:1.17-alpine
WORKDIR "/app"
EXPOSE 8000
COPY *.go .
CMD ["go", "run", "main.go", "schema.go", "/tree"]
