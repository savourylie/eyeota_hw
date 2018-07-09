package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"./lookupcache"
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

	// fmt.Println(lc.data["org1"])
	// fmt.Println(lc.GetSegmentForOrgAndKey("org1", "paramName1"))
	// lc.GetSegmentForOrgAndKey("org1", "paramName1")

	fmt.Println(lc.GetSegmentForOrgAndKeyAndVal("org1", "paramName1", "paramVal1"))
	// lc.GetSegmentForOrgAndKeyAndVal("org1", "paramName1", "paramVal2")

	// fmt.Println(lc.GetSegmentForOrgAndKeyAndVal("org1", "paramName1", "paramVal2\nparamVal3\nparamVal4\nparamVal5"))
	fmt.Println(lc.GetSegmentForOrgAndKeyAndVal("org1", "paramName1", "paramVal2"))
	// fmt.Println(lc.GetSegmentForOrgAndKeyAndVal("org1", "paramName1", "paramVal3"))
	// fmt.Println(lc.GetSegmentForOrgAndKeyAndVal("org1", "paramName1", "paramVal4"))
	// fmt.Println(lc.GetSegmentForOrgAndKeyAndVal("org1", "paramName1", "paramVal5"))
	// fmt.Println(lc.GetSegmentForOrgAndKeyAndVal("org1", "paramName1", "paramVal5"))

}

func initCache(filename string) lookupCacheCK {
	raw, _ := ioutil.ReadFile(filename)
	lc := lookupCacheCK{}

	// Initialize maps
	lc.data = make(map[string]map[string]map[string][]string)

	lc.orgJSON = string(raw)

	var tree []interface{}
	json.Unmarshal([]byte(lc.orgJSON), &tree)
	lc.dataRaw = tree

	// fmt.Println(lc.dataRaw)

	// Top level loop
	for i := range lc.dataRaw {
		orgMapNA := lc.dataRaw[i]
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
