package converter

import (
	"fmt"
	"github.com/samuelgarciastk/excel-converter/template"
	"github.com/tealeg/xlsx"
)

type Excel struct {
	Source   string
	Target   string
	Template template.Template
}

func (excel *Excel) Convert() {
	srcFile, err := xlsx.OpenFile(excel.Source)
	if err != nil {
		fmt.Printf("Cannot open source file: %s\n", excel.Source)
		panic(err)
	}
	srcSheetMap := srcFile.Sheet
	tgtFile := xlsx.NewFile()

	for sheetName, sheetTemplate := range excel.Template.Sheets {
		srcSheet := srcSheetMap[sheetName]

		tgtSheet, err := tgtFile.AddSheet(sheetName)
		if err != nil {
			fmt.Printf("Cannot add sheet: %s\n", sheetName)
			panic(err)
		}

		tgtRow := tgtSheet.AddRow()
		for _, header := range excel.Template.Header {
			tgtCell := tgtRow.AddCell()
			tgtCell.Value = header
		}

		for rowIndex, srcRow := range srcSheet.Rows {
			if rowIndex < sheetTemplate.Start-1 {
				continue
			}
			if rowIndex > sheetTemplate.End-1 {
				break
			}
			srcCells := srcRow.Cells
			tgtRow := tgtSheet.AddRow()
			for i := 0; i < len(excel.Template.Header); i++ {
				tgtRow.AddCell()
			}
			tgtCells := tgtRow.Cells
			for src, tgt := range sheetTemplate.Mapping {
				if tgt > len(tgtCells) {
					fmt.Printf("Write index [%d] out of range.\n", tgt)
				}
				if src <= len(srcCells) {
					tgtCells[tgt-1] = srcCells[src-1]
				}
			}
		}
	}

	err = tgtFile.Save(excel.Target)
	if err != nil {
		fmt.Printf("Cannot save file: %s\n", excel.Target)
		panic(err)
	}
}
