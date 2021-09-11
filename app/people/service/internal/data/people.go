package data

import (
	"context"

	"github.com/genglsh/go-trainning-map/app/people/service/internal/data/ent/people"

	"github.com/pkg/errors"

	"github.com/genglsh/go-trainning-map/app/people/service/internal/data/ent"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/genglsh/go-trainning-map/app/people/service/internal/biz"
)

type PeopleRepo struct {
	data *Data
	log  *log.Helper
}

func NewLabelRepo(data *Data, logger log.Logger) biz.PeopleRepo {
	return &PeopleRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/server-service")),
	}
}

func (p *PeopleRepo) CreateCollegeInfo(ctx context.Context, peoples []*biz.People) error {

	bulk := make([]*ent.PeopleCreate, 0, len(peoples))
	for _, ele := range peoples {
		tmp := &ent.PeopleCreate{}
		bulk = append(bulk, tmp.SetIdnum(ele.ID).SetName(ele.Name).SetTelephone(ele.TelePhone))
	}
	err := p.data.db.People.CreateBulk(bulk...).Exec(ctx)
	return errors.Wrap(err, "failed to insert data ")
}

func cvtent2biz(people2 *ent.People) *biz.People {
	if people2 == nil {
		return nil
	}
	return &biz.People{
		Name:      people2.Name,
		ID:        people2.Idnum,
		TelePhone: people2.Telephone,
	}
}

func (p *PeopleRepo) GetCollegeInfo(ctx context.Context, ID string) (*biz.People, error) {
	res, err := p.data.db.People.Query().Where(people.Idnum(ID)).First(ctx)
	if ent.IsNotFound(err) {
		p.log.Infof("failed to search %s people", ID)
	} else if err != nil {
		return nil, errors.Wrap(err, "mysql error")
	}
	return cvtent2biz(res), nil

}
