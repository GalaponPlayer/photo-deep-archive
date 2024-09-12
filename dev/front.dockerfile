FROM node:20

# Install Go
ENV GO_VERSION 1.21.1
RUN apt-get update && apt-get install -y --no-install-recommends \
    curl \
    git && \
    curl -LO https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz && \
    rm go${GO_VERSION}.linux-amd64.tar.gz && \
    rm -rf /var/lib/apt/lists/*

# Set Go environment variables
RUN mkdir -p /go && chmod -R 777 /go
ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/go"
ENV GOBIN="${GOPATH}/bin"


RUN mkdir /app && chown node:node /app
WORKDIR /app

USER node

COPY --chown=node:node ../manage-console/package.json ../manage-console/yarn.lock ./

RUN yarn install

COPY --chown=node:node . .

EXPOSE 80

# CMD [ "yarn", "dev" ]