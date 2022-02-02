package database

import (
	"context"
	"reflect"
	"testing"

	db "firebase.google.com/go/v4/db"
)

func TestSetup(t *testing.T) {
	tests := []struct {
		name string
		want *Instance
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Setup("" , ""); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Setup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_Init(t *testing.T) {
	type fields struct {
		Email    string
		Name     string
		Password string
		Username string
	}
	type args struct {
		email    string
		name     string
		username string
		password string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				Email:    tt.fields.Email,
				Name:     tt.fields.Name,
				Password: tt.fields.Password,
				Username: tt.fields.Username,
			}
			a.Init(tt.args.email, tt.args.name, tt.args.username, tt.args.password)
		})
	}
}

func TestAccount_Key(t *testing.T) {
	type fields struct {
		Email    string
		Name     string
		Password string
		Username string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				Email:    tt.fields.Email,
				Name:     tt.fields.Name,
				Password: tt.fields.Password,
				Username: tt.fields.Username,
			}
			if got := a.Key(); got != tt.want {
				t.Errorf("Account.Key() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_GetPath(t *testing.T) {
	type fields struct {
		Email    string
		Name     string
		Password string
		Username string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				Email:    tt.fields.Email,
				Name:     tt.fields.Name,
				Password: tt.fields.Password,
				Username: tt.fields.Username,
			}
			if got := a.GetPath(); got != tt.want {
				t.Errorf("Account.GetPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHash(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hash(tt.args.s); got != tt.want {
				t.Errorf("Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashAsString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HashAsString(tt.args.s); got != tt.want {
				t.Errorf("HashAsString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInstance_Init(t *testing.T) {
	type fields struct {
		database *db.Client
		ctx      context.Context
	}
	type args struct {
		ctx      context.Context
		database *db.Client
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Instance{
				database: tt.fields.database,
				ctx:      tt.fields.ctx,
			}
			i.Init(tt.args.ctx, tt.args.database)
		})
	}
}

func TestInstance_PostUserInfo(t *testing.T) {
	type fields struct {
		database *db.Client
		ctx      context.Context
	}
	type args struct {
		acc *Account
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
			i := &Instance{
				database: tt.fields.database,
				ctx:      tt.fields.ctx,
			}
			if err := i.PostUserInfo(tt.args.acc); (err != nil) != tt.wantErr {
				t.Errorf("Instance.PostUserInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInstance_GetUserInfo(t *testing.T) {
	type fields struct {
		database *db.Client
		ctx      context.Context
	}
	type args struct {
		userID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantAcc *Account
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Instance{
				database: tt.fields.database,
				ctx:      tt.fields.ctx,
			}
			gotAcc, err := i.GetUserInfo(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Instance.GetUserInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotAcc, tt.wantAcc) {
				t.Errorf("Instance.GetUserInfo() = %v, want %v", gotAcc, tt.wantAcc)
			}
		})
	}
}

func TestInstance_UpdateUserInfo(t *testing.T) {
	type fields struct {
		database *db.Client
		ctx      context.Context
	}
	type args struct {
		userID string
		newAcc *NilAccount
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
			i := &Instance{
				database: tt.fields.database,
				ctx:      tt.fields.ctx,
			}
			if err := i.UpdateUserInfo(tt.args.userID, tt.args.newAcc); (err != nil) != tt.wantErr {
				t.Errorf("Instance.UpdateUserInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInstance_AuthenticateUserInfo(t *testing.T) {
	type fields struct {
		database *db.Client
		ctx      context.Context
	}
	type args struct {
		form *AuthenticatorForm
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Instance{
				database: tt.fields.database,
				ctx:      tt.fields.ctx,
			}
			got, err := i.AuthenticateUserInfo(tt.args.form)
			if (err != nil) != tt.wantErr {
				t.Errorf("Instance.AuthenticateUserInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Instance.AuthenticateUserInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
