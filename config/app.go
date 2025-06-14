package config

import "time"

func MatchDisplayLowDelta() time.Duration {
	return cfg.MatchDisplayLowDelta
}

func MatchDisplayHighDelta() time.Duration {
	return cfg.MatchDisplayHighDelta
}

func MinCashOutAmount() float64 {
	return cfg.MinCashOutAmount
}

func MaxPicks() int {
	return cfg.MaxPicks
}

func NameChangeInterval() time.Duration {
	return cfg.NameChangeInterval
}

func DefaultChatRoomChannel() string {
	return cfg.DefaultChatRoomChannel
}
