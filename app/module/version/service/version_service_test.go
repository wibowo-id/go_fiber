package service

import (
	"errors"
	"github.com/google/uuid"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go_fiber_wibowo/app/database/schema"
	"go_fiber_wibowo/app/module/version/repository"
	"go_fiber_wibowo/app/module/version/request"
	"go_fiber_wibowo/utils/paginator"
)

func TestNewVersionService(t *testing.T) {
	type args struct {
		versionRepo repository.VersionRepository
	}
	tests := []struct {
		name string
		args args
		want VersionService
	}{
		{
			name: "TEST_PASS",
			args: args{
				versionRepo: &repository.MockVersionRepository{},
			},
			want: &versionService{
				Repo: &repository.MockVersionRepository{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewVersionService(tt.args.versionRepo)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestVersionService_All(t *testing.T) {
	type fields struct {
		versionRepo *repository.MockVersionRepository
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		wantErr bool
	}{
		{
			name: "TEST_PASS",
			prepare: func(f *fields) {
				f.versionRepo.EXPECT().GetVersions(request.VersionsRequest{}).Return([]*schema.Version{
					{
						Id:         uuid.Must(uuid.Parse("cca2ea8c-feaa-4823-b0d2-ce22dccbbd66")),
						Version:    "1.0.0",
						MinVersion: "1.0.0",
						Url:        "https://play.google.com/store/apps/details?id=com.oss_app&hl=id&pli=1",
					},
				}, paginator.Pagination{}, nil)
			},
			wantErr: false,
		},
		{
			name: "TEST_FAILED_GET_ARTICLES",
			prepare: func(f *fields) {
				f.versionRepo.EXPECT().GetVersions(request.VersionsRequest{}).Return(nil, paginator.Pagination{}, errors.New("error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				versionRepo: repository.NewMockVersionRepository(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := versionService{
				Repo: f.versionRepo,
			}

			if _, _, err := s.All(request.VersionsRequest{}); (err != nil) != tt.wantErr {
				t.Errorf("VersionService.All() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVersionService_Show(t *testing.T) {
	type fields struct {
		versionRepo *repository.MockVersionRepository
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		wantErr bool
	}{
		{
			name: "TEST_PASS",
			prepare: func(f *fields) {
				f.versionRepo.EXPECT().FindOne(uuid.Must(uuid.Parse("cca2ea8c-feaa-4823-b0d2-ce22dccbbd66"))).Return(&schema.Version{
					Id:         uuid.Must(uuid.Parse("cca2ea8c-feaa-4823-b0d2-ce22dccbbd66")),
					Version:    "1.0.0",
					MinVersion: "1.0.0",
					Url:        "https://play.google.com/store/apps/details?id=com.oss_app&hl=id&pli=1",
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "TEST_FAILED_FIND_ONE",
			prepare: func(f *fields) {
				f.versionRepo.EXPECT().FindOne(uuid.Must(uuid.Parse("cca2ea8c-feaa-4823-b0d2-ce22dccbbd66"))).Return(nil, errors.New("failed find one"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				versionRepo: repository.NewMockVersionRepository(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := versionService{
				Repo: f.versionRepo,
			}

			if _, err := s.Show(uuid.Must(uuid.Parse("cca2ea8c-feaa-4823-b0d2-ce22dccbbd66"))); (err != nil) != tt.wantErr {
				t.Errorf("VersionService.Show() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVersionService_Store(t *testing.T) {
	type fields struct {
		versionRepo *repository.MockVersionRepository
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		wantErr bool
	}{
		{
			name: "TEST_PASS",
			prepare: func(f *fields) {
				f.versionRepo.EXPECT().Create(&schema.Version{
					Id:         uuid.Must(uuid.Parse("cca2ea8c-feaa-4823-b0d2-ce22dccbbd66")),
					Version:    "1.0.0",
					MinVersion: "1.0.0",
					Url:        "https://play.google.com/store/apps/details?id=com.oss_app&hl=id&pli=1",
				}).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "TEST_FAILED_CREATE",
			prepare: func(f *fields) {
				f.versionRepo.EXPECT().Create(&schema.Version{
					Id:         uuid.Must(uuid.Parse("cca2ea8c-feaa-4823-b0d2-ce22dccbbd66")),
					Version:    "1.0.0",
					MinVersion: "1.0.0",
					Url:        "https://play.google.com/store/apps/details?id=com.oss_app&hl=id&pli=1",
				}).Return(errors.New("failed create"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				versionRepo: repository.NewMockVersionRepository(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := versionService{
				Repo: f.versionRepo,
			}

			if err := s.Store(request.VersionRequest{
				Id:         uuid.Must(uuid.Parse("cca2ea8c-feaa-4823-b0d2-ce22dccbbd66")),
				Version:    "1.0.0",
				MinVersion: "1.0.0",
				Url:        "https://play.google.com/store/apps/details?id=com.oss_app&hl=id&pli=1",
			}); (err != nil) != tt.wantErr {
				t.Errorf("VersionService.Store() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVersionService_Update(t *testing.T) {
	type fields struct {
		versionRepo *repository.MockVersionRepository
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		wantErr bool
	}{
		{
			name: "TEST_PASS",
			prepare: func(f *fields) {
				f.versionRepo.EXPECT().Update(uuid.Must(uuid.Parse("cca2ea8c-feaa-4823-b0d2-ce22dccbbd66")), &schema.Version{
					Version:    "1.0.0",
					MinVersion: "1.0.0",
					Url:        "https://play.google.com/store/apps/details?id=com.oss_app&hl=id&pli=1",
				}).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "TEST_FAILED_UPDATE",
			prepare: func(f *fields) {
				f.versionRepo.EXPECT().Update(uuid.Must(uuid.Parse("cca2ea8c-feaa-4823-b0d2-ce22dccbbd66")), &schema.Version{
					Version:    "1.0.0",
					MinVersion: "1.0.0",
					Url:        "https://play.google.com/store/apps/details?id=com.oss_app&hl=id&pli=1",
				}).Return(errors.New("failed update"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				versionRepo: repository.NewMockVersionRepository(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := versionService{
				Repo: f.versionRepo,
			}

			if err := s.Update(uuid.Must(uuid.Parse("cca2ea8c-feaa-4823-b0d2-ce22dccbbd66")), request.VersionRequest{
				Version:    "1.0.0",
				MinVersion: "1.0.0",
				Url:        "https://play.google.com/store/apps/details?id=com.oss_app&hl=id&pli=1",
			}); (err != nil) != tt.wantErr {
				t.Errorf("VersionService.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVersionService_Delete(t *testing.T) {
	type fields struct {
		versionRepo *repository.MockVersionRepository
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		wantErr bool
	}{
		{
			name: "TEST_PASS",
			prepare: func(f *fields) {
				f.versionRepo.EXPECT().Delete(uuid.Must(uuid.Parse("cca2ea8c-feaa-4823-b0d2-ce22dccbbd66"))).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "TEST_FAILED_DELETE",
			prepare: func(f *fields) {
				f.versionRepo.EXPECT().Delete(uuid.Must(uuid.Parse("cca2ea8c-feaa-4823-b0d2-ce22dccbbd66"))).Return(errors.New("failed delete"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				versionRepo: repository.NewMockVersionRepository(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := versionService{
				Repo: f.versionRepo,
			}

			if err := s.Destroy(uuid.Must(uuid.Parse("cca2ea8c-feaa-4823-b0d2-ce22dccbbd66"))); (err != nil) != tt.wantErr {
				t.Errorf("VersionService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
