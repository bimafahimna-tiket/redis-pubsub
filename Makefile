server:
	go run ./cmd/app/main.go

worker:
	go run ./cmd/worker/main.go

dockerup:
	docker compose up $(BUILD) $(D)

dockerdown:
	docker compose down $(BUILD) $(D)

image:
	docker build --ssh default=${HOME}/.ssh/id_ed25519 .  

redisup:
	docker compose -f ./deployment/docker-compose.redis.yml up $(BUILD) $(D)

redisdown:
	docker compose -f ./deployment/docker-compose.redis.yml down $(BUILD) $(D)

.PHONY: server backendup redisup redisdown image