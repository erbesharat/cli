package helper

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"text/template"

	"github.com/resumic/schema/docs/schema"
	"github.com/tidwall/gjson"
)

//InitResume initializes the resume.json file
func InitResume() {
	var conf, doc, tmpl string
	flag.StringVar(&conf, "conf", "../schema/schema.json", "Schema json file")
	flag.StringVar(&tmpl, "tmpl", "resume.tmpl", "Template file")
	flag.StringVar(&doc, "doc", "resume.json", "Output json file")
	config, err := ioutil.ReadFile(conf)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile(tmpl, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString("{")
	value := gjson.GetBytes(config, "properties")
	j := 0
	value.ForEach(func(key, value gjson.Result) bool {
		if j != 0 {
			f.WriteString("\n ,")
		}
		j++
		_, err = f.WriteString("\"" + key.String() + "\":") //write name of sections
		if err != nil {
			log.Fatal(err)
		}
		if (isArray(getType(value.String()))) == true {
			_, err = f.WriteString("[") //write for type array
			if err != nil {
				log.Fatal(err)
			}
		}
		_, err = f.WriteString("{\n")
		if err != nil {
			log.Fatal(err)
		}
		result := gjson.GetBytes(config, "properties."+key.String())
		result.ForEach(func(key1, value gjson.Result) bool {
			switch key1.String() {
			case "properties":
				{
					i := 0
					result1 := gjson.GetBytes(config, "properties."+key.String()+"."+key1.String())
					result1.ForEach(func(key2, value gjson.Result) bool {
						if i != 0 {
							f.WriteString(",")
						}
						i++
						f.WriteString("\"" + key2.String() + "\":")
						if isArray(getType(value.String())) == true {
							f.WriteString("[")
						}
						val := getFormat(value.String())
						_, err := f.WriteString(checkStr(value.String(), val) + "\n\n") //write properties of subsections
						if err != nil {
							log.Fatal(err)
						}
						if isArray(getType(value.String())) == true {
							f.WriteString("]")
						}
						return true
					})
				}
			case "items":
				{
					i := 0
					result1 := gjson.GetBytes(config, "properties."+key.String()+".items.properties")
					result1.ForEach(func(key2, value gjson.Result) bool {
						if i != 0 {
							f.WriteString(",")
						}
						i++
						f.WriteString("\"" + key2.String() + "\":")
						val := getFormat(value.String())

						if isArray(getType(value.String())) == true {
							f.WriteString("[")
							_, err := f.WriteString(checkStr(value.String(), val)) //write items of subsections
							if err != nil {
								log.Fatal(err)
							}
						} else {
							_, err := f.WriteString(checkStr(value.String(), val) + "\n\n") //write items of subsections
							if err != nil {
								log.Fatal(err)
							}
						}
						if isArray(getType(value.String())) == true {
							f.WriteString("]")
						}
						return true
					})
				}
			}
			return true
		})
		f.WriteString("}")
		if (isArray(getType(value.String()))) == true {
			_, err = f.WriteString("]")
			if err != nil {
				log.Fatal(err)
			}
		}
		return true
	})
	f.WriteString("}")
	f.Close()

	var schema schema.Data
	if err := json.Unmarshal(config, &schema); err != nil {
		log.Fatal(err)
	}

	tpl, err := template.ParseFiles(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create(doc)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if err = tpl.Execute(file, schema); err != nil {
		log.Fatal(err)
	}
	err = os.Remove(tmpl)
	if err != nil {
		log.Fatal(err)
		return
	}

}

func checkStr(x, typ string) string {
	d := "\"description\": "
	var val string
	if strings.Contains(x, d) {
		s := make([]string, 0, 2)
		s = strings.Split(x, d)
		switch typ {
		case "date":
			{
				val = "\"" + extractDate(x) + "\""
				return val
			}
		case "email":
			{
				s = strings.Split(s[1], "e.g. ")
				s = strings.Split(s[1], ",\n")
				val = "\"" + s[0]
				return val
			}

		case "uri":
			{
				pat := regexp.MustCompile(`https?://.*\.com`)
				s := "\"" + pat.FindString(x) + "\""
				return s
			}

		case "location":
			{
				str := "{\n \"lat\": \"0\",\n\"long\": \"0\"\n}"
				return str
			}

		default:
			{
				if strings.Contains(s[1], "e.g. ") {
					s = strings.Split(s[1], "e.g. ")
					s[1] = "\"" + s[1]
				}
				if strings.Contains(s[1], " - [") {
					s = strings.Split(s[1], " - [")
					s[1] = s[0] + "\""
				}
				s = strings.Split(s[1], "\n")
				val = s[0]

				val = strings.TrimSuffix(val, ",")
				return val
			}
		}
	}

	return "\"\""
}

func getFormat(x string) string {
	val := gjson.Get(x, "format")
	if !val.Exists() {
		return ""
	}
	return val.String()
}

func extractDate(date string) string {
	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	match := re.FindStringSubmatch(date)
	return match[0]
}

func getType(x string) string {
	val := gjson.Get(x, "type")
	if !val.Exists() {
		return ""
	}
	return val.String()

}

func isArray(x string) bool {
	if x == "array" {
		return true
	}
	return false
}
