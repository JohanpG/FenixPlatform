package handler

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"FenixPlatform/com.cambiosfenix.ftm.service/controller"
	"FenixPlatform/com.cambiosfenix.ftm.service/domain"
	"FenixPlatform/com.cambiosfenix.ftm.service/log"
	"FenixPlatform/com.cambiosfenix.ftm.service/panics"
	"FenixPlatform/com.cambiosfenix.ftm.service/util"
)

/*
Descripcion :
	Este metodo tiene como objetivo definir todos los mapeos entres URL y manegadores.
*/
func SetUpHandlers() {
	//se define el mapping entre las URL y los metodos
	// CUD Operations
	//http.HandleFunc("/CreatePartner", createPartner)
	//http.HandleFunc("/UpdatePartner", updatePartner)
	//http.HandleFunc("/DeletePartner", deletePartner)
	http.HandleFunc("/CreateTransaction", createTransaction)
	http.HandleFunc("/UpdateTransaction", updateTransaction)
	http.HandleFunc("/DeleteTransaction", deleteTransaction)
	// Management Operations
	//http.HandleFunc("/GetPartners", getPartners)
	http.HandleFunc("/GetTransactions", getTransactions)
	//http.HandleFunc("/GetEnumerations", getEnumerations)
	//http.HandleFunc("/GetCurrencies", getCurrencies)
	//http.HandleFunc("/GetTRMs", getTRMs)
}

/*
Descripcion : Funcion encargada de realizar el marshal de la respuesta en formato JSon
	de los servicios.
*/
func marshalJson(pAccept string, pResourceRs interface{}) []byte {
	var value []byte
	if pAccept == "application/json" || strings.Contains(pAccept, "application/json") {
		var err error
		if pResourceRs != nil {
			value, err = json.Marshal(pResourceRs)
		}
		if err != nil {
			log.Debugf("Error Marshalling json: %v", err)
		}
	}
	return value
}

/*
Descripcion : Funcion encargada de crear un recurso de acuerdo a la peticion de entrada.

Parametros :
      pResponse http.ResponseWriter :  contiene la respuesta que se enviara al usuario
	  pRequest *http.Request :         Contiene la peticion del usuario
*/
//func createPartner(pResponse http.ResponseWriter, pRequest *http.Request) {

//	startTime := time.Now()
//	defer panics.CatchPanic("CreatePartner")

//	message := new(domain.CreatePartnerRQ)
//	accept := pRequest.Header.Get("Accept")

//	var err error
//	if accept == "application/json" || strings.Contains(accept, "application/json") {
//		err = json.NewDecoder(pRequest.Body).Decode(&message)
//		if err != nil {
//			log.Error("Unmarshal error", err)
//		}
//	}
//	log.Info("Process Create Partner", message)
//	response := controller.ProcessCreatePartner(message)

//	// Se asigna tiempo de respuesta de todo el proceso.
//	if response != nil && response.Header != nil {
//		response.GetHeader().ResponseTime = util.Concatenate(response.GetHeader().ResponseTime)
//	}

//	value := marshalJson(accept, response)
//	pResponse.Header().Add("Content-Type", "application/json")
//	pResponse.Write(value)

//	processTime := time.Now().Sub(startTime)
//	log.Info("Process Time:", processTime.String())
//}

/*
Descripcion : Funcion encargada de actualizar un recurso de acuerdo a la peticion de entrada.

Parametros :
      pResponse http.ResponseWriter :  contiene la respuesta que se enviara al usuario
	  pRequest *http.Request :         Contiene la peticion del usuario
*/
//func updatePartner(pResponse http.ResponseWriter, pRequest *http.Request) {

//	startTime := time.Now()

//	defer panics.CatchPanic("UpdatePartner")

//	message := new(domain.UpdatePartnerRQ)
//	accept := pRequest.Header.Get("Accept")

//	var err error
//	if accept == "application/json" || strings.Contains(accept, "application/json") {
//		err = json.NewDecoder(pRequest.Body).Decode(&message)
//	}

//	if err != nil {
//		log.Error("Ha ocurrido un error al realizar el Unmarshal", err)
//	}

//	log.Info("Process Update Partner", message)
//	response := controller.ProcessUpdatePartner(message)

//	// Se asigna tiempo de respuesta de todo el proceso.
//	if response != nil && response.Header != nil {
//		response.GetHeader().ResponseTime = util.Concatenate(response.GetHeader().ResponseTime)
//	}

