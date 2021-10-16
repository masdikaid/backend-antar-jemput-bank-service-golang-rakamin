package service

import (
	"backend-a-antar-jemput/internal/contract"
	"backend-a-antar-jemput/internal/databases"
	"backend-a-antar-jemput/internal/repository"
	"backend-a-antar-jemput/internal/service/mocks"
	"reflect"
	"testing"

	mockVal "github.com/stretchr/testify/mock"
)

func Test_GetAgent(t *testing.T) {
	const id = uint(1)
	agent := contract.DetailAGent{
		ID:          uint(1),
		OutletName:  "Berkah Link",
		Name:        "Firman",
		PhoneNumber: "085771002552",
		Province:    "",
		City:        "SLEMAN",
		District:    "GAMPING",
		Address:     "",
		MaxTrx:      int(25000000),
		IsAvailable: true,
		Rating:      float64(0),
		Service:     nil,
	}

	s := struct {
		Repository repository.AgentRepositoryInterface
	}{Repository: repository.NewAgentRepo(databases.DBCon)}

	tests := []struct {
		name      string
		mockAgent func(mock *mocks.ServiceAgentInterface)
		want      *contract.DetailAGent
		wantErr   bool
	}{
		{
			name: "given id will retrun agent details",
			mockAgent: func(mock *mocks.ServiceAgentInterface) {
				mock.On("GetAgent", id).Return(mockVal.AnythingOfType("contract.DetailAgent"), nil)
			},
			want:    &agent,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAgent := new(mocks.ServiceAgentInterface)
			tt.mockAgent(mockAgent)

			res, err := s.Repository.GetByID(id)
			contract := contract.DetailAGent{}
			contract.FromEntity(*res)

			if (err != nil) != tt.wantErr {
				t.Errorf("err=%v, want=%v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(&contract, tt.want) {
				t.Errorf("err=%v, want=%v", &contract, tt.want)
			}
		})
	}
}
