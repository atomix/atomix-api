# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [atomix/primitive/metadata.proto](#atomix/primitive/metadata.proto)
    - [CreatePrimitiveRequest](#atomix.primitive.CreatePrimitiveRequest)
    - [CreatePrimitiveResponse](#atomix.primitive.CreatePrimitiveResponse)
    - [DeletePrimitiveRequest](#atomix.primitive.DeletePrimitiveRequest)
    - [DeletePrimitiveResponse](#atomix.primitive.DeletePrimitiveResponse)
    - [GetPrimitiveRequest](#atomix.primitive.GetPrimitiveRequest)
    - [GetPrimitiveResponse](#atomix.primitive.GetPrimitiveResponse)
    - [GetPrimitivesRequest](#atomix.primitive.GetPrimitivesRequest)
    - [GetPrimitivesResponse](#atomix.primitive.GetPrimitivesResponse)
    - [PrimitiveMetadata](#atomix.primitive.PrimitiveMetadata)
  
  
  
    - [PrimitiveService](#atomix.primitive.PrimitiveService)
  

- [atomix/primitive/primitive.proto](#atomix/primitive/primitive.proto)
    - [PrimitiveId](#atomix.primitive.PrimitiveId)
  
    - [PrimitiveType](#atomix.primitive.PrimitiveType)
  
  
  

- [Scalar Value Types](#scalar-value-types)



<a name="atomix/primitive/metadata.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## atomix/primitive/metadata.proto



<a name="atomix.primitive.CreatePrimitiveRequest"></a>

### CreatePrimitiveRequest
CreatePrimitiveRequest is a request to create a primitive


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| database | [atomix.database.DatabaseId](#atomix.database.DatabaseId) |  |  |
| primitive | [PrimitiveId](#atomix.primitive.PrimitiveId) |  |  |
| type | [PrimitiveType](#atomix.primitive.PrimitiveType) |  |  |






<a name="atomix.primitive.CreatePrimitiveResponse"></a>

### CreatePrimitiveResponse
CreatePrimitiveResponse is a response for creating a primitive


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| primitive | [PrimitiveMetadata](#atomix.primitive.PrimitiveMetadata) |  |  |






<a name="atomix.primitive.DeletePrimitiveRequest"></a>

### DeletePrimitiveRequest
DeletePrimitiveRequest is a request to delete a primitive


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| database | [atomix.database.DatabaseId](#atomix.database.DatabaseId) |  |  |
| primitive | [PrimitiveId](#atomix.primitive.PrimitiveId) |  |  |






<a name="atomix.primitive.DeletePrimitiveResponse"></a>

### DeletePrimitiveResponse
DeletePrimitiveResponse is a response for deleting a primitive


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| primitive | [PrimitiveMetadata](#atomix.primitive.PrimitiveMetadata) |  |  |






<a name="atomix.primitive.GetPrimitiveRequest"></a>

### GetPrimitiveRequest
GetPrimitiveRequest is a request for primitive metadata


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| database | [atomix.database.DatabaseId](#atomix.database.DatabaseId) |  |  |
| primitive | [PrimitiveId](#atomix.primitive.PrimitiveId) |  |  |






<a name="atomix.primitive.GetPrimitiveResponse"></a>

### GetPrimitiveResponse
GetPrimitiveResponse is a response containing primitive metadata


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| primitive | [PrimitiveMetadata](#atomix.primitive.PrimitiveMetadata) |  |  |






<a name="atomix.primitive.GetPrimitivesRequest"></a>

### GetPrimitivesRequest
GetPrimitivesRequest is a request for primitive metadata


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| database | [atomix.database.DatabaseId](#atomix.database.DatabaseId) |  |  |
| primitive | [PrimitiveId](#atomix.primitive.PrimitiveId) |  |  |
| type | [PrimitiveType](#atomix.primitive.PrimitiveType) |  |  |






<a name="atomix.primitive.GetPrimitivesResponse"></a>

### GetPrimitivesResponse
GetPrimitivesResponse is a response containing primitive metadata


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| primitives | [PrimitiveMetadata](#atomix.primitive.PrimitiveMetadata) | repeated |  |






<a name="atomix.primitive.PrimitiveMetadata"></a>

### PrimitiveMetadata
PrimitiveMetadata indicates the type and name of a primitive


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| database | [atomix.database.DatabaseId](#atomix.database.DatabaseId) |  |  |
| primitive | [PrimitiveId](#atomix.primitive.PrimitiveId) |  |  |
| type | [PrimitiveType](#atomix.primitive.PrimitiveType) |  |  |





 

 

 


<a name="atomix.primitive.PrimitiveService"></a>

### PrimitiveService
PrimitiveService is a service for providing partition/primitive metadata

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreatePrimitive | [CreatePrimitiveRequest](#atomix.primitive.CreatePrimitiveRequest) | [CreatePrimitiveResponse](#atomix.primitive.CreatePrimitiveResponse) | CreatePrimitive creates a new primitive |
| GetPrimitive | [GetPrimitiveRequest](#atomix.primitive.GetPrimitiveRequest) | [GetPrimitiveResponse](#atomix.primitive.GetPrimitiveResponse) | GetPrimitive returns a primitive in the system |
| GetPrimitives | [GetPrimitivesRequest](#atomix.primitive.GetPrimitivesRequest) | [GetPrimitivesResponse](#atomix.primitive.GetPrimitivesResponse) | GetPrimitives returns a list of primitives in the system |
| DeletePrimitive | [DeletePrimitiveRequest](#atomix.primitive.DeletePrimitiveRequest) | [DeletePrimitiveResponse](#atomix.primitive.DeletePrimitiveResponse) | DeletePrimitive deletes a primitive |

 



<a name="atomix/primitive/primitive.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## atomix/primitive/primitive.proto



<a name="atomix.primitive.PrimitiveId"></a>

### PrimitiveId
Namespaced primitive identifier


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| namespace | [string](#string) |  |  |





 


<a name="atomix.primitive.PrimitiveType"></a>

### PrimitiveType
PrimitiveType is a primitive type

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN | 0 |  |
| COUNTER | 1 |  |
| ELECTION | 2 |  |
| INDEXED_MAP | 3 |  |
| LEADER_LATCH | 4 |  |
| LIST | 5 |  |
| LOCK | 6 |  |
| LOG | 7 |  |
| MAP | 8 |  |
| SET | 9 |  |
| VALUE | 10 |  |


 

 

 



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

