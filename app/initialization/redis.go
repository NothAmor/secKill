package initialization

import (
	"fmt"
	"secKill/app/common"

	"github.com/redis/rueidis"
	"github.com/redis/rueidis/rueidisaside"
	"github.com/redis/rueidis/rueidiscompat"
)

func initRedis() {
	redisConfig := common.Config.Redis
	client, err := rueidisaside.NewClient(rueidisaside.ClientOption{
		ClientOption: rueidis.ClientOption{InitAddress: []string{fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port)}},
	})
	if err != nil {
		panic(err)
	}

	common.Redis = rueidiscompat.NewAdapter(client.Client())
}
