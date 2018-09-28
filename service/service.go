package service
import (
	log "github.com/sirupsen/logrus"

	"github.com/gordonrehling2/certavs/server/db"
	"github.com/gordonrehling2/certavs/entities"
)

type IRfeService interface {
	RfeList() *[]entities.RFE

}

type RfeService struct {
	db db.IDB
}

func NewRfeService(db db.IDB) *RfeService {
	return &RfeService{
		db,
	}
}

func (r RfeService) RfeList() *[]entities.RFE {
	rows, err := r.db.Query("SELECT id, description FROM rfe;")
	if err != nil {
		log.Fatal(err)
	}

	var refId int32
	var desc string

	var result []entities.RFE

	for rows.Next() {
		err := rows.Scan(&refId, &desc)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, entities.RFE {RfeId:  refId, Description:desc})
	}

	return &result
}