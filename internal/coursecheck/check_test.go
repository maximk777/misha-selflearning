package coursecheck

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCheck(t *testing.T) {
	tests := []struct {
		name    string
		mutate  func(t *testing.T, root string)
		wantErr string
	}{
		{
			name:    "accepts the repository contract before course content exists",
			wantErr: "",
		},
		{
			name: "reports missing root file",
			mutate: func(t *testing.T, root string) {
				t.Helper()
				mustRemove(t, root, "README.md")
			},
			wantErr: "отсутствует обязательный корневой файл: README.md",
		},
		{
			name: "reports missing required progress heading",
			mutate: func(t *testing.T, root string) {
				t.Helper()
				mustWrite(t, root, "progress/STATUS.md", "# Статус\n")
			},
			wantErr: "progress/STATUS.md: отсутствует обязательный заголовок: ## Текущий этап",
		},
		{
			name: "reports broken relative markdown link",
			mutate: func(t *testing.T, root string) {
				t.Helper()
				mustWrite(t, root, "README.md", "[Нерабочая ссылка](missing.md)\n")
			},
			wantErr: "README.md: неработающая относительная ссылка: missing.md",
		},
		{
			name: "reports lab without LAB.md",
			mutate: func(t *testing.T, root string) {
				t.Helper()
				mustWrite(t, root, "labs/go/01-syntax/CHECK.md", "# Проверка\n")
			},
			wantErr: "labs/go/01-syntax: отсутствует LAB.md",
		},
		{
			name: "reports lab without CHECK.md",
			mutate: func(t *testing.T, root string) {
				t.Helper()
				mustWrite(t, root, "labs/go/01-syntax/LAB.md", "# Лабораторная\n")
				mustWrite(t, root, "labs/go/01-syntax/starter/main.go", "package main\n")
			},
			wantErr: "labs/go/01-syntax: отсутствует CHECK.md",
		},
		{
			name: "reports lab without starter code",
			mutate: func(t *testing.T, root string) {
				t.Helper()
				mustWrite(t, root, "labs/go/01-syntax/LAB.md", "# Лабораторная\n")
				mustWrite(t, root, "labs/go/01-syntax/CHECK.md", "# Проверка\n")
			},
			wantErr: "labs/go/01-syntax: отсутствует код в starter/",
		},
		{
			name: "accepts normative labs go 01-syntax without treating starter as a lab",
			mutate: func(t *testing.T, root string) {
				t.Helper()
				mustWrite(t, root, "labs/go/01-syntax/LAB.md", "# Лабораторная\n")
				mustWrite(t, root, "labs/go/01-syntax/CHECK.md", "# Проверка\n")
				mustWrite(t, root, "labs/go/01-syntax/starter/main.go", "package main\n")
			},
			wantErr: "",
		},
		{
			name: "reports starter code under solutions",
			mutate: func(t *testing.T, root string) {
				t.Helper()
				mustWrite(t, root, "labs/go/solutions/01-syntax/starter/main.go", "package main\n")
			},
			wantErr: "labs/go/solutions/01-syntax/starter/main.go: код starter не должен храниться в директории solutions",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := t.TempDir()
			writeValidCourse(t, root)
			if tt.mutate != nil {
				tt.mutate(t, root)
			}

			errs := Check(root)
			if tt.wantErr == "" {
				if len(errs) != 0 {
					t.Fatalf("Check() errors = %v, want none", errs)
				}
				return
			}

			if !containsError(errs, tt.wantErr) {
				t.Fatalf("Check() errors = %v, want diagnostic %q", errs, tt.wantErr)
			}
		})
	}
}

func TestCheckSortsDiagnostics(t *testing.T) {
	root := t.TempDir()
	mustWrite(t, root, "README.md", "[broken](missing.md)\n")

	errs := Check(root)
	for i := 1; i < len(errs); i++ {
		if errs[i-1].Error() > errs[i].Error() {
			t.Fatalf("diagnostics are not sorted: %v", errs)
		}
	}
}

func writeValidCourse(t *testing.T, root string) {
	t.Helper()
	mustWrite(t, root, "README.md", "# Курс\n")
	mustWrite(t, root, "AGENTS.md", "# Наставник\n")
	mustWrite(t, root, "tool-versions.md", "# Версии\n")
	mustWrite(t, root, "go.mod", "module example.com/course\n\ngo 1.24.0\n")
	mustWrite(t, root, "progress/PROFILE.md", "# Профиль\n\n## Опыт\n\n## Цели\n\n## Ритм\n\n## Тон общения\n")
	mustWrite(t, root, "progress/STATUS.md", "# Статус\n\n## Текущий этап\n\n## Ближайшее задание\n\n## Честный пробел\n\n## Следующий шаг\n")
	mustWrite(t, root, "progress/EVIDENCE.md", "# Доказательства\n\n## Правила\n\n## Журнал\n")
	mustWrite(t, root, "progress/EXAMS.md", "# Экзамены\n\n## Журнал экзаменов\n")
	mustWrite(t, root, "progress/REVIEW_QUEUE.md", "# Очередь повторения\n\n## К повторению\n")
	mustWrite(t, root, "progress/README.md", "# Прогресс\n\n## Правила обновления\n")
}

func mustWrite(t *testing.T, root, name, content string) {
	t.Helper()
	path := filepath.Join(root, filepath.FromSlash(name))
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("MkdirAll(%q): %v", path, err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("WriteFile(%q): %v", path, err)
	}
}

func mustRemove(t *testing.T, root, name string) {
	t.Helper()
	if err := os.Remove(filepath.Join(root, filepath.FromSlash(name))); err != nil {
		t.Fatalf("Remove(%q): %v", name, err)
	}
}

func containsError(errs []error, want string) bool {
	for _, err := range errs {
		if err.Error() == want {
			return true
		}
	}
	return false
}
