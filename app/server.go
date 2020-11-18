// Server to run web application to interact with Ubiot smart contract

package main

import (
	"fmt"
	"log"
	"strings"
	"net/http"
	"encoding/json"
	"fabcar/functions"
)

func main() {
	getmachine := func(w http.ResponseWriter, r *http.Request) {
		
		// Get data from HTML forms
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
	   	}
	   	id := r.PostFormValue("machineID")
		js, err := json.Marshal(functions.GetMachine(id)) // To JSON

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		// Removing unnecessary symbols to the message
		cleanjs := strings.Replace(string(js), "\\", "", -1)
		
		if cleanjs == "" {
			w.Write([]byte("Machine does not exist"))
			return
		}

		// Print received message in the web screen
		w.Write([]byte(strings.Replace(cleanjs, "}},", "}}\n", -1)))
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

		// Print received message in the web screen
		w.Write(js)
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
		
		// Print received message in the web screen
		w.Write(js)
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
		
		// Print received message in the web screen
		w.Write(js)
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
		
		// Print received message in the web screen
		w.Write(js)
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
		
		// Print received message in the web screen
		w.Write(js)
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
		
		// Print received message in the web screen
		w.Write(js)
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
		
		// Print received message in the web screen
		w.Write(js)
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
		
		// Print received message in the web screen
		w.Write(js)
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
		
		// Print received message in the web screen
		w.Write(js)
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
		
		// Print received message in the web screen
		w.Write(js)
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
		
		// Print received message in the web screen
		w.Write(js)
	}

	// Routing calls from the HTML file
	http.Handle("/", http.FileServer(http.Dir("./static")))
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
	
	fmt.Println("Serving in port 8080...")

	log.Fatal(http.ListenAndServe(":8080", nil))
}