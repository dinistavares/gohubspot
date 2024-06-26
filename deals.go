package gohubspot

import "fmt"

type DealsService service

type Deal struct {
	PortalID     int                 `json:"portalId"`
	DealID       int                 `json:"dealId"`
	IsDeleted    bool                `json:"isDeleted"`
	Associations Associations        `json:"associations"`
	Properties   map[string]Property `json:"properties"`
}

type Associations struct {
	AssociatedVids       []int `json:"associatedVids"`
	AssociatedCompanyIds []any `json:"associatedCompanyIds"`
	AssociatedDealIds    []any `json:"associatedDealIds"`
	AssociatedTicketIds  []any `json:"associatedTicketIds"`
}

func (s *DealsService) GetById(dealID int) (*Deal, error) {
	url := fmt.Sprintf("/deals/v1/deal/%d", dealID)
	res := new(Deal)
	err := s.client.RunGet(url, res)
	return res, err
}
