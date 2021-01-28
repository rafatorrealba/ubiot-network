// Server to run web application to interact with Ubiot smart contract

package main

import (
	"fmt"
	"log"
	"path"
	"strings"
	"net/http"
	"html/template"
	"encoding/json"
	"fabcar/functions"

)

type BasicMachine struct {
	ID              string `json:"id"`
	Lessor          string `json:"lessor"`
	Status          string `json:"status"`
	ReservePrice    uint64 `json:"reserveprice"`
	WorkedHours     uint64 `json:"workedhours"`
	PricePerHour    uint64 `json:"priceperhour"`
	Lessee          string `json:"lessee"`
	RentalTime      string `json:"rentaltime"`
	PlaceOfDelivery string `json:"placeofdelivery"`
	WorkHours       uint64 `json:"workhours"`
}

func main() {
	
	getall := func(w http.ResponseWriter, r *http.Request) {

		result := functions.GetAll() // To JSON
		clean := string(result[1:(len(string(result))) - 1])

		str := strings.ReplaceAll(clean, "},", "},,")
		s := strings.Split(string(str), ",,")
		
		var bas []*BasicMachine
		for i, j := range s {
			if j == "" {}
			var ba BasicMachine
			_ = json.Unmarshal([]byte(s[i]), &ba)

			bas = append(bas, &ba)
			//fmt.Println(bas[(len(bas) - 1)], "\n")

			p := path.Join("static", "getAll.html")
			tmpl, err := template.ParseFiles(p)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if err := tmpl.Execute(w, bas[(len(bas) - 1)]); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			
		}

	}
	
	getmachine := func(w http.ResponseWriter, r *http.Request) {

		// Get data from HTML forms
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
		}
		id := r.PostFormValue("machineID")
		
		js := functions.GetMachine(id) // To JSON

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
				
		var ba BasicMachine
		err = json.Unmarshal(js, &ba)
		
		// Print response as HTML
		if  ba.ID == "" {
			p := path.Join("static", "base.html")
			tmpl, err := template.ParseFiles(p)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		
			if err := tmpl.Execute(w, "Machine does not exists"); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			
		} else {
			p := path.Join("static", "get.html")
			tmpl, err := template.ParseFiles(p)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		
			if err := tmpl.Execute(w, ba); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}	

	}

	newmachine := func(w http.ResponseWriter, r *http.Request) {

		// Get data from HTML forms
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
		}

		// Form data to function arguments
		id := r.PostFormValue("machineID")
		lessor := r.PostFormValue("lessor")
		reserveprice := r.PostFormValue("reserveprice")
		workedhours := r.PostFormValue("workedhours")
		priceperhour := r.PostFormValue("priceperhour")

		// Execute function with arguments
		js, err := json.Marshal(functions.NewMachine(id, lessor, reserveprice, workedhours, priceperhour))

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Print response as HTML
		p := path.Join("static", "base.html")
		tmpl, err := template.ParseFiles(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
		if err := tmpl.Execute(w, string(js)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	reservemachine := func(w http.ResponseWriter, r *http.Request) {

		// Get data from HTML forms
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
		}

		// Form data to function arguments
		id := r.PostFormValue("machineID")
		lesseeAdd := r.PostFormValue("lesseeAdd")
		rentaltimeAdd := r.PostFormValue("rentaltime")
		placeofdeliveryAdd := r.PostFormValue("placeofdeliveryAdd")

		// Execute function with arguments
		js, err := json.Marshal(functions.ReserveMachine(id, lesseeAdd, rentaltimeAdd, placeofdeliveryAdd))

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Print response as HTML
		p := path.Join("static", "base.html")
		tmpl, err := template.ParseFiles(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
		if err := tmpl.Execute(w, string(js)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	sentmachine := func(w http.ResponseWriter, r *http.Request) {

		// Get data from HTML forms
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
		}

		// Form data to function arguments
		id := r.PostFormValue("machineID")

		// Execute function with arguments
		js, err := json.Marshal(functions.SentMachine(id))

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Print response as HTML
		p := path.Join("static", "base.html")
		tmpl, err := template.ParseFiles(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
		if err := tmpl.Execute(w, string(js)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	receivedmachine := func(w http.ResponseWriter, r *http.Request) {

		// Get data from HTML forms
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
		}

		// Form data to function arguments
		id := r.PostFormValue("machineID")

		// Execute function with arguments
		js, err := json.Marshal(functions.ReceivedMachine(id))

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Print response as HTML
		p := path.Join("static", "base.html")
		tmpl, err := template.ParseFiles(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
		if err := tmpl.Execute(w, string(js)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	payperuse := func(w http.ResponseWriter, r *http.Request) {

		// Get data from HTML forms
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
		}

		// Form data to function arguments
		id := r.PostFormValue("machineID")
		workhours := r.PostFormValue("workhours")

		// Execute function with arguments
		js, err := json.Marshal(functions.PayPerUse(id, workhours))

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

			// Print response as HTML
		p := path.Join("static", "base.html")
		tmpl, err := template.ParseFiles(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
		if err := tmpl.Execute(w, string(js)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	returnmachine := func(w http.ResponseWriter, r *http.Request) {

		// Get data from HTML forms
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
		}

		// Form data to function arguments
		id := r.PostFormValue("machineID")

		// Execute function with arguments
		js, err := json.Marshal(functions.ReturnMachine(id))

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Print response as HTML
		p := path.Join("static", "base.html")
		tmpl, err := template.ParseFiles(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
		if err := tmpl.Execute(w, string(js)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	machineincompany := func(w http.ResponseWriter, r *http.Request) {

		// Get data from HTML forms
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
		}

		// Form data to function arguments
		id := r.PostFormValue("machineID")

		// Execute function with arguments
		js, err := json.Marshal(functions.MachineInCompany(id))

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Print response as HTML
		p := path.Join("static", "base.html")
		tmpl, err := template.ParseFiles(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
		if err := tmpl.Execute(w, string(js)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}

	machineinmaintenance := func(w http.ResponseWriter, r *http.Request) {

		// Get data from HTML forms
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
		}

		// Form data to function arguments
		id := r.PostFormValue("machineID")

		// Execute function with arguments
		js, err := json.Marshal(functions.MachineInMaintenance(id))

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Print response as HTML
		p := path.Join("static", "base.html")
		tmpl, err := template.ParseFiles(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
		if err := tmpl.Execute(w, string(js)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	availablemachine := func(w http.ResponseWriter, r *http.Request) {

		// Get data from HTML forms
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
		}

		// Form data to function arguments
		id := r.PostFormValue("machineID")

		// Execute function with arguments
		js, err := json.Marshal(functions.AvailableMachine(id))

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Print response as HTML
		p := path.Join("static", "base.html")
		tmpl, err := template.ParseFiles(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
		if err := tmpl.Execute(w, string(js)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	updatereserveprice := func(w http.ResponseWriter, r *http.Request) {

		// Get data from HTML forms
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
		}

		// Form data to function arguments
		id := r.PostFormValue("machineID")
		reservepriceadd := r.PostFormValue("reservepriceAdd")

		// Execute function with arguments
		js, err := json.Marshal(functions.ReservePrice(id, reservepriceadd))

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Print response as HTML
		p := path.Join("static", "base.html")
		tmpl, err := template.ParseFiles(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
		if err := tmpl.Execute(w, string(js)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	updatepriceperhour := func(w http.ResponseWriter, r *http.Request) {

		// Get data from HTML forms
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
		}

		// Form data to function arguments
		id := r.PostFormValue("machineID")
		priceperhour := r.PostFormValue("priceperhour")

		// Execute function with arguments
		js, err := json.Marshal(functions.UpdatePricePerHour(id, priceperhour))

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Print response as HTML
		p := path.Join("static", "base.html")
		tmpl, err := template.ParseFiles(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
		if err := tmpl.Execute(w, string(js)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	// Routing calls from the HTML file
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/getall", getall)
	http.HandleFunc("/getmachine", getmachine)
	http.HandleFunc("/newmachine", newmachine)
	http.HandleFunc("/reservemachine", reservemachine)
	http.HandleFunc("/sentmachine", sentmachine)
	http.HandleFunc("/receivedmachine", receivedmachine)
	http.HandleFunc("/payperuse", payperuse)
	http.HandleFunc("/returnmachine", returnmachine)
	http.HandleFunc("/machineincompany", machineincompany)
	http.HandleFunc("/machineinmaintenance", machineinmaintenance)
	http.HandleFunc("/availablemachine", availablemachine)
	http.HandleFunc("/updatereserveprice", updatereserveprice)
	http.HandleFunc("/updatepriceperhour", updatepriceperhour)

	fmt.Println("Serving in port 8081...")

	log.Fatal(http.ListenAndServe(":8081", nil))
}
