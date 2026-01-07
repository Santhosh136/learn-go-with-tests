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

	student := Student{"santhosh", "wit"}
	got := student.Introduce()
	want := "I am santhosh and I am wit"

	if got != want {
		t.Errorf("got %s, but want %s", got, want)
	}
}
