package webdriver

import (
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/dsmontoya/webdriver/api"
)

type ChromeOption func(*Chrome)

var DefaultChromeOptions = []ChromeOption{
	ChromePort(9515),
	ChromeURLBase(""),
	ChromeHTTPThreads(4),
	ChromeLogPath("./chrome.log"),
	ChromeStartTimeout(20 * time.Second),
}

type Chrome struct {
	URL *url.URL
	//The port that Chrome listens on. Default: 9515
	port int64
	//The URL path prefix to use for all incoming WebDriver REST requests. Default: ""
	urlBase string
	//The number of threads to use for handling HTTP requests. Default: 4
	httpThreads int64
	//The path to use for the Chrome server log. Default: ./chromedriver.log
	logPath string
	// Log file to dump chromedriver stdout/stderr. If "" send to terminal. Default: ""
	LogFile string
	// Start method fails if Chromedriver doesn't start in less than startTimeout. Default 20s.
	startTimeout time.Duration

	cmd        *exec.Cmd
	logFile    *os.File
	path       string
	stdoutPipe io.ReadCloser
	stderrPipe io.ReadCloser
}

func NewAPIClient() *api.APIClient {
	cfg := &api.Configuration{
		DefaultHeader: make(map[string]string),
		UserAgent:     "OpenAPI-Generator/1.0.0/go",
		Debug:         false,
		Servers: api.ServerConfigurations{
			{
				URL:         "{protocol}://{host}:{port}{path}",
				Description: "Selenium WebDriver API",
				Variables: map[string]api.ServerVariable{
					"protocol": {
						Description:  "No description provided",
						DefaultValue: "http",
						EnumValues: []string{
							"http",
							"https",
						},
					},
					"host": {
						Description:  "Default host name",
						DefaultValue: "localhost",
					},
					"port": {
						Description:  "Default port number",
						DefaultValue: "9515",
						EnumValues: []string{
							"9515",
							"4445",
						},
					},
					"path": {
						Description:  "Selenium API base path",
						DefaultValue: "",
						EnumValues: []string{
							"",
						},
					},
				},
			},
		},
		OperationServers: map[string]api.ServerConfigurations{},
	}

	return api.NewAPIClient(cfg)
}

// NewChrome returns a new Chrome. It sets the corresponding DefaultChromeOptions option for
// each option not provided.
func NewChrome(path string, options ...ChromeOption) *Chrome {
	chrome := &Chrome{path: path}
	for _, option := range append(DefaultChromeOptions, options...) {
		option(chrome)
	}
	return chrome
}

var switchesFormat = "-port=%d -url-base=%s -log-path=%s -http-threads=%d"

var cmdchan = make(chan error)

func (c *Chrome) Start() error {
	if c.cmd != nil {
		return nil
	}

	u, err := url.Parse(fmt.Sprintf("http://127.0.0.1:%d%s", c.port, c.urlBase))
	if err != nil {
		return errors.New("wrong url base")
	}
	c.URL = u

	if err := c.setupCommand(); err != nil {
		return err
	}

	if err := c.cmd.Start(); err != nil {
		return fmt.Errorf("failed to start command: %v", err)
	}
	if c.LogFile != "" {
		flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
		var err error
		c.logFile, err = os.OpenFile(c.LogFile, flags, 0640)
		if err != nil {
			return err
		}
		go io.Copy(c.logFile, c.stderrPipe)
		go io.Copy(c.logFile, c.stderrPipe)
	} else {
		go io.Copy(os.Stdout, c.stdoutPipe)
		go io.Copy(os.Stderr, c.stderrPipe)
	}
	if err := probePort(int(c.port), c.startTimeout); err != nil {
		return err
	}
	return nil
}

func (d *Chrome) Stop() error {
	if d.cmd == nil {
		return errors.New("stop failed: chromedriver not running")
	}
	defer func() {
		d.cmd = nil
	}()
	//TODO: use cmd.Wait
	d.cmd.Process.Signal(os.Interrupt)
	if d.logFile != nil {
		d.logFile.Close()
	}
	return nil
}

func (c *Chrome) setupCommand() error {
	var arguments []string
	arguments = append(arguments, "-port="+strconv.Itoa(int(c.port)))
	arguments = append(arguments, "-log-path="+c.logPath)
	arguments = append(arguments, "-http-threads="+strconv.Itoa(int(c.httpThreads)))
	if c.urlBase != "" {
		arguments = append(arguments, "-url-base="+c.urlBase)
	}
	c.cmd = exec.Command(c.path, arguments...)

	stdoutPipe, err := c.cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to connect to stdout: %v", err)
	}
	c.stdoutPipe = stdoutPipe

	stderrPipe, err := c.cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("failed to connect to stderr: %v", err)
	}
	c.stderrPipe = stderrPipe

	return nil
}

func ChromePort(port int64) ChromeOption {
	return func(chrome *Chrome) {
		chrome.port = port
	}
}

func ChromeURLBase(urlBase string) ChromeOption {
	return func(chrome *Chrome) {
		chrome.urlBase = urlBase
	}
}

func ChromeHTTPThreads(httpThreads int64) ChromeOption {
	return func(chrome *Chrome) {
		chrome.httpThreads = httpThreads
	}
}

func ChromeLogPath(logPath string) ChromeOption {
	return func(chrome *Chrome) {
		chrome.logPath = logPath
	}
}

func ChromeStartTimeout(startTimeout time.Duration) ChromeOption {
	return func(chrome *Chrome) {
		chrome.startTimeout = startTimeout
	}
}
