# bliz

A very simple cli app for getting and setting the key value.

Under the hook, it's writing and reading to the files. The files are partitioned by the remainder of the decimal of the first character of the key divided by 10. So, there will be only 10 partitions.

```
NAME:
   bliz - A new cli application

USAGE:
   bliz [global options] command [command options] [arguments...]

COMMANDS:
   get      get the value by key eg. get key
   set      set the value by key eg. set key value
   list     list all the keys, might be deleted in the future
   help, h  Shows a list of commands or help for one command
```

to be continued 😉
