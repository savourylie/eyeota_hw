package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"./lookupcache"
	"github.com/elgs/gojq"
)

type lookupCacheCK struct {
	orgJSON string
	dataRaw []interface{}
	data    map[string]map[string]map[string][]string
}

func main() {
	filename := "./data/test.json"
	// parseJSON(filename)

	// filename := "./data/data.json"
	// _ = initCache(filename)
	lc := initCache(filename)

	// fmt.Println(lc.GetSegmentForOrgAndKey("org1", "paramName1"))
	// lc.GetSegmentForOrgAndKey("org1", "paramName1")

	// fmt.Println(lc.GetSegmentForOrgAndKeyAndVal("org1", "paramName1", "paramVal1"))
	// lc.GetSegmentForOrgAndKeyAndVal("org1", "paramName1", "paramVal2")
	fmt.Println(lc.GetSegmentForOrgAndKeyAndVal("org1", "paramName1", "paramVal2\nparamVal3\nparamVal4\nparamVal5"))
	fmt.Println(lc.GetSegmentForOrgAndKeyAndVal("org1", "paramName1", "paramVal2"))
	fmt.Println(lc.GetSegmentForOrgAndKeyAndVal("org1", "paramName1", "paramVal3"))
	fmt.Println(lc.GetSegmentForOrgAndKeyAndVal("org1", "paramName1", "paramVal4"))
	fmt.Println(lc.GetSegmentForOrgAndKeyAndVal("org1", "paramName1", "paramVal5"))
	// fmt.Println(lc.GetSegmentForOrgAndKeyAndVal("org1", "paramName1", "paramVal5"))

}

// func parseJSON(filename string) {
// 	raw, _ := ioutil.ReadFile(filename)
// 	orgJSON := string(raw)

// 	// var result []map[string][]map[string][]map[string]string
// 	var result []interface{}
// 	json.Unmarshal([]byte(orgJSON), &result)

// 	orgArray := result

// 	for i := range orgArray {
// 		fmt.Println(i)
// 		orgMap := orgArray[i].(map[string]interface{})
// 		for org := range orgMap {
// 			fmt.Println(org)
// 			pNameInd := orgMap[org].([]interface{})
// 			for j := range pNameInd {
// 				pNameMap := pNameInd[j].(map[string]interface{})
// 				// paramName level loop
// 				for pName := range pNameMap {
// 					pValArray := pNameMap[pName].([]interface{})
// 					for k := range pValArray {
// 						pValMap := pValArray[k].(map[string]interface{})
// 						// paramVal level loop
// 						for pVal, seg := range pValMap {
// 							if strings.Contains(pVal, "\n") {
// 								pValSlice := strings.Split(pVal, "\n")
// 							}
// 							// fmt.Println(pVal, seg)
// 							segMap := pValMap[pVal].(map[string]interface{})
// 							// Segment level loop
// 							for _, segValNA := range segMap {
// 								segVal := segValNA.(string)
// 								fmt.Println(segVal)
// 							}
// 						}
// 					}
// 				}
// 			}

// 		}
// 	}

// fmt.Println(result[0])
// fmt.Println(result[0]["org1"])
// fmt.Println(result[0]["org1"][0])
// fmt.Println(result[0]["org1"][0]["paramName1"])
// fmt.Println(result[0]["org1"][0]["paramName1"][0])
// fmt.Println(result[0]["org1"][0]["paramName1"][0]["paramVal1"])

// fmt.Println(result[0]["org1"][0]["paramName1"][1])
// fmt.Println(result[0]["org1"][0]["paramName1"][1]["paramVal2\nparamVal3\nparamVal4\nparamVal5"])

// for key, value := range result {
// 	// Each value is an interface{} type, that is type asserted as a string
// 	fmt.Println(key, value)
// }

// }

func initCache(filename string) lookupCacheCK {
	raw, err := ioutil.ReadFile(filename)
	lc := lookupCacheCK{}

	// Initialize maps
	lc.data = make(map[string]map[string]map[string][]string)

	lc.orgJSON = string(raw)
	parser, err := gojq.NewStringQuery(lc.orgJSON)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	tree, err := parser.Query(".")
	lc.dataRaw = tree.([]interface{})

	// Top level loop
	for i := range lc.dataRaw {
		// lc.orgID = append(lc.orgID, i)
		orgMapNA, _ := parser.Query("[" + strconv.Itoa(i) + "]")
		orgMap := orgMapNA.(map[string]interface{})

		// OrgKey level loop
		for org := range orgMap {
			pNameInd := orgMap[org].([]interface{})
			for j := range pNameInd {
				pNameMap := pNameInd[j].(map[string]interface{})
				// paramName level loop
				for pName := range pNameMap {
					pValArray := pNameMap[pName].([]interface{})
					for k := range pValArray {
						pValMap := pValArray[k].(map[string]interface{})
						// paramVal level loop
						for pVal := range pValMap {
							// fmt.Println(pVal, seg)
							if strings.Contains(pVal, "\n") {
								pValSlice := strings.Split(pVal, "\n")

								segMap := pValMap[pVal].(map[string]interface{})
								// Segment level loop
								for _, segValNA := range segMap {
									segVal := segValNA.(string)

									if lc.data[org] == nil {
										lc.data[org] = make(map[string]map[string][]string)
									}

									if lc.data[org][pName] == nil {
										lc.data[org][pName] = make(map[string][]string)
									}

									for _, pv := range pValSlice {
										lc.data[org][pName][pv] = append(lc.data[org][pName][pv], segVal)
									}

								}
							} else {
								segMap := pValMap[pVal].(map[string]interface{})
								// Segment level loop
								for _, segValNA := range segMap {
									segVal := segValNA.(string)

									if lc.data[org] == nil {
										lc.data[org] = make(map[string]map[string][]string)
									}

									if lc.data[org][pName] == nil {
										lc.data[org][pName] = make(map[string][]string)
									}
									lc.data[org][pName][pVal] = append(lc.data[org][pName][pVal], segVal)
								}
							}
						}
					}
				}
			}
		}
	}
	return lc
}

func (lc lookupCacheCK) GetSegmentForOrgAndKey(orgKey string, paramKey string) []lookupcache.SegmentConfig {
	resultString := lc.data[orgKey][paramKey][""]
	// pValMap := lc.data[orgKey][paramKey]
	// resultString := []string{}

	// if len(pVal)
	// for key, segs := range pValMap {
	// 	fmt.Println("==============")
	// 	fmt.Println(key, pValMap[key])
	// 	resultString = append(resultString, segs...)
	// }

	// fmt.Println(resultString)

	result := []lookupcache.SegmentConfig{}

	for _, seg := range resultString {
		result = append(result, lookupcache.SegmentConfig{Id: seg})
	}
	return result
}

func (lc lookupCacheCK) GetSegmentForOrgAndKeyAndVal(orgKey string, paramKey string, paramVal string) []lookupcache.SegmentConfig {
	resultString := lc.data[orgKey][paramKey][paramVal]
	result := []lookupcache.SegmentConfig{}

	for _, seg := range resultString {
		result = append(result, lookupcache.SegmentConfig{Id: seg})
	}
	return result
}

// func Keys(m map[string]interface{}) (keys []string) {
// 	for k := range m {
// 		keys = append(keys, k)
// 	}
// 	return keys
// }
