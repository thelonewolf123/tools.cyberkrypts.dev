FROM golang:latest

WORKDIR /app

ARG APPLICATION_BASE_URL
ARG DATABASE_URL

ENV GIN_MODE=release
ENV APPLICATION_BASE_URL=${APPLICATION_BASE_URL}
ENV DATABASE_URL=${DATABASE_URL}


COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/a-h/templ/cmd/templ@latest

COPY . .

RUN templ generate
RUN go build -o /app/bin .

CMD [ "/app/bin" ]