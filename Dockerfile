FROM docker.io/golang
RUN mkdir wywdocker/
COPY . wywdocker/.
WORKDIR ./wywdocker
RUN   export GO111MODULE=on && go env && go build -v -o v0.0.1
EXPOSE 8080
#CMD go run wd.go

ENTRYPOINT ./wywdocker/v0.0.1