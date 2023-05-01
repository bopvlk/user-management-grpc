
up:
	sudo docker-compose -f build/docker-compose.yml up --build --remove-orphans

up1:
	sudo docker-compose -f build/docker-compose.yml up --build

up2:
	sudo docker-compose -f build/docker-compose.yml up

down:
	sudo docker-compose -f build/docker-compose.yml down -v