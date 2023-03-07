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
    __slots__ = ["base_fee_per_gas", "difficulty", "extra_data", "gas_limit", "gas_used", "hash", "logs_bloom", "miner", "mix_hash", "nonce", "number", "parent_hash", "receipts_root", "sha3_uncles", "size", "state_root", "timestamp", "total_difficulty", "transactions", "transactions_root", "uncles"]
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
    RECEIPTS_ROOT_FIELD_NUMBER: _ClassVar[int]
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
    number: str
    parent_hash: str
    receipts_root: str
    sha3_uncles: str
    size: str
    state_root: str
    timestamp: str
    total_difficulty: str
    transactions: _containers.RepeatedCompositeFieldContainer[Transaction]
    transactions_root: str
    uncles: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, number: _Optional[str] = ..., hash: _Optional[str] = ..., parent_hash: _Optional[str] = ..., nonce: _Optional[str] = ..., sha3_uncles: _Optional[str] = ..., logs_bloom: _Optional[str] = ..., transactions_root: _Optional[str] = ..., state_root: _Optional[str] = ..., receipts_root: _Optional[str] = ..., miner: _Optional[str] = ..., difficulty: _Optional[str] = ..., total_difficulty: _Optional[str] = ..., extra_data: _Optional[str] = ..., size: _Optional[str] = ..., gas_limit: _Optional[str] = ..., gas_used: _Optional[str] = ..., timestamp: _Optional[str] = ..., transactions: _Optional[_Iterable[_Union[Transaction, _Mapping]]] = ..., uncles: _Optional[_Iterable[str]] = ..., base_fee_per_gas: _Optional[str] = ..., mix_hash: _Optional[str] = ...) -> None: ...

class CallTrace(_message.Message):
    __slots__ = ["calls", "error", "gas", "gas_used", "input", "output", "revertReason", "time", "to", "type", "value"]
    CALLS_FIELD_NUMBER: _ClassVar[int]
    ERROR_FIELD_NUMBER: _ClassVar[int]
    FROM_FIELD_NUMBER: _ClassVar[int]
    GAS_FIELD_NUMBER: _ClassVar[int]
    GAS_USED_FIELD_NUMBER: _ClassVar[int]
    INPUT_FIELD_NUMBER: _ClassVar[int]
    OUTPUT_FIELD_NUMBER: _ClassVar[int]
    REVERTREASON_FIELD_NUMBER: _ClassVar[int]
    TIME_FIELD_NUMBER: _ClassVar[int]
    TO_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    VALUE_FIELD_NUMBER: _ClassVar[int]
    calls: _containers.RepeatedCompositeFieldContainer[CallTrace]
    error: str
    gas: str
    gas_used: str
    input: str
    output: str
    revertReason: str
    time: str
    to: str
    type: str
    value: str
    def __init__(self, type: _Optional[str] = ..., to: _Optional[str] = ..., value: _Optional[str] = ..., gas: _Optional[str] = ..., gas_used: _Optional[str] = ..., input: _Optional[str] = ..., output: _Optional[str] = ..., time: _Optional[str] = ..., error: _Optional[str] = ..., revertReason: _Optional[str] = ..., calls: _Optional[_Iterable[_Union[CallTrace, _Mapping]]] = ..., **kwargs) -> None: ...

