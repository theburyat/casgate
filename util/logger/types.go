package logger

type ObjectType string

const (
	ObjectTypeUser         ObjectType = "user"
	ObjectTypeRole         ObjectType = "role"
	ObjectTypeProvider     ObjectType = "provider"
	ObjectTypeApplication  ObjectType = "application"
	ObjectTypeUserApiToken ObjectType = "user-api-token"
	ObjectTypeLDAP         ObjectType = "ldap"
)

type OperationResult string

const (
	OperationResultSuccess OperationResult = "success"
	OperationResultFailure OperationResult = "failure"
)

type OperationName string

const (
	OperationNameAddOrganization    OperationName = "add-organization"
	OperationNameUpdateOrganization OperationName = "update-organization"

	OperationNameUserUpdate  OperationName = "user-update"
	OperationNameAddUser     OperationName = "add-user"
	OperationNameSetPassword OperationName = "set-password"
	OperationNameUserDelete  OperationName = "user-delete"

	OperationNameRoleUpdate     OperationName = "role-update"
	OperationNameSyncRoleToUser OperationName = "sync-role-to-user"

	OperationNameProviderUpdate OperationName = "provider-update"
	OperatoinNameAddProvider    OperationName = "add-provider"

	OperationNameApplicationUpdate OperationName = "application-update"

	OperationNameAddUserApiToken    OperationName = "add-user-api-token"
	OperationNameUserApiTokenDelete OperationName = "user-api-token-delete"

	OperationNameLdapSyncUsers OperationName = "ldap-sync-users"
)

type LogMsgDetailed map[string]interface{}
