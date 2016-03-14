package godatayes

import (
	"fmt"
	"testing"

)

func TestGetData(t *testing.T) {
	Init("b4be14a835d1550257a0423a58f1ec27b60ddd17e3ae7dade54e8976c6675697")
	path := "/api/macro/getChinaDataGDP.csv?field=&indicID=M010000002&indicName=&beginDate=&endDate="
	data, err := getData(path)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(data)
}
