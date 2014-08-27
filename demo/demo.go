package main

import (
	"log"
	"time"

	"github.com/galaktor/gostwriter"
	"github.com/galaktor/gostwriter/key"
)

func main() {
	kb, err := gostwriter.New("foo")

	guard(err)

	t, err    := kb.Get(key.CODE_T);         guard(err);
	e, err    := kb.Get(key.CODE_E);         guard(err);
	s, err    := kb.Get(key.CODE_S);         guard(err);
	ret, err  := kb.Get(key.CODE_ENTER);     guard(err);

	ctrl, err := kb.Get(key.CODE_LEFTCTRL);  guard(err);
	c, err    := kb.Get(key.CODE_C);         guard(err);

	log.Println("this demo will type the word 'test' and a newline 10 times")
	log.Println("then it will terminate itself by pressing CTRL + C")

	cnt := 0
	for {
		<-time.After(time.Millisecond*100)
		push(t)
		<-time.After(time.Millisecond*100)
		push(e)
		<-time.After(time.Millisecond*100)
		push(s)
		<-time.After(time.Millisecond*100)
		push(t)
		<-time.After(time.Millisecond*500)
		push(ret)
		
		if cnt = cnt + 1; cnt == 10 {
			press(ctrl)
			press(c)
		}

	}

	kb.Destroy()
}

/* presses and subsequently releases a key */
func push(k *gostwriter.K) {
	err := k.Push(); guard(err);
}

/* presses a key, if not already pressed. does not release! */
func press(k *gostwriter.K) {
	err := k.Press(); guard(err);
}

/* releases a key, if not aready released. */
func release(k *gostwriter.K) {
	err := k.Release(); guard(err);
}

func guard(e error) {
	if e != nil {
		log.Fatal(e)
	}
}











