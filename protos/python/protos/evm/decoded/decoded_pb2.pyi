from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Access(_message.Message):
    __slots__ = ["address", "storage_keys"]
    ADDRESS_FIELD_NUMBER: _ClassVar[int]
    STORAGE_KEYS_FIELD_NUMBER: _ClassVar[int]
    address: str
    storage_keys: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, address: _Optional[str] = ..., storage_keys: _Optional[_Iterable[str]] = ...) -> None: ...

class Block(_message.Message):
    __slots__ = ["base_fee_per_gas", "difficulty", "extra_data", "gas_limit", "gas_used", "hash", "logs_bloom", "miner", "mix_hash", "nonce", "number", "parent_hash", "receipt_root", "sha3_uncles", "size", "state_root", "timestamp", "total_difficulty", "transactions", "transactions_root", "uncles"]
    BASE_FEE_PER_GAS_FIELD_NUMBER: _ClassVar[int]
    DIFFICULTY_FIELD_NUMBER: _ClassVar[int]
    EXTRA_DATA_FIELD_NUMBER: _ClassVar[int]
    GAS_LIMIT_FIELD_NUMBER: _ClassVar[int]
    GAS_USED_FIELD_NUMBER: _ClassVar[int]
    HASH_FIELD_NUMBER: _ClassVar[int]
    LOGS_BLOOM_FIELD_NUMBER: _ClassVar[int]
    MINER_FIELD_NUMBER: _ClassVar[int]
    MIX_HASH_FIELD_NUMBER: _ClassVar[int]
    NONCE_FIELD_NUMBER: _ClassVar[int]
    NUMBER_FIELD_NUMBER: _ClassVar[int]
    PARENT_HASH_FIELD_NUMBER: _ClassVar[int]
    RECEIPT_ROOT_FIELD_NUMBER: _ClassVar[int]
    SHA3_UNCLES_FIELD_NUMBER: _ClassVar[int]
    SIZE_FIELD_NUMBER: _ClassVar[int]
    STATE_ROOT_FIELD_NUMBER: _ClassVar[int]
    TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
    TOTAL_DIFFICULTY_FIELD_NUMBER: _ClassVar[int]
    TRANSACTIONS_FIELD_NUMBER: _ClassVar[int]
    TRANSACTIONS_ROOT_FIELD_NUMBER: _ClassVar[int]
    UNCLES_FIELD_NUMBER: _ClassVar[int]
    base_fee_per_gas: str
    difficulty: str
    extra_data: str
    gas_limit: str
    gas_used: str
    hash: str
    logs_bloom: str
    miner: str
    mix_hash: str
    nonce: str
    number: int
    parent_hash: str
    receipt_root: str
    sha3_uncles: str
    size: int
    state_root: str
    timestamp: _timestamp_pb2.Timestamp
    total_difficulty: int
    transactions: _containers.RepeatedCompositeFieldContainer[Transaction]
    transactions_root: str
    uncles: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, base_fee_per_gas: _Optional[str] = ..., difficulty: _Optional[str] = ..., extra_data: _Optional[str] = ..., gas_limit: _Optional[str] = ..., gas_used: _Optional[str] = ..., hash: _Optional[str] = ..., logs_bloom: _Optional[str] = ..., miner: _Optional[str] = ..., mix_hash: _Optional[str] = ..., nonce: _Optional[str] = ..., number: _Optional[int] = ..., parent_hash: _Optional[str] = ..., receipt_root: _Optional[str] = ..., sha3_uncles: _Optional[str] = ..., state_root: _Optional[str] = ..., timestamp: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., transactions: _Optional[_Iterable[_Union[Transaction, _Mapping]]] = ..., transactions_root: _Optional[str] = ..., uncles: _Optional[_Iterable[str]] = ..., total_difficulty: _Optional[int] = ..., size: _Optional[int] = ...) -> None: ...

class CallTrace(_message.Message):
    __slots__ = ["calls", "gas", "gas_used", "input", "output", "to", "type", "value"]
    CALLS_FIELD_NUMBER: _ClassVar[int]
    FROM_FIELD_NUMBER: _ClassVar[int]
    GAS_FIELD_NUMBER: _ClassVar[int]
    GAS_USED_FIELD_NUMBER: _ClassVar[int]
    INPUT_FIELD_NUMBER: _ClassVar[int]
    OUTPUT_FIELD_NUMBER: _ClassVar[int]
    TO_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    VALUE_FIELD_NUMBER: _ClassVar[int]
    calls: _containers.RepeatedCompositeFieldContainer[CallTrace]
    gas: str
    gas_used: str
    input: str
    output: str
    to: str
    type: str
    value: str
    def __init__(self, type: _Optional[str] = ..., to: _Optional[str] = ..., value: _Optional[str] = ..., gas: _Optional[str] = ..., gas_used: _Optional[str] = ..., input: _Optional[str] = ..., output: _Optional[str] = ..., calls: _Optional[_Iterable[_Union[CallTrace, _Mapping]]] = ..., **kwargs) -> None: ...

