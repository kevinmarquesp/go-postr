FROM postgres:bookworm

COPY /migrate.sql /docker-entrypoint-initdb.d/

EXPOSE $POSTGRES_PORT
