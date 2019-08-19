module github.com/branthz/gotest/github/echo

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190618222545-ea8f1a30c443
	golang.org/x/net => github.com/golang/net v0.0.0-20190619014844-b5b0513f8c1b
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190620070143-6f217b454f45
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190619215442-4adf7a708c2d
)

require github.com/labstack/echo/v4 v4.1.0
