-- +goose Up
CREATE TABLE if not exists notification_types
(
    id          serial      not null,
    type        varchar(50) NOT NULL,
    description varchar(200),
    unique (id),
    unique (type)

);

INSERT INTO notification_types (type, description)
VALUES ('Email', 'Email notifications'),
       ('Sms', 'SMS notifications'),
       ('Web Hook', 'Notifications via http');

CREATE TABLE if not exists notifications
(
    id                serial not null,
    notification_type serial references notification_types (id),
    title              varchar(30),
    description       varchar(200),
    template          text   NOT NULL
);

CREATE TABLE if not exists email_notifications
(
    email       varchar(100),
    pwd         varchar(20),
    smtp_server varchar(20),
    smtp_port   int,
    unique (id)
) inherits (notifications);

-- +goose Down
DROP TABLE email_notifications;
DROP TABLE notifications;
DROP TABLE notification_types;
