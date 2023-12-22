package event

import "github.com/Shopify/sarama"

type Offset string

const (
	OffsetOldest Offset = "OLDEST"
	OffsetNewest Offset = "NEWEST"
)

func (o Offset) SaramaOffset() int64 {
	switch o {
	case OffsetOldest:
		return sarama.OffsetOldest
	case OffsetNewest:
		return sarama.OffsetNewest
	default:
		return sarama.OffsetNewest
	}
}
