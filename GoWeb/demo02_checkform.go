package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

func register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	if len(username) == 0 {
		fmt.Println("username cant be empty")
		fmt.Fprintf(w, "username cant be empty!")
	}
	age, err := strconv.Atoi(r.Form.Get("age"))
	if err != nil {
		fmt.Println("not number")
		fmt.Fprintf(w, "not number")
	}

	if age > 100 || age < 0 {
		//太大了或太小了
		fmt.Println("input the age must be at 0-100")
		fmt.Fprintf(w, "input the age must be at 0-100")
	}

	if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("age")); !m {
		fmt.Println("wrong")
		fmt.Fprintf(w, "wrong")
	}

	if m, _ := regexp.MatchString(`^[\x{4e00}-\x{9fa5}]+$`, r.Form.Get("zhname")); !m {
		fmt.Println("plz type chinese")
		fmt.Fprintf(w, "plz type chinese")
	}

	if m, _ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("enname")); !m {
		fmt.Println("plz type english")
		fmt.Fprintf(w, "plz type english")
	}

	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !m {
		fmt.Println("emgil address go wrong")
		fmt.Fprintf(w, "emgil address go wrong")
	}

	if m, _ := regexp.MatchString(`^(1[3|5|6|7|8][0-9]\d{8})$`, r.Form.Get("mobile")); !m {
		fmt.Println("phone number go wrong")
		fmt.Fprintf(w, "phone number go wrong")
	}

	xueli := r.Form.Get("xueli")
	res1 := checkSelect(xueli)
	if !res1 {
		fmt.Println("----------")
		fmt.Fprintf(w, "--------------")
	}

	sex := r.Form.Get("sex")
	res2 := checkSex(sex)
	if !res2 {
		fmt.Println("sex go wrong")
		fmt.Fprintf(w, "sex go wrong")
	}

	hobby := r.Form["hobby"]
	res3 := checkHobby(hobby)
	if !res3 {
		fmt.Println("----------")
		fmt.Fprintf(w, "----------")
	}


	if m, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, r.Form.Get("usercard")); !m {
		fmt.Println("id card go wrong")
		fmt.Fprintf(w, "id card go wrong")
	}

	//fmt.Println("验证成功！")
	//fmt.Fprintf(w, "验证成功！")

}


func checkSelect(xueli string) bool {
	slice := []string{"primary school", "junior high school", "High school", "College", "Undergraduate", "master's degree", "PhD"}
	for _, v := range slice {

		if v == xueli {
			return true
		}
	}
	return false
}


func checkSex(sex string) bool {
	slice := []string{"male", "female"}
	for _, v := range slice {
		if v == sex {
			return true
		}
	}
	return false
}


func checkHobby(hobby []string) bool {
	slice := []string{"game", "sport", "movie", "reading"}

	hobby2 := Slice_diff(hobby, slice)

	if hobby2 == nil {
		return true
	}
	return false
}

func Slice_diff(slice1, slice2 []string) (diffslice []string) {
	for _, v := range slice1 {
		if !InSlice(v, slice2) {
			diffslice = append(diffslice, v)
		}
	}
	return
}


func InSlice(val string, slice []string) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

func main() {
	http.HandleFunc("/register", register)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}