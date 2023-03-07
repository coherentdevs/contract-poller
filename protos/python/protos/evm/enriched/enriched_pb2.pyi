from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class AssetTransfer(_message.Message):
    __slots__ = ["amount", "asset_address", "asset_name", "standard", "symbol", "to", "token_id"]
    AMOUNT_FIELD_NUMBER: _ClassVar[int]
    ASSET_ADDRESS_FIELD_NUMBER: _ClassVar[int]
    ASSET_NAME_FIELD_NUMBER: _ClassVar[int]
    FROM_FIELD_NUMBER: _ClassVar[int]
    STANDARD_FIELD_NUMBER: _ClassVar[int]
    SYMBOL_FIELD_NUMBER: _ClassVar[int]
    TOKEN_ID_FIELD_NUMBER: _ClassVar[int]
    TO_FIELD_NUMBER: _ClassVar[int]
    amount: str
    asset_address: str
    asset_name: str
    standard: str
    symbol: str
    to: str
    token_id: str
    def __init__(self, asset_address: _Optional[str] = ..., amount: _Optional[str] = ..., asset_name: _Optional[str] = ..., to: _Optional[str] = ..., standard: _Optional[str] = ..., symbol: _Optional[str] = ..., token_id: _Optional[str] = ..., **kwargs) -> None: ...

class Block(_message.Message):
    __slots__ = ["block_hash", "block_number", "timestamp", "transactions"]
    BLOCK_HASH_FIELD_NUMBER: _ClassVar[int]
    BLOCK_NUMBER_FIELD_NUMBER: _ClassVar[int]
    TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
    TRANSACTIONS_FIELD_NUMBER: _ClassVar[int]
    block_hash: str
    block_number: int
    timestamp: _timestamp_pb2.Timestamp
    transactions: _containers.RepeatedCompositeFieldContainer[Transaction]
    def __init__(self, block_hash: _Optional[str] = ..., block_number: _Optional[int] = ..., timestamp: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., transactions: _Optional[_Iterable[_Union[Transaction, _Mapping]]] = ...) -> None: ...

class Transaction(_message.Message):
    __slots__ = ["assets_received", "assets_sent", "gas", "gas_price", "hash", "object", "to", "verb"]
    ASSETS_RECEIVED_FIELD_NUMBER: _ClassVar[int]
    ASSETS_SENT_FIELD_NUMBER: _ClassVar[int]
    FROM_FIELD_NUMBER: _ClassVar[int]
    GAS_FIELD_NUMBER: _ClassVar[int]
    GAS_PRICE_FIELD_NUMBER: _ClassVar[int]
    HASH_FIELD_NUMBER: _ClassVar[int]
    OBJECT_FIELD_NUMBER: _ClassVar[int]
    TO_FIELD_NUMBER: _ClassVar[int]
    VERB_FIELD_NUMBER: _ClassVar[int]
    assets_received: _containers.RepeatedCompositeFieldContainer[AssetTransfer]
    assets_sent: _containers.RepeatedCompositeFieldContainer[AssetTransfer]
    gas: str
    gas_price: str
    hash: str
    object: str
    to: str
    verb: str
    def __init__(self, assets_received: _Optional[_Iterable[_Union[AssetTransfer, _Mapping]]] = ..., assets_sent: _Optional[_Iterable[_Union[AssetTransfer, _Mapping]]] = ..., gas: _Optional[str] = ..., gas_price: _Optional[str] = ..., hash: _Optional[str] = ..., object: _Optional[str] = ..., to: _Optional[str] = ..., verb: _Optional[str] = ..., **kwargs) -> None: ...
