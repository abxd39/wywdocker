FROM docker.io/golang
RUN mkdir wywdocker/
COPY . wywdocker/.
RUN   export GO111MODULE=on && go build -v -o w0.0.1
EXPOSE 8080
#CMD go run wd.go
WORKDIR wywdocker
ENTRYPOINT w0.0.1