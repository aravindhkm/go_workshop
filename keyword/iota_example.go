package keyword

import "fmt"

type Day int

const (
    Sunday Day = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)

func (d Day) String() string {
    return [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}[d]
}

func IotaExampleOne() {
    var today Day = Wednesday
    fmt.Println(today.String()) // Output: Tuesday
}
