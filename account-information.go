package gohubspot

type AccountInformationService service

type AccountDetails struct {
	PortalID              int      `json:"portalId,omitempty"`
	AccountType           string   `json:"accountType,omitempty"`
	TimeZone              string   `json:"timeZone,omitempty"`
	CompanyCurrency       string   `json:"companyCurrency,omitempty"`
	AdditionalCurrencies  []string `json:"additionalCurrencies,omitempty"`
	UtcOffset             string   `json:"utcOffset,omitempty"`
	UtcOffsetMilliseconds int      `json:"utcOffsetMilliseconds,omitempty"`
	UIDomain              string   `json:"uiDomain,omitempty"`
	DataHostingLocation   string   `json:"dataHostingLocation,omitempty"`
}

func (s *AccountInformationService) GetAccountDetails() (*AccountDetails, error) {
	url := "/account-info/v3/details"
	res := new(AccountDetails)
	err := s.client.RunGet(url, res)
	return res, err
}
