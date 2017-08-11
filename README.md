# AppEngine (AE) Deps

No one likes lock-in. And no one likes slow tests.

This meta-package provides a series of functions and root APIs for working
with Google AppEngine services in a way that decouples your application from
the underlying AppEngine implementation.

The primary audience for these packages is anyone who is currently using
AppEngine, and would like to get away from it. One of the goals is to make it as
easy as possible to use this package as a drop in replacement (with some
find/replace) for the basic AppEngine implementations. Then, as time permits,
the abstracts/interfaces provided herein will allow your code to start depending
on standard golang abstractions OR a public interface that your application can
implement.

## Provided Packages

Currently, aedeps provides implementations/abstractions for:

* [TaskQueue](https://godoc.org/google.golang.org/appengine/taskqueue)
* [ContextFactory](https://cloud.google.com/appengine/docs/standard/go/reference#NewContext)
* [RequestIDGetter](https://cloud.google.com/appengine/docs/standard/go/reference#RequestID)
