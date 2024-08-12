package initialize

import (
	"fmt"
	"golang-BE/global"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	// read configuration
	err := viper.ReadInConfig();
	if err != nil {
		panic(fmt.Errorf("failed to read configuration %w \n ", err))
	}

	fmt.Println("Server Port ::", viper.GetInt("server.port"))
	fmt.Println("Server Security ::", viper.GetString("security.jwt.key"))

	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("unable to decode configuration %v", err)
	}
}