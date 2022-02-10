# go_logger [under development]
Self-made Golang logger base on original Golang log pkg for reuse in projects.

<summary><h2 style="display: inline-block">Table of Contents</h2></summary>
<ol>
  <li><a href="#overview">Overview</a></li>
  <li><a href="#usage">Usage</a></li>
  <li><a href="#tests">Tests</a></li>
  <li><a href="#contact">Contact</a></li>
</ol>

<br>

## Overview
This go logger package provides an easy way to write log to the log file and print log on the console at the same time.
Before using go_logger, you should set a level kind of like Debug, Info, Warning or Error to distinguish the situation. 
In your code, put the logger somewhere you want to record with level type.

The logger implements an interface Logger which have 5 methods to present log level information.

```
type Logger interface {
	Debug(v interface{})
	Info(v interface{})
	Warning(v interface{})
	Error(v interface{})
	Critical(v interface{})
}
```

## Usage
create 
```

```



## Tests
---




## Contact
---
email: adelberteng@gmail.com