// Package coursecheck validates the repository's learner-facing contract.
package coursecheck

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var requiredRootFiles = []string{
	"AGENTS.md",
	"README.md",
	"go.mod",
	"tool-versions.md",
	"progress/PROFILE.md",
	"progress/STATUS.md",
	"progress/EVIDENCE.md",
	"progress/EXAMS.md",
	"progress/REVIEW_QUEUE.md",
	"progress/README.md",
}

var requiredProgressHeadings = map[string][]string{
	"progress/PROFILE.md": {
		"## Опыт",
		"## Цели",
		"## Ритм",
		"## Тон общения",
	},
	"progress/STATUS.md": {
		"## Текущий этап",
		"## Ближайшее задание",
		"## Честный пробел",
		"## Следующий шаг",
	},
	"progress/EVIDENCE.md": {
		"## Правила",
		"## Журнал",
	},
	"progress/EXAMS.md": {
		"## Журнал экзаменов",
	},
	"progress/REVIEW_QUEUE.md": {
		"## К повторению",
	},
	"progress/README.md": {
		"## Правила обновления",
	},
}

// Check returns all contract violations in stable lexicographic order.
// It only reads files and never changes the course directory.
func Check(root string) []error {
	var diagnostics []error

	for _, name := range requiredRootFiles {
		if !isRegularFile(filepath.Join(root, filepath.FromSlash(name))) {
			diagnostics = append(diagnostics, fmt.Errorf("отсутствует обязательный корневой файл: %s", name))
		}
	}

	for name, headings := range requiredProgressHeadings {
		content, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(name)))
		if err != nil {
			continue
		}
		for _, heading := range headings {
			if !hasHeading(string(content), heading) {
				diagnostics = append(diagnostics, fmt.Errorf("%s: отсутствует обязательный заголовок: %s", name, heading))
			}
		}
	}

	walkErr := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			diagnostics = append(diagnostics, fmt.Errorf("%s: не удалось прочитать: %v", relativePath(root, path), err))
			return nil
		}
		if entry.IsDir() && entry.Name() == ".git" {
			return filepath.SkipDir
		}
		if entry.IsDir() {
			if isLabDirectory(root, path) {
				validateLabDirectory(&diagnostics, root, path)
			}
			return nil
		}

		relative := relativePath(root, path)
		if containsStarterUnderSolutions(relative) {
			diagnostics = append(diagnostics, fmt.Errorf("%s: код starter не должен храниться в директории solutions", relative))
		}
		if strings.EqualFold(filepath.Ext(path), ".md") {
			diagnostics = append(diagnostics, checkMarkdownLinks(root, path)...)
		}
		return nil
	})
	if walkErr != nil {
		diagnostics = append(diagnostics, fmt.Errorf("не удалось обойти корень курса: %v", walkErr))
	}

	sort.Slice(diagnostics, func(i, j int) bool {
		return diagnostics[i].Error() < diagnostics[j].Error()
	})
	return diagnostics
}

func isRegularFile(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.Mode().IsRegular()
}

func hasHeading(content, heading string) bool {
	for _, line := range strings.Split(content, "\n") {
		if strings.TrimSpace(line) == heading {
			return true
		}
	}
	return false
}

func isLabDirectory(root, path string) bool {
	relative := filepath.ToSlash(relativePath(root, path))
	parts := strings.Split(relative, "/")
	return len(parts) == 3 && parts[0] == "labs" && !containsPart(parts, "solutions")
}

func validateLabDirectory(diagnostics *[]error, root, path string) {
	relative := relativePath(root, path)
	if !isRegularFile(filepath.Join(path, "LAB.md")) {
		*diagnostics = append(*diagnostics, fmt.Errorf("%s: отсутствует LAB.md", relative))
	}
	if !isRegularFile(filepath.Join(path, "CHECK.md")) {
		*diagnostics = append(*diagnostics, fmt.Errorf("%s: отсутствует CHECK.md", relative))
	}
	if !hasStarterMaterial(filepath.Join(path, "starter")) {
		*diagnostics = append(*diagnostics, fmt.Errorf("%s: отсутствует код в starter/", relative))
	}
}

func hasStarterMaterial(path string) bool {
	found := false
	_ = filepath.WalkDir(path, func(_ string, entry fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if !entry.IsDir() && entry.Type().IsRegular() {
			found = true
			return filepath.SkipAll
		}
		return nil
	})
	return found
}

func containsStarterUnderSolutions(relative string) bool {
	parts := strings.Split(filepath.ToSlash(relative), "/")
	for index, part := range parts {
		if part != "solutions" {
			continue
		}
		return containsPart(parts[index+1:], "starter")
	}
	return false
}

func containsPart(parts []string, want string) bool {
	for _, part := range parts {
		if part == want {
			return true
		}
	}
	return false
}

func checkMarkdownLinks(root, path string) []error {
	content, err := os.ReadFile(path)
	if err != nil {
		return []error{fmt.Errorf("%s: не удалось прочитать Markdown: %v", relativePath(root, path), err)}
	}

	var diagnostics []error
	for _, target := range markdownLinkTargets(string(content)) {
		if !isRelativeTarget(target) {
			continue
		}
		linkPath := strings.SplitN(strings.SplitN(target, "#", 2)[0], "?", 2)[0]
		if linkPath == "" {
			continue
		}
		resolved := filepath.Clean(filepath.Join(filepath.Dir(path), filepath.FromSlash(linkPath)))
		relative, err := filepath.Rel(root, resolved)
		if err != nil || relative == ".." || strings.HasPrefix(relative, ".."+string(filepath.Separator)) {
			diagnostics = append(diagnostics, fmt.Errorf("%s: относительная ссылка выходит за пределы репозитория: %s", relativePath(root, path), target))
			continue
		}
		if _, err := os.Stat(resolved); err != nil {
			diagnostics = append(diagnostics, fmt.Errorf("%s: неработающая относительная ссылка: %s", relativePath(root, path), target))
		}
	}
	return diagnostics
}

func markdownLinkTargets(content string) []string {
	var targets []string
	for offset := 0; offset < len(content); {
		open := strings.Index(content[offset:], "](")
		if open == -1 {
			break
		}
		open += offset + len("](")
		close := strings.IndexByte(content[open:], ')')
		if close == -1 {
			break
		}
		close += open
		target := strings.TrimSpace(content[open:close])
		if firstSpace := strings.IndexAny(target, " \t"); firstSpace != -1 {
			target = target[:firstSpace]
		}
		targets = append(targets, strings.Trim(target, "<>"))
		offset = close + 1
	}
	return targets
}

func isRelativeTarget(target string) bool {
	lower := strings.ToLower(target)
	return target != "" && !strings.HasPrefix(target, "#") && !strings.HasPrefix(target, "/") &&
		!strings.Contains(lower, "://") && !strings.HasPrefix(lower, "mailto:") && !strings.HasPrefix(lower, "tel:")
}

func relativePath(root, path string) string {
	relative, err := filepath.Rel(root, path)
	if err != nil {
		return path
	}
	return filepath.ToSlash(relative)
}
