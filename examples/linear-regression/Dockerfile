FROM	golang:1.13-alpine3.10
RUN     apk update && apk add make gcc g++ linux-headers git perl musl-dev
RUN		git clone https://github.com/xianyi/OpenBLAS && cd OpenBLAS && make && make PREFIX=/usr install
RUN		mkdir -p /go/bin /go/pkg /go/srcgolearn
ENV		GOPATH=/go
# RUN     go version
RUN go get -u github.com/gonum/blas
RUN		go get -u  github.com/sjwhitworth/golearn github.com/kniren/gota/dataframe\ 
         && go get -u -t gonum.org/v1/gonum/... gonum.org/v1/plot/...

ENV foo src/golearn
WORKDIR ${foo}

COPY   ../../golearn src/golearn
RUN     go build ${foo}/examples/linear-regression/main.go
CMD     [.${foo}/examples/linear-regression/linear-regression]
# RUN     go run ${foo}/examples/linear-regression/main.go
