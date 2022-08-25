package config

import (
	"os"
	"strconv"
)

type Config struct {
	CompanyId string
	Secret    string
	AgentId   int
}

func NewConfig() *Config {
	var agentId int
	agentIdStr := os.Getenv("AGENT_ID")
	if len(agentIdStr) != 0 {
		agentId, _ = strconv.Atoi(agentIdStr)
	}
	return &Config{
		CompanyId: os.Getenv("COMPANY_ID"),
		Secret:    os.Getenv("SECRET"),
		AgentId:   agentId,
	}
}
