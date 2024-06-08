package mapache

import "testing"

func TestUser_TableName(t *testing.T) {
	u := User{}
	if u.TableName() != "user" {
		t.Error("Expected user, got", u.TableName())
	}
}

func TestUser_HasRole(t *testing.T) {
	t.Run("Test HasRole", func(t *testing.T) {
		u := User{
			Roles: []string{"admin", "user"},
		}
		if !u.HasRole("admin") {
			t.Error("Expected true, got", u.HasRole("admin"))
		}
		if !u.HasRole("user") {
			t.Error("Expected true, got", u.HasRole("user"))
		}
		if u.HasRole("guest") {
			t.Error("Expected false, got", u.HasRole("guest"))
		}
	})
}

func TestUserRole_TableName(t *testing.T) {
	ur := UserRole{}
	if ur.TableName() != "user_role" {
		t.Error("Expected user_role, got", ur.TableName())
	}
}
