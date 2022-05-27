-- This table is for registering clients.
CREATE TABLE Clients (chat_id VARCHAR(32), token VARCHAR(64), refresh_token VARCHAR(64), PRIMARY KEY(chat_id));
CREATE TABLE Codes (client_id VARCHAR(32), code VARCHAR(64), PRIMARY KEY(client_id, code)); 