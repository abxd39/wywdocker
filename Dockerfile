FROM glang
COPY . WORDIR/.
RUN go build -v -o v1.0
ENTRYPOINT["./wd.go"]