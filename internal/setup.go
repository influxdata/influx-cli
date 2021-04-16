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
)

var InfiniteRetention = 0

func (c *CLI) Setup(ctx context.Context, client api.SetupApi, params *SetupParams) error {
	// First, check if setup is even allowed.
	checkReq := client.GetSetup(ctx)
	if c.TraceId != "" {
		checkReq.ZapTraceSpan(c.TraceId)
	}
	checkResp, _, err := client.GetSetupExecute(checkReq)
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
	setupReq := client.PostSetup(ctx).OnboardingRequest(setupBody)
	if c.TraceId != "" {
		setupReq.ZapTraceSpan(c.TraceId)
	}
	resp, _, err := client.PostSetupExecute(setupReq)
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
		return fmt.Errorf("failed to write new config to local path: %w", err)
	}

	// TODO: Print info.

	return nil
}

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
	if err := c.Banner("Welcome to InfluxDB 2.0!"); err != nil {
		return req, err
	}
	if params.Username == "" {
		req.Username, err = c.GetStringInput("Please type your primary username", "")
		if err != nil {
			return req, err
		}
	}
	if params.Password == "" {
		for {
			pass1, err := c.GetPassword("Please type your password", true)
			if err != nil {
				return req, err
			}
			// Don't bother with the length check the 2nd time, since we check equality to pass1.
			pass2, err := c.GetPassword("Please type your password again", false)
			if err != nil {
				return req, err
			}
			if pass1 == pass2 {
				req.Password = &pass1
				break
			}
			if err := c.Error("Passwords do not match"); err != nil {
				return req, err
			}
		}
	}
	if params.Org == "" {
		req.Org, err = c.GetStringInput("Please type your primary organization name", "")
		if err != nil {
			return req, err
		}
	}
	if params.Bucket == "" {
		req.Bucket, err = c.GetStringInput("Please type your primary bucket name", "")
		if err != nil {
			return req, err
		}
	}
	if params.Retention == "" {
		infiniteStr := strconv.Itoa(InfiniteRetention)
		for {
			rpStr, err := c.GetStringInput("Please type your retention period in hours, or 0 for infinite", infiniteStr)
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
			} else if err := c.Error("Retention period cannot be negative"); err != nil {
				return req, err
			}
		}
	}

	if confirmed := c.GetConfirm(func() string {
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
		return api.OnboardingRequest{}, fmt.Errorf("setup canceled")
	}

	return req, nil
}
