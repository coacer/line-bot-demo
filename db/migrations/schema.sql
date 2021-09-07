CREATE TABLE LineBotChannel (
    Id                 INT64 NOT NULL,
    ChannelId          STRING(MAX) NOT NULL,
    ChannelSecretId    STRING(MAX) NOT NULL,
    ChannelAccessToken STRING(MAX) NOT NULL,
    CreatedAt          TIMESTAMP NOT NULL OPTIONS ( allow_commit_timestamp = true ),
    UpdatedAt          TIMESTAMP NOT NULL OPTIONS ( allow_commit_timestamp = true )
) PRIMARY KEY (Id);

CREATE TABLE User (
    Id                INT64 NOT NULL,
    LineBotChannelId  INT64 NOT NULL,
    LineUID           STRING(MAX) NOT NULL,
    CreatedAt          TIMESTAMP NOT NULL OPTIONS ( allow_commit_timestamp = true ),
    UpdatedAt          TIMESTAMP NOT NULL OPTIONS ( allow_commit_timestamp = true ),
    CONSTRAINT FK_LineBotChannelId FOREIGN KEY (LineBotChannelId) REFERENCES LineBotChannel (Id)
) PRIMARY KEY (Id);

CREATE TABLE UserEvent (
    Id                INT64 NOT NULL,
    UserId            INT64 NOT NULL,
    Type              STRING(256) NOT NULL,
    Content           STRING(MAX) NOT NULL,
    CreatedAt          TIMESTAMP NOT NULL OPTIONS ( allow_commit_timestamp = true ),
    UpdatedAt          TIMESTAMP NOT NULL OPTIONS ( allow_commit_timestamp = true ),
    CONSTRAINT FK_UserId FOREIGN KEY (UserId) REFERENCES User (Id)
) PRIMARY KEY (Id);

