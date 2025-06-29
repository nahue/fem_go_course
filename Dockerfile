FROM golang:1.24-alpine AS build
WORKDIR /app
COPY . .
RUN go build -o main .

FROM scratch
COPY --from=build /app/main .
CMD ["./main"]
EXPOSE 8080
ENTRYPOINT ["/main"]
