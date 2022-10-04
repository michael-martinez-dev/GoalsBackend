APP_NAME = goals
APP_VERSION = v1
APP_BIN = server

DOCKERHUB_USER=mixedmachine


.PHONY: db dev-api dev-rec pipeline image image-push image-run compose 


build: main.go api/*.go recommender/*.go
	go mod tidy
	go build -o ./bin/$(APP_BIN).exe main.go

db:
	docker compose -f ./build/docker-compose.db.yml up -d

dev-api: db build 
	./api/bin/$(APP_BIN).exe

dev-rec: db build 
	./recommender/bin/$(APP_BIN).exe

pipeline:
	go fmt ./...
	golangci-lint run

dockerfile:
	go build -o ./bin/$(APP_BIN) main.go

image:
	docker build -f ./build/api.Dockerfile -t $(APP_NAME)-api:latest .
	docker build -f ./build/api.Dockerfile -t $(APP_NAME)-api:$(APP_VERSION) .

	docker build -f ./build/rec.Dockerfile -t $(APP_NAME)-recommender:latest .
	docker build -f ./build/rec.Dockerfile -t $(APP_NAME)-recommender:$(APP_VERSION) .

image-push:
	docker tag $(APP_NAME)-api:latest $(DOCKERHUB_USER)/$(APP_NAME)-api:latest
	docker tag $(APP_NAME)-api:$(APP_VERSION) $(DOCKERHUB_USER)/$(APP_NAME)-api:$(APP_VERSION)
	docker tag $(APP_NAME)-recommender:latest $(DOCKERHUB_USER)/$(APP_NAME)-recommender:latest
	docker tag $(APP_NAME)-recommender:$(APP_VERSION) $(DOCKERHUB_USER)/$(APP_NAME)-recommender:$(APP_VERSION)

	docker push $(DOCKERHUB_USER)/$(APP_NAME)-api:latest
	docker push $(DOCKERHUB_USER)/$(APP_NAME)-api:$(APP_VERSION)
	docker push $(DOCKERHUB_USER)/$(APP_NAME)-recommender:latest
	docker push $(DOCKERHUB_USER)/$(APP_NAME)-recommender:$(APP_VERSION)

# This will not be able to connect to db unless you change the .env
# to a reachable host. Instead use compose.
image-run: image
	docker run -d \
	-p 8081:8080 \
	--env-file .env \
	--name $(APP_NAME) \
	$(APP_NAME)-api:latest

	docker run -d \
	-p 8082:8080 \
	--env-file .env \
	--name $(APP_NAME) \
	$(APP_NAME)-recommender:latest

compose: image
	docker compose -f ./build/docker-compose.db.yml up --build -d
	docker compose -f ./build/docker-compose.api.yml up --build -d

clean:
	rm -f ./bin/$(APP_BIN)
	docker rm -f $(APP_NAME)-api $(APP_NAME)-recommender
	docker compose -f ./build/docker-compose.db.yml down
	docker compose -f ./build/docker-compose.api.yml down
