package responses

type MarketingPermissionCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID            string `json:"ObjectID"`
			ETag                string `json:"ETag"`
			BusinessPartnerID   string `json:"BusinessPartner_ID"`
			BusinessPartnerUUID string `json:"BusinessPartnerUUID"`
			CreationDateTime    string `json:"CreationDateTime"`
			LastChangeDateTime  string `json:"LastChangeDateTime"`
			AlternativeID       string `json:"AlternativeID"`
			EntityLastChangedOn string `json:"EntityLastChangedOn"`
		} `json:"results"`
	} `json:"d"`
}
