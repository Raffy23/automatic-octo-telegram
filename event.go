package main

import (
	keptnv2 "github.com/keptn/go-utils/pkg/lib/v0_2_0"
)

const AutomaticOctoTelegramTriggeredType = "sh.keptn.event.echo.triggered"
const AutomaticOctoTelegramStartedEventType = "sh.keptn.event.echo.started"
const AutomaticOctoTelegramFinishedEventType = "sh.keptn.event.echo.finished"

type AutomaticOctoTelegramTriggeredEventData struct {
	keptnv2.EventData
	Action string `json:"action,omitempty"`
}

type AutomaticOctoTelegramStartedEventData struct {
	keptnv2.EventData
}

type AutomaticOctoTelegramFinishedEventData struct {
	keptnv2.EventData
}
