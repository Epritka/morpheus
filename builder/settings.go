package builder

type (
	registeredWordType string
	operatorType       string
	directedType       string
	funcNameType       string
)

const (
	_OPTIONAL registeredWordType = "OPTIONAL"
	_MATCH    registeredWordType = "MATCH"

	_MERGE     registeredWordType = "MERGE"
	_ON_MATCH  registeredWordType = "ON MATCH"
	_ON_CREATE registeredWordType = "ON CREATE"

	_CREATE registeredWordType = "CREATE"

	_WHERE registeredWordType = "WHERE"

	_SET registeredWordType = "SET"

	_WITH registeredWordType = "WITH"
	_DESC registeredWordType = "DESC"
	_AS   registeredWordType = "AS"

	_CALL  registeredWordType = "CALL"
	_YIELD registeredWordType = "YIELD"

	_DETACH registeredWordType = "DETACH"
	_DELETE registeredWordType = "DELETE"

	_RETURN   registeredWordType = "RETURN"
	_ORDER_BY registeredWordType = "ORDER BY"
	_LIMIT    registeredWordType = "LIMIT"

	_NOT operatorType = "NOT "
	_AND operatorType = " AND "
	_XOR operatorType = " XOR "
	_OR  operatorType = " OR "

	_EQUAL     operatorType = "="
	_NOT_EQUAL operatorType = "<>"
	_MORE      operatorType = ">"
	_LESS      operatorType = "<"

	_DIRECTED_OMITTED directedType = "-"
	_DIRECTED_IN      directedType = "->"
	_DIRECTED_OUT     directedType = "<-"
)
