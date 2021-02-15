package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/google/uuid"
)

type TuraiJson struct {
	ID         int64
	Metadata   map[string]interface{}
	AppVersion string
}

type MetaData1 struct {
	ID    string
	Kind  string
	Value int64
}

type MetaData2 struct {
	ID      string
	Content string
	Tags    []string
}

type MetaData3 struct {
	ID             string
	Value          string
	MetaData3Inner *MetaData3Inner
}

type MetaData3Inner struct {
	Hello string
}

type TuraiJsonOutputter struct {
	CurrentID   int64
	SleepSecond int64
}

func (outputter *TuraiJsonOutputter) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("done:TuraiJsonOutputter.Run()", ctx.Err())
			return
		default:
			now := time.Now()
			category := now.Day() % 3
			outputter.Log(outputter.GenerateTurai(category))
			outputter.CurrentID++
			time.Sleep(time.Duration(outputter.SleepSecond) * time.Second)
			outputter.SleepSecond += 1 + outputter.SleepSecond
		}
	}
}

func (outputter *TuraiJsonOutputter) Log(value *TuraiJson) {
	if err := json.NewEncoder(os.Stdout).Encode(value); err != nil {
		fmt.Println(err.Error())
	}
}

func (outputter *TuraiJsonOutputter) GenerateTurai(category int) *TuraiJson {
	metadata := make(map[string]interface{})

	switch category {
	case 1:
		metadata["hoge"] = MetaData3{
			ID:    uuid.NewString(),
			Value: "Hello Hoge",
			MetaData3Inner: &MetaData3Inner{
				Hello: "World",
			},
		}
	default:
		metadata["hoge"] = MetaData1{
			ID:    uuid.NewString(),
			Kind:  "TuraiKimoti",
			Value: rand.Int63(),
		}
		metadata["fuga"] = MetaData2{
			ID:      uuid.NewString(),
			Content: uuid.NewString(),
			Tags:    []string{"AppEngine", "ComputeEngine", "CloudSQL", "CloudStorage", "BigQuery"},
		}
	}

	return &TuraiJson{
		ID:         outputter.CurrentID,
		Metadata:   metadata,
		AppVersion: "v0.0.0",
	}
}
