package newEntity

import (
	"github.com/go-fuego/fuego"
)

type NewEntityResources struct {
	// TODO add resources
	NewEntityService NewEntityService
}

func (resource NewEntityResources) Routes(s *fuego.Server) {
	newEntityGroup := fuego.Group(s, "/newEntity")

	fuego.Get(newEntityGroup, "/", resource.getAllNewEntity)
	fuego.Post(newEntityGroup, "/", resource.postNewEntity)

	fuego.Get(newEntityGroup, "/{id}", resource.getNewEntity)
	fuego.Put(newEntityGroup, "/{id}", resource.putNewEntity)
	fuego.Delete(newEntityGroup, "/{id}", resource.deleteNewEntity)
}

func (resource NewEntityResources) getAllNewEntity(c fuego.ContextNoBody) ([]*NewEntity, error) {
	return resource.NewEntityService.GetAllNewEntity()
}

func (resource NewEntityResources) postNewEntity(c fuego.ContextWithBody[*NewEntity]) (*NewEntity, error) {
	body, err := c.Body()
	if err != nil {
		return nil, err
	}

	return resource.NewEntityService.CreateNewEntity(body)
}

func (resource NewEntityResources) getNewEntity(c fuego.ContextNoBody) (*NewEntity, error) {
	id := c.PathParam("id")

	return resource.NewEntityService.GetNewEntity(id)
}

func (resource NewEntityResources) putNewEntity(c fuego.ContextWithBody[*NewEntity]) (*NewEntity, error) {
	id := c.PathParam("id")

	body, err := c.Body()
	if err != nil {
		return nil, err
	}

	return resource.NewEntityService.UpdateNewEntity(id, body)
}

func (resource NewEntityResources) deleteNewEntity(c fuego.ContextNoBody) (any, error) {
	return resource.NewEntityService.DeleteNewEntity(c.PathParam("id"))
}
