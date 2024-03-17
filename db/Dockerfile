FROM postgres:bookworm

RUN apt-get update && apt-get install -y python3
COPY /mock/dummy_users.py /docker-entrypoint-initdb.d/

EXPOSE $POSTGRES_PORT
