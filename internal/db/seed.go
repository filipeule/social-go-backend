package db

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"social/internal/store"
)

var usernames = []string{
	"shadowbyte",
	"lunarfox",
	"codewizard",
	"neonpulse",
	"crypticwolf",
	"silentstorm",
	"pixelpirate",
	"ironclad",
	"midnightowl",
	"quantumdash",
	"stormrider",
	"darkmatter",
	"silverfang",
	"bytehunter",
	"embercore",
	"frostnova",
	"cyberhawk",
	"voidwalker",
	"rapidflare",
	"ghostcoder",
	"atomiczen",
	"wildcipher",
	"hexaspark",
	"bluecomet",
	"redvortex",
	"hypernova",
	"nightcrawler",
	"techphantom",
	"alphaecho",
	"betablaze",
	"gammawave",
	"deltaforge",
	"omegaedge",
	"ironpulse",
	"stormbyte",
	"lucidstrike",
	"zenithcore",
	"orbitflare",
	"shadowflux",
	"novahex",
	"rapidshadow",
	"pixelstorm",
	"darkzen",
	"silverbyte",
	"cryptoflare",
	"nebulacode",
	"voidspark",
	"emberstrike",
	"lunarshift",
	"codevoyager",
}

var titles = []string{
	"Getting Started with Go",
	"Understanding REST APIs",
	"Mastering SQL Basics",
	"Concurrency in Practice",
	"Clean Code Principles",
	"Debugging Like a Pro",
	"Intro to Docker",
	"Building Microservices",
	"Web Security Essentials",
	"Scaling Your Backend",
	"Writing Better Tests",
	"Async Programming 101",
	"Optimizing Database Queries",
	"Deploying with CI/CD",
	"Designing System Architecture",
	"Logging Best Practices",
	"Error Handling in Go",
	"Understanding Caching",
	"Working with JSON",
	"API Versioning Strategies",
}

var contents = []string{
	"Go is a powerful and efficient programming language designed for simplicity and performance. In this guide, you'll learn how to set up your environment and build your first application step by step.",
	"REST APIs are the backbone of modern web services. This article explains how they work, common HTTP methods, and best practices for designing clean and scalable APIs.",
	"Structured Query Language (SQL) is essential for working with relational databases. Here we cover basic queries, filtering, joins, and how to retrieve meaningful data efficiently.",
	"Concurrency allows your applications to handle multiple tasks at the same time. Learn how goroutines and channels make concurrent programming simpler in Go.",
	"Clean code is about readability, maintainability, and simplicity. Discover practical principles that help you write code your future self will thank you for.",
	"Debugging is a critical skill for every developer. This post shares tools, strategies, and mindset tips to help you identify and fix issues faster.",
	"Docker makes application deployment consistent and reliable. Learn how containers work and how to package your application for any environment.",
	"Microservices architecture helps teams scale independently. We explore its benefits, trade-offs, and how to design services that communicate effectively.",
	"Security should never be an afterthought. This article covers authentication, authorization, HTTPS, and other core practices to protect your applications.",
	"Scaling a backend system requires careful planning. Learn about horizontal scaling, load balancing, and database optimization techniques.",
	"Testing ensures your software works as expected. Discover different testing strategies, including unit tests, integration tests, and automation tips.",
	"Asynchronous programming improves responsiveness and performance. Understand how async workflows work and when to use them effectively.",
	"Slow queries can hurt performance significantly. This post explains indexing, query planning, and practical techniques to optimize your database.",
	"CI/CD pipelines automate testing and deployment. Learn how continuous integration and delivery streamline development workflows.",
	"System architecture defines how components interact. We discuss scalability, reliability, and maintainability considerations for modern systems.",
	"Logging helps you understand what your system is doing. Explore logging levels, structured logs, and monitoring integrations.",
	"Error handling is essential for robust applications. Learn how to handle, wrap, and propagate errors properly in Go.",
	"Caching reduces load and improves response times. Discover different caching strategies and when to apply them.",
	"JSON is widely used for data exchange. This article explains encoding, decoding, and best practices when working with JSON APIs.",
	"API versioning prevents breaking changes for clients. Learn strategies like URL versioning, headers, and backward compatibility techniques.",
}

var tags = []string{
	"go",
	"backend",
	"api",
	"rest",
	"database",
	"sql",
	"docker",
	"microservices",
	"security",
	"concurrency",
	"testing",
	"performance",
	"architecture",
	"devops",
	"cloud",
	"json",
	"caching",
	"logging",
	"ci-cd",
	"web-development",
}

var comments = []string{
	"Great article! This clarified a lot of doubts I had.",
	"I've been struggling with this topic for weeks. Thanks for the clear explanation!",
	"Very practical and straight to the point. Loved it.",
	"Could you write more about advanced use cases?",
	"This helped me refactor part of my project. Appreciate it!",
	"I disagree slightly, but I see your point. Interesting perspective.",
	"Exactly what I needed today. Thanks!",
	"Do you have any recommended resources to dive deeper?",
	"The examples were really helpful. Please keep posting!",
	"Nice overview! Looking forward to more content like this.",
}

func Seed(store store.Storage) {
	ctx := context.Background()

	users := generateUsers(100)
	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			log.Println("error creating user:", err)
			return
		}
	}

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("error creating post:", err)
			return
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("error creating post:", err)
			return
		}
	}

	log.Println("seeding completed")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := range num {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d@example.com", i),
			Password: "123",
		}
	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)

	for i := range num {
		user := users[rand.Intn(len(users))]

		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}

	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)
	for i := range num {
		cms[i] = &store.Comment{
			PostID: posts[rand.Intn(len(posts))].ID,
			UserID: users[rand.Intn(len(users))].ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}
	return cms
}