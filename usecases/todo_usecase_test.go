package usecases

import (
	"github.com/golang/mock/gomock"
	"github.com/greenteabiscuit/gomock-api-server/domain/model"
	"github.com/greenteabiscuit/gomock-api-server/mock/mock_repository"
	"reflect"
	"testing"
)

func TestList(t *testing.T) {
	cases := []struct {
		name string
		in []*model.Todo
	} {
		{
			"No results returned",
			nil,
		},
		{
			"2 results returned",
			[]*model.Todo{
					{
						ID: 1,
						Item: "homework",
					},
					{
						ID: 2,
						Item: "wash dishes",
					},
				},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, c := range cases {
		var err error
		mockSample := mock_repository.NewMockTodoRepository(ctrl)
		mockSample.EXPECT().List(gomock.Any(), gomock.Any()).Return(c.in, err)

		todoUsecase := NewTodoUsecase(mockSample)
		result, err := todoUsecase.FindAllTodos(nil, nil)

		if err != nil {
			t.Error("Is not same as expected")
		}

		if !reflect.DeepEqual(result, c.in) {
			t.Errorf("Is not same as expected")
		}
	}
}
