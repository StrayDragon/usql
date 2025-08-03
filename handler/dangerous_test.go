package handler

import (
	"testing"
)

// TestIsDangerousCommand tests the isDangerousCommand function
func TestIsDangerousCommand(t *testing.T) {
	tests := []struct {
		name     string
		prefix   string
		sqlstr   string
		expected bool
	}{
		{
			name:     "DELETE command",
			prefix:   "",
			sqlstr:   "DELETE FROM users WHERE id = 1",
			expected: true,
		},
		{
			name:     "DROP command",
			prefix:   "",
			sqlstr:   "DROP TABLE users",
			expected: true,
		},
		{
			name:     "UPDATE command",
			prefix:   "",
			sqlstr:   "UPDATE users SET name = 'test' WHERE id = 1",
			expected: true,
		},
		{
			name:     "ALTER command",
			prefix:   "",
			sqlstr:   "ALTER TABLE users ADD COLUMN email VARCHAR(255)",
			expected: true,
		},
		{
			name:     "TRUNCATE command",
			prefix:   "",
			sqlstr:   "TRUNCATE TABLE users",
			expected: true,
		},
		{
			name:     "SELECT command (safe)",
			prefix:   "",
			sqlstr:   "SELECT * FROM users",
			expected: false,
		},
		{
			name:     "INSERT command (safe)",
			prefix:   "",
			sqlstr:   "INSERT INTO users (name) VALUES ('test')",
			expected: false,
		},
		{
			name:     "Case insensitive DELETE",
			prefix:   "",
			sqlstr:   "delete from users where id = 1",
			expected: true,
		},
		{
			name:     "DELETE with prefix",
			prefix:   "EXPLAIN",
			sqlstr:   "DELETE FROM users WHERE id = 1",
			expected: true,
		},
		{
			name:     "Empty command",
			prefix:   "",
			sqlstr:   "",
			expected: false,
		},
		{
			name:     "CREATE INDEX command",
			prefix:   "",
			sqlstr:   "CREATE INDEX idx_name ON users (name)",
			expected: true,
		},
		{
			name:     "CREATE UNIQUE INDEX command",
			prefix:   "",
			sqlstr:   "CREATE UNIQUE INDEX idx_email ON users (email)",
			expected: true,
		},
		{
			name:     "GRANT command",
			prefix:   "",
			sqlstr:   "GRANT SELECT ON users TO user1",
			expected: true,
		},
		{
			name:     "REVOKE command",
			prefix:   "",
			sqlstr:   "REVOKE SELECT ON users FROM user1",
			expected: true,
		},
		{
			name:     "RENAME command",
			prefix:   "",
			sqlstr:   "RENAME TABLE old_table TO new_table",
			expected: true,
		},
		{
			name:     "CREATE TABLE command (safe)",
			prefix:   "",
			sqlstr:   "CREATE TABLE users (id INT, name VARCHAR(50))",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isDangerousCommand(tt.prefix, tt.sqlstr)
			if result != tt.expected {
				t.Errorf("isDangerousCommand(%q, %q) = %v, want %v", tt.prefix, tt.sqlstr, result, tt.expected)
			}
		})
	}
}
