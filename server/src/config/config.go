package config

import (
	"log"
	"os"
	"path"
	"strings"

	"github.com/spf13/viper"
)

type SecretManagerPostgres struct {
	DbClusterIdentifier string `json:"dbClusterIdentifier"`
	Password            string `json:"password"`
	Dbname              string `json:"dbname"`
	Engine              string `json:"engine"`
	Port                int    `json:"port"`
	Host                string `json:"host"`
	Username            string `json:"username"`
}

func Init() {
	home, err := os.UserHomeDir()
	if err != nil {
		home = "/root"
	}

	os.MkdirAll(path.Join(home, ".komiser", "data"), 0700)

	viper.SetDefault("debug", false)

	viper.SetConfigName("komiser")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.komiser")
	viper.AddConfigPath("/etc/komiser")
	viper.ReadInConfig()

	viper.SetEnvPrefix("komiser")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}

func Remember(ledger string) {
	ledgers := viper.Get("ledgers").([]interface{})

	for _, v := range ledgers {
		if ledger == v.(string) {
			return
		}
	}

	err := viper.WriteConfig()

	if err != nil {
		log.Printf(
			"failed to write config: komiser will not be remembered\n",
		)
	}
}
