package internal

import (
	"context"
	"errors"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/influxdata/influx-cli/v2/internal/api"
	"github.com/influxdata/influx-cli/v2/internal/config"
	"github.com/influxdata/influx-cli/v2/internal/duration"
)

type SetupParams struct {
	Username   string
	Password   string
	AuthToken  string
	Org        string
	Bucket     string
	Retention  string
	Force      bool
	ConfigName string
}

var (
	ErrPasswordIsTooShort = errors.New("password is too short")
	ErrAlreadySetUp       = errors.New("instance has already been set up")
	ErrConfigNameRequired = errors.New("config name is required if you already have existing configs")
	ErrSetupCanceled      = errors.New("setup was canceled")
)

const MinPasswordLen = 8

func (c *CLI) Setup(ctx context.Context, client api.SetupApi, params *SetupParams) error {
	// Ensure we'll be able to write onboarding results to local config.
	// Do this first so we catch any problems before modifying state on the server side.
	if err := c.validateNoNameCollision(params.ConfigName); err != nil {
		return err
	}

	// Check if setup is even allowed.
	checkResp, _, err := client.GetSetup(ctx).Execute()
	if err != nil {
		return fmt.Errorf("failed to check if already set up: %w", err)
	}
	if checkResp.Allowed == nil || !*checkResp.Allowed {
		return ErrAlreadySetUp
	}

	// Initialize the server.
	setupBody, err := c.onboardingRequest(params)
	if err != nil {
		return err
	}
	resp, _, err := client.PostSetup(ctx).OnboardingRequest(setupBody).Execute()
	if err != nil {
		return fmt.Errorf("failed to setup instance: %w", err)
	}

	cfg := config.Config{
		Name:  config.DefaultConfig.Name,
		Host:  c.ActiveConfig.Host,
		Token: *resp.Auth.Token,
		Org:   resp.Org.Name,
	}
	if params.ConfigName != "" {
		cfg.Name = params.ConfigName
	}

	if _, err := c.ConfigService.CreateConfig(cfg); err != nil {
		return fmt.Errorf("setup succeeded, but failed to write new config to local path: %w", err)
	}

	if c.PrintAsJSON {
		return c.PrintJSON(map[string]interface{}{
			"user":         resp.User.Name,
			"organization": resp.Org.Name,
			"bucket":       resp.Bucket.Name,
		})
	}

	return c.PrintTable([]string{"User", "Organization", "Bucket"}, map[string]interface{}{
		"User":         resp.User.Name,
		"Organization": resp.Org.Name,
		"Bucket":       resp.Bucket.Name,
	})
}

// validateNoNameCollision checks that we will be able to write onboarding results to local config:
//   - If a custom name was given, check that it doesn't collide with existing config
//   - If no custom name was given, check that we don't already have configs
func (c *CLI) validateNoNameCollision(configName string) error {
	existingConfigs, err := c.ConfigService.ListConfigs()
	if err != nil {
		return fmt.Errorf("error checking existing configs: %w", err)
	}
	if len(existingConfigs) == 0 {
		return nil
	}

	// If there are existing configs then require that a name be
	// specified in order to distinguish this new config from what's
	// there already.
	if configName == "" {
		return ErrConfigNameRequired
	}
	if _, ok := existingConfigs[configName]; ok {
		return fmt.Errorf("config name %q already exists", configName)
	}

	return nil
}

// onboardingRequest constructs a request body for the onboarding API.
// Unless the 'force' parameter is set, the user will be prompted to enter any missing information
// and to confirm the final request parameters.
func (c *CLI) onboardingRequest(params *SetupParams) (req api.OnboardingRequest, err error) {
	if (params.Force || params.Password != "") && len(params.Password) < MinPasswordLen {
		return req, ErrPasswordIsTooShort
	}

	// Populate the request with CLI args.
	req.Username = params.Username
	req.Org = params.Org
	req.Bucket = params.Bucket
	if params.Password != "" {
		req.Password = &params.Password
	}
	if params.AuthToken != "" {
		req.Token = &params.AuthToken
	}
	rpSecs := int64(InfiniteRetention)
	if params.Retention != "" {
		dur, err := duration.RawDurationToTimeDuration(params.Retention)
		if err != nil {
			return req, fmt.Errorf("failed to parse %q: %w", params.Retention, err)
		}
		secs, nanos := math.Modf(dur.Seconds())
		if nanos > 0 {
			return req, fmt.Errorf("retention policy %q is too precise, must be divisible by 1s", params.Retention)
		}
		rpSecs = int64(secs)
	}
	req.RetentionPeriodSeconds = &rpSecs

	if params.Force {
		return req, nil
	}

	// Ask the user for any missing information.
	if err := c.StdIO.Banner("Welcome to InfluxDB 2.0!"); err != nil {
		return req, err
	}
	if params.Username == "" {
		req.Username, err = c.StdIO.GetStringInput("Please type your primary username", "")
		if err != nil {
			return req, err
		}
	}
	if params.Password == "" {
		for {
			pass1, err := c.StdIO.GetPassword("Please type your password", MinPasswordLen)
			if err != nil {
				return req, err
			}
			// Don't bother with the length check the 2nd time, since we check equality to pass1.
			pass2, err := c.StdIO.GetPassword("Please type your password again", 0)
			if err != nil {
				return req, err
			}
			if pass1 == pass2 {
				req.Password = &pass1
				break
			}
			if err := c.StdIO.Error("Passwords do not match"); err != nil {
				return req, err
			}
		}
	}
	if params.Org == "" {
		req.Org, err = c.StdIO.GetStringInput("Please type your primary organization name", "")
		if err != nil {
			return req, err
		}
	}
	if params.Bucket == "" {
		req.Bucket, err = c.StdIO.GetStringInput("Please type your primary bucket name", "")
		if err != nil {
			return req, err
		}
	}
	if params.Retention == "" {
		infiniteStr := strconv.Itoa(InfiniteRetention)
		for {
			rpStr, err := c.StdIO.GetStringInput("Please type your retention period in hours, or 0 for infinite", infiniteStr)
			if err != nil {
				return req, err
			}
			rp, err := strconv.Atoi(rpStr)
			if err != nil {
				return req, err
			}
			if rp >= 0 {
				rpSeconds := int64((time.Duration(rp) * time.Hour).Seconds())
				req.RetentionPeriodSeconds = &rpSeconds
				break
			} else if err := c.StdIO.Error("Retention period cannot be negative"); err != nil {
				return req, err
			}
		}
	}

	if confirmed := c.StdIO.GetConfirm(func() string {
		rp := "infinite"
		if req.RetentionPeriodSeconds != nil && *req.RetentionPeriodSeconds > 0 {
			rp = (time.Duration(*req.RetentionPeriodSeconds) * time.Second).String()
		}
		return fmt.Sprintf(`Setup with these parameters?
  Username:          %s
  Organization:      %s
  Bucket:            %s
  Retention Period:  %s
`, req.Username, req.Org, req.Bucket, rp)
	}()); !confirmed {
		return api.OnboardingRequest{}, ErrSetupCanceled
	}

	return req, nil
}
