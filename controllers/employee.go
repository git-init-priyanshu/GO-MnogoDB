package controllers

import (
	"context"
	"go-MongoDB/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeRepo struct {
	MongoCollection *mongo.Collection
}

func (r *EmployeeRepo) FindByEmployeeId(empId string) (*models.Employee, error) {
	var emp models.Employee

	if err := r.MongoCollection.FindOne(
		context.Background(),
		bson.D{{Key: "employee_id", Value: empId}}).Decode(&emp); err != nil {
		return nil, err
	}

	return &emp, nil
}

func (r *EmployeeRepo) FindAllEmployee() ([]models.Employee, error) {
	res, err := r.MongoCollection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	var emps []models.Employee
	if err := res.All(context.Background(), &emps); err != nil {
		return nil, err
	}

	return emps, err
}

func (r *EmployeeRepo) InsertEmployee(emp *models.Employee) (interface{}, error) {
	res, err := r.MongoCollection.InsertOne(context.Background(), emp)
	if err != nil {
		return nil, err
	}

	return res.InsertedID, nil
}

func (r *EmployeeRepo) UpdateEmployeeById(empId string, emp *models.Employee) (int64, error) {
	res, err := r.MongoCollection.UpdateByID(
		context.Background(),
		bson.D{{Key: "employee_id", Value: empId}},
		bson.D{{Key: "$set", Value: emp}})
	if err != nil {
		return 0, err
	}

	return res.ModifiedCount, nil
}

func (r *EmployeeRepo) DeleteEmployeeById(empId string) (int64, error) {
	res, err := r.MongoCollection.DeleteOne(
		context.Background(),
		bson.D{{Key: "employee_id", Value: empId}})
	if err != nil {
		return 0, err
	}

	return res.DeletedCount, nil
}

func (r *EmployeeRepo) DeleteAllEmployee() (int64, error) {
	res, err := r.MongoCollection.DeleteMany(context.Background(), bson.D{})
	if err != nil {
		return 0, err
	}

	return res.DeletedCount, nil
}
