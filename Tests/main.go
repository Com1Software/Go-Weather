package main

import (
	"io"
	"math"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"

	"fmt"

	asciistring "github.com/Com1Software/Go-ASCII-String-Package"
)

// ----------------------------------------------------------------
func main() {
	fmt.Println("Weather")
	fmt.Printf("Operating System : %s\n", runtime.GOOS)
	//url := "https://forecast.weather.gov/MapClick.php?lat=41.5&lon=-81.7&unit=0&lg=english&FcstType=dwml"
	url := "https://forecast.weather.gov/MapClick.php?lat=41.25&lon=-81.44&unit=0&lg=english&FcstType=dwml"

	fmt.Printf("Current Time : %s\n", GetWeather(url, 0))
	fmt.Printf("Location : %s\n", GetWeather(url, 1))
	fmt.Printf("Hazard Warning : %s\n", GetWeather(url, 2))
	fmt.Printf("Current Teperature Fahrenheit : %s\n", GetWeather(url, 3))
	fmt.Printf("Current Teperature Celsius : %s\n", GetWeather(url, 4))
	fmt.Printf("Current Conditions : %s\n", GetWeather(url, 5))
	fmt.Printf("Current Conditions Icon: %s\n", GetWeather(url, 6))
	fmt.Printf("Sustaind Wind: %s\n", GetWeather(url, 7))
	fmt.Printf("Wind Gusts : %s\n", GetWeather(url, 8))
	fmt.Printf("Barometer : %s\n", GetWeather(url, 9))
	fmt.Printf("Humidity : %s\n", GetWeather(url, 10))
	fmt.Printf("Dewpoint Fahrenheit: %s\n", GetWeather(url, 11))
	fmt.Printf("Dewpoint Celsius: %s\n", GetWeather(url, 12))
	fmt.Printf("Visibility : %s\n", GetWeather(url, 13))
	fmt.Printf("Wind Chill : %s\n", GetWeather(url, 14))
	fmt.Printf("Forecast Slot 0 Period: %s\n", GetWeather(url, 15))
	fmt.Printf("Forecast Slot 0 Conditions : %s\n", GetWeather(url, 16))
	fmt.Printf("Forecast Slot 0 Icon : %s\n", GetWeather(url, 17))
	fmt.Printf("Forecast Slot 0 Text : %s\n", GetWeather(url, 18))
	fmt.Printf("Forecast Slot 1 Period: %s\n", GetWeather(url, 19))
	fmt.Printf("Forecast Slot 1 Conditions : %s\n", GetWeather(url, 20))
	fmt.Printf("Forecast Slot 1 Icon : %s\n", GetWeather(url, 21))
	fmt.Printf("Forecast Slot 1 Text : %s\n", GetWeather(url, 22))
	fmt.Printf("Forecast Slot 2 Period: %s\n", GetWeather(url, 23))
	fmt.Printf("Forecast Slot 2 Conditions : %s\n", GetWeather(url, 24))
	fmt.Printf("Forecast Slot 2 Icon : %s\n", GetWeather(url, 25))
	fmt.Printf("Forecast Slot 2 Text : %s\n", GetWeather(url, 26))
	fmt.Printf("Forecast Slot 3 Period: %s\n", GetWeather(url, 27))
	fmt.Printf("Forecast Slot 3 Conditions : %s\n", GetWeather(url, 28))
	fmt.Printf("Forecast Slot 3 Icon : %s\n", GetWeather(url, 29))
	fmt.Printf("Forecast Slot 3 Text : %s\n", GetWeather(url, 30))
	fmt.Printf("Forecast Slot 4 Period: %s\n", GetWeather(url, 31))
	fmt.Printf("Forecast Slot 4 Conditions : %s\n", GetWeather(url, 32))
	fmt.Printf("Forecast Slot 4 Icon : %s\n", GetWeather(url, 33))
	fmt.Printf("Forecast Slot 4 Text : %s\n", GetWeather(url, 34))
	fmt.Printf("Forecast Slot 5 Period: %s\n", GetWeather(url, 35))
	fmt.Printf("Forecast Slot 5 Conditions : %s\n", GetWeather(url, 36))
	fmt.Printf("Forecast Slot 5 Icon : %s\n", GetWeather(url, 37))
	fmt.Printf("Forecast Slot 5 Text : %s\n", GetWeather(url, 38))
	fmt.Printf("Forecast Slot 6 Period: %s\n", GetWeather(url, 39))
	fmt.Printf("Forecast Slot 6 Conditions : %s\n", GetWeather(url, 40))
	fmt.Printf("Forecast Slot 6 Icon : %s\n", GetWeather(url, 41))
	fmt.Printf("Forecast Slot 6 Text : %s\n", GetWeather(url, 42))
	fmt.Printf("Forecast Slot 7 Period: %s\n", GetWeather(url, 43))
	fmt.Printf("Forecast Slot 7 Conditions : %s\n", GetWeather(url, 44))
	fmt.Printf("Forecast Slot 7 Icon : %s\n", GetWeather(url, 45))
	fmt.Printf("Forecast Slot 7 Text : %s\n", GetWeather(url, 46))
	fmt.Printf("Forecast Slot 8 Period: %s\n", GetWeather(url, 47))
	fmt.Printf("Forecast Slot 8 Conditions : %s\n", GetWeather(url, 48))
	fmt.Printf("Forecast Slot 8 Icon : %s\n", GetWeather(url, 49))
	fmt.Printf("Forecast Slot 8 Text : %s\n", GetWeather(url, 50))
	fmt.Printf("Forecast Slot 9 Period: %s\n", GetWeather(url, 51))
	fmt.Printf("Forecast Slot 9 Conditions : %s\n", GetWeather(url, 52))
	fmt.Printf("Forecast Slot 9 Icon : %s\n", GetWeather(url, 53))
	fmt.Printf("Forecast Slot 9 Text : %s\n", GetWeather(url, 54))
	fmt.Printf("Forecast Slot 10 Period: %s\n", GetWeather(url, 55))
	fmt.Printf("Forecast Slot 10 Conditions : %s\n", GetWeather(url, 56))
	fmt.Printf("Forecast Slot 10 Icon : %s\n", GetWeather(url, 57))
	fmt.Printf("Forecast Slot 10 Text : %s\n", GetWeather(url, 58))
	fmt.Printf("Forecast Slot 11 Period: %s\n", GetWeather(url, 59))
	fmt.Printf("Forecast Slot 11 Conditions : %s\n", GetWeather(url, 60))
	fmt.Printf("Forecast Slot 11 Icon : %s\n", GetWeather(url, 61))
	fmt.Printf("Forecast Slot 11 Text : %s\n", GetWeather(url, 62))
	fmt.Printf("Forecast Slot 12 Period: %s\n", GetWeather(url, 63))
	fmt.Printf("Forecast Slot 12 Conditions : %s\n", GetWeather(url, 64))
	fmt.Printf("Forecast Slot 12 Icon : %s\n", GetWeather(url, 65))
	fmt.Printf("Forecast Slot 12 Text : %s\n", GetWeather(url, 66))

}

