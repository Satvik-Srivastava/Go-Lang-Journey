package main

import (
	"encoding/json"
	"fmt"
)

//
// we can also create aliases for struct fields using json tags using ``
type course struct { // course is small that means unexported
	Name     string `json:"coursename"`
	Price    int 	`json:"courseprice"`
	Platform string `json:"website"`
	Password string	`json:"-"` // this field will be ignored during marshaling
	Tags     []string	`json:"tags,omitempty"` // if the field is empty then it will be omitted
}

func main() {
	fmt.Println("Welcome to some more talk on json handling")
	// EncodeJson()
	DecodeJson()	
}

func EncodeJson() {
	lcoCourses := []course{
		{
			"Learn Go Lang",
			299,
			"learncodeonline.in",
			"abc123",
			[]string {"web-dev", "mobile-dev", "game-dev"},
		},
		{
			"Learn HTML Lang",
			299,
			"learncodeonline.in",
			"htmlcourse",
			[]string {"web-dev", "mobile-dev", "game-dev"},
		},
		{
			"Learn Python Lang",
			299,
			"learncodeonline.in",
			"pythoncourse",
			nil,
		},
	}

	// wraping this data as json and send it to the web server


	// marshal just the exported fields in raw format while json.MarshalIndent formats the json output with indentation
	
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	CheckNilErr(err)

	fmt.Printf("%s\n", finalJson)



}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}


func DecodeJson(){
	jsondatafromweb := []byte(`
		{
			"coursename": "Learn Go Lang",
            "courseprice": 299,
            "website": "learncodeonline.in",
            "tags": ["web-dev","mobile-dev","game-dev"]
        }
	`)


	var locCourse course 
	checkValid := json.Valid(jsondatafromweb)

	if checkValid {
		fmt.Println("Json is valid")
		json.Unmarshal(jsondatafromweb, &locCourse)
		fmt.Printf("%#v\n", locCourse)

	}else{
		fmt.Println("The json was incorrect")
	}
	
	
	// some other use case

	var myJson map[string] interface {}
	json.Unmarshal(jsondatafromweb, &myJson)
	fmt.Printf("%#v\n", myJson)

	for key, val := range myJson {
		fmt.Printf("The key is %v and value is %v\n", key, val)
	}

}