class Data(_message.Message):
    __slots__ = ["block", "call_traces", "transaction_receipts"]
    BLOCK_FIELD_NUMBER: _ClassVar[int]
    CALL_TRACES_FIELD_NUMBER: _ClassVar[int]
    TRANSACTION_RECEIPTS_FIELD_NUMBER: _ClassVar[int]
    block: Block
    call_traces: _containers.RepeatedCompositeFieldContainer[CallTrace]
    transaction_receipts: _containers.RepeatedCompositeFieldContainer[TransactionReceipt]
    def __init__(self, block: _Optional[_Union[Block, _Mapping]] = ..., transaction_receipts: _Optional[_Iterable[_Union[TransactionReceipt, _Mapping]]] = ..., call_traces: _Optional[_Iterable[_Union[CallTrace, _Mapping]]] = ...) -> None: ...

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
    block_number: str
    data: str
    log_index: str
    removed: bool
    topics: _containers.RepeatedScalarFieldContainer[str]
    transaction_hash: str
    transaction_index: str
    def __init__(self, removed: bool = ..., log_index: _Optional[str] = ..., transaction_index: _Optional[str] = ..., transaction_hash: _Optional[str] = ..., block_number: _Optional[str] = ..., block_hash: _Optional[str] = ..., address: _Optional[str] = ..., data: _Optional[str] = ..., topics: _Optional[_Iterable[str]] = ...) -> None: ...

class Transaction(_message.Message):
    __slots__ = ["access_list", "block_hash", "block_number", "chain_id", "gas", "gas_price", "hash", "index", "input", "l1_block_number", "l1_timestamp", "l1_tx_origin", "max_fee_per_gas", "max_priority_fee_per_gas", "nonce", "queue_index", "queue_origin", "r", "raw_transaction", "s", "to", "transaction_index", "tx_type", "type", "v", "value"]
    ACCESS_LIST_FIELD_NUMBER: _ClassVar[int]
    BLOCK_HASH_FIELD_NUMBER: _ClassVar[int]
    BLOCK_NUMBER_FIELD_NUMBER: _ClassVar[int]
    CHAIN_ID_FIELD_NUMBER: _ClassVar[int]
    FROM_FIELD_NUMBER: _ClassVar[int]
    GAS_FIELD_NUMBER: _ClassVar[int]
    GAS_PRICE_FIELD_NUMBER: _ClassVar[int]
    HASH_FIELD_NUMBER: _ClassVar[int]
    INDEX_FIELD_NUMBER: _ClassVar[int]
    INPUT_FIELD_NUMBER: _ClassVar[int]
    L1_BLOCK_NUMBER_FIELD_NUMBER: _ClassVar[int]
    L1_TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
    L1_TX_ORIGIN_FIELD_NUMBER: _ClassVar[int]
    MAX_FEE_PER_GAS_FIELD_NUMBER: _ClassVar[int]
    MAX_PRIORITY_FEE_PER_GAS_FIELD_NUMBER: _ClassVar[int]
    NONCE_FIELD_NUMBER: _ClassVar[int]
    QUEUE_INDEX_FIELD_NUMBER: _ClassVar[int]
    QUEUE_ORIGIN_FIELD_NUMBER: _ClassVar[int]
    RAW_TRANSACTION_FIELD_NUMBER: _ClassVar[int]
    R_FIELD_NUMBER: _ClassVar[int]
    S_FIELD_NUMBER: _ClassVar[int]
    TO_FIELD_NUMBER: _ClassVar[int]
    TRANSACTION_INDEX_FIELD_NUMBER: _ClassVar[int]
    TX_TYPE_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    VALUE_FIELD_NUMBER: _ClassVar[int]
    V_FIELD_NUMBER: _ClassVar[int]
    access_list: _containers.RepeatedCompositeFieldContainer[Access]
    block_hash: str
    block_number: str
    chain_id: str
    gas: str
    gas_price: str
    hash: str
    index: str
    input: str
    l1_block_number: str
    l1_timestamp: str
    l1_tx_origin: str
    max_fee_per_gas: str
    max_priority_fee_per_gas: str
    nonce: str
    queue_index: str
    queue_origin: str
    r: str
    raw_transaction: str
    s: str
    to: str
    transaction_index: str
    tx_type: str
    type: str
    v: str
    value: str
    def __init__(self, block_hash: _Optional[str] = ..., block_number: _Optional[str] = ..., gas: _Optional[str] = ..., gas_price: _Optional[str] = ..., hash: _Optional[str] = ..., input: _Optional[str] = ..., nonce: _Optional[str] = ..., to: _Optional[str] = ..., transaction_index: _Optional[str] = ..., value: _Optional[str] = ..., v: _Optional[str] = ..., r: _Optional[str] = ..., s: _Optional[str] = ..., access_list: _Optional[_Iterable[_Union[Access, _Mapping]]] = ..., chain_id: _Optional[str] = ..., max_fee_per_gas: _Optional[str] = ..., max_priority_fee_per_gas: _Optional[str] = ..., type: _Optional[str] = ..., queue_origin: _Optional[str] = ..., tx_type: _Optional[str] = ..., l1_tx_origin: _Optional[str] = ..., l1_block_number: _Optional[str] = ..., l1_timestamp: _Optional[str] = ..., index: _Optional[str] = ..., queue_index: _Optional[str] = ..., raw_transaction: _Optional[str] = ..., **kwargs) -> None: ...

