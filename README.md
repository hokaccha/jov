jov
===========

A CLI JSON viewer.

Installation
------------------

```
$ go get -u github.com/hokaccha/jov
```

Example
------------------

![Example capture](http://i.imgur.com/Bf8RmdA.png)

Usage
------------------

Example json file:

```
$ cat posts.json
{"status":200,"result":[{"id":1,"title":"foo","body":"But I must explain to you how all this mistaken idea of denouncing pleasure and praising pain was born and I will give you a complete account of the system.","created_at":"2011-04-22T13:33:48Z"},{"id":2,"title":"bar","body":"The European languages are members of the same family. Their separate existence is a myth. For science, music, sport, etc, Europe uses the same vocabulary.","created_at":"2012-04-22T13:33:48Z"},{"id":3,"title":"baz","body":"Nor again is there anyone who loves or pursues or desires to obtain pain of itself, because it is pain, but because occasionally circumstances occur in which toil and pain can procure him some great pleasure.","created_at":"2013-04-22T13:33:48Z"}]}
```

### show pretty

```
# From stdin
$ cat posts.json | jov

# From file
$ jov -f posts.json

{
  "result": [
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
    },
    {
      "body": "Nor again is there anyone who loves or pursues or desires to obtain pain of itself, because it is pain, but because occasionally circumstances occur in which toil and pain can procure him some great pleasure.",
      "created_at": "2013-04-22T13:33:48Z",
      "id": 3,
      "title": "baz"
    }
  ],
  "status": 200
}
```

### get field

```
$ cat posts.json | jov get result
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
  },
  {
    "body": "Nor again is there anyone who loves or pursues or desires to obtain pain of itself, because it is pain, but because occasionally circumstances occur in which toil and pain can procure him some great pleasure.",
    "created_at": "2013-04-22T13:33:48Z",
    "id": 3,
    "title": "baz"
  }
]
```

### select or reject fields

select:

```
$ cat posts.json \
  | jov get result \
  | jov select id title
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

reject:

```
$ cat posts.json \
  | jov get result \
  | jov reject id body
[
  {
    "created_at": "2011-04-22T13:33:48Z",
    "title": "foo"
  },
  {
    "created_at": "2012-04-22T13:33:48Z",
    "title": "bar"
  },
  {
    "created_at": "2013-04-22T13:33:48Z",
    "title": "baz"
  }
]
```

### head or tail

head:

```
$ cat posts.json \
  | jov get result \
  | jov select id title \
  | jov head 2
[
  {
    "id": 1,
    "title": "foo"
  },
  {
    "id": 2,
    "title": "bar"
  }
]
```

tail:

```
$ cat posts.json \
  | jov get result \
  | jov select id title \
  | jov tail 2
[
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

### truncate string

```
$ cat posts.json \
  | jov get result \
  | jov select title body \
  | jov cut 50
[
  {
    "body": "But I must explain to you how all this mistaken id...",
    "title": "foo"
  },
  {
    "body": "The European languages are members of the same fam...",
    "title": "bar"
  },
  {
    "body": "Nor again is there anyone who loves or pursues or ...",
    "title": "baz"
  }
]
```

Licence
------------------

MIT
