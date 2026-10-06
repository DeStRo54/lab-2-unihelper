package group

import (
	"context"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"homewormanager/internal/entity"
	"homewormanager/internal/errs"
	"testing"
)

func TestService_Create(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		ctx   context.Context
		group entity.Group
	}

	type want struct {
		group entity.GroupID
		err   error
	}

	tests := []struct {
		name    string
		mockFn  func(*MockGroupRepo)
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "Successful case",
			mockFn: func(m *MockGroupRepo) {
				m.EXPECT().GetByName(ctx, "БСБО-01-23").Times(1).Return(entity.Group{}, errs.GroupNotFound)
				m.EXPECT().Create(ctx, gomock.Any()).Times(1).Return(entity.GroupID(1), nil)
			},
			args: args{
				ctx: ctx,
				group: entity.Group{
					Name: "БСБО-01-23",
				},
			},
			want:    want{group: entity.GroupID(1), err: nil},
			wantErr: false,
		},
		{
			name: "Not successful case",
			mockFn: func(m *MockGroupRepo) {
				m.EXPECT().GetByName(ctx, "БСБО-01-23").Times(1).Return(entity.Group{}, nil)
			},
			args: args{
				ctx: ctx,
				group: entity.Group{
					Name: "БСБО-01-23",
				},
			},
			want:    want{group: entity.GroupID(0), err: errs.GroupExists},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			repo := NewMockGroupRepo(ctrl)
			tt.mockFn(repo)

			s := &Service{
				groupRepo: repo,
			}
			got, err := s.Create(tt.args.ctx, tt.args.group)

			if tt.wantErr {
				require.ErrorIs(t, tt.want.err, err)
			} else {
				require.NoError(t, err)
			}

			require.Equal(t, tt.want.group, got)

		})
	}
}
