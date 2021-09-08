package monitor

import (
	"github.com/ProjectAthenaa/sonic-core/fasttls"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent/product"
)

var(
	//eastbay, champs, footlocker, footaction
	siteMap = map[product.Site]string{
		product.SiteFootlocker_US: "footlocker",
		product.SiteFootaction_US: "footaction",
	}
)

func (tk *Task) GenerateDefaultHeaders(referrer string) fasttls.Headers {
	return fasttls.Headers{
		`user-agent`:         {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36"},
		`accept`:             {`application/json`},
		`accept-encoding`:    {`gzip, deflate, br`},
		`accept-language`:    {`en-us`},
		`content-type`:       {`application/json`},
		`sec-ch-ua`:          {`"Chromium";v="91", " Not A;Brand";v="99", "Google Chrome";v="91"`},
		`sec-ch-ua-mobile`:   {`?0`},
		`Sec-Fetch-Site`:     {`same-site`},
		`Sec-Fetch-Dest`:     {`empty`},
		`Sec-Fetch-Mode`:     {`cors`},
		`x-application-name`: {`web`},
		`referer`:            {referrer},
		`origin`:             {`https://www.target.com`},
		`Pragma`:             {`no-cache`},
		`Cache-Control`:      {`no-cache`},
		`Connection`:         {`keep-alive`},
	}
}
