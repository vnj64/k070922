package permissions

import (
	"fmt"
	"project/pkg/enum"
)

type Permission string // CreateUser

type PermissionMeta struct {
	Description string // "create_user"
}

var (
	permissions     = enum.New[string, Permission]()
	permissionsData = make(map[Permission]PermissionMeta, 100)
)

func register(enumCase Permission, description ...string) Permission {
	if _, exists := permissions.From(enumCase.Value()); exists {
		panic(fmt.Sprintf("permission '%s' is already registered", enumCase))
	}

	if len(description) > 0 {
		permissionsData[enumCase] = PermissionMeta{Description: description[0]}
	} else {
		permissionsData[enumCase] = PermissionMeta{Description: ""}
	}

	return permissions.Register(enumCase)
}

var (
	CreateUser = register("create_user", "only super admin creates user")
)

type RolePermissions map[Role][]Permission

var rolePermissions = RolePermissions{
	SuperAdmin: {
		CreateUser,
	},
	Default: {},
}

func PermissionsForRole(role Role) []Permission {
	return rolePermissions[role]
}

func (s Permission) Value() string {
	return string(s)
}
