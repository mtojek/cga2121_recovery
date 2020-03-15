# cga2121_recovery

Recover password from Technicolor CGA2121 using brute-force method.

## Usage

```bash
$ cga2121_recovery --help
Usage of cga2121_recovery:
  -hostPort string
    	host and port (default "192.168.0.1")
  -numWorkers int
    	number of workers (default 10)
  -passwordMinLength int
    	minimal password length (default 1)
  -passwordPattern string
    	character patterns (default "aA1_")
  -user string
    	user (default "admin")
```

Sample execution:

```bash
$ cga2121_recovery
2020/03/15 20:18:15 Check password: a
2020/03/15 20:18:15 Check password: b
2020/03/15 20:18:15 Check password: d
2020/03/15 20:18:15 Check password: e
2020/03/15 20:18:15 Check password: f
2020/03/15 20:18:15 Check password: g
2020/03/15 20:18:15 Check password: h
2020/03/15 20:18:15 Check password: i
2020/03/15 20:18:15 Check password: j
2020/03/15 20:18:15 Check password: c
...
```
