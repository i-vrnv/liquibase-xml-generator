package main

import (
	"bufio"
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"regexp"

	"github.com/tealeg/xlsx"
)

var author = flag.String("author", "liquibase", "Set ChangeSet author name")
var id = flag.String("id", "1", "Set ChangeSet id")
var sourceFile = flag.String("source", "", "Select the Excel's file with data")
var destinationFile = flag.String("destination", fmt.Sprintf("changeset-%s.xml", *id), "Change filename of output XML file")

func main() {
	flag.Parse()

	fmt.Printf("Start reading file %s...\n", *sourceFile)
	xlsFile, err := xlsx.OpenFile(*sourceFile)
	if err != nil {
		fmt.Printf("File reading error\n")
		os.Exit(1)
	}

	changeSet := createChangeSetFromFile(xlsFile)
	dataBaseChangeLog := createDataBaseChangeLog(changeSet)
	xml := generateXML(dataBaseChangeLog)

	writeToFile(*destinationFile, changeCloseTags(xml))

}

func (table *CreateTable) addColumn(column Column) {
	table.Columns = append(table.Columns, column)
}

func createChangeSetFromFile(xlsFile *xlsx.File) ChangeSet {
	var changeSet ChangeSet
	for _, sheet := range xlsFile.Sheets {
		table := CreateTable{
			TableName: sheet.Name,
		}
		for _, row := range sheet.Rows[1:] {
			column := Column{
				Name:    row.Cells[0].String(),
				Type:    row.Cells[1].String(),
				Remarks: row.Cells[3].String(),
				Constraints: Constraints{
					Nullable: row.Cells[2].Bool(),
				},
			}
			table.addColumn(column)
		}
		changeSet = ChangeSet{
			Author:      *author,
			ID:          *id,
			CreateTable: table,
		}
	}
	return changeSet
}

func createDataBaseChangeLog(changeSet ChangeSet) DatabaseChangeLog {
	dbcl := DatabaseChangeLog{
		Xmlns:             Xmlns,
		XmlnsXsi:          XmlnsXsi,
		XsiSchemaLocation: XsiSchemaLocation,
		ChangeSet:         changeSet,
	}
	return dbcl
}

func generateXML(dbChangelog DatabaseChangeLog) string {
	xmlBytes, err := xml.MarshalIndent(dbChangelog, "", "    ")
	if err != nil {
		panic(err)
	}
	return Header + string(xmlBytes)
}

func writeToFile(fileName string, xml string) {
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	w := bufio.NewWriter(f)
	b, err := w.WriteString(xml)
	if b < len(xml) {
		if err != nil {
			panic(err)
		}
	}
	w.Flush()
}

func changeCloseTags(xml string) string {
	r := regexp.MustCompile("></[[:alnum:]]*>")
	return r.ReplaceAllString(xml, "/>")
}
