package homework

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"homewormanager/internal/entity"
	"testing"
)

func TestService_Create(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		ctx      context.Context
		homework entity.Homework
	}

	type want struct {
		group entity.HomeworkID
		err   error
	}

	tests := []struct {
		name    string
		mockFn  func(*MockHomeworkRepo)
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "Successful case",
			mockFn: func(m *MockHomeworkRepo) {
				m.EXPECT().Create(ctx, gomock.Any()).Times(1).Return(entity.HomeworkID(1), nil)
			},
			args: args{
				ctx:      ctx,
				homework: entity.Homework{},
			},
			want: want{
				group: 1,
				err:   nil,
			},
			wantErr: false,
		},
		{
			name: "Not successful case",
			mockFn: func(m *MockHomeworkRepo) {
				m.EXPECT().Create(ctx, gomock.Any()).Times(1).Return(entity.HomeworkID(0), sql.ErrConnDone)
			},
			args: args{
				ctx:      ctx,
				homework: entity.Homework{},
			},
			want: want{
				group: 0,
				err:   sql.ErrConnDone,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewMockHomeworkRepo(ctrl)
			tt.mockFn(repo)

			s := &Service{
				homeworkRepo: repo,
			}
			got, err := s.Create(tt.args.ctx, tt.args.homework)
			if tt.wantErr {
				require.ErrorIs(t, tt.want.err, err)
			} else {
				require.NoError(t, err)
			}

			require.Equal(t, tt.want.group, got)
		})
	}
}
