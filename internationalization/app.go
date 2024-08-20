package main

import (
	"fmt"

	"github.com/steveanlorn/learning-go/internationalization/translations"
)

func Print() {
	en := translations.GetStatisMessage("en")
	ja := translations.GetStatisMessage("ja")
	zhTW := translations.GetStatisMessage("zh-TW")
	id := translations.GetStatisMessage("id")

	fmt.Println(en[translations.HelloWorldMessageID])
	fmt.Println(ja[translations.HelloWorldMessageID])
	fmt.Println(zhTW[translations.HelloWorldMessageID])
	fmt.Println(id[translations.HelloWorldMessageID])

	fmt.Println(en[translations.GoodByeMessageID])
	fmt.Println(ja[translations.GoodByeMessageID])
	fmt.Println(zhTW[translations.GoodByeMessageID])
	fmt.Println(id[translations.GoodByeMessageID])
}
