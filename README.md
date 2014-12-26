jov
===========

A CLI JSON viewer.

Installation
------------------

```
$ go get -u github.com/hokaccha/jov
```

Usage
------------------

Example json file:

```
$ cat foo.json
[{"id":1,"title":"foo","body":"But I must explain to you how all this mistaken idea of denouncing pleasure and praising pain was born and I will give you a complete account of the system.","created_at":"2011-04-22T13:33:48Z"},{"id":2,"title":"bar","body":"The European languages are members of the same family. Their separate existence is a myth. For science, music, sport, etc, Europe uses the same vocabulary.","created_at":"2012-04-22T13:33:48Z"},{"id":3,"title":"baz","body":"Nor again is there anyone who loves or pursues or desires to obtain pain of itself, because it is pain, but because occasionally circumstances occur in which toil and pain can procure him some great pleasure.","created_at":"2013-04-22T13:33:48Z"}]
```

### show pretty

```
# From stdin
$ cat foo.json | jov

# From file
$ jov -f foo.json
```

[image]

### with command

```
$ cat foo.json | jov select id title
[
  {
    "id": 1,
    "title": "foo"
  },
  {
    "id": 2,
    "title": "bar"
  },
  {
    "id": 3,
    "title": "baz"
  }
]
```

```
$ cat foo.json | jov head 2
[
  {
    "body": "But I must explain to you how all this mistaken idea of denouncing pleasure and praising pain was born and I will give you a complete account of the system.",
    "created_at": "2011-04-22T13:33:48Z",
    "id": 1,
    "title": "foo"
  },
  {
    "body": "The European languages are members of the same family. Their separate existence is a myth. For science, music, sport, etc, Europe uses the same vocabulary.",
    "created_at": "2012-04-22T13:33:48Z",
    "id": 2,
    "title": "bar"
  }
]
```

### truncate string

```
$ cat foo.json | jov -t 50
[
  {
    "body": "But I must explain to you how all this mistaken id...",
    "created_at": "2011-04-22T13:33:48Z",
    "id": 1,
    "title": "foo"
  },
  {
    "body": "The European languages are members of the same fam...",
    "created_at": "2012-04-22T13:33:48Z",
    "id": 2,
    "title": "bar"
  },
  {
    "body": "Nor again is there anyone who loves or pursues or ...",
    "created_at": "2013-04-22T13:33:48Z",
    "id": 3,
    "title": "baz"
  }
]
```

### with pipe

```
$ cat foo.json \
    | jov head 2 \
    | jov select title body \
    | jov -t 50
[
  {
    "body": "But I must explain to you how all this mistaken id...",
    "title": "foo"
  },
  {
    "body": "The European languages are members of the same fam...",
    "title": "bar"
  }
]
```

Commands
------------------

### get

```
$ jov get <key>
```

Retrieve the value of a object.

Input:

```json
{
  "status": 200,
  "results": [
    { "key1": "val1" },
    { "key2": "val2" },
    { "key3": "val3" }
  ]
}
```

Output:

```
$ cat input.json | jov get results
[
  {
    "key1": "val1"
  },
  {
    "key2": "val2"
  },
  {
    "key3": "val3"
  }
]
```

### select

```
$ jov select <property>...
```

Select properties of a collection (collection means `[Object, Object, ...]`)

Input:

```json
[
  {
    "id": 1,
    "title": "foo",
    "created_at": "2011-04-22T13:33:48Z"
  },
  {
    "id": 2,
    "title": "bar",
    "created_at": "2012-04-22T13:33:48Z"
  },
  {
    "id": 3,
    "title": "baz",
    "created_at": "2013-04-22T13:33:48Z"
  }
]
```

Output:

```
$ cat input.json | jov select id title
[
  {
    "id": 1,
    "title": "foo"
  },
  {
    "id": 2,
    "title": "bar"
  },
  {
    "id": 3,
    "title": "baz"
  }
]
```

### reject

```
$ jov reject <property>...
```

Reject properties of a collection.

Input:

```json
[
  {
    "id": 1,
    "title": "foo",
    "created_at": "2011-04-22T13:33:48Z"
  },
  {
    "id": 2,
    "title": "bar",
    "created_at": "2012-04-22T13:33:48Z"
  },
  {
    "id": 3,
    "title": "baz",
    "created_at": "2013-04-22T13:33:48Z"
  }
]
```

Output:

```
$ cat input.json | jov reject title
[
  {
    "created_at": "2011-04-22T13:33:48Z",
    "id": 1
  },
  {
    "created_at": "2012-04-22T13:33:48Z",
    "id": 2
  },
  {
    "created_at": "2013-04-22T13:33:48Z",
    "id": 3
  }
]
```

### head

```
$ jov head <length>
```

Return the first `<length>` elements of a array.

Input:

```json
[
  { "id": 1 },
  { "id": 2 },
  { "id": 3 },
  { "id": 4 },
  { "id": 5 },
  { "id": 6 },
  { "id": 7 },
  { "id": 8 },
  { "id": 9 },
  { "id": 10 }
]
```

Output:

```
$ cat input.json | jov head 3
[
  {
    "id": 1
  },
  {
    "id": 2
  },
  {
    "id": 3
  }
]
```

### tail

```
$ jov tail <length>
```

Return the last `<length>` elements of a array.

Input:

```json
[
  { "id": 1 },
  { "id": 2 },
  { "id": 3 },
  { "id": 4 },
  { "id": 5 },
  { "id": 6 },
  { "id": 7 },
  { "id": 8 },
  { "id": 9 },
  { "id": 10 }
]
```

Output:

```
$ cat input.json | jov tail 3
[
  {
    "id": 8
  },
  {
    "id": 9
  },
  {
    "id": 10
  }
]
```

### slice

Return a array starting at the `<start>` index and continuing for `<length>` elements of a array.

Input:

```json
[
  { "id": 1 },
  { "id": 2 },
  { "id": 3 },
  { "id": 4 },
  { "id": 5 },
  { "id": 6 },
  { "id": 7 },
  { "id": 8 },
  { "id": 9 },
  { "id": 10 }
]
```

Output:

```
$ cat input2.json | jov slice 5 3
[
  {
    "id": 6
  },
  {
    "id": 7
  },
  {
    "id": 8
  }
]
```

Licence
------------------

MIT
