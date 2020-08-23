package database

import (
	"fmt"
)

type TableData struct {
	columns []*Column
	values  [][]*RowValue
}

func NewTableDataStruct(columns ...*Column) *TableData {

	return &TableData{
		columns: columns,
		values:  make([][]*RowValue, 0),
	}
}

// mohammed.ayubi@razorpay.com
func (td *TableData) AddDataRow(rowValues []*RowValue) error {

	for i, row := range rowValues {
		valid, err := td.columns[i].dType.IsValidValue(row.Value)

		if !valid {
			return err
		}
	}

	td.values = append(td.values, rowValues)
	return nil
}

func (td *TableData) FindDataRow(columnId int, expCellValue interface{}) ([][]*RowValue, error) {

	resultSet := make([][]*RowValue, 0)

	column := td.columns[columnId]

	if vOk, vErr := column.dType.IsValidValue(expCellValue); !vOk{
		return nil, fmt.Errorf("you tried to filter on invalid Value, error[%v]", vErr)
	}

	for _, cellValues := range td.values {
		if cellValues[columnId].Value == expCellValue {
			resultSet = append(resultSet, cellValues)
		}
	}

	return resultSet, nil
}

func (td *TableData) PrintData() {
	for _, column := range td.columns {
		fmt.Printf("'%s' ", column.name)
	}

	fmt.Println()

	for _, rowData := range td.values {

		for _, cell := range rowData {
			fmt.Printf("'%v' ", cell.Value)
		}
		fmt.Println()
	}
}