class Data(_message.Message):
    __slots__ = ["block", "call_trace", "transaction_receipt"]
    BLOCK_FIELD_NUMBER: _ClassVar[int]
    CALL_TRACE_FIELD_NUMBER: _ClassVar[int]
    TRANSACTION_RECEIPT_FIELD_NUMBER: _ClassVar[int]
    block: Block
    call_trace: _containers.RepeatedCompositeFieldContainer[CallTrace]
    transaction_receipt: _containers.RepeatedCompositeFieldContainer[TransactionReceipt]
    def __init__(self, block: _Optional[_Union[Block, _Mapping]] = ..., transaction_receipt: _Optional[_Iterable[_Union[TransactionReceipt, _Mapping]]] = ..., call_trace: _Optional[_Iterable[_Union[CallTrace, _Mapping]]] = ...) -> None: ...

class Input(_message.Message):
    __slots__ = ["args", "function", "method_id"]
    ARGS_FIELD_NUMBER: _ClassVar[int]
    FUNCTION_FIELD_NUMBER: _ClassVar[int]
    METHOD_ID_FIELD_NUMBER: _ClassVar[int]
    args: _containers.RepeatedCompositeFieldContainer[Param]
    function: str
    method_id: str
    def __init__(self, function: _Optional[str] = ..., method_id: _Optional[str] = ..., args: _Optional[_Iterable[_Union[Param, _Mapping]]] = ...) -> None: ...

class Log(_message.Message):
    __slots__ = ["address", "block_hash", "block_number", "data", "log_index", "removed", "topics", "transaction_hash", "transaction_index"]
    ADDRESS_FIELD_NUMBER: _ClassVar[int]
    BLOCK_HASH_FIELD_NUMBER: _ClassVar[int]
    BLOCK_NUMBER_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    LOG_INDEX_FIELD_NUMBER: _ClassVar[int]
    REMOVED_FIELD_NUMBER: _ClassVar[int]
    TOPICS_FIELD_NUMBER: _ClassVar[int]
    TRANSACTION_HASH_FIELD_NUMBER: _ClassVar[int]
    TRANSACTION_INDEX_FIELD_NUMBER: _ClassVar[int]
    address: str
    block_hash: str
    block_number: int
    data: _containers.RepeatedCompositeFieldContainer[LogParams]
    log_index: int
    removed: bool
    topics: Topic
    transaction_hash: str
    transaction_index: int
    def __init__(self, address: _Optional[str] = ..., topics: _Optional[_Union[Topic, _Mapping]] = ..., data: _Optional[_Iterable[_Union[LogParams, _Mapping]]] = ..., block_number: _Optional[int] = ..., transaction_hash: _Optional[str] = ..., transaction_index: _Optional[int] = ..., block_hash: _Optional[str] = ..., log_index: _Optional[int] = ..., removed: bool = ...) -> None: ...

