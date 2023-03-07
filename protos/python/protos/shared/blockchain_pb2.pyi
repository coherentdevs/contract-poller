from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from typing import ClassVar as _ClassVar

ARBITRUM: Blockchain
AVALANCHE: Blockchain
DESCRIPTOR: _descriptor.FileDescriptor
ETHEREUM: Blockchain
GOERLI: Blockchain
OPTIMISM: Blockchain
POLYGON: Blockchain

class Blockchain(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
