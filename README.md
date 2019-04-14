
Excel spreadsheet to LiquiBase XML converter
==
[![Build Status](https://travis-ci.org/voronovim/liquibase-xml-generator.svg?branch=master)](https://travis-ci.org/voronovim/liquibase-xml-generator)


Description
--
This repository contains a utility for creating a LiquiBase XML file from an Excel spreadsheet. A typical use case is the creation of initial ChangeSet by Excel's table. 

Important
--
This is my first Go app. I know the app probably has many problems and non-idiomatic solutions. So I would be appreciated for advice or PR with fixes. Thank you!

Requirements
--
go v1.11

Install
--

Download and unzip the project. In project root folder:
```sh
go build .
```

Usage
--
Prepare an Excel table with data. For now, the columns are hardcoded. It means that you should place your data in the table in proper order. Stick to the following order: Column name, type, nullable, remark. Spreadsheet name determines table name in DB. Then run the application.

Arguments
--
| Argument | Description | Default |
| ------------- |:-------------| :-----|
| -author | Set ChangeSet author name | liquibase |
| -id | Set ChangeSet id | 1 |
| source | Select the Excel's file with data | none |
| destination | Select the Excel's file with data | changeset-%id.xml |

To Do:
--
- Improve installation process
- Make columns more flexible
- Add more options for constraints(pk,fk,  etc.)
- 