package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/elgs/gojq"
)

type lookupCacheCK struct {
	orgJSON string
	data    []interface{}
	dataCK  map[string]map[string]map[string][]string
}

func main() {
	filename := "./data/test.json"
	// filename := "./data/data.json"
	lc := initCache(filename)

	fmt.Println(lc.GetSegmentForOrgAndKey("org1", "paramName1"))

	// raw, err := ioutil.ReadFile(filename)

	// l := lookupCacheCK{}
	// l.orgJSON = string(raw)
	// parser, err := gojq.NewStringQuery(l.orgJSON)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// tree, err := parser.Query(".")
	// l.data = tree.([]interface{})
}

func Keys(m map[string]interface{}) (keys []string) {
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func initCache(filename string) lookupCacheCK {
	raw, err := ioutil.ReadFile(filename)
	lc := lookupCacheCK{}

	// Initialize maps
	lc.dataCK = make(map[string]map[string]map[string][]string)
	// lc.keyID = make(map[string]int)
	// lc.keyParamNameID = make(map[string][]int)
	// lc.paramNameID = make(map[string]int)

	lc.orgJSON = string(raw)
	parser, err := gojq.NewStringQuery(lc.orgJSON)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	tree, err := parser.Query(".")
	lc.data = tree.([]interface{})

	// Top level loop
	for i := range lc.data {
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
							segMap := pValMap[pVal].(map[string]interface{})
							// Segment level loop
							for _, segValNA := range segMap {
								segVal := segValNA.(string)

								if lc.dataCK[org] == nil {
									lc.dataCK[org] = make(map[string]map[string][]string)
								}

								if lc.dataCK[org][pName] == nil {
									lc.dataCK[org][pName] = make(map[string][]string)
								}
								lc.dataCK[org][pName][pVal] = append(lc.dataCK[org][pName][pVal], segVal)
							}
						}
					}
				}
			}
		}
	}

	// fmt.Println(lc.dataCK)

	return lc
}

func (l lookupCacheCK) GetSegmentForOrgAndKey(orgKey string, paramKey string) []string {
	pValMap := l.dataCK[orgKey][paramKey]
	
	resultString := []string{}
	
	for _, segs := range pValMap {
		resultString = append(resultString, segs...)
	}

	result := []SegmentConfig{}

	for _, seg := range resultString {
		result = append(SegmentConfig{ID})
	}
	return result
}

// raw, err := ioutil.ReadFile("./data/data.json")
0
func parseJSON(filename string) string {
	// Get raw []byte
	raw, err := ioutil.ReadFile(filename)

	// Exception handling
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	orgJSON := raw

	var f interface{}
	err = json.Unmarshal(orgJSON, &f)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	m := f.(map[string]interface{})

	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}

	return "Test completed"
}

// 	var result map[string]interface{}
// 	json.Unmarshal([]byte(orgJSON), &result)

// 	orgs := result["orgs"].(map[string]interface{})

// 	for key, value := range orgs {
// 		// Each value is an interface{} type, that is type asserted as a string
// 		fmt.Println(key, value.(string))
// 	}
// 	// var c []string
// json.Unmarshal(raw, &c)
// 	// fmt.Println(c)

// 	// return c
// }

// type LookupCache interface {
// 	GetSegmentForOrgAndKey(orgKey string, paramKey string) []SegmentConfig
// 	GetSegmentForOrgAndKeyAndVal(orgKey string, paramKey string, paramVal string) []SegmentConfig
// }

// func (cfg *SegmentConfig) GetId() string {
// 	return cfg.Id
// }
