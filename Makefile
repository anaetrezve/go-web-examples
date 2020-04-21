start_db:
	docker run --name gosql -e MYSQL_ROOT_PASSWORD=rootDBPass \
	-e MYSQL_DATABASE=godb \
	-e MYSQL_USER=go \
	-e MYSQL_PASSWORD=pass \
	-d -p "3306:3306" mysql:latest

stop_db:
	docker stop gosql

delete_db:
	docker rm $(docker ps -q --filter ancestor=gosql )


