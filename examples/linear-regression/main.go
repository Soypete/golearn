package main

import (
	"fmt"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	regression "github.com/sjwhitworth/golearn/linear_models"
)

const filepath = `src/github.com/Soypete/golearn/examples/linear-regression`

func main() {

	//TODO: add graph export
	// // Get data from csv file that contains advertising statics.
	// advertFile, err := os.Open(filepath + "/data/Advertising.csv")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// advertDF := dataframe.ReadCSV(advertFile)
	// advertSummary := advertDF.Describe()
	// fmt.Println(advertSummary)
	// advertFile.Close()
	// // visualize data information
	// for _, colName := range advertDF.Names() {
	// 	plotVals := make(plotter.Values, advertDF.Nrow())
	// 	for i, floatVal := range advertDF.Col(colName).Float() {
	// 		plotVals[i] = floatVal
	// 	}
	// 	p, err := plot.New()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	p.Title.Text = fmt.Sprintf("Histogram of a %s", colName)
	// 	h, err := plotter.NewHist(plotVals, 16)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	h.Normalize(1)
	// 	p.Add(h)
	// 	if err := p.Save(4*vg.Inch, 4*vg.Inch, filepath+"graphs/"+colName+"_hist.png"); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	// // we are comparing all of the other columns to the Sales column making Sales the dependent Variable
	// ySales := advertDF.Col("Sales").Float()

	// for _, colName := range advertDF.Names() {
	// 	pts := make(plotter.XYs, advertDF.Nrow())
	// 	for i, floatVal := range advertDF.Col(colName).Float() {
	// 		pts[i].X = floatVal
	// 		pts[i].Y = ySales[i]
	// 	}
	// 	p, err := plot.New()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	p.X.Label.Text = colName
	// 	p.Y.Label.Text = "y"
	// 	p.Add(plotter.NewGrid())
	// 	s, err := plotter.NewScatter(pts)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	s.GlyphStyle.Radius = vg.Points(3)
	// 	p.Add(s)
	// 	if err := p.Save(4*vg.Inch, 4*vg.Inch, filepath+"graphs/"+colName+"_scatter.png"); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	data, err := base.ParseCSVToInstances(filepath+"/data/Advertising.csv", true)
	if err != nil {
		panic(err)
	}
	trainData, testData := base.InstancesTrainTestSplit(data, 0.70)
	r := regression.NewLinearRegression()
	err = r.Fit(trainData)
	if err != nil {
		panic(err)
	}
	predictions, err := r.Predict(testData)
	if err != nil {
		panic(err)
	}
	// Evaluate
	fmt.Println("Linear Regression (information gain)")
	cf, err := evaluation.GetConfusionMatrix(testData, predictions)
	if err != nil {
		panic(fmt.Errorf("Unable to get confusion matrix: %s", err.Error()))
	}
	fmt.Println(evaluation.GetSummary(cf))

}
