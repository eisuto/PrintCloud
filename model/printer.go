package model

// 打印机
type Printer struct {
	ID     int64  `json:"id"`
	Status string `json:"status"`
}

// 获取打印机列表
func (printer *Printer) Printers() (printers []Printer, err error) {
	if err = Db.Find(&printers).Error; err != nil {
		return
	}
	return
}
