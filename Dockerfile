FROM golang as base

WORKDIR /url-shortener

COPY . .

RUN go mod download

RUN go build -o /bin/main .

ENTRYPOINT [ "/bin/main" ]





