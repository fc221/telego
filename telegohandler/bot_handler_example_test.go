package telegohandler_test

import (
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

func ExampleNewMultiBotHandler() {
	botOne, _ := telego.NewBot("1234567890:aaaabbbbaaaabbbbaaaabbbbaaaabbbbccc")
	botTwo, _ := telego.NewBot("1234567891:aaaabbbbaaaabbbbaaaabbbbaaaabbbbccc")

	updatesWithBot := make(chan th.UpdateWithBot, 2)
	bh, _ := th.NewMultiBotHandler(updatesWithBot)

	bh.Handle(func(ctx *th.Context, update telego.Update) error {
		_ = ctx.Bot()
		_ = update.UpdateID
		return nil
	})

	updatesWithBot <- th.UpdateWithBot{Update: telego.Update{UpdateID: 1}, Bot: botOne}
	updatesWithBot <- th.UpdateWithBot{Update: telego.Update{UpdateID: 2}, Bot: botTwo}

	_ = bh
}
