## GetIfaceBot

### Description
The bot is triggered every hour. Gets the network interface of the machine and if it has been changed, sends the changes to the chat

### Install
```bash
git clone <this repo> /opt/getifacebot
cd /opt/getifacebot
make build # Create ./dist/getifacebot
vim getifacebot.service # Change TOKEN and CHAT_ID
sudo cp getifacebot.service /etc/systemd/system/
sudo systemctl enable --now getifacebot.service
```