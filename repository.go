package main

import (
	"errors"
	"gopkg.in/mgo.v2"
)
import pb "github.com/binhgo/GoMicro-Vessel/proto/vessel"

const dbName  = "Shippy"
const vesselCollection = "Vessels"

type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Create(*pb.Vessel) error
	Close()
}

type VesselRepository struct {
	session *mgo.Session
}

func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error)  {
	vessels, err := repo.GetAll()
	if err != nil {
		return nil, err
	}

	for _, vessel := range vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}
	return nil, errors.New("No vessel found by that spec")
}

func (repo *VesselRepository) Create(vessel *pb.Vessel) error {
	return repo.Collection().Insert(vessel)
}

func (repo *VesselRepository) GetAll() ([]*pb.Vessel, error) {
	var vessels []*pb.Vessel
	err := repo.Collection().Find(nil).All(&vessels)
	if err != nil {
		return nil, err
	}

	return vessels, nil
}

func (repo *VesselRepository) Close()  {
	repo.session.Close()
}

func (repo *VesselRepository) Collection() *mgo.Collection  {
	return repo.session.DB(dbName).C(vesselCollection)
}
