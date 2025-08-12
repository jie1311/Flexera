package main

import (
	"os"
	"testing"
)

func TestCalculateCopiesFromCsv(t *testing.T) {
	// Test cases
	tests := []struct {
		name          string
		csvContent    string
		applicationID string
		wantCopies    int
		wantErr       bool
	}{
		{
			name: "valid single user single computer",
			csvContent: `ComputerID,UserID,ApplicationID,ComputerType,Comment
C1,U1,APP1,Desktop,test`,
			applicationID: "APP1",
			wantCopies:    1,
			wantErr:       false,
		},
		{
			name: "valid single user multiple computers",
			csvContent: `ComputerID,UserID,ApplicationID,ComputerType,Comment
C1,U1,APP1,Desktop,test
C2,U1,APP1,Laptop,test`,
			applicationID: "APP1",
			wantCopies:    2,
			wantErr:       false,
		},
		{
			name: "invalid format",
			csvContent: `Wrong,Format
C1,U1,APP1`,
			applicationID: "APP1",
			wantCopies:    0,
			wantErr:       true,
		},
		{
			name: "multiple applications",
			csvContent: `ComputerID,UserID,ApplicationID,ComputerType,Comment
C1,U1,APP1,Desktop,test
C2,U1,APP2,Desktop,test`,
			applicationID: "APP1",
			wantCopies:    1,
			wantErr:       false,
		},
		{
			name: "duplicate computer entries",
			csvContent: `ComputerID,UserID,ApplicationID,ComputerType,Comment
C1,U1,APP1,Desktop,test
C1,U1,APP1,Desktop,test`,
			applicationID: "APP1",
			wantCopies:    1,
			wantErr:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temporary file
			tmpfile, err := os.CreateTemp("", "test*.csv")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tmpfile.Name())

			// Write test content
			if _, err := tmpfile.Write([]byte(tt.csvContent)); err != nil {
				t.Fatal(err)
			}
			if err := tmpfile.Close(); err != nil {
				t.Fatal(err)
			}

			// Run test
			gotCopies, err := CalculateCopiesFromCsv(tmpfile.Name(), tt.applicationID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateCopiesFromCsv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCopies != tt.wantCopies {
				t.Errorf("CalculateCopiesFromCsv() = %v, want %v", gotCopies, tt.wantCopies)
			}
		})
	}
}

func TestCheckFormat(t *testing.T) {
	tests := []struct {
		name        string
		firstRecord []string
		wantErr     bool
	}{
		{
			name:        "valid format",
			firstRecord: []string{"ComputerID", "UserID", "ApplicationID", "ComputerType", "Comment"},
			wantErr:     false,
		},
		{
			name:        "invalid format - wrong fields",
			firstRecord: []string{"Wrong", "Format", "Fields"},
			wantErr:     true,
		},
		{
			name:        "invalid format - empty",
			firstRecord: []string{},
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkFormat(tt.firstRecord)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkFormat() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHasUser(t *testing.T) {
	users := map[string]User{
		"U1": {ID: "U1"},
		"U2": {ID: "U2"},
	}

	tests := []struct {
		name     string
		users    map[string]User
		id       string
		wantBool bool
	}{
		{
			name:     "existing user",
			users:    users,
			id:       "U1",
			wantBool: true,
		},
		{
			name:     "non-existing user",
			users:    users,
			id:       "U3",
			wantBool: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotBool := hasUser(tt.users, tt.id); gotBool != tt.wantBool {
				t.Errorf("hasUser() = %v, want %v", gotBool, tt.wantBool)
			}
		})
	}
}
