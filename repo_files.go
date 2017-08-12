package gogitlab

import "encoding/json"

const (
	repo_url_file = "/projects/:id/repository/files/:file_path"
)

type RepoFile struct {
	FileName     string `json:"file_name"`
	FilePath     string `json:"file_path"`
	Size         int    `json:"size"`
	Encoding     string `json:"encoding"`
	Content      []byte `json:"content"`
	Ref          string `json:"ref"`
	BlobID       string `json:"blob_id"`
	CommitID     string `json:"commit_id"`
	LastCommitID string `json:"last_commit_id"`
}

// RepoFile gets the specified file at the specified ref.
func (g *Gitlab) RepoFile(id, ref, filepath string) (*RepoFile, error) {
	url, opaque := g.ResourceUrlRaw(repo_url_file, map[string]string{
		":id":        id,
		":file_path": filepath,
	})
	url += "&ref=" + ref

	var fileInfo RepoFile

	contents, err := g.buildAndExecRequestRaw("GET", url, opaque, nil)
	if err == nil {
		err = json.Unmarshal(contents, &fileInfo)
	}

	return &fileInfo, err
}
