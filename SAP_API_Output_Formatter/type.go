package sap_api_output_formatter

type MarketingPermission struct {
	ConnectionKey           string `json:"connection_key"`
	Result                  bool   `json:"result"`
	RedisKey                string `json:"redis_key"`
	Filepath                string `json:"filepath"`
	APISchema               string `json:"api_schema"`
	MarketingPermissionCode string `json:"marketing_permission_code"`
	Deleted                 bool   `json:"deleted"`
}

type MarketingPermissionCollection struct {
	ObjectID            string `json:"ObjectID"`
	ETag                string `json:"ETag"`
	BusinessPartnerID   string `json:"BusinessPartner_ID"`
	BusinessPartnerUUID string `json:"BusinessPartnerUUID"`
	CreationDateTime    string `json:"CreationDateTime"`
	LastChangeDateTime  string `json:"LastChangeDateTime"`
	AlternativeID       string `json:"AlternativeID"`
	EntityLastChangedOn string `json:"EntityLastChangedOn"`
}