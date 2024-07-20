package odd_number

import "test-neversitup/model"

func FindOddNumber(text []int) int {
	// TODO : start your code here
	var ResultData []model.CollectData

	for _, data := range text {
		var foundData bool
		///check data from array
		for i, collect := range ResultData {
			///if have data in array update count data
			if data == collect.Value {
				ResultData[i].Count += 1
				foundData = true
			}
		}
		//if not have data add new data
		if foundData == false {
			result := model.CollectData{
				Value: data,
				Count: 1,
			}
			ResultData = append(ResultData, result)
		}
	}

	var FinishData int
	/// check data from array ResultData have count is odd number
	for _, jjj := range ResultData {
		if jjj.Count%2 != 0 {
			FinishData = jjj.Value
		}
	}
	return FinishData
}
