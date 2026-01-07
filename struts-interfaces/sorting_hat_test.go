package strutsinterfaces

import "testing"

func TestSortStudent(t *testing.T) {

	t.Run("sort the student to a house", func(t *testing.T) {
		got := SortStudent(Student{"santhosh", "wit"})
		want := "Revanclaw"

		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	})

	t.Run("sort as muggle if no character", func(t *testing.T) {
		got := SortStudent(Student{"legolas", ""})
		want := "Muggle"

		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	})
}

func TestIntroduce(t *testing.T) {

	// decoupled from concrete types student or teacher
	checkIntroduction := func(t testing.TB, person Person, want string) {
		t.Helper()
		got := person.Introduce()
		if got != want {
			t.Errorf("got %s, but want %s", got, want)
		}
	}

	t.Run("student introduction", func(t *testing.T) {
		student := Student{"santhosh", "wit"}
		want := "I am santhosh, I am wit"

		checkIntroduction(t, student, want)

	})

	t.Run("teacher introduction", func(t *testing.T) {
		teacher := Teacher{name: "snape", skill: "potions"}
		want := "I am snape, I teach potions"

		checkIntroduction(t, teacher, want)
	})

}
