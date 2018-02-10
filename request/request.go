package request

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/ysugimoto/ginger/config"
)

var debug string = ""

func generateStatementId(sType string) string {
	return fmt.Sprintf("ginger-statement-%s-%d", sType, time.Now().UnixNano())
}

func createAWSSession(c *config.Config) *session.Session {
	conf := aws.NewConfig().WithRegion(c.Project.Region)
	if c.Project.Profile != "" {
		conf = conf.WithCredentials(
			credentials.NewSharedCredentials("", c.Project.Profile),
		)
	}
	return session.New(conf)
}

func debugRequest(obj fmt.Stringer) {
	if debug != "enable" {
		return
	}
	fmt.Println("[DEBUG] ", obj)
}
