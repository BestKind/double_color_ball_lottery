package services

import (
	"double_color_ball_lottery/backend/db"
	"double_color_ball_lottery/backend/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type LotteryService struct {
}

func NewLotteryService() *LotteryService {
	return &LotteryService{}
}

func (ls *LotteryService) InitHistoryData() {
	fmt.Println("init history data start")
	pageSize := 30
	page := 100
	client := http.DefaultClient
	for page >= 1 {
		var records []models.Lottery
		url := fmt.Sprintf("https://www.cwl.gov.cn/cwl_admin/front/cwlkj/search/kjxx/findDrawNotice?name=ssq&pageNo=%d&pageSize=%d&systemType=PC", page, pageSize)
		req, _ := http.NewRequest(http.MethodGet, url, nil)
		req.Header.Add("Accept", "application/json")
		// req.Header.Add("Cache-Control", "no-cache")
		req.Header.Add("Cookie", "HMF_CI=cb65f6fe110df193cd702234258501f2b7e1fe8d173ac2ac104821d6294f0264fa172846275ca7b1c2a98da36ff245741ea484281874bc25116818f9b808f0f075; 21_vq=3")
		// req.Header.Add("Host", "www.cwl.gov.cn")
		// req.Header.Add("Pragma", "no-cache")
		// req.Header.Add("Referer", "https://www.cwl.gov.cn/ygkj/wqkjgg/")
		// req.Header.Add("Sec-Fetch-Dest", "empty")
		// req.Header.Add("Sec-Fetch-Mode", "cors")
		// req.Header.Add("Sec-Fetch-Site", "same-origin")
		req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36")
		// req.Header.Add("X-Requested-With", "XMLHttpRequest")
		// req.Header.Add("sec-ch-ua-mobile", "?0")
		// req.Header.Add("sec-ch-ua-platform", "macOS")
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("init history data get page:%d err:%v\n", page, err)
			os.Exit(1)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("init history data get page:%d body err:%v\n", page, err)
			os.Exit(1)
		}
		if resp.StatusCode != 200 {
			fmt.Printf("init history data get page:%d status code:%d body:%s\n", page, resp.StatusCode, body)
			os.Exit(1)
		}
		_ = resp.Body.Close()
		var res models.LotteryRes
		json.Unmarshal(body, &res)
		page--
		if len(res.Result) <= 0 {
			continue
		}
		data := res.Result
		sort.Slice(data, func(i, j int) bool {
			code1, _ := strconv.Atoi(data[i].Code)
			code2, _ := strconv.Atoi(data[j].Code)
			return code1 < code2
		})
		for _, val := range res.Result {
			reds := strings.Split(val.Red, ",")
			level := val.Prizegrades
			sort.Slice(level, func(i, j int) bool {
				return level[i].Type < level[j].Type
			})
			records = append(records, models.Lottery{
				Version:  val.Code,
				OpenTime: val.Date,
				Week:     val.Week,
				Red1:     reds[0],
				Red2:     reds[1],
				Red3:     reds[2],
				Red4:     reds[3],
				Red5:     reds[4],
				Red6:     reds[5],
				Blue:     val.Blue,
				L1Val:    level[0].Typemoney,
				L1Cnt:    level[0].Typenum,
				L2Val:    level[1].Typemoney,
				L2Cnt:    level[1].Typenum,
				L3Cnt:    level[2].Typenum,
				Total:    val.Poolmoney,
				Sale:     val.Sales,
			})
		}
		db.DB.Create(records)
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println(page)
	fmt.Println("init history data finish")
}
