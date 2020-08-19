FROM postgres:12.3

ENV POSTGRES_HOST: "localhost"
ENV POSTGRES_USER polina
ENV POSTGRES_PASSWORD super_secret
ENV POSTGRES_DB mychat

COPY ./scripts/sql/init.sql /docker-entrypoint-initdb.d/