class TransactionReceipt(_message.Message):
    __slots__ = ["block_hash", "block_number", "contract_address", "cumulative_gas_used", "effectiveGasPrice", "gas_used", "l1_fee", "l1_fee_scalar", "l1_gas_price", "l1_gas_used", "logs", "logs_bloom", "root", "status", "to", "transaction_hash", "transaction_index", "type"]
    BLOCK_HASH_FIELD_NUMBER: _ClassVar[int]
    BLOCK_NUMBER_FIELD_NUMBER: _ClassVar[int]
    CONTRACT_ADDRESS_FIELD_NUMBER: _ClassVar[int]
    CUMULATIVE_GAS_USED_FIELD_NUMBER: _ClassVar[int]
    EFFECTIVEGASPRICE_FIELD_NUMBER: _ClassVar[int]
    FROM_FIELD_NUMBER: _ClassVar[int]
    GAS_USED_FIELD_NUMBER: _ClassVar[int]
    L1_FEE_FIELD_NUMBER: _ClassVar[int]
    L1_FEE_SCALAR_FIELD_NUMBER: _ClassVar[int]
    L1_GAS_PRICE_FIELD_NUMBER: _ClassVar[int]
    L1_GAS_USED_FIELD_NUMBER: _ClassVar[int]
    LOGS_BLOOM_FIELD_NUMBER: _ClassVar[int]
    LOGS_FIELD_NUMBER: _ClassVar[int]
    ROOT_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    TO_FIELD_NUMBER: _ClassVar[int]
    TRANSACTION_HASH_FIELD_NUMBER: _ClassVar[int]
    TRANSACTION_INDEX_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    block_hash: str
    block_number: str
    contract_address: str
    cumulative_gas_used: str
    effectiveGasPrice: str
    gas_used: str
    l1_fee: str
    l1_fee_scalar: str
    l1_gas_price: str
    l1_gas_used: str
    logs: _containers.RepeatedCompositeFieldContainer[Log]
    logs_bloom: str
    root: str
    status: str
    to: str
    transaction_hash: str
    transaction_index: str
    type: str
    def __init__(self, transaction_hash: _Optional[str] = ..., transaction_index: _Optional[str] = ..., block_hash: _Optional[str] = ..., block_number: _Optional[str] = ..., to: _Optional[str] = ..., cumulative_gas_used: _Optional[str] = ..., effectiveGasPrice: _Optional[str] = ..., gas_used: _Optional[str] = ..., contract_address: _Optional[str] = ..., logs: _Optional[_Iterable[_Union[Log, _Mapping]]] = ..., logs_bloom: _Optional[str] = ..., type: _Optional[str] = ..., root: _Optional[str] = ..., status: _Optional[str] = ..., l1_fee: _Optional[str] = ..., l1_fee_scalar: _Optional[str] = ..., l1_gas_price: _Optional[str] = ..., l1_gas_used: _Optional[str] = ..., **kwargs) -> None: ...
