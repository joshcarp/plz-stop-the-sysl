FROM golang:alpine AS build
ENV GOPATH /go
WORKDIR /work
COPY . ./
RUN ls
RUN go build ./cmd/simple

FROM alpine
COPY --from=build /work/simple /
EXPOSE 5751
ENTRYPOINT ["/simple"]
