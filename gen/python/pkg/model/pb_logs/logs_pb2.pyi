from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Severity(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    SEVERITY_UNSPECIFIED: _ClassVar[Severity]
    SEVERITY_DEBUG: _ClassVar[Severity]
    SEVERITY_INFO: _ClassVar[Severity]
    SEVERITY_WARN: _ClassVar[Severity]
    SEVERITY_ERROR: _ClassVar[Severity]
    SEVERITY_CRITICAL: _ClassVar[Severity]
    SEVERITY_SUCCESS: _ClassVar[Severity]
SEVERITY_UNSPECIFIED: Severity
SEVERITY_DEBUG: Severity
SEVERITY_INFO: Severity
SEVERITY_WARN: Severity
SEVERITY_ERROR: Severity
SEVERITY_CRITICAL: Severity
SEVERITY_SUCCESS: Severity

class RuntimeLogs(_message.Message):
    __slots__ = ("log_date", "severity", "source", "message", "context", "path", "details", "tags")
    class ContextEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    class TagsEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    LOG_DATE_FIELD_NUMBER: _ClassVar[int]
    SEVERITY_FIELD_NUMBER: _ClassVar[int]
    SOURCE_FIELD_NUMBER: _ClassVar[int]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    CONTEXT_FIELD_NUMBER: _ClassVar[int]
    PATH_FIELD_NUMBER: _ClassVar[int]
    DETAILS_FIELD_NUMBER: _ClassVar[int]
    TAGS_FIELD_NUMBER: _ClassVar[int]
    log_date: _timestamp_pb2.Timestamp
    severity: Severity
    source: str
    message: str
    context: _containers.ScalarMap[str, str]
    path: str
    details: str
    tags: _containers.ScalarMap[str, str]
    def __init__(self, log_date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., severity: _Optional[_Union[Severity, str]] = ..., source: _Optional[str] = ..., message: _Optional[str] = ..., context: _Optional[_Mapping[str, str]] = ..., path: _Optional[str] = ..., details: _Optional[str] = ..., tags: _Optional[_Mapping[str, str]] = ...) -> None: ...

class SendResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class BatchSendRequest(_message.Message):
    __slots__ = ("logs",)
    LOGS_FIELD_NUMBER: _ClassVar[int]
    logs: _containers.RepeatedCompositeFieldContainer[RuntimeLogs]
    def __init__(self, logs: _Optional[_Iterable[_Union[RuntimeLogs, _Mapping]]] = ...) -> None: ...

class BatchSendResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...
