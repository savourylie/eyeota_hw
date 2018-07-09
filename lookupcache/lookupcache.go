package lookupcache

//LookupCache is an interface your API should implement
type LookupCache interface {
	GetSegmentForOrgAndKey(orgKey string, paramKey string) []SegmentConfig
	GetSegmentForOrgAndKeyAndVal(orgKey string, paramKey string, paramVal string) []SegmentConfig
}

//SegmentConfig is a struct that holds an id for 1 segment
type SegmentConfig struct {
	Id string
}

// func (cfg *SegmentConfig) GetId() string {
//     return cfg.Id
// }
