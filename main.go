package main

import (
	"encoding/json"
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
}

func initCache(filename string) lookupCacheCK {
	// Read file
	raw, _ := ioutil.ReadFile(filename)
	lc := lookupCacheCK{}
	// Initialize maps
	lc.data = make(map[string]map[string]map[string][]string)
	// Convert binary to string
	lc.orgJSON = string(raw)
	// Parse JSON and assign to our data structure
	// This interface hack is driving me nuts!
	var tree []interface{}
	json.Unmarshal([]byte(lc.orgJSON), &tree)
	lc.dataRaw = tree

	// Top level loop
	for i := range lc.dataRaw {
		orgMapNA := lc.dataRaw[i]
		// Type assertion
		orgMap := orgMapNA.(map[string]interface{})
		// OrgKey level loop
		for org := range orgMap {
			// Type assertion
			pNameInd := orgMap[org].([]interface{})
			for j := range pNameInd {
				// Type assertion
				pNameMap := pNameInd[j].(map[string]interface{})
				// paramName level loop
				for pName := range pNameMap {
					// Type assertion
					pValArray := pNameMap[pName].([]interface{})
					for k := range pValArray {
						// Type assertion
						pValMap := pValArray[k].(map[string]interface{})
						// paramVal level loop
						for pVal := range pValMap {
							if strings.Contains(pVal, "\n") {
								// Take care of compound keys
								pValSlice := strings.Split(pVal, "\n")
								// Type assertion
								segMap := pValMap[pVal].(map[string]interface{})
								// Segment level loop
								for _, segValNA := range segMap {
									// Type assertion
									segVal := segValNA.(string)
									if lc.data[org] == nil {
										// Initialize orgKey level map
										lc.data[org] = make(map[string]map[string][]string)
									}

									if lc.data[org][pName] == nil {
										// Initialize parameterName level map
										lc.data[org][pName] = make(map[string][]string)
									}

									for _, pv := range pValSlice {
										// Split compound keys and reassign values
										lc.data[org][pName][pv] = append(lc.data[org][pName][pv], segVal)
									}
								}
							} else {
								segMap := pValMap[pVal].(map[string]interface{})
								// Segment level loop
								for _, segValNA := range segMap {
									segVal := segValNA.(string)
									if lc.data[org] == nil {
										// Initilize orgKey level map
										lc.data[org] = make(map[string]map[string][]string)
									}

									if lc.data[org][pName] == nil {
										// Initialize parameterName level map
										lc.data[org][pName] = make(map[string][]string)
									}
									// Assign for non-compound keys
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

// Need two receiver implementations as specified since Go doesn't seem to support overloading
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
