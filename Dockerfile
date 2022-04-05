FROM golang

RUN mkdir '/blockhain'

WORKDIR "/blockhain"

COPY . .

RUN go mod download

RUN go build .

CMD ["./blockhain"]