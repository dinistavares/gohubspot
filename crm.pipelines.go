package gohubspot

import "fmt"

type PipelinesService service

type GetPipelinesResponse struct {
	Results *[]Pipeline `json:"results,omitempty"`
}

type Pipeline struct {
	Label        string          `json:"label,omitempty"`
	DisplayOrder int             `json:"displayOrder,omitempty"`
	ID           string          `json:"id,omitempty"`
	Stages       *[]PipelineStage `json:"stages,omitempty"`
	CreatedAt    string          `json:"createdAt,omitempty"`
	UpdatedAt    string          `json:"updatedAt,omitempty"`
	Archived     bool            `json:"archived,omitempty"`
}

type PipelineStage struct {
	Label            string                `json:"label,omitempty"`
	DisplayOrder     int                   `json:"displayOrder,omitempty"`
	Metadata        *PipelineStageMetadata `json:"metadata,omitempty"`
	ID               string                `json:"id,omitempty"`
	CreatedAt        string                `json:"createdAt,omitempty"`
	UpdatedAt        string                `json:"updatedAt,omitempty"`
	Archived         bool                  `json:"archived,omitempty"`
	WritePermissions string                `json:"writePermissions,omitempty"`
}

type PipelineStageMetadata struct {
	IsClosed    string `json:"isClosed,omitempty"`
	Probability string `json:"probability,omitempty"`
}


func (s *PipelinesService) GetObjectPipelines(objectType CRMObjectType) (*GetPipelinesResponse, error) {
	url := fmt.Sprintf("/crm/v3/pipelines/%s", objectType)
	res := new(GetPipelinesResponse)
	err := s.client.RunGet(url, res)
	return res, err
}
