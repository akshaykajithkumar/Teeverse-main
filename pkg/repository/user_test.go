package repository

import (
	"errors"
	"log"
	"reflect"
	"testing"

	"Teeverse/pkg/domain"
	"Teeverse/pkg/utils/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-playground/assert/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_CheckUserAvailabilty(t *testing.T) {
	tests := []struct {
		name string
		args string
		stub func(sqlmock.Sqlmock)
		want bool
	}{
		{
			name: "User Available",
			args: "akshay@gmail.com",
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^select count\(\*\) from users(.+)$`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(5))

			},
			want: true,
		},
		{
			name: "user not available",
			args: "akshay@gmail.com",
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^select count\(\*\) from users(.+)$`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

			},

			want: false,
		},
		{
			name: "error from database",
			args: "akshay@gmail.com",
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^select count\(\*\) from users(.+)$`

				mockSQL.ExpectQuery(expectedQuery).WithArgs("").
					WillReturnError(errors.New("text string"))

			},

			want: false,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)

			result := u.CheckUserAvailability(tt.args)

			assert.Equal(t, tt.want, result)
		})
	}

}

func Test_UserBlockStatus(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		stub    func(sqlmock.Sqlmock)
		want    bool
		wantErr error
	}{
		{
			name: "User is blocked",
			args: "akshay@gmail.com",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `^select permission from users(.+)$`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnRows(sqlmock.NewRows([]string{"permission"}).AddRow(true))
			},
			want:    true,
			wantErr: nil,
		},
		{
			name: "user is not blocked",
			args: "akshay@gmail.com",
			stub: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `^select permission from users(.+)$`

				mockSQL.ExpectQuery(expectedQuery).WithArgs().WillReturnRows(sqlmock.NewRows([]string{"permission"}).AddRow(false))
			},
			want:    false,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})
			tt.stub(mockSQL)
			u := NewUserRepository(gormDB)
			result, err := u.UserBlockStatus(tt.args)
			assert.Equal(t, result, tt.want)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_SignUp(t *testing.T) {

	type args struct {
		input models.UserDetails
	}

	tests := []struct {
		name    string
		args    args
		stub    func(sqlmock.Sqlmock)
		want    models.UserResponse
		wantErr error
	}{
		{
			name: "success signup user",
			args: args{
				input: models.UserDetails{Name: "UserNew", Email: "usernew@gmail.com", Username: "usernew", Phone: "1234567890", Password: "userpass", ConfirmPassword: "userpass"},
			},

			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^INSERT INTO users (.+)$`
				mockSQL.ExpectQuery(expectedQuery).WithArgs("UserNew", "usernew@gmail.com", "userpass", "1234567890", "usernew").
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "phone"}).AddRow(1, "UserNew", "usernew@gmail.com", "1234567890"))

			},

			want:    models.UserResponse{Id: 1, Name: "UserNew", Email: "usernew@gmail.com", Phone: "1234567890"},
			wantErr: nil,
		},

		{
			name: "error signup user",
			args: args{
				input: models.UserDetails{Name: "", Email: "", Username: "", Phone: "", Password: "", ConfirmPassword: ""},
			},
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^INSERT INTO users (.+)$`
				mockSQL.ExpectQuery(expectedQuery).WithArgs("", "", "", "", "").
					WillReturnError(errors.New("Query 'INSERT INTO users (name, email,phone,password,username) VALUES ($1, $2, $3, $4,$5) RETURNING id, name, email, phone', arguments do not match: argument 4 expected [string - ] does not match actual [string - 12345]"))

			},

			want:    models.UserResponse{},
			wantErr: errors.New("Query 'INSERT INTO users (name, email,phone,password,username) VALUES ($1, $2, $3, $4,$5) RETURNING id, name, email, phone', arguments do not match: argument 4 expected [string - ] does not match actual [string - 12345]"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)

			got, err := u.SignUp(tt.args.input)
			log.Println("log :", tt.name, got, tt.want)
			log.Println("log :", tt.name, tt.wantErr, err)
			assert.Equal(t, tt.wantErr, err)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepo.UserSignUp() = %v, want %v", got, tt.want)
			}
		})
	}

}

func Test_FindUserByEmail(t *testing.T) {
	tests := []struct {
		name    string
		args    models.UserLogin
		stub    func(sqlmock.Sqlmock)
		want    models.UserResponse
		wantErr error
	}{
		{
			name: "Success",
			args: models.UserLogin{
				Email:    "akshay@gmail.com",
				Password: "akshay2000",
			},
			stub: func(s sqlmock.Sqlmock) {
				expectedQuery := `^SELECT .* FROM users (.+)$`
				s.ExpectQuery(expectedQuery).WithArgs("akshay@gmail.com").WillReturnRows(sqlmock.NewRows([]string{"id", "name", "username", "email", "phone", "password"}).AddRow(1, "akshay", "akshay", "akshay@gmail.com", "8593008099", "akshay2000"))
			},
			want: models.UserResponse{Id: 1,
				Name:     "akshay",
				Email:    "akshay@gmail.com",
				Username: "akshay",
				Phone:    "8593008099",
				Password: "akshay2000",
			},
			wantErr: nil,
		},
		{
			name: "failure",
			args: models.UserLogin{
				Email:    "akshay@gmail.com",
				Password: "akshay2000",
			},
			stub: func(s sqlmock.Sqlmock) {
				expectedQuery := `^SELECT \* FROM users(.+)$`

				s.ExpectQuery(expectedQuery).WithArgs("akshay@gmail.com").
					WillReturnError(errors.New("new error"))
			},
			want:    models.UserResponse{},
			wantErr: errors.New("error checking user details"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)
			got, err := u.FindUserByEmail(tt.args)

			log.Println("log :", tt.name, tt.want, got)
			log.Println("log :", tt.name, tt.wantErr, err)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, err, tt.wantErr)
		})

	}

}

func Test_FindUserIDByOrderID(t *testing.T) {
	tests := []struct {
		name    string
		args    int
		stub    func(sqlmock.Sqlmock)
		want    int
		wantErr error
	}{
		{
			name: "Success",
			args: 4,
			stub: func(s sqlmock.Sqlmock) {
				expectedQuery := `^SELECT user_id FROM orders (.+)$`
				s.ExpectQuery(expectedQuery).WithArgs(4).WillReturnRows(sqlmock.NewRows([]string{"user_id"}).AddRow(1))
			},
			want:    1,
			wantErr: nil,
		},
		{
			name: "failure",
			args: 5,
			stub: func(s sqlmock.Sqlmock) {
				expectedQuery := `^SELECT user_id FROM orders(.+)$`

				s.ExpectQuery(expectedQuery).WithArgs(5).
					WillReturnError(errors.New("error checking user details"))
			},
			want:    0,
			wantErr: errors.New("error checking user details"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)
			got, err := u.FindUserIDByOrderID(tt.args)

			log.Println("log :", tt.name, tt.want, got)
			log.Println("log :", tt.name, tt.wantErr, err)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, err, tt.wantErr)
		})

	}

}

func Test_AddAddress(t *testing.T) {

	tests := []struct {
		name string
		args struct {
			id      int
			address models.AddAddress
			result  bool
		}
		stub    func(sqlmock.Sqlmock)
		wantErr error
	}{
		{
			name: "success",
			args: struct {
				id      int
				address models.AddAddress
				result  bool
			}{
				id: 1,
				address: models.AddAddress{
					Name:      "akshay",
					HouseName: "thundathil",
					Street:    "mundathicode",
					City:      "thrissur",
					State:     "kerala",
					Pin:       "680601",
				},
				result: true,
			},
			stub: func(mockSQL sqlmock.Sqlmock) {

				mockSQL.ExpectExec("INSERT INTO addresses").WithArgs(1, "akshay", "thundathil", "mundathicode", "thrissur", "kerala", "680601", true).WillReturnResult(sqlmock.NewResult(1, 1))

			},
			wantErr: nil,
		},
		{
			name: "error",
			args: struct {
				id      int
				address models.AddAddress
				result  bool
			}{
				id: 1,
				address: models.AddAddress{
					Name:      "akshay",
					HouseName: "thundathil",
					Street:    "mundathicode",
					City:      "thrissur",
					State:     "kerala",
					Pin:       "680601",
				},
				result: true,
			},
			stub: func(mockSQL sqlmock.Sqlmock) {

				mockSQL.ExpectExec("INSERT INTO addresses").WithArgs(1, "akshay", "thundathil", "mundathicode", "thrissur", "kerala", "680601", true).WillReturnError(errors.New("could not add address"))

			},
			wantErr: errors.New("could not add address"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)
			u := NewUserRepository(gormDB)

			err := u.AddAddress(1, tt.args.address, true)

			assert.Equal(t, tt.wantErr, err)
		})
	}

}

// fg
func Test_CheckIfFirstAddress(t *testing.T) {

	tests := []struct {
		name string
		args int
		stub func(sqlmock.Sqlmock)
		want bool
	}{
		{
			name: "first address",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^select count\(\*\) from addresses(.+)$`

				mockSQL.ExpectQuery(expectedQuery).WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

			},

			want: false,
		},
		{
			name: "error occured",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^select count\(\*\) from addresses(.+)$`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnError(errors.New("error"))

			},

			want: false,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)

			result := u.CheckIfFirstAddress(tt.args)
			log.Println("log :", tt.name, tt.want, result)
			//log.Println("log :", tt.name, tt.wantErr, err)

			assert.Equal(t, tt.want, result)
		})
	}

}

func Test_GetAddresses(t *testing.T) {

	tests := []struct {
		name    string
		args    int
		stub    func(sqlmock.Sqlmock)
		want    []domain.Address
		wantErr error
	}{
		{
			name: "successfully got all addresses",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^select \* from addresses(.+)$`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "name", "house_name", "street", "city", "state", "pin", "default"}).AddRow(1, 1, "a", "b", "c", "d", "e", "f", true).AddRow(2, 1, "a", "b", "c", "d", "e", "f", false))

			},

			want: []domain.Address{
				{ID: 1,
					UserID:    1,
					Name:      "a",
					HouseName: "b",
					Street:    "c",
					City:      "d",
					State:     "e",
					Pin:       "f",
					Default:   true,
				}, {
					ID:        2,
					UserID:    1,
					Name:      "a",
					HouseName: "b",
					Street:    "c",
					City:      "d",
					State:     "e",
					Pin:       "f",
					Default:   false,
				},
			},
			wantErr: nil,
		},
		{
			name: "error",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^select \* from addresses(.+)$`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnError(errors.New("error in getting addresses"))

			},

			want:    []domain.Address{},
			wantErr: errors.New("error in getting addresses"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)

			result, err := u.GetAddresses(tt.args)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, result)
		})
	}

}
func Test_GetUserDetails(t *testing.T) {

	tests := []struct {
		name    string
		args    int
		stub    func(sqlmock.Sqlmock)
		want    models.UserResponse
		wantErr error
	}{
		{
			name: "successfully got details",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				// expectedQuery := `^select \* from users(.+)$`,

				mockSQL.ExpectQuery(`^select id\,name\,username\,email\,phone from users(.+)$`).WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "username", "email", "phone"}).AddRow(1, "akshay", "akshay", "akshay@gmail.com", "8593098099"))
			},

			want: models.UserResponse{
				Id:       1,
				Name:     "akshay",
				Username: "akshay",
				Email:    "akshay@gmail.com",
				Phone:    "8593098099",
			},
			wantErr: nil,
		},
		{
			name: "error",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				mockSQL.ExpectQuery(`^select id\,name\,username\,email\,phone from users(.+)$`).
					WillReturnError(errors.New("could not get user details"))
			},

			want:    models.UserResponse{},
			wantErr: errors.New("could not get user details"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)

			result, err := u.GetUserDetails(tt.args)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, result)
		})
	}

}

func Test_ChangePassword(t *testing.T) {
	tests := []struct {
		name string
		args struct {
			id       int
			password string
		}
		stub    func(sqlmock.Sqlmock)
		wantErr error
	}{
		{
			name: "success",
			args: struct {
				id       int
				password string
			}{id: 1, password: "newpass"},
			stub: func(mockSQL sqlmock.Sqlmock) {

				mockSQL.ExpectExec("UPDATE users SET").WithArgs().WillReturnResult(sqlmock.NewResult(1, 1))

			},
			wantErr: nil,
		},
		{
			name: "error",
			args: struct {
				id       int
				password string
			}{id: 1, password: "nil"},
			stub: func(mockSQL sqlmock.Sqlmock) {

				mockSQL.ExpectExec("UPDATE users SET").WithArgs().WillReturnError(errors.New("couldnt change password"))

			},
			wantErr: errors.New("couldnt change password"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)
			u := NewUserRepository(gormDB)

			err := u.ChangePassword(tt.args.id, tt.args.password)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_FindIdFromPhone(t *testing.T) {

	tests := []struct {
		name    string
		args    string
		stub    func(sqlmock.Sqlmock)
		want    int
		wantErr error
	}{
		{
			name: "success",
			args: "1234567890",
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^select id from users(.+)$`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

			},

			want:    1,
			wantErr: nil,
		},
		{
			name: "failure",
			args: "1234567890",
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^select id from users(.+)$`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnError(errors.New("error"))

			},

			want:    0,
			wantErr: errors.New("error"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)

			result, err := u.FindIdFromPhone(tt.args)

			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}

}

func Test_EditName(t *testing.T) {
	tests := []struct {
		name string
		args struct {
			id   int
			name string
		}
		stub    func(sqlmock.Sqlmock)
		wantErr error
	}{
		{
			name: "success",
			args: struct {
				id   int
				name string
			}{id: 1, name: "akshay"},
			stub: func(mockSQL sqlmock.Sqlmock) {

				mockSQL.ExpectExec("update users set").WithArgs().WillReturnResult(sqlmock.NewResult(1, 1))

			},
			wantErr: nil,
		},
		{
			name: "error",
			args: struct {
				id   int
				name string
			}{id: 1, name: "akshay"},
			stub: func(mockSQL sqlmock.Sqlmock) {

				mockSQL.ExpectExec("update users set").WithArgs().WillReturnError(errors.New("error"))

			},
			wantErr: errors.New("error"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)
			u := NewUserRepository(gormDB)

			err := u.EditName(tt.args.id, tt.args.name)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_EditEmail(t *testing.T) {
	tests := []struct {
		name string
		args struct {
			id    int
			email string
		}
		stub    func(sqlmock.Sqlmock)
		wantErr error
	}{
		{
			name: "success",
			args: struct {
				id    int
				email string
			}{id: 1, email: "akshay@gmail.com"},
			stub: func(mockSQL sqlmock.Sqlmock) {

				mockSQL.ExpectExec("update users set").WithArgs().WillReturnResult(sqlmock.NewResult(1, 1))

			},
			wantErr: nil,
		},
		{
			name: "error",
			args: struct {
				id    int
				email string
			}{id: 1, email: "akshay@gmail.com"},
			stub: func(mockSQL sqlmock.Sqlmock) {

				mockSQL.ExpectExec("update users set").WithArgs().WillReturnError(errors.New("error"))

			},
			wantErr: errors.New("error"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)
			u := NewUserRepository(gormDB)

			err := u.EditEmail(tt.args.id, tt.args.email)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_EditUserame(t *testing.T) {
	tests := []struct {
		name string
		args struct {
			id       int
			username string
		}
		stub    func(sqlmock.Sqlmock)
		wantErr error
	}{
		{
			name: "success",
			args: struct {
				id       int
				username string
			}{id: 1, username: "akshay"},
			stub: func(mockSQL sqlmock.Sqlmock) {

				mockSQL.ExpectExec("update users set").WithArgs().WillReturnResult(sqlmock.NewResult(1, 1))

			},
			wantErr: nil,
		},
		{
			name: "error",
			args: struct {
				id       int
				username string
			}{id: 1, username: "akshay"},
			stub: func(mockSQL sqlmock.Sqlmock) {

				mockSQL.ExpectExec("update users set").WithArgs().WillReturnError(errors.New("error"))

			},
			wantErr: errors.New("error"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)
			u := NewUserRepository(gormDB)

			err := u.EditUsername(tt.args.id, tt.args.username)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_EditPhone(t *testing.T) {
	tests := []struct {
		name string
		args struct {
			id    int
			phone string
		}
		stub    func(sqlmock.Sqlmock)
		wantErr error
	}{
		{
			name: "success",
			args: struct {
				id    int
				phone string
			}{id: 1, phone: "1234567890"},
			stub: func(mockSQL sqlmock.Sqlmock) {

				mockSQL.ExpectExec("update users set").WithArgs().WillReturnResult(sqlmock.NewResult(1, 1))

			},
			wantErr: nil,
		},
		{
			name: "error",
			args: struct {
				id    int
				phone string
			}{id: 1, phone: "1234567890"},
			stub: func(mockSQL sqlmock.Sqlmock) {

				mockSQL.ExpectExec("update users set").WithArgs().WillReturnError(errors.New("error"))

			},
			wantErr: errors.New("error"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)
			u := NewUserRepository(gormDB)

			err := u.EditPhone(tt.args.id, tt.args.phone)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_RemoveFromCart(t *testing.T) {

	tests := []struct {
		name string
		args struct {
			cartId int
			invId  int
		}
		stub    func(sqlmock.Sqlmock)
		wantErr error
	}{
		{
			name: "success",
			args: struct {
				cartId int
				invId  int
			}{
				cartId: 1,
				invId:  1,
			},
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `delete from line_items`

				mockSQL.ExpectExec(expectedQuery).
					WillReturnResult(sqlmock.NewResult(1, 1))

			},
			wantErr: nil,
		},
		{
			name: "error",
			args: struct {
				cartId int
				invId  int
			}{
				cartId: 1,
				invId:  1,
			},
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `delete from line_items`

				mockSQL.ExpectExec(expectedQuery).
					WillReturnError(errors.New("error"))

			},
			wantErr: errors.New("error"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)

			err := u.RemoveFromCart(tt.args.cartId, tt.args.invId)

			assert.Equal(t, tt.wantErr, err)
		})
	}

}

func Test_ClearCart(t *testing.T) {

	tests := []struct {
		name    string
		args    int
		stub    func(sqlmock.Sqlmock)
		wantErr error
	}{
		{
			name: "success",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `delete from line_items`

				mockSQL.ExpectExec(expectedQuery).
					WillReturnResult(sqlmock.NewResult(1, 1))

			},
			wantErr: nil,
		},
		{
			name: "error",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `delete from line_items`

				mockSQL.ExpectExec(expectedQuery).
					WillReturnError(errors.New("error"))

			},
			wantErr: errors.New("error"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)

			err := u.ClearCart(tt.args)

			assert.Equal(t, tt.wantErr, err)
		})
	}

}

func Test_GetCartID(t *testing.T) {

	tests := []struct {
		name    string
		args    int
		stub    func(sqlmock.Sqlmock)
		want    int
		wantErr error
	}{
		{
			name: "success",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `select id from carts`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

			},
			want:    1,
			wantErr: nil,
		},
		{
			name: "error",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `select id from carts`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnError(errors.New("error"))

			},
			want:    0,
			wantErr: errors.New("error"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)

			result, err := u.GetCartID(tt.args)

			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}

}

func Test_GetProductsInCart(t *testing.T) {

	tests := []struct {
		name string
		args struct {
			cart_id int
			page    int
			limit   int
		}
		stub    func(sqlmock.Sqlmock)
		want    []int
		wantErr error
	}{
		{
			name: "success",
			args: struct {
				cart_id int
				page    int
				limit   int
			}{
				cart_id: 1, page: 1, limit: 2,
			},
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `select inventory_id from line_items`

				mockSQL.ExpectQuery(expectedQuery).WithArgs().
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1).AddRow(2))

			},
			want:    []int{1, 2},
			wantErr: nil,
		},
		{
			name: "error",
			args: struct {
				cart_id int
				page    int
				limit   int
			}{
				cart_id: 1, page: 1, limit: 2,
			},
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `select inventory_id from line_items`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnError(errors.New("error"))

			},
			want:    []int{},
			wantErr: errors.New("error"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)

			result, err := u.GetProductsInCart(tt.args.cart_id, tt.args.page, tt.args.limit)

			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}

}

func Test_FindProductsNames(t *testing.T) {

	tests := []struct {
		name    string
		args    int
		stub    func(sqlmock.Sqlmock)
		want    string
		wantErr error
	}{
		{
			name: "success",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `select product_name from inventories`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnRows(sqlmock.NewRows([]string{"product_name"}).AddRow("vintage product"))

			},
			want:    "vintage product",
			wantErr: nil,
		},
		{
			name: "error",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `select product_name from inventories`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnError(errors.New("error"))

			},
			want:    "",
			wantErr: errors.New("error"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)

			result, err := u.FindProductNames(tt.args)

			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}

}

func Test_FindCartQuantity(t *testing.T) {

	tests := []struct {
		name string
		args struct {
			cart_id int
			inv_id  int
		}
		stub    func(sqlmock.Sqlmock)
		want    int
		wantErr error
	}{
		{
			name: "success",
			args: struct {
				cart_id int
				inv_id  int
			}{
				cart_id: 1,
				inv_id:  1,
			},
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `select quantity from line_items`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnRows(sqlmock.NewRows([]string{"quantity"}).AddRow(1))

			},
			want:    1,
			wantErr: nil,
		},
		{
			name: "error",
			args: struct {
				cart_id int
				inv_id  int
			}{
				cart_id: 1,
				inv_id:  1,
			},
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `select quantity from line_items`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnError(errors.New("error"))

			},
			want:    0,
			wantErr: errors.New("error"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)

			result, err := u.FindCartQuantity(tt.args.cart_id, tt.args.inv_id)

			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}

}

func Test_FindPrice(t *testing.T) {

	tests := []struct {
		name    string
		args    int
		stub    func(sqlmock.Sqlmock)
		want    float64
		wantErr error
	}{
		{
			name: "success",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `select price from inventories`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnRows(sqlmock.NewRows([]string{"price"}).AddRow(400))

			},
			want:    400,
			wantErr: nil,
		},
		{
			name: "error",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `select price from inventories`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnError(errors.New("error"))

			},
			want:    0,
			wantErr: errors.New("error"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)

			result, err := u.FindPrice(tt.args)

			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}

}
