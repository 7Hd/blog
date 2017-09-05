# Purpose

統一換行符號

# `.gitattributes`

```
* text eol=lf
```

# Git Config

```sh
# Check setting
git config core.autocrlf

# Mac & Linux
git config --global core.autocrlf input
# Windows
git config --global core.autocrlf false
git config --global core.safecrlf true
```

# Git Reset

```sh
# Remove every file from the git index
git rm --cached -r .

# Rewrite the git index
git reset --hard
```

# Ref

1. https://stackoverflow.com/questions/36727469/bitbucket-crlf-issue
2. https://blog.miniasp.com/post/2013/09/15/Git-for-Windows-Line-Ending-Conversion-Notes.aspx
3. https://help.github.com/articles/dealing-with-line-endings/
4. https://github.com/cssmagic/blog/issues/22
