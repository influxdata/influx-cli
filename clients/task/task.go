package task

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/influxdata/influx-cli/v2/api"
	"github.com/influxdata/influx-cli/v2/clients"
)

type Client struct {
	clients.CLI
	api.TasksApi
	// AllowEmptyOrg will be useful for Kapacitor which doesn't use org / orgID
	AllowEmptyOrg bool
}

type CreateParams struct {
	clients.OrgParams
	FluxQuery    string
	Name         string
	Every        string
	Cron         string
	ScriptID     string
	ScriptParams map[string]interface{}
}

type NameOrID struct {
	Name string
	ID   string
}

func (n NameOrID) NameOrNil() *string {
	if n.Name == "" {
		return nil
	}
	return &n.Name
}

func (n NameOrID) IDOrNil() *string {
	if n.ID == "" {
		return nil
	}
	return &n.ID
}

func addOrg(n NameOrID, g api.ApiGetTasksRequest) api.ApiGetTasksRequest {
	if n.ID != "" {
		return g.OrgID(n.ID)
	}
	if n.Name != "" {
		return g.Org(n.Name)
	}
	return g
}

func (c Client) getOrg(params *clients.OrgParams) (NameOrID, error) {
	if params.OrgID != "" {
		return NameOrID{ID: params.OrgID}, nil
	}
	if params.OrgName != "" {
		return NameOrID{Name: params.OrgName}, nil
	}
	if c.ActiveConfig.Org != "" {
		return NameOrID{Name: c.ActiveConfig.Org}, nil
	}
	if c.AllowEmptyOrg {
		return NameOrID{}, nil
	}
	return NameOrID{}, clients.ErrMustSpecifyOrg
}

func (c Client) Create(ctx context.Context, params *CreateParams) error {
	org, err := c.getOrg(&params.OrgParams)
	if err != nil {
		return err
	}

	scriptID := &params.ScriptID
	if len(params.ScriptID) == 0 {
		scriptID = nil
	}

	scriptParams := &params.ScriptParams
	if len(params.ScriptParams) == 0 {
		scriptParams = nil
	}

	createRequest := api.TaskCreateRequest{
		Flux:             &params.FluxQuery,
		OrgID:            org.IDOrNil(),
		Org:              org.NameOrNil(),
		Name:             &params.Name,
		Cron:             &params.Cron,
		ScriptID:         scriptID,
		ScriptParameters: scriptParams,
	}

	// The FluxQuery can also contain the "every" field, so we only want to override if it is actually defined.
	if params.Every != "" {
		createRequest.Every = &params.Every
	}

	task, err := c.PostTasks(ctx).TaskCreateRequest(createRequest).Execute()
	if err != nil {
		return err
	}
	return c.printTasks(taskPrintOpts{
		task: &task,
	})
}

type FindParams struct {
	clients.OrgParams
	TaskID   string
	UserID   string
	ScriptID string
	Limit    int
}

func (c Client) Find(ctx context.Context, params *FindParams) error {
	if params.Limit < 1 {
		return fmt.Errorf("must specify a positive limit, not %d", params.Limit)
	}

	var tasks []api.Task
	// If we get an id, just find the one task
	if params.TaskID != "" {
		task, err := c.GetTasksID(ctx, params.TaskID).Execute()
		if err != nil {
			return err
		}
		tasks = append(tasks, task)
	} else {
		org, err := c.getOrg(&params.OrgParams)
		if err != nil {
			return err
		}
		// filter on all tasks
		if params.Limit > math.MaxInt32 {
			return fmt.Errorf("limit too large %d > %d", params.Limit, math.MaxInt32)
		}
		getTask := c.GetTasks(ctx).Limit(int32(params.Limit))
		getTask = addOrg(org, getTask)
		if params.UserID != "" {
			getTask = getTask.User(params.UserID)
		}
		if params.ScriptID != "" {
			getTask = getTask.ScriptID(params.ScriptID)
		}
		tasksResult, err := getTask.Execute()
		if err != nil {
			return err
		}
		tasks = *tasksResult.Tasks
	}
	return c.printTasks(taskPrintOpts{
		tasks: tasks,
	})
}

func (c Client) appendRuns(ctx context.Context, prev []api.Run, taskID string, filter RunFilter) ([]api.Run, error) {
	if filter.Limit < 1 {
		return nil, fmt.Errorf("must specify a positive run limit, not %d", filter.Limit)
	}
	if filter.Limit > math.MaxInt32 {
		return nil, fmt.Errorf("limit too large %d > %d", filter.Limit, math.MaxInt32)
	}
	getRuns := c.GetTasksIDRuns(ctx, taskID).Limit(int32(filter.Limit))
	if filter.After != "" {
		afterTime, err := time.Parse(time.RFC3339, filter.After)
		if err != nil {
			return nil, err
		}
		getRuns = getRuns.AfterTime(afterTime)
	}
	if filter.Before != "" {
		beforeTime, err := time.Parse(time.RFC3339, filter.Before)
		if err != nil {
			return nil, err
		}
		getRuns = getRuns.BeforeTime(beforeTime)
	}
	runs, err := getRuns.Execute()
	if err != nil {
		return nil, err
	}
	for _, run := range *runs.Runs {
		if filter.Status == "" {
			prev = append(prev, run)
		} else if run.Status != nil && *run.Status == filter.Status {
			prev = append(prev, run)
		}
	}
	return prev, nil
}

