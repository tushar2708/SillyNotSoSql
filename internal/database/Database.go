package database


type Database struct {
	tableMap map[string]*TableData
}


func NewDatabase(name string) *Database{
	return &Database{
		tableMap: make(map[string]*TableData, 0),
	}
}

func (d *Database) AddTable(name string, data *TableData){
	d.tableMap[name] = data
}
