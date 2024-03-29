FROM golang:latest 

WORKDIR /app


RUN export GO111MODULE=on

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./

RUN cd /app/ && git clone https://github.com/kaeppen/disys-mandatory2.git 
RUN cd /app/disys-mandatory2 && go build -o main .

ENV PORT $port

CMD ["/app/disys-mandatory2/main"]
