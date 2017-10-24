package util

import (
	"bytes"
	"errors"
	"math"
	"strconv"
	"strings"
	"time"

	"FenixPlatform/com.cambiosfenix.ftm.service/domain"
	"FenixPlatform/com.cambiosfenix.ftm.service/log"
)

//	"es": "Español",
//	"en": "Inglés",
//	"ca": "Catalán",
//	"de": "Alemán",
//	"it": "Italiano",
//	"nl": "Holandés",
//	"fr": "Francés",
//	"pt": "Portugués",
//	"ru": "Ruso",
//	"sv": "Sueco",
//	"zh": "Chino",
//	"ja": "Japonés",
//	"da": "Danés",
//	"ar": "Árabe",
//	"id": "Indonesio",
//	"ko": "Coreano",
//	"ms": "Malayo",
//	"no": "Noruego",
//	"tl": "Tagalo",
//	"th": "Tailandés",
//	"tr": "Turco",
//	"vi": "Vietnamita",

const (
	LANGUAGECODE = "es"
	LANGUAGES    = "es,en,ca,de,it,nl,fr,pt,ru,sv,zh,ja,da,ar,id,ko,ms,no,tl,th,tr,vi"
)

func StringToBool(pString string) bool {
	if pString == "true" {
		return true
	}
	return false
}

// Concatenates the sent string in a single string
func Concatenate(pStrings ...string) string {
	listStrings := bytes.Buffer{}
	for _, str := range pStrings {
		listStrings.WriteString(str)
	}
	return listStrings.String()
}

// Convierte un entero a string
func IntToString(pNumero int) string {
	return strconv.Itoa(pNumero)
}

// Convierte un entero a string
func Int64ToString(pNumero int64) string {
	return strconv.FormatInt(pNumero, 10)
}

/**
* Funcion engargada de obtener la descripcion necesaria de acuerdo al idioma.
 */
func LenguageManagement(pListLanguages, pLanguage string) string {
	if pListLanguages != "" && pLanguage != "" {
		var descripcion string
		// Separamos la lista para obtener cada pareja de lenguaje y valor asociado
		lenguajes := strings.Split(pListLanguages, "¬")
		for _, leng := range lenguajes {
			// Se descompone cada pareja, primera posicion es el lenguaje
			// La segunda posicion corresponde al valor asociado.
			lengDesc := strings.Split(leng, "|")
			if pLanguage == lengDesc[0] {
				// Retornamos el nombre que nos interesa segun el idioma de la peticion
				descripcion = lengDesc[1]
			}
		}
		if descripcion == "" {
			if pLanguage != "es" {
				for _, leng := range lenguajes {
					lengDesc := strings.Split(leng, "|")
					if lengDesc[0] == "es" {
						// Agregamos la descripcion en español
						descripcion = lengDesc[1]
					}
				}
			}
			if descripcion == "" {
				if pLanguage != "en" {
					for _, leng := range lenguajes {
						lengDesc := strings.Split(leng, "|")
						if lengDesc[0] == "en" {
							// Agregamos la descripcion en ingles
							descripcion = lengDesc[1]
						}
					}
				}
			}
		}
		return descripcion
	}
	log.Debug("No se encontro el idioma %v en el listado %v", pLanguage, pListLanguages)
	return ""
}

// Funcion encargada de convertir un string a entero, en caso de que el valor no sea un numero informa del error
// Sin embargo retorna cero
func StringToInt(pString string) int {
	if pString != "" {
		number, err := strconv.Atoi(pString)
		if err != nil {
			log.Errorf("Ocurrio un error al convertir el valor %v a entero, se retorna cero (0)", pString)
			return 0
		}
		return number
	}
	return 0
}

// GetNumAdultNinBebByOcup Obtiene el número de adultos, niños y bebes; a partir de un código de ocupación.
func GetNumAdultNinBebByOcup(pCodigo string) (a, n, b int, err error) {
	ocupacionSplit := strings.Split(pCodigo, "-")
	if len(ocupacionSplit) == 3 {
		a, err = strconv.Atoi(ocupacionSplit[0])
		if err != nil {
			log.Error("Ocurrio un error obteniendo la cantidad de adultos de la ocupación:", pCodigo)
		}
		n, err = strconv.Atoi(ocupacionSplit[1])
		if err != nil {
			log.Error("Ocurrio un error obteniendo la cantidad de ninos de la ocupación:", pCodigo)
		}
		b, err = strconv.Atoi(ocupacionSplit[2])
		if err != nil {
			log.Error("Ocurrio un error obteniendo la cantidad de bebes de la ocupación:", pCodigo)
		}
	} else {
		log.Error("El código de la ocupación no tiene el formato A-N-B, sino", pCodigo)
		err = errors.New("No se obtiene la cantidad de valores esperados")
	}
	return a, n, b, err
}

// Reemplazamos los saltos de linea por BR
func SaltosToBr(pCadena string) string {
	return strings.Replace(pCadena, "\n", "<br/>", -1)
}

/**
 * Funcion encargada de transformar una lista de string en una lista de int
 */
func SliceStringToInt(pSlice []string) []int {
	if pSlice == nil || len(pSlice) == 0 {
		log.Debug("Lista nil o vacia")
		return []int{}
	}
	var resultList []int
	// Se recorre la lista de strings
	for _, ele := range pSlice {
		eleInt, err := strconv.Atoi(ele)
		if err != nil {
			log.Debug("No se puede convertir la lista de strings")
			return []int{}
		}
		resultList = append(resultList, eleInt)
	}
	return resultList
}