//	value := marshalJson(accept, response)
//	pResponse.Header().Add("Content-Type", "application/json")
//	pResponse.Write(value)

//	processTime := time.Now().Sub(startTime)
//	log.Info("Process Time:", processTime.String())

//	// Send ProcessTime for updating service metrics
//	go func(pResponse *domain.UpdatePartnerRS) {
//		if pResponse != nil {
//			//TODO Insertar codigo aqui
//		}
//	}(response)
//}

///*
//Descripcion : Funcion encargada de eliminar un recurso de acuerdo a la peticion de entrada.

//Parametros :
//      pResponse http.ResponseWriter :  contiene la respuesta que se enviara al usuario
//	  pRequest *http.Request :         Contiene la peticion del usuario
//*/
//func deletePartner(pResponse http.ResponseWriter, pRequest *http.Request) {

//	startTime := time.Now()

//	defer panics.CatchPanic("DeleteResource")

//	message := new(domain.DeletePartnerRQ)
//	accept := pRequest.Header.Get("Accept")

//	var err error
//	if accept == "application/json" || strings.Contains(accept, "application/json") {
//		err = json.NewDecoder(pRequest.Body).Decode(&message)
//		if err != nil {
//			log.Error("Ha ocurrido un error al realizar el Unmarshal", err)
//		}
//	}

//	log.Info("Process Delete Partner", message)
//	response := controller.ProcessDeletePartner(message)

//	// Se asigna tiempo de respuesta de todo el proceso.
//	if response != nil && response.Header != nil {
//		response.GetHeader().ResponseTime = util.Concatenate(response.GetHeader().ResponseTime)
//	}

//	value := marshalJson(accept, response)
//	pResponse.Header().Add("Content-Type", "application/json")
//	pResponse.Write(value)

//	processTime := time.Now().Sub(startTime)
//	log.Info("Process Time:", processTime.String())
//}

///*
//Description : Function to get a resources according to filters input request.

//Params :
//      pResponse http.ResponseWriter :  Contain the response that will be sent to the user
//	  pRequest *http.Request :         Contain the user's request
//*/
//func getPartners(pResponse http.ResponseWriter, pRequest *http.Request) {

//	startTime := time.Now()
//	defer panics.CatchPanic("GetPartners")

//	message := new(domain.GetPartnersRQ)
//	accept := pRequest.Header.Get("Accept")

//	var err error
//	if accept == "application/json" || strings.Contains(accept, "application/json") {
//		err = json.NewDecoder(pRequest.Body).Decode(&message)
//		if err != nil {
//			log.Error("Error in Unmarshal process", err)
//		}
//	}

//	log.Info("Process Get Partners", message)

//	response := controller.ProcessGetPartners(message)

//	// Set response time to all process.
//	if response != nil && response.Header != nil {
//		response.GetHeader().ResponseTime = util.Concatenate(response.GetHeader().ResponseTime)
//	}

//	value := marshalJson(accept, response)
//	pResponse.Header().Add("Content-Type", "application/json")
//	pResponse.Write(value)

//	processTime := time.Now().Sub(startTime)
//	log.Info("Process Time:", processTime.String())
//}

//---------------------------------------------------------------------------------------------------
/*
Descripcion : Funcion encargada de crear un Transaction de acuerdo a la peticion de entrada.

Parametros :
      pResponse http.ResponseWriter :  contiene la respuesta que se enviara al usuario
	  pRequest *http.Request :         Contiene la peticion del usuario
*/
func createTransaction(pResponse http.ResponseWriter, pRequest *http.Request) {

	startTime := time.Now()

	defer panics.CatchPanic("CreateTransaction")

	message := new(domain.CreateTransactionRQ)
	accept := pRequest.Header.Get("Accept")

	var err error
	if accept == "application/json" || strings.Contains(accept, "application/json") {
		err = json.NewDecoder(pRequest.Body).Decode(&message)
		if err != nil {
			log.Error("Ha ocurrido un error al realizar el Unmarshal", err)
		}
	}

	log.Info("Process Create Transaction", message)

	response := controller.ProcessCreateTransaction(message)

	// Se asigna tiempo de respuesta de todo el proceso.
	if response != nil && response.Header != nil {
		response.GetHeader().ResponseTime = util.Concatenate(response.GetHeader().ResponseTime)
	}

	value := marshalJson(accept, response)
	pResponse.Header().Add("Content-Type", "application/json")
	pResponse.Write(value)

	processTime := time.Now().Sub(startTime)
	log.Info("Process Time:", processTime.String())
}

