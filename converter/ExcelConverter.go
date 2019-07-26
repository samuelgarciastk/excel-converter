package converter

import (
	"fmt"
	"github.com/samuelgarciastk/excel-converter/utils"
	"github.com/tealeg/xlsx"
	"log"
)

type Excel struct {
	Source      string
	Destination string
	Template    utils.Template
}

func (excel *Excel) Convert() error {
	srcFile, err := xlsx.OpenFile(excel.Source)
	if err != nil {
		return fmt.Errorf("cannot open source file: %s, due to %v", excel.Source, err)
	}
	srcSheetMap := srcFile.Sheet

	if err = utils.CopyFile(excel.Template.File, excel.Destination); err != nil {
		return fmt.Errorf("cannot copy template file: %s, due to %v", excel.Template.File, err)
	}

	dstFile, err := xlsx.OpenFile(excel.Destination)
	if err != nil {
		return fmt.Errorf("cannot open destination file: %s, due to %v", excel.Destination, err)
	}
	dstSheetMap := dstFile.Sheet

	for srcSheetName, sheetTmpl := range excel.Template.Sheets {
		srcSheet := srcSheetMap[srcSheetName]
		dstSheet := dstSheetMap[sheetTmpl.DstSheet]
		if srcSheet == nil || dstSheet == nil {
			log.Printf("source sheet [%s] or destination sheet [%s] not found", srcSheetName, sheetTmpl.DstSheet)
			continue
		}

		dstRowIdx := sheetTmpl.DstStart - 1
		for srcRowIdx, srcRow := range srcSheet.Rows {
			if srcRowIdx < sheetTmpl.SrcStart-1 {
				continue
			}
			if srcRowIdx > sheetTmpl.SrcEnd-1 {
				break
			}
			srcCells := srcRow.Cells
			dstRow := dstSheet.Row(dstRowIdx)
			dstRowIdx++
			dstCells := dstRow.Cells

			for srcCellIdx, dstCellIdx := range sheetTmpl.Mapping {
				if dstCellIdx > len(dstCells) {
					for i := 0; i < dstCellIdx-len(dstCells); i++ {
						dstRow.AddCell()
					}
					dstCells = dstRow.Cells
				}
				if srcCellIdx <= len(srcCells) {
					dstCells[dstCellIdx-1].SetString(srcCells[srcCellIdx-1].String())
				}
			}
		}
	}

	if err = dstFile.Save(excel.Destination); err != nil {
		return fmt.Errorf("cannot save converted file: %s, due to %v", excel.Destination, err)
	}
	return nil
}
