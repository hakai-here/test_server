package utils

import (
	"demoproject/api/db"
	"demoproject/api/structs"
	"encoding/xml"
	"io"
	"log"
	"os"

	"github.com/spf13/viper"
)

// the code inside the file is used to parse and add data to postgres sql database

func InsertData() {
	k := ParseXml()
	db.InsertMany(k.ProceedingInformation.ProceedinGentry)
}

func ParseXml() structs.Xmldata { // parsing the XML data in the data folder
	xmlpath, err := os.Open(viper.GetString("XMLPATH"))
	if err != nil {
		log.Fatal(err)
	}

	defer xmlpath.Close()

	xmlData, err := io.ReadAll(xmlpath)
	if err != nil {
		log.Fatalf("[ERROR] Error occured while reading the xml file : %s", err.Error())
	}

	var data structs.Xmldata
	err = xml.Unmarshal(xmlData, &data)

	if err != nil {
		log.Fatalf("[ERROR] Error occured while parsing the xml file : %s", err.Error())
	}

	return data
}
