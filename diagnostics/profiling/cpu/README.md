# Golang CPU Profiling

## Overview

My learning note on how CPU profiling works and how to use it.

***

## CPU Profiling Internals
Video: https://www.youtube.com/watch?v=MbSTTrwDbQ explaining how CPU profiling works behind the scene. (I delivered in Bahasa Indonesia, but the slide content is in English)

[![Video](http://img.youtube.com/vi/MbSTTrwDbQc/0.jpg)](https://www.youtube.com/watch?v=MbSTTrwDbQc "Video")  

References:
- Go profiler internal
https://www.instana.com/blog/go-profiler-internals/

- Settimer Linux manual page
https://man7.org/linux/man-pages/man2/setitimer.2.html

- How To Build a User-Level CPU Profiler
https://research.swtch.com/pprof

- Proposal: hardware performance counters for CPU profiling
https://go.googlesource.com/proposal/+/refs/changes/08/219508/2/design/36821-perf-counter-pprof.md

- Linux System Calls
https://www.informit.com/articles/article.aspx?p=23618&seqNum=14

- User CPU time vs System CPU time?
https://stackoverflow.com/questions/4310039/user-cpu-time-vs-system-cpu-time

- Circular buffer
https://en.wikipedia.org/wiki/Circular_buffer

- Profiler labels in Go 
https://rakyll.org/profiler-labels/

- Program Counter
https://www.techopedia.com/definition/13114/program-counter-pc

- IPC
https://youtu.be/dJuYKfR8vec

- Signals
https://youtu.be/m0tUe0PgGCs  

## CPU Profiling Implementation
Video: https://www.youtube.com/watch?v=sn4_wctJJ0A explaining how CPU profiling can be used in Go. (I delivered in Bahasa Indonesia, but the slide content is in English)

[![Video](http://img.youtube.com/vi/sn4_wctJJ0A/0.jpg)](https://www.youtube.com/watch?v=sn4_wctJJ0A "Video")  

References:

- Remote Profiling of go Programs 
  https://www.farsightsecurity.com/blog/txt-record/go-remote-profiling-20161028/
  
- Go: The Complete Guide to Profiling Your Code | Hacker Noon
  https://hackernoon.com/go-the-complete-guide-to-profiling-your-code-h51r3waz
  
-  pprof - The Go Programming Language (golang.org)
  https://golang.org/pkg/net/http/pprof/
  
- CVE-2019-11248: /debug/pprof exposed on kubelet's healthz port
  https://github.com/kubernetes/kubernetes/issues/81023
