package agency

import (
	"github.com/hiagomf/bank-api/server/infraestructure/records/agency"
	"github.com/hiagomf/bank-api/server/utils"
)

type IAgency interface {
	SelectPaginated(parameters *utils.ParametrosRequisicao) (res *agency.AgencyPag, err error)
}
