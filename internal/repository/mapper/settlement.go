package mapper

import (
	"smolathon/internal/entity"
	"smolathon/internal/repository/models"
)

func SettlementSliceMapFromDb(settlementsModels []models.Settlement) []entity.Settlement {
	var settlements = make([]entity.Settlement, 0, len(settlementsModels))
	for _, settlement := range settlementsModels {
		settlements = append(settlements, SettlementMapFromDb(settlement))
	}
	return settlements
}

func SettlementMapFromDb(settlement models.Settlement) entity.Settlement {
	return entity.Settlement{
		Id:        settlement.Id,
		Name:      settlement.Name,
		Latitude:  settlement.Location.P.X,
		Longitude: settlement.Location.P.Y,
	}
}
