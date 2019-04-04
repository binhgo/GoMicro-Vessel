package main

import (
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2"
)
import pb "github.com/binhgo/GoMicro-Vessel/proto/vessel"

type Handler struct {
	session *mgo.Session
}

func (han *Handler) GetRepo() Repository {
	repo := &VesselRepository{han.session.Clone()}
	return repo
}

func (han *Handler) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	// Find the next available vessel
	vessel, err := han.GetRepo().FindAvailable(req)
	if err != nil {
		return err
	}
	// Set the vessel as part of the response message type
	res.Vessel = vessel
	return nil
}

func (han *Handler) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	repo := han.GetRepo()
	defer repo.Close()

	err := repo.Create(req)
	if err != nil {
		return err
	}

	res.Created = true
	res.Vessel = req
	return nil
}