from sqlalchemy import ForeignKey
from sqlalchemy.orm import Mapped, declared_attr, mapped_column, relationship

from project_noah.db.database import Base, Session


class TitleMixin:
    title: Mapped[str]


class NotesMixin:
    notes: Mapped[str]


class CampaignItem:
    campaign_id = mapped_column(ForeignKey("campaigns.id"))

    @declared_attr
    def campaign(self):
        return relationship("Campaign")


class Campaign(TitleMixin, Base):
    __tablename__ = "campaigns"

    description: Mapped[str]


class Mission(TitleMixin, NotesMixin, CampaignItem, Base):
    __tablename__ = "missions"


class Scene(TitleMixin, NotesMixin, CampaignItem, Base):
    __tablename__ = "scenes"


class Combat(TitleMixin, NotesMixin, CampaignItem, Base):
    __tablename__ = "combats"


class NPC(NotesMixin, CampaignItem, Base):
    __tablename__ = "npcs"

    name: Mapped[str]
