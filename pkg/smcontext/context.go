package smcontext

import (
	"context"
	"time"
)

type keyType struct{}

var propertiesKey = keyType{}

type properties struct {
	name  string
	class string

	club string

	organiser string
	event     string
	date      string
}

func get(ctx context.Context) *properties {
	props, ok := ctx.Value(propertiesKey).(*properties)
	if !ok {
		props = &properties{
			club:  "Kubikenborgs USKF",
			date:  time.Now().Format(time.DateOnly),
			event: "Träning",
		}

		props.organiser = props.club
	}
	return props
}

func store(ctx context.Context, props *properties) context.Context {
	return context.WithValue(ctx, propertiesKey, props)
}

func SetNameAndClass(ctx context.Context, name, class string) context.Context {
	props := get(ctx)

	props.name = name
	props.class = class

	return store(ctx, props)
}

func Class(ctx context.Context) string {
	return get(ctx).class
}

func Club(ctx context.Context) string {
	return get(ctx).club
}

func Date(ctx context.Context) string {
	return get(ctx).date
}

func Event(ctx context.Context) string {
	return get(ctx).event
}

func Name(ctx context.Context) string {
	return get(ctx).name
}

func Organiser(ctx context.Context) string {
	return get(ctx).organiser
}
