package dummy

import (
	"io"
	"strconv"

	"github.com/golang/protobuf/proto"

	"github.com/ketchuphq/ketchup/db"
	"github.com/ketchuphq/ketchup/proto/ketchup/api"
	"github.com/ketchuphq/ketchup/proto/ketchup/models"
)

type DummyDB struct {
	Users  map[string]*models.User
	Pages  map[string]*models.Page
	Routes map[string]*models.Route
	Data   map[string]*models.Data

	counter int
}

var _ db.Backend = &DummyDB{}

func New() *DummyDB {
	return &DummyDB{
		Users:  map[string]*models.User{},
		Pages:  map[string]*models.Page{},
		Routes: map[string]*models.Route{},
		Data:   map[string]*models.Data{},
	}
}

func (d *DummyDB) GetUser(uuid string) (*models.User, error) {
	return d.Users[uuid], nil
}

func (d *DummyDB) GetUserByEmail(email string) (*models.User, error) {
	for _, u := range d.Users {
		if u.GetEmail() == email {
			return u, nil
		}
	}
	return nil, nil
}

func (d *DummyDB) GetUserByToken(token string) (*models.User, error) {
	for _, u := range d.Users {
		if u.GetToken() == token {
			return u, nil
		}
	}
	return nil, nil
}

func (d *DummyDB) UpdateUser(u *models.User) error {
	// todo: set uuid, timestamp?
	d.Users[u.GetEmail()] = u
	return nil
}

func (d *DummyDB) GetPage(uuid string) (*models.Page, error) {
	return d.Pages[uuid], nil
}

func (d *DummyDB) UpdatePage(p *models.Page) error {
	if p.GetUuid() == "" {
		p.Uuid = proto.String(strconv.Itoa(d.counter))
		d.counter++
	}
	d.Pages[p.GetUuid()] = p
	return nil
}

func (d *DummyDB) DeletePage(p *models.Page) error {
	delete(d.Pages, p.GetUuid())
	return nil
}

func (d *DummyDB) ListPages(_ *api.ListPageRequest) ([]*models.Page, error) {
	pages := []*models.Page{}
	for _, p := range d.Pages {
		pages = append(pages, p)
	}
	return pages, nil
}

func (d *DummyDB) GetRoute(uuid string) (*models.Route, error) {
	return d.Routes[uuid], nil
}

func (d *DummyDB) UpdateRoute(r *models.Route) error {
	if r.GetUuid() == "" {
		r.Uuid = proto.String(strconv.Itoa(d.counter))
		d.counter++
	}
	d.Routes[r.GetUuid()] = proto.Clone(r).(*models.Route)
	return nil
}

func (d *DummyDB) DeleteRoute(r *models.Route) error {
	delete(d.Routes, r.GetUuid())
	return nil
}

func (d *DummyDB) ListRoutes(req *api.ListRouteRequest) ([]*models.Route, error) {
	routes := []*models.Route{}
	for _, r := range d.Routes {
		uuid := req.GetOptions().GetPageUuid()
		if uuid != "" && uuid != r.GetPageUuid() {
			continue
		}
		routes = append(routes, r)
	}
	return routes, nil
}

func (d *DummyDB) Debug(w io.Writer) error {
	return nil
}

func (d *DummyDB) GetData(key string) (*models.Data, error) {
	return d.Data[key], nil
}

func (d *DummyDB) UpdateData(data *models.Data) error {
	d.Data[data.GetKey()] = data
	return nil
}

func (d *DummyDB) UpdateDataBatch(data []*models.Data) error {
	for _, datum := range data {
		d.Data[datum.GetKey()] = datum
	}
	return nil
}

func (d *DummyDB) DeleteData(data *models.Data) error {
	delete(d.Data, data.GetKey())
	return nil
}

func (d *DummyDB) ListData() ([]*models.Data, error) {
	data := []*models.Data{}
	for _, d := range d.Data {
		data = append(data, d)
	}
	return data, nil
}
