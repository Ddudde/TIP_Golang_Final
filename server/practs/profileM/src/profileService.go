package src

import (
	"context"
	"net/http"

	"practs/profileM/db"
	"practs/profileM/pb"
)

type Server struct {
	H db.Handler
	pb.UnimplementedProfileServiceServer
}

func (s *Server) GetProfile(ctx context.Context, prof *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {
	var user db.User

	if result := s.H.DB.First(&user, "id = ?", prof.UserId); result.Error != nil {
		return &pb.GetProfileResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	s.H.DB.Create(&user)

	return &pb.GetProfileResponse{
		Status:  http.StatusOK,
		Logname: user.Logname,
	}, nil
}
