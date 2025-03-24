With goroutine synchronization disabled all the list-related anf cache directory-related operations immediately become unreliable and buggy and in most cases result in Undefined Behaviour. The reason for this is that list and directory alteration processes are not immediate and same cache list and directory is used by all of the goroutines simultaneously. Therefore, for example, a goroutine may access a directory to check some cache line status after cache list was updated by another routine but before that routine had a chance to update the cache directory. This means that the information in the directory accessed by the first goroutine is simply invalid and using it in most cases would lead to errors. Many other buggy scenarious may take place at least because even list and directory opertaions are not atomic and can be "interrupted" by other routines, reading or modifying same data structure instances at the same time and in a conflicting way.

The most common error message generated after launching the 'fail' version of the code looks like this:

```bash
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x10 pc=0x4cc220]

goroutine 9 [running]:
container/list.(*List).Remove(...)
        /usr/lib/go-1.23/src/container/list/list.go:135
github.com/quaiion/go-practice/lru-cache/internal/CacheDir.(*CacheDir).processMiss(0xc00008c4c0, 0x12)
        /home/quaiion/Documents/code/go-practice/korney.ivanishin/task-4/internal/CacheDir/CacheDir.go:84 +0x1a0
github.com/quaiion/go-practice/lru-cache/internal/CacheDir.(*CacheDirSync).GetRequest(0x7e2cb0?, 0x12)
        /home/quaiion/Documents/code/go-practice/korney.ivanishin/task-4/internal/CacheDir/CacheDir.go:58 +0x145
github.com/quaiion/go-practice/lru-cache/internal/Requester.Requester.Request({0x0?}, 0xc00008c4c0)
        /home/quaiion/Documents/code/go-practice/korney.ivanishin/task-4/internal/Requester/Requester.go:23 +0x55
github.com/quaiion/go-practice/lru-cache/internal/Requester.Requester.RequestN({0x0?}, 0xc00008c4c0, 0x64)
        /home/quaiion/Documents/code/go-practice/korney.ivanishin/task-4/internal/Requester/Requester.go:37 +0x71
main.goroutRequestN({0x0?}, 0x0?, 0x0?, 0xc000028460)
        /home/quaiion/Documents/code/go-practice/korney.ivanishin/task-4/cmd/service/main.go:46 +0x18
main.launchRequesters.func1()
        /home/quaiion/Documents/code/go-practice/korney.ivanishin/task-4/cmd/service/main.go:68 +0x38
golang.org/x/sync/errgroup.(*Group).Go.func1()
        /home/quaiion/go/pkg/mod/golang.org/x/sync@v0.12.0/errgroup/errgroup.go:78 +0x50
created by golang.org/x/sync/errgroup.(*Group).Go in goroutine 1
        /home/quaiion/go/pkg/mod/golang.org/x/sync@v0.12.0/errgroup/errgroup.go:75 +0x96
```

As you see, the asynchronous use of cache list and cache directory in most cases simply leads to dereferencing a nil pointer. Boring? Yeah, but still hazardous. Be a good developer. Don't forget about synchronization.
