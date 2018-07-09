## Objective:
The objective of this assignment is to simulate a typical development task.  You will design and implement a component which encapsulates a data structure which will used to lookup reference data in a latency-sensitive data processing pipeline.

## Data Format:
The data which needs to be ingested is included in the attached file: **data.json**.  It is organized as follows:

* Outer Array of Organization Objects keyed by the orgKey and contain an array of ParameterNameObjects
    * ParameterNameObjects are keyed by parameterName and contain an array of 1...n SegmentConfig objects.
        * SegmentConfig objects are keyed by parameterValue which apply to that segment.  These can be an empty string, a single value, or a newline-delimited
string of values. Multiple segment objects in a ParameterNameObject can share the same key.  Internally, Segment Objects have a segmentId field.

### Functional Requirements:
Functionally, the following interface must be supported:

```golang
type LookupCache interface {
	GetSegmentForOrgAndKey(orgKey string, paramKey string) []SegmentConfig
	GetSegmentForOrgAndKeyAndVal(orgKey string, paramKey string, paramVal string) []SegmentConfig
}

//SegmentConfig is a struct that holds an id for 1 segment
type SegmentConfig struct {
	Id string
}

func (cfg *SegmentConfig) GetId() string {
	return cfg.Id
}
```

The LookupCache interface specifies the public API your cache should expose.  

The first function will return an array of SegmentConfig objects given an orgKey and a parameterName.  This call should return all SegmentConfigs associated
with that org and parameter, and with an empty parameterValue

The second function will return an array of SegmentConfig objects given an orgKey, a parameterName, and a parameterValue.  This call should return all SegmentConfigs associated
with that org, parameter, and parameter value.


### Example:
Given the following JSON Structure:
```json
[
    {
        "org1": [
            {
                "paramName1": [
                    {
                        "paramVal1": {
                            "segmentId": "seg_1234"
                        }
                    },
                    {
                        "paramVal2\nparamVal3\nparamVal4\nparamVal5": {
                            "segmentId": "intr.edu"
                        }
                    },
                    {
                        "paramVal6": {
                            "segmentId": "dem.infg.m"
                        }
                    },
                    {
                         "paramVal6": {
                             "segmentId": "intr.heal"
                         }
                    },
                    {
                          "paramVal6": {
                             "segmentId": "dem.infg.f"
                          }
                    }
                ]
            },
            {
                "testedu": [
                    {
                        "": {
                            "segmentId": "n277"
                        }
                    }
                ]
            },
            {
                "sid": [
                    {
                        "": {
                            "segmentId": "dem.life.expat"
                        }
                    }
                ]
            },
            {
                "gen": [
                    {
                        "Female": {
                            "segmentId": "dem.g.f"
                        }
                    },
                    {
                        "Male": {
                            "segmentId": "dem.g.m"
                        }
                    }
                ]
            }
        ]
    }
]
```
* The query getSegmentFor("org1", "paramName1")  will return an empty SegmentConfig array.
* The query getSegmentFor("org1", "paramName1", "paramVal1")  will return a 1-element SegmentConfigArray containing a SegmentConfig object for seg_1234
* The query getSegmentFor("org1", "paramName1", <<"paramVal2" OR  "paramVal3" OR  "paramVal4" OR  "paramVal5">> )  will return a 1-element SegmentConfigArray containing a SegmentConfig object with id: "intr.edu". (Note that in the data file paramVals may be in a "\n" delimited string.)
* The query getSegmentFor("org1", "paramName1", "paramVal6" )  will return a 3-element SegmentConfigArray containing SegmentConfig objects with ids: dem.infg.m, intr.heal, dem.infg.f
* The query getSegmentFor("org1", "testedu")  will return a 1-element SegmentConfigArray containing a SegmentConfig object with id "n277"
* The query getSegmentFor("org1", "testedu", "")  will return a 1-element SegmentConfigArray containing a SegmentConfig object with id "n277"
* The query getSegmentFor("org1", "testedu", <<any value other than an empty string>>) will return an empty SegmentConfig array.
* The query getSegmentFor("org1", "gen", "Female") will return a 1-element SegmentConfigArray containing a SegmentConfig object with id "dem.g.f"
* The query getSegmentFor("org1", "gen", "Male") will return a 1-element SegmentConfigArray containing a SegmentConfig object with id "dem.g.m"
* The query getSegmentFor("org1", "gen", <<any value other than "Male" or "Female">>) will return an empty SegmentConfig array.

## Nonfunctional requirements:
* The data structure should be as small as possible in memory
* **HIGH IMPORTANCE:** The design should be optimized for speed of lookup (i.e., the methods defined in the interface).
* **HIGH IMPORTANCE:** The design should be optimized to be as memory stable as possible for lookups; meaning high rates of lookup calls should not
 generate memory churn and significant gc activity.
## Notes:
*  Please implement your solution in your own project.  We have provided a basic library which should be imported into your solution as a dependency.  
* The go tool should be used for all dependency sourcing, building, test runs, etc... Do not use any other toolchain.  
* Do not change any of the supplied files.
* You can design your solution with the assumption that the data will fit in memory.
* All projects must build and should have runnable unit tests that demonstrate that the requirements have been met. only the standard go test framework should be used for tests.
* Solutions should target Go 1.9 and above.  No other dependencies (including external libraries) should be added. This includes both application and test code.
* You may use any resources available to you in terms of research, however you need to work alone.

Please submit your solution (without binaries) in the form of a tarball (created with command: tar -czvf) or a zip file to your Eyeota recruiting contact within 7 days of receiving these instructions.  


Good luck, and have fun!

The Eyeota Recruiting Team
