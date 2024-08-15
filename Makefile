all:
	go build -o overcv

install: all
	systemctl --user stop overland-receiver || true
	cp overcv ~/.local/bin
	cp overland-receiver.service ~/.config/systemd/user
	systemctl --user daemon-reload
	systemctl --user enable --now overland-receiver