func RedondearFloat64(val float64, places int) (newVal float64) {
	// Valor sumado para mejorar redondeo
	var aux = 0.00000000000001
	val = val + aux
	var roundOn float64
	roundOn = 0.5
	esNegativo := false
	if val < 0 {
		esNegativo = true
		val = val * (-1)
	}
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	if esNegativo {
		newVal = newVal * (-1)
	}
	return
}

func RedondearFiltros(pValor float64, pDecimales int) float64 {
	valor := strconv.FormatFloat(pValor, 'f', pDecimales, 64)
	valorRedondeado, _ := strconv.ParseFloat(valor, 64)
	return valorRedondeado
}

// Funcion encargada de convertir un string t o f a booleano de go
func BoolPSQLToBool(pString string) bool {
	if pString == "t" {
		return true
	}
	return false
}

/**
* Function to mapping create partner request to business entity partner.
 */
func MappingCreatePartner(pRequest *domain.CreatePartnerRQ) *domain.Partner {
	partner := new(domain.Partner)
	partner.Code = pRequest.Code
	partner.Name = pRequest.Name
	partner.Address = pRequest.Address
	partner.City = pRequest.City
	partner.ZipCode = pRequest.ZipCode
	partner.State = pRequest.State
	partner.CountryCode = pRequest.CountryCode
	partner.Phone = pRequest.Phone
	partner.Email = pRequest.Email

	return partner
}

/**
* Function to mapping create transaction request to business entity transaction.
 */
func MappingCreateTransaction(pRequest *domain.CreateTransactionRQ) *domain.Transaction {
	transaction := new(domain.Transaction)
	transaction.Code = pRequest.Code
	transaction.Type = pRequest.Type
	startDate := new(string)
	startDate = &pRequest.Date
	if startDate == nil || *startDate == "" {
		log.Error("Dates undefined")
		return nil
	}
	startDateInt, err := ConvertirFechasPeticion(startDate)
	if err != nil {
		log.Error(err)
		return nil
	}
	transaction.Date = time.Unix(startDateInt, 0)
	transaction.TRM = BuildTRM(pRequest.TRMID)
	transaction.TRM.TRMValue = pRequest.TRMValue
	transaction.Partner = BuildPartner(pRequest.PartnerID)
	transaction.CurrencyFromQuantity = pRequest.CurrencyFromQuantity
	transaction.CurrencyToQuantity = pRequest.CurrencyToQuantity
	transaction.PayType = pRequest.PayType
	transaction.IncoTerm = pRequest.IncoTerm
	transaction.Casher = pRequest.Casher
	transaction.Notes = pRequest.Notes
	return transaction
}

func BuildEnumeration(pEnum int) *domain.Enumeration {

	newEnum := new(domain.Enumeration)
	newEnum.ID = pEnum
	return newEnum
}

func BuildCurrency(pCurrencyID int) *domain.Currency {

	newCurrency := new(domain.Currency)
	newCurrency.ID = pCurrencyID
	return newCurrency
}

func BuildTRM(pTRMID int) *domain.TRM {

	newCurrency := new(domain.TRM)
	newCurrency.ID = pTRMID
	return newCurrency
}

func BuildPartner(pPartner int) *domain.Partner {

	newPartner := new(domain.Partner)
	newPartner.ID = pPartner
	return newPartner
}

/**
* Function to mapping create TRM request to business entity TRM.
 */
func MappingCreateTRM(pRequest *domain.CreateTRMRQ) *domain.TRM {
	trm := new(domain.TRM)
	trm.CurrencyFromID = pRequest.CurrencyFromID
	trm.CurrencyToID = pRequest.CurrencyToID
	trm.TRMValue = pRequest.TRMValue
	trm.TRMSaleValue = pRequest.TRMSaleValue
	return trm
}

/**
* Function to mapping request to get partner in a partner entity.
 */
func MappingFiltersPartner(pRequest *domain.GetPartnersRQ) *domain.Partner {
	if pRequest != nil {
		filters := domain.Partner{}

		if pRequest.ID != 0 {
			filters.ID = pRequest.ID
		}
		if pRequest.Code != "" {
			filters.Code = pRequest.Code
		}
		if pRequest.Name != "" {
			filters.Name = pRequest.Name
		}

		return &filters
	}
	return nil
}

/**
* Function to mapping request to get Transactions in a Transactions entity.
 */
func MappingFiltersTransactions(pRequest *domain.GetTransactionsRQ) *domain.Transaction {
	if pRequest != nil {
		filters := domain.Transaction{}

		if pRequest.ID != 0 {
			filters.ID = pRequest.ID
		}
		if pRequest.Code != "" {
			filters.Code = pRequest.Code
		}
		if pRequest.StartDate != "" {
			startDate, err := time.Parse("2006-01-02", pRequest.StartDate)
			if err == nil {
				filters.Date = startDate
			}
		}
		if pRequest.Type != 0 {
			filters.Type = pRequest.Type
		}
		if pRequest.TRMID != 0 {
			filters.TRM = BuildTRM(pRequest.TRMID)
		}

		return &filters
	}
	return nil
}

func MappingCurrencyRQ(pDomain *domain.CurrencyRQ) *domain.Currency {
	currency := new(domain.Currency)
	currency.ID = pDomain.ID
	currency.Code = pDomain.Code

	return currency
}

/**
* Function to mapping Types request to business entity project.
 */
func MappingEnumeration(pRequest *domain.EnumerationRQ) *domain.Enumeration {
	enumeration := new(domain.Enumeration)
	enumeration.Value = pRequest.Value
	return enumeration
}

func BuildHeaderResponse(timeResponse time.Time) *domain.Response_Header {
	header := new(domain.Response_Header)
	header.RequestDate = time.Now().String()
	responseTime := time.Now().Sub(timeResponse)
	header.ResponseTime = responseTime.String()

	return header
}
