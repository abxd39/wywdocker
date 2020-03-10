FROM docker.io/golang
RUN mkdir wywdocker/
COPY . wywdocker/.
WORKDIR ./wywdocker
RUN   export GO111MODULE=auto && go build -o wdc
ADD wdc /wywdocker/
EXPOSE 8080
#CMD go run wd.go

ENTRYPOINT ./wywdocker/wdc