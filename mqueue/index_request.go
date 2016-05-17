package mqueue

type IndexRequest struct {
	DownloadUrl   string `json:"downloadUrl"`
	ArtifactoryId string `json:"artifactoryId"`
	RepoKey       string `json:"repoKey"`
	Path          string `json:"path"`
	Sha           string `json:"sha"`
}
