# linkedlist
<a href="https://coveralls.io/github/badges/shields">
    <img src="https://img.shields.io/coveralls/github/badges/shields" alt="coverage">
</a>

Library of generic singly, doubly, and circularly linked-list data structures for Go.

* [singly linked-list](./singly/)
* [doubly linked-list](./doubly/)
* [circularly linked-list](./circularly/)


## Install
```
$ go get github.com/golang-ds/linkedlist
```

## Singly LinkedList Usage
#### Import
```
import "github.com/golang-ds/linkedlist/singly"
```
#### Use
```
list := singly.New[int]()
list.AddFirst(1)
```

## Doubly LinkedList Usage
#### Import
```
import "github.com/golang-ds/linkedlist/doubly"
```
#### Use
```
list := doubly.New[int]()
list.AddFirst(1)
```

## Circularly LinkedList Usage
#### Import
```
import "github.com/golang-ds/linkedlist/circularly"
```
#### Use
```
list := circularly.New[int]()
list.AddFirst(1)
```


## Todo
* OfRange factory function
* OfSlice factory function
* Init(range, func) factory function
