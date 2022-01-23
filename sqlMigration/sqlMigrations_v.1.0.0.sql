-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE IF NOT EXISTS userAccount (
	userID   varchar(20) NOT NULL ,
	userName varchar(50) NOT NULL ,
	joinDate datetime NOT NULL ,
	PRIMARY KEY (userID)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS notesType (
	typeID   varchar(3) NOT NULL ,
	typeName varchar(20) NOT NULL ,
	PRIMARY KEY (typeID)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS walletGroup (
	walletGroupID varchar(15) NOT NULL ,
	userID        varchar(20) NOT NULL ,
	groupName     varchar(50) NOT NULL ,
	PRIMARY KEY (walletGroupID),
	KEY fkIdx_84 (userID),
	CONSTRAINT FK_82 FOREIGN KEY fkIdx_84 (userID) REFERENCES userAccount (userID)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS wallet (
	walletID      varchar(15) NOT NULL ,
	userID        varchar(20) NOT NULL ,
	walletGroupID varchar(15) NOT NULL ,
	walletName    varchar(50) NOT NULL ,
	amount        bigint NOT NULL ,
	description   varchar(100) NULL ,

	PRIMARY KEY (walletID),
	KEY fkIdx_114 (userID),
	CONSTRAINT FK_112 FOREIGN KEY fkIdx_114 (userID) REFERENCES userAccount (userID),
	KEY fkIdx_131 (walletGroupID),
	CONSTRAINT FK_129 FOREIGN KEY fkIdx_131 (walletGroupID) REFERENCES walletGroup (walletGroupID)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS notesCategory (
	categoryID   varchar(15) NOT NULL ,
	userID       varchar(20) NOT NULL ,
	categoryName varchar(20) NOT NULL ,

	PRIMARY KEY (categoryID),
	KEY fkIdx_107 (userID),
	CONSTRAINT FK_105 FOREIGN KEY fkIdx_107 (userID) REFERENCES userAccount (userID)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS notes (
	notesID         bigint NOT NULL ,
	userID          varchar(20) NOT NULL ,
	walletID        varchar(15) NOT NULL ,
	notesTypeID     varchar(3) NOT NULL ,
	notesCategoryID varchar(15) NOT NULL ,
	notesName       varchar(50) NOT NULL ,
	amount          bigint NOT NULL ,
	fromAccount     varchar(20) NULL ,
	toAccount       varchar(20) NULL ,
	description     varchar(100) NULL ,
	notesDate       datetime NOT NULL ,
	
	PRIMARY KEY (notesID),
	KEY fkIdx_100 (userID),
	CONSTRAINT FK_98 FOREIGN KEY fkIdx_100 (userID) REFERENCES userAccount (userID),
	KEY fkIdx_125 (walletID),
	CONSTRAINT FK_123 FOREIGN KEY fkIdx_125 (walletID) REFERENCES wallet (walletID),
	KEY fkIdx_25 (notesTypeID),
	CONSTRAINT FK_24 FOREIGN KEY fkIdx_25 (notesTypeID) REFERENCES notesType (typeID),
	KEY fkIdx_35 (notesCategoryID),
	CONSTRAINT FK_34 FOREIGN KEY fkIdx_35 (notesCategoryID) REFERENCES notesCategory (categoryID)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

#define notes type (Income, Expence, Transfer)
INSERT IGNORE INTO notesType(typeID, typeName)
VALUES('inc', 'Income'), ('exp', 'Expence'), ('trf', 'Transfer');
-- +migrate StatementEnd