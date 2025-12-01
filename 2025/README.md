# aoc-2025

I am practicing C++17 this year to be better in the fundamental syntax. I also included a Makefile to automate the days and I added vscode config to automate lldb (clang debugger) on vscode.

```
> make help
  make run-sample [DAY=N]  - Build and run day N with sample.txt
  make fmt              - Format all days with clang-format
  make clean [DAY=N]       - Clean day N
  make clean-all           - Clean all days

Examples:
  make DAY=5               - Build day 5
  make run DAY=5           - Run day 5 with input.txt
  make run-sample DAY=3    - Run day 3 with sample.txt
```

**disclaimer: I don't love C++**