FROM golang:1.17-alpine
WORKDIR "/peekfile"
EXPOSE 8000
CMD ["go", "run", "src/main.go", "src/schema.go", "/tree"]
