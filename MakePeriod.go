package exsrapi

import (
	"time"
)

func MakePeriod(
	started_at int64,	//	開始日時（Unixtime）
	ended_at int64,	//	終了日時（Unixtime）
) (
	period string,	//	期間　Ex."Jul 31, 2023 6:00 PM - Aug 6, 2023 9:59 PM"
	err error,
) {
	st := time.Unix(started_at, 0)
	ed := time.Unix(ended_at, 0)
	period = st.Format("Jan 2, 2006 3:04 PM") + " - " + ed.Format("Jan 2, 2006 3:04 PM")
	return
}
