package utils

import (
	"crspy2/licenses/database"
)

var nameToPermission = map[string]database.Permission{
	"Approve Staff":         database.ApproveStaffPermission,
	"Reset HWID":            database.HWIDResetPermission,
	"Reset Password":        database.PassResetPermission,
	"View License Keys":     database.ViewKeysPermission,
	"Generate Keys":         database.KeyGenPermission,
	"Compensate Keys":       database.CompensatePermission,
	"Change Product Status": database.StatusChangePermission,
	"Manage Products":       database.ProductsPermission,
	"Manage User Actions":   database.UserActionsPermission,
	"Set/Edit Offsets":      database.OffsetsPermission,
}

func PermissionFromName(name string) (database.Permission, bool) {
	perm, exists := nameToPermission[name]
	return perm, exists
}

func PermissionFromInt(value int64) database.Permission {
	return database.Permission(value)
}
