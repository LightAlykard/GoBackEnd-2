package community

import (
	"context"
	"fmt"

	"github.com/LightAlykard/GoBackEnd-2/hw1/app/models/user"
	"github.com/google/uuid"
)

type Community struct {
	ID    uuid.UUID
	Name  string
	Users []uuid.UUID
}

type CommunityStore interface {
	Create(ctx context.Context, c Community) (*uuid.UUID, error)
	Read(ctx context.Context, communityId uuid.UUID) (*Community, error)
	Update(ctx context.Context, communityID uuid.UUID) (*Community, error)
	SearchComms(ctx context.Context, s string) (chan Community, error)
}

type Communities struct {
	commstore CommunityStore
}

func NewCommunities(commstore CommunityStore) *Communities {
	return &Communities{
		commstore: commstore,
	}
}

func (cs *Communities) AddUserToCommunity(ctx context.Context, commynityID uuid.UUID, u user.User) (*Community, error) {
	c, err := cs.commstore.Read(ctx, commynityID)
	if err != nil {
		return nil, fmt.Errorf("community not found: %w", err)
	}

	return &Community{
		ID:    c.ID,
		Name:  c.Name,
		Users: append(c.Users, u),
	}, nil
}

func (us *Communities) SearchComms(ctx context.Context, s string) (chan Community, error) {

	chin, err := us.commstore.SearchComms(ctx, s)
	if err != nil {
		return nil, err
	}
	chout := make(chan Community, 100)
	go func() {
		defer close(chout)
		for {
			select {
			case <-ctx.Done():
				return
			case u, ok := <-chin:
				if !ok {
					return
				}
				chout <- u
			}
		}
	}()
	return chout, nil
}

func (us *Communities) Create(ctx context.Context, u Community) (*Community, error) {
	id, err := us.commstore.Create(ctx, u)
	if err != nil {
		return nil, fmt.Errorf("create item error: %w", err)
	}
	u.ID = *id
	return &u, nil
}

func (us *Communities) Update(ctx context.Context, userId uuid.UUID) (*Community, error) {
	u, err := us.commstore.Update(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("search item error: %w", err)
	}
	return u, nil
}

func (us *Communities) Read(ctx context.Context, uid uuid.UUID) (*Community, error) {
	u, err := us.commstore.Read(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("read item error: %w", err)
	}
	return u, nil
}
