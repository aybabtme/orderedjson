# `orderedjson`

```go
import "github.com/aybabtme/orderedjson"
```

Sometimes you decode a JSON object in Go, and despite the JSON spec technically requiring that keys be unordered... well, someone somewhere is relying on the order of the keys. Example: you have a JSON job file specifying commands to run, and these commands need to run in the order specified in the JSON config file... for legacy reasons.

If you just use the Go `encoding/json` package to decode into a `map[string]...`, order will be lost since map iteration is random. There's no official way of maintaining this order, so this package offers a type that does that, and only that. It doesn't decode any further, it stricly gives you access to the JSON payload, keys and values, in the order presented. You can then further continue to unmarshal the keys and values as you wish.

# usage


Given this input:
```json
{
  "0": null,
  "1": 0,
  "2": "s",
  "3": [null, 0, "string", [], {}],
  "4": {"0": null, "1": 0, "2": "s", "3": [], "4": {}},
}
```

Use the `orderedjson.Map` type to unmarshal:

```go

var object orderedjson.Map
err := json.Unmarshal(input, &object)
```

The content of `object` will then be:

```go
object := orderedjson.Map{
  {Key: json.RawMessage(`"0"`), Value: json.RawMessage(`null`)},
  {Key: json.RawMessage(`"1"`), Value: json.RawMessage(`0`)},
  {Key: json.RawMessage(`"2"`), Value: json.RawMessage(`"s"`)},
  {Key: json.RawMessage(`"3"`), Value: json.RawMessage(`[null, 0, "string", [], {}]`)},
  {Key: json.RawMessage(`"4"`), Value: json.RawMessage(`{"0": null, "1": 0, "2": "s", "3": [], "4": {}}`)},
}
```

And you can continue unmarshalling `Key` and `Value` however you wish.

# license

MIT.