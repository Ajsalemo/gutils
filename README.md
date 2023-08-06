# gutils

```
                  _     _   _
   __ _   _   _  | |_  (_) | |  ___
  / _` | | | | | | __| | | | | / __|
 | (_| | | |_| | | |_  | | | | \__ \
  \__, |  \__,_|  \__| |_| |_| |___/
  |___/
```

Rewrite of a personal git command line helper in Go.

This combines `git add .`, `git commit -m`, and `git push` into one executable, with the option to specify the commit message, remote, and branch.

Usage:
-  -b (string)
    - The branch to push to - defaults to main (default "main")
-  -c (string)
    - A commit message - defaults to 'initial commit' (default "initial commit")
-  -r (string)
    - The specified upstream git remote to push to - defaults to origin (default "origin")


