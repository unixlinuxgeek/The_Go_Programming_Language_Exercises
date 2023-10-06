### Рекурсия


```shell
$ ./fetch https://golang.org | ./findLinks1 
/
#main-content
#
/solutions/case-studies
/solutions/use-cases
/security/
/learn/
#
/doc/effective_go
/doc
https://pkg.go.dev/std
/doc/devel/release
https://pkg.go.dev
#
/talks/
https://www.meetup.com/pro/go
https://github.com/golang/go/wiki/Conferences
/blog
/help
https://groups.google.com/g/golang-nuts
https://github.com/golang
https://twitter.com/golang
https://www.reddit.com/r/golang/
https://invite.slack.golangbridge.org/
https://stackoverflow.com/tags/go
/
#
#
/solutions/case-studies
/solutions/use-cases
/security/
/learn/
#
#
/doc/effective_go
/doc
https://pkg.go.dev/std
/doc/devel/release
https://pkg.go.dev
#
#
/talks/
https://www.meetup.com/pro/go
https://github.com/golang/go/wiki/Conferences
/blog
/help
https://groups.google.com/g/golang-nuts
https://github.com/golang
https://twitter.com/golang
https://www.reddit.com/r/golang/
https://invite.slack.golangbridge.org/
https://stackoverflow.com/tags/go
/learn/
/dl
/dl/
/dl
/solutions/
/solutions/google/
/solutions/paypal
/solutions/americanexpress
/solutions/mercadolibre
https://bitly.com/blog/why-we-write-everything-in-go/?utm_source=go-dev&utm_medium=referral&utm_campaign=go-dev&utm_content=case-study
https://medium.com/capital-one-tech/a-serverless-and-go-journey-credit-offers-api-74ef1f9fde7f
https://www.cockroachlabs.com/blog/why-go-was-the-right-choice-for-cockroachdb/
https://blogs.dropbox.com/tech/2014/07/open-sourcing-our-go-libraries/
https://blog.cloudflare.com/graceful-upgrades-in-go/
https://entgo.io/blog/2019/10/03/introducing-ent/
https://cloudblogs.microsoft.com/opensource/2018/02/21/go-lang-brian-ketelsen-explains-fast-growth/
https://medium.com/tech-at-wildlife-studios/pitaya-wildlifes-golang-go-af57865f7a11
https://medium.com/netflix-techblog/application-data-caching-using-ssds-5bf25df851ef
https://technology.riotgames.com/news/leveraging-golang-game-development-and-operations
https://www.zdnet.com/article/salesforce-why-we-ditched-python-for-googles-go-language-in-einstein-analytics/
https://blog.twitch.tv/en/2016/07/05/gos-march-to-low-latency-gc-a6fa96f06eb7/
https://blog.twitter.com/engineering/en_us/a/2015/handling-five-billion-sessions-a-day-in-real-time.html
https://eng.uber.com/aresdb/
/tour/
https://cloud.google.com/go/home
https://aws.amazon.com/sdk-for-go/
https://github.com/Azure/azure-sdk-for-go
/solutions/cloud/
https://github.com/spf13/cobra
https://github.com/spf13/viper
https://github.com/urfave/cli
https://github.com/go-delve/delve
https://github.com/chzyer/readline
/solutions/clis/
/pkg/net/http/
/pkg/html/template/
https://github.com/flosch/pongo2
/pkg/database/sql/
https://github.com/elastic/go-elasticsearch
/solutions/webdev/
https://github.com/open-telemetry/opentelemetry-go
https://github.com/istio/istio
https://github.com/urfave/cli
/solutions/devops/
/solutions/use-cases
/learn/
/doc/install/
/learn#guided-learning-journeys
/learn#online-learning
/learn#featured-books
/learn#self-paced-labs
https://www.ardanlabs.com/
https://www.gopherguides.com/
https://bosssauce.it/services/training
https://github.com/shijuvar/gokit/tree/master/training
/solutions/
/solutions/use-cases
/solutions/case-studies
/learn/
/play
/tour/
https://stackoverflow.com/questions/tagged/go?tab=Newest
/help/
https://pkg.go.dev
/pkg/
https://pkg.go.dev/about
/project
/dl/
/blog/
https://github.com/golang/go/issues
/doc/devel/release
/brand
/conduct
https://www.twitter.com/golang
https://www.twitter.com/golang
https://github.com/golang
https://invite.slack.golangbridge.org/
https://reddit.com/r/golang
https://www.meetup.com/pro/go
https://golangweekly.com/
/copyright
/tos
http://www.google.com/intl/en/policies/privacy/
/s/website-issue
https://google.com
https://policies.google.com/technologies/cookies
geek@lenovo:~/GolandProjects/The_Go_Programming_Language_Exercises/ch5$ 

```

