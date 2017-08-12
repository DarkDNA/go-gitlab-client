package gogitlab

const (
	repo_url_file_raw = "/projects/:id/repository/files/:file_path/raw"
)

// RepoFileRaw gets the specified file at the specified ref.
func (g *Gitlab) RepoFileRaw(id, ref, filepath string) ([]byte, error) {
	url, opaque := g.ResourceUrlRaw(repo_url_file_raw, map[string]string{
		":id":        id,
		":file_path": filepath,
	})
	url += "&ref=" + ref

	contents, err := g.buildAndExecRequestRaw("GET", url, opaque, nil)
	return contents, err
}
