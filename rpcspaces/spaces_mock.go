package rpcspaces

import (
	"context"
	"github.com/pepeunlimited/microservice-kit/errorz"
)

type SpacesMock struct {
	Errors 		errorz.Stack
	IsFile      bool
	File        *File
}

func (s *SpacesMock) GetFile(context.Context, *GetFileParams) (*File, error) {
	s.IsFile = true
	if s.Errors.IsEmpty() {
		if s.File == nil {
			return &File{}, nil
		}
		return s.File, nil
	}
	return nil, s.Errors.Pop()
}

func (s *SpacesMock) GetFiles(context.Context, *GetFilesParams) (*GetFilesResponse, error) {
	panic("implement me")
}

func (s *SpacesMock) GetSpaces(context.Context, *GetSpacesParams) (*GetSpacesResponse, error) {
	panic("implement me")
}

func (s *SpacesMock) Cut(context.Context, *CutParams) (*CutResponse, error) {
	panic("implement me")
}

func (s *SpacesMock) Delete(context.Context, *DeleteParams) (*DeleteResponse, error) {
	panic("implement me")
}

func (s *SpacesMock) Wipe(context.Context, *WipeParams) (*WipeParamsResponse, error) {
	panic("implement me")
}

func (s *SpacesMock) CreateSpaces(context.Context, *CreateSpacesParams) (*CreateSpacesResponse, error) {
	panic("implement me")
}

func NewSpacesMock(errors []error) SpacesService {
	return &SpacesMock{
		Errors: errorz.NewErrorStack(errors),
	}
}