package main

type ScoreType int

const (
	Food ScoreType = iota
	Beverage
	Water
	Cheese
)

type NutritionalScore struct {
	Value     int
	Positive  int
	Negative  int
	ScoreType ScoreType
}

type EnergyKJ float64

type SugarGram float64

type SaturatedFattyaAcids float64

type SodiumMilligram float64

type FruitsPercent float64

type FibreGram float64

type ProteinGram float64

type NutricionalData struct {
	Energy               EnergyKJ
	Sugars               SugarGram
	SaturatedFattyaAcids SaturatedFattyaAcids
	Sodium               SodiumMilligram
	Fruits               FruitsPercent
	Fibre                FibreGram
	Protein              ProteinGram
	isWater              bool
}

var scoreToLetter = []string{"A", "B", "C", "D", "E"}
var energyLevels = []float64{3350, 3015, 2680, 2345, 2010, 1675, 1340, 1005, 670, 335}
var sugarLevels = []float64{45, 60, 36, 31, 27, 22.5, 18, 13.5, 9, 4.5}
var saturatedLevels = []float64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
var sodiumLevels = []float64{900, 810, 720, 630, 540, 450, 360, 270, 180, 90}
var fibreLevels = []float64{4.7, 3.7, 2.8, 1.9, 0.9}
var proteinLevels = []float64{8, 6.4, 4.8, 3.2, 1.6}

var energyLevelsBeverage = []float64{270, 240, 210, 180, 150, 120, 90, 60, 30, 0}
var sugarsLevelsBeverage = []float64{13.5, 12, 10.5, 9, 7.5, 6, 4.5, 3, 1.5, 0}

func (energy EnergyKJ) GetPoints(st ScoreType) int {
	if st == Beverage {
		return getPointsFromRange(float64(energy), energyLevelsBeverage)
	}
	return getPointsFromRange(float64(energy), energyLevels)
}

func (sugar SugarGram) GetPoints(st ScoreType) int {
	if st == Beverage {
		return getPointsFromRange(float64(sugar), sugarsLevelsBeverage)
	}
	return getPointsFromRange(float64(sugar), sugarLevels)
}

func (saturated SaturatedFattyaAcids) GetPoints(st ScoreType) int {
	return getPointsFromRange(float64(saturated), saturatedLevels)
}

func (sodium SodiumMilligram) GetPoints(st ScoreType) int {
	return getPointsFromRange(float64(sodium), sodiumLevels)

}

func (fruit FruitsPercent) GetPoints(st ScoreType) int {
	var multiplier int
	switch st {
	case Beverage:
		multiplier = 2
	default:
		multiplier = 1
	}

	points := 0
	if fruit >= 80 {
		points = 5 * multiplier
	} else if fruit >= 60 {
		points = 2 * multiplier
	} else if fruit >= 40 {
		points = 1 * multiplier
	}

	return points
	// if st == Beverage {
	// 	if fruit > 80 {
	// 		return 10
	// 	} else if fruit > 60 {
	// 		return 4
	// 	} else if fruit > 40 {
	// 		return 2
	// 	}
	// 	return 0
	// }
	// if fruit > 80 {
	// 	return 5
	// } else if fruit > 60 {
	// 	return 2
	// } else if fruit > 40 {
	// 	return 1
	// }
	// return 0
}

func (fibre FibreGram) GetPoints(st ScoreType) int {
	return getPointsFromRange(float64(fibre), fibreLevels)

}
func (protein ProteinGram) GetPoints(st ScoreType) int {
	return getPointsFromRange(float64(protein), proteinLevels)
}

func EnergyFromKcl(kcal float64) EnergyKJ {
	return EnergyKJ(kcal * 4.184)
}

func SodiumFromSalt(SaltMg float64) SodiumMilligram {
	return SodiumMilligram(SaltMg / 2.5)
}

func GetNutritionalScore(n NutricionalData, st ScoreType) NutritionalScore {

	value := 0
	positive := 0
	negative := 0

	if st != Water {
		fruitPoints := n.Fruits.GetPoints(st)
		fibrePoints := n.Fibre.GetPoints(st)

		negative = n.Energy.GetPoints(st) + n.Sugars.GetPoints(st) + n.SaturatedFattyaAcids.GetPoints(st) + n.Sodium.GetPoints(st)
		positive = fruitPoints + fibrePoints + n.Protein.GetPoints(st)

		if st == Cheese {
			value = negative - positive
		} else {
			if negative >= 11 && fruitPoints < 5 {
				value = negative - positive - fruitPoints
			} else {
				value = negative - positive
			}
		}
	}

	return NutritionalScore{
		Value:     value,
		Positive:  positive,
		Negative:  negative,
		ScoreType: st,
	}
}
func (ns NutritionalScore) GetNutriScore() string {
	if ns.ScoreType == Food {
		return scoreToLetter[getPointsFromRange(float64(ns.Value), []float64{18, 10, 2, -1})]

	}
	if ns.ScoreType == Water {
		return scoreToLetter[0]
	}
	return scoreToLetter[getPointsFromRange(float64(ns.Value), []float64{9, 5, 1, -2})]
}

func getPointsFromRange(v float64, steps []float64) int {
	lenSteps := len(steps)
	for i, l := range steps {
		if v > l {
			return lenSteps - i
		}
	}
	return 0
}
