package plutogo

// PageMargins represents page margin configurations
type PageMargins struct {
	Top    float64
	Right  float64
	Bottom float64
	Left   float64
}

// Pre-defined page margins
var (
	PageMarginsNone     = PageMargins{0, 0, 0, 0}
	PageMarginsNormal   = PageMargins{72, 72, 72, 72} // 1 inch
	PageMarginsNarrow   = PageMargins{36, 36, 36, 36} // 0.5 inches
	PageMarginsModerate = PageMargins{72, 54, 72, 54}
	PageMarginsWide     = PageMargins{72, 144, 72, 144}
)
