package zuora

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type describeService struct {
	http               Doer
	authHeaderProvider AuthHeaderProvider
	baseURL            string
}

type xmlObject struct {
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

func newDescribeService(http Doer, authHeaderProvider AuthHeaderProvider, baseURL string) *describeService {
	return &describeService{
		http:               http,
		authHeaderProvider: authHeaderProvider,
		baseURL:            baseURL,
	}
}

func (t *describeService) Model(ctx context.Context, objectName ObjecName) (string, error) {
	authHeader, err := t.authHeaderProvider.AuthHeaders(ctx)

	if err != nil {
		return "", responseError{isTemporary: false, message: fmt.Sprintf("error while trying to set auth headers: %v", err)}
	}

	url := fmt.Sprintf("%v/v1/describe/%v", t.baseURL, objectName)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return "", responseError{isTemporary: false, message: fmt.Sprintf("error while trying to create an HTTP request: %v", err)}
	}

	req.Header.Add("Authorization", authHeader)
	req.Header.Add("Content-Type", "application/json")

	if ctx.Value(ContextKeyZuoraEntityIds) != nil {
		req.Header.Add("Zuora-Entity-Ids", ctx.Value(ContextKeyZuoraEntityIds).(string))
	}

	if ctx.Value(ContextKeyZuoraTrackID) != nil {
		req.Header.Add("Zuora-Track-Id", ctx.Value(ContextKeyZuoraTrackID).(string))
	}

	if ctx.Value(ContextKeyZuoraVersion) != nil {
		req.Header.Add("zuora-version", ctx.Value(ContextKeyZuoraVersion).(string))
	}

	res, err := t.http.Do(req.WithContext(ctx))
	defer res.Body.Close()

	if err != nil {
		return "", responseError{isTemporary: false, message: fmt.Sprintf("error while trying to make request: %v", err)}
	}

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode < 200 || res.StatusCode > 299 {
		var isTemporary bool
		if http.StatusRequestTimeout == res.StatusCode ||
			http.StatusTooManyRequests == res.StatusCode ||
			http.StatusInternalServerError == res.StatusCode ||
			http.StatusServiceUnavailable == res.StatusCode {
			isTemporary = true
		}

		if err != nil {
			return "", responseError{isTemporary: isTemporary, message: fmt.Sprintf("error while trying to read body response into memory. Response Code: %v - Error: %v", res.StatusCode, err)}
		}

		return "", responseError{isTemporary: isTemporary, message: fmt.Sprintf("got an invalid http status. Response Code: %v - Body: %v", res.StatusCode, string(body))}
	}

	var objectAsXML xmlObject

	if err := xml.Unmarshal(body, &objectAsXML); err != nil {
		return "", responseError{isTemporary: false, message: fmt.Sprintf("error while Unmarshal json response. Error: %v. JSON: %v", err, string(body))}
	}

	var b strings.Builder
	fmt.Fprintf(&b, "type %v struct{\n", objectAsXML.Name)

	for _, field := range objectAsXML.Fields.Fields {
		fmt.Println("Field types", field.Name, field.Type)
		currentType := getType(field.Required, field.Type)
		if field.Required {
			fmt.Fprintf(&b, "%v\t%v\t`json:\"%v\"`\n", field.Name, currentType, field.Name)
		} else {
			fmt.Fprintf(&b, "%v\t%v\t`json:\"%v,omitempty\"`\n", field.Name, currentType, field.Name)
		}
	}

	fmt.Fprint(&b, "}\n")
	return b.String(), nil
}

func getType(fieldRequired bool, fieldType string) string {
	var goType string
	switch fieldType {
	case "boolean":
		goType = "bool"
	case "picklist", "text":
		goType = "string"
	case "integer":
		goType = "int"
	case "decimal":
		goType = "float64"
	default:
		goType = "string"
	}

	if fieldRequired {
		return goType
	}

	return fmt.Sprintf("*%v", goType)
}
