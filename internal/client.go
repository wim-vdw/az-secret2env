package internal

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
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
	if err := c.ReadExtraEnvsFromFile(); err != nil {
		return err
	}
	for _, env := range os.Environ() {
		parts := strings.SplitN(env, "=", 2)
		name := parts[0]
		value := parts[1]
		prefix := "azure://"
		if strings.HasPrefix(value, prefix) {
			value = strings.TrimPrefix(value, prefix)
			newEnv := EnvVariable{Original: env, Name: name, Value: value, IsSecret: true}
			c.envs = append(c.envs, newEnv)
			c.containsSecrets = true
		} else {
			newEnv := EnvVariable{Original: env, Name: name, Value: value, IsSecret: false}
			c.envs = append(c.envs, newEnv)
		}
	}
	return nil
}

func (c *Client) ConvertSecrets() error {
	for _, env := range c.envs {
		if err := env.Convert(); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) PrintDryRunResults() {
	if c.containsSecrets {
		fmt.Println("Following environment variables are converted with success:")
		fmt.Println("===========================================================")
		for _, env := range c.envs {
			if env.IsSecret {
				fmt.Println(env.Original)
			}
		}
	} else {
		fmt.Println("Environment variables do not contain secrets, no conversions done.")
	}
}

func (c *Client) ReadExtraEnvsFromFile() error {
	if c.filename != "" {
		if err := godotenv.Load(c.filename); err != nil {
			return err
		}
	}
	return nil
}
