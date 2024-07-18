FROM golang:1.22

WORKDIR /app
ADD ./ /app/
RUN go mod download all
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/usermgt
RUN chmod a+x /app/usermgt

EXPOSE 8080

CMD [ "/app/usermgt" ]