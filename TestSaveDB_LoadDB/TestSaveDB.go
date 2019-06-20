package TestSaveDB_LoadDB
import  "github.com/json-iterator/go"
var json = jsoniter.ConfigCompatibleWithStandardLibrary

type TestStruct struct {
	Gg int64
	Hh string
}
//
func ( self  *TestStruct)SaveDB() (jsonStr string){
	jsonStr,error:=json.MarshalToString(self)
	if error!=nil {
		return jsonStr
	}
	return
}
//
func UnmarshalFromDBString(DBStr string,struct_ interface{}) ( interface_  interface{} ) {
	switch struct_.(type) {
	case TestStruct:
		interface_:=&TestStruct{}
		json.UnmarshalFromString(DBStr,interface_)
	}
	return
}

