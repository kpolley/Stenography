FROM golang:1.16-alpine

WORKDIR /app

# Copys everything into root dir
COPY . ./

RUN go build -o /start-server

EXPOSE 80
CMD [ "/start-server" ]