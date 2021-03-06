# Godig (Golang)

![Godig](https://siasky.net/IAAJhHVXd-SkfrdvpshHWJCOSPiNAuTgtpynfdQpgrEEFA)
Domain resolver (DNS) -> Like DIG

This script is equivalent to :
`dig a domain.com +short "@IP" && dig mx domain.com +short "@IP" && dig txt domain.com +short && dig _dmarc.domain.com TXT +short && dig selector._domainkey.domain.com TXT +short`

## How to use
```
GODIG - Domain DNS Resolver in Golang
        Usage:   godig [domain] [selector | @IPresolver]

        Example: godig domain.com
                 godig domain.com mailjet
                 godig domain.com @8.8.8.8
                 godig domain.com @8.8.8.8
		 godig domain.com google @8.8.8.8
                 
Use [godig help] to show this message
```

You can build this program
- Linux
```bash
go build -o godig godig.go
#and execute the built program
./godig domain.com
```

**You can make this script easier**
```bash
sudo mv godig /usr/bin/godig
#Just lanch the script
godig domain.com
```

- Windows
```bash
env GOOS=windows GOARCH=amd64 go build -o godig.exe godig.go
.\godig.exe domain.com
```


You can now launch this program (without build)
```bash
go run godig.go domain.com
```

## Output

```bash
godig domain.com mail

DNS (A):
216.58.206.238
2a00:1450:4007:816::200e

MX fields:
10 aspmx.l.google.com
20 alt1.aspmx.l.google.com
30 alt2.aspmx.l.google.com
40 alt3.aspmx.l.google.com
50 alt4.aspmx.l.google.com

TXT records:
globalsign-smime-dv=CDYX+XFHUw2wml6/Gb8+59BsH31KzUr6c1l2BPvqKX8=
docusign=05958488-4752-4ef2-95eb-aa7ba8a3bd0e
v=spf1 include:_spf.google.com ~all
docusign=1b0a6754-49b1-4db5-8540-d2c12664b289
facebook-domain-verification=22rm551cu4k0ab0bxsw536tlds4h95

DMARC key:
v=DMARC1; p=reject; rua=mailto:mailauth-reports@google.com

DKIM key:
No DKIM key found
```

## Docker
- Build
```bash
docker build -t godig .
```
- Run
```bash
docker run --rm godig domain.com
```
- Run without building - Download from Docker Hub
```bash
docker run --rm deedwark/godig domain.com
```

## Demo
[![asciicast](https://asciinema.org/a/2Sc2uQqRosGsC97IbrjbFqsvE.svg)](https://asciinema.org/a/2Sc2uQqRosGsC97IbrjbFqsvE)

## Me
[LinkedIn](https://fr.linkedin.com/in/kenji-duriez-9b93bb141)
