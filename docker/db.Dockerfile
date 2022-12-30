FROM postgres:15

ENV POSTGRES_DB nenga
ENV POSTGRES_PASSWORD thisisunsafepassword

RUN apt-get update && \
    rm -rf /var/lib/apt/lists/*

RUN mkdir -p /docker-entrypoint-initdb.d
COPY ./sql/ /docker-entrypoint-initdb.d/

EXPOSE 5432
