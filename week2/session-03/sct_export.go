package main

import (
	"fmt"
	"hacktiv8-go/session-03/testexport"
)

func sctExport() {
	fmt.Println(testexport.ExportableVariable)
	// fmt.Println(testexport.inExportableVariable)

	var exportableStruct testexport.ExportableStruct
	exportableStruct.ExportableField = "some string"
	exportableStruct.SetInexportableField("tidak bisa diekspor")
	exportableStruct.SetInexportableField("tidak bisa diekspor yang sudah diganti nilainya")
	fmt.Println("EXPORTABLE FIELD = ", exportableStruct.ExportableField)
	fmt.Println("INEXPORTABLE FIELD = ", exportableStruct.GetInexportableField())

	// var inExportableStruct testexport.inExportableStruct
	// fmt.Println(inExportableStruct)
}
