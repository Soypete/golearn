package fifa

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/trees"
)

// DecisionTree Write and train DecisionTree model.
func DecisionTree(file string) error {
	var wg sync.WaitGroup
	rand.Seed(44111342)

	// Load in the words dataset
	words, err := base.ParseCSVToInstances(file, true)
	if err != nil {
		return err
	}
	wg.Add(4)
	// Create a 70-30 training-test split
	trainData, testData := base.InstancesTrainTestSplit(words, 0.70)
	go func() error {
		defer wg.Done()
		tree := trees.NewID3DecisionTree(0.6)
		// Train the ID3 tree
		err = tree.Fit(trainData)
		if err != nil {
			return err
		}
		// Generate predictions
		predictions, err := tree.Predict(testData)
		if err != nil {
			return err
		}
		// Evaluate
		fmt.Println("ID3 Performance (information gain)")
		cf, err := evaluation.GetConfusionMatrix(testData, predictions)
		if err != nil {
			return fmt.Errorf("Unable to get confusion matrix: %s", err.Error())
		}
		fmt.Println(evaluation.GetSummary(cf))
		tree.Save("models/DecisonTree.h")
		return nil
	}()
	go func() error {
		defer wg.Done()
		tree := trees.NewID3DecisionTreeFromRule(0.6, new(trees.InformationGainRatioRuleGenerator))
		// (Parameter controls train-prune split.)

		// Train the ID3 tree
		err := tree.Fit(trainData)
		if err != nil {
			return err
		}

		// Generate predictions
		predictions, err := tree.Predict(testData)
		if err != nil {
			return err
		}

		// Evaluate
		fmt.Println("ID3 Performance (information gain ratio)")
		cf, err := evaluation.GetConfusionMatrix(testData, predictions)
		if err != nil {
			panic(fmt.Sprintf("Unable to get confusion matrix: %s", err.Error()))
		}
		fmt.Println(evaluation.GetSummary(cf))
		tree.Save("models/DecisonTreeInformationGain.h")
		return nil
	}()

	go func() error {
		defer wg.Done()
		tree := trees.NewID3DecisionTreeFromRule(0.6, new(trees.GiniCoefficientRuleGenerator))
		// (Parameter controls train-prune split.)

		// Train the ID3 tree
		err := tree.Fit(trainData)
		if err != nil {
			return err
		}

		// Generate predictions
		predictions, err := tree.Predict(testData)
		if err != nil {
			return err
		}

		// Evaluate
		fmt.Println("ID3 Performance (gini index generator)")
		cf, err := evaluation.GetConfusionMatrix(testData, predictions)
		if err != nil {
			panic(fmt.Sprintf("Unable to get confusion matrix: %s", err.Error()))
		}
		fmt.Println(evaluation.GetSummary(cf))
		tree.Save("models/DecisonTreeGiniCoefficient.h")
		return nil
	}()

	go func() error {
		defer wg.Done()
		// Consider two randomly-chosen attributes
		tree := trees.NewRandomTree(4)
		err = tree.Fit(trainData)
		if err != nil {
			return err
		}
		predictions, err := tree.Predict(testData)
		if err != nil {
			return err
		}
		fmt.Println("RandomTree Performance")
		cf, err := evaluation.GetConfusionMatrix(testData, predictions)
		if err != nil {
			panic(fmt.Sprintf("Unable to get confusion matrix: %s", err.Error()))
		}
		fmt.Println(evaluation.GetSummary(cf))
		tree.Save("models/RandomTree.h")
		return nil
	}()
	return err
}
