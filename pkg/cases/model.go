package cases

type Case struct {
	ID             uint `gorm:"primaryKey"` // Every object should have ID
	PatientID      uint // A has many relationship should be on this
	DoctorID       uint
	Department     string
	Complaint      string // Use urls to locate pictures
	Diagnosis      string
	PastHistory    string
	Prescriptions  []Prescription
	PreviousCase   *Case // Previous case (the lastest one). If there is none prevous case, set nil
	PreviousCaseID *uint
}

type Prescription struct {
	ID         uint `gorm:"primarykey"`
	CaseID     uint
	Advice     string
	Guidelines []Guideline
}

type Guideline struct {
	ID       uint `gorm:"primarykey"`
	Medicine Medicine
	Dosage   string
	Quantity uint
}

type Medicine struct {
	ID               uint `gorm:"primarykey"`
	Name             string
	Price            float32
	Contraindication string
}