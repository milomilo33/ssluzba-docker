# ## Build
# FROM golang:1.19-alpine AS build

# ENV GO111MODULE=on

# WORKDIR /app

# COPY go.mod .
# COPY go.sum .

# RUN go mod download

# COPY . .

# # RUN go get

# RUN go build -o /uns

# ## Deploy
# FROM gcr.io/distroless/base-debian10

# WORKDIR /

# COPY --from=build /uns /uns

# CMD ["/uns"]

FROM golang:1.19-alpine

ENV GO111MODULE=on

WORKDIR /app

COPY . .

RUN go get

RUN go build -o /uns

EXPOSE 8080

CMD [ "/uns" ]