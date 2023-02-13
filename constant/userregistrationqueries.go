package constant

const (
	PutUpdateQuery    = "UPDATE users SET id=?, username=?, email=?, password=? WHERE id=?"
	PostDetailsQuery  = "Insert into users (id, username, email, password) values (?,?,?,?)"
	GetRecordQuery    = "SELECT id, username, email, password FROM users WHERE id = ?"
	GetAllRecordQuery = "SELECT id,username,email,password from users"
	DeleteRecord      = "DELETE from users WHERE id = ?"
)
