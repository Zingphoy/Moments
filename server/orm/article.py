# encoding:utf-8

from sqlalchemy import Column, INT, TEXT
from sqlalchemy.ext.declarative import declarative_base

Base = declarative_base()


class Article(Base):
    __tablename__ = "artile"

    aid = Column(INT(), primary_key=True)
    send_time = Column(INT(), nullable=False)
    content = Column(TEXT(), nullable=False)
    photo_list = Column(TEXT(), default=None)
    privacy = Column(INT(), default=0)
    is_deleted = Column(INT(), default=0)

    def __repr__(self):
        return f"<Artile(aid={self.aid},send_time={self.send_time},content={self.content},photo_list={self.photo_list},privacy={self.privacy},is_delete={self.is_deleted})"
