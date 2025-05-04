// MayTheFourth will turn your webpage into a Jedi koan

package main

import (
	"context"
)

type MayTheFourth struct{}

func (m *MayTheFourth) Enlighten(ctx context.Context, url string) (string, error) {
	summary, err := dag.TechSummarizerAgent().Summarize(ctx, url)
	if err != nil {
		return "", err
	}
	environment := dag.Env().
		WithStringInput("summary", summary, "short page summary").
		WithStringOutput("koan", "a short Jedi koan based on $summary")
	return dag.LLM().
		WithEnv(environment).
		WithPrompt("You're a Jedi master, and a Zen master. Write accordingly.").
		Loop().
		Env().
		Output("koan").
		AsString(ctx)
}
