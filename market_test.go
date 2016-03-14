package godatayes

import (
	"fmt"
	"testing"
)

func TestGetMktEqud(t *testing.T) {
	Init("b4be14a835d1550257a0423a58f1ec27b60ddd17e3ae7dade54e8976c6675697")
	data, err := GetMktEqudByTicker("000060", "20150202", "20160202")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(data)
}
