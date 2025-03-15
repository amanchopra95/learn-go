package main

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

type AuditLogView struct {
	EntityId          string
	CreatedAt         any
	OperationDateTime any
	Details           Detail
	UserContext       UserContext
}

type ChangeLog struct {
	OldValue string
	NewValue string
}

type Detail struct {
	Status ChangeLog
}

type UserContext struct {
	FirstName string
	LastName  string
	UserRole  string
}

type BookingRequestModel struct {
	Id                int64
	AccountName       string
	Status            string
	BillToInfoId      int64
	ServiceTypeIds    any
	PassengerName     string
	PassengerId       int64
	PassengerDOB      sql.NullString
	PassengerGender   sql.NullString
	Payer             string
	TripPurpose       sql.NullString
	Notes             sql.NullString
	VehicleCategoryId int64
	UserName          sql.NullString
}

var db *sql.DB

func Connect() {
	connStr, status := os.LookupEnv("CONN_STR")
	if !status {
		log.Fatalln("Missing environment variable to connect to database")
	}

	var err error
	db, err = sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database %s", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Unable to ping the database %s", err)
	}

	log.Println("Database connected")
}

func ReadJsonFile() []AuditLogView {
	jsonFile, err := os.Open(os.Getenv("FILE_PATH"))
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)

	var auditLogViews []AuditLogView
	json.Unmarshal(byteValue, &auditLogViews)

	return auditLogViews
}

func WriteCSVFile(payload []BookingRequestModel) {
	file, err := os.Create("records.csv")
	if err != nil {
		log.Fatal("Failed to open a file", err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	records := make([][]string, len(payload)+1)
	records[0] = []string{"Id", "AccountName", "Status", "BillToInfoId", "ServiceTypeIds", "PassengerName", "PassengerId", "PassengerDOB", "PassengerGender", "Payer", "TripPurpose", "Notes", "VehicleCategoryId", "UserName"}

	for i, record := range payload {
		records[i+1] = []string{
			strconv.Itoa(int(record.Id)),
			record.AccountName,
			record.Status,
			strconv.Itoa(int(record.BillToInfoId)),
			fmt.Sprint(record.ServiceTypeIds),
			record.PassengerName,
			strconv.Itoa(int(record.PassengerId)),
			record.PassengerDOB.String,
			record.PassengerGender.String,
			record.Payer,
			record.TripPurpose.String,
			record.Notes.String,
			strconv.Itoa(int(record.VehicleCategoryId)),
			record.UserName.String,
		}
	}

	w.WriteAll(records)
}

func QuerySQLData(payload []AuditLogView) []BookingRequestModel {
	var entityIds []int64
	var ids []string
	for i := range payload {
		id, err := strconv.ParseInt(payload[i].EntityId, 10, 64)
		if err != nil {
			log.Println("Could not parse the EntityId", payload[i].EntityId)
			continue
		}
		entityIds = append(entityIds, id)
		ids = append(ids, payload[i].EntityId)
	}

	stmt, err := db.Prepare(`
		SELECT br."Id", br."AccountName", br."BillToInfoId", br."PassengerId", br."PassengerName", br."ServiceTypeIds", lp."Name" as "Status",
		pa."DateOfBirth", pa."Gender", no."Notes", tp."Name" as "TripPurpose", bi."Name" as "Payer", br."VehicleCategoryId", st."Name" as "UserName"
		FROM mekong_trip."BookingRequest" br
		join mekong_trip."Passenger" pa on pa."Id" = br."PassengerId"
		left join mekong_trip."Notes" no on no."BookingRequestId" = br."Id" and no."NotesTypeId" = 1
		left join mekong_trip."TripPurpose" tp on tp."Id" = br."TripPurposeId"
		join mekong_trip."BillingInfo" bi on bi."Id" = br."BillToInfoId"
		left join mekong_fna."Staff" st on st."SystemUserId" = br."updatedById"
		left join mekong_trip."Lookup" lp on lp."Id" = br."StatusId"
		WHERE br."Id" = ANY($1);
	`)

	if err != nil {
		log.Fatalln("Could not generate SQL statement", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(entityIds)

	if err != nil {
		log.Printf("Could not fetch data from DB %s", err)
	}
	defer rows.Close()

	bookings := make([]BookingRequestModel, 0)
	for rows.Next() {
		var b BookingRequestModel
		err := rows.Scan(&b.Id, &b.AccountName, &b.BillToInfoId, &b.PassengerId, &b.PassengerName, &b.ServiceTypeIds, &b.Status, &b.PassengerDOB, &b.PassengerGender, &b.Notes, &b.TripPurpose, &b.Payer, &b.VehicleCategoryId, &b.UserName)

		if err != nil {
			log.Println("Error in scanning the rows", err)
		}

		bookings = append(bookings, b)
	}

	return bookings
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load .env")
	}
	// Read JSON file from a given path
	var auditLogViews []AuditLogView = ReadJsonFile()
	// loop through entities and query the data from DB
	Connect()
	bookings := QuerySQLData(auditLogViews)
	fmt.Println("Bookings returned ", len(bookings))
	// write the data in excel
	WriteCSVFile(bookings)

}
