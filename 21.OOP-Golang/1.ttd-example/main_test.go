package main

import (
	"database/sql"
	"reflect"
	"testing"
)

func TestMessageService_SendMessage(t *testing.T) {
	type fields struct {
		storage MessageStorage
	}
	type args struct {
		senderID   int
		receiverID int
		text       string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := MessageService{
				storage: tt.fields.storage,
			}
			if err := s.SendMessage(tt.args.senderID, tt.args.receiverID, tt.args.text); (err != nil) != tt.wantErr {
				t.Errorf("SendMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSQLiteMessageStorage_GetMessages(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		userID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Message
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SQLiteMessageStorage{
				db: tt.fields.db,
			}
			got, err := s.GetMessages(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMessages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMessages() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSQLiteMessageStorage_SaveMessage(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		message Message
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SQLiteMessageStorage{
				db: tt.fields.db,
			}
			if err := s.SaveMessage(tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("SaveMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_createMessagesTable(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := createMessagesTable(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("createMessagesTable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
