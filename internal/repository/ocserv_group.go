package repository

import (
	"context"
	"fmt"
	"ocserv/internal/repository/ocserv"
	"os"
	"reflect"
)

var ocservGroupDir = "/etc/ocserv/groups"

type OcsGroupRepository struct {
	cmd ocserv.ServiceOcservRepositoryInterface
}

type OcservGroupRepositoryInterface interface {
	GroupList() []string
	GroupCreateOrUpdate(context.Context) error
	GroupDelete(string) error
	DefaultGroupUpdate(ctx context.Context) error
}

func NewOcservGroupRepository() *OcsGroupRepository {
	return &OcsGroupRepository{
		cmd: ocserv.NewOcservRepository(),
	}
}

func (g *OcsGroupRepository) GroupList() []string {
	var groups []string
	groupsFiles, err := os.ReadDir(ocservGroupDir)
	if err != nil {
		return groups
	}
	for _, file := range groupsFiles {
		groups = append(groups, file.Name())
	}
	return groups
}

func (g *OcsGroupRepository) GroupCreateOrUpdate(ctx context.Context) error {
	name := ctx.Value("name").(string)
	value := ctx.Value("config").(reflect.Value)
	filePath := fmt.Sprintf("%s/%s", ocservGroupDir, name)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	ch := updateConfigFile(filePath, value)
	select {
	case err := <-ch:
		if err != nil {
			cancel()
			_ = os.Remove(filePath)
			return err
		}
		_ = g.cmd.ReloadService()
		return nil
	case <-ctx.Done():
		_ = os.Remove(filePath)
		return ctx.Err()
	}

}

func (g *OcsGroupRepository) GroupDelete(groupName string) error {
	err := os.Remove(fmt.Sprintf("%s/%s", ocservGroupDir, groupName))
	if err != nil {
		return err
	}
	_ = g.cmd.ReloadService()
	return nil
}

func (g *OcsGroupRepository) DefaultGroupUpdate(ctx context.Context) error {
	filePath := fmt.Sprintf("%s/defaults", ocservGroupDir)
	value := ctx.Value("config").(reflect.Value)
	ch := updateConfigFile(filePath, value)
	_ = g.cmd.ReloadService()
	return <-ch
}
