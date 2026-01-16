package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name    string `json:"Spyname"`
	Age     int
	Email   string
	Phone   int      `json:"-"`
	Hobbies []string `json",omitempty"`
}

func main() {
	fmt.Println("Herew we convert data into JSON")

	Data := []course{
		{"Shreyansh", 22, "shreyansh1418@gmail.com", 1223445556, []string{"Playing", "music", "Bgmi"}},
		{"uday", 25, "uday111@gmail.com", 14144141214, []string{"cricket", "freefire"}},
		{"rao", 29, "srao@gmail.com", 75775474, nil},
	}
	final, err := json.MarshalIndent(Data, "", "\t")
	check(err)
	fmt.Println(string(final))
	decodejson()
}

func decodejson() {
	jsondata := []byte(`
	    {
                "Spyname": "Shreyansh",
                "Age": 22,
                "Email": "shreyansh1418@gmail.com",
				"phone": 344443434,
                "Hobbies": ["Playing","music","Bgmi"]
        }
	`)
	valid :=json.Valid(jsondata)
    var format course
	if(valid){
		fmt.Println("JSON is valid")
		json.Unmarshal(jsondata, &format)
		fmt.Printf("%#v\n",format)
	}else{
		fmt.Println("JSON is not valid")
	}

	var record map[string]any
	json.Unmarshal(jsondata,&record)
	fmt.Printf("%#v\n",record)

	for k,v := range record{
		fmt.Printf("the kay is %v the value is %v and its type is %T\n",k,v,v)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}


//any works smae as interface{}
//any is an alias of interfece