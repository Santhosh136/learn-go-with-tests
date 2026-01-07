package strutsinterfaces

const (
	gryffindor = "Gryffindor"
	slytherin  = "Slytherin"
	hufflepuff = "Hufflepuff"
	revanclaw  = "Revanclaw"
)

type Student struct {
	name      string
	character string
}

func (student Student) Introduce() string {
	return "I am " + student.name + " and I am " + student.character
}

func SortStudent(student Student) string {
	return sort(student)
}

func sort(student Student) string {
	var house string
	switch student.character {
	case "brave":
		house = gryffindor
	case "cunning":
		house = slytherin
	case "justice":
		house = hufflepuff
	case "wit":
		house = revanclaw
	default:
		house = "Muggle"
	}
	return house
}
