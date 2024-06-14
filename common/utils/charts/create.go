package charts

import (
	"os"

	"github.com/vicanso/go-charts/v2"
)

func CreateCharts() {
	xValue := []string{"A1", "A2"}
	yValue := [][]float64{
		{1, 2},
		{100000, 200000},
	}

	f := false
	p, err := charts.LineRender(
		yValue,
		charts.FontFamilyOptionFunc("noto"),
		charts.TitleTextOptionFunc("1"),
		charts.XAxisDataOptionFunc(xValue),
		func(opt *charts.ChartOption) {
			opt.XAxis.BoundaryGap = &f
			opt.Padding = charts.Box{Left: 20, Right: 50, Top: 20, Bottom: 20}
		},
		charts.ThemeOptionFunc("grafana"),
		charts.WidthOptionFunc(1000),
	)
	if err != nil {
		return
	}
	buf, _ := p.Bytes()

	file, _ := os.Create("test.png")
	_, _ = file.Write(buf)
}
