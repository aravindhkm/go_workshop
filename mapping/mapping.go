package mapping

import "fmt"

type Store string

func (d *Store) PrintStr(a string) string {
	return fmt.Sprintf("%s : %s", *d,a)
}

func MappingExecute() {

}