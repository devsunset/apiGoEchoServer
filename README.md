# apiServer
apiServer (Tesla Chargers App - Api Server)


1. postgres install

* postgres
  $ docker run -d -p 5432:5432 --name postgres -e POSTGRES_PASSWORD=PASSWORD -v /workspace/app/postgres/pgdata:/var/lib/postgresql/data postgres

  * postgres docker shll 접속 
  $ docker exec -it postgres /bin/bash

  * psql 명령 실행 
  psql --username=postgres --dbname=postgres

  CREATE ROLE teslapsql LOGIN CREATEDB PASSWORD 'PASSWORD';
  \du

  CREATE DATABASE teslapsql OWNER teslapsql;
  \l
  
2. config/config.json (
