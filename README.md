

```shell
make init
```






## Scalar Value Types 

float, uint32

string


---

- optional
- repeated
- map


Maps

Map fields are just a shorthand for a special kind of repeated field.
If we have

```protobuf
message Test6 {
  map<string, int32> g = 7;
}
```

this is actually the same as

```protobuf
message Test6 {
  message g_Entry {
    optional string key = 1;
    optional int32 value = 2;
  }
  repeated g_Entry g = 7;
}
```

## Nested Types 

message
