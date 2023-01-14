package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type ElitanResponseFind struct {
	Str   string `json:"str"`
	Delay string `json:"delay"`
	Mfg   string `json:"mfg"`
	Items struct {
		Ids            []string      `json:"ids"`
		Articles       []interface{} `json:"articles"`
		Aliass         []string      `json:"aliass"`
		CountValFilter []interface{} `json:"count_val_filter"`
		CountCurValue  int           `json:"count_cur_value"`
		Sort           interface{}   `json:"sort"`
		Data           []struct {
			Ntovara       string  `json:"Ntovara"`
			PartnameFull  string  `json:"partname_full"`
			Partname      string  `json:"partname"`
			Images        *string `json:"images"`
			Housing       *string `json:"housing"`
			Bignote       *string `json:"bignote"`
			Namemfg       string  `json:"namemfg"`
			Linkmfg       string  `json:"linkmfg"`
			Id            string  `json:"id"`
			Idmfg         string  `json:"idmfg"`
			Artik         string  `json:"artik"`
			ArtikOriginal string  `json:"artik_original"`
			Danger        string  `json:"danger"`
			Rayting       string  `json:"rayting"`
			TermDelay     string  `json:"term_delay"`
			ParamSearch   string  `json:"param_search"`
			Datasheets    *string `json:"datasheets"`
			MoreSearch    string  `json:"more_search"`
			Visible       int     `json:"visible"`
			CountRace     int     `json:"count_race"`
			CountStock    int     `json:"count_stock"`
			Stock         []struct {
				IdStock              string      `json:"id_stock"`
				Id                   string      `json:"id"`
				HashPartname         string      `json:"hash_partname"`
				PartnameHidden       string      `json:"partname_hidden"`
				IdSearch             string      `json:"id_search"`
				Partclear            string      `json:"partclear"`
				Rayting              string      `json:"rayting"`
				MinCount             string      `json:"min_count"`
				Ntovara              string      `json:"Ntovara"`
				Nka                  string      `json:"Nka"`
				IdtovarKA            string      `json:"idtovarKA"`
				Partname             string      `json:"partname"`
				Artik                string      `json:"artik"`
				Discount             string      `json:"Discount"`
				Note                 string      `json:"note"`
				Mfg                  string      `json:"mfg"`
				Stock                string      `json:"stock"`
				Group                string      `json:"group"`
				Normoupakovka        string      `json:"normoupakovka"`
				TermDelay            string      `json:"term_delay"`
				Dostup               string      `json:"dostup"`
				Abs                  string      `json:"abs"`
				Race                 string      `json:"race"`
				Danger               string      `json:"danger"`
				Valuta               string      `json:"valuta"`
				Extra                string      `json:"extra"`
				Exportcontrol        string      `json:"exportcontrol"`
				Manufpartnumber      string      `json:"manufpartnumber"`
				Norm                 string      `json:"norm"`
				Alias1               string      `json:"alias1"`
				Alias2               string      `json:"alias2"`
				Alias3               string      `json:"alias3"`
				CountInPack          string      `json:"countInPack"`
				RiskLogistikEAC      string      `json:"risk_logistik_EAC"`
				RiskReserv           string      `json:"risk_reserv"`
				RiskLogistikDelivery string      `json:"risk_logistik_delivery"`
				Prices               string      `json:"prices"`
				Cache                string      `json:"cache"`
				ReReelCost           string      `json:"reReelCost"`
				TargetYear           string      `json:"targetYear"`
				TargetYearCost       string      `json:"targetYearCost"`
				Reklama              string      `json:"reklama"`
				CodeNomForSpec       string      `json:"codeNomForSpec"`
				Blitz                string      `json:"blitz"`
				Weight               string      `json:"weight"`
				DataCreate           string      `json:"DataCreate"`
				Namemfg              string      `json:"namemfg"`
				Linkmfg              string      `json:"linkmfg"`
				Idmfg                string      `json:"idmfg"`
				ParamSearch          string      `json:"param_search"`
				Images               *string     `json:"images"`
				Big                  *string     `json:"big"`
				Datasheets           *string     `json:"datasheets"`
				Big2                 *string     `json:"big2"`
				Basic                interface{} `json:"basic"`
				PartnameFull         string      `json:"partname_full"`
				MinRayting           string      `json:"min_rayting"`
				ArtikOriginal        string      `json:"artik_original"`
				PricesCache          string      `json:"prices_cache"`
				StockCache           string      `json:"stock_cache"`
				Housing              *string     `json:"housing"`
				MoreSearch           string      `json:"more_search"`
				DataCreateCache      string      `json:"DataCreateCache"`
				Price                []struct {
					Price string `json:"price"`
					Count string `json:"count"`
					Total string `json:"total"`
				} `json:"price"`
				Typebasket      string `json:"typebasket"`
				Mincount        string `json:"mincount"`
				TermDelayString string `json:"term_delay_string"`
				Pack            string `json:"pack"`
				ShowPack        int    `json:"show_pack"`
				PackNoteShow    int    `json:"pack_note_show"`
				PriceFull       []struct {
					Price string `json:"price"`
					Count string `json:"count"`
					Total string `json:"total"`
				} `json:"price_full,omitempty"`
				PackNoteShowPercent int `json:"pack_note_show_percent,omitempty"`
			} `json:"stock"`
			Sample          string `json:"sample"`
			Index           int    `json:"index"`
			Normupakovka    string `json:"normupakovka"`
			YearIssue       string `json:"year_issue"`
			Small           string `json:"small,omitempty"`
			DatasheetsArray struct {
				Field1 []struct {
					Value    string `json:"value"`
					Type     string `json:"type"`
					Image    string `json:"image"`
					TypeText string `json:"type_text"`
					First    int    `json:"first"`
				} `json:"1"`
			} `json:"datasheets_array,omitempty"`
			Alias []string `json:"alias"`
		} `json:"data"`
		Time        interface{} `json:"time"`
		CountCurVal []struct {
			Value  int `json:"value"`
			Select int `json:"select"`
		} `json:"count_cur_val"`
	} `json:"items"`
	Sort              string      `json:"sort"`
	Search            int         `json:"search"`
	User              int         `json:"user"`
	Showhelp          int         `json:"showhelp"`
	TimeWork          interface{} `json:"time_work"`
	MinTargetSum      int         `json:"min_target_sum"`
	MinTargetStockSum interface{} `json:"min_target_stock_sum"`
	Url               string      `json:"url"`
	NtovaraArray      interface{} `json:"Ntovara_array"`
	Group             interface{} `json:"group"`
	Bignotesearch     string      `json:"bignotesearch"`
	FirstRow          int         `json:"first_row"`
	CountSearch       int         `json:"count_search"`
	CountPath         int         `json:"count_path"`
	ShowProject       int         `json:"show_project"`
	ShowSale          int         `json:"show_sale"`
	ShowBom           int         `json:"show_bom"`
	ShowBlitz         int         `json:"show_blitz"`
	SumBalance        string      `json:"sum_balance"`
	SumNewExtra       int         `json:"sum_new_extra"`
	MaxBasketSum      int         `json:"max_basket_sum"`
}

