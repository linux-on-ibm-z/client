// +build ece408ProjectMode

package client

import (
	"time"

	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/rai-project/auth/provider"
	"github.com/rai-project/config"
	"github.com/rai-project/database/mongodb"
	"github.com/spf13/cast"
)

func (c *Client) RecordJob() error {

	if c.jobBody == nil {
		return errors.New("ranking uninitialized")
	}

	body, ok := c.jobBody.(*Ece408JobResponseBody)
	if !ok {
		panic("invalid job type")
	}

	// body.ID = ""
	body.UpdatedAt = time.Now()
	body.IsSubmission = cast.ToBool(c.options.ctx.Value(isSubmissionKey{}))
	body.SubmissionTag = cast.ToString(c.options.ctx.Value(submissionKindKey{}))

	prof, err := provider.New()
	user := prof.Info()
	body.Username = user.Username
	body.UserAccessKey = user.AccessKey

	body.Teamname, err = FindTeamName(body.Username)
	if err != nil && body.IsSubmission {
		color.Red("no team name found.\n")
		body.Teamname = user.Team.Name
	}

	db, err := mongodb.NewDatabase(config.App.Name)
	if err != nil {
		return err
	}
	defer db.Close()

	col, err := NewEce408JobResponseBodyCollection(db)
	if err != nil {
		return err
	}
	defer col.Close()

	err = col.Insert(body)
	if err != nil {
		log.WithError(err).Error("Failed to insert job record:", body)
		return err
	}

	return nil
}
