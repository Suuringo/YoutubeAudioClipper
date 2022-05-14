# Build
FROM golang:1.18.2-alpine AS build

WORKDIR /app

COPY . ./

WORKDIR /app/cmd/serve-http-clipper

RUN go build -o /app/serve-http-clipper

# Deploy
FROM alpine:3.15.4
WORKDIR /
COPY --from=build /app/serve-http-clipper /app/serve-http-clipper

RUN apk add ffmpeg youtube-dl

EXPOSE 3030

CMD ["./app/serve-http-clipper"]
