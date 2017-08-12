package gogitlab

import (
	"encoding/json"
)

const (
	project_url_hooks = "/projects/:id/hooks"          // Get list of project hooks
	project_url_hook  = "/projects/:id/hooks/:hook_id" // Get single project hook
)

type Hook struct {
	Id           int    `json:"id,omitempty"`
	Url          string `json:"url,omitempty"`
	CreatedAtRaw string `json:"created_at,omitempty"`
}

type HookEvents struct {
	Push          bool `json:"push_events"`
	Issues        bool `json:"issue_events"`
	MergeRequests bool `json:"merge_request_events"`
	TagPush       bool `json:"tag_push_events"`
	Note          bool `json:"note_events"`
	Job           bool `json:"job_events"`
	Wiki          bool `json:"wiki_events"`
}

/*
Get list of project hooks.

    GET /projects/:id/hooks

Parameters:

    id The ID of a project

*/
func (g *Gitlab) ProjectHooks(id string) ([]*Hook, error) {

	url, opaque := g.ResourceUrlRaw(project_url_hooks, map[string]string{":id": id})

	var err error
	var hooks []*Hook

	contents, err := g.buildAndExecRequestRaw("GET", url, opaque, nil)
	if err != nil {
		return hooks, err
	}

	err = json.Unmarshal(contents, &hooks)

	return hooks, err
}

/*
Get single project hook.

    GET /projects/:id/hooks/:hook_id

Parameters:

    id      The ID of a project
    hook_id The ID of a hook

*/
func (g *Gitlab) ProjectHook(id, hook_id string) (*Hook, error) {

	url, opaque := g.ResourceUrlRaw(project_url_hook, map[string]string{
		":id":      id,
		":hook_id": hook_id,
	})

	var err error
	hook := new(Hook)

	contents, err := g.buildAndExecRequestRaw("GET", url, opaque, nil)
	if err != nil {
		return hook, err
	}

	err = json.Unmarshal(contents, &hook)

	return hook, err
}

/*
Add new project hook.

    POST /projects/:id/hooks

Parameters:

    id                    The ID or NAMESPACE/PROJECT_NAME of a project
    hook_url              The hook URL
    push_events           Trigger hook on push events
    issues_events         Trigger hook on issues events
    merge_requests_events Trigger hook on merge_requests events

*/
func (g *Gitlab) AddProjectHook(id, hook_url string, events HookEvents) error {

	url, opaque := g.ResourceUrlRaw(project_url_hooks, map[string]string{":id": id})

	var err error

	body, _ := json.Marshal(struct {
		HookEvents `json:",inline"`
		URL        string `json:"url"`
	}{URL: hook_url, HookEvents: events})
	_, err = g.buildAndExecRequestRaw("POST", url, opaque, body)

	return err
}

/*
Edit existing project hook.

    PUT /projects/:id/hooks/:hook_id

Parameters:

    id                    The ID or NAMESPACE/PROJECT_NAME of a project
    hook_id               The ID of a project hook
    hook_url              The hook URL
    push_events           Trigger hook on push events
    issues_events         Trigger hook on issues events
    merge_requests_events Trigger hook on merge_requests events

*/
func (g *Gitlab) EditProjectHook(id, hook_id, hook_url string, events HookEvents) error {

	url, opaque := g.ResourceUrlRaw(project_url_hook, map[string]string{
		":id":      id,
		":hook_id": hook_id,
	})

	var err error

	body, _ := json.Marshal(struct {
		HookEvents `json:",inline"`
		URL        string `json:"url"`
	}{URL: hook_url, HookEvents: events})
	_, err = g.buildAndExecRequestRaw("PUT", url, opaque, body)

	return err
}

/*
Remove hook from project.

    DELETE /projects/:id/hooks/:hook_id

Parameters:

    id      The ID or NAMESPACE/PROJECT_NAME of a project
    hook_id The ID of hook to delete

*/
func (g *Gitlab) RemoveProjectHook(id, hook_id string) error {

	url, opaque := g.ResourceUrlRaw(project_url_hook, map[string]string{
		":id":      id,
		":hook_id": hook_id,
	})

	var err error

	_, err = g.buildAndExecRequestRaw("DELETE", url, opaque, nil)

	return err
}
