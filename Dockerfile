FROM golang:1.12.6

WORKDIR $GOPATH/

RUN git clone https://github.com/angelhack2019/user_svc.git

WORKDIR $GOPATH/food_svc

RUN go mod download

RUN go install

ENTRYPOINT ["/go/bin/user_svc"]

EXPOSE 8080