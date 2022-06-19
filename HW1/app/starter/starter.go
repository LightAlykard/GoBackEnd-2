package starter

import (
	"context"
	"sync"

	"github.com/LightAlykard/GoBackEnd-2/hw1/app/models/community"
	"github.com/LightAlykard/GoBackEnd-2/hw1/app/models/user"
)

type App struct {
	us *user.Users
	cs *community.Communities
}

func NewApp(us user.UserStore, cs community.CommunityStore) *App {
	a := &App{
		us: user.NewUsers(us),
		cs: community.NewCommunities(cs),
	}
	return a
}

type APIServer interface {
	Start(us *user.Users, cs *community.Communities)
	Stop()
}

func (a *App) Serve(ctx context.Context, wg *sync.WaitGroup, hs APIServer) {
	defer wg.Done()
	hs.Start(a.us, a.cs)
	<-ctx.Done()
	hs.Stop()
}