/*
Descripcion : Funcion encargada de actualizar un proyecto de acuerdo a la peticion de entrada.

Parametros :
      pResponse http.ResponseWriter :  contiene la respuesta que se enviara al usuario
	  pRequest *http.Request :         Contiene la peticion del usuario
*/
func updateTransaction(pResponse http.ResponseWriter, pRequest *http.Request) {

	startTime := time.Now()

	defer panics.CatchPanic("UpdateTransaction")

	message := new(domain.UpdateTransactionRQ)
	accept := pRequest.Header.Get("Accept")

	var err error
	if accept == "application/json" || strings.Contains(accept, "application/json") {
		err = json.NewDecoder(pRequest.Body).Decode(&message)
		if err != nil {
			log.Error("Ha ocurrido un error al realizar el Unmarshal", err)
		}
	}
	log.Info("Process Update Transaction", message)
	response := controller.ProcessUpdateTransaction(message)

	// Se asigna tiempo de respuesta de todo el proceso.
	if response != nil && response.Header != nil {
		response.GetHeader().ResponseTime = util.Concatenate(response.GetHeader().ResponseTime)
	}

	value := marshalJson(accept, response)
	pResponse.Header().Add("Content-Type", "application/json")
	pResponse.Write(value)

	processTime := time.Now().Sub(startTime)
	log.Info("Process Time:", processTime.String())
}

/*
Descripcion : Funcion encargada de eliminar un proyecto de acuerdo a la peticion de entrada.

Parametros :
      pResponse http.ResponseWriter :  contiene la respuesta que se enviara al usuario
	  pRequest *http.Request :         Contiene la peticion del usuario
*/
func deleteTransaction(pResponse http.ResponseWriter, pRequest *http.Request) {

	startTime := time.Now()

	defer panics.CatchPanic("DeleteTransaction")

	message := new(domain.DeleteTransactionRQ)
	accept := pRequest.Header.Get("Accept")

	var err error
	if accept == "application/json" || strings.Contains(accept, "application/json") {
		err = json.NewDecoder(pRequest.Body).Decode(&message)
		if err != nil {
			log.Error("Ha ocurrido un error al realizar el Unmarshal", err)
		}
	}

	log.Info("Process Delete Transaction", message)
	response := controller.ProcessDeleteTransaction(message)

	// Se asigna tiempo de respuesta de todo el proceso.
	if response != nil && response.Header != nil {
		response.GetHeader().ResponseTime = util.Concatenate(response.GetHeader().ResponseTime)
	}

	value := marshalJson(accept, response)
	pResponse.Header().Add("Content-Type", "application/json")
	pResponse.Write(value)

	processTime := time.Now().Sub(startTime)
	log.Info("Process Time:", processTime.String())
}

/*
Description : Function to get a projects according to filters input request.

Params :
      pResponse http.ResponseWriter :  Contain the response that will be sent to the user
	  pRequest *http.Request :         Contain the user's request
*/
func getTransactions(pResponse http.ResponseWriter, pRequest *http.Request) {

	startTime := time.Now()

	defer panics.CatchPanic("GetTransactions")

	message := new(domain.GetTransactionsRQ)
	accept := pRequest.Header.Get("Accept")

	var err error
	if accept == "application/json" || strings.Contains(accept, "application/json") {
		err = json.NewDecoder(pRequest.Body).Decode(&message)
		if err != nil {
			log.Error("Error in Unmarshal process", err)
		}
	}

	log.Info("Process Get Transactions", message)

	response := controller.ProcessGetTransactions(message)

	// Set response time to all process.
	if response != nil && response.Header != nil {
		response.GetHeader().ResponseTime = util.Concatenate(response.GetHeader().ResponseTime)
	}

	value := marshalJson(accept, response)
	pResponse.Header().Add("Content-Type", "application/json")
	pResponse.Write(value)

	processTime := time.Now().Sub(startTime)
	log.Info("Process Time:", processTime.String())
}

//------------------------------------------------------------------------------------------

//func getEnumerations(pResponse http.ResponseWriter, pRequest *http.Request) {
//	startTime := time.Now()
//	defer panics.CatchPanic("GetEnumerations")

