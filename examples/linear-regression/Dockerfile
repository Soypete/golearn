
FROM golang:alpine
RUN  apk update && apk add make linux-headers git
RUN mkdir go/src/github.com/Soypete -p
RUN cd go/src/github.com/Soypete
CMD ls
RUN git clone https://github.com/Soypete/golearn.git

RUN cd golearn/examples/linear-regression
CMD ls
CMD go build ./...
ENTRYPOINT [ ./linear-regression ]