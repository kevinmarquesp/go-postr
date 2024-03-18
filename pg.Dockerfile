FROM postgres:bookworm

RUN apt-get update && apt-get install -y python3

COPY /db/migrate.sql /docker-entrypoint-initdb.d/

EXPOSE $POSTGRES_PORT
