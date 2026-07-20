package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestCoursecheckReportsSuccessfulValidationInRussian(t *testing.T) {
	root := t.TempDir()
	writeValidCourse(t, root)

	output, err := runCoursecheck(t, root)
	if err != nil {
		t.Fatalf("coursecheck failed: %v\n%s", err, output)
	}
	if got, want := string(output), "Проверка структуры курса пройдена.\n"; got != want {
		t.Fatalf("coursecheck output = %q, want %q", got, want)
	}
}

func TestCoursecheckReportsDiagnosticsInRussian(t *testing.T) {
	root := t.TempDir()

	output, err := runCoursecheck(t, root)
	if err == nil {
		t.Fatal("coursecheck succeeded for an invalid course")
	}
	if want := "отсутствует обязательный корневой файл: README.md"; !strings.Contains(string(output), want) {
		t.Fatalf("coursecheck output = %q, want diagnostic %q", output, want)
	}
}

func TestCoursecheckReportsUsageInRussian(t *testing.T) {
	command := exec.Command("go", "run", ".")
	command.Dir = "."
	command.Env = append(os.Environ(), "GOCACHE="+filepath.Join(t.TempDir(), "go-build"))

	output, err := command.CombinedOutput()
	if err == nil {
		t.Fatal("coursecheck succeeded without a course root")
	}
	if want := "Использование: coursecheck <корень-курса>"; !strings.Contains(string(output), want) {
		t.Fatalf("coursecheck output = %q, want usage %q", output, want)
	}
}

func runCoursecheck(t *testing.T, root string) ([]byte, error) {
	t.Helper()
	command := exec.Command("go", "run", ".", root)
	command.Dir = "."
	command.Env = append(os.Environ(), "GOCACHE="+filepath.Join(t.TempDir(), "go-build"))
	return command.CombinedOutput()
}

func writeValidCourse(t *testing.T, root string) {
	t.Helper()
	files := map[string]string{
		"README.md":                "# Курс\n",
		"AGENTS.md":                "# Наставник\n",
		"tool-versions.md":         "# Версии\n",
		"go.mod":                   "module example.com/course\n\ngo 1.24.0\n",
		"progress/PROFILE.md":      "## Опыт\n## Цели\n## Ритм\n## Тон общения\n",
		"progress/STATUS.md":       "## Текущий этап\n## Ближайшее задание\n## Честный пробел\n## Следующий шаг\n",
		"progress/EVIDENCE.md":     "## Правила\n## Журнал\n",
		"progress/EXAMS.md":        "## Журнал экзаменов\n",
		"progress/REVIEW_QUEUE.md": "## К повторению\n",
		"progress/README.md":       "## Правила обновления\n",
	}
	for name, content := range files {
		path := filepath.Join(root, filepath.FromSlash(name))
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatalf("MkdirAll(%q): %v", path, err)
		}
		if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
			t.Fatalf("WriteFile(%q): %v", path, err)
		}
	}
}
