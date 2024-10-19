package internal

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azsecrets"
)

const (
	ColorRed    = "\u001b[31m"
	ColorGreen  = "\u001b[32m"
	ColorYellow = "\u001b[33m"
	ColorReset  = "\u001b[0m"
)

type EnvVariable struct {
	original string
	name     string
	value    string
	isSecret bool
}

type EnvVariables []EnvVariable

func (ev *EnvVariable) Convert(verboseError, showStatus bool) error {
	if ev.isSecret {
		if showStatus {
			fmt.Printf("%s[PROCESS]%s %s", ColorYellow, ColorReset, ev.original)
		}
		parts := strings.Split(ev.value, "/")
		if len(parts) != 2 {
			if showStatus {
				fmt.Printf("\r%s[FAIL]   %s %s\n", ColorRed, ColorReset, ev.original)
			}
			return fmt.Errorf("environment variable %s does not consist of 2 parts", ev.name)
		}
		keyvaultURL := "https://" + parts[0]
		secretName := parts[1]
		secret, err := GetSecretFromKeyvault(keyvaultURL, secretName)
		if err != nil {
			if showStatus {
				fmt.Printf("\r%s[FAIL]   %s %s\n", ColorRed, ColorReset, ev.original)
			}
			if verboseError {
				return fmt.Errorf("could not retrieve secret %q from keyvault %q\n%s", secretName, keyvaultURL, err)
			}
			return fmt.Errorf("could not retrieve secret %q from keyvault %q (use --verbose switch for more info)", secretName, keyvaultURL)
		}
		if err := os.Setenv(ev.name, secret); err != nil {
			return err
		}
		if showStatus {
			fmt.Printf("\r%s[OK]     %s %s\n", ColorGreen, ColorReset, ev.original)
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