//	message := new(domain.EnumerationRQ)
//	accept := pRequest.Header.Get("Accept")

//	var err error
//	if accept == "application/json" || strings.Contains(accept, "application/json") {
//		err = json.NewDecoder(pRequest.Body).Decode(&message)
//		if err != nil {
//			log.Error("Error in Unmarshal process", err)
//		}
//	}

//	log.Info("Process Get Enumerations", message)
//	response := controller.ProcessGetEnumerations(message)

//	// Set response time to all process.
//	if response != nil && response.Header != nil {
//		response.GetHeader().ResponseTime = util.Concatenate(response.GetHeader().ResponseTime)
//	}

//	value := marshalJson(accept, response)
//	pResponse.Header().Add("Content-Type", "application/json")
//	pResponse.Write(value)

//	processTime := time.Now().Sub(startTime)
//	log.Info("Process Time:", processTime.String())

//}

//func getCurrencies(pResponse http.ResponseWriter, pRequest *http.Request) {
//	startTime := time.Now()
//	defer panics.CatchPanic("GetCurrencies")

//	message := new(domain.CurrencyRQ)
//	accept := pRequest.Header.Get("Accept")

//	var err error
//	if accept == "application/json" || strings.Contains(accept, "application/json") {
//		err = json.NewDecoder(pRequest.Body).Decode(&message)
//		if err != nil {
//			log.Error("Error in Unmarshal process", err)
//		}
//	}

//	log.Info("Process Get Currencies", message)
//	response := controller.ProcessGetCurrencies(message)

//	// Set response time to all process.
//	if response != nil && response.Header != nil {
//		response.GetHeader().ResponseTime = util.Concatenate(response.GetHeader().ResponseTime)
//	}

//	value := marshalJson(accept, response)
//	pResponse.Header().Add("Content-Type", "application/json")
//	pResponse.Write(value)

//	processTime := time.Now().Sub(startTime)
//	log.Info("Process Time:", processTime.String())

//}

//func getTRMs(pResponse http.ResponseWriter, pRequest *http.Request) {
//	startTime := time.Now()
//	defer panics.CatchPanic("GetTRMs")

//	message := new(domain.GetTRMRQ)
//	accept := pRequest.Header.Get("Accept")

//	var err error
//	if accept == "application/json" || strings.Contains(accept, "application/json") {
//		err = json.NewDecoder(pRequest.Body).Decode(&message)
//		if err != nil {
//			log.Error("Error in Unmarshal process", err)
//		}
//	}

//	log.Info("Process Get TRMs", message)
//	response := controller.ProcessGetTRMs(message)

//	// Set response time to all process.
//	if response != nil && response.Header != nil {
//		response.GetHeader().ResponseTime = util.Concatenate(response.GetHeader().ResponseTime)
//	}

//	value := marshalJson(accept, response)
//	pResponse.Header().Add("Content-Type", "application/json")
//	pResponse.Write(value)

//	processTime := time.Now().Sub(startTime)
//	log.Info("Process Time:", processTime.String())

//}

///*
//Description : Function to update a TRM according to input request.

//Params :
//      pResponse http.ResponseWriter :  Contain the response that will be sent to the user
//	  pRequest *http.Request :         Contain the user's request
//*/
//func updateTRM(pResponse http.ResponseWriter, pRequest *http.Request) {

//	startTime := time.Now()

//	defer panics.CatchPanic("UpdateTRM")

//	message := new(domain.UpdateTRMRQ)
//	accept := pRequest.Header.Get("Accept")

//	var err error
//	if accept == "application/json" || strings.Contains(accept, "application/json") {
//		err = json.NewDecoder(pRequest.Body).Decode(&message)
//		if err != nil {
//			log.Error("Error in Unmarshal process", err)
//		}
//	}
//	log.Info("Process Update TRM", message)
//	response := controller.ProcessUpdateTRM(message)

//	// Set response time to all process.
//	if response != nil && response.Header != nil {
//		response.GetHeader().ResponseTime = util.Concatenate(response.GetHeader().ResponseTime)
//	}

//	value := marshalJson(accept, response)
//	pResponse.Header().Add("Content-Type", "application/json")
//	pResponse.Write(value)

//	processTime := time.Now().Sub(startTime)
//	log.Info("Process Time:", processTime.String())
//}
