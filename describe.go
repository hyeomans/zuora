package zuora

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//DescribeService access to describe endpoint. Don't use in production or use at your own risk
type DescribeService struct {
	config       *Config
	tokenService *TokenService
}

type object struct {
	XMLName xml.Name `xml:"object"`
	Name    string   `xml:"name"`
	Fields  fields   `xml:"fields"`
}

type fields struct {
	XMLName xml.Name `xml:"fields"`
	Fields  []field  `xml:"field"`
}

type field struct {
	XMLName    xml.Name `xml:"field"`
	Name       string   `xml:"name"`
	Custom     bool     `xml:"custom"`
	Label      string   `xml:"label"`
	Selectable bool     `xml:"selectable"`
	Createable bool     `xml:"createable"`
	Updateable bool     `xml:"updateable"`
	Filterable bool     `xml:"filterable"`
	Maxlength  string   `xml:"maxlength"`
	Required   bool     `xml:"required"`
	Type       string   `xml:"type"`
}

func newDescribeService(config *Config, tokenService *TokenService) *DescribeService {
	return &DescribeService{
		config:       config,
		tokenService: tokenService,
	}
}

//Model returns useful information about a Zuora Object
func (s *DescribeService) Model(ctx context.Context, objectName ObjecName) (string, error) {
	token, err := s.tokenService.Token(ctx)

	if err != nil {
		return "", err
	}

	url := fmt.Sprint(s.config.BaseURL, "/v1/describe/", objectName)
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", fmt.Sprint("Bearer ", token.AccessToken))
	req.Header.Add("Content-Type", "application/json")

	if ctx.Value("Zuora-Entity-Ids") != nil {
		req.Header.Add("Zuora-Entity-Ids", ctx.Value("Zuora-Entity-Ids").(string))
	}

	if ctx.Value("Zuora-Track-Id") != nil {
		req.Header.Add("Zuora-Track-Id", ctx.Value("Zuora-Track-Id").(string))
	}

	res, err := s.config.HTTPClient.Do(req)

	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	var objectAsXML object

	if err := xml.Unmarshal(body, &objectAsXML); err != nil {
		return "", err
	}

	var b strings.Builder
	fmt.Fprintf(&b, "type %v struct{\n", objectAsXML.Name)

	for _, field := range objectAsXML.Fields.Fields {
		currentType := getType(field.Type)
		if field.Required {
			fmt.Fprintf(&b, "%v\t%v\t`json:\"%v\"`\n", field.Name, currentType, field.Name)
		} else {
			fmt.Fprintf(&b, "%v\t%v\t`json:\"%v,omitempty\"`\n", field.Name, currentType, field.Name)
		}
	}

	fmt.Fprint(&b, "}\n")
	return b.String(), nil
}

//ModelNonCustom returns only original fields in Zuora object
func (s *DescribeService) ModelNonCustom(ctx context.Context, objectName ObjecName) (string, error) {
	token, err := s.tokenService.Token(ctx)

	if err != nil {
		return "", err
	}

	url := fmt.Sprint(s.config.BaseURL, "/v1/describe/", objectName)
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", fmt.Sprint("Bearer ", token.AccessToken))
	req.Header.Add("Content-Type", "application/json")

	if ctx.Value("Zuora-Entity-Ids") != nil {
		req.Header.Add("Zuora-Entity-Ids", ctx.Value("Zuora-Entity-Ids").(string))
	}

	if ctx.Value("Zuora-Track-Id") != nil {
		req.Header.Add("Zuora-Track-Id", ctx.Value("Zuora-Track-Id").(string))
	}

	res, err := s.config.HTTPClient.Do(req)

	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	var objectAsXML object

	if err := xml.Unmarshal(body, &objectAsXML); err != nil {
		return "", err
	}

	var b strings.Builder
	fmt.Fprintf(&b, "type %v struct{\n", objectAsXML.Name)

	for _, field := range objectAsXML.Fields.Fields {
		if field.Custom {
			continue
		}
		currentType := getType(field.Type)
		if field.Required {
			fmt.Fprintf(&b, "%v\t%v\t`json:\"%v\"`\n", field.Name, currentType, field.Name)
		} else {
			fmt.Fprintf(&b, "%v\t%v\t`json:\"%v,omitempty\"`\n", field.Name, currentType, field.Name)
		}
	}

	fmt.Fprint(&b, "}\n")
	return b.String(), nil
}

func getType(fieldType string) string {
	switch fieldType {
	case "boolean":
		return "bool"
	case "picklist", "text":
		return "string"
	}
	return "string"
}
