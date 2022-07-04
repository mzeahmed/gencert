package cert

import "testing"

func TestValidCertData(t *testing.T) {
	c, err := New("Golang", "Bob", "2018-05-31")
	if err != nil {
		t.Errorf("Cert data should be valid. err=%v", err)
	}
	if c == nil {
		t.Errorf("Cert should be a valid reference. got=nil")
	}

	if c.Course != "GOLANG COURSE" {
		t.Errorf("Course name is not valid. expected='GOLANG COOURSE', got=%v", c.Course)
	}
}

func TestCourseEmptyValue(t *testing.T) {
	_, err := New("", "Bob", "2018-05-31")
	if err == nil {
		t.Errorf("Error should be returned on an empty course")
	}
}

func TestCourseToLong(t *testing.T) {
	course := "azertyuiopqsdfghjklmmwxcvbnazertyuiopqsdfghjklmmwxcvbnazertyuiopqsdfghj"
	_, err := New(course, "Bob", "2018-05-31")
	if err == nil {
		t.Errorf("Error should be returned on a too long course name (course=%s)", course)
	}
}

func TestNameEmptyValue(t *testing.T) {
	_, err := New("Golang", "", "2018-05-31")
	if err == nil {
		t.Errorf("Error should be returned on an empty name")
	}
}

func TestNameToLong(t *testing.T) {
	name := "azertyuiopqsdfghjklmmwxcvbnazertyuiopqsdfghjklmmwxcvbnazertyuiopqsdfghj"
	_, err := New("Golang", name, "2018-05-31")
	if err == nil {
		t.Errorf("Error should be returned on a too long course name (name=%s)", name)
	}
}