type RunFilter struct {
	After  string
	Before string
	Limit  int
	Status string
}

type RetryFailedParams struct {
	clients.OrgParams
	TaskID    string
	DryRun    bool
	TaskLimit int
	RunFilter RunFilter
}

func (c Client) retryRun(ctx context.Context, run api.Run, dryRun bool) error {
	// Note that this output does not respect json flag, in line with original influx cli
	// The server should fill in the empty id's so this shouldn't happen
	if run.Id == nil {
		_ = c.StdIO.Error("skipping empty run id from influxdb")
		return nil
	}
	if run.TaskID == nil {
		_ = c.StdIO.Error("skipping empty task id from influxdb")
		return nil
	}
	if dryRun {
		_, _ = fmt.Fprintf(c.StdIO, "Would retry for %s run for Task %s.\n", *run.Id, *run.TaskID)
	} else {
		newRun, err := c.PostTasksIDRunsIDRetry(ctx, *run.TaskID, *run.Id).Execute()
		if err != nil {
			return err
		}
		_, _ = fmt.Fprintf(c.StdIO, "Retry for task %s's run %s queued as run %s.\n", *run.TaskID, *run.Id, *newRun.Id)
	}
	return nil
}

func (c Client) RetryFailed(ctx context.Context, params *RetryFailedParams) error {
	if params.TaskLimit < 1 {
		return fmt.Errorf("must specify a positive task limit, not %d", params.TaskLimit)
	}
	var taskIds []string
	if params.TaskID != "" {
		taskIds = []string{params.TaskID}
	} else {
		org, err := c.getOrg(&params.OrgParams)
		if err != nil {
			return err
		}

		if params.TaskLimit > math.MaxInt32 {
			return fmt.Errorf("limit too large %d > %d", params.TaskLimit, math.MaxInt32)
		}
		getTask := c.GetTasks(ctx).Limit(int32(params.TaskLimit))
		getTask = addOrg(org, getTask)
		tasks, err := getTask.Execute()
		if err != nil {
			return err
		}
		taskIds = make([]string, 0, len(*tasks.Tasks))
		for _, t := range *tasks.Tasks {
			taskIds = append(taskIds, t.Id)
		}
	}
	var failedRuns []api.Run
	for _, taskId := range taskIds {
		var err error
		runFilter := params.RunFilter
		runFilter.Status = "failed"
		failedRuns, err = c.appendRuns(ctx, failedRuns, taskId, runFilter)
		if err != nil {
			return err
		}
	}

	for _, run := range failedRuns {
		err := c.retryRun(ctx, run, params.DryRun)
		if err != nil {
			return err
		}
	}
	if params.DryRun {
		_, _ = fmt.Fprintf(c.StdIO, `Dry run complete. Found %d tasks with a total of %d runs to be retried
Rerun without '--dry-run' to execute
`, len(taskIds), len(failedRuns))
	}
	return nil
}

type UpdateParams struct {
	FluxQuery    string
	TaskID       string
	Status       string
	ScriptID     string
	ScriptParams map[string]interface{}
}

func (c Client) Update(ctx context.Context, params *UpdateParams) error {
	var flux *string
	if params.FluxQuery != "" {
		flux = &params.FluxQuery
	}
	var status *api.TaskStatusType
	if params.Status != "" {
		var s api.TaskStatusType
		err := s.UnmarshalJSON([]byte(fmt.Sprintf("%q", params.Status)))
		if err != nil {
			return err
		}
		status = &s
	}

	scriptID := &params.ScriptID
	if len(params.ScriptID) == 0 {
		scriptID = nil
	}

	scriptParams := &params.ScriptParams
	if len(params.ScriptParams) == 0 {
		scriptParams = nil
	}

	task, err := c.PatchTasksID(ctx, params.TaskID).TaskUpdateRequest(api.TaskUpdateRequest{
		Status:           status,
		Flux:             flux,
		ScriptID:         scriptID,
		ScriptParameters: scriptParams,
	}).Execute()
	if err != nil {
		return err
	}
	return c.printTasks(taskPrintOpts{
		task: &task,
	})
}

type DeleteParams struct {
	TaskID string
}

func (c Client) Delete(ctx context.Context, params *DeleteParams) error {
	task, err := c.GetTasksID(ctx, params.TaskID).Execute()
	if err != nil {
		return fmt.Errorf("while finding task: %w", err)
	}
	err = c.DeleteTasksID(ctx, params.TaskID).Execute()
	if err != nil {
		return fmt.Errorf("while deleting: %w", err)
	}
	return c.printTasks(taskPrintOpts{
		task: &task,
	})
}

