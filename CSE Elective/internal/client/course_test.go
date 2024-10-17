package client

import "testing"

func TestClient_GetCourses(t *testing.T) {
	c := NewClient()
	c.SetToken("eeSI6IjE5ZmNhYmEzLTE3MDEtNDNhYS1iODk4LTkxNDBhMjRjOGE3YiJ9.Kjnbq7MIBGvJbaG_aOAkGDQWJJ0RNdKW9gFOhS_256L1dziAmKybHtixOYrdVEkrn1PdxlMEpjb5dP0vcd7ydA")
	courses, err := c.GetCourses()
	if err != nil {
		t.Errorf("GetCourses() failed: %v", err)
		return
	}

	t.Logf("GetCourses() success: courses=%+v", courses)
}

func TestClient_SelectCourse(t *testing.T) {
	c := NewClient()
	c.SetToken("eyJhbGciOiIjE5ZmNhYmEzLTE3MDEtNDNhYS1iODk4LTkxNDBhMjRjOGE3YiJ9.Kjnbq7MIBGvJbaG_aOAkGDQWJJ0RNdKW9gFOhS_256L1dziAmKybHtixOYrdVEkrn1PdxlMEpjb5dP0vcd7ydA")
	course := Course{
		CourseId: 1141,
	}
	err := c.SelectCourse(course)
	if err != nil {
		t.Errorf("SelectCourse() failed: %v", err)
		return
	}

	t.Logf("SelectCourse() success")
}
