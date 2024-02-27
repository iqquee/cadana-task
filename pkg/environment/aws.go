package environment

import (
	"cadana/pkg/helper"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

// fetchFromUpstream retrieves the API key for a given service from AWS Secret Manager
func (e *Env) fetchFromUpstream() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(e.Region),
	})
	if err != nil {
		fmt.Println(err)
	}

	client := secretsmanager.New(sess)

	result, err := client.ListSecrets(&secretsmanager.ListSecretsInput{})
	if err != nil {
		fmt.Println(err)
	}

	// Map to store API keys
	mapResponse := make(map[string]string)

	// Iterate through each secret and retrieve its value
	for _, secret := range result.SecretList {
		// Input to retrieve the secret value
		getSecretInput := &secretsmanager.GetSecretValueInput{
			SecretId: secret.ARN,
		}

		// Retrieve the secret value
		secretValue, err := client.GetSecretValue(getSecretInput)
		if err != nil {
			fmt.Println(err)
		}

		// Store the API key in the map
		mapResponse[*secret.Name] = *secretValue.SecretString
	}

	var aRetry string
	if e.attemptPullFromCloud {
		aRetry = "A re-pull attempt call... "
	}

	printMsg := fmt.Sprintf("%stotal env read: %d", aRetry, len(mapResponse))
	fmt.Println(printMsg)
	if len(mapResponse) == 0 {
		// error occurred.
		fmt.Println("could not use the upstream log at this time")
	}

	e.isFromCloud = true
	// only override what is in e.envCache if the length of newly pulled variables is more than or equal to the prev one
	if len(mapResponse) >= len(e.envCache) {
		e.envCache = mapResponse
	}

	// then recursively call itself. this will keep the environment varialable in sync with what's in cloud
	time.AfterFunc(time.Minute*5, func() {
		e.fetchFromUpstream()
	})
}

// CreateNewSecret creates a new secrete on AWS Secret Manager
func (e *Env) CreateNewSecret(secretName, secretValue string) (string, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(e.Region),
	})
	if err != nil {
		return "", err
	}

	client := secretsmanager.New(sess)

	input := &secretsmanager.CreateSecretInput{
		Name:         aws.String(secretName),
		SecretString: aws.String(secretValue),
	}

	result, err := client.CreateSecret(input)
	if err != nil {
		return "", err
	}

	return *result.Name, nil
}

// GetSecret retrieves a mock secret key for testing
func (Env) GetSecret(secretName string) (string, error) {
	// Simulation for fetching the API key from AWS
	time.Sleep(time.Millisecond * 100)

	if len(secretName) == 0 {
		return "", helper.CustomError("secret name not found")
	}

	return "apiKeyResponse", nil
}
