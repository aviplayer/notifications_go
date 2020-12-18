package accessors

import (
	. "andrew.com/notifications/cmd/app/models"
	"context"
	"errors"
	"github.com/georgysavva/scany/pgxscan"
	"log"
)

const (
	insertNotification = `insert into email_notifications
(title, notification_type, template, email, pwd, smtp_server, smtp_port, description)
values($1, $2, $3, $4, $5, $6, $7, $8)
returning id`

	updateNotification = `insert into email_notifications
(id, title, notification_type, template, email, pwd, smtp_server, smtp_port, description)
values($1, $2, $3, $4, $5, $6, $7, $8, $9)
on conflict (id)
do update set template = EXCLUDED.template,
title = EXCLUDED.title,
notification_type = EXCLUDED.notification_type,
email = EXCLUDED.email, 
pwd = EXCLUDED.pwd, 
smtp_server = EXCLUDED.smtp_server,
description = EXCLUDED.description
returning id`

	selectNotifications = `select title, notification_type, template, email, pwd, smtp_server, smtp_port, description
from email_notifications
`
	selectNotification = selectNotifications + `
where id=$1
`

	deleteNotification = `DELETE FROM email_notifications
	WHERE type=$1;`
)

type EmailAccessor struct {
	Conn *DbConnection
}

func (acc EmailAccessor) InsertNotification(notification EmailNotification) (EmailNotification, error) {
	id := 0
	n := notification
	b := n.BaseNotification
	usr, psw, server, port := n.GetAuth()
	log.Printf("Values to be inserted in to email_notification usr: %s, pass: %s, server: %s, port: %d, description: %s \n",
		usr, psw, server, port, b.Description)
	conn, err := acc.Conn.getConnection()
	if err != nil {
		return EmailNotification{}, err
	}
	if err := conn.QueryRow(
		context.Background(),
		insertNotification,
		b.Title, Types[Email], b.Template, usr, psw, server, port, b.Description,
	).Scan(&id); err != nil {
		return EmailNotification{}, err
	}
	notification.BaseNotification.Id = int8(id)
	return notification, nil
}

func (acc EmailAccessor) upsertNotification(notification EmailNotification) error {
	id := 0
	n := notification
	b := n.BaseNotification
	usr, psw, server, port := n.GetAuth()
	log.Printf("Values to be updated in to email_notification id: %d, usr: %s, pass: %s, server: %s, port: %d, description: %s \n",
		b.Id, usr, psw, server, port, b.Description)
	conn, err := acc.Conn.getConnection()
	if err != nil {
		return err
	}
	if err := conn.QueryRow(
		context.Background(),
		updateNotification,
		b.Id, b.Title, Types[Email], b.Template, usr, psw, server, port, b.Description,
	).Scan(&id); err != nil {
		return err
	}
	return nil
}

func (acc EmailAccessor) GetNotification(id int) (EmailNotification, error) {
	//defer acc.Conn.getConnection().Close()
	notification := EmailNotification{}
	var records []*DbENRecord
	ctx := context.Background()
	conn, err := acc.Conn.getConnection()
	if err != nil {
		return notification, err
	}
	defer conn.Release()
	err = pgxscan.Select(ctx, conn, &records, selectNotification, id)
	if err != nil {
		return notification, err
	}
	records[0].SetConfigToNotification(&notification)
	return notification, err
}

func (acc EmailAccessor) GetNotifications() ([]*DbENRecord, error) {
	var records []*DbENRecord
	ctx := context.Background()
	conn, err := acc.Conn.getConnection()
	if err != nil {
		return records, err
	}
	err = pgxscan.Select(ctx, conn, &records, selectNotifications)
	return records, err
}

func (acc EmailAccessor) deleteNotification() error {
	conn, err := acc.Conn.getConnection()
	if err != nil {
		return err
	}
	defer conn.Release()
	commandTag, err := conn.Exec(
		context.Background(),
		deleteNotification,
		"email",
	)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return errors.New("no row found to delete")
	}
	return nil
}
