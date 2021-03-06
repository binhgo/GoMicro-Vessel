build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/binhgo/GoMicro-Vessel \
	proto/vessel/vessel.proto
	docker login --username huynhbinh -p Cicevn2007
	docker build -t vessel-service .

run:
	docker run --link=GoMicroMongoDB:mongodb --name GoMicroVesselService -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns -e DB_HOST=mongodb:27017 vessel-service