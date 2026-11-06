package domain

import "errors"

var (
	ErrCampaignNotFound      = errors.New("campaign not found")
	ErrCampaignInvalid       = errors.New("campaign invalid")
	ErrCampaignCreate        = errors.New("failed to create campaign")
	ErrCampaignUpdate        = errors.New("failed to update campaign")
	ErrCampaignAlreadyExists = errors.New("campaign already exists")
)
