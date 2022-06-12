package db

import (
	"fmt"
	"time"
)

/*
 id |          time_stamp           |                           task                            | done | present |         last_modified
----+-------------------------------+-----------------------------------------------------------+------+---------+-------------------------------
  1 | 2022-06-05 17:35:19.945537+03 | Create a script for DB initialization.                    | f    | t       | 2022-06-06 11:44:35.373912+03
  2 | 2022-06-05 17:37:25.017739+03 | Create appropriate user permissions.                      | f    | t       | 2022-06-06 11:44:35.373912+03
  4 | 2022-06-06 01:26:42.308235+03 | NEW TASK                                                  | t    | t       | 2022-06-06 11:44:35.373912+03
  3 | 2022-06-05 17:38:31.12035+03  | Decide if ORM or Raw SQL.                                 | f    | t       | 2022-06-06 11:44:35.373912+03
  5 | 2022-06-06 02:39:42.879486+03 | handle messages of those who are already finished         | f    | t       | 2022-06-06 11:44:35.373912+03
  6 | 2022-06-06 11:45:56.270987+03 | Handle updating the not present and already done elements | f    | t       | 2022-06-06 11:45:56.270987+03
*/

type Task struct {
	ID            int
	Time_Stamp    time.Time
	Task_Text     string
	Is_Done       bool
	Present       bool
	Last_Modified time.Time
}

var cross_mark = '\u274C'
var tick = '\u2714'

// var skull = '\u1F480'
// var baby = '\u1F476'

var done_set map[bool]rune = map[bool]rune{true: tick, false: cross_mark}
var present_set map[bool]rune = map[bool]rune{true: '+', false: '-'}

func (t *Task) String() string {
	return fmt.Sprintf("(%d) %s \n", t.ID, t.Task_Text)
}

// Verbose string
func (t *Task) V_String() string {
	return fmt.Sprintf("[%v](%d)[%c][%c] %s \n", t.Last_Modified, t.ID, done_set[t.Is_Done], present_set[t.Present], t.Task_Text)
}

// Completed string
func (t *Task) C_String() string {
	return fmt.Sprintf(" %c (%d) %s \n", done_set[t.Is_Done], t.ID, t.Task_Text)
}
