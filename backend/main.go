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
	"time"

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

type RAMRec struct {
	ID      int    `json:"id,omitempty"`
	Fecha   string `json:"fecha,omitempty"`
	Libre   int    `json:"libre,omitempty"`
	Ocupado int    `json:"ocupado,omitempty"`
}

type Student struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Datasys struct {
	Process []Process `json:"process,omitempty"`
	Ram     DataUse   `json:"ram,omitempty"`
	Cpu     DataUse   `json:"cpu,omitempty"`
}

type DataUse struct {
	Free     float64 `json:"free"`
	Occupied float64 `json:"occupied"`
	Aux      string  `json:"aux"`
}

type Process struct {
	Pid   string    `json:"pid,omitempty"`
	Comm  string    `json:"comm,omitempty"`
	State string    `json:"state,omitempty"`
	Owner string    `json:"owner,omitempty"`
	Child []Process `json:"child,omitempty"`
}

type RAM struct {
	Free  int `json:"free"`
	Total int `json:"total"`
}

type Students struct {
	Students []Student `json:"students,omitempty"`
}

type RAMS struct {
	Ram []RAMRec `json:"ram,omitempty"`
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

func getDataSys(c *fiber.Ctx) error {
	procArray := getProcess()
	ram := getRam()
	cpu := getCPU()
	insertCPU(cpu)
	insertRAM(ram)
	dataSys := Datasys{Process: procArray, Ram: ram, Cpu: cpu}
	return c.JSON(dataSys)
}

func getProcess() []Process {
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
	return proc
}

func getRam() DataUse {
	cmd := exec.Command("sh", "-c", "cat /proc/ram_201901849")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	//output := string(out[:])
	var ram RAM
	err = json.Unmarshal(out, &ram)
	if err != nil {
		return DataUse{Free: -1, Occupied: -2, Aux: err.Error()}
	}
	perc := float64(ram.Free) / float64(ram.Total)
	occupied := (1 - perc) * 100
	free := perc * 100
	s := fmt.Sprintf("%f", perc)
	return DataUse{Free: free, Occupied: occupied, Aux: s}
}

func getCPU() DataUse {
	cmd := exec.Command("sh", "-c", "top -b -n1 | tee aorpprkd004.out | grep 'Cpu(s):'")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	output := string(out[:])
	cpu := ""
	index := 9
	if output[9] == ' ' {
		index++
	}
	for i := index; output[i] != ' '; i++ {
		cpu += string(output[i])
	}
	free := 0.0
	occupied, err := strconv.ParseFloat(cpu, 64)
	if err == nil {
		free = 100 - occupied
		return DataUse{Free: free, Occupied: occupied, Aux: output}
	} else {
		return DataUse{Free: -1, Occupied: -2, Aux: err.Error()}
	}

}

func testCPU(c *fiber.Ctx) error {
	return c.JSON(getCPU())
}

func testRAM(c *fiber.Ctx) error {
	return c.JSON(getRam())
}

func testProcess(c *fiber.Ctx) error {
	return c.JSON(getProcess())
}

func getUsers(c *fiber.Ctx) error {
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

func devFecha() string {
	t := time.Now()
	fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	return fecha
}

func createTable() {
	res, err := db.Query("CREATE TABLE Cpu (id INT AUTO_INCREMENT PRIMARY KEY , fecha VARCHAR(40) , libre INT , ocupado INT);")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

func insertRAM(u DataUse) error {
	/*u := new(DataUse)
	if err := c.BodyParser(u); err != nil {
		fmt.Println("Es un error")
		return c.Status(400).SendString(err.Error())
	}*/
	res, err := db.Query("INSERT INTO Ram (FECHA, LIBRE, OCUPADO)  VALUES (?,?,?)", devFecha(), u.Free, u.Occupied)
	if err != nil {
		return err
	}
	log.Println(res)
	return nil
}

func insertCPU(u DataUse) error {
	/*u := new(DataUse)
	if err := c.BodyParser(u); err != nil {
		fmt.Println("Es un error")
		return c.Status(400).SendString(err.Error())
	}*/
	res, err := db.Query("INSERT INTO Cpu (FECHA, LIBRE, OCUPADO)  VALUES (?,?,?)", devFecha(), u.Free, u.Occupied)
	if err != nil {
		return err
	}
	log.Println(res)
	return nil
}

func selectRAM(c *fiber.Ctx) error {
	rows, err := db.Query("SELECT * from Ram")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	result := RAMS{}

	for rows.Next() {
		user := RAMRec{}
		if err := rows.Scan(&user.ID, &user.Fecha, &user.Libre, &user.Ocupado); err != nil {
			return err
		}
		result.Ram = append(result.Ram, user)
	}
	fmt.Println(result.Ram)
	return c.JSON(result)
}

func selectCPU(c *fiber.Ctx) error {
	rows, err := db.Query("SELECT * from Cpu")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	result := RAMS{}

	for rows.Next() {
		user := RAMRec{}
		if err := rows.Scan(&user.ID, &user.Fecha, &user.Libre, &user.Ocupado); err != nil {
			return err
		}
		result.Ram = append(result.Ram, user)
	}
	fmt.Println(result.Ram)
	return c.JSON(result)
}

func main() {
	if err := Connect(); err != nil {
		log.Fatal(err)
	}
	var err error
	//db1, err = connectWithConnector()
	//createTable()
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
	app.Get("/readModCPU", getDataSys)
	app.Post("/readCPU", insertUser)
	app.Get("/testRAM", testRAM)
	app.Get("/testCPU", testCPU)
	app.Get("/testProcess", testProcess)
	//app.Post("/insertRAM", insertRAM)
	//app.Post("/insertCPU", insertCPU)
	app.Get("/getRAM", selectRAM)
	app.Get("/getCPU2", selectCPU)
	log.Fatal(app.Listen(":4000"))
}
