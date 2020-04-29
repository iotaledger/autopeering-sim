package config

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

// Config keys
const (
	numberNodes  = "NumberNodes"
	duration     = "Duration"
	saltLifetime = "SaltLifetime"
	vEnabled     = "VisualEnabled"
	dropOnUpdate = "DropOnUpdate"
	mana         = "mana"
	r            = "R"
	ro           = "Ro"
	zipf         = "zipf"
)

func init() {
	viper.SetDefault(numberNodes, 100)
	viper.SetDefault(duration, 60)
	viper.SetDefault(saltLifetime, 60*60)
	viper.SetDefault(vEnabled, false)
	viper.SetDefault(dropOnUpdate, false)
	viper.SetDefault(mana, false)
	viper.SetDefault(r, 10)
	viper.SetDefault(ro, 2.)
	viper.SetDefault(zipf, 1.)
}

func Load() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Using default config")
		} else {
			log.Fatal(err)
		}
	}
}

func PrintConfig() {
	settings := viper.AllSettings()
	if cfg, err := json.MarshalIndent(settings, "", "  "); err == nil {
		fmt.Println(string(cfg))
	}
}

func NumberNodes() int {
	return viper.GetInt(numberNodes)
}

func Duration() time.Duration {
	return time.Duration(viper.GetInt(duration)) * time.Second
}

func SaltLifetime() time.Duration {
	return time.Duration(viper.GetInt(saltLifetime)) * time.Second
}

func DropOnUpdate() bool {
	return viper.GetBool(dropOnUpdate)
}

func VisEnabled() bool {
	return viper.GetBool(vEnabled)
}

func Mana() bool {
	return viper.GetBool(mana)
}

func R() int {
	return viper.GetInt(r)
}

func Ro() float64 {
	return viper.GetFloat64(ro)
}

func Zipf() float64 {
	return viper.GetFloat64(zipf)
}
