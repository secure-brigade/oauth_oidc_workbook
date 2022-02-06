package env

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var requiredKeys = []string{
	"HOST",
	"ALLOWED_ORIGINS",
}

// SetupVars read either config/*.toml or environment variables and map to app vars
func SetupVars() {
	if os.Getenv("ENV") == "" {
		log.Panicln("ENV environment variable is required")
	}

	// read from file
	file := "config/" + os.Getenv("ENV") + ".toml"
	viper.SetConfigFile(file)
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicf("Fatal error config file: %s", err)
	}

	// read from env
	replacer := strings.NewReplacer(".", "__")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix("sts")
	viper.AutomaticEnv()
	viper.Set("ENV", os.Getenv("ENV"))

	if IsDebug() {
		printVars()
	}

	validateVars()
}

func validateVars() {
	printPanic := func(v string) {
		log.Panicf("missing required environment variable: %s", v)
	}

	for key, envVar := range viper.AllSettings() {
		isRequired := false
		for _, v := range requiredKeys {
			if key == v {
				isRequired = true
			}
		}

		if !isRequired {
			continue
		}

		if envVar == nil || envVar == "" || len(envVar.([]string)) == 0 {
			printPanic(key)
		}
	}
}

// PrintVars prints environment variables
func printVars() {
	fmt.Println("====configured variables====")
	for key, value := range viper.AllSettings() {
		fmt.Printf("%s=%v\n", key, value)
	}
	fmt.Println("===========================")
}

// IsLocal ... very self-explanatory
func IsLocal() bool {
	d := viper.GetString("ENV")
	return d == "" || d == "local"
}

// IsDebug ... very self-explanatory
func IsDebug() bool {
	return viper.GetBool("DEBUG")
}

// IsDevelop ... very self-explanatory
func IsDevelop() bool {
	return viper.GetString("ENV") == "dev"
}

func GetURLAuthority() string {
	host := viper.GetString("host")
	if host == "" {
		host = "localhost" // default host
	}
	port := viper.GetString("port")
	if port == "" {
		port = "80" // default port
	}

	return host + ":" + port
}

func GetAllowedOrigins() []string {
	return viper.GetStringSlice("ALLOWED_ORIGINS")
}
