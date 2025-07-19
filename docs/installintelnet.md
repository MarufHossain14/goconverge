To install Telnet on your system, follow these steps:

## Windows
1. Open Command Prompt as Administrator.
2. Run the following command:
```bash
dism /online /Enable-Feature /dism /online /Enable-Feature /FeatureName:TelnetClient
```
3. Once installed, you can use Telnet by typing:
```bash
telnet localhost 8888
```
## Linux
Open your terminal.
Install Telnet using your package manager:
Debian/Ubuntu:
```bash
sudo apt install telnet
```
CentOS/Red Hat:
```bash
sudo yum install telnet
```
Test Telnet:
```bash
telnet localhost 8888
```
## Mac
Open your terminal.
Install Telnet using Homebrew:
```bash
brew install telnet
```
Test Telnet:
```bash
telnet localhost 8888
```
