package stdio

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTerminalStdIO_Banner(t *testing.T) {
	t.Parallel()

	tmp, err := os.MkdirTemp("", "")
	require.NoError(t, err)
	defer os.RemoveAll(tmp)

	stdin, err := os.Create(filepath.Join(tmp, "stdin"))
	require.NoError(t, err)
	defer stdin.Close()
	stdout, err := os.Create(filepath.Join(tmp, "stdout"))
	require.NoError(t, err)
	defer stdout.Close()
	stderr, err := os.Create(filepath.Join(tmp, "stderr"))
	require.NoError(t, err)
	defer stderr.Close()

	io := newTerminalStdio(stdin, stdout, stderr)
	require.NoError(t, io.Banner("Hello world!"))
	outBytes, err := os.ReadFile(stdout.Name())
	require.NoError(t, err)
	require.Equal(t, "> Hello world!\n", string(outBytes))
}

func TestTerminalStdIO_Error(t *testing.T) {
	t.Parallel()

	tmp, err := os.MkdirTemp("", "")
	require.NoError(t, err)
	defer os.RemoveAll(tmp)

	stdin, err := os.Create(filepath.Join(tmp, "stdin"))
	require.NoError(t, err)
	defer stdin.Close()
	stdout, err := os.Create(filepath.Join(tmp, "stdout"))
	require.NoError(t, err)
	defer stdout.Close()
	stderr, err := os.Create(filepath.Join(tmp, "stderr"))
	require.NoError(t, err)
	defer stderr.Close()

	io := newTerminalStdio(stdin, stdout, stderr)
	require.NoError(t, io.Error("Oh no"))
	errBytes, err := os.ReadFile(stderr.Name())
	require.NoError(t, err)
	require.Equal(t, "X Oh no\n", string(errBytes))
}

func TestTerminalStdIO_GetStringInput(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		defVal string
		lines  []string
	}{
		{
			name: "empty, no default",
		},
		{
			name:   "empty with default",
			defVal: "foo",
		},
		{
			name:  "one line",
			lines: []string{"foo"},
		},
		{
			name:  "multi-line",
			lines: []string{"foo", "bar"},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tmp, err := os.MkdirTemp("", "")
			require.NoError(t, err)
			defer os.RemoveAll(tmp)

			stdin, err := os.Create(filepath.Join(tmp, "stdin"))
			require.NoError(t, err)
			defer stdin.Close()
			for _, l := range tc.lines {
				_, err := stdin.WriteString(fmt.Sprintln(l))
				require.NoError(t, err)
			}
			_, err = stdin.Seek(0, 0)
			require.NoError(t, err)

			stdout, err := os.Create(filepath.Join(tmp, "stdout"))
			require.NoError(t, err)
			defer stdout.Close()
			stderr, err := os.Create(filepath.Join(tmp, "stderr"))
			require.NoError(t, err)
			defer stderr.Close()

			io := newTerminalStdio(stdin, stdout, stderr)

			if len(tc.lines) == 0 {
				val, err := io.GetStringInput("my prompt", tc.defVal)
				if tc.defVal != "" {
					require.NoError(t, err)
					require.Equal(t, tc.defVal, val)

				} else {
					require.Error(t, err)
				}
				return
			}

			for _, l := range tc.lines {
				val, err := io.GetStringInput("my prompt", tc.defVal)
				require.NoError(t, err)
				require.Equal(t, l, val)
			}
		})
	}
}

func TestTerminalStdIO_GetSecret(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		minLength int
		lines     []string
	}{
		{
			name: "empty, no min length",
		},
		{
			name:      "empty with min length",
			minLength: 1,
		},
		{
			name:      "non-empty, too short",
			minLength: 3,
			lines:     []string{"oh"},
		},
		{
			name:  "one line",
			lines: []string{"foo"},
		},
		{
			name:      "multi-line",
			minLength: 3,
			lines:     []string{"foo", "bar"},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tmp, err := os.MkdirTemp("", "")
			require.NoError(t, err)
			defer os.RemoveAll(tmp)

			stdin, err := os.Create(filepath.Join(tmp, "stdin"))
			require.NoError(t, err)
			defer stdin.Close()
			for _, l := range tc.lines {
				_, err := stdin.WriteString(fmt.Sprintln(l))
				require.NoError(t, err)
			}
			_, err = stdin.Seek(0, 0)
			require.NoError(t, err)

			stdout, err := os.Create(filepath.Join(tmp, "stdout"))
			require.NoError(t, err)
			defer stdout.Close()
			stderr, err := os.Create(filepath.Join(tmp, "stderr"))
			require.NoError(t, err)
			defer stderr.Close()

			io := newTerminalStdio(stdin, stdout, stderr)

			if len(tc.lines) == 0 {
				val, err := io.GetSecret("my prompt", tc.minLength)
				if tc.minLength == 0 {
					require.NoError(t, err)
					require.Empty(t, val)
				} else {
					require.Error(t, err)
				}
				return
			}

			for _, l := range tc.lines {
				val, err := io.GetSecret("my prompt", tc.minLength)
				if len(l) < tc.minLength {
					require.Error(t, err)
				} else {
					require.NoError(t, err)
					require.Equal(t, l, val)
				}
			}
		})
	}
}

func TestTerminalStdIO_GetConfirm(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		answer   string
		expected bool
	}{
		{
			name:     "empty answer",
			expected: false,
		},
		{
			name:     "short affirmative",
			answer:   "y",
			expected: true,
		},
		{
			name:   "short negative",
			answer: "n",
		},
		{
			name:     "long affirmative",
			answer:   "yes",
			expected: true,
		},
		{
			name:     "long negative",
			answer:   "no",
			expected: false,
		},
		{
			name:     "nonsense answer",
			answer:   "I dunno",
			expected: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tmp, err := os.MkdirTemp("", "")
			require.NoError(t, err)
			defer os.RemoveAll(tmp)

			stdin, err := os.Create(filepath.Join(tmp, "stdin"))
			require.NoError(t, err)
			defer stdin.Close()
			_, err = stdin.WriteString(fmt.Sprintln(tc.answer))
			require.NoError(t, err)
			_, err = stdin.Seek(0, 0)
			require.NoError(t, err)

			stdout, err := os.Create(filepath.Join(tmp, "stdout"))
			require.NoError(t, err)
			defer stdout.Close()
			stderr, err := os.Create(filepath.Join(tmp, "stderr"))
			require.NoError(t, err)
			defer stderr.Close()

			io := newTerminalStdio(stdin, stdout, stderr)
			confirmed := io.GetConfirm("?")
			require.Equal(t, tc.expected, confirmed)
		})
	}
}
