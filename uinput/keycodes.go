package uinput

/*
  #include <linux/uinput.h>
*/
import "C"

type KeyCode C.__u16

const (
	/* letters */
	KEY_A KeyCode = C.KEY_A
	KEY_B         = C.KEY_B
	KEY_C         = C.KEY_C
	KEY_D         = C.KEY_D
	KEY_E         = C.KEY_E
	KEY_F         = C.KEY_F
	KEY_G         = C.KEY_G
	KEY_H         = C.KEY_H
	KEY_I         = C.KEY_I
	KEY_J         = C.KEY_J
	KEY_K         = C.KEY_K
	KEY_L         = C.KEY_L
	KEY_M         = C.KEY_M
	KEY_N         = C.KEY_N
	KEY_O         = C.KEY_O
	KEY_P         = C.KEY_P
	KEY_Q         = C.KEY_Q
	KEY_R         = C.KEY_R
	KEY_S         = C.KEY_S
	KEY_T         = C.KEY_T
	KEY_U         = C.KEY_U
	KEY_V         = C.KEY_V
	KEY_W         = C.KEY_W
	KEY_X         = C.KEY_X
	KEY_Y         = C.KEY_Y
	KEY_Z         = C.KEY_Z

	/* numbers */
	KEY_0         = C.KEY_0
	KEY_1         = C.KEY_1
	KEY_2         = C.KEY_2
	KEY_3         = C.KEY_3
	KEY_4         = C.KEY_4
	KEY_5         = C.KEY_5
	KEY_6         = C.KEY_6
	KEY_7         = C.KEY_7
	KEY_8         = C.KEY_8
	KEY_9         = C.KEY_9

	/* special */
	KEY_ENTER    = C.KEY_ENTER
	KEY_ESC      = C.KEY_ESC
	KEY_TAB      = C.KEY_TAB
	KEY_BACK     = C.KEY_BACK
	KEY_CTRL_L   = C.KEY_LEFTCTRL
	KEY_CTRL_R   = C.KEY_RIGHTCTRL
	KEY_SHIFT_L  = C.KEY_LEFTSHIFT
	KEY_SHIFT_R  = C.KEY_RIGHTSHIFT
)
