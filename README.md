# linkedlist

Library of generic singly, doubly, and circularly linked-list data structures for Go.

* [singly linked-list](./singly/)
* [doubly linked-list](./doubly/)
* [circularly linked-list](./circularly/)

## Install

```Go
$ go get github.com/golang-ds/linkedlist
```

## Singly LinkedList Usage

### Import

```Go
import "github.com/golang-ds/linkedlist/singly"
```

### Use

```Go
list := singly.New[int]()
list.AddFirst(1)
```

## Doubly LinkedList Usage

### Import

```Go
import "github.com/golang-ds/linkedlist/doubly"
```

### Use

```Go
list := doubly.New[int]()
list.AddFirst(1)
```

## Circularly LinkedList Usage

### Import

```Go
import "github.com/golang-ds/linkedlist/circularly"
```

### Use

```Go
list := circularly.New[int]()
list.AddFirst(1)
```

## Todo

* OfRange factory function
* OfSlice factory function
* Init(range, func) factory function
