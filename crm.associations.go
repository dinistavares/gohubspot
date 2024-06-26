package gohubspot

import "fmt"

type CRMAssociationsService service

// Used to define association definitions (Reference: https://legacydocs.hubspot.com/docs/methods/crm-associations/crm-associations-overview)
type CRMAssociationType int

const (
  // Defines the association definition 'Contact to company'
  CRMAssociationContactToCompany              CRMAssociationType = 1

  // Defines the association definition 'Company to contact (default)'
  CRMAssociationCompanyToContactDefault       CRMAssociationType = 2

  // Defines the association definition 'Company to contact (all labels)'
  CRMAssociationCompanyToContactAllLabels     CRMAssociationType = 280

  // Defines the association definition 'Deal to contact'
  CRMAssociationDealToContact                 CRMAssociationType = 3

  // Defines the association definition 'Contact to deal'
  CRMAssociationContactToDeal                 CRMAssociationType = 4

  // Defines the association definition 'Deal to company'
  CRMAssociationDealToCompany                 CRMAssociationType = 5

  // Defines the association definition 'Company to deal'
  CRMAssociationCompanyToDeal                 CRMAssociationType = 6

  // Defines the association definition 'Company to engagement'
  CRMAssociationCompanyToEngagement           CRMAssociationType = 7

  // Defines the association definition 'Engagement to company'
  CRMAssociationEngagementToCompany           CRMAssociationType = 8

  // Defines the association definition 'Contact to engagement'
  CRMAssociationContactToEngagement           CRMAssociationType = 9

  // Defines the association definition 'Engagement to contact'
  CRMAssociationEngagementToContact           CRMAssociationType = 10

  // Defines the association definition 'Deal to engagement'
  CRMAssociationDealToEngagement              CRMAssociationType = 11

  // Defines the association definition 'Engagement to deal'
  CRMAssociationEngagementToDeal              CRMAssociationType = 12

  // Defines the association definition 'Parent company to child company'
  CRMAssociationParentCompanyToChildCompany   CRMAssociationType = 13

  // Defines the association definition 'Child company to parent company'
  CRMAssociationChildCompanyToParentCompany   CRMAssociationType = 14

  // Defines the association definition 'Contact to ticket'
  CRMAssociationContactToTicket               CRMAssociationType = 15

  // Defines the association definition 'Ticket to contact'
  CRMAssociationTicketToContact               CRMAssociationType = 16

  // Defines the association definition 'Ticket to engagement'
  CRMAssociationTicketToEngagement            CRMAssociationType = 17

  // Defines the association definition 'Engagement to ticket'
  CRMAssociationEngagementToTicket            CRMAssociationType = 18

  // Defines the association definition 'Deal to line item'
  CRMAssociationDealToLineItem                CRMAssociationType = 19

  // Defines the association definition 'Line item to deal'
  CRMAssociationLineItemToDeal                CRMAssociationType = 20

  // Defines the association definition 'Company to ticket'
  CRMAssociationCompanyToTicket               CRMAssociationType = 25

  // Defines the association definition 'Ticket to company'
  CRMAssociationTicketToCompany               CRMAssociationType = 26

  // Defines the association definition 'Deal to ticket'
  CRMAssociationDealToTicket                  CRMAssociationType = 27

  // Defines the association definition 'Ticket to deal'
  CRMAssociationTicketToDeal                  CRMAssociationType = 28
)

type CRMAssocition struct {
	Results []int `json:"results"`
	HasMore bool  `json:"hasMore"`
	Offset  int   `json:"offset"`
}

func (s *CRMAssociationsService) GetAssociationByID(associationType CRMAssociationType, objectID int) (*CRMAssocition, error) {
	url := fmt.Sprintf("/crm-associations/v1/associations/%d/HUBSPOT_DEFINED/%d", objectID, associationType)
	res := new(CRMAssocition)
	err := s.client.RunGet(url, res)
	return res, err
}
