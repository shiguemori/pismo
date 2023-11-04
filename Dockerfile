FROM postgres:14

ENV POSTGRES_DB=app_db
ENV POSTGRES_USER=app_user
ENV POSTGRES_PASSWORD=app_pass
  
COPY db/script.sql /docker-entrypoint-initdb.d/

EXPOSE 5432
