CREATE TABLE likes (
    likeId int auto_increment primary key,
    contentId int NOT NULL,
    userId int NOT NULL,
    status int NOT NULL,
    createdAt datetime default CURRENT_TIMESTAMP NOT NULL,
    updatedAt datetime default CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP
);
CREATE INDEX `likeId` ON likes (likeId);
CREATE INDEX `contentId` ON likes (contentId);
CREATE INDEX `userId` ON likes (userId);

CREATE TABLE readInteraction (
    readId int auto_increment primary key,
    contentId int NOT NULL,
    userId int NOT NULL,
    status int NOT NULL,
    createdAt datetime default CURRENT_TIMESTAMP NOT NULL,
    updatedAt datetime default CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP
);
CREATE INDEX `readId` ON readInteraction (readId);
CREATE INDEX `contentId` ON readInteraction (contentId);
CREATE INDEX `userId` ON readInteraction (userId);