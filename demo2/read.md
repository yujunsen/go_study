main.go
	data
		data.json
	matchers
		rss.go
	search
		defult.go
		feed.go
		match.go
		search.go
	
main.go -> import matchers search
	search->defult.go->init->Register
	matchers->rss.go->init->Register
	main.go->init()->main()
	search.go->Run()->
		feed.go->RetrieveFeeds()->open(data.json)->return feeds
	for
		go func(matcher Matcher, feed *Feed) 
			Match(matcher, feed, searchTerm, results)
				Search
	go func() 
		close(results)
	Display

Register	matcher与string进行绑定

每个代码文件都属于一个包，而包名应该与代码文件所在的文件夹同名。
Go 语言提供了多种声明和初始化变量的方式。如果变量的值没有显式初始化，编译器会将变量初始化为零值。
使用指针可以在函数间或者 goroutine 间共享数据。
通过启动 goroutine 和使用通道完成并发和同步
Go 语言提供了内置函数来支持 Go 语言内部的数据结构。
标准库包含很多包，能做很多很有用的事情
使用 Go 接口可以编写通用的代码和框