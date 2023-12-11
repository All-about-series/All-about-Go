## Todos

- learn `makefile` from https://makefiletutorial.com
- read about `UTF-8` (from _learning go_ or _internet_)
- change the folder structure accordingly

## References

> ### Books-
>
> 1. Learning Go An Idiomatic Approach to Real-world Go Programming, 2nd Edition (Jon Bodner)  
>    with [Exercises solution](<./Exercises%20Sol/Learning%20Go%20An%20Idiomatic%20Approach%20to%20Real-world%20Go%20Programming%20(2nd%20ed)/>)
> 2. 

<hr/>

### Folder structure
```go
All-about-Go/
├── All Topics/
│   ├── 1. <topic>/
│   │   ├── codes/
│   │   │   ├── <sub-topic>/          // package
│   │   │   │   ├── <file>.go
│   │   │   │   └── ...
│   │   │   ├── .../
│   │   │   │   ├── ...
│   │   │   │   └── ...
│   │   │   └── ...
│   │   │   ├── Makefile
│   │   │   ├── main.go
│   │   │   └── go.mod                // module name
│   │   ├── theories/
│   │   │   ├── <sub-topic>/
│   │   │   │   ├── <file>.md
│   │   │   │   ├── ...
│   │   │   │   └── README.md
│   │   │   ├── .../
│   │   │   │   ├── ...
│   │   │   │   └── README.md
│   │   │   └── ...
│   │   └── README.md
│   ├── 2. <topic>/
│   │   ├── codes/
│   │   │   ├── ...
│   │   │   └── ...                   
│   │   ├── theories/
│   │   │   ├── ...
│   │   │   └── ...
│   │   └── README.md
│   └── ...
├── Exercises Sol/
│   ├── <book/tutorial>/
│   │   ├── <chapter/topic>/
│   │   │   ├── <sub-topic>/     
│   │   │   │   └── sol.go         
│   │   │   ├── .../     
│   │   │   │   └── sol.go         
│   │   │   ├── Makefile           
│   │   │   ├── main.go            
│   │   │   └── go.mod             
│   │   ├── ...                     
│   └── ...
└── README.md
```