package main

import "fmt"

// Определяем интерфейс строителя отчета
type ReportBuilder interface {
	SetTitle(title string)
	AddText(text string)
	AddTable(table string)
	AddFooter(footer string)
	GetReport() string
}

// Строитель для PDF-отчета
type PDFBuilder struct {
	report string
}

func (b *PDFBuilder) SetTitle(title string) {
	b.report += "PDF Title: " + title + "\n"
}

func (b *PDFBuilder) AddText(text string) {
	b.report += "PDF Text: " + text + "\n"
}

func (b *PDFBuilder) AddTable(table string) {
	b.report += "PDF Table: " + table + "\n"
}

func (b *PDFBuilder) AddFooter(footer string) {
	b.report += "PDF Footer: " + footer + "\n"
}

func (b *PDFBuilder) GetReport() string {
	return b.report
}

// Строитель для HTML-отчета
type HTMLBuilder struct {
	report string
}

func (b *HTMLBuilder) SetTitle(title string) {
	b.report += "<h1>" + title + "</h1>\n"
}

func (b *HTMLBuilder) AddText(text string) {
	b.report += "<p>" + text + "</p>\n"
}

func (b *HTMLBuilder) AddTable(table string) {
	b.report += "<table>" + table + "</table>\n"
}

func (b *HTMLBuilder) AddFooter(footer string) {
	b.report += "<footer>" + footer + "</footer>\n"
}

func (b *HTMLBuilder) GetReport() string {
	return b.report
}

// Директор, который управляет процессом создания отчета
type ReportDirector struct {
	builder ReportBuilder
}

func (d *ReportDirector) SetBuilder(b ReportBuilder) {
	d.builder = b
}

func (d *ReportDirector) BuildReport() {
	d.builder.SetTitle("Annual Report")
	d.builder.AddText("This is the report content.")
	d.builder.AddTable("Table data: [1, 2, 3, 4]")
	d.builder.AddFooter("End of Report")
}

// Клиентский код
func main() {
	// Создаем директора
	director := &ReportDirector{}

	// Строим PDF отчет
	pdfBuilder := &PDFBuilder{}
	director.SetBuilder(pdfBuilder)
	director.BuildReport()
	pdfReport := pdfBuilder.GetReport()
	fmt.Println("PDF Report:")
	fmt.Println(pdfReport)

	// Строим HTML отчет
	htmlBuilder := &HTMLBuilder{}
	director.SetBuilder(htmlBuilder)
	director.BuildReport()
	htmlReport := htmlBuilder.GetReport()
	fmt.Println("HTML Report:")
	fmt.Println(htmlReport)
}
