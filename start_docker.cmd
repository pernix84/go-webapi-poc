@rem Adjust if u wanna test

SETLOCAL

SET servicename=mysql-authservice
@rem You need full directory, LIMITATION in Docker for Windows
SET dbdir=d:/git/priv/go-webapi-poc/sql/db
SET startupscriptdir=d:/git/priv/go-webapi-poc/sql

docker network create timemanager_default

docker run -d -p 3336:3306 --name=%servicename% --net=timemanager_default -v %startupscriptdir%:/docker-entrypoint-initdb.d/ -v %dbdir%:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 mysql

@rem only use at home
TIMEOUT 5

docker run -d -p 8000:8000 --name=go-webapi-poc --net=timemanager_default pernix84/go-webapi-poc

ENDLOCAL