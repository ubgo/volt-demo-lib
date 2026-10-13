# volt-demo-lib

The **single-library repo** demo for [volt](https://github.com/ubgo/voltkit): an importable package at the repo root, bare `vX.Y.Z` tags, releases with notes but no binaries — the artifact that matters is the module proxy entry, which `volt release` warms and verifies.

```go
import "github.com/ubgo/volt-demo-lib"
```

Released with `volt release . v0.1.0`. Doubles as a regression fixture for volt.
