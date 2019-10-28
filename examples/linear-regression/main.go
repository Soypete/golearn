package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	regression "github.com/sjwhitworth/golearn/linear_models"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Get data from csv file that contains advertising statics.
	advertFile, err := os.Open("../datasets/Advertising.csv")
	if err != nil {
		log.Fatal(err)
	}

	advertDF := dataframe.ReadCSV(advertFile)
	advertSummary := advertDF.Describe()
	fmt.Println(advertSummary)
	advertFile.Close()
	// visualize data information
	for _, colName := range advertDF.Names() {
		plotVals := make(plotter.Values, advertDF.Nrow())
		for i, floatVal := range advertDF.Col(colName).Float() {
			plotVals[i] = floatVal
		}
		p, err := plot.New()
		if err != nil {
			log.Fatal(err)
		}
		p.Title.Text = fmt.Sprintf("Histogram of a %s", colName)
		h, err := plotter.NewHist(plotVals, 16)
		if err != nil {
			log.Fatal(err)
		}
		h.Normalize(1)
		p.Add(h)
		if err := p.Save(4*vg.Inch, 4*vg.Inch, "graphs/"+colName+"_hist.png"); err != nil {
			log.Fatal(err)
		}
	}
	// we are comparing all of the other columns to the Sales column making Sales the dependent Variable
	ySales := advertDF.Col("Sales").Float()

	for _, colName := range advertDF.Names() {
		pts := make(plotter.XYs, advertDF.Nrow())
		for i, floatVal := range advertDF.Col(colName).Float() {
			pts[i].X = floatVal
			pts[i].Y = ySales[i]
		}
		p, err := plot.New()
		if err != nil {
			log.Fatal(err)
		}
		p.X.Label.Text = colName
		p.Y.Label.Text = "y"
		p.Add(plotter.NewGrid())
		s, err := plotter.NewScatter(pts)
		if err != nil {
			log.Fatal(err)
		}
		s.GlyphStyle.Radius = vg.Points(3)
		p.Add(s)
		if err := p.Save(4*vg.Inch, 4*vg.Inch, "graphs/"+colName+"_scatter.png"); err != nil {
			log.Fatal(err)
		}
	}
	// // manual train/test split
	// trainingNum := (4 * advertDF.Nrow()) / 5
	// testNum := advertDF.Nrow() / 5
	// // adjust for odd amount
	// if trainingNum+testNum < advertDF.Nrow() {
	// 	trainingNum++
	// }
	// trainingIndex := make([]int, trainingNum)
	// testIndex := make([]int, testNum)
	// for i := 0; i < trainingNum; i++ {
	// 	trainingIndex[i] = i
	// }
	// for i := 0; i < testNum; i++ {
	// 	testIndex[i] = trainingNum + i
	// }
	// // randomize data
	// trainingDF := advertDF.Subset(trainingIndex)
	// testDF := advertDF.Subset(testIndex)
	// setMap := map[int]dataframe.DataFrame{
	// 	0: trainingDF,
	// 	1: testDF,
	// }
	// for i, setName := range []string{"data/training.csv", "data/test.csv"} {
	// 	f, err := os.Create(setName)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	w := bufio.NewWriter(f)
	// 	if err := setMap[i].WriteCSV(w); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	// trainCSV, err := os.Open("data/training.csv")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// readerTrain := csv.NewReader(trainCSV)
	// readerTrain.FieldsPerRecord = 4
	// trainingData, err := readerTrain.ReadAll()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	data, err := base.ParseCSVToInstances("../datasets/Advertising.csv", true)
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