func main() {

	var v = make(url.Values)
	v.Set("find", "LME49726MY/NOPB")
	v.Set("t", strconv.FormatInt(time.Now().Unix(), 10))
	v.Set("delay", "-1")
	v.Set("mfg", "all")
	v.Set("seenform", "y")

	var u = url.URL{
		Scheme:   "https",
		Host:     "www.elitan.ru",
		Path:     "price/index_json.php",
		RawQuery: v.Encode(),
	}
	log.Println("URL/net: ")
	log.Println(u.String())

	resp, err := http.Get(u.String())
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	elitanResponse := ElitanResponseFind{}
	err = json.Unmarshal(body, &elitanResponse)
	if err != nil {
		log.Fatal("Unmarshal failed", err)
	}
	fmt.Println("Done", elitanResponse.Items.Ids)
	fmt.Println("Done", len(elitanResponse.Items.Data))
	var minPrice int64

	for i, data := range elitanResponse.Items.Data {
		//fmt.Println("Len Data: ", len(data.Stock))
		for i2, stock := range data.Stock {
			fmt.Println("Number", i, "Stock", i2, "ID: ", stock.Id)
		}
	}

	var addRequest = make(url.Values)
	addRequest.Set("id", elitanResponse.Items.Data[0].Stock[0].Id)
	addRequest.Set("count", "6")
	addRequest.Set("target_price", "")
	addRequest.Set("pricenull", "0")
	addRequest.Set("reReelCost", "0")
	addRequest.Set("targetYear", "0")
	addRequest.Set("targetYearCost", "0")
	addRequest.Set("new_extra", "0")
	addRequest.Set("time", strconv.FormatInt(time.Now().Unix(), 10))

	var addUrl = url.URL{
		Scheme:   "https",
		Host:     "www.elitan.ru",
		Path:     "order/additemorder",
		RawQuery: addRequest.Encode(),
	}
	log.Println("URL/net: ")
	log.Println(addUrl.String())

	client := &http.Client{}
	req, err := http.NewRequest("GET", addUrl.String(), nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("Accept-Language", "ru,en;q=0.9,zh;q=0.8,de;q=0.7,la;q=0.6")
	req.Header.Add("Cache-Control", "max-age=0")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cookie", "city=7800000000000000000000000; region=7800000000000000000000000; __utma=84796621.1324461965.1564412988.1624354689.1627636495.28; adres_registr=192007, Санкт-Петербург г, Лиговский пр-кт, дом № 166, лит.А, пом.13Н; inn=7802719547; user_adres=34d295e136ed8347b738a24e8905fce9; deliveryglobal=0; sticker=0; OrderGarant=0; paperaccount=0; UPD=2BM; UPD98948=2BM; UPDcheck=1; lastname=Дойников; firstname=Антон; middlename=Евгеньевич; tel=+7, 985-787-65-61; paperaccountcheck=1; OrderGarantcheck=1; post=инженер; show_popup_basket=1; OrderDesigned=STROYKA; exp=Для изготовления стендов проверки электронного оборудования; TypeAccount=СО; OrderDateValid=0; OrderDateValidcheck=1; stickercheck=1; OrderDesignedcheck=1; delivery=16778000000000; order=2fftvi6s3g5d4e83oqj5empdj3; PHPSESSID=5jh61q8f2vm0du7gd82rhrt4ar; type_sort=sort_price; cur_date=11; count_cur_val=1; current_basket=d51d3e3c828fdb745faf458b73ed389b; mail_hash=5d691068de4771e05a2193142d93dfbb; email_aut=aantonov@track-me.pro; show_bom=1")
	req.Header.Add("Sec-Fetch-Dest", "document")
	req.Header.Add("Sec-Fetch-Mode", "navigate")
	req.Header.Add("Sec-Fetch-Site", "none")
	req.Header.Add("Sec-Fetch-User", "?1")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 YaBrowser/23.1.0.2539 (beta) Yowser/2.5 Safari/537.36")
	req.Header.Add("sec-ch-ua", "\"Not?A_Brand\";v=\"8\", \"Chromium\";v=\"108\", \"Yandex\";v=\"23\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")

	respAdd, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer respAdd.Body.Close()

	//bodyAdd, err := ioutil.ReadAll(respAdd.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Println("BodyAdd: ", string(bodyAdd))
}
