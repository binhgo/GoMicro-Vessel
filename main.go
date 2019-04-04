package main

import (
	"fmt"
	pb "github.com/binhgo/GoMicro-Vessel/proto/vessel"
	"github.com/micro/go-micro"
	"log"
	"os"
)

const defaultHost = "localhost:27017"


func main() {

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = defaultHost
	}
	session, err := CreateSession(host)
	defer session.Close()

	if err != nil {
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}

	repo := &VesselRepository{session.Copy()}
	createDummyData(repo)

	//vessels := []*pb.Vessel {
	//	&pb.Vessel{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
	//}

	//repo := &VesselRepository{vessels:vessels}

	service := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)

	service.Init()

	pb.RegisterVesselServiceHandler(service.Server(), &Handler{session:session})
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}


func createDummyData(repo Repository) {
	defer repo.Close()
	vessels := []*pb.Vessel{
		{Id: "vessel-001", Name: "Kane's Salty Secret", MaxWeight: 5000, Capacity: 100},
		{Id: "vessel-002", Name: "Peter's Salty Secret", MaxWeight: 10000, Capacity: 300},
		{Id: "vessel-003", Name: "Big's Salty Secret", MaxWeight: 15000, Capacity: 500},
		{Id: "vessel-004", Name: "Grand's Salty Secret", MaxWeight: 200000, Capacity: 700},
	}
	for _, v := range vessels {
		repo.Create(v)
	}
}
