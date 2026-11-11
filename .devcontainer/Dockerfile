FROM golang:1.23

RUN apt-get update && apt-get install -y --no-install-recommends \
    curl \
    && rm -rf /var/lib/apt/lists/*

RUN curl -fsSL https://github.com/sqlc-dev/sqlc/releases/download/v1.29.0/sqlc_1.29.0_linux_amd64.tar.gz \
    | tar -xz -C /usr/local/bin sqlc

# templ para compilar componentes SSR (.templ -> .go)
RUN curl -fsSL https://github.com/a-h/templ/releases/download/v0.3.833/templ_Linux_x86_64.tar.gz \
    | tar -xz -C /usr/local/bin templ

# hurl para correr las pruebas de la API (requests.hurl)
RUN curl -fsSL -o /tmp/hurl.deb https://github.com/Orange-OpenSource/hurl/releases/download/5.0.1/hurl_5.0.1_amd64.deb \
    && apt-get update && apt-get install -y --no-install-recommends /tmp/hurl.deb \
    && rm -rf /tmp/hurl.deb /var/lib/apt/lists/*

WORKDIR /workspace
