package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	//"strings"
)

type Food struct {
	Id string `json:"id"`
	Species     string `json:"species"`
	Description string `json:"description"`
}

type FoodIds struct {
	IDs string `json:""`
}

func getFoodHandler(w http.ResponseWriter, r *http.Request) {
	food, err := store.GetFood()

	foodListBytes, err := json.Marshal(food)

	if err != nil {
		fmt.Println(fmt.Errorf("Error %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(foodListBytes)
}

func createFoodHandler(w http.ResponseWriter, r *http.Request) {
	singleFood := Food{}

	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	singleFood.Species = r.Form.Get("species")
	singleFood.Description = r.Form.Get("description")

	err = store.CreateFood(&singleFood)
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/static/templates/food.html", http.StatusFound)
}

func testFoodHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/static/templates/food.html", http.StatusFound)
}

func deleteFoodHandler(w http.ResponseWriter, r *http.Request) {
	/*err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println(r.Body)*/
//	http.Redirect(w, r, "/static/templates/food.html", http.StatusFound)

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}


	// type DeleteItem struct {
	// 	val string `json:"Value"`
	// }
	fmt.Printf("%s", body)

	var deleteItems []string
	err = json.Unmarshal(body, &deleteItems)

	if err != nil {
		fmt.Println(fmt.Errorf("Error %v", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}


	retErr := ""
	//errDup := "duplicate value found"
	errMyPrecious := "you cannot delete noodles"
	code := 200
	for _, item := range deleteItems {
		if item == "0" {
			retErr = errMyPrecious
			code = 500
			break
		}
		// verify if parsing the string to int is succesful
	}

	type errResponse struct {
		ErrCode int `json:"code"`
		Message string `json:"message"`
	}

	if retErr != "" {
		res := errResponse{ErrCode: code, Message: retErr}
		w.WriteHeader(http.StatusBadRequest)

		bytes, err := json.Marshal(res)
		
		if err != nil{
			return
		}
		w.Write(bytes)
		return
	}

	//err = store.DeleteFood(strings.Join(deleteItems, ",")) 
	for _, item := range deleteItems {

		err = store.DeleteFood(item)

		if err != nil {
			fmt.Println(fmt.Errorf("Error %v", err))
			
			res := errResponse{ErrCode: 500, Message: err.Error()}
			bytes, err := json.Marshal(res)
			if err != nil{
				return
			}
			w.Write(bytes)
	
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	}
	
	w.Write([]byte("{\"code\": 200, \"message\": \"Success\"}"))
	w.WriteHeader(http.StatusOK)
}
