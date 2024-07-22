package impl_repository

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	"github.com/hrm1810884/works-hai-backend/config"
	"github.com/hrm1810884/works-hai-backend/domain/entity/user"
	"github.com/hrm1810884/works-hai-backend/infrastructure/impl/database"
)

type ImplUserRepository struct {
	Client *firestore.Client
}

func NewImplUserRepository(ctx context.Context) (*ImplUserRepository, error) {
	app, err := config.InitializeApp()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Firebase app: %w", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Firebase app: %w", err)
	}

	return &ImplUserRepository{Client: client}, nil
}

func (ur *ImplUserRepository) Create(user user.User) error {
	ctx := context.Background()
	userData := ConvertUserToData(user)

	err := database.Create(ur.Client, ctx, *userData)
	if err != nil {
		return err
	}

	return nil
}

func (ur *ImplUserRepository) FindById(userId user.UserId) (*user.User, error) {
	ctx := context.Background()

	userData, err := database.FindById(ur.Client, ctx, userId.ToId())
	if err != nil {
		return nil, err
	}

	return ConvertDataToUser(*userData)
}

func (ur *ImplUserRepository) FindByPos(pos user.Position) (*user.User, error) {
	ctx := context.Background()

	userData, err := database.FindByPos(ur.Client, ctx, pos.GetX(), pos.GetY())
	if err != nil {
		return nil, err
	}

	return ConvertDataToUser(*userData)
}

func (ur *ImplUserRepository) FindLatest() (*user.User, error) {
	ctx := context.Background()

	userData, err := database.FindLatest(ur.Client, ctx)
	if err != nil {
		return nil, err
	}

	return ConvertDataToUser(*userData)
}

func (ur *ImplUserRepository) Update(user user.User) error {
	ctx := context.Background()
	userData := ConvertUserToData(user)

	err := database.Update(ur.Client, ctx, *userData)
	if err != nil {
		return err
	}

	return nil
}

func (ur *ImplUserRepository) Delete(userId user.UserId) error {
	ctx := context.Background()

	err := database.Delete(ur.Client, ctx, userId.ToId())
	if err != nil {
		return err
	}

	return nil
}

func ConvertDataToUser(data database.UserData) (*user.User, error) {
	id, err := uuid.Parse(data.UserId)
	if err != nil {
		return nil, fmt.Errorf("failed to convert id to uuid: %w", err)
	}

	userId, err := user.NewUserId(id)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}

	position := user.NewPosition(data.PosX, data.PosY)
	user := user.NewUser(*userId, *position, data.Url, data.IsDrawn, data.CreatedAt, data.UpdatedAt)
	return user, nil
}

func ConvertUserToData(user user.User) *database.UserData {
	now := time.Now()
	return &database.UserData{
		UserId:    user.GetId().ToId(),
		PosX:      user.GetPosition().GetX(),
		PosY:      user.GetPosition().GetY(),
		Url:       user.GetUrl(),
		IsDrawn:   user.IsDrawn(),
		CreatedAt: user.GetCreatedAt(),
		UpdatedAt: now,
	}
}
