# build stage
FROM golang:1.22-bullseye

RUN mkdir /app
WORKDIR /app

RUN apt-get update && apt-get upgrade -y && apt-get install -y git gcc libc-dev

RUN go install -v github.com/cespare/reflex@latest
EXPOSE 8080
ENTRYPOINT ["reflex", "-c", "reflex.conf"]
