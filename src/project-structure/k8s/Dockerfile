FROM golang:1.13.6 as build
WORKDIR /go/src
COPY . .
RUN CGO_ENABLED=0 make build-mdserver // HL

FROM scratch // HL
COPY --from=build /go/src/build/mdserver /entrypoint
ENTRYPOINT ["/entrypoint"]
