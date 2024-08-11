package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)



type Storage interface {
	  CreateAccount(*Account) error
	  DeleteAccount(int)error
	  UpdateAccount(*Account)error
	  GetAccountById(int)( *Account,error)
     GetAccounts() ([]*Account, error)
}


type PostgresStore struct {
	  db *sql.DB
}

func NewPostgresStore()(*PostgresStore , error){
   conStr := "postgresql://rssagg_owner:ydT13IVBALuU@ep-still-meadow-a1hc2aci.ap-southeast-1.aws.neon.tech/rssagg?sslmode=require"

   db,err := sql.Open("postgres" , conStr)

   if err!=nil {
   	  return nil,err
   }
   
   if err:=db.Ping();err!=nil{
   	  return nil,err
   }
    
   return &PostgresStore{
   	  db:db,
   },nil
}




func (s *PostgresStore) Init() error{
 return s.createAccountTable()
}

func (s *PostgresStore) createAccountTable() error{
   
   query := `CREATE TABLE IF NOT EXISTS Account(
     
     id SERIAL PRIMARY KEY,
     first_name VARCHAR(100),
     last_name VARCHAR(100),
     number SERIAL,
     balance SERIAL,
     created_at TIMESTAMP
    )` 
   
   _ , err := s.db.Exec(query)
   return err
}



func (s *PostgresStore) CreateAccount( acc *Account)error {
   
    query := `INSERT INTO account (first_name , last_name , number , balance , created_at)  VALUES ($1 , $2, $3 , $4 , $5)`
    resp , err := s.db.Query(query , acc.FirstName , acc.LastName , acc.Number , acc.Balance ,acc.CreatedAt)
    
    if err!=nil {
       return err
    }
    fmt.Printf("%+v" , resp)
	 
    return nil
}


func (s *PostgresStore) GetAccounts() ([]*Account ,error ) {
   
   rows ,err := s.db.Query("SELECT * FROM ACCOUNT")
   if err!=nil {
        return  nil , err
   }

    accounts :=  []*Account{} 
   
   for rows.Next() {
       account :=  &Account{}
       if err := rows.Scan(&account.ID ,&account.FirstName ,&account.LastName , &account.Number , &account.Balance , &account.CreatedAt ); err!= nil {
           return nil ,err
       }
       accounts = append(accounts, account)
   }
   return accounts,nil

   
}

func (s *PostgresStore) DeleteAccount( id int)error {
	 return nil
}
func (s *PostgresStore) UpdateAccount(acc *Account)error {
	 return nil
}
func (s *PostgresStore) GetAccountById(id int)(*Account , error) {
	 return nil,nil
}


var st ,e = NewPostgresStore()