package entity

type Information struct {
	Id             int32 `pg:",pk"`
	Name           string
	Age            int32
	Phone          int64           `sql:",unique"`
	CollegeDetails *CollegeDetails `pg:"rel:has-one,fk:Id"`
}
type CollegeDetails struct {
	Id                 int32
	CollegeCode        string `sql:",unique"`
	CollegeName        string
	CollegeLocation    string
	CollegeContactInfo collagecontactinfo
}
type collagecontactinfo struct {
	Phone int64
	Email string
}
