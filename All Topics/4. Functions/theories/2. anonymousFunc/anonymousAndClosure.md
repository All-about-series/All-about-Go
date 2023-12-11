Combined [code examples](../../codes/anonymousFn/anonymousAndClosures.go)  
## Anonymous func
A func without a name
## Closures
- a specific type of anonymous func
- child func inside a parent func can inherit the properties from parent func
- useful when a parent func variable needs to be used outside of the variable (can be done when returing from a closure)
- **`defer`** keyword: Programs often create temporary resources, like files or network connections, that need to be cleaned up. This cleanup has to happen, no matter how many exit points a function has, or whether a function completed successfully or not. In Go, the cleanup
code is attached to the function with the `defer` keyword.
    - `defer` is use only for functions, methods or closures
    - It is only executed when the surrounding function returns
    - It can be used on multiple `func` at a time, they run in (**LIFO** order, last registered defer runs first)
    - A common pattern in Go is for a function that allocates a resource to also return a closure that cleans up the resource.
  - Unlike Go, Java/Js/Python uses try/catch/finally block to clean up resources (downside of these resource-cleanup block is that it creates more indentations and makes code harder to read) 
     > In research described in a 2017 paper in Empirical Software Engineering, Vard Antinyan et al. discovered that “Of...
    eleven proposed code characteristics, only two markedly influence
    complexity growth: the `nesting depth` and the `lack of structure`.”

    