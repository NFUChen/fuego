package newEntity

type NewEntity struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type NewEntityService interface {
	GetNewEntity(id string) (*NewEntity, error)
	CreateNewEntity(*NewEntity) (*NewEntity, error)
	GetAllNewEntity() ([]*NewEntity, error)
	UpdateNewEntity(id string, entity *NewEntity) (*NewEntity, error)
	DeleteNewEntity(id string) (any, error)
}
