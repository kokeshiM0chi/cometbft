package progressbar

import "fmt"

// the progressbar indicates the current status and progress would be desired.
// ref: https://www.pixelstech.net/article/1596946473-A-simple-example-on-implementing-progress-bar-in-GoLang

type Bar struct {
	rate    string
	graph   string
	percent int64
	cur     int64
	start   int64
	total   int64
}

func (bar *Bar) NewOption(start, total int64) {
	bar.cur = start
	bar.start = start
	bar.total = total
	bar.graph = "█"
	bar.percent = bar.getPercent()
}

func (bar *Bar) getPercent() int64 {
	return int64(float32(bar.cur-bar.start) / float32(bar.total-bar.start) * 100)
}

func (bar *Bar) Play(cur int64) {
	bar.cur = cur
	last := bar.percent
	bar.percent = bar.getPercent()
	if bar.percent != last && bar.percent%2 == 0 {
		bar.rate += bar.graph
	}
	fmt.Printf("\r[%-50s]%3d%% %8d/%d", bar.rate, bar.percent, bar.cur, bar.total)
}

func (bar *Bar) Finish() {
	fmt.Println()
}
