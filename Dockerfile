FROM	golang:1.13-alpine3.10
RUN     apk update && apk add make gcc g++ linux-headers git perl musl-dev
RUN		git clone https://github.com/xianyi/OpenBLAS && cd OpenBLAS && make && make PREFIX=/usr install
RUN		mkdir -p /go/bin /go/pkg /go/src/github.com/Soypete/golear
ENV		GOPATH=/go
RUN go get -u github.com/gonum/blas
RUN		go get -u  github.com/sjwhitworth/golearn github.com/kniren/gota/dataframe\ 
         && go get -u -t gonum.org/v1/gonum/... gonum.org/v1/plot/...

ENV foo src/github.com/Soypete/golearn
WORKDIR ${foo}   # WORKDIR /bar

COPY   . src/github.com/Soypete/golearn