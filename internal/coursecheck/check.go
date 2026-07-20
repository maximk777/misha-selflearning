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
			diagnostics = append(diagnostics, fmt.Errorf("missing required root file: %s", name))
		}
	}

	for name, headings := range requiredProgressHeadings {
		content, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(name)))
		if err != nil {
			continue
		}
		for _, heading := range headings {
			if !hasHeading(string(content), heading) {
				diagnostics = append(diagnostics, fmt.Errorf("%s: missing required heading: %s", name, heading))
			}
		}
	}

	walkErr := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			diagnostics = append(diagnostics, fmt.Errorf("%s: cannot inspect: %v", relativePath(root, path), err))
			return nil
		}
		if entry.IsDir() && entry.Name() == ".git" {
			return filepath.SkipDir
		}
		if entry.IsDir() {
			if isLabDirectory(root, path) && !isRegularFile(filepath.Join(path, "LAB.md")) {
				diagnostics = append(diagnostics, fmt.Errorf("%s: missing LAB.md", relativePath(root, path)))
			}
			return nil
		}

		relative := relativePath(root, path)
		if containsStarterUnderSolutions(relative) {
			diagnostics = append(diagnostics, fmt.Errorf("%s: starter code must not be stored under a solutions directory", relative))
		}
		if strings.EqualFold(filepath.Ext(path), ".md") {
			diagnostics = append(diagnostics, checkMarkdownLinks(root, path)...)
		}
		return nil
	})
	if walkErr != nil {
		diagnostics = append(diagnostics, fmt.Errorf("cannot walk course root: %v", walkErr))
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
	if len(parts) < 3 || parts[0] != "labs" || containsPart(parts, "solutions") {
		return false
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return false
	}
	for _, entry := range entries {
		name := entry.Name()
		if name == "CHECK.md" || name == "LAB.md" || name == "go.mod" || name == "starter" || strings.HasSuffix(name, ".go") {
			return true
		}
	}
	return false
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
		return []error{fmt.Errorf("%s: cannot read Markdown: %v", relativePath(root, path), err)}
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
			diagnostics = append(diagnostics, fmt.Errorf("%s: relative link escapes repository: %s", relativePath(root, path), target))
			continue
		}
		if _, err := os.Stat(resolved); err != nil {
			diagnostics = append(diagnostics, fmt.Errorf("%s: broken relative link: %s", relativePath(root, path), target))
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