class LogParams(_message.Message):
    __slots__ = ["name", "type", "value"]
    NAME_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    VALUE_FIELD_NUMBER: _ClassVar[int]
    name: str
    type: str
    value: str
    def __init__(self, name: _Optional[str] = ..., type: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...

class Param(_message.Message):
    __slots__ = ["data", "name", "type"]
    DATA_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    data: str
    name: str
    type: str
    def __init__(self, name: _Optional[str] = ..., type: _Optional[str] = ..., data: _Optional[str] = ...) -> None: ...

class Topic(_message.Message):
    __slots__ = ["event_id", "indexed_params", "signature"]
    EVENT_ID_FIELD_NUMBER: _ClassVar[int]
    INDEXED_PARAMS_FIELD_NUMBER: _ClassVar[int]
    SIGNATURE_FIELD_NUMBER: _ClassVar[int]
    event_id: str
    indexed_params: _containers.RepeatedCompositeFieldContainer[LogParams]
    signature: str
    def __init__(self, signature: _Optional[str] = ..., indexed_params: _Optional[_Iterable[_Union[LogParams, _Mapping]]] = ..., event_id: _Optional[str] = ...) -> None: ...

class Transaction(_message.Message):
    __slots__ = ["access_list", "block_hash", "block_number", "chain_id", "fee", "gas", "gas_price", "hash", "input", "is_transfer", "max_fee_per_gas", "max_priority_fee_per_gas", "nonce", "r", "s", "successful", "to", "type", "v", "value"]
    ACCESS_LIST_FIELD_NUMBER: _ClassVar[int]
    BLOCK_HASH_FIELD_NUMBER: _ClassVar[int]
    BLOCK_NUMBER_FIELD_NUMBER: _ClassVar[int]
    CHAIN_ID_FIELD_NUMBER: _ClassVar[int]
    FEE_FIELD_NUMBER: _ClassVar[int]
    FROM_FIELD_NUMBER: _ClassVar[int]
    GAS_FIELD_NUMBER: _ClassVar[int]
    GAS_PRICE_FIELD_NUMBER: _ClassVar[int]
    HASH_FIELD_NUMBER: _ClassVar[int]
    INPUT_FIELD_NUMBER: _ClassVar[int]
    IS_TRANSFER_FIELD_NUMBER: _ClassVar[int]
    MAX_FEE_PER_GAS_FIELD_NUMBER: _ClassVar[int]
    MAX_PRIORITY_FEE_PER_GAS_FIELD_NUMBER: _ClassVar[int]
    NONCE_FIELD_NUMBER: _ClassVar[int]
    R_FIELD_NUMBER: _ClassVar[int]
    SUCCESSFUL_FIELD_NUMBER: _ClassVar[int]
    S_FIELD_NUMBER: _ClassVar[int]
    TO_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    VALUE_FIELD_NUMBER: _ClassVar[int]
    V_FIELD_NUMBER: _ClassVar[int]
    access_list: _containers.RepeatedCompositeFieldContainer[Access]
    block_hash: str
    block_number: int
    chain_id: int
    fee: float
    gas: int
    gas_price: int
    hash: str
    input: Input
    is_transfer: bool
    max_fee_per_gas: int
    max_priority_fee_per_gas: int
    nonce: int
    r: str
    s: str
    successful: bool
    to: str
    type: int
    v: str
    value: int
    def __init__(self, block_hash: _Optional[str] = ..., block_number: _Optional[int] = ..., gas: _Optional[int] = ..., gas_price: _Optional[int] = ..., hash: _Optional[str] = ..., input: _Optional[_Union[Input, _Mapping]] = ..., max_fee_per_gas: _Optional[int] = ..., max_priority_fee_per_gas: _Optional[int] = ..., nonce: _Optional[int] = ..., r: _Optional[str] = ..., s: _Optional[str] = ..., to: _Optional[str] = ..., type: _Optional[int] = ..., v: _Optional[str] = ..., value: _Optional[int] = ..., access_list: _Optional[_Iterable[_Union[Access, _Mapping]]] = ..., chain_id: _Optional[int] = ..., fee: _Optional[float] = ..., successful: bool = ..., is_transfer: bool = ..., **kwargs) -> None: ...

class TransactionReceipt(_message.Message):
    __slots__ = ["block_hash", "block_number", "contract_address", "cumulative_gas_used", "effective_gas_price", "gas_used", "logs", "logs_bloom", "root", "status", "to", "transaction_hash", "transaction_index", "type"]
    BLOCK_HASH_FIELD_NUMBER: _ClassVar[int]
    BLOCK_NUMBER_FIELD_NUMBER: _ClassVar[int]
    CONTRACT_ADDRESS_FIELD_NUMBER: _ClassVar[int]
    CUMULATIVE_GAS_USED_FIELD_NUMBER: _ClassVar[int]
    EFFECTIVE_GAS_PRICE_FIELD_NUMBER: _ClassVar[int]
    FROM_FIELD_NUMBER: _ClassVar[int]
    GAS_USED_FIELD_NUMBER: _ClassVar[int]
    LOGS_BLOOM_FIELD_NUMBER: _ClassVar[int]
    LOGS_FIELD_NUMBER: _ClassVar[int]
    ROOT_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    TO_FIELD_NUMBER: _ClassVar[int]
    TRANSACTION_HASH_FIELD_NUMBER: _ClassVar[int]
    TRANSACTION_INDEX_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    block_hash: str
    block_number: int
    contract_address: str
    cumulative_gas_used: int
    effective_gas_price: int
    gas_used: int
    logs: _containers.RepeatedCompositeFieldContainer[Log]
    logs_bloom: str
    root: str
    status: int
    to: str
    transaction_hash: str
    transaction_index: str
    type: int
    def __init__(self, block_hash: _Optional[str] = ..., block_number: _Optional[int] = ..., contract_address: _Optional[str] = ..., cumulative_gas_used: _Optional[int] = ..., effective_gas_price: _Optional[int] = ..., gas_used: _Optional[int] = ..., logs: _Optional[_Iterable[_Union[Log, _Mapping]]] = ..., logs_bloom: _Optional[str] = ..., status: _Optional[int] = ..., to: _Optional[str] = ..., transaction_hash: _Optional[str] = ..., transaction_index: _Optional[str] = ..., type: _Optional[int] = ..., root: _Optional[str] = ..., **kwargs) -> None: ...
