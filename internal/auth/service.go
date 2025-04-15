package auth

import (
    "context"
    "errors"
    "time"

    "github.com/Potagashev/breddit_threads/internal/users"
    "google.golang.org/grpc"
    pb "github.com/Potagashev/breddit_threads/internal/auth/proto" // Adjust the import path to the correct location of your proto package
)

type AuthService struct {
    authClient pb.AuthServiceClient
}

func NewAuthService(conn *grpc.ClientConn) *AuthService {
    return &AuthService{
        authClient: pb.NewAuthServiceClient(conn),
    }
}

func (s *AuthService) VerifyToken(ctx context.Context, in *pb.VerifyTokenRequest, opts ...grpc.CallOption) (*pb.VerifyTokenResponse, error) {
    return &pb.VerifyTokenResponse{
        Id: user.UserId,
        Username: user.Username,
        Email: user.Email,
    }, nil
}

func (s *AuthService) GetUserByToken(token string) (*users.User, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    req := &pb.VerifyTokenRequest{Token: token}
    resp, err := s.authClient.VerifyToken(ctx, req)
    if err != nil {
        return nil, err
    }

    if resp == nil {
        return nil, errors.New("failed to validate token")
    }

    return &users.User{
        UserId:   resp.Id,
        Username: resp.Username,
        Email:    resp.Email,
    }, nil
}
