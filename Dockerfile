FROM docker.io/golang
RUN mkdir wywdocker/
COPY . wywdocker/.
ADD wywdocker /wywdocker/
WORKDIR ./wywdocker
RUN   export GO111MODULE=auto
EXPOSE 8080
#CMD go run wd.go

ENTRYPOINT ./wywdocker/wywdocker