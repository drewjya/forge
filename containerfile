# docker/postgres/Dockerfile
FROM postgres:16

# Install build tools + pg_ulid
RUN apt-get update \
 && apt-get install -y git make gcc postgresql-server-dev-16 \
 && git clone https://github.com/andrielfn/pg-ulid /pg-ulid \
 && cd /pg-ulid \
 && make && make install \
 && apt-get remove -y git make gcc \
 && apt-get autoremove -y \
 && rm -rf /var/lib/apt/lists/* /pg-ulid
