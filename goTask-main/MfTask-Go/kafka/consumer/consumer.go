package consumer

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"keycloak-demo/database"
	"keycloak-demo/model"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
	"gorm.io/gorm"
)

var (
	Topic         string
	ConsumerGroup string
	// ConsumeMsg    = make(chan *kgo.Record, 30)
	// repo database.HOLDINGSDB
)

func ConsumeTopic(db *gorm.DB) {

	repo := database.NewHoldingsDB(db)

	flag.StringVar(&Topic, "topic", "ordersv1", "ordersv1")
	flag.Parse()

	seeds := []string{"kafka1", "kafka2", "kafka3"}

	cl, err := kgo.NewClient(
		kgo.SeedBrokers(seeds...),
		kgo.ConsumerGroup(ConsumerGroup),
		kgo.ConsumeTopics(Topic),
	)
	if err != nil {
		panic(err)
	}
	defer cl.Close()

	ctx := context.Background()
	time.Sleep(time.Second * 5)
	for {
		fetches := cl.PollFetches(ctx)
		if errs := fetches.Errors(); len(errs) > 0 {

			panic(fmt.Sprint(errs))
		}

		// We can iterate through a record iterator...
		iter := fetches.RecordIter()
		for !iter.Done() {
			record := iter.Next()
			// ConsumeMsg <- record
			obj := new(model.ORDER)
			err := json.Unmarshal(record.Value, obj)
			if err != nil {
				return
			}

			if obj.Status == "confirmed" {

				fmt.Println("Here its adding the holdings")
				holdObj := new(model.HOLDINGS)
				holdObj.SchemeCode = obj.Scheme_code
				holdObj.UserId = obj.UserId

				repo.AddHoldings(holdObj)

			}

			fmt.Println("Partition-->", record.Partition, "Topic-->", record.Topic, string(record.Value), "from an iterator!")
		}

	}

}
