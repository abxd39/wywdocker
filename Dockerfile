FROM docker.io/golang
RUN mkdir wywdocker/
COPY . wywdocker/.
WORKDIR ./wywdocker
RUN   export GO111MODULE=on && go env
EXPOSE 8080
#CMD go run wd.go

ENTRYPOINT w0.0.1