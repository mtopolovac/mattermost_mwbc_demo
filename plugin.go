package main

import (
	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/plugin"
)

// HelloWorldPlugin implements the interface expected by the Mattermost server to communicate
// between the server and plugin processes.
type MessageWillBeConsumedDemoPlugin struct {
	plugin.MattermostPlugin
}

func (p *MessageWillBeConsumedDemoPlugin) MessageWillBeConsumed(post *model.Post) *model.Post {
	post.Message = "wbc_plugin:" + post.Message
	return post
}

// This example demonstrates a plugin that handles HTTP requests which respond by greeting the
// world.
func main() {
	plugin.ClientMain(&MessageWillBeConsumedDemoPlugin{})
}
