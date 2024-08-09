package main

type storage interface {
	  CreateAccount(*Account) error
	  DeleteAccount(int)error
	  updateAccount(*Account)error
	  GetAccountById(int)( *Account,error)
}