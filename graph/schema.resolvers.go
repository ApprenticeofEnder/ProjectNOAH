package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.60

import (
	"context"
	"fmt"

	"github.com/ApprenticeofEnder/ProjectNOAH/graph/model"
	"github.com/google/uuid"
)

// User is the resolver for the user field.
func (r *campaignResolver) User(ctx context.Context, obj *model.Campaign) (*model.User, error) {
	return &model.User{
		SupabaseID:       "DUMMY-ID",
		Email:            "test@test.com",
		MarketingConsent: true,
	}, nil
}

// Scenes is the resolver for the scenes field.
func (r *campaignResolver) Scenes(ctx context.Context, obj *model.Campaign) ([]*model.Scene, error) {
	var scenes []*model.Scene
	var dummyScene = model.Scene{
		Title: "Scene 1",
		Notes: "Some markdown here!",
	}
	scenes = append(scenes, &dummyScene)
	return scenes, nil
}

// Npcs is the resolver for the npcs field.
func (r *campaignResolver) Npcs(ctx context.Context, obj *model.Campaign) ([]*model.Npc, error) {
	var npcs []*model.Npc
	var dummyNpc = model.Npc{
		Name: "Jason Fullboar",
		CombatClass: &model.NpcClass{
			Name: "Ace",
			Role: "Striker",
		},
	}
	npcs = append(npcs, &dummyNpc)
	return npcs, nil
}

// Combats is the resolver for the combats field.
func (r *campaignResolver) Combats(ctx context.Context, obj *model.Campaign) ([]*model.Combat, error) {
	var combats []*model.Combat
	var dummyCombat = model.Combat{
		Title: "Combat 1",
		Notes: "This is the first combat of the campaign.",
	}
	combats = append(combats, &dummyCombat)
	return combats, nil
}

// CreateCampaign is the resolver for the createCampaign field.
func (r *mutationResolver) CreateCampaign(ctx context.Context, input model.NewCampaign) (*model.Campaign, error) {
	var campaign model.Campaign
	var user model.User
	campaign.Title = input.Title
	campaign.Description = input.Description
	user.Email = "test@test.com"
	campaign.User = &user
	return &campaign, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	var user model.User
	user.Email = input.Email
	user.MarketingConsent = input.MarketingConsent
	user.SupabaseID = input.SupabaseID
	return &user, nil
}

// Campaigns is the resolver for the campaigns field.
func (r *queryResolver) Campaigns(ctx context.Context) ([]*model.Campaign, error) {
	var campaigns []*model.Campaign
	dummyCampaign := model.Campaign{
		Title:       "Operation Shattered Vision",
		Description: "SSMR but More AC6 Vibes",
	}
	campaigns = append(campaigns, &dummyCampaign)
	return campaigns, nil
}

// Scenes is the resolver for the scenes field.
func (r *queryResolver) Scenes(ctx context.Context, campaignID uuid.UUID) ([]*model.Scene, error) {
	panic(fmt.Errorf("not implemented: Scenes - scenes"))
}

// Combats is the resolver for the combats field.
func (r *queryResolver) Combats(ctx context.Context, campaignID uuid.UUID) ([]*model.Combat, error) {
	panic(fmt.Errorf("not implemented: Combats - combats"))
}

// Npcs is the resolver for the npcs field.
func (r *queryResolver) Npcs(ctx context.Context, campaignID uuid.UUID) ([]*model.Npc, error) {
	panic(fmt.Errorf("not implemented: Npcs - npcs"))
}

// Campaign returns CampaignResolver implementation.
func (r *Resolver) Campaign() CampaignResolver { return &campaignResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type campaignResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
