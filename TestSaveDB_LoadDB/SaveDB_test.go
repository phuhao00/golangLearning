package TestSaveDB_LoadDB

//
func ExampleTestStruct_SaveDB() {
	struct_:=&TestStruct{}
	jsonStr:=struct_.SaveDB()

}
