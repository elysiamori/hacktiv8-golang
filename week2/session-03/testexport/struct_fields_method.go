package testexport

type ExportableStruct struct {
	ExportableField   string
	inExportableField string
}

func (e *ExportableStruct) SetInexportableField(in string) {
	if e.inExportableField != "" {
		return
	}
	e.inExportableField = in
}

func (e *ExportableStruct) GetInexportableField() string {
	return e.inExportableField
}

func (e *ExportableStruct) deleteInexportableField(in string) {
	e.inExportableField = ""
}

type inExportableStruct struct {
}
