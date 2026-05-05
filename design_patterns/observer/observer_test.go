package observer

import "testing"

func TestObserver(t *testing.T) {
	subject := NewSubject()

	r1 := NewReader("reader1")
	r2 := NewReader("reader2")
	r3 := NewReader("reader3")

	subject.Attach(r1)
	subject.Attach(r2)
	subject.Attach(r3)

	subject.UpdateContext("observer mode")

	readers := []*Reader{r1, r2, r3}
	for _, r := range readers {
		received := r.Received()
		if len(received) != 1 {
			t.Fatalf("%s expect 1 update got %d", r.Name(), len(received))
		}
		if received[0] != "observer mode" {
			t.Fatalf("%s expect observer mode got %s", r.Name(), received[0])
		}
	}
}
