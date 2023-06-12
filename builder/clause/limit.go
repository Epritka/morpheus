package clause

import "fmt"

type limit int

func Limit(number int) limit {
	return limit(number)
}

func (l limit) String() string {
	return fmt.Sprintf("LIMIT %d", l)
}
