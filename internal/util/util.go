package util

import (
	"fmt"
	"internal/internal/config"
)

// CtxKey Context Key
type CtxKey string

type ViaHttpRequest struct {
	
}

// SiteTemplate loads the correct template directory for the site
func SiteTemplate(path string) (string, error) {

	cfg, err := config.GetConfig()
	if err != nil {
		return "", fmt.Errorf("error loading template directory %s", err)
	}
	return "sites/" + cfg.Site + path, nil

}
