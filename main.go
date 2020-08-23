package main

import (
	"fmt"

	"InMemorySqlDB/internal/database"
	"InMemorySqlDB/internal/types"
)

func main() {
	db := database.NewDatabase("sample")

	column1, err1 := database.NewColumn("kidName", types.STRING, false, 5, 15)
	if err1 != nil {
		fmt.Printf("invalid column err:[%v]", err1)
	}

	column2, err2 := database.NewColumn("kidAge", types.INT, false, 5, 15)
	if err2 != nil {
		fmt.Printf("invalid column err:[%v]", err2)
	}

	// create table
	table := database.NewTableDataStruct(column1, column2)

	// add table
	db.AddTable("kids", table)


	iErr1 := table.AddDataRow([]*database.RowValue{
		{
			Value: "tushar",
		},
		{
			Value: 10,
		},
	})
	if iErr1 != nil {
		fmt.Printf("invalid row err:[%v]", iErr1)
	}

	iErr2 := table.AddDataRow([]*database.RowValue{
		{
			Value: "someone",
		},
		{
			Value: 12,
		},
	})

	if iErr2 != nil {
		fmt.Printf("invalid row err:[%v]", iErr2)
	}

	row, sErr := table.FindDataRow(0, "tushar")
	if sErr != nil {
		fmt.Printf("invalid query err:[%v]", sErr)
	} else {
		fmt.Printf("data: [%v]", row)
	}

}
