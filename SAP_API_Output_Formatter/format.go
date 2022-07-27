package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-marketing-permission-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToMarketingPermissionCollection(raw []byte, l *logger.Logger) ([]MarketingPermissionCollection, error) {
	pm := &responses.MarketingPermissionCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to MarketingPermissionCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	marketingPermissionCollection := make([]MarketingPermissionCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		marketingPermissionCollection = append(marketingPermissionCollection, MarketingPermissionCollection{
			ObjectID:            data.ObjectID,
			ETag:                data.ETag,
			BusinessPartnerID:   data.BusinessPartnerID,
			BusinessPartnerUUID: data.BusinessPartnerUUID,
			CreationDateTime:    data.CreationDateTime,
			LastChangeDateTime:  data.LastChangeDateTime,
			AlternativeID:       data.AlternativeID,
			EntityLastChangedOn: data.EntityLastChangedOn,
		})
	}

	return marketingPermissionCollection, nil
}