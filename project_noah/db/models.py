import uuid
from datetime import datetime, timezone
from typing import List, Optional

from sqlalchemy import DateTime, ForeignKey, String, Uuid
from sqlalchemy.orm import DeclarativeBase, Mapped, Session, mapped_column, relationship


class Base(DeclarativeBase):
    __abstract__ = True

    id: Mapped[int] = mapped_column(primary_key=True)
    public_id: Mapped[Uuid] = mapped_column(server_default=lambda: uuid.uuid4())
    created_at: Mapped[DateTime] = mapped_column(
        server_default=lambda: datetime.now(timezone.utc)
    )
    updated_at: Mapped[DateTime] = mapped_column(
        server_default=lambda: datetime.now(timezone.utc)
    )


class Campaign(Base):
    pass


class Mission(Base):
    pass


class Scene(Base):
    pass


class Combat(Base):
    pass


class NPC(Base):
    pass
