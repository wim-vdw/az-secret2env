package internal

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Client struct {
	filename        string
	envs            EnvVariables
	containsSecrets bool
}

func NewClient(filename string) *Client {
	return &Client{
		filename: filename,
	}
}

func (c *Client) LoadEnvs() error {
	verboseError := viper.GetBool("verbose")
	if err := c.ReadExtraEnvsFromFile(verboseError); err != nil {
		return err
	}
	for _, env := range os.Environ() {
		parts := strings.SplitN(env, "=", 2)
		name := parts[0]
		value := parts[1]
		prefix := "azure://"
		if strings.HasPrefix(value, prefix) {
			value = strings.TrimPrefix(value, prefix)
			newEnv := EnvVariable{original: env, name: name, value: value, isSecret: true}
			c.envs = append(c.envs, newEnv)
			c.containsSecrets = true
		} else {
			newEnv := EnvVariable{original: env, name: name, value: value, isSecret: false}
			c.envs = append(c.envs, newEnv)
		}
	}
	return nil
}

func (c *Client) ConvertSecrets() error {
	verboseError := viper.GetBool("verbose")
	for _, env := range c.envs {
		if err := env.Convert(verboseError, true); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) PrintDryRunResults() {
	if c.containsSecrets {
		fmt.Println("All secrets in environment variables converted with success.")
	} else {
		fmt.Println("Environment variables do not contain secrets, no conversions done.")
	}
}

func (c *Client) ReadExtraEnvsFromFile(verboseError bool) error {
	if c.filename != "" {
		if err := godotenv.Load(c.filename); err != nil {
			if verboseError {
				return fmt.Errorf("could not read or parse env file %q\n%s", c.filename, err)
			}
			return fmt.Errorf("could not read or parse env file %q (use --verbose switch for more info)", c.filename)
		}
	}
	return nil
}
