package main

import (
	"fmt"
	"net/url"
	"strings"
)

type Eurovent struct {
	client Client
}

func (e *Eurovent) GetTotalCount(program string, productType string, brands []string) (int, error) {
	values := url.Values{}
	values.Set("program", program)
	values.Set("product_type", productType)
	for _, brand := range brands {
		values.Add("brand[]", strings.ToUpper(brand))
	}
	var t TotalCountResponse
	err := e.client.Get("https://www.eurovent-certification.com/en/totalcount", values, &t)
	return t.TotalCount, err
}

func (e *Eurovent) GetData(program string, productType string, brands []string, limit int) (AdvancedSearchResponse, error) {
	values := url.Values{}
	values.Set("limit", fmt.Sprintf("%d", limit))
	values.Set("program", program)
	values.Set("product_type", productType)
	for _, brand := range brands {
		values.Add("brand[]", strings.ToUpper(brand))
	}
	var r AdvancedSearchResponse
	err := e.client.Post("https://www.eurovent-certification.com/en/advancedsearch/ajax", values, &r)
	return r, err
}

func (e *Eurovent) DataToCsv(program string, productType string, brands []string, data []int) (string, error) {

	params := url.Values{}
	params.Set("program", program)
	params.Set("product_type", productType)

	serviceUrl := "https://www.eurovent-certification.com/en/catalog/program/certificate/participant/model/csv?" + params.Encode()

	payload := url.Values{}
	payload.Set("head_result", "false")
	payload.Set("range_brand_col", "true")
	for _, brand := range brands {
		payload.Add("brand", strings.ToUpper(brand))
	}
	models := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(data)), ","), "[]")
	payload.Set("model", models)

	res, err := e.client.PostRaw(serviceUrl, payload)
	return string(res), err
}

type TotalCountResponse struct {
	TotalCount int `json:"total_count"`
}

type AdvancedSearchResponse struct {
	Draw            interface{} `json:"draw"`
	Total           int         `json:"total"`
	RecordsFiltered int         `json:"recordsFiltered"`
	Rows            []struct {
		ID              int    `json:"id"`
		Ppr             string `json:"ppr"`
		ModelName       string `json:"model_name"`
		Range           string `json:"range"`
		Brand           string `json:"brand"`
		Country         string `json:"country"`
		City            string `json:"city"`
		SoftwareName    string `json:"software_name"`
		SoftwareVersion string `json:"software_version"`
		Champ23         string `json:"champ_23"`
		Champ24         string `json:"champ_24"`
		Champ25         string `json:"champ_25"`
		Champ26         string `json:"champ_26"`
		Champ27         string `json:"champ_27"`
		Champ28         string `json:"champ_28"`
		Champ29         string `json:"champ_29"`
		Champ30         string `json:"champ_30"`
		Champ34         string `json:"champ_34"`
		Champ35         string `json:"champ_35"`
		Champ36         string `json:"champ_36"`
		Champ38         string `json:"champ_38"`
		Champ46         string `json:"champ_46"`
		Champ48         string `json:"champ_48"`
		Champ50         string `json:"champ_50"`
		Champ51         string `json:"champ_51"`
		Champ52         string `json:"champ_52"`
		Champ59         string `json:"champ_59"`
		Champ61         string `json:"champ_61"`
		Champ63         string `json:"champ_63"`
		Champ65         string `json:"champ_65"`
		Champ68         string `json:"champ_68"`
		Champ69         string `json:"champ_69"`
		Champ72         string `json:"champ_72"`
		Champ73         string `json:"champ_73"`
		Champ75         string `json:"champ_75"`
		Champ77         string `json:"champ_77"`
		Champ79         string `json:"champ_79"`
		Champ81         string `json:"champ_81"`
		Champ83         string `json:"champ_83"`
		Champ85         string `json:"champ_85"`
		Champ87         string `json:"champ_87"`
		Champ89         string `json:"champ_89"`
		Champ90         string `json:"champ_90"`
		Champ91         string `json:"champ_91"`
		Champ92         string `json:"champ_92"`
		Champ93         string `json:"champ_93"`
		Champ94         string `json:"champ_94"`
		Champ95         string `json:"champ_95"`
		Champ96         string `json:"champ_96"`
		Champ98         string `json:"champ_98"`
		Champ99         string `json:"champ_99"`
		Champ100        string `json:"champ_100"`
		Champ101        string `json:"champ_101"`
		Champ102        string `json:"champ_102"`
		Champ103        string `json:"champ_103"`
		Champ104        string `json:"champ_104"`
		Champ105        string `json:"champ_105"`
		Champ106        string `json:"champ_106"`
		Champ108        string `json:"champ_108"`
		Champ109        string `json:"champ_109"`
		Champ110        string `json:"champ_110"`
		Champ111        string `json:"champ_111"`
		Champ113        string `json:"champ_113"`
		Champ114        string `json:"champ_114"`
		Champ138        string `json:"champ_138"`
		Champ139        string `json:"champ_139"`
		Champ140        string `json:"champ_140"`
		Champ141        string `json:"champ_141"`
		Champ142        string `json:"champ_142"`
		Champ143        string `json:"champ_143"`
		Champ144        string `json:"champ_144"`
		Champ145        string `json:"champ_145"`
		Champ146        string `json:"champ_146"`
	} `json:"rows"`
}