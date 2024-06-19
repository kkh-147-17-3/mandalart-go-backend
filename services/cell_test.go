package services

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	repo "mandalart.com/repositories"
)

// Mock the DBTX type
type MockDBTX struct {
	mock.Mock
}

func (m *MockDBTX) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	argsArray := m.Called(ctx, query, args)
	return argsArray.Get(0).(pgconn.CommandTag), argsArray.Error(1)
}

func (m *MockDBTX) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	argsArray := m.Called(ctx, query, args)
	return argsArray.Get(0).(pgx.Rows), argsArray.Error(1)
}

func (m *MockDBTX) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	argsArray := m.Called(ctx, query, args)
	return argsArray.Get(0).(pgx.Row)
}

func (m *MockDBTX) CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	argsArray := m.Called(ctx, tableName, columnNames, rowSrc)
	return argsArray.Get(0).(int64), argsArray.Error(1)
}

// Mock the Queries type
type MockQueries struct {
	*repo.Queries
	mock.Mock
}

func NewMockQueries() *MockQueries {
	mockDBTX := &MockDBTX{}
	return &MockQueries{Queries: repo.New(mockDBTX)}
}

func (m *MockQueries) GetCellById(ctx context.Context, id int) (repo.Cell, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(repo.Cell), args.Error(1)
}

func (m *MockQueries) GetChildrenCellsByParentId(ctx context.Context, parentID int) ([]repo.Cell, error) {
	args := m.Called(ctx, parentID)
	if _, ok := args.Get(0).([]repo.Cell); !ok {
		return nil, args.Error(1)
	}
	
	return args.Get(0).([]repo.Cell), args.Error(1)
}

func TestGetChildrenCellsByParentID(t *testing.T) {
	mockQueries := NewMockQueries()
	cellService := NewCellService(mockQueries)

	ctx := context.Background()
	userID := 1
	parentID := 10

	parentCell := repo.Cell{
		ID:      parentID,
		OwnerID: userID,
	}

	color1 := "red"
	goal1 := "Goal1"
	color2 := "blue"
	goal2 := "Goal2"
	childrenCells := []repo.Cell{
		{ID: 11, Color: &color1, Goal: &goal1, IsCompleted: false},
		{ID: 12, Color: &color2, Goal: &goal2, IsCompleted: true},
	}

	// Mocking the GetCellById call
	mockQueries.On("GetCellById", ctx, parentID).Return(parentCell, nil)

	// Mocking the GetChildrenCellsByParentId call
	mockQueries.On("GetChildrenCellsByParentId", ctx, parentID).Return(childrenCells, nil)

	cells, err := cellService.GetChildrenCellsByParentID(ctx, userID, parentID)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(cells))
	assert.Equal(t, cells[0].ID, childrenCells[0].ID)
	assert.Equal(t, cells[1].ID, childrenCells[1].ID)

	mockQueries.AssertExpectations(t)
}

func TestGetChildrenCellsByParentID_NotAuthorized(t *testing.T) {
	mockQueries := NewMockQueries()
	cellService := NewCellService(mockQueries)

	ctx := context.Background()
	userID := 1
	parentID := 10

	parentCell := repo.Cell{
		ID:      parentID,
		OwnerID: 2,
	}

	mockQueries.On("GetCellById", ctx, parentID).Return(parentCell, nil)

	cells, err := cellService.GetChildrenCellsByParentID(ctx, userID, parentID)
	assert.Error(t, err)
	assert.Nil(t, cells)
	assert.Equal(t, "not authorized", err.Error())

	mockQueries.AssertExpectations(t)
}

func TestGetChildrenCellsByParentID_ErrorFetchingParent(t *testing.T) {
	mockQueries := NewMockQueries()
	cellService := NewCellService(mockQueries)

	ctx := context.Background()
	userID := 1
	parentID := 10

	mockQueries.On("GetCellById", ctx, parentID).Return(repo.Cell{}, errors.New("db error"))

	cells, err := cellService.GetChildrenCellsByParentID(ctx, userID, parentID)
	assert.Error(t, err)
	assert.Nil(t, cells)
	assert.Equal(t, "db error", err.Error())

	mockQueries.AssertExpectations(t)
}

func TestGetChildrenCellsByParentID_ErrorFetchingChildren(t *testing.T) {
	mockQueries := NewMockQueries()
	cellService := NewCellService(mockQueries)

	ctx := context.Background()
	userID := 1
	parentID := 10

	parentCell := repo.Cell{
		ID:      parentID,
		OwnerID: userID,
	}

	mockQueries.On("GetCellById", ctx, parentID).Return(parentCell, nil)
	mockQueries.On("GetChildrenCellsByParentId", ctx, parentID).Return(nil, errors.New("db error"))

	cells, err := cellService.GetChildrenCellsByParentID(ctx, userID, parentID)
	assert.Error(t, err)
	assert.Nil(t, cells)
	assert.Equal(t, "db error", err.Error())

	mockQueries.AssertExpectations(t)
}
