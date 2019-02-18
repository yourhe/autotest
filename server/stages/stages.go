package stages

import (
	"context"

	"github.com/yourhe/autotest/utils"

	"github.com/yourhe/autotest/proto/stages"
	"gitlab.iyorhe.com/wfgz/reverseproxy/service/stores"
)

//Stages server ....
type Stages struct {
	store stores.Store
}

// s *Stages stages.StagesServer
func (s *Stages) methodName() {

}

func (s *Stages) CreateStages(context.Context, *stages.CreateStagesRequest) (*stages.CreateStagesResponse, error) {
	key := utils.NewID()
	s.store.Write(key, nil)
	panic("not implemented")

}

func (s *Stages) GetStages(context.Context, *stages.SearchStagesRequest) (*stages.CreateStagesResponse, error) {
	panic("not implemented")
}

func (s *Stages) ListStages(context.Context, *stages.SearchStagesRequest) (*stages.CreateStagesResponse, error) {
	panic("not implemented")
}

func (s *Stages) UpdateStages(context.Context, *stages.UpdateStagesRequest) (*stages.CreateStagesResponse, error) {
	panic("not implemented")
}

func (s *Stages) RemoveStages(context.Context, *stages.CreateStagesRequest) (*stages.CreateStagesResponse, error) {
	panic("not implemented")
}

//New instace Stages Server
func New(opts ...Option) *Stages {
	options := defaultOptions()
	return &Stages{
		store: options.Store,
	}
}
