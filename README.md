# easygo

Some general golang code

```
import "github.com/powerpuffpenguin/easygo"
```

Other Module list:

* [bytes](bytes/README.md)
* [option](option/README.md)

# Pair

Pair is a struct with two properties

```
easygo.Pair[int, string]{1, "ok"}

easygo.MakePair(1, "ok")

easygo.NewPair(1, "ok")
```

# Tuple

Tuple is similar to Pair, but it allows 1 to many properties (Tuple1 Tuple2 ... Tuple30)

```
easygo.Tuple2[int, string]{1, "ok"}

easygo.MakeTuple2(1, "ok")

easygo.NewTuple2(1, "ok")
```
