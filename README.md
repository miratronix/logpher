# logpher [![Documentation](https://godoc.org/github.com/miratronix/logpher?status.svg)](http://godoc.org/github.com/miratronix/logpher)
logpher is a basic [autumn](https://github.com/miratronix/autumn)-enabled golang logger package with various writers.

## Configuration
Logpher is built around the concept of named loggers. Each logger has its own level, which can be specified via 
configuration. Additionally, Logpher supports 3 writers out of the box:
- A combination writer
- A console writer
- A file writer
- A rolling file writer

All of these settings are controlled via a configuration object:
```go
config := &logpher.Configuration{
    Type:       "console",          // This can be "combination", "console", "file", or "rolling"
    Combine:    "console,rolling"   // The writers to combine when using the "combination" type
    File:       "./mylog.txt",      // The name of the file to log to when the type is "file" or "rolling"        
    Size:       8,                  // The maximum log file size in MB when the type is "rolling"
    Count:      5,                  // The number of files to keep when the type is "rolling"
    Levels: map[string]string{      // The levels to use for the various loggers
    	"default": "info",          // The default log level for new loggers
    	"main": "debug",            // Overrides the log level for the "main" logger
    }
}
```

## Standard Usage
Standard usage is as simple as initializing Logpher and creating a logger:
```go
// Initialize Logpher
l := logpher.New(config)

// Create a logger
mainLogger := l.NewLogger("main")

// Log something
mainLogger.Trace("something")
mainLogger.Debug("something")
mainLogger.Info("something")
mainLogger.Warn("something")
mainLogger.Error("something")

// Var args will be concatenated with spaces
mainLogger.Debug("something", "happened")

// Close open files
l.Close()
```

## Autumn Usage
Logpher is designed to work nicely with Autumn:
```go
// Initialize the autumn tree
tree := autumn.NewTree()

// Add the configuration object
tree.AddNamedLeaf("logConfiguration", config)

// Add the main logpher leaf
tree.AddLeaf(&logpher.Logpher{})

// Add loggers
tree.AddLeaf(logpher.NewLogger("main"))

// The logger will be wired into the following structure automatically
type main struct {
	Logger *logpher.Logger `autumn:"mainLogger"`
}

// The log file will be closed when the tree is chopped
tree.Chop()
```
