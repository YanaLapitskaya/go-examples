package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type Person struct {
	firstName string
	lastName  string
	birthday  time.Time
}

type People []Person

var dateLayout string = "2006-01-02"

func (p People) Len() int {
	return len(p)
}

func (p Person) String() string {
	return fmt.Sprintf("%v, %v, %v \n", p.firstName, p.lastName, p.birthday.Format(dateLayout))
}

func (p People) Less(i, j int) bool {
	if p[i].birthday.Equal(p[j].birthday) {
		return strings.Compare(
			p[i].firstName+p[i].lastName,
			p[j].firstName+p[j].lastName,
		) < 0
	}
	return p[i].birthday.After(p[j].birthday)
}

func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func parseDate(dateString string) time.Time {
	date, err := time.Parse(dateLayout, dateString)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return date
}

func main() {
	p := People{
		{"Ivan", "Ivanov", parseDate("2005-08-10")},
		{"Ivan", "Ivanov", parseDate("2003-08-05")},
		{"Artiom", "Ivanov", parseDate("2005-08-10")},
	}
	sort.Sort(p)
	fmt.Println(p)
}
