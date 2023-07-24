package kookvoice

func Play(token string, channelId string, input string) {
	gatewayUrl := GetGatewayUrl(token, channelId)
	connect, rtpUrl := InitWebsocketClient(gatewayUrl)
	defer connect.Close()
	go KeepWebsocketClientAlive(connect)
	go KeepRecieveMessage(connect)
	StreamAudio(rtpUrl, input)
}

func New(token string, channelId string) (*voiceInstance, error) {
	vi := voiceInstance{
		Token:     token,
		ChannelId: channelId,
	}
	err := vi.Init()
	if err != nil {
		return nil, err
	}
	return &vi, nil
}
