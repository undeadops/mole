# Build Project

webserver:
	cd website && python -m SimpleHTTPServer 8080

run:
	go run api/*.go

docker:
	docker build -t undeadops/mole:latest .

release: docker
	docker push undeadops/mole:latest
