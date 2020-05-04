# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [atomix/pb/pb.proto](#atomix/pb/pb.proto)
    - [JoinReplicaGroupRequest](#atomix.pb.JoinReplicaGroupRequest)
    - [JoinReplicaGroupResponse](#atomix.pb.JoinReplicaGroupResponse)
    - [Partition](#atomix.pb.Partition)
    - [PartitionId](#atomix.pb.PartitionId)
    - [Replica](#atomix.pb.Replica)
    - [ReplicaGroup](#atomix.pb.ReplicaGroup)
    - [ReplicaGroupId](#atomix.pb.ReplicaGroupId)
    - [ReplicaId](#atomix.pb.ReplicaId)
  
  
  
    - [ReplicaGroupService](#atomix.pb.ReplicaGroupService)
  

- [Scalar Value Types](#scalar-value-types)



<a name="atomix/pb/pb.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## atomix/pb/pb.proto



<a name="atomix.pb.JoinReplicaGroupRequest"></a>

### JoinReplicaGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| replica_id | [ReplicaId](#atomix.pb.ReplicaId) |  |  |
| group_id | [ReplicaGroupId](#atomix.pb.ReplicaGroupId) |  |  |
| partitions | [uint32](#uint32) |  |  |
| replication_factor | [uint32](#uint32) |  |  |






<a name="atomix.pb.JoinReplicaGroupResponse"></a>

### JoinReplicaGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group | [ReplicaGroup](#atomix.pb.ReplicaGroup) |  |  |






<a name="atomix.pb.Partition"></a>

### Partition
Partition


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [PartitionId](#atomix.pb.PartitionId) |  |  |
| term | [uint64](#uint64) |  |  |
| leader | [ReplicaId](#atomix.pb.ReplicaId) |  |  |
| replicas | [Replica](#atomix.pb.Replica) | repeated |  |






<a name="atomix.pb.PartitionId"></a>

### PartitionId
Partition identifier


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| namespace | [string](#string) |  |  |
| index | [uint32](#uint32) |  |  |






<a name="atomix.pb.Replica"></a>

### Replica
Replica is a partition replica


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [ReplicaId](#atomix.pb.ReplicaId) |  |  |
| host | [string](#string) |  |  |
| port | [int32](#int32) |  |  |






<a name="atomix.pb.ReplicaGroup"></a>

### ReplicaGroup
Replica group state


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [ReplicaGroupId](#atomix.pb.ReplicaGroupId) |  |  |
| partitions | [Partition](#atomix.pb.Partition) | repeated |  |






<a name="atomix.pb.ReplicaGroupId"></a>

### ReplicaGroupId
Replica group identifier


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| namespace | [string](#string) |  |  |






<a name="atomix.pb.ReplicaId"></a>

### ReplicaId
Replica identifier


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| namespace | [string](#string) |  |  |





 

 

 


<a name="atomix.pb.ReplicaGroupService"></a>

### ReplicaGroupService
Atomix replica group service

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| JoinReplicaGroup | [JoinReplicaGroupRequest](#atomix.pb.JoinReplicaGroupRequest) | [JoinReplicaGroupResponse](#atomix.pb.JoinReplicaGroupResponse) stream | Joins a member to a replica group |

 



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
