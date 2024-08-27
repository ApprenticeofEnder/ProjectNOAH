import uuid
from datetime import datetime, timezone
from typing import List, Optional

from sqlalchemy import ForeignKey, String, create_engine
from sqlalchemy.orm import (
    DeclarativeBase,
    Mapped,
    declared_attr,
    mapped_column,
    relationship,
    sessionmaker,
)

# TODO: Change this for an environment variable
engine = create_engine(
    "postgresql://postgres:postgres@127.0.0.1:54322/postgres", echo=True
)

Session = sessionmaker(engine)


class Base(DeclarativeBase):
    id: Mapped[int] = mapped_column(primary_key=True)
    public_id: Mapped[uuid.UUID] = mapped_column(default=uuid.uuid4())
    created_at: Mapped[datetime] = mapped_column(
        default=lambda: datetime.now(timezone.utc)
    )
    updated_at: Mapped[datetime] = mapped_column(
        default=lambda: datetime.now(timezone.utc)
    )
