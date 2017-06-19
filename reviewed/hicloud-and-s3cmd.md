
# Ref

* [s3cmd](http://s3tools.org/s3cmd)
* [hinet s3 help](http://s3help.cloudbox.hinet.net/index.php/)


# Install

* [Windows](http://s3help.cloudbox.hinet.net/index.php/s3cmd-start)
* MacOS
  * `brew install s3cmd`
* [Ubuntu](http://s3help.cloudbox.hinet.net/index.php/2015-02-12-07-08-03)
  * `sudo apt-get install s3cmd`

# Setting

```bash
# set config file
# need Access Key and Secret Key
s3cmd --configure

# change amazonaws setting to hicloud setting
vi .s3cfg
```

```ini
host_base = s3.amazonaws.com
host_bucket = %(bucket)s.s3.amazonaws.com
signature_v2 = False
```

Change to :

```ini
host_base = s3.hicloud.net.tw
host_bucket = %(bucket)s.s3.hicloud.net.tw
signature_v2 = True
```

# Upload

```bash
s3cmd path/to/file.ipa s3://bucket/path/to/file.ipa -P
```