func GetWeather(url string, opt int) string {
	xdata := ""
	chr := ""
	ton := false
	word := ""
	loc := ""

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching URL: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "HTTP error: %v\n", resp.Status)
		os.Exit(1)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading response body: %v\n", err)
		os.Exit(1)
	}
	switch {
	case opt == 0:
		t := time.Now()
		formattedTime := t.Format(time.Kitchen)
		xdata = formattedTime
	//------------------------------------------------------------------------ Location
	case opt == 1:
		for x := 1; x < len(body); x++ {
			chr = string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<location" {
				tmp := ""
				tdata := string(body[x+20 : x+170])
				for xx := 1; xx < len(tdata)-18; xx++ {
					if tdata[xx:xx+18] == "<area-description>" {
						xx = xx + 18
						for xx := xx; xx < len(tdata)-18; xx++ {
							chr = string(tdata[xx : xx+1])
							if chr == "<" {
								break
							}
							tmp = tmp + chr
						}

					}

				}
				loc = tmp
			}

		}
		xdata = loc

		//------------------------------------------------------------------------ Hazzard Warning
	case opt == 2:
		hazctl := false
		haz := ""
		for x := 1; x < len(body); x++ {
			chr = string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<hazard headline" {
				haz = ""
				hazctl = true
				tdata := string(body[x : x+30])
				tt := false
				for xx := 1; xx < len(tdata); xx++ {
					chr = string(tdata[xx : xx+1])
					if tt {
						if asciistring.StringToASCII(chr) != 34 {
							haz = haz + chr
						}
					}
					switch {
					case asciistring.StringToASCII(chr) == 34 && tt == false:
						tt = true
					case asciistring.StringToASCII(chr) == 34 && tt == true:
						tt = false
					}
				}
			}

		}
		if hazctl {
			xdata = haz
		}

		//------------------------------------------------------------------------ Temperature
	case opt == 3 || opt == 4:
		for x := 1; x < len(body); x++ {
			chr = string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<temperature" {
				if string(body[x+8:x+16]) == "apparent" {
					temp := ""
					tdata := string(body[x+20 : x+100])
					for xx := 1; xx < len(tdata)-7; xx++ {
						if tdata[xx:xx+7] == "<value>" {
							xx = xx + 7
							for xx := xx; xx < len(tdata)-7; xx++ {
								chr = string(tdata[xx : xx+1])
								if chr == "<" {
									break
								}
								temp = temp + chr
							}
						}
					}
					if opt == 3 {
						xdata = temp
					}
					if opt == 4 {
						fahrenheit, err := strconv.ParseFloat(temp, 64)
						if err != nil {
							fmt.Println("Error converting Fahrenheit to float:", err)

						}
						celsius := (fahrenheit - 32) * 5 / 9
						xdata = fmt.Sprintf("%.0f", celsius)

					}
				}

			}

		}

		//------------------------------------------------------------------------ Current Conditions
	case opt == 5:
		cond := ""
		for x := 1; x < len(body); x++ {
			chr = string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {

				word = word + chr
			}

			if word == "<weather-conditions w" {
				cond = ""
				tdata := string(body[x+10 : x+50])
				tt := false
				for xx := 1; xx < len(tdata); xx++ {
					chr = string(tdata[xx : xx+1])
					if tt {
						if asciistring.StringToASCII(chr) != 34 {
							cond = cond + chr
						}
					}
					switch {
					case asciistring.StringToASCII(chr) == 34 && tt == false:
						tt = true
					case asciistring.StringToASCII(chr) == 34 && tt == true:
						tt = false
					}
				}
			}

		}
		xdata = cond

		//------------------------------------------------------------------------ Current Conditions Icon
	case opt == 6:
		cond := ""
		for x := 1; x < len(body); x++ {
			chr = string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {

				word = word + chr
			}

			if word == "<conditions-icon" {
				cond = ""
				tdata := string(body[x+0 : x+255])
				for xx := 1; xx < len(tdata)-11; xx++ {
					if tdata[xx:xx+11] == "<icon-link>" {
						xx = xx + 11
						for xx := xx; xx < len(tdata)-7; xx++ {
							chr = string(tdata[xx : xx+1])
							if chr == "<" {
								break
							}
							cond = cond + chr
						}
					}
				}
			}

		}
		xdata = cond

		//------------------------------------------------------------------------ Wind
	case opt == 7 || opt == 8:
		gust := ""
		sust := ""
		for x := 1; x < len(body); x++ {
			chr = string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<wind-speed" {
				if string(body[x+8:x+12]) == "gust" {
					tdata := string(body[x+20 : x+100])
					for xx := 1; xx < len(tdata)-7; xx++ {
						if tdata[xx:xx+7] == "<value>" {
							xx = xx + 7
							for xx := xx; xx < len(tdata)-7; xx++ {
								chr = string(tdata[xx : xx+1])
								if chr == "<" {
									break
								}
								gust = gust + chr
							}
						}
					}
				}
				if gust == "NA" {
					gust = ""
				}
				if string(body[x+8:x+17]) == "sustained" {
					tdata := string(body[x+20 : x+100])
					for xx := 1; xx < len(tdata)-7; xx++ {
						if tdata[xx:xx+7] == "<value>" {
							xx = xx + 7
							for xx := xx; xx < len(tdata)-7; xx++ {
								chr = string(tdata[xx : xx+1])
								if chr == "<" {
									break
								}
								sust = sust + chr
							}
						}
					}

				}

			}
		}
		if opt == 7 {
			xdata = sust
		}
		if opt == 8 {
			xdata = gust
		}
	//------------------------------------------------------------------------ Barometer
	case opt == 9:
		for x := 1; x < len(body); x++ {
			chr = string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<pressure" {
				if string(body[x+8:x+17]) == "barometer" {
					temp := ""
					tdata := string(body[x+20 : x+100])
					for xx := 1; xx < len(tdata)-7; xx++ {
						if tdata[xx:xx+7] == "<value>" {
							xx = xx + 7
							for xx := xx; xx < len(tdata)-7; xx++ {
								chr = string(tdata[xx : xx+1])
								if chr == "<" {
									break
								}
								temp = temp + chr
							}
						}
					}
					xdata = temp
				}

			}

		}

		//------------------------------------------------------------------------ Humidity
	case opt == 10:
		for x := 1; x < len(body); x++ {
			chr = string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<humidity" {
				temp := ""
				tdata := string(body[x+20 : x+100])
				for xx := 1; xx < len(tdata)-7; xx++ {
					if tdata[xx:xx+7] == "<value>" {
						xx = xx + 7
						for xx := xx; xx < len(tdata)-7; xx++ {
							chr = string(tdata[xx : xx+1])
							if chr == "<" {
								break
							}
							temp = temp + chr
						}
					}
					xdata = temp
				}

			}

		}

		//------------------------------------------------------------------------ Dewpoint
	case opt == 11 || opt == 12:
		for x := 1; x < len(body); x++ {
			chr = string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<temperature" {
				if string(body[x+8:x+11]) == "dew" {
					temp := ""
					tdata := string(body[x+20 : x+100])
					for xx := 1; xx < len(tdata)-7; xx++ {
						if tdata[xx:xx+7] == "<value>" {
							xx = xx + 7
							for xx := xx; xx < len(tdata)-7; xx++ {
								chr = string(tdata[xx : xx+1])
								if chr == "<" {
									break
								}
								temp = temp + chr
							}
						}
					}
					if opt == 11 {
						xdata = temp
					}
					if opt == 12 {
						fahrenheit, err := strconv.ParseFloat(temp, 64)
						if err != nil {
							fmt.Println("Error converting Fahrenheit to float:", err)

						}
						celsius := (fahrenheit - 32) * 5 / 9
						xdata = fmt.Sprintf("%.0f", celsius)
					}
				}

			}

		}

		//------------------------------------------------------------------------ Visibility
	case opt == 13:
		for x := 1; x < len(body); x++ {
			chr = string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<visibility" {
				temp := ""
				tdata := string(body[x+20 : x+100])
				for xx := 1; xx < len(tdata)-1; xx++ {
					if tdata[xx:xx+1] == ">" {
						xx = xx + 1
						for xx := xx; xx < len(tdata)-7; xx++ {
							chr = string(tdata[xx : xx+1])
							if chr == "<" {
								break
							}
							temp = temp + chr
						}
					}
					xdata = temp
				}

			}

		}
		//------------------------------------------------------------------------ Wind Chill
	case opt == 14:
		temp := ""
		sust := ""
		for x := 1; x < len(body); x++ {
			chr = string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<temperature" {
				if string(body[x+8:x+16]) == "apparent" {
					tdata := string(body[x+20 : x+100])
					for xx := 1; xx < len(tdata)-7; xx++ {
						if tdata[xx:xx+7] == "<value>" {
							xx = xx + 7
							for xx := xx; xx < len(tdata)-7; xx++ {
								chr = string(tdata[xx : xx+1])
								if chr == "<" {
									break
								}
								temp = temp + chr
							}
						}
					}
				}

			}

		}

		for x := 1; x < len(body); x++ {
			chr = string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<wind-speed" {
				if string(body[x+8:x+17]) == "sustained" {
					tdata := string(body[x+20 : x+100])
					for xx := 1; xx < len(tdata)-7; xx++ {
						if tdata[xx:xx+7] == "<value>" {
							xx = xx + 7
							for xx := xx; xx < len(tdata)-7; xx++ {
								chr = string(tdata[xx : xx+1])
								if chr == "<" {
									break
								}
								sust = sust + chr
							}
						}
					}

				}

			}
		}

		windSpeed, err := strconv.ParseFloat(sust, 64)
		if err != nil {
			fmt.Errorf("invalid wind speed input: %v", err)
		}

		tfv, err := strconv.ParseFloat(temp, 64)
		if err != nil {
			fmt.Println("Error:", err)
		}
		windChill := 35.74 + 0.6215*tfv - 35.75*math.Pow(windSpeed, 0.16) + 0.4275*tfv*math.Pow(windSpeed, 0.16)
		xdata = fmt.Sprintf("%.0f", windChill)

		//------------------------------------------------------------------------ Forecast Slot 0 Period
	case opt == 15:
		cond := ""
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<start-valid-time" {
				occ++
				if occ == 1 {
					tdata := string(body[x+10 : x+50])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if asciistring.StringToASCII(chr) != 34 {
								cond = cond + chr
							}
						}
						switch {
						case asciistring.StringToASCII(chr) == 34 && !tt:
							tt = true
						case asciistring.StringToASCII(chr) == 34 && tt:
							tt = false
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 0 Conditions
	case opt == 16:
		cond := ""
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<weather-conditions" {
				occ++
				if occ == 1 {
					tdata := string(body[x+10 : x+50])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if asciistring.StringToASCII(chr) != 34 {
								cond = cond + chr
							}
						}
						switch {
						case asciistring.StringToASCII(chr) == 34 && !tt:
							tt = true
						case asciistring.StringToASCII(chr) == 34 && tt:
							tt = false
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 0 Icon
	case opt == 17:
		cond := ""
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<icon" {
				occ++
				if occ == 1 {
					tdata := string(body[x : x+90])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if chr != "<" {
								if asciistring.StringToASCII(chr) != 10 {
									cond = cond + chr
								}
							}
						}
						switch {
						case chr == ">" && !tt:
							tt = true
						case chr == "<" && tt:
							tt = false
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 0 Text
	case opt == 18:
		cond := ""
		pass := 1
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<text" {
				occ++
				if occ == 1 {
					tdata := string(body[x : x+255])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if chr != "<" {

								if pass == 1 {
									cond = cond + chr
								}
							}
						}
						switch {
						case chr == ">" && !tt:
							tt = true
						case chr == "<" && tt:
							tt = false
							pass++
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 1 Period
	case opt == 19:
		cond := ""
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<start-valid-time" {
				occ++
				if occ == 2 {
					tdata := string(body[x+10 : x+50])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if asciistring.StringToASCII(chr) != 34 {
								cond = cond + chr
							}
						}
						switch {
						case asciistring.StringToASCII(chr) == 34 && !tt:
							tt = true
						case asciistring.StringToASCII(chr) == 34 && tt:
							tt = false
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 1 Conditions
	case opt == 20:
		cond := ""
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<weather-conditions" {
				occ++
				if occ == 2 {
					tdata := string(body[x+10 : x+50])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if asciistring.StringToASCII(chr) != 34 {
								cond = cond + chr
							}
						}
						switch {
						case asciistring.StringToASCII(chr) == 34 && !tt:
							tt = true
						case asciistring.StringToASCII(chr) == 34 && tt:
							tt = false
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 1 Icon
	case opt == 21:
		cond := ""
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<icon" {
				occ++
				if occ == 2 {
					tdata := string(body[x : x+90])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if chr != "<" {
								if asciistring.StringToASCII(chr) != 10 {
									cond = cond + chr
								}
							}
						}
						switch {
						case chr == ">" && !tt:
							tt = true
						case chr == "<" && tt:
							tt = false
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 1 Text
	case opt == 22:
		cond := ""
		pass := 1
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<text" {
				occ++
				if occ == 2 {
					tdata := string(body[x : x+255])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if chr != "<" {

								if pass == 1 {
									cond = cond + chr
								}
							}
						}
						switch {
						case chr == ">" && !tt:
							tt = true
						case chr == "<" && tt:
							tt = false
							pass++
						}
					}
					break
				}
			}
		}
		xdata = cond

		//	}

		//------------------------------------------------------------------------ Forecast Slot 2 Period
	case opt == 23:
		cond := ""
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<start-valid-time" {
				occ++
				if occ == 3 {
					tdata := string(body[x+10 : x+50])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if asciistring.StringToASCII(chr) != 34 {
								cond = cond + chr
							}
						}
						switch {
						case asciistring.StringToASCII(chr) == 34 && !tt:
							tt = true
						case asciistring.StringToASCII(chr) == 34 && tt:
							tt = false
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 2 Conditions
	case opt == 24:
		cond := ""
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<weather-conditions" {
				occ++
				if occ == 3 {
					tdata := string(body[x+10 : x+50])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if asciistring.StringToASCII(chr) != 34 {
								cond = cond + chr
							}
						}
						switch {
						case asciistring.StringToASCII(chr) == 34 && !tt:
							tt = true
						case asciistring.StringToASCII(chr) == 34 && tt:
							tt = false
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 2 Icon
	case opt == 25:
		cond := ""
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<icon" {
				occ++
				if occ == 3 {
					tdata := string(body[x : x+90])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if chr != "<" {
								if asciistring.StringToASCII(chr) != 10 {
									cond = cond + chr
								}
							}
						}
						switch {
						case chr == ">" && !tt:
							tt = true
						case chr == "<" && tt:
							tt = false
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 2 Text
	case opt == 26:
		cond := ""
		pass := 1
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<text" {
				occ++
				if occ == 3 {
					tdata := string(body[x : x+255])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if chr != "<" {

								if pass == 1 {
									cond = cond + chr
								}
							}
						}
						switch {
						case chr == ">" && !tt:
							tt = true
						case chr == "<" && tt:
							tt = false
							pass++
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 3 Period
	case opt == 27:
		cond := ""
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<start-valid-time" {
				occ++
				if occ == 4 {
					tdata := string(body[x+10 : x+50])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if asciistring.StringToASCII(chr) != 34 {
								cond = cond + chr
							}
						}
						switch {
						case asciistring.StringToASCII(chr) == 34 && !tt:
							tt = true
						case asciistring.StringToASCII(chr) == 34 && tt:
							tt = false
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 3 Conditions
	case opt == 28:
		cond := ""
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<weather-conditions" {
				occ++
				if occ == 4 {
					tdata := string(body[x+10 : x+50])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if asciistring.StringToASCII(chr) != 34 {
								cond = cond + chr
							}
						}
						switch {
						case asciistring.StringToASCII(chr) == 34 && !tt:
							tt = true
						case asciistring.StringToASCII(chr) == 34 && tt:
							tt = false
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 3 Icon
	case opt == 29:
		cond := ""
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<icon" {
				occ++
				if occ == 4 {
					tdata := string(body[x : x+90])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if chr != "<" {
								if asciistring.StringToASCII(chr) != 10 {
									cond = cond + chr
								}
							}
						}
						switch {
						case chr == ">" && !tt:
							tt = true
						case chr == "<" && tt:
							tt = false
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 3 Text
	case opt == 30:
		cond := ""
		pass := 1
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<text" {
				occ++
				if occ == 4 {
					tdata := string(body[x : x+255])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if chr != "<" {

								if pass == 1 {
									cond = cond + chr
								}
							}
						}
						switch {
						case chr == ">" && !tt:
							tt = true
						case chr == "<" && tt:
							tt = false
							pass++
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 4 Period
	case opt == 31:
		cond := ""
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<start-valid-time" {
				occ++
				if occ == 5 {
					tdata := string(body[x+10 : x+50])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if asciistring.StringToASCII(chr) != 34 {
								cond = cond + chr
							}
						}
						switch {
						case asciistring.StringToASCII(chr) == 34 && !tt:
							tt = true
						case asciistring.StringToASCII(chr) == 34 && tt:
							tt = false
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 4 Conditions
	case opt == 32:
		cond := ""
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<weather-conditions" {
				occ++
				if occ == 5 {
					tdata := string(body[x+10 : x+50])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if asciistring.StringToASCII(chr) != 34 {
								cond = cond + chr
							}
						}
						switch {
						case asciistring.StringToASCII(chr) == 34 && !tt:
							tt = true
						case asciistring.StringToASCII(chr) == 34 && tt:
							tt = false
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 4 Icon
	case opt == 33:
		cond := ""
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<icon" {
				occ++
				if occ == 5 {
					tdata := string(body[x : x+90])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if chr != "<" {
								if asciistring.StringToASCII(chr) != 10 {
									cond = cond + chr
								}
							}
						}
						switch {
						case chr == ">" && !tt:
							tt = true
						case chr == "<" && tt:
							tt = false
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 4 Text
	case opt == 34:
		cond := ""
		pass := 1
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<text" {
				occ++
				if occ == 5 {
					tdata := string(body[x : x+255])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if chr != "<" {

								if pass == 1 {
									cond = cond + chr
								}
							}
						}
						switch {
						case chr == ">" && !tt:
							tt = true
						case chr == "<" && tt:
							tt = false
							pass++
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 5 Period
	case opt == 35:
		cond := ""
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<start-valid-time" {
				occ++
				if occ == 6 {
					tdata := string(body[x+10 : x+50])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if asciistring.StringToASCII(chr) != 34 {
								cond = cond + chr
							}
						}
						switch {
						case asciistring.StringToASCII(chr) == 34 && !tt:
							tt = true
						case asciistring.StringToASCII(chr) == 34 && tt:
							tt = false
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 5 Conditions
	case opt == 36:
		cond := ""
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<weather-conditions" {
				occ++
				if occ == 6 {
					tdata := string(body[x+10 : x+50])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if asciistring.StringToASCII(chr) != 34 {
								cond = cond + chr
							}
						}
						switch {
						case asciistring.StringToASCII(chr) == 34 && !tt:
							tt = true
						case asciistring.StringToASCII(chr) == 34 && tt:
							tt = false
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 5 Icon
	case opt == 37:
		cond := ""
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<icon" {
				occ++
				if occ == 6 {
					tdata := string(body[x : x+90])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if chr != "<" {
								if asciistring.StringToASCII(chr) != 10 {
									cond = cond + chr
								}
							}
						}
						switch {
						case chr == ">" && !tt:
							tt = true
						case chr == "<" && tt:
							tt = false
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 5 Text
	case opt == 38:
		cond := ""
		pass := 1
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<text" {
				occ++
				if occ == 6 {
					tdata := string(body[x : x+255])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if chr != "<" {

								if pass == 1 {
									cond = cond + chr
								}
							}
						}
						switch {
						case chr == ">" && !tt:
							tt = true
						case chr == "<" && tt:
							tt = false
							pass++
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 6 Period
	case opt == 39:
		cond := ""
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<start-valid-time" {
				occ++
				if occ == 7 {
					tdata := string(body[x+10 : x+50])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if asciistring.StringToASCII(chr) != 34 {
								cond = cond + chr
							}
						}
						switch {
						case asciistring.StringToASCII(chr) == 34 && !tt:
							tt = true
						case asciistring.StringToASCII(chr) == 34 && tt:
							tt = false
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 6 Conditions
	case opt == 40:
		cond := ""
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<weather-conditions" {
				occ++
				if occ == 7 {
					tdata := string(body[x+10 : x+50])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if asciistring.StringToASCII(chr) != 34 {
								cond = cond + chr
							}
						}
						switch {
						case asciistring.StringToASCII(chr) == 34 && !tt:
							tt = true
						case asciistring.StringToASCII(chr) == 34 && tt:
							tt = false
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 6 Icon
	case opt == 41:
		cond := ""
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<icon" {
				occ++
				if occ == 7 {
					tdata := string(body[x : x+90])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if chr != "<" {
								if asciistring.StringToASCII(chr) != 10 {
									cond = cond + chr
								}
							}
						}
						switch {
						case chr == ">" && !tt:
							tt = true
						case chr == "<" && tt:
							tt = false
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 6 Text
	case opt == 42:
		cond := ""
		pass := 1
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<text" {
				occ++
				if occ == 7 {
					tdata := string(body[x : x+255])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if chr != "<" {

								if pass == 1 {
									cond = cond + chr
								}
							}
						}
						switch {
						case chr == ">" && !tt:
							tt = true
						case chr == "<" && tt:
							tt = false
							pass++
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 7 Period
	case opt == 43:
		cond := ""
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<start-valid-time" {
				occ++
				if occ == 8 {
					tdata := string(body[x+10 : x+50])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if asciistring.StringToASCII(chr) != 34 {
								cond = cond + chr
							}
						}
						switch {
						case asciistring.StringToASCII(chr) == 34 && !tt:
							tt = true
						case asciistring.StringToASCII(chr) == 34 && tt:
							tt = false
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 7 Conditions
	case opt == 44:
		cond := ""
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<weather-conditions" {
				occ++
				if occ == 8 {
					tdata := string(body[x+10 : x+50])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if asciistring.StringToASCII(chr) != 34 {
								cond = cond + chr
							}
						}
						switch {
						case asciistring.StringToASCII(chr) == 34 && !tt:
							tt = true
						case asciistring.StringToASCII(chr) == 34 && tt:
							tt = false
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 7 Icon
	case opt == 45:
		cond := ""
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<icon" {
				occ++
				if occ == 8 {
					tdata := string(body[x : x+90])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if chr != "<" {
								if asciistring.StringToASCII(chr) != 10 {
									cond = cond + chr
								}
							}
						}
						switch {
						case chr == ">" && !tt:
							tt = true
						case chr == "<" && tt:
							tt = false
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 7 Text
	case opt == 46:
		cond := ""
		pass := 1
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<text" {
				occ++
				if occ == 8 {
					tdata := string(body[x : x+255])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if chr != "<" {

								if pass == 1 {
									cond = cond + chr
								}
							}
						}
						switch {
						case chr == ">" && !tt:
							tt = true
						case chr == "<" && tt:
							tt = false
							pass++
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 8 Period
	case opt == 47:
		cond := ""
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<start-valid-time" {
				occ++
				if occ == 9 {
					tdata := string(body[x+10 : x+50])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if asciistring.StringToASCII(chr) != 34 {
								cond = cond + chr
							}
						}
						switch {
						case asciistring.StringToASCII(chr) == 34 && !tt:
							tt = true
						case asciistring.StringToASCII(chr) == 34 && tt:
							tt = false
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 8 Conditions
	case opt == 48:
		cond := ""
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<weather-conditions" {
				occ++
				if occ == 9 {
					tdata := string(body[x+10 : x+50])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if asciistring.StringToASCII(chr) != 34 {
								cond = cond + chr
							}
						}
						switch {
						case asciistring.StringToASCII(chr) == 34 && !tt:
							tt = true
						case asciistring.StringToASCII(chr) == 34 && tt:
							tt = false
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 8 Icon
	case opt == 49:
		cond := ""
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<icon" {
				occ++
				if occ == 9 {
					tdata := string(body[x : x+90])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if chr != "<" {
								if asciistring.StringToASCII(chr) != 10 {
									cond = cond + chr
								}
							}
						}
						switch {
						case chr == ">" && !tt:
							tt = true
						case chr == "<" && tt:
							tt = false
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 8 Text
	case opt == 50:
		cond := ""
		pass := 1
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<text" {
				occ++
				if occ == 9 {
					tdata := string(body[x : x+255])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if chr != "<" {

								if pass == 1 {
									cond = cond + chr
								}
							}
						}
						switch {
						case chr == ">" && !tt:
							tt = true
						case chr == "<" && tt:
							tt = false
							pass++
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 9 Period
	case opt == 51:
		cond := ""
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<start-valid-time" {
				occ++
				if occ == 10 {
					tdata := string(body[x+10 : x+50])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if asciistring.StringToASCII(chr) != 34 {
								cond = cond + chr
							}
						}
						switch {
						case asciistring.StringToASCII(chr) == 34 && !tt:
							tt = true
						case asciistring.StringToASCII(chr) == 34 && tt:
							tt = false
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 9 Conditions
	case opt == 52:
		cond := ""
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<weather-conditions" {
				occ++
				if occ == 10 {
					tdata := string(body[x+10 : x+50])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if asciistring.StringToASCII(chr) != 34 {
								cond = cond + chr
							}
						}
						switch {
						case asciistring.StringToASCII(chr) == 34 && !tt:
							tt = true
						case asciistring.StringToASCII(chr) == 34 && tt:
							tt = false
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 9 Icon
	case opt == 53:
		cond := ""
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<icon" {
				occ++
				if occ == 10 {
					tdata := string(body[x : x+90])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if chr != "<" {
								if asciistring.StringToASCII(chr) != 10 {
									cond = cond + chr
								}
							}
						}
						switch {
						case chr == ">" && !tt:
							tt = true
						case chr == "<" && tt:
							tt = false
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 9 Text
	case opt == 54:
		cond := ""
		pass := 1
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<text" {
				occ++
				if occ == 10 {
					tdata := string(body[x : x+255])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if chr != "<" {

								if pass == 1 {
									cond = cond + chr
								}
							}
						}
						switch {
						case chr == ">" && !tt:
							tt = true
						case chr == "<" && tt:
							tt = false
							pass++
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 10 Period
	case opt == 55:
		cond := ""
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<start-valid-time" {
				occ++
				if occ == 11 {
					tdata := string(body[x+10 : x+50])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if asciistring.StringToASCII(chr) != 34 {
								cond = cond + chr
							}
						}
						switch {
						case asciistring.StringToASCII(chr) == 34 && !tt:
							tt = true
						case asciistring.StringToASCII(chr) == 34 && tt:
							tt = false
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 10 Conditions
	case opt == 56:
		cond := ""
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<weather-conditions" {
				occ++
				if occ == 11 {
					tdata := string(body[x+10 : x+50])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if asciistring.StringToASCII(chr) != 34 {
								cond = cond + chr
							}
						}
						switch {
						case asciistring.StringToASCII(chr) == 34 && !tt:
							tt = true
						case asciistring.StringToASCII(chr) == 34 && tt:
							tt = false
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 10 Icon
	case opt == 57:
		cond := ""
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<icon" {
				occ++
				if occ == 11 {
					tdata := string(body[x : x+90])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if chr != "<" {
								if asciistring.StringToASCII(chr) != 10 {
									cond = cond + chr
								}
							}
						}
						switch {
						case chr == ">" && !tt:
							tt = true
						case chr == "<" && tt:
							tt = false
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 10 Text
	case opt == 58:
		cond := ""
		pass := 1
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<text" {
				occ++
				if occ == 11 {
					tdata := string(body[x : x+255])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if chr != "<" {

								if pass == 1 {
									cond = cond + chr
								}
							}
						}
						switch {
						case chr == ">" && !tt:
							tt = true
						case chr == "<" && tt:
							tt = false
							pass++
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 11 Period
	case opt == 59:
		cond := ""
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<start-valid-time" {
				occ++
				if occ == 12 {
					tdata := string(body[x+10 : x+50])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if asciistring.StringToASCII(chr) != 34 {
								cond = cond + chr
							}
						}
						switch {
						case asciistring.StringToASCII(chr) == 34 && !tt:
							tt = true
						case asciistring.StringToASCII(chr) == 34 && tt:
							tt = false
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 11 Conditions
	case opt == 60:
		cond := ""
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<weather-conditions" {
				occ++
				if occ == 12 {
					tdata := string(body[x+10 : x+50])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if asciistring.StringToASCII(chr) != 34 {
								cond = cond + chr
							}
						}
						switch {
						case asciistring.StringToASCII(chr) == 34 && !tt:
							tt = true
						case asciistring.StringToASCII(chr) == 34 && tt:
							tt = false
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 11 Icon
	case opt == 61:
		cond := ""
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<icon" {
				occ++
				if occ == 12 {
					tdata := string(body[x : x+90])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if chr != "<" {
								if asciistring.StringToASCII(chr) != 10 {
									cond = cond + chr
								}
							}
						}
						switch {
						case chr == ">" && !tt:
							tt = true
						case chr == "<" && tt:
							tt = false
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 11 Text
	case opt == 62:
		cond := ""
		pass := 1
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<text" {
				occ++
				if occ == 12 {
					tdata := string(body[x : x+255])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if chr != "<" {

								if pass == 1 {
									cond = cond + chr
								}
							}
						}
						switch {
						case chr == ">" && !tt:
							tt = true
						case chr == "<" && tt:
							tt = false
							pass++
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 12 Period
	case opt == 63:
		cond := ""
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<start-valid-time" {
				occ++
				if occ == 13 {
					tdata := string(body[x+10 : x+50])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if asciistring.StringToASCII(chr) != 34 {
								cond = cond + chr
							}
						}
						switch {
						case asciistring.StringToASCII(chr) == 34 && !tt:
							tt = true
						case asciistring.StringToASCII(chr) == 34 && tt:
							tt = false
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 12 Conditions
	case opt == 64:
		cond := ""
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<weather-conditions" {
				occ++
				if occ == 13 {
					tdata := string(body[x+10 : x+50])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if asciistring.StringToASCII(chr) != 34 {
								cond = cond + chr
							}
						}
						switch {
						case asciistring.StringToASCII(chr) == 34 && !tt:
							tt = true
						case asciistring.StringToASCII(chr) == 34 && tt:
							tt = false
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 12 Icon
	case opt == 65:
		cond := ""
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<icon" {
				occ++
				if occ == 13 {
					tdata := string(body[x : x+90])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if chr != "<" {
								if asciistring.StringToASCII(chr) != 10 {
									cond = cond + chr
								}
							}
						}
						switch {
						case chr == ">" && !tt:
							tt = true
						case chr == "<" && tt:
							tt = false
						}
					}
					break
				}
			}
		}
		xdata = cond

		//------------------------------------------------------------------------ Forecast Slot 12 Text
	case opt == 66:
		cond := ""
		pass := 1
		occ := 0
		for x := 1; x < len(body); x++ {
			chr := string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<text" {
				occ++
				if occ == 13 {
					tdata := string(body[x : x+255])
					tt := false
					for xx := 1; xx < len(tdata); xx++ {
						chr = string(tdata[xx : xx+1])
						if tt {
							if chr != "<" {

								if pass == 1 {
									cond = cond + chr
								}
							}
						}
						switch {
						case chr == ">" && !tt:
							tt = true
						case chr == "<" && tt:
							tt = false
							pass++
						}
					}
					break
				}
			}
		}
		xdata = cond

	}
	return xdata

}
