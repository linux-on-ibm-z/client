package client

import "github.com/rai-project/aws"

// creates an AWS session. this uses STS
// which allows us to provide a session
// that's only valid for certain amount
// of time.
func (c *Client) createAWSSession() error {

	// Create an AWS session
	session, err := aws.NewSession(
		aws.Region(aws.AWSRegionUSEast1),
		aws.AccessKey(aws.Config.AccessKey),
		aws.SecretKey(aws.Config.SecretKey),
		aws.Sts(c.ID.Hex()),
	)
	if err != nil {
		return err
	}
	c.awsSession = session
	return nil
}

// create an authentication token for AWS and fix
// the docker credientials in the job request
func (c *Client) Authenticate() error {
	if err := c.createAWSSession(); err != nil {
		return err
	}

	if err := c.fixDockerPushCredentials(); err != nil {
		return err
	}

	return nil
}
