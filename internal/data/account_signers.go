package data

type AccountSigner struct {
	SignerID string `db:"signer_id"`
	RoleID   uint64 `db:"role_id"`
	Weight   uint32 `db:"weight"`
	Identity uint32 `db:"identity"`
	//TODO Details?
}
