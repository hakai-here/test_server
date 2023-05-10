package structs

type Xmldata struct {
	Version struct {
		Versionno   string `xml:"version-no" json:"version-no"`
		Versiondate uint   `xml:"version-date" json:"version-date"`
	} `xml:"version" json:"version"`
	Actionkeycode         string                `xml:"action-key-code" json:"action-key-code"`
	Transactiondate       uint                  `xml:"transaction-date" json:"transaction-date"`
	ProceedingInformation Proceedinginformation `xml:"proceeding-information" json:"proceeding-information"`
}

type Proceedinginformation struct {
	ProceedinGentry []Proceedingentry `xml:"proceeding-entry" json:"proceeding-entry"`
}

type Proceedingentry struct {
	Number                    uint   `xml:"number" json:"number" gorm:"primaryKey"`
	Typecode                  string `xml:"type-code" json:"type-code"`
	Filingdate                uint   `xml:"filing-date" json:"filing-date"`
	Employeenumber            uint   `xml:"employee-number" json:"employee-number"`
	Interlocutoryattorneyname string `xml:"interlocutory-attorney-name" json:"interlocutory-attorney-name"`
	Locationcode              string `xml:"location-code" json:"location-code"`
	Dayinlocation             uint   `xml:"day-in-location" json:"day-in-location"`
	Statusupdatedate          uint   `xml:"status-update-date" json:"status-update-date"`
	Statuscode                uint   `xml:"status-code" json:"status-code"`
}
