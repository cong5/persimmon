package service

import (
	"encoding/json"
	"github.com/cong5/persimmon/app/models"
	"github.com/cong5/persimmon/app/utils"
	"github.com/revel/revel"
	"github.com/revel/revel/cache"
	"time"
)

type NavigationService struct {
	NavKey string
}

func (this *NavigationService) GetNavigation(real bool) ([]models.Navigation, error) {
	NavigationMenu := make([]models.Navigation, 0)
	cKey := utils.CacheKey("NavigationService", "Navigation")

	if real {
		go cache.Delete(cKey)
	}

	if err := cache.Get(cKey, &NavigationMenu); err != nil {
		navigation, navErr := optionService.GetValueByName(this.GetNavKey(), false)
		if navErr != nil {
			return nil, navErr
		}

		err := json.Unmarshal([]byte(navigation), &NavigationMenu)
		if err != nil {
			revel.AppLog.Errorf("Json unmarshal failed. data: %s : error: %s", navigation, err)
			return nil, err
		}
		go cache.Set(cKey, NavigationMenu, 30*time.Minute)
	}

	return NavigationMenu, nil
}

func (this *NavigationService) GetNavKey() string {
	return "navigations"
}
