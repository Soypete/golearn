FROM	golang:1.13-alpine3.10
RUN     apk update && apk add make gcc g++ linux-headers git perl musl-dev
RUN		git clone https://github.com/xianyi/OpenBLAS && cd OpenBLAS && make && make PREFIX=/usr install
RUN		mkdir -p /go/src /go/bin /go/pkg
ENV		GOPATH=/go
# RUN     go version
RUN go get -u github.com/gonum/blas
RUN		go get -u  github.com/sjwhitworth/golearn github.com/kniren/gota/dataframe\ 
         && go get -u -t gonum.org/v1/gonum/... gonum.org/v1/plot/...

ENV foo src/github.com/Soypete/golearn
RUN mkdir -p src/github.com/Soypete/golearn
RUN ls
RUN ls src
RUN ls src/github.com
WORKDIR ${foo}   # WORKDIR /bar
RUN ls 

COPY   . src/github.com/Soypete/golearn
RUN     ls src/github.com/Soypete/golearn/examples/linear-regression/graphs
# RUN     cd src/github.com/Soypete/golearn
# RUN     pwd
# RUN     go build ${foo}/examples/linear-regression/main.go
# CMD     [.${foo}/examples/linear-regression/linear-regression]
RUN     go run ${foo}/examples/linear-regression/main.go
