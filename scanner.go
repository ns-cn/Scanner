package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"strconv"
	"time"
)

const (
	HEADER = "¥"
)

func scanner(url string, user User, season Season) float64 {
	// 创建一个上下文
	basicCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 启动 Chrome 实例，并在超时后自动关闭
	ctx, cancel := chromedp.NewContext(basicCtx)
	defer cancel()

	// 打开页面
	err := chromedp.Run(ctx, chromedp.Navigate(url))
	if err != nil {
		fmt.Printf("Error navigating to %s: %s\n", url, err)
		return 0
	}

	// wait for the page to load completely
	err = chromedp.Run(ctx, chromedp.WaitVisible("searchBtn", chromedp.ByID))
	if err != nil {
		fmt.Println(err)
		return 0
	}

	// 输入数据
	err = chromedp.Run(ctx, chromedp.SendKeys("#inputTab > tbody > tr:nth-child(1) > td.item_cont > input[type=text]:nth-child(1)", user.Id))
	err = chromedp.Run(ctx, chromedp.SendKeys("#inputTab > tbody > tr:nth-child(2) > td.item_cont > input[type=text]:nth-child(1)", user.Name))
	err = chromedp.Run(ctx, chromedp.SendKeys("#inputTab > tbody > tr:nth-child(3) > td.item_cont > input[type=text]:nth-child(1)", season.Code))
	err = chromedp.Run(ctx, chromedp.SendKeys("#inputTab > tbody > tr:nth-child(4) > td.item_cont > input[type=text]:nth-child(1)", season.Name))
	if err != nil {
		fmt.Printf("Error sending keys to input element: %s\n", err)
		return 0
	}

	// 点击提交按钮
	err = chromedp.Run(ctx, chromedp.Click("chaxunbtn", chromedp.ByID))
	if err != nil {
		fmt.Printf("Error clicking submit button: %s\n", err)
		return 0
	}

	chromedp.Sleep(2 * time.Second)
	//err = chromedp.Run(ctx, chromedp.WaitVisible("tab1", chromedp.ByID))
	//if err != nil {
	//	fmt.Println(err)
	//	return 0
	//}

	// 获取提交后的页面内容
	var res string

	//err = chromedp.Run(ctx, chromedp.Text("#tab1 > li:nth-child(1) > span.s2 > span.font3 > input", &res, chromedp.NodeVisible))
	//if err != nil {
	//	fmt.Printf("Error getting result element HTML: %s\n", err)
	//	return 0
	//}
	//fmt.Println(res[len(HEADER):])
	err = chromedp.Run(ctx, chromedp.Value("#tab1 > li:nth-child(1) > span.s2 > span.font3 > input", &res, chromedp.NodeVisible))
	if err != nil {
		//fmt.Printf("Error getting result element HTML: %s\n", err)
		return 0
	}
	//fmt.Println(res[len(HEADER):])
	float, err := strconv.ParseFloat(res[len(HEADER):], 64)
	if err != nil {
		fmt.Println(err)
	}
	return float
}
