FROM golang:alpine

ENV GO111MODULE=on
ENV API_PORT=8080
ENV EXPORT_FILE_PATH=/./data
ENV RECORD_CYCLE = 1
WORKDIR /app
COPY . ./
RUN go build -o /goland-memory

CMD [ "/golang-memory" ]