# tarcing

```md
+--------+                                                +--------+
|        |                  Start Span                    |        |
| Client | <---------------------------------------------- | Server |
|        |               (with client context)             |        |
+--------+                                                +--------+
    |                                                           |
    |             Send request with Span context                |
    | --------------------------------------------------------> |
    |                                                           |
    |             Receive request with Span context             |
    | <-------------------------------------------------------- |
    |                                                           |
    |                  Start Span on Server                     |
    |               (with request Span context)                 |
    |                                                           |
    |                                                           |
    |                                                           |
    | ---------------------- RPC processing -------------------- |
    |                                                           |
    |                                                           |
    |                                                           |
    |                                                           |
    |                  Add attributes to Span                    |
    |                     (e.g. response code)                   |
    |                                                           |
    |                                                           |
    |                                                           |
    |                  End Span on Server                         |
    |                                                           |
    |                                                           |
    |             Update client Span with response               |
    | <-------------------------------------------------------- |
    |                                                           |
    |                                                           |
    |                                                           |
    |                      Close connection                       |
    | --------------------------------------------------------> |
    |                                                           |
+--------+                                                +--------+
|        |                                                |        |
| Client |                                                | Server |
|        |                                                |        |
+--------+                                                +--------+

```