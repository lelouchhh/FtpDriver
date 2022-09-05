package main

import (
	"crypto/tls"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"gopkg.in/gomail.v2"
	"log"
)

const (
	DATABASE = "dungeon"
	PASSWORD = ""
	USER     = "slave"
	HOST     = "185.200.241.2"
)

func main() {
	var email string
	var msg string
	flag.StringVar(&email, "email", "llchh@yahoo.com", "user email")
	flag.StringVar(&msg, "msg", "hash", "message type")
	flag.Parse()
	fmt.Println(email)
	db, err := DbConnect()
	if err != nil {
		fmt.Println(err)
	}
	hash := GetHash(db, email)
	login := GetUser(db, email)
	if msg == "pass" {
	} else {
		SendHashEmail(email, hash, login)

	}
}
func GetHash(db *sql.DB, email string) (hash string) {
	userSql := "SELECT email_pass FROM auth.t_user WHERE email = $1"

	err := db.QueryRow(userSql, email).Scan(&hash)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}
	return hash
}
func DbConnect() (*sql.DB, error) {
	var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=require", HOST, USER, PASSWORD, DATABASE)

	// Initialize connection object.
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("err")
		return nil, err
	}
	return db, err
}
func GetUser(db *sql.DB, email string) (login string) {
	userSql := "SELECT login FROM auth.t_user WHERE email = $1"

	err := db.QueryRow(userSql, email).Scan(&login)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}
	return login
}
func SendHashEmail(to, hash, login string) {
	m := gomail.NewMessage()

	m.SetHeader("From", "info@xn--80agm.com")
	fmt.Println("to: ", to)
	m.SetHeader("To", to)

	m.SetHeader("Subject", "Подтверждение почты!")

	m.SetBody("text/plain", fmt.Sprintf("Ваша ссылка для подтверждения почты!: %s", fmt.Sprintf("https://агз.com/ref/%s/%s", login, hash)))

	d := gomail.NewDialer("mail.smtp2go.com", 2525, "xn--80agm.com", "dvK6m2lWfKjZ6E8c")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return
}
func SendPassEmail() {

}
