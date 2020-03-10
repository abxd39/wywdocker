FROM glang
RUN mkdir wywdocker/
COPY . wywdocker/.
RUN  go mod && go build -v -o v1.0
EXPOSE 8080
#CMD go run wd.go
WORKDIR wywdocker
ENTRYPOINT wywdocker