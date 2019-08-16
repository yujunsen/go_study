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