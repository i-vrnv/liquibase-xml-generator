package main

import "encoding/xml"

const (
	// Header with version and encoding
	Header = `<?xml version="1.1" encoding="UTF-8" standalone="no"?>` + "\n"

	// Xmlns namespace
	Xmlns = `http://www.liquibase.org/xml/ns/dbchangelog`

	// XmlnsXsi Xsi namespace
	XmlnsXsi = `http://www.w3.org/2001/XMLSchema-instance`

	// XsiSchemaLocation schema location
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
	Columns   []Column `xml:"column"`
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
