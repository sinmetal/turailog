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

type MetaData4 struct {
	KV               map[string]int
	LabelAnnotations []*LabelAnnotation
}

type LabelAnnotation struct {
	ID          string
	Description string
	Score       float64
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
			category := int(now.Unix() % 3)
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
	case 2:
		kv := make(map[string]int)
		for i := 0; i < 30; i++ {
			kv[uuid.NewString()] = rand.Int()
		}

		metadata["KV"] = MetaData4{
			KV:               kv,
			LabelAnnotations: RandomLabelAnnotations(30),
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

func RandomLabelAnnotations(count int) []*LabelAnnotation {
	m := make(map[string]*LabelAnnotation)
	for i := 0; i < count; i++ {
		la := RandomLabelAnnotation()
		m[la.ID] = la
	}
	var las []*LabelAnnotation
	for _, v := range m {
		las = append(las, v)
	}
	return las
}

func RandomLabelAnnotation() *LabelAnnotation {
	i := rand.Int() % 10
	switch i {
	case 0:
		return &LabelAnnotation{ID: "m001", Description: "dog"}
	case 1:
		return &LabelAnnotation{ID: "m002", Description: "cat"}
	case 2:
		return &LabelAnnotation{ID: "m003", Description: "bird"}
	case 3:
		return &LabelAnnotation{ID: "m004", Description: "fish"}
	case 4:
		return &LabelAnnotation{ID: "m005", Description: "shark"}
	case 5:
		return &LabelAnnotation{ID: "m006", Description: "phoenix"}
	default:
		return &LabelAnnotation{ID: uuid.NewString(), Description: "random"}
	}
}
