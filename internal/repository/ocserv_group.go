package repository

type OcsGroupRepository struct{}

type OcservGroupRepositoryInterface interface{}

var ocservPath = "/etc/ocserv"

func NewOcservGroupRepository() *OcsGroupRepository {
	return &OcsGroupRepository{}
}

func (g *OcsGroupRepository) GroupList() []string {
	var groups []string

	return groups
}
