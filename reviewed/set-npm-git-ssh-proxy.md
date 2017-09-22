# Ref.

* [ssh in git behind proxy on windows 7](http://stackoverflow.com/questions/5103083/ssh-in-git-behind-proxy-on-windows-7)
* [NPM behind a proxy](https://github.com/npm/npm/issues/9401)
* [如何在強制使用代理伺服器的環境下設定 git, npm, bower, gem, ionic 工具](http://blog.miniasp.com/post/2015/09/02/proxy-settings-for-git-npm-bower-gem-ionic.aspx)
* [Angular 2 學習筆記 00 - 緣起與開發環境建置](https://dotblogs.com.tw/topcat/2016/12/19/153702)
* [多重 SSH Keys 與 Github 帳號](https://kuanyui.github.io/2016/08/01/git-multiple-ssh-key/)

# Hint

* `http://USERNAME:PASSWORD@IP:PORT` 可以改為 `http://IP:PORT`。
<!-- * For Allianz: `http://wsg1.allianz.corp:8080` -->

# Setting

<!-- TOC depthFrom:2 -->

- [NPM](#npm)
- [GIT](#git)
- [SSH](#ssh)
- [VS CODE](#vs-code)
- [NVM](#nvm)

<!-- /TOC -->

## NPM

```sh
npm config set http-proxy http://USERNAME:PASSWORD@IP:PORT
npm config set https-proxy http://USERNAME:PASSWORD@IP:PORT
npm set strict-ssl false
```

in `C:\Users\USERNAME\.npmrc`:

```
proxy=http://USERNAME:PASSWORD@IP:PORT
http-proxy=http://USERNAME:PASSWORD@IP:PORT
https-proxy=http://USERNAME:PASSWORD@IP:PORT
strict-ssl=false

```

## GIT


```sh
git config --global http.proxy http://USERNAME:PASSWORD@IP:PORT
git config --global https.proxy http://USERNAME:PASSWORD@IP:PORT
```

in `C:\Users\USERNAME\.gitconfig`:

```
[http]
	proxy = http://USERNAME:PASSWORD@IP:PORT
[https]
	proxy = http://USERNAME:PASSWORD@IP:PORT
```

## SSH

in `C:\Users\USERNAME\.ssh\config`:

```
ProxyCommand connect -H IP:PORT %h %p

Host github.com
  User git
  Port 22
  Hostname github.com
  IdentityFile "C:\users\USERNAME\.ssh\id_rsa"
  TCPKeepAlive yes
  IdentitiesOnly yes

Host ssh.github.com
  User git
  Port 443
  Hostname ssh.github.com
  IdentityFile "C:\users\USERNAME\.ssh\id_rsa"
  TCPKeepAlive yes
  IdentitiesOnly yes

Host bitbucket.org
  User git
  Port 22
  Hostname bitbucket.org
  IdentityFile "C:\users\USERNAME\.ssh\id_rsa"
  TCPKeepAlive yes
  IdentitiesOnly yes
```

## VS CODE

in `setting.json`:

```json
{
  "http.proxy": "http://USERNAME:PASSWORD@IP:PORT",
}
```


## NVM

* [nvm github](https://github.com/creationix/nvm)
* [nvm-window github](https://github.com/coreybutler/nvm-windows)

```sh
# set proxy
nvm proxy http://USERNAME:PASSWORD@IP:PORT

# show a list of versions available for download.
nvm list available

# install node
nvm install NODE_VERSION

# set use node version
nvm use NODE_VERSION
```
