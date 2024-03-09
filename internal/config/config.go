package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func ReadConfig(cmd *cobra.Command, path string) (conf *Config, err error) {
	// set config file name and path
	viper.SetConfigName(path)
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	// read config file
	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "an error occurred while reading config file")
	}

	// unmarshal config file into Config struct
	if err := viper.Unmarshal(&conf); err != nil {
		return nil, errors.Wrap(err, "an error occurred while unmarshaling config file")
	}

	// if a flag is set, it overrides the value in config file
	if cmd.Flags().Changed("verbose") {
		verbose, _ := cmd.Flags().GetBool("verbose")
		conf.Verbose = verbose
	}

	return conf, nil
}
