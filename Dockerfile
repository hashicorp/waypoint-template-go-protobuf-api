# Copyright (c) HashiCorp, Inc.

#--------------------------------------------------------------------
# builder
#--------------------------------------------------------------------

FROM golang:1.19.3-alpine3.17 AS builder

WORKDIR /app-src

COPY go.mod ./
COPY go.sum ./

RUN go mod tidy

RUN go mod download

COPY . ./

RUN go build -o /tmp/%%wp_project%% cmd/%%wp_project%%-api/main.go

# Copy in Waypoint Entrypoint to app image
FROM hashicorp/waypoint:0.11.1 AS ceb

#--------------------------------------------------------------------
# final image
#--------------------------------------------------------------------

FROM alpine:3.17

# Config file will be delivered here at runtime by waypoint
RUN mkdir /opt/config

COPY --from=ceb /usr/bin/waypoint-entrypoint /waypoint-entrypoint

COPY --from=builder /tmp/%%wp_project%% /%%wp_project%%
EXPOSE 8080
CMD [ "/waypoint-entrypoint", "/%%wp_project%%" ]