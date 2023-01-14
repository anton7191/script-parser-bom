package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://www.elitan.ru/order/additemorder?count=6&id=0eb52066965810daa0cb7baa5334731a&new_extra=0&pricenull=0&reReelCost=0&targetYear=0&tar%2520getYearCost=0&target_price=&time=1673369720"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

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

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
