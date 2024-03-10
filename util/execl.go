package util

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

// excel文件导出

// sheet中的数据

const defaultSheetName = "Sheet1"

type TableData struct {
	DataList     [][]interface{} `desc:"一个table的数据"`
	MarginBottom int             `desc:"下个表格和本表格的纵向距离"`
}
type SheetData struct {
	SheetName string      `desc:"sheet的名称"`
	TableList []TableData `desc:"一张sheet中的数据"`
}
type myExcel struct {
	*excelize.File

	headStyle int
}

type MyExcelOption func(*myExcel)

func MyExcelOptionHeadStyle() MyExcelOption {
	return func(me *myExcel) {
		headStyle, _ := me.NewStyle(&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "center",
			},
			Font: &excelize.Font{
				Bold: true,
			},
		})

		me.headStyle = headStyle
	}
}

func NewMyExcel(options ...MyExcelOption) *myExcel {
	me := &myExcel{
		File: excelize.NewFile(),
	}

	for _, option := range options {
		option(me)
	}

	return me
}

func (me *myExcel) Fill(sheetDataList ...SheetData) *myExcel {
	// 获取sheet的个数
	isCan := true
	for _, sheetData := range sheetDataList {
		if sheetData.SheetName == "" {
			sheetData.SheetName = defaultSheetName
			isCan = false
		}

		sheetIndex, _ := me.NewSheet(sheetData.SheetName)

		// 目前最多支持到A-Z
		rowIndex := 0
		for _, table := range sheetData.TableList {
			isHead := true
			for _, columnList := range table.DataList {
				rowIndex += 1       // 行
				column := rune('A') // 列
				for _, value := range columnList {
					// 每个table的表头设置样式
					if isHead {
						me.SetCellStyle(me.GetSheetName(sheetIndex), fmt.Sprintf("%c%d", column, rowIndex), fmt.Sprintf("%c%d", column, rowIndex), me.headStyle)
					}
					me.SetCellValue(me.GetSheetName(sheetIndex), fmt.Sprintf("%c%d", column, rowIndex), value)
					column += 1 // 行不变, 列加一
				}
				isHead = false
			}

			rowIndex += table.MarginBottom
		}
	}

	if isCan {
		me.DeleteSheet(defaultSheetName)
	}

	return me
}

func (me *myExcel) Save(fileName string) error {
	return me.SaveAs(fileName)
}
