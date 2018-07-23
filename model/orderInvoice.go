package model

import (
	"mugg/guapin/app/db"
)

// 订单发票表 (order_invoice)
// |-- 自动编号 (invoice_id)
// |-- 订单编号 (order_id)
// |-- 是否增值税发票 (is_vat, 普通发票,增值发票)
// |-- 发票抬头名称 (invoice_title)
// |-- 发票抬头内容 (invoice_content)
// |-- 发票金额 (invoice_amount)
// |-- 发票税号 (invoice_tax_no)
// |-- 开票税金 (invoice_tax)
// |-- 公司名称[增值税] (vat_company_name)
// |-- 公司地址[增值税] (vat_company_address)
// |-- 联系电话[增值税] (vat_telphone)
// |-- 开户银行[增值税] (vat_bank_name)
// |-- 银行帐号[增值税] (vat_bank_account)
// |-- 开票时间 (created_time)

type (
	// OrderInvoice is
	OrderInvoice struct {
		IDAutoModel
		OrderIDModel
		IsVat             bool   `json:"is_vat"`
		InvoiceTitle      string `json:"invoice_title"`
		InvoiceContent    string `json:"invoice_content"`
		InvoiceAmount     string `json:"invoice_amount"`
		InvoiceTaxNo      string `json:"invoice_tax_no"`
		InvoiceTax        string `json:"invoice_tax"`
		VatCompanyName    string `json:"vat_company_name"`
		VatCompanyAddress string `json:"vat_company_address"`
		VatTelphone       string `json:"vat_telphone"`
		VatBankName       string `json:"vat_bank_name"`
		VatBankAccount    string `json:"vat_bank_account"`
		TimeAllModel
	}
)

// Update is
func (m *OrderInvoice) Update(data *OrderInvoice) error {
	var (
		err error
	)
	tx := gorm.MysqlConn().Begin()
	if err = tx.Save(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}

// Delete is
func (m *OrderInvoice) Delete() error {
	var err error
	tx := gorm.MysqlConn().Begin()
	if err = tx.Delete(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}
