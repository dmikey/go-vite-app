# client build
from node:21-bookworm as nodebuilder
WORKDIR /
COPY . .
RUN make build-client

# server build
from golang:1.21.7-bookworm as gobuilder
WORKDIR /
COPY . .
COPY --from=nodebuilder /server/assets /server/assets
RUN make build-server

# final image
from debian:bookworm-slim
WORKDIR /
COPY --from=gobuilder myapp myapp

EXPOSE 8080 8080
ENTRYPOINT ["/myapp", "--headless", "--port", "8080"]