package handlers

type objectStr map[string]string
type sliceStr []string

var GetUsersPerms = sliceStr{
	"view_list_of_users",
}

var CreateMessagePerms = sliceStr{
	"post_messages",
}

var GetRolesPerms = sliceStr{
	"view_list_of_roles",
}

var CreateRolePerms = sliceStr{
	"create_roles",
	"view_list_of_roles",
	"assign_roles",
}
var GetPermissionsPerms = sliceStr{
	"view_list_of_permissions",
}

var CreatePermissionsPerms = sliceStr{
	"create_permissions",
}

var DeleteUserPerms = sliceStr{
	"view_list_of_users",
	"delete_users",
}

var DeleteMessagePerms = sliceStr{
	"delete_messages",
}
