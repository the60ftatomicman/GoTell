burn:
	pkill -f *gotell | echo done
	ps

#stty -icanon && nc localhost 9002
run:
	- $(MAKE) burn
	go run .

build-app:
	- docker stop gotell && docker rm -f gotell && docker image rm --force gotell
	docker build -t gotell ./

run-app: 
	- docker stop gotell && docker rm -f gotell
	docker run -it -p 9002:9002 --rm --name gotell gotell

run-docker: build-app run-app

get-ip:
	docker inspect gotell | grep IPAddress