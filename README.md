# apiGoEchoServer

apiGoEchoServer (Tesla Chargers App - Api Server)


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

2. tb_tesla_location.sql 스크립트 사용하여 테이블 생성 

3. config/config.json 설정 정보 환경에 맞게 수정 

4. ./swagger_generater.sh

5. ./execute_batch.sh
   cf) tesla api가 값을 제대로 리턴 하지 않는 경우가 많음 여러번 시도 필요 

6. go build main.go

7. ./startup.sh

8. http://ip:8383/swagger/index.html




