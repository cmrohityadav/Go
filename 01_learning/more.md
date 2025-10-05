| Type             | Zero-value      | Nil?  | make() effect                             |
| ---------------- | --------------- | ----- | ----------------------------------------- |
| int, float, bool | 0, 0.0, false   | ❌ no  | N/A                                       |
| string           | "" (empty)      | ❌ no  | N/A                                       |
| pointer `*T`     | nil             | ✅ yes | `new(T)` → alloc pointer, non-nil         |
| slice `[]T`      | nil             | ✅ yes | `make([]T,n)` → non-nil, len=n, cap≥n     |
| map `map[K]V`    | nil             | ✅ yes | `make(map[K]V)` → non-nil, ready-to-use   |
| channel `chan T` | nil             | ✅ yes | `make(chan T, n)` → non-nil, ready-to-use |
| interface        | nil             | ✅ yes | N/A                                       |
| struct           | all fields zero | ❌ no  | N/A                                       |




---

VALUE TYPE (int, string, bool)
+--------+
| x = 0  |  <- zero-value
+--------+
Stack memory

POINTER / REFERENCE TYPE (slice, map, channel)
Nil slice / map / channel
Stack header:
+---------------------+
| pointer = nil       |  <- not allocated
| length=0 / cap=0    |
+---------------------+
Heap: none allocated

Make slice/map/channel
Stack header:
+---------------------+
| pointer ------------> Heap allocated
| length >=0          |
| capacity >=0        |
+---------------------+
Heap (slice example):
+---+---+---+---+---+
| 0 | 0 | 0 | 0 | 0 |
+---+---+---+---+---+



---

