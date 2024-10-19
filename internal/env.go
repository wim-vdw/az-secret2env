package internal

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azsecrets"
	"github.com/spf13/viper"
)

type EnvVariable struct {
	Original string
	Name     string
	Value    string
	IsSecret bool
}

type EnvVariables []EnvVariable

func (e *EnvVariable) Convert() error {
	if e.IsSecret {
		parts := strings.Split(e.Value, "/")
		if len(parts) != 2 {
			return fmt.Errorf("environment variable %s does not consist of 2 parts", e.Name)
		}
		keyvaultURL := "https://" + parts[0]
		secretName := parts[1]
		verbose := viper.GetBool("verbose")
		secret, err := GetSecretFromKeyvault(keyvaultURL, secretName)
		if err != nil {
			if verbose {
				return fmt.Errorf("could not retrieve secret from keyvault %s\n%s", keyvaultURL, err)
			}
			return fmt.Errorf("could not retrieve secret from keyvault %s (run with --verbose switch for more info)", keyvaultURL)
		}
		if err := os.Setenv(e.Name, secret); err != nil {
			return err
		}
	}
	return nil
}

var cred *azidentity.DefaultAzureCredential

func GetAuth() error {
	var err error
	cred, err = azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return err
	}
	return nil
}

func GetSecretFromKeyvault(keyvaultURL, secretName string) (string, error) {
	err := GetAuth()
	if err != nil {
		return "", err
	}
	client, err := azsecrets.NewClient(keyvaultURL, cred, nil)
	if err != nil {
		return "", err
	}
	ctx := context.Background()
	response, err := client.GetSecret(ctx, secretName, "", nil)
	if err != nil {
		return "", err
	}
	return *response.Value, err
}
