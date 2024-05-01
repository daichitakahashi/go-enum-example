package usecase

import (
	"context"
	"crypto/rand"
	"log"
	"math/big"
	"time"

	"github.com/daichitakahashi/go-enum"
)

type AppCentral struct{}

//go:generate go run github.com/daichitakahashi/go-enum/cmd/enumgen@latest --out=results.gen.go --visitor="*:*Handler:On*" --accept="*:Handle" --visitor-impl="*"

type (
	CheckHealthStatusResult interface {
		enum.VisitorReturns[error]
		CheckHealthStatusResultEnum
	}

	CheckHealthStatusHealthy struct {
		enum.MemberOf[CheckHealthStatusResult]

		StartedAt  time.Time
		FinishedAt time.Time
	}
	CheckHealthStatusUnhealthy struct {
		enum.MemberOf[CheckHealthStatusResult]

		StartedAt  time.Time
		FinishedAt time.Time
		Cause      string
	}
)

func (a *AppCentral) CheckHealthStatus(ctx context.Context) (CheckHealthStatusResult, error) {
	startedAt := time.Now()

	time.Sleep(time.Second)
	r, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		return nil, err
	}

	if r.Int64()%2 == 0 {
		return CheckHealthStatusHealthy{
			StartedAt:  startedAt,
			FinishedAt: time.Now(),
		}, nil
	}
	return CheckHealthStatusUnhealthy{
		StartedAt:  startedAt,
		FinishedAt: time.Now(),
		Cause:      "Today is cloudy",
	}, nil
}

type (
	GreetingResult interface {
		enum.VisitorReturns[error]
		GreetingResultEnum
	}

	GreetingHello struct {
		enum.MemberOf[GreetingResult]

		Message string
	}
	GreetingAbsent struct {
		enum.MemberOf[GreetingResult]
	}
)

func (a *AppCentral) Greeting(ctx context.Context, handler GreetingResultHandler) error {
	r, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		return err
	}

	if r.Int64()%2 == 0 {
		err := GreetingHello{
			Message: "Hi!",
		}.Handle(handler)
		if err != nil {
			return err
		}
	} else {
		err := GreetingAbsent{}.Handle(handler)
		if err != nil {
			return err
		}
	}

	log.Println("Greeting responded")
	return nil
}
