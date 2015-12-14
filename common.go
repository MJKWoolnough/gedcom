package gedcom

import "strconv"

type Address struct {
	Addr     string        // ? Address Line (Cont)
	Addr1    string        // ? Address line 1
	Addr2    string        // ? Address line 2
	City     string        // ? Address city
	State    string        // ? Address state
	PostCode string        // ? Address Postal Code
	Country  string        // ? Address Country
	Phone    []PhoneNumber // ? Phone Number
}

type ErrTooMany struct {
	Tag string
	Max int
}

func setCheck(set *bool, name string) {
	if *set {
		panic(ErrTooMany{
			name,
			1,
		})
	}
	*set = true
}

func (e ErrTooMany) Error() string {
	return "too many " + e.Tag + " tags - maximum should be " + strconv.Itoa(e.Max)
}