geek@lenovo:~/GolandProjects/The_Go_Programming_Language_Exercises/ch5$
```shell
$ ./fetch https://golang.org | ./outLine
[html]
[html head]
[html head link]
[html head script]
[html head meta]
[html head meta]
[html head meta]
[html head link]
[html head link]
[html head link]
[html head link]
[html head link]
[html head script]
[html head script]
[html head meta]
[html head meta]
[html head title]
[html head meta]
[html head meta]
[html head meta]
[html head meta]
[html head meta]
[html head meta]
[html body]
[html body noscript]
[html body header]
[html body header div]
[html body header div nav]
[html body header div nav a]
[html body header div nav a img]
[html body header div nav div]
[html body header div nav div a]
[html body header div nav div]
[html body header div nav div ul]
[html body header div nav div ul li]
[html body header div nav div ul li a]
[html body header div nav div ul li a i]
[html body header div nav div ul li div]
[html body header div nav div ul li ul]
[html body header div nav div ul li ul li]
[html body header div nav div ul li ul li div]
[html body header div nav div ul li ul li div a]
[html body header div nav div ul li ul li p]
[html body header div nav div ul li ul li]
[html body header div nav div ul li ul li div]
[html body header div nav div ul li ul li div a]
[html body header div nav div ul li ul li p]
[html body header div nav div ul li ul li]
[html body header div nav div ul li ul li div]
[html body header div nav div ul li ul li div a]
[html body header div nav div ul li ul li p]
[html body header div nav div ul li]
[html body header div nav div ul li a]
[html body header div nav div ul li div]
[html body header div nav div ul li]
[html body header div nav div ul li a]
[html body header div nav div ul li a i]
[html body header div nav div ul li div]
[html body header div nav div ul li ul]
[html body header div nav div ul li ul li]
[html body header div nav div ul li ul li div]
[html body header div nav div ul li ul li div a]
[html body header div nav div ul li ul li p]
[html body header div nav div ul li ul li]
[html body header div nav div ul li ul li div]
[html body header div nav div ul li ul li div a]
[html body header div nav div ul li ul li p]
[html body header div nav div ul li ul li]
[html body header div nav div ul li ul li div]
[html body header div nav div ul li ul li div a]
[html body header div nav div ul li ul li p]
[html body header div nav div ul li ul li]
[html body header div nav div ul li ul li div]
[html body header div nav div ul li ul li div a]
[html body header div nav div ul li ul li p]
[html body header div nav div ul li]
[html body header div nav div ul li a]
[html body header div nav div ul li div]
[html body header div nav div ul li]
[html body header div nav div ul li a]
[html body header div nav div ul li a i]
[html body header div nav div ul li div]
[html body header div nav div ul li ul]
[html body header div nav div ul li ul li]
[html body header div nav div ul li ul li div]
[html body header div nav div ul li ul li div a]
[html body header div nav div ul li ul li p]
[html body header div nav div ul li ul li]
[html body header div nav div ul li ul li div]
[html body header div nav div ul li ul li div a]
[html body header div nav div ul li ul li div a i]
[html body header div nav div ul li ul li p]
[html body header div nav div ul li ul li]
[html body header div nav div ul li ul li div]
[html body header div nav div ul li ul li div a]
[html body header div nav div ul li ul li div a i]
[html body header div nav div ul li ul li p]
[html body header div nav div ul li ul li]
[html body header div nav div ul li ul li div]
[html body header div nav div ul li ul li div a]
[html body header div nav div ul li ul li p]
[html body header div nav div ul li ul li]
[html body header div nav div ul li ul li div]
[html body header div nav div ul li ul li div a]
[html body header div nav div ul li ul li p]
[html body header div nav div ul li ul li]
[html body header div nav div ul li ul li div]
[html body header div nav div ul li ul li p]
[html body header div nav div ul li ul li div]
[html body header div nav div ul li ul li div a]
[html body header div nav div ul li ul li div a img]
[html body header div nav div ul li ul li div a]
[html body header div nav div ul li ul li div a img]
[html body header div nav div ul li ul li div a]
[html body header div nav div ul li ul li div a img]
[html body header div nav div ul li ul li div a]
[html body header div nav div ul li ul li div a img]
[html body header div nav div ul li ul li div a]
[html body header div nav div ul li ul li div a img]
[html body header div nav div ul li ul li div a]
[html body header div nav div ul li ul li div a img]
[html body header div nav div button]
[html body aside]
[html body aside nav]
[html body aside nav div]
[html body aside nav div a]
[html body aside nav div a img]
[html body aside nav ul]
[html body aside nav ul li]
[html body aside nav ul li a]
[html body aside nav ul li a span]
[html body aside nav ul li a i]
[html body aside nav ul li div]
[html body aside nav ul li div nav]
[html body aside nav ul li div nav div]
[html body aside nav ul li div nav div a]
[html body aside nav ul li div nav div a i]
[html body aside nav ul li div nav ul]
[html body aside nav ul li div nav ul li]
[html body aside nav ul li div nav ul li a]
[html body aside nav ul li div nav ul li]
[html body aside nav ul li div nav ul li a]
[html body aside nav ul li div nav ul li]
[html body aside nav ul li div nav ul li a]
[html body aside nav ul li]
[html body aside nav ul li a]
[html body aside nav ul li]
[html body aside nav ul li a]
[html body aside nav ul li a span]
[html body aside nav ul li a i]
[html body aside nav ul li div]
[html body aside nav ul li div nav]
[html body aside nav ul li div nav div]
[html body aside nav ul li div nav div a]
[html body aside nav ul li div nav div a i]
[html body aside nav ul li div nav ul]
[html body aside nav ul li div nav ul li]
[html body aside nav ul li div nav ul li a]
[html body aside nav ul li div nav ul li]
[html body aside nav ul li div nav ul li a]
[html body aside nav ul li div nav ul li]
[html body aside nav ul li div nav ul li a]
[html body aside nav ul li div nav ul li]
[html body aside nav ul li div nav ul li a]
[html body aside nav ul li]
[html body aside nav ul li a]
[html body aside nav ul li]
[html body aside nav ul li a]
[html body aside nav ul li a span]
[html body aside nav ul li a i]
[html body aside nav ul li div]
[html body aside nav ul li div nav]
[html body aside nav ul li div nav div]
[html body aside nav ul li div nav div a]
[html body aside nav ul li div nav div a i]
[html body aside nav ul li div nav ul]
[html body aside nav ul li div nav ul li]
[html body aside nav ul li div nav ul li a]
[html body aside nav ul li div nav ul li]
[html body aside nav ul li div nav ul li a]
[html body aside nav ul li div nav ul li a i]
[html body aside nav ul li div nav ul li]
[html body aside nav ul li div nav ul li a]
[html body aside nav ul li div nav ul li a i]
[html body aside nav ul li div nav ul li]
[html body aside nav ul li div nav ul li a]
[html body aside nav ul li div nav ul li]
[html body aside nav ul li div nav ul li a]
[html body aside nav ul li div nav ul li]
[html body aside nav ul li div nav ul li div]
[html body aside nav ul li div nav ul li div]
[html body aside nav ul li div nav ul li div a]
[html body aside nav ul li div nav ul li div a img]
[html body aside nav ul li div nav ul li div a]
[html body aside nav ul li div nav ul li div a img]
[html body aside nav ul li div nav ul li div a]
[html body aside nav ul li div nav ul li div a img]
[html body aside nav ul li div nav ul li div a]
[html body aside nav ul li div nav ul li div a img]
[html body aside nav ul li div nav ul li div a]
[html body aside nav ul li div nav ul li div a img]
[html body aside nav ul li div nav ul li div a]
[html body aside nav ul li div nav ul li div a img]
[html body div]
[html body main]
[html body main section]
[html body main section div]
[html body main section div div]
[html body main section div div h1]
[html body main section div div ul]
[html body main section div div ul li]
[html body main section div div ul li svg]
[html body main section div div ul li svg path]
[html body main section div div ul li]
[html body main section div div ul li svg]
[html body main section div div ul li svg path]
[html body main section div div ul li]
[html body main section div div ul li svg]
[html body main section div div ul li svg path]
[html body main section div div ul li]
[html body main section div div ul li svg]
[html body main section div div ul li svg path]
[html body main section div div]
[html body main section div div div]
[html body main section div div div a]
[html body main section div div div a]
[html body main section div div div div]
[html body main section div div div div]
[html body main section div div div]
[html body main section div div div p]
[html body main section div div div p a]
[html body main section div div div p a]
[html body main section div div div p a]
[html body main section div div div p a]
[html body main section div div div p]
[html body main section div div div p code]
[html body main section div div div p a]
[html body main section div div]
[html body main section div div]
[html body main section div div img]
[html body main section]
[html body main section div]
[html body main section div div]
[html body main section div div h2]
[html body main section div div p]
[html body main section div div p a]
[html body main section div div]
[html body main section div div ul]
[html body main section div div ul li]
[html body main section div div ul li a]
[html body main section div div ul li a img]
[html body main section div div ul li a p]
[html body main section div div ul li]
[html body main section div div ul li a]
[html body main section div div ul li a img]
[html body main section div div ul li a p]
[html body main section div div ul li]
[html body main section div div ul li a]
[html body main section div div ul li a img]
[html body main section div div ul li a p]
[html body main section div div ul li]
[html body main section div div ul li a]
[html body main section div div ul li a img]
[html body main section div div ul li a p]
[html body main section div div ul li]
[html body main section div div ul li a]
[html body main section div div ul li a img]
[html body main section div div ul li]
[html body main section div div ul li a]
[html body main section div div ul li a img]
[html body main section div div ul li]
[html body main section div div ul li a]
[html body main section div div ul li a img]
[html body main section div div ul li]
[html body main section div div ul li a]
[html body main section div div ul li a img]
[html body main section div div ul li]
[html body main section div div ul li a]
[html body main section div div ul li a img]
[html body main section div div ul li]
[html body main section div div ul li a]
[html body main section div div ul li a img]
[html body main section div div ul li]
[html body main section div div ul li a]
[html body main section div div ul li a img]
[html body main section div div ul li]
[html body main section div div ul li a]
[html body main section div div ul li a img]
[html body main section div div ul li]
[html body main section div div ul li a]
[html body main section div div ul li a img]
[html body main section div div ul li]
[html body main section div div ul li a]
[html body main section div div ul li a img]
[html body main section div div ul li]
[html body main section div div ul li a]
[html body main section div div ul li a img]
[html body main section div div ul li]
[html body main section div div ul li a]
[html body main section div div ul li a img]
[html body main section div div ul li]
[html body main section div div ul li a]
[html body main section div div ul li a img]
[html body main section div div ul li]
[html body main section div div ul li a]
[html body main section div div ul li a img]
[html body main section]
[html body main section div]
[html body main section div div]
[html body main section div div div]
[html body main section div div div ul]
[html body main section div div div ul li]
[html body main section div div div ul li div]
[html body main section div div div ul li div div]
[html body main section div div div ul li div div p]
[html body main section div div div ul li div div p strong]
[html body main section div div div ul li div div div]
[html body main section div div div ul li div div div span]
[html body main section div div div ul li div div div span]
[html body main section div div div ul li]
[html body main section div div div ul li div]
[html body main section div div div ul li div div]
[html body main section div div div ul li div div p]
[html body main section div div div ul li div div p strong]
[html body main section div div div ul li div div p strong]
[html body main section div div div ul li div div div]
[html body main section div div div ul li div div div span]
[html body main section div div div ul li div div div span]
[html body main section div div div ul li]
[html body main section div div div ul li div]
[html body main section div div div ul li div div]
[html body main section div div div ul li div div p]
[html body main section div div div ul li div div p strong]
[html body main section div div div ul li div div p strong]
[html body main section div div div ul li div div div]
[html body main section div div div ul li div div div span]
[html body main section div div div ul li div div div span]
[html body main section div div div ul li]
[html body main section div div div ul li div]
[html body main section div div div ul li div div]
[html body main section div div div ul li div div p]
[html body main section div div div ul li div div p strong]
[html body main section div div div ul li div div div]
[html body main section div div div ul li div div div span]
[html body main section div div div ul li div div div span]
[html body main section div div div ul li]
[html body main section div div div ul li div]
[html body main section div div div ul li div div]
[html body main section div div div ul li div div p]
[html body main section div div div ul li div div p strong]
[html body main section div div div ul li div div div]
[html body main section div div div ul li div div div span]
[html body main section div div div ul li div div div span]
[html body main section div div div ul li]
[html body main section div div div ul li div]
[html body main section div div div ul li div div]
[html body main section div div div ul li div div p]
[html body main section div div div ul li div div div]
[html body main section div div div ul li div div div span]
[html body main section div div div ul li div div div span]
[html body main section div div button]
[html body main section div div button i]
[html body main section div div button]
[html body main section div div button i]
[html body main section]
[html body main section div]
[html body main section div div]
[html body main section div div h2]
[html body main section div div]
[html body main section div div div]
[html body main section div div textarea]
[html body main section div div]
[html body main section div div]
[html body main section div div pre]
[html body main section div div pre noscript]
[html body main section div div]
[html body main section div div select]
[html body main section div div select option]
[html body main section div div select option]
[html body main section div div select option]
[html body main section div div select option]
[html body main section div div select option]
[html body main section div div select option]
[html body main section div div select option]
[html body main section div div select option]
[html body main section div div div]
[html body main section div div div button]
[html body main section div div div div]
[html body main section div div div div button]
[html body main section div div div div a]
[html body main section]
[html body main section div]
[html body main section div div]
[html body main section div div h2]
[html body main section div div p]
[html body main section div ul]
[html body main section div ul li]
[html body main section div ul li div]
[html body main section div ul li div div]
[html body main section div ul li div div img]
[html body main section div ul li div div img]
[html body main section div ul li div div]
[html body main section div ul li div div h3]
[html body main section div ul li div div p]
[html body main section div ul li div]
[html body main section div ul li div div]
[html body main section div ul li div div div]
[html body main section div ul li div div div img]
[html body main section div ul li div div ul]
[html body main section div ul li div div ul li]
[html body main section div ul li div div ul li a]
[html body main section div ul li div div ul li]
[html body main section div ul li div div ul li a]
[html body main section div ul li div div ul li]
[html body main section div ul li div div ul li a]
[html body main section div ul li div div]
[html body main section div ul li div div a]
[html body main section div ul li div div a i]
[html body main section div ul li]
[html body main section div ul li div]
[html body main section div ul li div div]
[html body main section div ul li div div img]
[html body main section div ul li div div img]
[html body main section div ul li div div]
[html body main section div ul li div div h3]
[html body main section div ul li div div p]
[html body main section div ul li div]
[html body main section div ul li div div]
[html body main section div ul li div div div]
[html body main section div ul li div div div img]
[html body main section div ul li div div ul]
[html body main section div ul li div div ul li]
[html body main section div ul li div div ul li a]
[html body main section div ul li div div ul li]
[html body main section div ul li div div ul li a]
[html body main section div ul li div div ul li]
[html body main section div ul li div div ul li a]
[html body main section div ul li div div ul li]
[html body main section div ul li div div ul li a]
[html body main section div ul li div div ul li]
[html body main section div ul li div div ul li a]
[html body main section div ul li div div]
[html body main section div ul li div div a]
[html body main section div ul li div div a i]
[html body main section div ul li]
[html body main section div ul li div]
[html body main section div ul li div div]
[html body main section div ul li div div img]
[html body main section div ul li div div img]
[html body main section div ul li div div]
[html body main section div ul li div div h3]
[html body main section div ul li div div p]
[html body main section div ul li div]
[html body main section div ul li div div]
[html body main section div ul li div div div]
[html body main section div ul li div div div img]
[html body main section div ul li div div ul]
[html body main section div ul li div div ul li]
[html body main section div ul li div div ul li a]
[html body main section div ul li div div ul li]
[html body main section div ul li div div ul li a]
[html body main section div ul li div div ul li]
[html body main section div ul li div div ul li a]
[html body main section div ul li div div ul li]
[html body main section div ul li div div ul li a]
[html body main section div ul li div div ul li]
[html body main section div ul li div div ul li a]
[html body main section div ul li div div]
[html body main section div ul li div div a]
[html body main section div ul li div div a i]
[html body main section div ul li]
[html body main section div ul li div]
[html body main section div ul li div div]
[html body main section div ul li div div img]
[html body main section div ul li div div img]
[html body main section div ul li div div]
[html body main section div ul li div div h3]
[html body main section div ul li div div p]
[html body main section div ul li div]
[html body main section div ul li div div]
[html body main section div ul li div div div]
[html body main section div ul li div div div img]
[html body main section div ul li div div ul]
[html body main section div ul li div div ul li]
[html body main section div ul li div div ul li a]
[html body main section div ul li div div ul li]
[html body main section div ul li div div ul li a]
[html body main section div ul li div div ul li]
[html body main section div ul li div div ul li a]
[html body main section div ul li div div]
[html body main section div ul li div div a]
[html body main section div ul li div div a i]
[html body main section div ul li]
[html body main section div ul li div]
[html body main section div ul li div div]
[html body main section div ul li div div img]
[html body main section div ul li div div]
[html body main section div ul li div div a]
[html body main section div ul li div div a i]
[html body main section]
[html body main section div]
[html body main section div div]
[html body main section div div h2]
[html body main section div div p]
[html body main section div div div]
[html body main section div div div a]
[html body main section div div div a]
[html body main section div div]
[html body main section div div ul]
[html body main section div div ul li]
[html body main section div div ul li]
[html body main section div div ul li a]
[html body main section div div ul li div]
[html body main section div div ul li]
[html body main section div div ul li a]
[html body main section div div ul li div]
[html body main section div div ul li]
[html body main section div div ul li a]
[html body main section div div ul li div]
[html body main section div div ul li]
[html body main section div div ul li a]
[html body main section div div ul li div]
[html body main section div div ul]
[html body main section div div ul li]
[html body main section div div ul li]
[html body main section div div ul li a]
[html body main section div div ul li div]
[html body main section div div ul li]
[html body main section div div ul li a]
[html body main section div div ul li div]
[html body main section div div ul li]
[html body main section div div ul li a]
[html body main section div div ul li div]
[html body main section div div ul li]
[html body main section div div ul li a]
[html body main section div div ul li div]
[html body main script]
[html body footer]
[html body footer div]
[html body footer div div]
[html body footer div div div]
[html body footer div div div div]
[html body footer div div div div a]
[html body footer div div div div a]
[html body footer div div div div a]
[html body footer div div div div]
[html body footer div div div div a]
[html body footer div div div div a]
[html body footer div div div div a]
[html body footer div div div div a]
[html body footer div div div div a]
[html body footer div div div div]
[html body footer div div div div a]
[html body footer div div div div a]
[html body footer div div div div a]
[html body footer div div div div]
[html body footer div div div div a]
[html body footer div div div div a]
[html body footer div div div div a]
[html body footer div div div div a]
[html body footer div div div div a]
[html body footer div div div div a]
[html body footer div div div div a]
[html body footer div div div div]
[html body footer div div div div a]
[html body footer div div div div a]
[html body footer div div div div a]
[html body footer div div div div a]
[html body footer div div div div a]
[html body footer div div div div a]
[html body footer div div div div a]
[html body footer div]
[html body footer div]
[html body footer div div]
[html body footer div div div]
[html body footer div div div img]
[html body footer div div div ul]
[html body footer div div div ul li]
[html body footer div div div ul li a]
[html body footer div div div ul li]
[html body footer div div div ul li a]
[html body footer div div div ul li]
[html body footer div div div ul li a]
[html body footer div div div ul li]
[html body footer div div div ul li a]
[html body footer div div div ul li]
[html body footer div div div ul li button]
[html body footer div div div ul li button img]
[html body footer div div div ul li button img]
[html body footer div div div ul li button img]
[html body footer div div div a]
[html body footer div div div a img]
[html body footer script]
[html body footer script]
[html body footer script]
[html body footer script]
[html body footer script]
[html body footer script]
[html body footer script]
[html body footer script]
[html body section]
[html body section div]
[html body section div a]
[html body section div]
[html body section div button]
geek@lenovo:~/GolandProjects/The_Go_Programming_Language_Exercises/ch5$ 
```