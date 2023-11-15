import os

from features.steps.logger import httpclient_logging_patch

TARGET = os.getenv("TARGET", "localhost:31002")


class Config:
    httpclient_logging_patch()
    address = TARGET.split(":")
    ports = address[1].split(",")

    @classmethod
    def get_transaction_url(cls):
        return f"{cls.address[0]}:{cls.ports[0]}"
    