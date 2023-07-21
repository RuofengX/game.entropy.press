from typing import Any


class SingletonMixin:
    __singleton: dict[str, Any] = dict()  # 保存所有单例的字典

    def __new__(cls, *args: Any, **kwargs: Any):
        if cls.__name__ not in cls.__singleton:
            cls.__singleton[cls.__name__] = super().__new__(cls, *args, **kwargs)
        return cls.__singleton[cls.__name__]
