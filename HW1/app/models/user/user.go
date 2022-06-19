package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	ID              uuid.UUID
	Name            string
	UserCommunities []uuid.UUID
}

type UserStore interface {
	Create(ctx context.Context, u User) (*uuid.UUID, error)
	Read(ctx context.Context, uid uuid.UUID) (*User, error)
	Update(ctx context.Context, userID uuid.UUID) (*User, error)
	Search(ctx context.Context, s string) (chan User, error)
}

type Users struct {
	ustore UserStore
}

func NewUsers(us UserStore) *Users {
	return &Users{
		ustore: us,
	}
}

func (u *User) AddCommunity(communityName uuid.UUID) *User {
	updtUserCommunities := append(u.UserCommunities, communityName)
	return &User{
		ID:              u.ID,
		Name:            u.Name,
		UserCommunities: updtUserCommunities,
	}
}

func (u *User) DeleteCommunity(communityName string) *User {
	//TODO: используется Update
	return &User{}
}

func (us *Users) Search(ctx context.Context, s string) (chan User, error) {
	chin, err := us.ustore.Search(ctx, s)
	if err != nil {
		return nil, err
	}
	chout := make(chan User, 100)
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

func (us *Users) Read(ctx context.Context, uid uuid.UUID) (*User, error) {
	u, err := us.ustore.Read(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("read item error: %w", err)
	}
	return u, nil
}

func (us *Users) Create(ctx context.Context, u User) (*User, error) {
	id, err := us.ustore.Create(ctx, u)
	if err != nil {
		return nil, fmt.Errorf("create item error: %w", err)
	}
	u.ID = *id
	return &u, nil
}

func (us *Users) Update(ctx context.Context, userId uuid.UUID) (*User, error) {
	u, err := us.ustore.Update(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("search item error: %w", err)
	}
	return u, nil
}
