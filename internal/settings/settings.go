package settings

import "github.com/caarlos0/env/v11"

type Settings struct {
	App      `envPrefix:"APP_"`
	Database `envPrefix:"DATABASE_"`
}

func Load() (Settings, error) {
	return env.ParseAs[Settings]()
}
