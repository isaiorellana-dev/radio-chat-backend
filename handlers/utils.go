package handlers

import m "github.com/isaiorellana-dev/livechat-backend/models"

type objectStr map[string]string
type sliceStr []string

const POST_MESSAGES = "post_messages"
const DELETE_MESSAGES = "delete_messages"
const DELETE_USERS = "delete_users"
const VIEW_LIST_OF_USERS = "view_list_of_users"
const VIEW_LIST_OF_ROLES = "view_list_of_roles"
const VIEW_LIST_OF_PERMISSIONS = "view_list_of_permissions"
const ASSIGN_ROLES = "assign_roles"
const CREATE_ROLES = "create_roles"
const CREATE_PERMISSIONS = "create_permissions"

const ADMIN = "admin"
const GUEST = "guest"
const DEV = "dev"

var admin = m.Role{Name: ADMIN}
var dev = m.Role{Name: DEV}
var guest = m.Role{Name: GUEST}

var postMessages = m.Permission{Name: POST_MESSAGES}
var deleteMessages = m.Permission{Name: DELETE_MESSAGES}
var deleteUsers = m.Permission{Name: DELETE_USERS}
var viewListOfUsers = m.Permission{Name: VIEW_LIST_OF_USERS}
var viewListOfRoles = m.Permission{Name: VIEW_LIST_OF_ROLES}
var viewListOfPermissions = m.Permission{Name: VIEW_LIST_OF_PERMISSIONS}
var assignRoles = m.Permission{Name: ASSIGN_ROLES}
var createRoles = m.Permission{Name: CREATE_ROLES}
var createPermissions = m.Permission{Name: CREATE_PERMISSIONS}

var GetUsersPerms = sliceStr{
	VIEW_LIST_OF_USERS,
}

var CreateMessagePerms = sliceStr{
	POST_MESSAGES,
}

var GetRolesPerms = sliceStr{
	VIEW_LIST_OF_ROLES,
}

var CreateRolePerms = sliceStr{
	CREATE_ROLES,
	VIEW_LIST_OF_ROLES,
	ASSIGN_ROLES,
}
var GetPermissionsPerms = sliceStr{
	VIEW_LIST_OF_PERMISSIONS,
}

var CreatePermissionsPerms = sliceStr{
	CREATE_PERMISSIONS,
}

var DeleteUserPerms = sliceStr{
	VIEW_LIST_OF_USERS,
	DELETE_USERS,
}

var DeleteMessagePerms = sliceStr{
	DELETE_MESSAGES,
}
