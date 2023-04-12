package services

import "github.com/runan13/echo-rest-api/types"

type SettingsService interface {
	GetAllLocales() ([]types.Locale, string, error)
}

type settingService struct {
}

func NewSettingsService() SettingsService {
	return &settingService{}
}

func (s *settingService) GetAllLocales() ([]types.Locale, string, error) {
	return nil, "", nil
}
