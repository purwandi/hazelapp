package storage

import (
	"time"

	"github.com/purwandi/hazelapp/project/domain"
)

type ProjectStorage struct {
	ProjectMap []domain.Project
}

func NewProjectStorage() *ProjectStorage {
	return &ProjectStorage{
		ProjectMap: []domain.Project{},
	}
}

func (storage *ProjectStorage) Demo() error {
	storage.ProjectMap = []domain.Project{
		domain.Project{ID: 1, OwnerID: 1, Name: "project-website", Description: "Add new features", CreatedAt: time.Now()},
		domain.Project{ID: 2, OwnerID: 1, Name: "marketing-website", Description: "Implement project management for linear issues", CreatedAt: time.Now()},
		domain.Project{ID: 3, OwnerID: 1, Name: "mobile-app", Description: "Robus filtering on all views", CreatedAt: time.Now()},
		domain.Project{ID: 4, OwnerID: 1, Name: "cycles", Description: "Initial version of cycles", CreatedAt: time.Now()},
		domain.Project{ID: 5, OwnerID: 1, Name: "react", Description: "A declarative, efficient, and flexible JavaScript library for building user interfaces.", CreatedAt: time.Now()},
		domain.Project{ID: 6, OwnerID: 1, Name: "react-native", Description: "A framework for building native apps with React.", CreatedAt: time.Now()},
		domain.Project{ID: 7, OwnerID: 1, Name: "litho", Description: "A declarative framework for building efficient UIs on Android.", CreatedAt: time.Now()},
		domain.Project{ID: 8, OwnerID: 1, Name: "idb", Description: "idb is a flexible command line interface for automating iOS simulators and devices", CreatedAt: time.Now()},
		domain.Project{ID: 9, OwnerID: 1, Name: "rocksdb", Description: "A library that provides an embeddable, persistent key-value store for fast storage.", CreatedAt: time.Now()},
		domain.Project{ID: 10, OwnerID: 1, Name: "componentkit", Description: "A React-inspired view framework for iOS.", CreatedAt: time.Now()},
		domain.Project{ID: 11, OwnerID: 1, Name: "flipper", Description: "A desktop debugging platform for mobile developers.", CreatedAt: time.Now()},
		domain.Project{ID: 12, OwnerID: 1, Name: "buck", Description: "A fast build system that encourages the creation of small, reusable modules over a variety of platforms and languages.", CreatedAt: time.Now()},
		domain.Project{ID: 13, OwnerID: 1, Name: "pyre-check", Description: "Performant type-checking for python.", CreatedAt: time.Now()},
		domain.Project{ID: 14, OwnerID: 1, Name: "fbt", Description: "A JavaScript Internationalization Framework.", CreatedAt: time.Now()},
		domain.Project{ID: 15, OwnerID: 1, Name: "facebook-ios-sdk", Description: "Used to integrate the Facebook Platform with your iOS & tvOS apps.", CreatedAt: time.Now()},
		domain.Project{ID: 16, OwnerID: 1, Name: "proxygen", Description: "A collection of C++ HTTP libraries including an easy to use HTTP server.", CreatedAt: time.Now()},

		domain.Project{ID: 17, OwnerID: 2, Name: "ax", Description: "Adaptive Experimentation Platform.", CreatedAt: time.Now()},
		domain.Project{ID: 18, OwnerID: 2, Name: "infer", Description: "A static analyzer for Java, C, C++, and Objective-C", CreatedAt: time.Now()},
		domain.Project{ID: 19, OwnerID: 2, Name: "fbshipit", Description: "Copy commits between repositories - git => git, git => hg, hg => hg, or hg => git", CreatedAt: time.Now()},
		domain.Project{ID: 20, OwnerID: 2, Name: "relay", Description: "Relay is a JavaScript framework for building data-driven React applications.", CreatedAt: time.Now()},
		domain.Project{ID: 21, OwnerID: 2, Name: "fboss", Description: "Facebook Open Switching System Software for controlling network switches.", CreatedAt: time.Now()},
		domain.Project{ID: 22, OwnerID: 2, Name: "fbzmq", Description: "Facebook ZeroMQ wrappers.", CreatedAt: time.Now()},
		domain.Project{ID: 23, OwnerID: 2, Name: "wangle", Description: "Wangle is a framework providing a set of common client/server abstractions for building services in a consistent, modular, and composable way.", CreatedAt: time.Now()},
		domain.Project{ID: 24, OwnerID: 2, Name: "openr", Description: "Distributed platform for building autonomic network functions.", CreatedAt: time.Now()},
		domain.Project{ID: 25, OwnerID: 2, Name: "fbmeshd", Description: "fbmeshd is the core technology behind Facebook's Self-organizing Mesh Access (SoMA) network.", CreatedAt: time.Now()},
		domain.Project{ID: 26, OwnerID: 2, Name: "fbthrift", Description: "A React-inspired view framework for iOS.", CreatedAt: time.Now()},
		domain.Project{ID: 27, OwnerID: 2, Name: "bistro", Description: "Bistro is a flexible distributed scheduler, a high-performance framework supporting multiple paradigms while retaining ease of configuration, management, and monitoring.", CreatedAt: time.Now()},
		domain.Project{ID: 28, OwnerID: 2, Name: "watchman", Description: "Watches files and records, or triggers actions, when they change.", CreatedAt: time.Now()},
		domain.Project{ID: 29, OwnerID: 2, Name: "fb303", Description: "fb303 is a core set of thrift functions that provide a common mechanism for querying stats and other information from a service.", CreatedAt: time.Now()},
		domain.Project{ID: 30, OwnerID: 2, Name: "folly", Description: "An open-source C++ library developed and used at Facebook.", CreatedAt: time.Now()},

		domain.Project{ID: 31, OwnerID: 3, Name: "flow", Description: "Adds static typing to JavaScript to improve developer productivity and code quality.", CreatedAt: time.Now()},
		domain.Project{ID: 32, OwnerID: 3, Name: "hhvm", Description: "A virtual machine for executing programs written in Hack.", CreatedAt: time.Now()},
	}

	return nil
}
