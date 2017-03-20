package drawimage

import "github.com/jung-kurt/gofpdf"

//Drawcircle is a demo for draw kinds of circles in pdf file
func Drawcircle(fileStr string) error {
	const (
		thin  = 0.2
		thick = 3.0
	)
	pdf := gofpdf.New("", "", "", "")
	pdf.SetFont("Helvetica", "", 12)
	pdf.SetFillColor(200, 200, 220)
	pdf.AddPage()

	y := 15.0
	pdf.Text(10, y, "Circles")
	pdf.SetFillColor(200, 200, 220)
	pdf.SetLineWidth(thin)
	pdf.Circle(20, y+15, 10, "D")
	pdf.Circle(45, y+15, 10, "F")
	pdf.Circle(70, y+15, 10, "FD")
	pdf.SetLineWidth(thick)
	pdf.Circle(95, y+15, 10, "FD")
	pdf.SetLineWidth(thin)

	y += 40.0
	pdf.Text(10, y, "Ellipses")
	pdf.SetFillColor(220, 200, 200)
	pdf.Ellipse(30, y+15, 20, 10, 0, "D")
	pdf.Ellipse(75, y+15, 20, 10, 0, "F")
	pdf.Ellipse(120, y+15, 20, 10, 0, "FD")
	pdf.SetLineWidth(thick)
	pdf.Ellipse(165, y+15, 20, 10, 0, "FD")
	pdf.SetLineWidth(thin)

	y += 40.0
	pdf.Text(10, y, "Curves (quadratic)")
	pdf.SetFillColor(220, 220, 200)
	pdf.Curve(10, y+30, 15, y-20, 40, y+30, "D")
	pdf.Curve(45, y+30, 50, y-20, 75, y+30, "F")
	pdf.Curve(80, y+30, 85, y-20, 110, y+30, "FD")
	pdf.SetLineWidth(thick)
	pdf.Curve(115, y+30, 120, y-20, 145, y+30, "FD")
	pdf.SetLineCapStyle("round")
	pdf.Curve(150, y+30, 155, y-20, 180, y+30, "FD")
	pdf.SetLineWidth(thin)
	pdf.SetLineCapStyle("butt")

	y += 40.0
	pdf.Text(10, y, "Curves (cubic)")
	pdf.SetFillColor(220, 200, 220)
	pdf.CurveBezierCubic(10, y+30, 15, y-20, 10, y+30, 40, y+30, "D")
	pdf.CurveBezierCubic(45, y+30, 50, y-20, 45, y+30, 75, y+30, "F")
	pdf.CurveBezierCubic(80, y+30, 85, y-20, 80, y+30, 110, y+30, "FD")
	pdf.SetLineWidth(thick)
	pdf.CurveBezierCubic(115, y+30, 120, y-20, 115, y+30, 145, y+30, "FD")
	pdf.SetLineCapStyle("round")
	pdf.CurveBezierCubic(150, y+30, 155, y-20, 150, y+30, 180, y+30, "FD")
	pdf.SetLineWidth(thin)
	pdf.SetLineCapStyle("butt")

	y += 40.0
	pdf.Text(10, y, "Arcs")
	pdf.SetFillColor(200, 220, 220)
	pdf.SetLineWidth(thick)
	pdf.Arc(45, y+35, 20, 10, 0, 0, 180, "FD")
	pdf.SetLineWidth(thin)
	pdf.Arc(45, y+35, 25, 15, 0, 90, 270, "D")
	pdf.SetLineWidth(thick)
	pdf.Arc(45, y+35, 30, 20, 0, 0, 360, "D")
	pdf.SetLineCapStyle("round")
	pdf.Arc(135, y+35, 20, 10, 135, 0, 180, "FD")
	pdf.SetLineWidth(thin)
	pdf.Arc(135, y+35, 25, 15, 135, 90, 270, "D")
	pdf.SetLineWidth(thick)
	pdf.Arc(135, y+35, 30, 20, 135, 0, 360, "D")
	pdf.SetLineWidth(thin)
	pdf.SetLineCapStyle("butt")

	err := pdf.OutputFileAndClose(fileStr)
	return err
}
