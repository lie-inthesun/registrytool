package harbor

import "time"

type projectsResp []struct {
	ChartCount         int               `json:"chart_count"`
	CreationTime       time.Time         `json:"creation_time"`
	CurrentUserRoleId  int               `json:"current_user_role_id"`
	CurrentUserRoleIds []int             `json:"current_user_role_ids"`
	Metadata           map[string]string `json:"metadata"`
	Name               string            `json:"name"`
	OwnerId            int               `json:"owner_id"`
	OwnerName          string            `json:"owner_name"`
	ProjectId          int               `json:"project_id"`
	RepoCount          int               `json:"repo_count"`
	UpdateTime         time.Time         `json:"update_time"`
	//CveAllowlist       struct {} `json:"cve_allowlist"`
}

type repositoriesResp []struct {
	Id            int       `json:"id"`
	Name          string    `json:"name"`
	ProjectID     int       `json:"project_id"`
	ArtifactCount int       `json:"artifact_count"`
	PullCount     int       `json:"pull_count"`
	CreationTime  time.Time `json:"creation_time"`
	UpdateTime    time.Time `json:"update_time"`
}

type imageDetailResp struct {
	Digest            string               `json:"digest"`
	ExtraAttrs        imageDetailExtraAttr `json:"extra_attrs"`
	Icon              string               `json:"icon"`
	Id                int                  `json:"id"`
	ManifestMediaType string               `json:"manifest_media_type"`
	MediaType         string               `json:"media_type"`
	ProjectId         int                  `json:"project_id"`
	PullTime          time.Time            `json:"pull_time"`
	PushTime          time.Time            `json:"push_time"`
	RepositoryId      int                  `json:"repository_id"`
	Size              int                  `json:"size"`
	Tags              []imageDetailTag     `json:"tags"`
	//Labels		[]struct{}
	//References	struct{}
}

type imageDetailExtraAttr struct {
	Architecture string    `json:"architecture"`
	Author       string    `json:"author"`
	Created      time.Time `json:"created"`
	Os           string    `json:"os"`
	//Config	struct{}
}

type imageDetailTag struct {
	ArtifactId   int       `json:"artifact_id"`
	Id           int       `json:"id"`
	Immutable    bool      `json:"immutable"`
	Name         string    `json:"name"`
	PullTime     time.Time `json:"pull_time"`
	PushTime     time.Time `json:"push_time"`
	RepositoryId int       `json:"repository_id"`
	Signed       bool      `json:"signed"`
}
