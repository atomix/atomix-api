# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [atomix/headers/headers.proto](#atomix/headers/headers.proto)
    - [RequestHeader](#atomix.headers.RequestHeader)
    - [ResponseHeader](#atomix.headers.ResponseHeader)
    - [StreamHeader](#atomix.headers.StreamHeader)
  
    - [ResponseStatus](#atomix.headers.ResponseStatus)
    - [ResponseType](#atomix.headers.ResponseType)
  
  
  

- [Scalar Value Types](#scalar-value-types)



<a name="atomix/headers/headers.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## atomix/headers/headers.proto



<a name="atomix.headers.RequestHeader"></a>

### RequestHeader



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| partition | [uint32](#uint32) |  |  |
| primitive | [atomix.primitive.PrimitiveId](#atomix.primitive.PrimitiveId) |  |  |
| session_id | [uint64](#uint64) |  |  |
| request_id | [uint64](#uint64) |  |  |
| index | [uint64](#uint64) |  |  |
| streams | [StreamHeader](#atomix.headers.StreamHeader) | repeated |  |






<a name="atomix.headers.ResponseHeader"></a>

### ResponseHeader



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| session_id | [uint64](#uint64) |  |  |
| stream_id | [uint64](#uint64) |  |  |
| response_id | [uint64](#uint64) |  |  |
| index | [uint64](#uint64) |  |  |
| leader | [string](#string) |  |  |
| status | [ResponseStatus](#atomix.headers.ResponseStatus) |  |  |
| type | [ResponseType](#atomix.headers.ResponseType) |  |  |
| message | [string](#string) |  |  |






<a name="atomix.headers.StreamHeader"></a>

### StreamHeader



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stream_id | [uint64](#uint64) |  |  |
| response_id | [uint64](#uint64) |  |  |





 


<a name="atomix.headers.ResponseStatus"></a>

### ResponseStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| OK | 0 |  |
| ERROR | 1 |  |
| NOT_LEADER | 2 |  |
| UNKNOWN | 3 |  |
| CANCELED | 4 |  |
| NOT_FOUND | 5 |  |
| ALREADY_EXISTS | 6 |  |
| UNAUTHORIZED | 7 |  |
| FORBIDDEN | 8 |  |
| CONFLICT | 9 |  |
| INVALID | 10 |  |
| UNAVAILABLE | 11 |  |
| NOT_SUPPORTED | 12 |  |
| TIMEOUT | 13 |  |
| INTERNAL | 14 |  |



<a name="atomix.headers.ResponseType"></a>

### ResponseType


| Name | Number | Description |
| ---- | ------ | ----------- |
| RESPONSE | 0 |  |
| OPEN_STREAM | 1 |  |
| CLOSE_STREAM | 2 |  |


 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

