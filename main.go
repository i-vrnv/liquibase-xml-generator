package main

import "fmt"
import "encoding/xml"
import "github.com/tealeg/xlsx"

const (
	// Header for liquidbase
	Header            = `<?xml version="1.1" encoding="UTF-8" standalone="no"?>` + "\n"
	Xmlns             = `http://www.liquibase.org/xml/ns/dbchangelog`
	XmlnsXsi          = `http://www.w3.org/2001/XMLSchema-instance`
	XsiSchemaLocation = `http://www.liquibase.org/xml/ns/dbchangelog http://www.liquibase.org/xml/ns/dbchangelog/dbchangelog-3.5.xsd`
)

// DatabaseChangeLog structure
type DatabaseChangeLog struct {
	XMLName           xml.Name  `xml:"databaseChangeLog"`
	Xmlns             string    `xml:"xmlns,attr"`
	XmlnsXsi          string    `xml:"xmlns:xsi,attr"`
	XsiSchemaLocation string    `xml:"xsi:schemaLocation,attr"`
	ChangeSet         ChangeSet `xml:"changeSet"`
}

// ChangeSet structure
type ChangeSet struct {
	XMLName     xml.Name    `xml:"changeSet"`
	Author      string      `xml:"author,attr"`
	ID          string      `xml:"id,attr"`
	CreateTable CreateTable `xml:"createTable"`
}

// CreateTable structure
type CreateTable struct {
	XMLName   xml.Name `xml:"createTable"`
	TableName string   `xml:"tableName,attr"`
	Remarks   string   `xml:"remarks,attr"`
	Column    Column   `xml:"column"`
}

// Column structure
type Column struct {
	XMLName     xml.Name    `xml:"column"`
	Name        string      `xml:"name,attr"`
	Type        string      `xml:"type,attr"`
	Remarks     string      `xml:"remarks,attr,omitempty"`
	Constraints Constraints `xml:"constraints"`
}

// Constraints structure
type Constraints struct {
	XMLName    xml.Name `xml:"constraints"`
	Nullable   bool     `xml:"nullable,attr"`
	PrimaryKey bool     `xml:"primaryKey,attr,omitempty"`
}

func main() {
	dataBaseChangeLog := createDataBaseChangeLog()
	dataBaseChangeLog.ChangeSet = ChangeSet{
		Author: "Me",
		ID:     "New Table",
		CreateTable: CreateTable{
			TableName: "new_big_table",
			Remarks:   "Table for nothing",
			Column: Column{
				Name: "code",
				Type: "varchar",
				Constraints: Constraints{
					PrimaryKey: false,
				},
			},
		},
	}

	if xmlString, err := xml.MarshalIndent(dataBaseChangeLog, "", "    "); err == nil {
		xmlString = []byte(Header + string(xmlString))
		fmt.Printf("%s\n", xmlString)
	}
}

func createDataBaseChangeLog() DatabaseChangeLog {
	dbcl := DatabaseChangeLog{
		Xmlns:             Xmlns,
		XmlnsXsi:          XmlnsXsi,
		XsiSchemaLocation: XsiSchemaLocation,
	}
	return dbcl
}

func extractData() {
	fmt.Printf("Start reading file...\n")
	excelFileName := "/home/user/Documents/GPB/test.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Printf("File reading error\n")
	}
	for _, sheet := range xlFile.Sheets {
		for i, row := range sheet.Rows {
			fmt.Printf("Start reading a row %v\n", i)
			for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%s\n", text)
			}
		}
	}
}
