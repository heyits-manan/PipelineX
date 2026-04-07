package video

import "context"
import "errors"
import "time"

type IDGenerator interface {
	NewID() string
}

type JobPublisher interface {
	PublishVideoUploaded(ctx context.Context, video Video) error
}

type Service struct {
	store Store
	ids   IDGenerator
	jobs  JobPublisher
}

func NewService(store Store, ids IDGenerator, jobs JobPublisher) *Service {
	return &Service{
		store: store,
		ids: ids,
		jobs: jobs,
	}
}

func (s *Service) CreateVideo(ctx context.Context, input *CreateVideoInput) (Video, error){
	if input.Filename == ""{
		return Video{}, errors.New("filename is required")
	}

	now := time.Now()
	video := Video{
		ID:        s.ids.NewID(),
		Filename:  input.Filename,
		Status:    StatusUploaded,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := s.store.Create(ctx, video); err != nil{
		return Video{}, err
	}

	if s.jobs != nil{
		if err := s.jobs.PublishVideoUploaded(ctx, video); err != nil{
			return Video{}, err
		}
	}
	return video, nil
}

func (s *Service) GetVideo(ctx context.Context, id string) (Video, error){
	if id == ""{
		return Video{}, errors.New("video id is required")
	}
	video, err := s.store.GetByID(ctx, id)
	if err != nil{
		return Video{}, err
	}

	return video, nil

}

func (s *Service) ListVideos(ctx context.Context) ([]Video, error){
	videos, err := s.store.ListVideos(ctx)

	if err != nil{
		return nil, err
	}

	return videos, nil
}	

// func (s *Service) MarkProcessing(ctx context.Context, input UpdateVideoStatusInput) error{

// }


// func (s *Service) MarkReady(ctx context.Context, input UpdateVideoStatusInput) error{

// }

// func (s *Service) MarkFailed(ctx context.Context, input UpdateVideoStatusInput) error{

// }

