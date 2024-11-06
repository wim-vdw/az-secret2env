package internal

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Client struct {
	filename        string
	envs            envVariables
	containsSecrets bool
}

func NewClient(filename string) *Client {
	return &Client{
		filename: filename,
	}
}

func readExtraEnvsFromFile(filename string, verboseError bool) error {
	if filename != "" {
		if err := godotenv.Load(filename); err != nil {
			if verboseError {
				return fmt.Errorf("unable to read or parse the specified environment file %q\n%s", filename, err)
			}
			return fmt.Errorf("unable to read or parse the specified environment file %q (use the --verbose flag for more details)", filename)
		}
	}
	return nil
}

func (c *Client) LoadEnvs(verboseError bool) error {
	if err := readExtraEnvsFromFile(c.filename, verboseError); err != nil {
		return err
	}
	for _, env := range os.Environ() {
		parts := strings.SplitN(env, "=", 2)
		name := parts[0]
		value := parts[1]
		prefix := "azure://"
		if strings.HasPrefix(value, prefix) {
			value = strings.TrimPrefix(value, prefix)
			newEnv := envVariable{original: env, name: name, value: value, isSecret: true}
			c.envs = append(c.envs, newEnv)
			c.containsSecrets = true
		} else {
			newEnv := envVariable{original: env, name: name, value: value, isSecret: false}
			c.envs = append(c.envs, newEnv)
		}
	}
	return nil
}

func (c *Client) ConvertSecrets(verboseError, showStatus bool) error {
	for _, env := range c.envs {
		if err := env.convert(verboseError, showStatus); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) PrintDryRunResults() {
	if c.containsSecrets {
		fmt.Println("All secret references in environment variables were successfully converted.")
	} else {
		fmt.Println("No secret references found in environment variables. No conversions were made.")
	}
}
