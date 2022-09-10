package main

import (
	//"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	//"net"
	"encoding/json"
	"os/exec"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

/*
*	INFORMACION PARA LA CONEXION A LA BASE DE DATOS
*
 */
var db *sql.DB
var db1 *sql.DB

const (
	host     = "108.59.80.233"
	port     = 3306 // Default port
	user     = "root"
	password = "root"
	dbname   = "db_t1"
)

type Student struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Process struct {
	Pid   string    `json:"pid,omitempty"`
	Comm  string    `json:"comm,omitempty"`
	State string    `json:"state,omitempty"`
	Owner string    `json:"owner,omitempty"`
	Child []Process `json:"child,omitempty"`
}

type Students struct {
	Students []Student `json:"students,omitempty"`
}

func Connect() error {
	var err error
	// Use DSN string to open
	fmt.Printf("%s:%s@%s(%s:%s)/%s", user, password, "tcp", host, strconv.Itoa(port), dbname)
	fmt.Println("\n//")
	//db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", user, password, dbname))
	//fmt.Printf("%s:%s@%s(%s:%s)/%s", user, password, "tcp", host, strconv.Itoa(port), dbname)
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@%s(%s:%s)/%s", user, password, "tcp", host, strconv.Itoa(port), dbname))
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}

func getCPU(c *fiber.Ctx) error {
	fmt.Println("DATOS OBTENIDOS DESDE EL MODULO :")
	fmt.Println("")

	cmd := exec.Command("sh", "-c", "cat /proc/cpu201901849")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	output := string(out[:])
	correct_out := strings.Replace(output, "},\n]", "}]", -1)
	correct_bytes := []byte(correct_out)
	var proc []Process
	err = json.Unmarshal(correct_bytes, &proc)
	if err != nil {
		return c.JSON(err.Error())
	}

	//fmt.Println(output)
	return c.JSON(proc)
}

func getUsers(c *fiber.Ctx) error {
	fmt.Println("Hola test")
	rows, err := db.Query("SELECT * from students")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	result := Students{}

	for rows.Next() {
		user := Student{}
		if err := rows.Scan(&user.Id, &user.Name); err != nil {
			return err
		}

		result.Students = append(result.Students, user)
	}
	fmt.Println(result.Students)
	return c.JSON(result)
}

func insertUser(c *fiber.Ctx) error {
	u := new(Student)
	if err := c.BodyParser(u); err != nil {
		fmt.Println("Es un error")
		return c.Status(400).SendString(err.Error())
	}
	fmt.Println("el u es")
	fmt.Println(u)
	res, err := db.Query("INSERT INTO students (ID, NAME)  VALUES (?,?)", u.Id, u.Name)
	if err != nil {
		return err
	}
	log.Println(res)
	return c.JSON("Operacion Completada")
}

func main() {
	if err := Connect(); err != nil {
		log.Fatal(err)
	}
	var err error
	//db1, err = connectWithConnector()

	if err != nil {
		log.Fatal(err)
	}
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Get("/getCPU", getUsers)
	app.Get("/readModCPU", getCPU)
	app.Post("/readCPU", insertUser)
	log.Fatal(app.Listen(":4000"))
}