type taskPrintOpts struct {
	task  *api.Task
	tasks []api.Task
}

func derefOrEmpty(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func (c Client) printTasks(printOpts taskPrintOpts) error {
	if c.PrintAsJSON {
		var v interface{} = printOpts.tasks
		if printOpts.task != nil {
			v = printOpts.task
		}
		return c.PrintJSON(v)
	}

	headers := []string{
		"ID",
		"Name",
		"Organization ID",
		"Organization",
		"Status",
		"Every",
		"Cron",
		"ScriptID",
	}

	if printOpts.task != nil {
		printOpts.tasks = append(printOpts.tasks, *printOpts.task)
	}

	var rows []map[string]interface{}
	for _, t := range printOpts.tasks {
		row := map[string]interface{}{
			"ID":              t.Id,
			"Name":            t.Name,
			"Organization ID": t.OrgID,
			"Organization":    derefOrEmpty(t.Org),
			"Status":          derefOrEmpty((*string)(t.Status)),
			"Every":           derefOrEmpty(t.Every),
			"Cron":            derefOrEmpty(t.Cron),
			"ScriptID":        derefOrEmpty(t.ScriptID),
		}
		rows = append(rows, row)
	}

	return c.PrintTable(headers, rows...)
}

type LogFindParams struct {
	TaskID string
	RunID  string
}

func (c Client) FindLogs(ctx context.Context, params *LogFindParams) error {
	var logs api.Logs
	if params.RunID != "" {
		var err error
		logs, err = c.GetTasksIDRunsIDLogs(ctx, params.TaskID, params.RunID).Execute()
		if err != nil {
			return err
		}
	} else {
		var err error
		logs, err = c.GetTasksIDLogs(ctx, params.TaskID).Execute()
		if err != nil {
			return err
		}
	}
	if logs.Events == nil {
		return c.printLogs(nil)
	}
	return c.printLogs(*logs.Events)
}

func (c Client) printLogs(logs []api.LogEvent) error {
	if c.PrintAsJSON {
		var v interface{} = logs
		return c.PrintJSON(v)
	}

	headers := []string{
		"RunID",
		"Time",
		"Message",
	}

	var rows []map[string]interface{}
	for _, l := range logs {
		row := map[string]interface{}{
			"RunID":   derefOrEmpty(l.RunID),
			"Time":    l.Time,
			"Message": derefOrEmpty(l.Message),
		}
		rows = append(rows, row)
	}

	return c.PrintTable(headers, rows...)
}

type RunFindParams struct {
	TaskID string
	RunID  string
	Filter RunFilter
}

func (c Client) FindRuns(ctx context.Context, params *RunFindParams) error {
	if params.Filter.Limit < 1 {
		return fmt.Errorf("must specify a positive run limit, not %d", params.Filter.Limit)
	}

	runs := make([]api.Run, 0)
	if params.RunID != "" {
		run, err := c.GetTasksIDRunsID(ctx, params.TaskID, params.RunID).Execute()
		if err != nil {
			return err
		}
		runs = append(runs, run)
	} else {
		var err error
		runs, err = c.appendRuns(ctx, runs, params.TaskID, params.Filter)
		if err != nil {
			return err
		}
	}

	return c.printRuns(runs)
}

func (c Client) printRuns(runs []api.Run) error {
	if c.PrintAsJSON {
		var v interface{} = runs
		return c.PrintJSON(v)
	}

	headers := []string{
		"ID",
		"TaskID",
		"Status",
		"ScheduledFor",
		"StartedAt",
		"FinishedAt",
		"RequestedAt",
	}

	derefAndFormat := func(t *time.Time, layout string) string {
		if t == nil {
			return ""
		}
		return t.Format(layout)
	}

	var rows []map[string]interface{}
	for _, r := range runs {
		row := map[string]interface{}{
			"ID":           derefOrEmpty(r.Id),
			"TaskID":       derefOrEmpty(r.TaskID),
			"Status":       derefOrEmpty(r.Status),
			"ScheduledFor": derefAndFormat(r.ScheduledFor, time.RFC3339),
			"StartedAt":    derefAndFormat(r.StartedAt, time.RFC3339Nano),
			"FinishedAt":   derefAndFormat(r.FinishedAt, time.RFC3339Nano),
			"RequestedAt":  derefAndFormat(r.RequestedAt, time.RFC3339Nano),
		}
		rows = append(rows, row)
	}

	return c.PrintTable(headers, rows...)
}

type RunRetryParams struct {
	TaskID string
	RunID  string
}

func (c Client) RetryRun(ctx context.Context, params *RunRetryParams) error {
	newRun, err := c.PostTasksIDRunsIDRetry(ctx, params.TaskID, params.RunID).Execute()
	if err != nil {
		return err
	}
	_, _ = fmt.Fprintf(c.StdIO, "Retry for task %s's run %s queued as run %s.\n", params.TaskID, params.RunID, *newRun.Id)
	return nil
}
