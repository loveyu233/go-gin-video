package initialization

import "github.com/spf13/viper"

func InitViper(isTest bool) error {
	path := ""
	if isTest {
		path = "../config"
	} else {
		path = "./config"
	}
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}
